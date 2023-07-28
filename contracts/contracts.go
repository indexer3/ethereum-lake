package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/indexer3/ethereum-lake/contracts/erc1155"
	"github.com/indexer3/ethereum-lake/contracts/erc20"
	"github.com/indexer3/ethereum-lake/contracts/erc721"
	"github.com/indexer3/ethereum-lake/contracts/multicall"
	"github.com/indexer3/ethereum-lake/contracts/one_inch_oracle/one_inch_oracle"
	"github.com/indexer3/ethereum-lake/contracts/opensea/seaport"
	"github.com/indexer3/ethereum-lake/contracts/uniswap/uniswap_v2_factory"
	"github.com/indexer3/ethereum-lake/contracts/uniswap/uniswap_v2_pair"
	"github.com/indexer3/ethereum-lake/contracts/uniswap/uniswap_v3_factory"
	"github.com/indexer3/ethereum-lake/contracts/uniswap/uniswap_v3_nft_position_manager"
	"github.com/indexer3/ethereum-lake/contracts/uniswap/uniswap_v3_pool"
	"github.com/indexer3/ethereum-lake/contracts/uniswap/uniswap_v3_router"
)

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

var ABIs = make(map[ContractType]*abi.ABI)

func init() {
	erc20ABI, _ := erc20.ERC20MetaData.GetAbi()
	erc721ABI, _ := erc721.ERC721MetaData.GetAbi()
	erc1155ABI, _ := erc1155.ERC1155MetaData.GetAbi()
	multicallABI, _ := multicall.MulticallMetaData.GetAbi()
	seaportABI, _ := seaport.SeaportMetaData.GetAbi()
	oneInchOracleABI, _ := one_inch_oracle.OneInchOracleMetaData.GetAbi()
	uniswapv2FactoryABI, _ := uniswap_v2_factory.UniswapV2FactoryMetaData.GetAbi()
	uniswapv2PairABI, _ := uniswap_v2_pair.UniswapV2PairMetaData.GetAbi()
	uniswapv3FactoryABI, _ := uniswap_v3_factory.UniswapV3FactoryMetaData.GetAbi()
	uniswapv3PoolABI, _ := uniswap_v3_pool.UniswapV3PoolMetaData.GetAbi()
	uniswapv3NFTPositionManagerABI, _ := uniswap_v3_nft_position_manager.UniswapV3NFTPositionManagerMetaData.GetAbi()
	uniswapv3RouterABI, _ := uniswap_v3_router.UniswapV3RouterMetaData.GetAbi()

	ABIs[ContractTypeERC20] = erc20ABI
	ABIs[ContractTypeERC721] = erc721ABI
	ABIs[ContractTypeERC1155] = erc1155ABI
	ABIs[ContractTypeMulticall] = multicallABI
	ABIs[ContractTypeSeaport] = seaportABI
	ABIs[ContractTypeOneInchOracle] = oneInchOracleABI
	ABIs[ContractTypeUniswapV2Factory] = uniswapv2FactoryABI
	ABIs[ContractTypeUniswapV2Pair] = uniswapv2PairABI
	ABIs[ContractTypeUniswapV3Factory] = uniswapv3FactoryABI
	ABIs[ContractTypeUniswapV3Pool] = uniswapv3PoolABI
	ABIs[ContractTypeUniswapV3NFTPositionManager] = uniswapv3NFTPositionManagerABI
	ABIs[ContractTypeUniswapV3Router] = uniswapv3RouterABI
}
