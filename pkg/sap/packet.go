package sap

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"io"
	"net"
	"strings"

	"github.com/pkg/errors"
)

const (
	SDPPayloadType = "application/sdp"
)

type Version int

const (
	Version1 = Version(1)
)

const (
	compressedFlag      = uint8(1 << 0)
	encryptedFlag       = uint8(1 << 1)
	messageTypeDeletion = uint8(1 << 2)
	addressV6Flag       = uint8(1 << 4)
)

type MessageType int

const (
	MessageTypeAnnouncement = MessageType(iota)
	MessageTypeDeletion
)

type Packet struct {
	Type               MessageType
	IDHash             uint16
	Origin             net.IP
	Encrypted          bool
	Compressed         bool
	PayloadType        string
	AuthenticationData []byte
	Payload            []byte
}

func (p *Packet) UniqueID() string {
	b := append(p.Origin, byte(p.IDHash&0xff), byte(p.IDHash>>8))

	return base64.StdEncoding.EncodeToString(b)
}

var ErrPacketTooShort = errors.New("packet too short")
var ErrAuthenticationDataTooLong = errors.New("authentication data too long")
var ErrPacketInvalidIntegrity = errors.New("packet integrity error")

func (p *Packet) Encode() ([]byte, error) {
	writer := new(bytes.Buffer)

	flags := uint8(0)

	// version field
	flags |= 0b00100000

	if p.Type == MessageTypeDeletion {
		flags |= messageTypeDeletion
	}

	if p.Encrypted {
		flags |= encryptedFlag
	}

	if p.Compressed {
		flags |= compressedFlag
	}

	var origin net.IP

	if ipv4 := p.Origin.To4(); ipv4 != nil {
		origin = ipv4
	} else {
		origin = p.Origin.To16()
		flags |= addressV6Flag
	}

	writer.WriteByte(flags)

	if len(p.AuthenticationData) >= 0x100 {
		return nil, ErrAuthenticationDataTooLong
	}

	writer.WriteByte(uint8(len(p.AuthenticationData)))

	binary.Write(writer, binary.BigEndian, p.IDHash)
	writer.Write(origin)

	writer.Write(p.AuthenticationData)

	payloadBytes := new(bytes.Buffer)
	var payloadWriter io.Writer = payloadBytes

	if p.Compressed {
		payloadWriter = zlib.NewWriter(payloadWriter)
	}

	if len(p.PayloadType) != 0 {
		if _, err := payloadWriter.Write([]byte(p.PayloadType)); err != nil {
			return nil, err
		}

		if _, err := payloadWriter.Write([]byte{0}); err != nil {
			return nil, err
		}
	}

	if _, err := payloadWriter.Write(p.Payload); err != nil {
		return nil, err
	}

	// We need to close the zWriter before we can access its bytes
	if closer, ok := payloadWriter.(io.Closer); ok {
		closer.Close()
	}

	writer.Write(payloadBytes.Bytes())

	return writer.Bytes(), nil
}

func DecodePacket(raw []byte) (*Packet, error) {
	p := &Packet{}

	reader := bytes.NewBuffer(raw)

	var flags uint8
	if err := binary.Read(reader, binary.BigEndian, &flags); err != nil {
		return nil, err
	}

	if flags&messageTypeDeletion == messageTypeDeletion {
		p.Type = MessageTypeDeletion
	} else {
		p.Type = MessageTypeAnnouncement
	}

	if flags&0b11100000 != 0b00100000 {
		return nil, ErrPacketInvalidIntegrity
	}

	p.Compressed = flags&compressedFlag == compressedFlag
	p.Encrypted = flags&encryptedFlag == encryptedFlag

	var authLen uint8
	if err := binary.Read(reader, binary.BigEndian, &authLen); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &p.IDHash); err != nil {
		return nil, err
	}

	var origin net.IP

	if flags&addressV6Flag == addressV6Flag {
		origin = make(net.IP, net.IPv6len)
	} else {
		origin = make(net.IP, net.IPv4len)
	}

	if err := binary.Read(reader, binary.BigEndian, &origin); err != nil {
		return nil, err
	}

	if len(origin) == net.IPv6len {
		p.Origin = origin
	} else {
		p.Origin = net.IPv4(origin[0], origin[1], origin[2], origin[3])
	}

	if authLen > 0 {
		p.AuthenticationData = make([]byte, authLen)
		if err := binary.Read(reader, binary.BigEndian, p.AuthenticationData); err != nil {
			return nil, err
		}
	}

	var payload bytes.Buffer

	if p.Compressed {
		zReader, err := zlib.NewReader(reader)
		if err != nil {
			return nil, err
		}

		if _, err := payload.ReadFrom(zReader); err != nil {
			return nil, err
		}

		zReader.Close()
	} else {
		if _, err := payload.ReadFrom(reader); err != nil {
			return nil, err
		}
	}

	// RFC 2974, section 6:
	// The absence of a payload type field may be noted since the payload
	// section of such a packet will start with an SDP `v=0' field, which is
	// not a legal MIME content type specifier.
	sdpMagic := []byte("v=0")

	if bytes.HasPrefix(payload.Bytes(), sdpMagic) {
		p.PayloadType = SDPPayloadType
	} else {
		var err error

		p.PayloadType, err = payload.ReadString(0)
		if err != nil {
			return nil, ErrPacketInvalidIntegrity
		}

		p.PayloadType = strings.TrimRight(p.PayloadType, "\000")
	}

	p.Payload = make([]byte, payload.Len())
	copy(p.Payload, payload.Bytes())

	return p, nil
}
