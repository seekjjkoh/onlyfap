build-rabbitmq:
	docker build -t rabbitmq-with-hash:0.0.1 -f Dockerfile.rabbitmq .

start-rabbitmq:	
	docker run --rm -d \
	--hostname localrabbitmanagement \
	--name myrabbitmqcse \
	-p 15672:15672 \
	-p 5672:5672 \
	rabbitmq-with-hash:0.0.1
