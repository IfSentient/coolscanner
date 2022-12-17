package manager

import "fmt"

type Network interface {
	IsBiDirectional() bool
	Send(string, []byte) error
	Listen(chan<- []byte) error
}

type Manager struct {
	Network  Network
	PoolSize int
}

func (m *Manager) Run(stopCh <-chan struct{}) error {
	msgChan := make(chan []byte, m.PoolSize)
	errChan := make(chan error)
	go func() {
		err := m.Network.Listen(msgChan)
		errChan <- err
	}()
	select {
	case msg := <-msgChan:
		fmt.Println(msg)
	case <-stopCh:
		return nil
	case err := <-errChan:
		return err
	}

	return fmt.Errorf("unexpected select exit")
}
