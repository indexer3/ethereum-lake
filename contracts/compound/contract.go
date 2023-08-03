package compound

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./cerc20.abi --pkg compound --type CompoundCErc20 --out ./cerc20.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./cether.abi --pkg compound --type CompoundCEther --out ./cether.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./compound_lens.abi --pkg compound --type CompoundLens --out ./compound_lens.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./comptroller.abi --pkg compound --type Comptroller --out ./comptroller.go
