.PHONY:

test:
	go test ./... -count=1

build:
	cd src && go build -o ../bin/challenge.exe

u-dep:
	cd src && go get -u && go mod tidy

run:
	cd src && go run . local
