run-server:
	escript erlbenchmark.erl

docker-build:
	docker build -t erl-benchmark:0.0.1 .

docker-run:
	docker run --rm -p 2345:2345 erl-benchmark:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 2345:2345 erl-benchmark:0.0.1
