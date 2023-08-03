package insta_dapp

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./guni_pool_static.abi --pkg insta_dapp --type InstaDappGUniPoolStatic --out ./guni_pool_static.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./insta_dsa_resolver.abi --pkg insta_dapp --type InstaDappDSAResolver --out ./insta_dsa_resolver.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./staking_rewards.abi --pkg insta_dapp --type InstaDappStakingRewards --out ./staking_rewards.go
