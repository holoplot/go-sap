package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"time"

	"github.com/holoplot/go-sap/pkg/sap"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	destFlag := flag.String("dest", "239.255.255.255", "Multicast group to listen to")
	originFlag := flag.String("origin", "192.168.1.100", "Origin to use in sent packets")
	timeoutFlag := flag.Int("timeout", 0, "Timeout in seconds (0 for disable)")
	flag.Parse()

	consoleWriter := zerolog.ConsoleWriter{
		Out: colorable.NewColorableStdout(),
	}

	log.Logger = log.Output(consoleWriter)

	ip := net.ParseIP(*destFlag)

	p := &sap.Packet{
		Type:        sap.MessageTypeAnnouncement,
		IDHash:      0x2342,
		Origin:      net.ParseIP(*originFlag),
		PayloadType: sap.SDPPayloadType,
		Payload: []byte(
			`xxx`,
		),
	}

	ctx := context.Background()

	if *timeoutFlag > 0 {
		ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Duration(*timeoutFlag)*time.Second))
	}

	log.Info().
		IPAddr("dest", ip).
		IPAddr("origin", p.Origin).
		Int("timeout", *timeoutFlag).
		Str("payload-type", p.PayloadType).
		Msg("Sending announcements periodically")

	if err := sap.AnnouncePeriodically(ctx, ip, p); err != nil && !errors.Is(err, context.DeadlineExceeded) {
		log.Fatal().Err(err).Msg("Failed to announce periodically")
	}
}
