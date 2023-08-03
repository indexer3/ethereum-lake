package curve

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./crypto_swap_factory.abi --pkg curve --type CurveCryptoSwapFactory --out ./crypto_swap_factory.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./crypto_swap_registry.abi --pkg curve --type CurveCryptoSwapRegistry --out ./crypto_swap_registry.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./curve_lp_token.abi --pkg curve --type CurveLpToken --out ./curve_lp_token.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./curve_pool_info_getter.abi --pkg curve --type CurvePoolInfoGetter --out ./curve_pool_info_getter.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./curve_pool.abi --pkg curve --type CurvePool --out ./curve_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./factory.abi --pkg curve --type CurveFactory --out ./factory.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./liquidity_gauge.abi --pkg curve --type CurveLiquidityGauge --out ./liquidity_gauge.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./metapool_factory.abi --pkg curve --type CurveMetaPoolFactory --out ./metapool_factory.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./registry.abi --pkg curve --type CurveRegistry --out ./registry.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./staking_rewards.abi --pkg curve --type CurveStakingRewards --out ./staking_rewards.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./vecrv.abi --pkg curve --type CurveVecrv --out ./vecrv.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./vesting_crv.abi --pkg curve --type CurveVestingCrv --out ./vesting_crv.go
