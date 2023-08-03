package shibaswap

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./bone_locker.abi --pkg shibaswap --type ShibaswapBoneLocker --out ./bone_locker.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./top_dog.abi --pkg shibaswap --type ShibaswapTopDog --out ./top_dog.go
