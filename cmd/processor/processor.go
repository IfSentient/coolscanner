package main

import (
	"coolscanner/pkg/processor"
	"coolscanner/pkg/protobuf"
	"coolscanner/pkg/scanners"
	"coolscanner/pkg/transport"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := processor.New()
	p.AddScanner("foo", &scanners.SimpleMatcher{})
	l := transport.TCPListener{}
	l.Listen(45679, func(data []byte) ([]byte, error) {
		var d protobuf.SystemInfo
		proto.Unmarshal(data, &d)
		problems, err := p.Process(&d)
		if err != nil {
			panic(err)
		}
		return []byte(fmt.Sprintf("Problems: %d", len(problems))), nil
	})
}
