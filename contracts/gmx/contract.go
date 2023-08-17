package gmx

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./gmx_vester.abi --pkg gmx --type GmxVester --out ./gmx_vester.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./reward_tracker.abi --pkg gmx --type GmxRewardTracker --out ./reward_tracker.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./staked_glp.abi --pkg gmx --type GmxStakedGlp --out ./staked_glp.go

var (
	rewardTrackerSGMXAddress_Arbitrum       = common.HexToAddress("0x908c4d94d34924765f1edc22a1dd098397c59dd4")
	rewardTrackerSbfGmxAddress_Arbitrum = common.HexToAddress("0xd2D1162512F927a7e282Ef43a362659E4F2a728F")
	rewardTrackerGlpAddress_Arbitrum    = common.HexToAddress("0x4e971a87900b931ff39d1aad67697f49835400b6")
	stakedGlpAddress_Arbitrum           = common.HexToAddress("0x5402B5F40310bDED796c7D0F3FF6683f5C0cFfdf")
	gmxTokenAddress_Arbitrum            = common.HexToAddress("0xfc5A1A6EB076a2C7aD06eD22C90d7E710E35ad0a")
	esTokenAddress_Arbitrum             = common.HexToAddress("0xf42ae1d54fd613c9bb14810b0588faaa09a426ca")
	gmxLpAddress_Arbitrum               = common.HexToAddress("0x4277f8f2c384827b5273592ff7cebd9f2c1ac258")
	vesterGMXAddress_Arbitrum           = common.HexToAddress("0x199070DDfd1CFb69173aa2F7e20906F26B363004")
	vesterGlpAddress_Arbitrum           = common.HexToAddress("0xA75287d2f8b217273E7FCD7E86eF07D33972042E")
)

func RewardTrackerSGMXAddress() common.Address {
	return rewardTrackerSGMXAddress_Arbitrum
}

func RewardTrackerSbfGmxAddress() common.Address {
	return rewardTrackerSbfGmxAddress_Arbitrum
}

func RewardTrackerGlpAddress() common.Address {
	return rewardTrackerGlpAddress_Arbitrum
}

func StakedGlpAddress() common.Address {
	return stakedGlpAddress_Arbitrum
}

func GMXTokenAddress() common.Address {
	return gmxTokenAddress_Arbitrum
}

func EsTokenAddress() common.Address {
	return esTokenAddress_Arbitrum
}

func GMXLpAddress() common.Address {
	return gmxLpAddress_Arbitrum
}

func VesterGMXAddress() common.Address {
	return vesterGMXAddress_Arbitrum
}

func VesterGlpAddress() common.Address {
	return vesterGlpAddress_Arbitrum
}
