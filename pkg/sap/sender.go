package sap

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	bandwidthLimitBits = 4000
	minIntervalDefault = 300 * time.Second
)

type config struct {
	minInterval time.Duration
}

type Option func(o *config)

func WithMinInterval(interval time.Duration) Option {
	return func(c *config) {
		c.minInterval = interval
	}
}

func AnnouncePeriodically(ctx context.Context, ip net.IP, p *Packet, opts ...Option) error {
	c := config{
		minInterval: minIntervalDefault,
	}

	for _, opt := range opts {
		opt(&c)
	}

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
	interval := time.Duration(8*len(raw)/bandwidthLimitBits) * time.Second

	if interval < c.minInterval {
		interval = c.minInterval
	}

	intervalSec := int(interval / time.Second)

	for {
		_, err := conn.Write(raw)
		if err != nil {
			return fmt.Errorf("sending announcement package: %w", err)
		}

		// RFC 2974, section 3.1
		offset := time.Duration(rand.Intn(intervalSec*2/3)-intervalSec/3) * time.Second

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

		case <-time.After(interval + offset):
		}
	}
}
