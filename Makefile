all: build test

build: src/main.go
	go build -o bin/client src/main.go

test: