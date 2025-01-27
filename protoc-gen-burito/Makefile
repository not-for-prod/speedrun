.PHONY: run
run:
	protoc --plugin=protoc-gen-burito=./bin/protoc-gen-burito --burito_out=./example/out ./example/in/aboba.proto

.PHONY: build
build:
	go build -o bin/protoc-gen-burito main.go