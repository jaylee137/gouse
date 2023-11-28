export GO111MODULE=on

.PHONY: install
install:
	go get .
	go mod download

.PHONY: run
run:
	go run main.go

.PHONY: buid
build:
	go build ./...

.PHONY: test
test:
	go test -v -count=1 -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: format
format:
	gofmt -w -s .