package pongrun

import "sync"

type PongPublisher struct {
	P map[string](chan []byte)
	M sync.Mutex
}

func NewPongPublisher() *PongPublisher {
	return &PongPublisher{
		P: make(map[string](chan []byte)),
	}
}
