package aave

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

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
var (
	aaveV2Incentives_Ethereum      = common.HexToAddress("0xd784927Ff2f95ba542BfC824c8a8a98F3495f6b5")
	aaveV2DataProvider_Ethereum    = common.HexToAddress("0x057835ad21a177dbdd3090bb1cae03eacf78fc6d")
	aaveV2LendingPool_Ethereum     = common.HexToAddress("0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9")
	aaveV2StakedToken_Ethereum     = common.HexToAddress("0x4da27a545c0c5b758a6ba100e3a049001de870f5")
	aaveV2StakedAbptToken_Ethereum = common.HexToAddress("0xa1116930326d21fb917d5a27f1e9943a9595fb47")

	aaveV2Incentives_Polygon   = common.HexToAddress("0x357D51124f59836DeD84c8a1730D72B749d8BC23")
	aaveV2DataProvider_Polygon = common.HexToAddress("0x7551b5D2763519d4e37e8B81929D336De671d46d")
	aaveV2LendingPool_Polygon  = common.HexToAddress("0x8dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf")
	// aaveV2StakedTokenPolygon     = common.HexToAddress("0x4da27a545c0c5b758a6ba100e3a049001de870f5")
	aaveV2StakedAbptToken_Polygon = common.HexToAddress("0xa1116930326d21fb917d5a27f1e9943a9595fb47")
)

func AaveV2Incentives(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return aaveV2Incentives_Ethereum
	case constant.NetworkPolygon:
		return aaveV2Incentives_Polygon
	default:
		return aaveV2Incentives_Ethereum
	}
}

func AaveV2DataProvider(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return aaveV2DataProvider_Ethereum
	case constant.NetworkPolygon:
		return aaveV2DataProvider_Polygon
	default:
		return aaveV2DataProvider_Ethereum
	}
}

func AaveV2LendingPool(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return aaveV2LendingPool_Ethereum
	case constant.NetworkPolygon:
		return aaveV2LendingPool_Polygon
	default:
		return aaveV2LendingPool_Ethereum
	}
}

func AaveV2StakedToken(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return aaveV2StakedToken_Ethereum
	// case constant.NetworkPolygon:
	// 	return aaveV2StakedTokenPolygon
	default:
		return aaveV2StakedToken_Ethereum
	}
}

func AaveV2StakedAbptToken(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return aaveV2StakedAbptToken_Ethereum
	case constant.NetworkPolygon:
		return aaveV2StakedAbptToken_Polygon
	default:
		return aaveV2StakedAbptToken_Ethereum
	}
}
