.PHONY: build


build: 
	go build -o build/ethereum-lake .

lint: 
	go mod tidy 
	golangci-lint run 

test: 
	go test --cover ./...
