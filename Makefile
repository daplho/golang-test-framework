TARGET := golang-testing-framework

build:
	go build -o bin/${TARGET}

run: build
	./bin/${TARGET}

test:
	go test -v ./...
