package sap

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func AnnouncePeriodically(ctx context.Context, ip net.IP, p *Packet) error {
	p.Type = MessageTypeAnnouncement

	raw, err := p.Encode()
	if err != nil {
		return fmt.Errorf("encoding announcement package: %w", err)
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
		return err
	}

	defer conn.Close()

	// RFC 2974, section 3.1
	bandwidthLimit := 4000
	intervalSeconds := (8 * len(raw)) / bandwidthLimit

	if intervalSeconds < 300 {
		intervalSeconds = 300
	}

	for {
		_, err := conn.Write(raw)
		if err != nil {
			return fmt.Errorf("sending announcement package: %w", err)
		}

		// RFC 2974, section 3.1
		offsetSeconds := intervalSeconds
		offsetSeconds += rand.Intn(intervalSeconds*2/3) - intervalSeconds/3

		select {
		case <-ctx.Done():
			p.Type = MessageTypeDeletion

			raw, err := p.Encode()
			if err != nil {
				return fmt.Errorf("encoding deletion package: %w", err)
			}

			_, err = conn.Write(raw)
			if err != nil {
				return fmt.Errorf("sending deletion package: %w", err)
			}

			return ctx.Err()

		case <-time.After(time.Duration(offsetSeconds) * time.Second):
		}
	}
}
