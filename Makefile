.PHONY: build

generate: 
	@echo "Generating..."
	go generate ./contracts/...

build: 
	go build -o build/ethereum-lake .

lint: 
	go mod tidy 
	golangci-lint run --max-same-issues=0 --fix

test: 
	go test --cover ./...
