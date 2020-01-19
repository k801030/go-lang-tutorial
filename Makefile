export IMGAE_NAME=go-docker
export CONTAINER_NAME=go-tutorial

.PHONY: all test build run clean

all: stop test build run

test:
	go test -v ./pkg

build:
	docker build -t ${IMGAE_NAME} .

run:
	docker run -d -p 80:8080 --name ${CONTAINER_NAME} --rm ${IMGAE_NAME}
	scripts/health-check.sh

clean:
	docker ps -q --filter "name=${CONTAINER_NAME}" | grep -q . && docker stop ${CONTAINER_NAME} || echo "no container found"