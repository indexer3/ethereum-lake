package frax

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./convex_staking_wrapped_frax.abi --pkg frax --type FraxConvexStakingWrappedFrax --out ./convex_staking_wrapped_frax.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./frax_farm_uni_v3.abi --pkg frax --type FraxFarmUniV3 --out ./frax_farm_uni_v3.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./frax_lp_token.abi --pkg frax --type FraxLpToken --out ./frax_lp_token.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./frax_unified_farm_erc20_convex/frax_unified_farm_erc20_convex.abi --pkg frax_unified_farm_erc20_convex --type FraxUnifiedFarmErc20Convex --out ./frax_unified_farm_erc20_convex/frax_unified_farm_erc20_convex.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./frax_unified_farm_erc20_fraxswap/frax_unified_farm_erc20_fraxswap_v2.abi --pkg frax_unified_farm_erc20_fraxswap --type FraxUnifiedFarmErc20FraxswapV2 --out ./frax_unified_farm_erc20_fraxswap/frax_unified_farm_erc20_fraxswap_v2.go

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./sfrx_eth.abi --pkg frax --type FraxSfrxEth --out ./sfrx_eth.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./ve_fxs_yield_distributor.abi --pkg frax --type FraxVeFxsYieldDistributor --out ./ve_fxs_yield_distributor.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./ve_fxs.abi --pkg frax --type FraxVeFxs --out ./ve_fxs.go
