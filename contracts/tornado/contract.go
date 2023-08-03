package tornado

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./tornado_locked.abi --pkg tornado --type TornadoLocked --out ./tornado_locked.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./tornado_staking.abi --pkg tornado --type TornadoStaking --out ./tornado_staking.go
