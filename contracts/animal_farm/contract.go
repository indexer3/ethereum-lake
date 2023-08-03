package animal_farm

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./dog_pound_auto_pool.abi --pkg animal_farm --type AnimalFarm --out ./dog_pound_auto_pool.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./master_chef_dogs_v2.abi --pkg animal_farm --type MasterChefDogsV2 --out ./master_chef_dogs_v2.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./master_chef_pigs_v2.abi --pkg animal_farm --type MasterChefPigsV2 --out ./master_chef_pigs_v2.go
var (
	
)
