package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Message struct {
	HashID  string
	Content []byte
}

func main() {
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

	exchange := "header.exchange"
	hashHeader := "my-hash-header"
	rKey := "header.routekey"

	keys := []int{10001, 10003, 20001, 30001, 40001, 40002, 40033, 50001, 60001}

	for _, k := range keys {
		for i := 0; i < 100; i++ {
			id := k / 10000
			msg := Message{
				HashID:  fmt.Sprint(id),
				Content: []byte(fmt.Sprintf(`{"hash_id":"%d"}`, id)),
			}
			headers := amqp.Table{}
			headers[hashHeader] = id
			msgID := uuid.New()
			pb := amqp.Publishing{
				ContentType: "application/json",
				Body:        msg.Content,
				Headers:     headers,
				MessageId:   msgID.String(),
			}
			err = ch.Publish(exchange, rKey, false, false, pb)
			if err != nil {
				panic(err)
			}
		}
	}
}
