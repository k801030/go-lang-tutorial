export IMGAE_NAME=go-docker
export CONTAINER_NAME=go-tutorial

all: build run

build:
	docker build -t ${IMGAE_NAME} .

run:
	docker run -d -p 80:8080 --name ${CONTAINER_NAME} --rm ${IMGAE_NAME}

stop:
	docker stop ${CONTAINER_NAME}