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

	exchange := "hashheader.exchange"
	hashHeader := "hh"

	keys := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "a", "b", "c"}

	for _, k := range keys {
		for i := 0; i < 100; i++ {
			id := k
			msg := Message{
				HashID:  fmt.Sprint(id),
				Content: []byte(fmt.Sprintf(`{"hash_id":"%s"}`, id)),
			}
			headers := amqp.Table{}
			headers[hashHeader] = k
			msgID := uuid.New()
			pb := amqp.Publishing{
				ContentType: "application/json",
				Body:        msg.Content,
				Headers:     headers,
				MessageId:   msgID.String(),
			}
			err = ch.Publish(exchange, "", false, false, pb)
			if err != nil {
				panic(err)
			}
		}
	}
}
