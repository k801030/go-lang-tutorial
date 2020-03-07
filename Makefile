export IMAGE_NAME=go-docker
export CONTAINER_NAME=go-tutorial
export ECR=734574508434.dkr.ecr.ap-northeast-1.amazonaws.com/go-fargate

.PHONY: all test build run clean

all: clean test build run

test:
	go test -v ./web

build:
	docker build -t ${IMAGE_NAME} .

run:
	docker run -d -p 80:80 --name ${CONTAINER_NAME} --rm ${IMAGE_NAME}
	scripts/health-check.sh

clean:
	docker ps -q --filter "name=${CONTAINER_NAME}" | grep -q . && docker stop ${CONTAINER_NAME} || echo "no container found"

publish-to-ecr:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${ECR}
	docker tag ${IMAGE_NAME}:latest ${ECR}:latest
	docker push ${ECR}:latest
