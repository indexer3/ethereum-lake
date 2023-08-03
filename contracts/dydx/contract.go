package dydx

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./staked_dydx.abi --pkg dydx --type StakedDydx --out ./staked_dydx.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./staked_usdc.abi --pkg dydx --type StakedUSDC --out ./staked_usdc.go
