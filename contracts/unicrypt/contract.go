package unicrypt

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./unicrypt_token_vesting.abi --pkg unicrypt --type UnicryptTokenVesting --out ./unicrypt_token_vesting.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./unicrypt_uniswap_v2_locker.abi --pkg unicrypt --type UnicryptUniswapV2Locker --out ./unicrypt_uniswap_v2_locker.go
