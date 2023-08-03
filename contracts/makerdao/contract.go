package makerdao

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./cdps.abi --pkg makerdao --type MakerDaoCDPS --out ./cdps.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./ilk_registry.abi --pkg makerdao --type MakerDaoIlkRegistry --out ./ilk_registry.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./mcd_vat.abi --pkg makerdao --type MakerDaoMcdVat --out ./mcd_vat.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./proxy_registry.abi --pkg makerdao --type MakerDaoProxyRegistry --out ./proxy_registry.go
