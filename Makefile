export GO111MODULE=on

install:
	go get -v ./...
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

run:
	go run main.go

build:
	go build ./...

test:
	@echo "Running tests..."
    # go clean -testcache && go vet -v ./... && govulncheck ./...
    # go test -timeout 30s ./pkg/helpers -run ^TestParallelize$ -v
	go test -v -count=1 -cover -coverprofile=coverage.out ./array/... ./constants/... ./shared/...
	go tool cover -func=coverage.out
	@echo "Done!"

lint:
	@echo "Running linter..."
	gofmt -w -s . && goimports -w . && go fmt ./...
	golangci-lint version
	golangci-lint run -c .golangci.yml ./...
	@echo "Done!"

clean:
	go clean -i -x -cache -testcache -modcache
	rm -rf coverage.out
	rm -rf ./tmp