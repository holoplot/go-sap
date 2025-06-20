package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/holoplot/go-sap/pkg/sap"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ipFlag := flag.String("dest", "239.255.255.255", "Multicast group to listen to")
	ifaceFlag := flag.String("iface", "", "Interface name to use")
	writeFileFlag := flag.Bool("write-file", false, "Write packets to files in the current directory")
	flag.Parse()

	consoleWriter := zerolog.ConsoleWriter{
		Out: colorable.NewColorableStdout(),
	}

	log.Logger = log.Output(consoleWriter)

	var ifi *net.Interface

	if *ifaceFlag != "" {
		var err error

		ifi, err = net.InterfaceByName(*ifaceFlag)
		if err != nil {
			log.Fatal().Err(err).Msg("No such interface")
		}
	}

	ip := net.ParseIP(*ipFlag)

	l, err := sap.NewListener(ip, ifi)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

	log.Info().Msg("Listening for packets")

	for {
		b, err := l.ReadPacketRaw()
		if err != nil {
			log.Error().Err(err).Msg("Failed to read raw packet")

			return
		}

		p, err := sap.DecodePacket(b)
		if err != nil {
			log.Error().Err(err).Stack().Msg("Failed to decode packet")

			continue
		}

		log.Info().
			IPAddr("origin", p.Origin).
			Bool("compressed", p.Compressed).
			Bool("is-announcement", p.Type == sap.MessageTypeAnnouncement).
			Str("id-hash", fmt.Sprintf("%04x", p.IDHash)).
			Str("payload-type", p.PayloadType).
			Msg("Packet received")

		if *writeFileFlag {
			filename := fmt.Sprintf("%04x.sdp", p.IDHash)
			f, err := os.Create(filename)
			if err != nil {
				log.Error().Err(err).Msg("Failed to create file")
				continue
			}

			if _, err := f.Write(p.Payload); err != nil {
				log.Error().Err(err).Msg("Failed to write packet to file")
			}

			f.Close()
		}
	}
}
