IMAGE_NAME ?= animals-api:latest

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
	go build -o bin/animals-api .

build-and-push: fmt vet test
	docker build . -t ${IMAGE_NAME}
	docker push ${IMAGE_NAME}

deploy:
	kubectl apply -f k8s/namespace.yaml
	(cd k8s/; kustomize edit set image animals-api=${IMAGE_NAME})
	(cd k8s/; kustomize build . | kubectl apply -f -)
