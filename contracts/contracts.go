package contracts

type ContractType string

const (
	ContractTypeERC20   ContractType = "ERC20"
	ContractTypeERC721  ContractType = "ERC721"
	ContractTypeERC1155 ContractType = "ERC1155"

	ContractTypeSeaport       ContractType = "Seaport"
	ContractTypeMulticall     ContractType = "Multicall"
	ContractTypeOneInchOracle ContractType = "OneInchOracle"

	ContractTypeUniswapV2Factory ContractType = "UniswapV2Factory"
	ContractTypeUniswapV2Pair    ContractType = "UniswapV2Pair"

	ContractTypeUniswapV3Factory            ContractType = "UniswapV3Factory"
	ContractTypeUniswapV3Router             ContractType = "UniswapV3Router"
	ContractTypeUniswapV3Pool               ContractType = "UniswapV3Pool"
	ContractTypeUniswapV3NFTPositionManager ContractType = "UniswapV3NFTPositionManager"
)
