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
	go run docs/generate.go array console date function helper io math number regex strings structs types
	@echo "Done!"

test:
	@echo "Running tests..."
	go clean -testcache
	go test -v -count=1 -cover -coverprofile=coverage.out ./date/... ./number/... ./regex/... ./strings/... ./types/...
	go tool cover -func=coverage.out
	@echo "Done!"

lint:
	@echo "Running linter..."
	gofmt -w -s . && goimports -w . && go fmt ./...
	@echo "Done!"

count:
	@echo "Counting lines..."
	sh count.sh public/count.svg true 13708a array console date function helper io math number regex strings structs types
	@echo "Done!"

precommit:
	make build && make test && make doc && make lint
	git add .

clean:
	go clean -i -x -cache -testcache -modcache
	rm -rf coverage.out
	rm -rf ./tmp