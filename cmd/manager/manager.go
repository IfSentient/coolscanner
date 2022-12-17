package main

import (
	"coolscanner/pkg/protobuf"
	"coolscanner/pkg/transport"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Test
	l := transport.TCPListener{}
	l.Listen(45678, func(data []byte) ([]byte, error) {
		var d protobuf.SystemInfo
		proto.Unmarshal(data, &d)
		fmt.Println("Unmarshaled")
		fmt.Println(d)
		go func() {
			resp, _ := transport.SendTCP("localhost:45679", data)
			fmt.Println(string(resp))
		}()
		return []byte("ok"), nil
	})
}
