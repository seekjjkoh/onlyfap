run-server:
	cd play-scala-seed && sbt run

docker-build:
	cd play-scala-seed && sbt docker:publishLocal

docker-run:
	docker run --rm -p 9000:9000 play-scala-seed:0.0.1

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 9000:9000 play-scala-seed:0.0.1
