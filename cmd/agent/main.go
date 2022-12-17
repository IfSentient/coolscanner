package main

import (
	"coolscanner/pkg/agent/linux"
	"coolscanner/pkg/protobuf"
	"coolscanner/pkg/transport"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	c := linux.RPMCollector{}
	p, e := c.Collect()
	if e != nil {
		panic(e)
	}
	for _, pkg := range p {
		fmt.Println(pkg)
	}

	// Convert to protobuf
	list := protobuf.PackageList{
		Packages: make([]*protobuf.Package, 0),
	}
	for _, pkg := range p {
		list.Packages = append(list.Packages, &protobuf.Package{
			Type:    string(pkg.Type),
			Name:    pkg.Name,
			Epoch:   &pkg.Epoch,
			Version: pkg.Version,
			Release: &pkg.Release,
			Arch:    &pkg.Release,
		})
	}

	info := protobuf.SystemInfo{
		Packakges: &list,
	}

	out, err := proto.Marshal(&info)
	if err != nil {
		panic(err)
	}

	resp, err := transport.SendTCP("localhost:45678", out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}
