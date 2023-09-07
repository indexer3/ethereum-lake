package synthetix

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./fee_pool.abi --pkg synthetix --type SynthetixFeePool --out ./fee_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./issuer.abi --pkg synthetix --type SynthetixIssuer --out ./issuer.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./liquidator_rewards.abi --pkg synthetix --type SynthetixLiquidatorRewards --out ./liquidator_rewards.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./multi_collateral_synth.abi --pkg synthetix --type SynthetixMultiCollateral --out ./multi_collateral_synth.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./reward_escrow_v2.abi --pkg synthetix --type SynthetixRwardEscrowV2 --out ./reward_escrow_v2.go
var (
	SynthetixIssuerAddress_Ethereum     = common.HexToAddress("0x8D1201b074EeF198D5e6957F4574e2e4FC5Ba9Cf")
	SynthetixFeePoolAddress_Ethereum    = common.HexToAddress("0x3B2f389AeE480238A49E3A9985cd6815370712eB")
	SynthetixRewardEscrowV2             = common.HexToAddress("0xAc86855865CbF31c8f9FBB68C749AD5Bd72802e3")
	SynthetixLiquidatorRewards_Ethereum = common.HexToAddress("0xf79603a71144e415730C1A6f57F366E4Ea962C00")
	SynthetixNetworkToken_Ethereum      = common.HexToAddress("0x6b10E5Ce50e3A062731d83Cd3cAD1964e5F93DA6")
	SUSD_Ethereum                       = common.HexToAddress("0x10A5F7D9D65bCc2734763444D4940a31b109275f")
	SETH_Ethereum                       = common.HexToAddress("0x5D4C724BFe3a228Ff0E29125Ac1571FE093700a4")

	SynthetixStakingReward_Ethereum = common.HexToAddress("0x46C15afca591F7E4709Dd4369077Fa2dAA11FAf2")

	SynthetixIssueAddress_Ethereum = common.HexToAddress("0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f")



	Opt_SynthetixIssuer            = "0x59B01789bF268C7C77451D02758621990bB50BBF"
	Opt_SynthetixFeePool           = "0xD3739A5F06747e148E716Dcb7147B9BA15b70fcc"
	Opt_SynthetixVester            = "0x6330D5F08f51057F36F46d6751eCDc0c65Ef7E9e"
	Opt_sUSD                       = "0xDfA2d3a0d32F870D87f8A0d7AA6b9CdEB7bc5AdB"
	Opt_sUSDReal                   = "0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9"
	Opt_sETH                       = "0xe9dceA0136FEFC76c4E639Ec60CCE70482E2aCF7"
	Opt_sETHReal                   = "0xe405de8f52ba7559f9df3c368500b6e6ae6cee49"
	Opt_SynthetixLiquidatorRewards = "0xF4EebDD0704021eF2a6Bbe993fdf93030Cd784b4"
	Opt_SynthetixToken             = "0x8700dAec35aF8Ff88c16BdF0418774CB3D7599B4"
)



