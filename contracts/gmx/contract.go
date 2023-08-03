package gmx

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./gmx_vester.abi --pkg gmx --type GmxVester --out ./gmx_vester.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./reward_tracker.abi --pkg gmx --type GmxRewardTracker --out ./reward_tracker.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./staked_glp.abi --pkg gmx --type GmxStakedGlp --out ./staked_glp.go
