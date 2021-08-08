package pongrun

import (
	"sync"

	"github.com/gorilla/websocket"
)

type PongSubscriber struct {
	S map[string]*Subcriber
	M sync.Mutex
}

type Subcriber struct {
	Messages chan []byte
	Conns    []*websocket.Conn
	M        sync.Mutex
}

func NewPongSubscriber() *PongSubscriber {
	return &PongSubscriber{
		S: make(map[string]*Subcriber),
	}
}
