run-server:
	cargo run

go-build:
	cargo build --release

docker-build:
	docker build -t rust-benchmark:0.0.1 .

docker-run:
	docker run --rm -p 8000:8000 rust-benchmark:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 8000:8000 rust-benchmark:0.0.1
