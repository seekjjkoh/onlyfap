builder_dockerfile := Dockerfile.build
server_dockerfile := Dockerfile.serve
luvit_builder := luvitbuilder
builder_version := 0.0.1
luvit_server := luvitserver
server_version := 0.0.1

docker-build-luvit:
	docker build -t ${luvit_builder}:${builder_version} -f ${builder_dockerfile} .
	docker run -v ${PWD}:/workdir ${luvit_builder}:${builder_version} luvi app -o server

docker-build-app:
	docker build -t ${luvit_server}:${server_version} -f ${server_dockerfile} .

docker-run-benchmark:
	docker run --rm --cpus="0.1" --memory="128m" -p 1337:1337 ${luvit_server}:${server_version}
