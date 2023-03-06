SHELL := /bin/bash


.PHONY: generate
generate:
	echo "Generating protobuf..."
	./generate.sh

.PHONY: build
build:
	mkdir -p ./bin
	go mod vendor
	go build -o ./bin/middleware-server ./main.go

.PHONY: clean
clean:
	echo "cleaning project"
	rm -rf ./src/generated/
	rm -rf ./temp/
	rm -rf ./bin
	rm -rf ./mocks/