all: build

build:
	@go generate
	@go build -o nvltr main.go

run:
	@make build
	@./nvltr

vet:
	@go vet ./...

fmt:
	@go fmt ./...

test:
	@go test ./...

deploy:
	@go generate
	GOOS=linux GOARCH=amd64 go build -v -o nvltr
	@make sendscp

sendscp:
	@scp ./nvltr user@host:/home/user/bin/nvltr
	@scp ./config.yml user@host:/home/user/bin/config.yml
	@scp -r ./cert user@host:/home/user/bin/cert

.PHONY: all build vet test deploy fmt run sendscp deploy
