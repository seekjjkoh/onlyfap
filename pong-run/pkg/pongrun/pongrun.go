package pongrun

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

/* Write to redis */
// send data (channelName, data)
// publish to redis
// no need emit to ws (subscribers of redis to handle this)

/* Read from redis */
// subscribers of the channelName
// receive data and emit to all ws

// PongRun is the main type holding connections to redis
// first publish messages to respective redis channels/topics
// then other PongRun instances thats subscribed to the channel
// broadcast the message back to clients through websockets
type PongRun struct {
	Pool       *redis.Pool
	Publisher  *PongPublisher
	Subscriber *PongSubscriber
}

func New(redisHost, redisPort string, redisMaxConn int) *PongRun {
	// fmt.Println("redis", fmt.Sprintf("%s:%s", redisHost, redisPort))
	return &PongRun{
		Pool: &redis.Pool{
			MaxIdle: redisMaxConn,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", fmt.Sprintf("%s:%s", redisHost, redisPort))
			},
		},
		Publisher:  NewPongPublisher(),
		Subscriber: NewPongSubscriber(),
	}
}

func (pr *PongRun) Publish(channelName string, data []byte) {
	if publisher, exist := pr.Publisher.P[channelName]; exist {
		// if there's an existing channel
		// trivial - just push to publisher
		// goBroadcast will pick up and handle this
		publisher <- data
		return
	}

	pr.Publisher.M.Lock()
	// create a publisher channel
	pr.Publisher.P[channelName] = make(chan []byte)
	pr.Publisher.M.Unlock()
	// spawn goroutine listneing to channel
	go pr.goBroadcast(channelName)
	// push message to channel
	pr.Publisher.P[channelName] <- data
}

func (pr *PongRun) goBroadcast(channelName string) error {
	conn := pr.Pool.Get()
	defer conn.Close()
	for m := range pr.Publisher.P[channelName] {
		// publish message to redis
		if err := conn.Send("PUBLISH", channelName, m); err != nil {
			zap.S().Error(err)
			return err
		}
		if err := conn.Flush(); err != nil {
			zap.S().Error(err)
			return err
		}
	}
	return nil
}

func (pr *PongRun) Subscribe(channelName string, wsConn *websocket.Conn) {
	if subs, exist := pr.Subscriber.S[channelName]; exist {
		subs.M.Lock()
		subs.Conns = append(subs.Conns, wsConn)
		subs.M.Unlock()
		return
	}
	pr.Subscriber.M.Lock()
	// create a publisher channel
	pr.Subscriber.S[channelName] = &Subcriber{
		Messages: make(chan []byte),
		Conns:    []*websocket.Conn{wsConn},
		M:        sync.Mutex{},
	}
	pr.Subscriber.M.Unlock()
	go pr.goBroadcastWs(channelName)
	go pr.goSubscribeChannel(channelName)
}

func (pr *PongRun) goBroadcastWs(channelName string) error {
	for msg := range pr.Subscriber.S[channelName].Messages {
		toRemove := make([]*websocket.Conn, 0)
		for _, wsConn := range pr.Subscriber.S[channelName].Conns {
			if err := wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
				zap.S().Error(err)
				toRemove = append(toRemove, wsConn)
			}
		}
		for _, wsConn := range toRemove {
			pr.DeRegisterWs(channelName, wsConn)
		}
	}
	return nil
}

func (pr *PongRun) goSubscribeChannel(channelName string) error {
	conn := pr.Pool.Get()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channelName)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			if _, exist := pr.Subscriber.S[channelName]; !exist {
				return errors.New("channel closed")
			}
			pr.Subscriber.S[channelName].Messages <- v.Data
		case redis.Subscription:
			zap.S().Debug("New subscription")
		case error:
			return v
		}
	}
}

func (pr *PongRun) SetState(key string, state []byte) error {
	conn := pr.Pool.Get()
	defer conn.Close()
	if err := conn.Send("SET", key, state); err != nil {
		return err
	}
	if err := conn.Flush(); err != nil {
		return err
	}
	return nil
}

func (pr *PongRun) GetState(key string) ([]byte, error) {
	conn := pr.Pool.Get()
	defer conn.Close()
	if err := conn.Send("GET", key); err != nil {
		return nil, err
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}
	v, err := conn.Receive()
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, errors.New("empty key")
	}
	return v.([]byte), err
}

func (pr *PongRun) DeRegisterWs(channelName string, ws *websocket.Conn) {
	if _, exist := pr.Subscriber.S[channelName]; !exist {
		return
	}
	wss := pr.Subscriber.S[channelName].Conns
	toRemove := -1
	for i, c := range wss {
		if c == ws {
			toRemove = i
			break
		}
	}
	if toRemove == 0 {
		pr.Subscriber.M.Lock()
		delete(pr.Subscriber.S, channelName)
		pr.Subscriber.M.Unlock()
		return
	}
	if toRemove > 0 {
		copy(pr.Subscriber.S[channelName].Conns[toRemove:], pr.Subscriber.S[channelName].Conns[toRemove+1:])
		pr.Subscriber.S[channelName].Conns = pr.Subscriber.S[channelName].Conns[:len(pr.Subscriber.S[channelName].Conns)-1]
	}
}
