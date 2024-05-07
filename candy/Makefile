all: build

build: build_server

build_server:
	go build -o candy_server cmd/candyServer/candyServer.go

clean:
	rm -rf candy_server