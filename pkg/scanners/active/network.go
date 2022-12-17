package active

import (
	"fmt"
	"net"
)

type PortScanner struct {
}

type HostDiscovery struct {
}

func (h *HostDiscovery) PingSweep(broadcast net.IPNet) ([]net.IP, error) {
	broadcast.String()
	return nil, nil
}

func (p *PortScanner) Check(ip string, port int) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return err
	}
	defer conn.Close()
	// TODO
	return nil
}
