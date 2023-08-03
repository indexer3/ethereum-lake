package aura

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aura_booster.abi --pkg aura --type AuraBooster --out ./aura_booster.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aura_locker.abi --pkg aura --type AuraLocker --out ./aura_locker.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aura_token.abi --pkg aura --type AuraToken --out ./aura_token.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./base_reward_pool.abi --pkg aura --type BaseRewardPool --out ./base_reward_pool.go
