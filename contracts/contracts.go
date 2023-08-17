package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	one_inch "github.com/indexer3/ethereum-lake/contracts/1inch"
	"github.com/indexer3/ethereum-lake/contracts/aave"
	aave_data_provider "github.com/indexer3/ethereum-lake/contracts/aave/aave_data_provider"
	aave_staked_contract "github.com/indexer3/ethereum-lake/contracts/aave/aave_staked_contract"
	"github.com/indexer3/ethereum-lake/contracts/erc1155"
	"github.com/indexer3/ethereum-lake/contracts/erc20"
	"github.com/indexer3/ethereum-lake/contracts/erc721"
	"github.com/indexer3/ethereum-lake/contracts/gmx"
	"github.com/indexer3/ethereum-lake/contracts/multicall"
	"github.com/indexer3/ethereum-lake/contracts/opensea/seaport"
	"github.com/indexer3/ethereum-lake/contracts/uniswap"
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

	ContractTypeAaveV2Incentives   ContractType = "AaveV2Incentives"
	ContractTypeAaveV2DataProvider ContractType = "AaveV2DataProvider"
	ContractTypeAaveV2LendingPool  ContractType = "AaveV2LendingPool"
	ContractTypeAaveV2StakedToken  ContractType = "AaveV2StakedToken"
	ContractTypeAaveV2StakedAbpt   ContractType = "AaveV2StakedAbpt"

	ContractTypeGMXVester        ContractType = "GMXVester"
	ContractTypeGMXRewardTracker ContractType = "GMXRewardTracker"
	ContractTypeGMXStakedGLP     ContractType = "GMXStakedGLP"
)

var ABIs = make(map[ContractType]*abi.ABI)

func init() {
	erc20ABI, _ := erc20.ERC20MetaData.GetAbi()
	erc721ABI, _ := erc721.ERC721MetaData.GetAbi()
	erc1155ABI, _ := erc1155.ERC1155MetaData.GetAbi()

	multicallABI, _ := multicall.MulticallMetaData.GetAbi()

	seaportABI, _ := seaport.SeaportMetaData.GetAbi()

	oneInchOracleABI, _ := one_inch.OneInchOracleMetaData.GetAbi()

	uniswapv2FactoryABI, _ := uniswap.UniswapV2FactoryMetaData.GetAbi()
	uniswapv2PairABI, _ := uniswap.UniswapV2PairMetaData.GetAbi()

	uniswapv3FactoryABI, _ := uniswap.UniswapV3FactoryMetaData.GetAbi()
	uniswapv3PoolABI, _ := uniswap.UniswapV3PoolMetaData.GetAbi()
	uniswapv3NFTPositionManagerABI, _ := uniswap.UniswapV3NFTPositionManagerMetaData.GetAbi()
	uniswapv3RouterABI, _ := uniswap.UniswapV3RouterMetaData.GetAbi()

	aaveV2IncentivesABI, _ := aave.AaveIncentivesControllerMetaData.GetAbi()
	aaveV2DataProviderABI, _ := aave_data_provider.AaveDataProviderMetaData.GetAbi()
	aaveV2LendingPoolABI, _ := aave.AaveLendingPoolMetaData.GetAbi()
	aaveV2StakedTokenABI, _ := aave_staked_contract.AaveStakedContractMetaData.GetAbi()
	aaveV2StakedAbptABI, _ := aave.AaveAbptStakedMetaData.GetAbi()

	gmxVesterABI, _ := gmx.GmxVesterMetaData.GetAbi()
	gmxRewardTracker, _ := gmx.GmxRewardTrackerMetaData.GetAbi()
	gmxStakedGLP, _ := gmx.GmxStakedGlpMetaData.GetAbi()

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

	ABIs[ContractTypeAaveV2Incentives] = aaveV2IncentivesABI
	ABIs[ContractTypeAaveV2DataProvider] = aaveV2DataProviderABI
	ABIs[ContractTypeAaveV2LendingPool] = aaveV2LendingPoolABI
	ABIs[ContractTypeAaveV2StakedToken] = aaveV2StakedTokenABI
	ABIs[ContractTypeAaveV2StakedAbpt] = aaveV2StakedAbptABI

	ABIs[ContractTypeGMXVester] = gmxVesterABI
	ABIs[ContractTypeGMXRewardTracker] = gmxRewardTracker
	ABIs[ContractTypeGMXStakedGLP] = gmxStakedGLP

}
