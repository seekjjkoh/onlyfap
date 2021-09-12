package main

import "github.com/streadway/amqp"

func main() {
	cs := "amqp://guest:guest@localhost:5672/"
	exchange := "header.exchange"
	queues := map[string]int{
		"h1": 1,
		"h2": 2,
		"h3": 3,
		"h4": 4,
	}
	// this is the alternate-exchange
	// when there's no possible route
	defaultQ := "h0"
	rKey := "header.routekey"
	headerName := "my-hash-header"

	conn, err := amqp.Dial(cs)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// header exchange
	err = ch.ExchangeDeclare(exchange, "headers", true, false, false, false, amqp.Table{
		"alternate-exchange": defaultQ,
	})
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare(defaultQ, "fanout", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// declare queues
	_, err = ch.QueueDeclare(defaultQ, true, false, false, false, amqp.Table{})
	if err != nil {
		panic(err)
	}
	for qName := range queues {
		_, err = ch.QueueDeclare(qName, true, false, false, false, amqp.Table{})
		if err != nil {
			panic(err)
		}
	}

	// bind queues to exchange with its routing key
	err = ch.QueueBind(defaultQ, "", defaultQ, false, nil)
	if err != nil {
		panic(err)
	}
	for qName, qKey := range queues {
		err = ch.QueueBind(qName, rKey, exchange, false, amqp.Table{
			"x-match":  "all",
			headerName: qKey,
		})
		if err != nil {
			panic(err)
		}
	}
}
