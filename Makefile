export NAME=go-docker

all: build run

build:
	docker build -t ${NAME} .

run:
	docker run -d -p 80:8080 --name test ${NAME}