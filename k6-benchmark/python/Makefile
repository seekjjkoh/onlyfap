run-server:
	python main.py

docker-build:
	docker build -t python-benchmark:0.0.1 .

docker-run:
	docker run --rm -p 1802:1802 python-benchmark:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 1802:1802 python-benchmark:0.0.1
