run-server:
	go run main.go

go-build:
	CGO_ENABLED=0 GOOS=linux go build -v -o server ./main.go

docker-build:
	docker build -t go-benchmark:0.0.1 .

docker-run:
	docker run --rm -p 1800:1800 go-benchmark:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 1800:1800 go-benchmark:0.0.1
