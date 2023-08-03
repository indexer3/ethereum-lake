package venus

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./vai_vault_proxy.abi --pkg venus --type VenusVaiVaultProxy --out ./vai_vault_proxy.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./venus_comproller.abi --pkg venus --type VenusComproller --out ./venus_comproller.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./venus_token.abi --pkg venus --type VenusToken --out ./venus_token.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./xvs_vault_proxy.abi --pkg venus --type VenusXvsVaultProxy --out ./xvs_vault_proxy.go
