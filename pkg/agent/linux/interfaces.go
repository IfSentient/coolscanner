package linux

import (
	"fmt"
	"net"
)

type NetworkCollector struct {
}

func (n *NetworkCollector) Collect() error {
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, i := range ifaces {
		fmt.Println(i.Name)
		fmt.Println(i.Flags)

	}
	return nil
}
