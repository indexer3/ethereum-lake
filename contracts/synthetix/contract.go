package synthetix

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./fee_pool.abi --pkg synthetix --type SynthetixFeePool --out ./fee_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./issuer.abi --pkg synthetix --type SynthetixIssuer --out ./issuer.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./liquidator_rewards.abi --pkg synthetix --type SynthetixLiquidatorRewards --out ./liquidator_rewards.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./multi_collateral_synth.abi --pkg synthetix --type SynthetixMultiCollateral --out ./multi_collateral_synth.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./reward_escrow_v2.abi --pkg synthetix --type SynthetixRwardEscrowV2 --out ./reward_escrow_v2.go
var (
	
)
