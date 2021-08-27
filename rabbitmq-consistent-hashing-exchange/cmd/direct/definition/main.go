package main

import "github.com/streadway/amqp"

func main() {
	cs := "amqp://guest:guest@localhost:5672/"
	qName := "naive.exchange"
	conn, err := amqp.Dial(cs)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(qName, true, false, false, false, amqp.Table{})
	if err != nil {
		panic(err)
	}
}
