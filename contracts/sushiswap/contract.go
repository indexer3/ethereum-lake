package sushiswap

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./l2_master_chef_v2.abi --pkg sushiswap --type SushiswapL2MasterChefV2 --out ./l2_master_chef_v2.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./master_chef_v2.abi --pkg sushiswap --type SushiswapMasterChefV2 --out ./master_chef_v2.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./master_chef.abi --pkg sushiswap --type SushiswapMasterChef --out ./master_chef.go
