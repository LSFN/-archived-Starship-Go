package main

import (
	"net"
	"sync"
)

type Subscription struct {
	sync.RWMutex
	conn  net.Conn
	state *State

	// data
}

func (s *Subscription) Process() error {
	var data []byte
	// process state.

	s.Lock()
	defer s.Unlock()
	return write(s.conn, data)
}
