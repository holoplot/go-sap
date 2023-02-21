package sap

import (
	"net"
)

type Listener struct {
	conn *net.UDPConn
}

func NewListener(ip net.IP, ifi *net.Interface) (*Listener, error) {
	udpAddr := &net.UDPAddr{
		IP:   ip,
		Port: sapPort,
	}

	conn, err := net.ListenMulticastUDP("udp4", ifi, udpAddr)
	if err != nil {
		return nil, err
	}

	if err := conn.SetReadBuffer(maxDatagramSize); err != nil {
		return nil, err
	}

	return &Listener{
		conn: conn,
	}, nil
}

func (l *Listener) Close() {
	l.conn.Close()
}

func (l *Listener) ReadPacketRaw() ([]byte, error) {
	buf := make([]byte, maxDatagramSize)

	n, _, err := l.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}
