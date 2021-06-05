url := "http://host.docker.internal:1803"
vus = 100
duration = 30s

docker-run:
	docker run \
	-i --rm \
	--add-host=host.docker.internal:host-gateway \
	-v ${PWD}:/workdir \
	-e URL=${url} \
	loadimpact/k6 run \
	--vus ${vus} --duration ${duration} \
	- <script.js 
