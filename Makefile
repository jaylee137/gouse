export GO111MODULE=on

install:
	go get -v ./...
	go mod download
	go install golang.org/x/tools/cmd/goimports@latest

run:
	go run main.go

build:
	go build ./...

doc:
	@echo "Generating docs..."
	go run tools/doc.go array console date function helper io math number shared strings types
	@echo "Done!"

test:
	@echo "Running tests..."
	go clean -testcache && go vet -v ./...
    # go test -timeout 30s ./pkg/helpers -run ^TestParallelize$ -v
	go test -v -count=1 -cover -coverprofile=coverage.out ./array/... ./constants/... ./shared/...
	go tool cover -func=coverage.out
	@echo "Done!"

lint:
	@echo "Running linter..."
	gofmt -w -s . && goimports -w . && go fmt ./...
	@echo "Done!"

clean:
	go clean -i -x -cache -testcache -modcache
	rm -rf coverage.out
	rm -rf ./tmp