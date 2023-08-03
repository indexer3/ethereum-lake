package maple

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./maple_pool.abi --pkg maple --type MaplePool --out ./maple_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./maple_stake_locker.abi --pkg maple --type MapleStakeLocker --out ./maple_stake_locker.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./xmpl.abi --pkg maple --type MapleXmpl --out ./xmpl.go
