package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/holoplot/go-multicast/pkg/multicast"
	"github.com/holoplot/go-sap/pkg/logger"
	"github.com/holoplot/go-sap/pkg/sap"
)

func main() {
	ipFlag := flag.String("dest", "239.255.255.255", "Multicast group to listen to")
	ifaceFlag := flag.String("iface", "", "Interface name to use")
	writeFileFlag := flag.Bool("write-file", false, "Write packets to files in the current directory")
	flag.Parse()

	logger.Setup()

	ifis := []*net.Interface{}

	if *ifaceFlag != "" {
		var err error

		ifi, err := net.InterfaceByName(*ifaceFlag)
		if err != nil {
			slog.Error("No such interface", "error", err)
			os.Exit(1)
		}

		ifis = append(ifis, ifi)
	} else {
		allIfis, err := net.Interfaces()
		if err != nil {
			slog.Error("Failed to get network interfaces", "error", err)
			os.Exit(1)
		}

		for _, ifi := range allIfis {
			if ifi.Flags&net.FlagMulticast == 0 {
				continue
			}

			slog.Debug("Adding interface", "name", ifi.Name)

			ifis = append(ifis, &ifi)
		}
	}

	ml := multicast.NewListener(ifis)

	addr := &net.UDPAddr{
		IP:   net.ParseIP(*ipFlag),
		Port: 9875,
	}

	if _, err := ml.AddConsumer(addr, func(ifi *net.Interface, src net.Addr, payload []byte) {
		p, err := sap.DecodePacket(payload)
		if err != nil {
			slog.Error("Failed to decode packet", "error", err)

			return
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
				return
			}

			if _, err := f.Write(p.Payload); err != nil {
				slog.Error("Failed to write packet to file", "error", err)
			}

			f.Close()
		}
	}); err != nil {
		slog.Error("Failed to add consumer", "error", err)
		os.Exit(1)
	}

	slog.Info("Listening for packets")

	<-context.Background().Done()
}
