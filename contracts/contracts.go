package contracts

type ContractType string

const (
	ContractTypeERC20   ContractType = "ERC20"
	ContractTypeERC721  ContractType = "ERC721"
	ContractTypeERC1155 ContractType = "ERC1155"

	ContractTypeSeaport       ContractType = "Seaport"
	ContractTypeMulticall     ContractType = "Multicall"
	ContractTypeOneInchOracle ContractType = "OneInchOracle"
)
