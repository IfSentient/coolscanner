package transport

import (
	"encoding/binary"
	"fmt"
	"net"
)

type TCPListener struct {
}

func (t *TCPListener) Listen(port int, handler func([]byte) ([]byte, error)) error {
	// Listen for incoming connections
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			// TODO: handle
			continue
		}

		go t.handleConnection(conn, handler)
	}
}

func (t *TCPListener) handleConnection(conn net.Conn, handler func([]byte) ([]byte, error)) error {
	defer conn.Close()
	msg, err := ReadMessage(conn)
	response, err := handler(msg)
	if err != nil {
		// TODO: handle
		return err
	}
	return WriteMessage(conn, response)
}

func SendTCP(dest string, msg []byte) ([]byte, error) {
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	WriteMessage(conn, msg)
	return ReadMessage(conn)
}

func WriteMessage(conn net.Conn, msg []byte) error {
	msgSize := make([]byte, 8) // 64-bit
	binary.BigEndian.PutUint64(msgSize, uint64(len(msg)))
	_, err := conn.Write(msgSize)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(msg)
	return err
}

func ReadMessage(conn net.Conn) ([]byte, error) {
	msgSize := make([]byte, 8)
	_, err := conn.Read(msgSize)
	if err != nil {
		return nil, err
	}
	msg := make([]byte, binary.BigEndian.Uint64(msgSize))
	_, err = conn.Read(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
