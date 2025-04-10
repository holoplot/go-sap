package sap

import (
	"context"
	"math/rand"
	"net"
	"time"
)

type Sender struct {
	conn            *net.UDPConn
	raw             []byte
	intervalSeconds int
}

func NewSender(ip net.IP, p *Packet) (*Sender, error) {
	raw, err := p.Encode()
	if err != nil {
		return nil, err
	}

	udpAddr := &net.UDPAddr{
		IP:   ip,
		Port: sapPort,
	}

	network := "udp4"
	if ip.To4() == nil {
		network = "udp6"
	}

	conn, err := net.DialUDP(network, nil, udpAddr)
	if err != nil {
		return nil, err
	}

	// RFC 2974, section 3.1
	bandwidthLimit := 4000
	intervalSeconds := (8 * len(raw)) / bandwidthLimit

	if intervalSeconds < 300 {
		intervalSeconds = 300
	}

	return &Sender{
		conn:            conn,
		raw:             raw,
		intervalSeconds: intervalSeconds,
	}, nil
}

func (s *Sender) AnnouncePeriodically(ctx context.Context) error {
	defer s.conn.Close()

	for {
		_, err := s.conn.Write(s.raw)
		if err != nil {
			return err
		}

		// RFC 2974, section 3.1
		offsetSeconds := s.intervalSeconds
		offsetSeconds += rand.Intn(s.intervalSeconds*2/3) - s.intervalSeconds/3

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(offsetSeconds) * time.Second):
		}
	}
}
