package pancake

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./pancake_cake_pool.abi --pkg pancake --type PancakeCakePool --out ./pancake_cake_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./pancake_masterchef.abi --pkg pancake --type PancakeMasterChef --out ./pancake_masterchef.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./pancake_pair.abi --pkg pancake --type PancakePair --out ./pancake_pair.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./pancake_syrup.abi --pkg pancake --type PancakeSyrup --out ./pancake_syrup.go
