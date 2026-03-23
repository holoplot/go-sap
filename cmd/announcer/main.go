package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/holoplot/go-sap/pkg/logger"
	"github.com/holoplot/go-sap/pkg/sap"
)

func main() {
	destFlag := flag.String("dest", "239.255.255.255", "Multicast group to listen to")
	originFlag := flag.String("origin", "192.168.1.100", "Origin to use in sent packets")
	timeoutFlag := flag.Int("timeout", 0, "Timeout in seconds (0 for disable)")
	sdpFlag := flag.String("sdp", "sdp.txt", "SDP file to use as payload")
	flag.Parse()

	logger.Setup()

	ip := net.ParseIP(*destFlag)

	b, err := os.ReadFile(*sdpFlag)
	if err != nil {
		slog.Error("Failed to read SDP file", "error", err)

		os.Exit(1)
	}

	p := &sap.Packet{
		Type:        sap.MessageTypeAnnouncement,
		IDHash:      0x2342,
		Origin:      net.ParseIP(*originFlag),
		PayloadType: sap.SDPPayloadType,
		Payload:     b,
	}

	ctx := context.Background()

	if *timeoutFlag > 0 {
		var cancel context.CancelFunc

		ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Duration(*timeoutFlag)*time.Second))
		defer cancel()
	}

	slog.Info("Sending announcements periodically",
		"dest", ip,
		"origin", p.Origin,
		"timeout", *timeoutFlag,
		"payload-type", p.PayloadType,
	)

	if err := sap.AnnouncePeriodically(ctx, ip, p); err != nil && !errors.Is(err, context.DeadlineExceeded) {
		slog.Error("Failed to announce periodically", "error", err)
	}
}
