IMAGE_NAME ?= frontend:latest

all: test docker-build

fmt:
	go fmt ./...

vet:
	go vet ./...

test: fmt vet
	go test -v ./...

# local run
run: fmt vet
	go run main.go "${@:-}"

build: fmt vet test
	go build -o bin/frontend .

build-and-push: fmt vet test
	docker build . -t ${IMAGE_NAME}
	docker push ${IMAGE_NAME}

deploy:
	cd k8s/
	kustomize build . | kubectl apply -f -
