package stargate

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./stargate_lp_staking.abi --pkg stargate --type StargateLpStaking --out ./stargate_lp_staking.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./stargate_pool.abi --pkg stargate --type StargatePool --out ./stargate_pool.go
