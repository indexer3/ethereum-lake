package aave

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_a_token.abi --pkg aave --type AaveAToken --out ./aave_a_token.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_abpt_staked.abi --pkg aave --type AaveAbptStaked --out ./aave_abpt_staked.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_data_provider/aave_data_provider.abi --pkg aave --type AaveDataProvider --out ./aave_data_provider/aave_data_provider.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_incentives_controller.abi --pkg aave --type AaveIncentivesController --out ./aave_incentives_controller.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_lending_pool.abi --pkg aave --type AaveLendingPool --out ./aave_lending_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_stable_debt.abi --pkg aave --type AaveStableDebt --out ./aave_stable_debt.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_staked_contract/aave_staked_contract.abi --pkg aave --type AaveStakedContract --out ./aave_staked_contract/aave_staked_contract.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_staked/aave_staked.abi --pkg aave --type AaveStaked --out ./aave_staked/aave_staked.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_v3_protocol_data_provider/aave_v3_protocol_data_provider.abi --pkg aave --type AaveV3ProtocolDataProvider --out ./aave_v3_protocol_data_provider/aave_v3_protocol_data_provider.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_v3_reward_controller.abi --pkg aave --type AaveV3RewardController --out ./aave_v3_reward_controller.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./aave_variable_debt.abi --pkg aave --type AaveVariableDebt --out ./aave_variable_debt.go
var ()
