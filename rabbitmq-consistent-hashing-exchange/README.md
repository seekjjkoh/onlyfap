# RabbitMQ Consistent Hashing Exchange

## tldr
```bash
# build rabbitmq
$ make build-rabbitmq
# start rabbitmq
$ make start-rabbitmq

# to try che-header
# declare exchange/queue definition
$ go run cmd/che-header/definition/main.go
# spin each in a separate terminal
$ go run cmd/che-header/consumer/main.go --qName=h1
$ go run cmd/che-header/consumer/main.go --qName=h2
$ go run cmd/che-header/consumer/main.go --qName=h3
$ go run cmd/che-header/consumer/main.go --qName=h4
# run producer
$ go run cmd/che-header/producer/main.go

## you should see the counter of receive key (hashid)
```
