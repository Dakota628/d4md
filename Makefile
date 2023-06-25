.PHONY: lint
lint:
	go mod tidy

.PHONY: build
build:
	go build

.PHONY: start
start: build
	./d4md

.PHONY: clean
clean:
	rm -f d4md d4md.exe

