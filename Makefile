.PHONY: setup
.PHONY: build

build:
	@which go > /dev/null || (echo "needs go installed" && exit 1)
	go build -o main

setup:
	@which go > /dev/null || (echo "needs go installed" && exit 1)
	go build -o main
	chmod +x main