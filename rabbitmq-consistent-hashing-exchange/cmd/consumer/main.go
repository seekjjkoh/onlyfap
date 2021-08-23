package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/streadway/amqp"
)

var (
	qName = flag.String("qName", "queue1", "Queue name")
)

type Message struct {
	HashID string `json:"hash_id"`
}

func main() {
	flag.Parse()
	cs := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(cs)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(*qName, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	counter := make(map[string]int, 10)

	for {
		for d := range msgs {
			var msg Message
			json.Unmarshal(d.Body, &msg)
			counter[msg.HashID]++ // safe, no need lock
			d.Ack(true)
			b, _ := json.Marshal(counter)
			fmt.Println(*qName, string(b))
		}
	}

}
