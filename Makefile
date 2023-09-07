.PHONY: build

generate: 
	@echo "Generating..."
	go generate ./contracts/...

build: 
	go build -o build/ethereum-lake .

run-indexer:
	go run . ethereum-lake indexer --config ./configs/indexer.yaml

lint: 
	go mod tidy 
	golangci-lint run --max-same-issues=0 --fix ./...

test: 
	go test --race --cover ./...
