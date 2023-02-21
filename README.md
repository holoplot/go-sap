[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/holoplot/go-sap) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/holoplot/go-sap/main/LICENSE)

# go-sap

A Go package to decode and encode SAP packets as described in [RFC 2974](https://www.rfc-editor.org/rfc/rfc2974).

# Example

Also see the demo applications in the `cmd/` folder.

## Encode and send

```go
import "github.com/holoplot/go-sap/pkg/sap"

func main() {
	sdp := []byte{
		// your SDP here
	}

	p := &sap.Packet{
		Type:        sap.MessageTypeAnnouncement,
		IDHash:      0x2342,
		Origin:      net.ParseIP("192.168.1.100"),
		PayloadType: sap.SDPPayloadType,
		Payload:     sdp,
	}

	s, err := sap.NewSender(net.ParseIP("239.255.255.255"), p)
	if err != nil {
		panic(err)
	}

	s.AnnouncePeriodically(context.Background())
}
```

## Receive and decode

```go
import "github.com/holoplot/go-sap/pkg/sap"

func main() {
	l, err := sap.NewListener(net.ParseIP("239.255.255.255"), nil)
	if err != nil {
		panic(err)
	}

	for {
		b, err := l.ReadPacketRaw()
		if err != nil {
			panic(err)
		}

		p, err := sap.DecodePacket(b)
		if err != nil {
			panic(err)
		}

		// Use the content of the packet
	}
}
```

# License

MIT
