all: test build

clean:
	rm -rf bin

build: src/client.go
	go build -o bin/client.exe src/client.go

test:
	go test ./src/