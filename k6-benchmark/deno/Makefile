run-server:
	deno run --allow-net main.ts

docker-build:
	docker build -t deno-benchmark:0.0.1 .

docker-run:
	docker run --rm -p 1803:1803 deno-benchmark:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 1803:1803 deno-benchmark:0.0.1
