package liquity

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./lqty_staking.abi --pkg liquity --type LiquityLQTYStaking --out ./lqty_staking.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./stability_pool.abi --pkg liquity --type LiquityStabilityPool --out ./stability_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./trove_manager.abi --pkg liquity --type LiquityTroveManager --out ./trove_manager.go
