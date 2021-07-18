package main

import (
	"net"
	"sync"
)

type AtomicSocket struct {
	m    sync.Mutex
	sock net.Conn
}

func (s *AtomicSocket) Write(p []byte) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.sock.Write(p)
}

func Dial(t, addr string) (*AtomicSocket, error) {
	sock, err := net.Dial(t, addr)
	if err != nil {
		return nil, err
	}
	return &AtomicSocket{
		m:    sync.Mutex{},
		sock: sock,
	}, nil
}
