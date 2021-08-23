package main

import "github.com/streadway/amqp"

func main() {
	cs := "amqp://guest:guest@localhost:5672/"
	topic := "hash.exchange"
	queues := map[string]string{
		"queue1": "1",
		"queue2": "2",
		"queue3": "3",
		"queue4": "4",
	}

	conn, err := amqp.Dial(cs)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// declare an exchange (x-consistent-hash)
	err = ch.ExchangeDeclare(topic, "x-consistent-hash", true, false, false, false, amqp.Table{})
	if err != nil {
		panic(err)
	}

	// declare queues
	for qName := range queues {
		_, err = ch.QueueDeclare(qName, true, false, false, false, amqp.Table{})
		if err != nil {
			panic(err)
		}
	}

	// bind queues to exchange with its routing key
	for qName, qKey := range queues {
		err = ch.QueueBind(qName, qKey, topic, false, nil)
		if err != nil {
			panic(err)
		}
	}
}
