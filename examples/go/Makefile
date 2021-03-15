
.PHONY: all
all: install

install:
	cd ../.. && go install .
	mkdir -p build
	frugal -gen go ../idl/music.frugal
	go build -v -o build/bin-httpClient ./httpClient
	go build -v -o build/bin-httpServer ./httpServer
	go build -v -o build/bin-natsClient ./natsClient
	go build -v -o build/bin-natsPublisher ./natsPublisher
	go build -v -o build/bin-natsServer ./natsServer
	go build -v -o build/bin-natsSubscriber ./natsSubscriber
	go build -v -o build/bin-stompPublisher ./stompPublisher
	go build -v -o build/bin-stompSubscriber ./stompSubscriber
