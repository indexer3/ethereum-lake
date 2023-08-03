package rocketpool

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./rocket_mini_pool_delegate.abi --pkg rocketpool --type RocketPoolMiniPoolDelegate --out ./rocket_mini_pool_delegate.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./rocket_mini_pool_manager.abi --pkg rocketpool --type RocketPoolMiniPoolManager --out ./rocket_mini_pool_manager.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./rocket_node_staking.abi --pkg rocketpool --type RocketPoolNodeStaking --out ./rocket_node_staking.go
