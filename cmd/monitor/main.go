package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/holoplot/go-sap/pkg/logger"
	"github.com/holoplot/go-sap/pkg/sap"
)

func main() {
	ipFlag := flag.String("dest", "239.255.255.255", "Multicast group to listen to")
	ifaceFlag := flag.String("iface", "", "Interface name to use")
	writeFileFlag := flag.Bool("write-file", false, "Write packets to files in the current directory")
	flag.Parse()

	logger.Setup()

	var ifi *net.Interface

	if *ifaceFlag != "" {
		var err error

		ifi, err = net.InterfaceByName(*ifaceFlag)
		if err != nil {
			slog.Error("No such interface", "error", err)
			os.Exit(1)
		}
	}

	ip := net.ParseIP(*ipFlag)

	l, err := sap.NewListener(ip, ifi)
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}

	slog.Info("Listening for packets")

	for {
		b, err := l.ReadPacketRaw()
		if err != nil {
			slog.Error("Failed to read packet", "error", err)

			return
		}

		p, err := sap.DecodePacket(b)
		if err != nil {
			slog.Error("Failed to decode packet", "error", err)

			continue
		}

		slog.Info("Packet received",
			"origin", p.Origin,
			"compressed", p.Compressed,
			"is-announcement", p.Type == sap.MessageTypeAnnouncement,
			"id-hash", fmt.Sprintf("%04x", p.IDHash),
			"payload-type", p.PayloadType,
		)

		if *writeFileFlag {
			filename := fmt.Sprintf("%04x.sdp", p.IDHash)
			f, err := os.Create(filename)
			if err != nil {
				slog.Error("Failed to create file", "error", err)
				continue
			}

			if _, err := f.Write(p.Payload); err != nil {
				slog.Error("Failed to write packet to file", "error", err)
			}

			f.Close()
		}
	}
}
