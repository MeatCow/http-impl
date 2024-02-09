all: build test

clean:
	rm -rf bin

build: src/main.go
	go build -o bin/client.exe src/main.go

test:
	./bin/client.exe