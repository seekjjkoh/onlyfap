service_name := pongrun
version := 0.0.14
project_id := jjkoh95

dev:
	docker run --name redis --rm -d -p 6379:6379 redis
	go run cmd/server/main.go

push:
	CGO_ENABLED=0 GOOS=linux go build -v -o server ./cmd/server/main.go
	sudo docker build -t asia.gcr.io/${project_id}/${service_name}:${version} .
	sudo docker push asia.gcr.io/${project_id}/${service_name}:${version}

clean:
	docker rmi asia.gcr.io/${project_id}/${service_name}:${version}
	rm server
