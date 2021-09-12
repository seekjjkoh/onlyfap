package main

import "github.com/streadway/amqp"

func main() {
	cs := "amqp://guest:guest@localhost:5672/"
	exchange := "hashheader.exchange"
	queues := map[string]string{
		"h1": "1",
		"h2": "2",
		"h3": "3",
		"h4": "4",
	}
	hashHeader := "hh"

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
	// note `hash-header` will define where to hash
	err = ch.ExchangeDeclare(exchange, "x-consistent-hash", true, false, false, false, amqp.Table{"hash-header": hashHeader})
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
		err = ch.QueueBind(qName, qKey, exchange, false, nil)
		if err != nil {
			panic(err)
		}
	}
}
