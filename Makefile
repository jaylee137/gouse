install:
	go get -v ./...
	go mod download
	go install golang.org/x/tools/cmd/goimports@latest
	go get golang.org/x/perf/cmd/benchstat
	go install golang.org/x/perf/cmd/benchstat
	go install honnef.co/go/tools/cmd/staticcheck@latest

start:
	go run main.go

dev:
	go run main.go -isDev=true

build:
	go build ./...

doc:
	@echo "Generating docs..."
	go run docs/generate.go ./sample
	@echo "Done!"

test:
	@echo "Running tests..."
	go clean -testcache
	go test -v -count=1 -cover -coverprofile=coverage.out ./date/... ./number/... ./regex/... ./strings/... ./structs/... ./types/...
	go tool cover -func=coverage.out
	@echo "Done!"

BENCH_CMD = go test -count 5 -run=^\# -bench=. ./number/... ./regex/...

bench:
	@echo "Running benchmarks..."
	$(BENCH_CMD)
	@echo "Done!"

bench_filter:
	BENCH_CMD -benchmem > $(f)

bench_stat:
	@echo "Running stat filter benchmarks..."
	make bench_filter f=old.txt
	make bench_filter f=new.txt
	benchstat old.txt new.txt
	@echo "Done!"

format:
	@echo "Running format..."
	gofmt -w -s . && goimports -w . && go fmt ./...
	@echo "Done!"

lint:
	@echo "Running lint..."
	staticcheck ./...
	@echo "Done!"

count:
	@echo "Counting lines..."
	bash count.sh public/count.svg true 13708a array cache crypto console date function helper io log math net number regex strings structs tools types
	@echo "Done!"

pre:
	make build && make test && make format && make lint && make doc
	git add .

clean:
	go clean -i -x -cache -testcache -modcache
	rm -rf coverage.out
	rm -rf ./tmp