APP=interview-cli

default:
	echo ${APP}

.PHONEY: build
build: clean
	GOFLAGS=-mod=mod go build -o ${APP} main.go

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	rm -rf ${APP}