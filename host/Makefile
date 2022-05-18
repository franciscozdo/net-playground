proto:
	protoc -I=.. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import ../proto/*

build:
	go build

mod-update:
	go mod tidy

all: mod-update proto build

.PHONY: all
