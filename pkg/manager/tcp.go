package manager

import "coolscanner/pkg/transport"

type TCPNetwork struct {
	transport.TCPListener
	Port int
}

func (t *TCPNetwork) IsBiDirectional() bool {
	return true
}

func (t *TCPNetwork) Send(string, []byte) error {
	return nil
}

func (t *TCPNetwork) Accept(ch chan<- []byte) error {
	t.Listen(t.Port, func(bytes []byte) ([]byte, error) {
		ch <- bytes
		return nil, nil
	})
	return nil
}
