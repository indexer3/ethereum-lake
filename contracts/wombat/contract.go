package wombat

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./master_wombat_v2.abi --pkg wombat --type MasterWomatV2 --out ./master_wombat_v2.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./vewom.abi --pkg wombat --type Vewom --out ./vewom.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./wombat_asset.abi --pkg wombat --type WombatAsset --out ./wombat_asset.go
