run-server:
	node index.js

docker-build:
	docker build -t node-benchmark:0.0.1 .

docker-run:
	docker run --rm -p 1801:1801 node-benchmark:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 1801:1801 node-benchmark:0.0.1
