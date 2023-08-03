package lido

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./lido.abi --pkg lido --type Lido --out ./lido.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./stmatic.abi --pkg lido --type LidoStmatic --out ./stmatic.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./wsteth.abi --pkg lido --type LidoWsteth --out ./wsteth.go
