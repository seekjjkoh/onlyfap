VERSION := 0.0.19
NETWORK := datadognetwork

create-network:
	docker network create ${NETWORK}

start-datadog-agent:
	docker run --rm --name datadog-agent \
		--network ${NETWORK} \
		-v /var/run/docker.sock:/var/run/docker.sock:ro \
		-v /proc/:/host/proc/:ro \
		-v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
		-p 127.0.0.1:8126:8126/tcp \
		-e DD_API_KEY=${DATADOG_API_KEY} \
		-e DD_APM_ENABLED=true \
		-e DD_APM_NON_LOCAL_TRAFFIC=true \
		-e DD_LOGS_ENABLED=true \
		-e DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL=true \
		-e DD_LOGS_CONFIG_DOCKER_CONTAINER_USE_FILE=true \
		-e DD_CONTAINER_EXCLUDE="name:datadog-agent" \
		gcr.io/datadoghq/agent:latest

build-server:
	CGO_ENABLED=0 GOOS=linux go build -v -o server ./cmd/main.go
	docker build -t datadogapm:${VERSION} .

start-server-8080:
	docker run --rm --name server8080 \
		--network ${NETWORK} \
		-e DD_AGENT_HOST=datadog-agent \
		-p 8080:8080 \
		datadogapm:${VERSION} \
		./server --datadogAgentAddress=datadog-agent:8126 --timeout=1 --targetURL=http://host8081:8081

start-server-8081:
	docker run --rm --name host8081 \
		--network ${NETWORK} \
		-e DD_AGENT_HOST=datadog-agent \
		-p 8081:8081 \
		datadogapm:${VERSION} \
		./server --datadogAgentAddress=datadog-agent:8126 --timeout=5 --port=8081 --targetURL=http://host8082:8082

start-server-8082:
	docker run --rm --name host8082 \
		--network ${NETWORK} \
		-e DD_AGENT_HOST=datadog-agent \
		-p 8082:8082 \
		datadogapm:${VERSION} \
		./server --datadogAgentAddress=datadog-agent:8126 --timeout=2 --port=8082
