APP=interview-cli

default:
	echo ${APP}

build:
	GOFLAGS=-mod=mod go build -o ${APP} main.go

run:
	go run main.go

clean:
	rm -rf ${APP}