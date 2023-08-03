// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RewardsDataTypesRewardsConfigInput is an auto generated low-level Go binding around an user-defined struct.
type RewardsDataTypesRewardsConfigInput struct {
	EmissionPerSecond *big.Int
	TotalSupply       *big.Int
	DistributionEnd   uint32
	Asset             common.Address
	Reward            common.Address
	TransferStrategy  common.Address
	RewardOracle      common.Address
}

// AaveV3RewardControllerMetaData contains all meta data concerning the AaveV3RewardController contract.
var AaveV3RewardControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsAccrued\",\"type\":\"uint256\"}],\"name\":\"Accrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEmission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEmission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDistributionEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDistributionEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"AssetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldEmissionManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEmissionManager\",\"type\":\"address\"}],\"name\":\"EmissionManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardOracle\",\"type\":\"address\"}],\"name\":\"RewardOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferStrategy\",\"type\":\"address\"}],\"name\":\"TransferStrategyInstalled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimAllRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimAllRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"claimAllRewardsToSelf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"claimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"claimRewardsToSelf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint88\",\"name\":\"emissionPerSecond\",\"type\":\"uint88\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"distributionEnd\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"contractITransferStrategyBase\",\"name\":\"transferStrategy\",\"type\":\"address\"},{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"rewardOracle\",\"type\":\"address\"}],\"internalType\":\"structRewardsDataTypes.RewardsConfigInput[]\",\"name\":\"config\",\"type\":\"tuple[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllUserRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardsList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"unclaimedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getDistributionEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmissionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getRewardOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getRewardsByAsset\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getRewardsData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getTransferStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getUserAccruedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getUserAssetIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"getUserRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"}],\"name\":\"handleAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"setClaimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newDistributionEnd\",\"type\":\"uint32\"}],\"name\":\"setDistributionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"}],\"name\":\"setEmissionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"rewards\",\"type\":\"address[]\"},{\"internalType\":\"uint88[]\",\"name\":\"newEmissionsPerSecond\",\"type\":\"uint88[]\"}],\"name\":\"setEmissionPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"rewardOracle\",\"type\":\"address\"}],\"name\":\"setRewardOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"contractITransferStrategyBase\",\"name\":\"transferStrategy\",\"type\":\"address\"}],\"name\":\"setTransferStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveV3RewardControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV3RewardControllerMetaData.ABI instead.
var AaveV3RewardControllerABI = AaveV3RewardControllerMetaData.ABI

// AaveV3RewardController is an auto generated Go binding around an Ethereum contract.
type AaveV3RewardController struct {
	AaveV3RewardControllerCaller     // Read-only binding to the contract
	AaveV3RewardControllerTransactor // Write-only binding to the contract
	AaveV3RewardControllerFilterer   // Log filterer for contract events
}

// AaveV3RewardControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV3RewardControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3RewardControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV3RewardControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3RewardControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV3RewardControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3RewardControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV3RewardControllerSession struct {
	Contract     *AaveV3RewardController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AaveV3RewardControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV3RewardControllerCallerSession struct {
	Contract *AaveV3RewardControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AaveV3RewardControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV3RewardControllerTransactorSession struct {
	Contract     *AaveV3RewardControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AaveV3RewardControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV3RewardControllerRaw struct {
	Contract *AaveV3RewardController // Generic contract binding to access the raw methods on
}

// AaveV3RewardControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV3RewardControllerCallerRaw struct {
	Contract *AaveV3RewardControllerCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV3RewardControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV3RewardControllerTransactorRaw struct {
	Contract *AaveV3RewardControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV3RewardController creates a new instance of AaveV3RewardController, bound to a specific deployed contract.
func NewAaveV3RewardController(address common.Address, backend bind.ContractBackend) (*AaveV3RewardController, error) {
	contract, err := bindAaveV3RewardController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardController{AaveV3RewardControllerCaller: AaveV3RewardControllerCaller{contract: contract}, AaveV3RewardControllerTransactor: AaveV3RewardControllerTransactor{contract: contract}, AaveV3RewardControllerFilterer: AaveV3RewardControllerFilterer{contract: contract}}, nil
}

// NewAaveV3RewardControllerCaller creates a new read-only instance of AaveV3RewardController, bound to a specific deployed contract.
func NewAaveV3RewardControllerCaller(address common.Address, caller bind.ContractCaller) (*AaveV3RewardControllerCaller, error) {
	contract, err := bindAaveV3RewardController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerCaller{contract: contract}, nil
}

// NewAaveV3RewardControllerTransactor creates a new write-only instance of AaveV3RewardController, bound to a specific deployed contract.
func NewAaveV3RewardControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV3RewardControllerTransactor, error) {
	contract, err := bindAaveV3RewardController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerTransactor{contract: contract}, nil
}

// NewAaveV3RewardControllerFilterer creates a new log filterer instance of AaveV3RewardController, bound to a specific deployed contract.
func NewAaveV3RewardControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV3RewardControllerFilterer, error) {
	contract, err := bindAaveV3RewardController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerFilterer{contract: contract}, nil
}

// bindAaveV3RewardController binds a generic wrapper to an already deployed contract.
func bindAaveV3RewardController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV3RewardControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3RewardController *AaveV3RewardControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3RewardController.Contract.AaveV3RewardControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3RewardController *AaveV3RewardControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.AaveV3RewardControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3RewardController *AaveV3RewardControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.AaveV3RewardControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3RewardController *AaveV3RewardControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3RewardController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3RewardController *AaveV3RewardControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3RewardController *AaveV3RewardControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.contract.Transact(opts, method, params...)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) REVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) REVISION() (*big.Int, error) {
	return _AaveV3RewardController.Contract.REVISION(&_AaveV3RewardController.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) REVISION() (*big.Int, error) {
	return _AaveV3RewardController.Contract.REVISION(&_AaveV3RewardController.CallOpts)
}

// GetAllUserRewards is a free data retrieval call binding the contract method 0x4c0369c3.
//
// Solidity: function getAllUserRewards(address[] assets, address user) view returns(address[] rewardsList, uint256[] unclaimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetAllUserRewards(opts *bind.CallOpts, assets []common.Address, user common.Address) (struct {
	RewardsList      []common.Address
	UnclaimedAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getAllUserRewards", assets, user)

	outstruct := new(struct {
		RewardsList      []common.Address
		UnclaimedAmounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardsList = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UnclaimedAmounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAllUserRewards is a free data retrieval call binding the contract method 0x4c0369c3.
//
// Solidity: function getAllUserRewards(address[] assets, address user) view returns(address[] rewardsList, uint256[] unclaimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetAllUserRewards(assets []common.Address, user common.Address) (struct {
	RewardsList      []common.Address
	UnclaimedAmounts []*big.Int
}, error) {
	return _AaveV3RewardController.Contract.GetAllUserRewards(&_AaveV3RewardController.CallOpts, assets, user)
}

// GetAllUserRewards is a free data retrieval call binding the contract method 0x4c0369c3.
//
// Solidity: function getAllUserRewards(address[] assets, address user) view returns(address[] rewardsList, uint256[] unclaimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetAllUserRewards(assets []common.Address, user common.Address) (struct {
	RewardsList      []common.Address
	UnclaimedAmounts []*big.Int
}, error) {
	return _AaveV3RewardController.Contract.GetAllUserRewards(&_AaveV3RewardController.CallOpts, assets, user)
}

// GetAssetDecimals is a free data retrieval call binding the contract method 0x9efd6f72.
//
// Solidity: function getAssetDecimals(address asset) view returns(uint8)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetAssetDecimals(opts *bind.CallOpts, asset common.Address) (uint8, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getAssetDecimals", asset)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAssetDecimals is a free data retrieval call binding the contract method 0x9efd6f72.
//
// Solidity: function getAssetDecimals(address asset) view returns(uint8)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetAssetDecimals(asset common.Address) (uint8, error) {
	return _AaveV3RewardController.Contract.GetAssetDecimals(&_AaveV3RewardController.CallOpts, asset)
}

// GetAssetDecimals is a free data retrieval call binding the contract method 0x9efd6f72.
//
// Solidity: function getAssetDecimals(address asset) view returns(uint8)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetAssetDecimals(asset common.Address) (uint8, error) {
	return _AaveV3RewardController.Contract.GetAssetDecimals(&_AaveV3RewardController.CallOpts, asset)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetClaimer(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getClaimer", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _AaveV3RewardController.Contract.GetClaimer(&_AaveV3RewardController.CallOpts, user)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _AaveV3RewardController.Contract.GetClaimer(&_AaveV3RewardController.CallOpts, user)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0x1b839c77.
//
// Solidity: function getDistributionEnd(address asset, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetDistributionEnd(opts *bind.CallOpts, asset common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getDistributionEnd", asset, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionEnd is a free data retrieval call binding the contract method 0x1b839c77.
//
// Solidity: function getDistributionEnd(address asset, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetDistributionEnd(asset common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetDistributionEnd(&_AaveV3RewardController.CallOpts, asset, reward)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0x1b839c77.
//
// Solidity: function getDistributionEnd(address asset, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetDistributionEnd(asset common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetDistributionEnd(&_AaveV3RewardController.CallOpts, asset, reward)
}

// GetEmissionManager is a free data retrieval call binding the contract method 0x92074b08.
//
// Solidity: function getEmissionManager() view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetEmissionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getEmissionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmissionManager is a free data retrieval call binding the contract method 0x92074b08.
//
// Solidity: function getEmissionManager() view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetEmissionManager() (common.Address, error) {
	return _AaveV3RewardController.Contract.GetEmissionManager(&_AaveV3RewardController.CallOpts)
}

// GetEmissionManager is a free data retrieval call binding the contract method 0x92074b08.
//
// Solidity: function getEmissionManager() view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetEmissionManager() (common.Address, error) {
	return _AaveV3RewardController.Contract.GetEmissionManager(&_AaveV3RewardController.CallOpts)
}

// GetRewardOracle is a free data retrieval call binding the contract method 0x2a17bf60.
//
// Solidity: function getRewardOracle(address reward) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetRewardOracle(opts *bind.CallOpts, reward common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getRewardOracle", reward)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardOracle is a free data retrieval call binding the contract method 0x2a17bf60.
//
// Solidity: function getRewardOracle(address reward) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetRewardOracle(reward common.Address) (common.Address, error) {
	return _AaveV3RewardController.Contract.GetRewardOracle(&_AaveV3RewardController.CallOpts, reward)
}

// GetRewardOracle is a free data retrieval call binding the contract method 0x2a17bf60.
//
// Solidity: function getRewardOracle(address reward) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetRewardOracle(reward common.Address) (common.Address, error) {
	return _AaveV3RewardController.Contract.GetRewardOracle(&_AaveV3RewardController.CallOpts, reward)
}

// GetRewardsByAsset is a free data retrieval call binding the contract method 0x6657732f.
//
// Solidity: function getRewardsByAsset(address asset) view returns(address[])
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetRewardsByAsset(opts *bind.CallOpts, asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getRewardsByAsset", asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardsByAsset is a free data retrieval call binding the contract method 0x6657732f.
//
// Solidity: function getRewardsByAsset(address asset) view returns(address[])
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetRewardsByAsset(asset common.Address) ([]common.Address, error) {
	return _AaveV3RewardController.Contract.GetRewardsByAsset(&_AaveV3RewardController.CallOpts, asset)
}

// GetRewardsByAsset is a free data retrieval call binding the contract method 0x6657732f.
//
// Solidity: function getRewardsByAsset(address asset) view returns(address[])
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetRewardsByAsset(asset common.Address) ([]common.Address, error) {
	return _AaveV3RewardController.Contract.GetRewardsByAsset(&_AaveV3RewardController.CallOpts, asset)
}

// GetRewardsData is a free data retrieval call binding the contract method 0x7eff4ba8.
//
// Solidity: function getRewardsData(address asset, address reward) view returns(uint256, uint256, uint256, uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetRewardsData(opts *bind.CallOpts, asset common.Address, reward common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getRewardsData", asset, reward)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetRewardsData is a free data retrieval call binding the contract method 0x7eff4ba8.
//
// Solidity: function getRewardsData(address asset, address reward) view returns(uint256, uint256, uint256, uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetRewardsData(asset common.Address, reward common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AaveV3RewardController.Contract.GetRewardsData(&_AaveV3RewardController.CallOpts, asset, reward)
}

// GetRewardsData is a free data retrieval call binding the contract method 0x7eff4ba8.
//
// Solidity: function getRewardsData(address asset, address reward) view returns(uint256, uint256, uint256, uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetRewardsData(asset common.Address, reward common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AaveV3RewardController.Contract.GetRewardsData(&_AaveV3RewardController.CallOpts, asset, reward)
}

// GetRewardsList is a free data retrieval call binding the contract method 0xb45ac1a9.
//
// Solidity: function getRewardsList() view returns(address[])
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetRewardsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getRewardsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardsList is a free data retrieval call binding the contract method 0xb45ac1a9.
//
// Solidity: function getRewardsList() view returns(address[])
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetRewardsList() ([]common.Address, error) {
	return _AaveV3RewardController.Contract.GetRewardsList(&_AaveV3RewardController.CallOpts)
}

// GetRewardsList is a free data retrieval call binding the contract method 0xb45ac1a9.
//
// Solidity: function getRewardsList() view returns(address[])
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetRewardsList() ([]common.Address, error) {
	return _AaveV3RewardController.Contract.GetRewardsList(&_AaveV3RewardController.CallOpts)
}

// GetTransferStrategy is a free data retrieval call binding the contract method 0x5f130b24.
//
// Solidity: function getTransferStrategy(address reward) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetTransferStrategy(opts *bind.CallOpts, reward common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getTransferStrategy", reward)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTransferStrategy is a free data retrieval call binding the contract method 0x5f130b24.
//
// Solidity: function getTransferStrategy(address reward) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetTransferStrategy(reward common.Address) (common.Address, error) {
	return _AaveV3RewardController.Contract.GetTransferStrategy(&_AaveV3RewardController.CallOpts, reward)
}

// GetTransferStrategy is a free data retrieval call binding the contract method 0x5f130b24.
//
// Solidity: function getTransferStrategy(address reward) view returns(address)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetTransferStrategy(reward common.Address) (common.Address, error) {
	return _AaveV3RewardController.Contract.GetTransferStrategy(&_AaveV3RewardController.CallOpts, reward)
}

// GetUserAccruedRewards is a free data retrieval call binding the contract method 0xb022418c.
//
// Solidity: function getUserAccruedRewards(address user, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetUserAccruedRewards(opts *bind.CallOpts, user common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getUserAccruedRewards", user, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAccruedRewards is a free data retrieval call binding the contract method 0xb022418c.
//
// Solidity: function getUserAccruedRewards(address user, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetUserAccruedRewards(user common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetUserAccruedRewards(&_AaveV3RewardController.CallOpts, user, reward)
}

// GetUserAccruedRewards is a free data retrieval call binding the contract method 0xb022418c.
//
// Solidity: function getUserAccruedRewards(address user, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetUserAccruedRewards(user common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetUserAccruedRewards(&_AaveV3RewardController.CallOpts, user, reward)
}

// GetUserAssetIndex is a free data retrieval call binding the contract method 0x533f542a.
//
// Solidity: function getUserAssetIndex(address user, address asset, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetUserAssetIndex(opts *bind.CallOpts, user common.Address, asset common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getUserAssetIndex", user, asset, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetIndex is a free data retrieval call binding the contract method 0x533f542a.
//
// Solidity: function getUserAssetIndex(address user, address asset, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetUserAssetIndex(user common.Address, asset common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetUserAssetIndex(&_AaveV3RewardController.CallOpts, user, asset, reward)
}

// GetUserAssetIndex is a free data retrieval call binding the contract method 0x533f542a.
//
// Solidity: function getUserAssetIndex(address user, address asset, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetUserAssetIndex(user common.Address, asset common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetUserAssetIndex(&_AaveV3RewardController.CallOpts, user, asset, reward)
}

// GetUserRewards is a free data retrieval call binding the contract method 0x70674ab9.
//
// Solidity: function getUserRewards(address[] assets, address user, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCaller) GetUserRewards(opts *bind.CallOpts, assets []common.Address, user common.Address, reward common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3RewardController.contract.Call(opts, &out, "getUserRewards", assets, user, reward)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRewards is a free data retrieval call binding the contract method 0x70674ab9.
//
// Solidity: function getUserRewards(address[] assets, address user, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) GetUserRewards(assets []common.Address, user common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetUserRewards(&_AaveV3RewardController.CallOpts, assets, user, reward)
}

// GetUserRewards is a free data retrieval call binding the contract method 0x70674ab9.
//
// Solidity: function getUserRewards(address[] assets, address user, address reward) view returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerCallerSession) GetUserRewards(assets []common.Address, user common.Address, reward common.Address) (*big.Int, error) {
	return _AaveV3RewardController.Contract.GetUserRewards(&_AaveV3RewardController.CallOpts, assets, user, reward)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xbb492bf5.
//
// Solidity: function claimAllRewards(address[] assets, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ClaimAllRewards(opts *bind.TransactOpts, assets []common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "claimAllRewards", assets, to)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xbb492bf5.
//
// Solidity: function claimAllRewards(address[] assets, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerSession) ClaimAllRewards(assets []common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimAllRewards(&_AaveV3RewardController.TransactOpts, assets, to)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0xbb492bf5.
//
// Solidity: function claimAllRewards(address[] assets, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ClaimAllRewards(assets []common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimAllRewards(&_AaveV3RewardController.TransactOpts, assets, to)
}

// ClaimAllRewardsOnBehalf is a paid mutator transaction binding the contract method 0x9ff55db9.
//
// Solidity: function claimAllRewardsOnBehalf(address[] assets, address user, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ClaimAllRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "claimAllRewardsOnBehalf", assets, user, to)
}

// ClaimAllRewardsOnBehalf is a paid mutator transaction binding the contract method 0x9ff55db9.
//
// Solidity: function claimAllRewardsOnBehalf(address[] assets, address user, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerSession) ClaimAllRewardsOnBehalf(assets []common.Address, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimAllRewardsOnBehalf(&_AaveV3RewardController.TransactOpts, assets, user, to)
}

// ClaimAllRewardsOnBehalf is a paid mutator transaction binding the contract method 0x9ff55db9.
//
// Solidity: function claimAllRewardsOnBehalf(address[] assets, address user, address to) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ClaimAllRewardsOnBehalf(assets []common.Address, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimAllRewardsOnBehalf(&_AaveV3RewardController.TransactOpts, assets, user, to)
}

// ClaimAllRewardsToSelf is a paid mutator transaction binding the contract method 0xbf90f63a.
//
// Solidity: function claimAllRewardsToSelf(address[] assets) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ClaimAllRewardsToSelf(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "claimAllRewardsToSelf", assets)
}

// ClaimAllRewardsToSelf is a paid mutator transaction binding the contract method 0xbf90f63a.
//
// Solidity: function claimAllRewardsToSelf(address[] assets) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerSession) ClaimAllRewardsToSelf(assets []common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimAllRewardsToSelf(&_AaveV3RewardController.TransactOpts, assets)
}

// ClaimAllRewardsToSelf is a paid mutator transaction binding the contract method 0xbf90f63a.
//
// Solidity: function claimAllRewardsToSelf(address[] assets) returns(address[] rewardsList, uint256[] claimedAmounts)
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ClaimAllRewardsToSelf(assets []common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimAllRewardsToSelf(&_AaveV3RewardController.TransactOpts, assets)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x236300dc.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ClaimRewards(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "claimRewards", assets, amount, to, reward)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x236300dc.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimRewards(&_AaveV3RewardController.TransactOpts, assets, amount, to, reward)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x236300dc.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimRewards(&_AaveV3RewardController.TransactOpts, assets, amount, to, reward)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x33028b99.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ClaimRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, user common.Address, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "claimRewardsOnBehalf", assets, amount, user, to, reward)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x33028b99.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimRewardsOnBehalf(&_AaveV3RewardController.TransactOpts, assets, amount, user, to, reward)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x33028b99.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimRewardsOnBehalf(&_AaveV3RewardController.TransactOpts, assets, amount, user, to, reward)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x57b89883.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ClaimRewardsToSelf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "claimRewardsToSelf", assets, amount, reward)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x57b89883.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimRewardsToSelf(&_AaveV3RewardController.TransactOpts, assets, amount, reward)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x57b89883.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount, address reward) returns(uint256)
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int, reward common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ClaimRewardsToSelf(&_AaveV3RewardController.TransactOpts, assets, amount, reward)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x955c2ad7.
//
// Solidity: function configureAssets((uint88,uint256,uint32,address,address,address,address)[] config) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) ConfigureAssets(opts *bind.TransactOpts, config []RewardsDataTypesRewardsConfigInput) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "configureAssets", config)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x955c2ad7.
//
// Solidity: function configureAssets((uint88,uint256,uint32,address,address,address,address)[] config) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) ConfigureAssets(config []RewardsDataTypesRewardsConfigInput) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ConfigureAssets(&_AaveV3RewardController.TransactOpts, config)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x955c2ad7.
//
// Solidity: function configureAssets((uint88,uint256,uint32,address,address,address,address)[] config) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) ConfigureAssets(config []RewardsDataTypesRewardsConfigInput) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.ConfigureAssets(&_AaveV3RewardController.TransactOpts, config)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) HandleAction(opts *bind.TransactOpts, user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "handleAction", user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.HandleAction(&_AaveV3RewardController.TransactOpts, user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.HandleAction(&_AaveV3RewardController.TransactOpts, user, totalSupply, userBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address emissionManager) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) Initialize(opts *bind.TransactOpts, emissionManager common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "initialize", emissionManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address emissionManager) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) Initialize(emissionManager common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.Initialize(&_AaveV3RewardController.TransactOpts, emissionManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address emissionManager) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) Initialize(emissionManager common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.Initialize(&_AaveV3RewardController.TransactOpts, emissionManager)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) SetClaimer(opts *bind.TransactOpts, user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "setClaimer", user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetClaimer(&_AaveV3RewardController.TransactOpts, user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetClaimer(&_AaveV3RewardController.TransactOpts, user, caller)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0xc5a7b538.
//
// Solidity: function setDistributionEnd(address asset, address reward, uint32 newDistributionEnd) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) SetDistributionEnd(opts *bind.TransactOpts, asset common.Address, reward common.Address, newDistributionEnd uint32) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "setDistributionEnd", asset, reward, newDistributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0xc5a7b538.
//
// Solidity: function setDistributionEnd(address asset, address reward, uint32 newDistributionEnd) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) SetDistributionEnd(asset common.Address, reward common.Address, newDistributionEnd uint32) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetDistributionEnd(&_AaveV3RewardController.TransactOpts, asset, reward, newDistributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0xc5a7b538.
//
// Solidity: function setDistributionEnd(address asset, address reward, uint32 newDistributionEnd) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) SetDistributionEnd(asset common.Address, reward common.Address, newDistributionEnd uint32) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetDistributionEnd(&_AaveV3RewardController.TransactOpts, asset, reward, newDistributionEnd)
}

// SetEmissionManager is a paid mutator transaction binding the contract method 0x4f7459d5.
//
// Solidity: function setEmissionManager(address emissionManager) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) SetEmissionManager(opts *bind.TransactOpts, emissionManager common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "setEmissionManager", emissionManager)
}

// SetEmissionManager is a paid mutator transaction binding the contract method 0x4f7459d5.
//
// Solidity: function setEmissionManager(address emissionManager) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) SetEmissionManager(emissionManager common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetEmissionManager(&_AaveV3RewardController.TransactOpts, emissionManager)
}

// SetEmissionManager is a paid mutator transaction binding the contract method 0x4f7459d5.
//
// Solidity: function setEmissionManager(address emissionManager) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) SetEmissionManager(emissionManager common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetEmissionManager(&_AaveV3RewardController.TransactOpts, emissionManager)
}

// SetEmissionPerSecond is a paid mutator transaction binding the contract method 0xf996868b.
//
// Solidity: function setEmissionPerSecond(address asset, address[] rewards, uint88[] newEmissionsPerSecond) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) SetEmissionPerSecond(opts *bind.TransactOpts, asset common.Address, rewards []common.Address, newEmissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "setEmissionPerSecond", asset, rewards, newEmissionsPerSecond)
}

// SetEmissionPerSecond is a paid mutator transaction binding the contract method 0xf996868b.
//
// Solidity: function setEmissionPerSecond(address asset, address[] rewards, uint88[] newEmissionsPerSecond) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) SetEmissionPerSecond(asset common.Address, rewards []common.Address, newEmissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetEmissionPerSecond(&_AaveV3RewardController.TransactOpts, asset, rewards, newEmissionsPerSecond)
}

// SetEmissionPerSecond is a paid mutator transaction binding the contract method 0xf996868b.
//
// Solidity: function setEmissionPerSecond(address asset, address[] rewards, uint88[] newEmissionsPerSecond) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) SetEmissionPerSecond(asset common.Address, rewards []common.Address, newEmissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetEmissionPerSecond(&_AaveV3RewardController.TransactOpts, asset, rewards, newEmissionsPerSecond)
}

// SetRewardOracle is a paid mutator transaction binding the contract method 0x5453ba10.
//
// Solidity: function setRewardOracle(address reward, address rewardOracle) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) SetRewardOracle(opts *bind.TransactOpts, reward common.Address, rewardOracle common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "setRewardOracle", reward, rewardOracle)
}

// SetRewardOracle is a paid mutator transaction binding the contract method 0x5453ba10.
//
// Solidity: function setRewardOracle(address reward, address rewardOracle) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) SetRewardOracle(reward common.Address, rewardOracle common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetRewardOracle(&_AaveV3RewardController.TransactOpts, reward, rewardOracle)
}

// SetRewardOracle is a paid mutator transaction binding the contract method 0x5453ba10.
//
// Solidity: function setRewardOracle(address reward, address rewardOracle) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) SetRewardOracle(reward common.Address, rewardOracle common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetRewardOracle(&_AaveV3RewardController.TransactOpts, reward, rewardOracle)
}

// SetTransferStrategy is a paid mutator transaction binding the contract method 0xe15ac623.
//
// Solidity: function setTransferStrategy(address reward, address transferStrategy) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactor) SetTransferStrategy(opts *bind.TransactOpts, reward common.Address, transferStrategy common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.contract.Transact(opts, "setTransferStrategy", reward, transferStrategy)
}

// SetTransferStrategy is a paid mutator transaction binding the contract method 0xe15ac623.
//
// Solidity: function setTransferStrategy(address reward, address transferStrategy) returns()
func (_AaveV3RewardController *AaveV3RewardControllerSession) SetTransferStrategy(reward common.Address, transferStrategy common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetTransferStrategy(&_AaveV3RewardController.TransactOpts, reward, transferStrategy)
}

// SetTransferStrategy is a paid mutator transaction binding the contract method 0xe15ac623.
//
// Solidity: function setTransferStrategy(address reward, address transferStrategy) returns()
func (_AaveV3RewardController *AaveV3RewardControllerTransactorSession) SetTransferStrategy(reward common.Address, transferStrategy common.Address) (*types.Transaction, error) {
	return _AaveV3RewardController.Contract.SetTransferStrategy(&_AaveV3RewardController.TransactOpts, reward, transferStrategy)
}

// AaveV3RewardControllerAccruedIterator is returned from FilterAccrued and is used to iterate over the raw logs and unpacked data for Accrued events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerAccruedIterator struct {
	Event *AaveV3RewardControllerAccrued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerAccrued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerAccrued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerAccrued represents a Accrued event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerAccrued struct {
	Asset          common.Address
	Reward         common.Address
	User           common.Address
	AssetIndex     *big.Int
	UserIndex      *big.Int
	RewardsAccrued *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrued is a free log retrieval operation binding the contract event 0x3303facd24627943a92e9dc87cfbb34b15c49b726eec3ad3487c16be9ab8efe8.
//
// Solidity: event Accrued(address indexed asset, address indexed reward, address indexed user, uint256 assetIndex, uint256 userIndex, uint256 rewardsAccrued)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterAccrued(opts *bind.FilterOpts, asset []common.Address, reward []common.Address, user []common.Address) (*AaveV3RewardControllerAccruedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "Accrued", assetRule, rewardRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerAccruedIterator{contract: _AaveV3RewardController.contract, event: "Accrued", logs: logs, sub: sub}, nil
}

// WatchAccrued is a free log subscription operation binding the contract event 0x3303facd24627943a92e9dc87cfbb34b15c49b726eec3ad3487c16be9ab8efe8.
//
// Solidity: event Accrued(address indexed asset, address indexed reward, address indexed user, uint256 assetIndex, uint256 userIndex, uint256 rewardsAccrued)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchAccrued(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerAccrued, asset []common.Address, reward []common.Address, user []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "Accrued", assetRule, rewardRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerAccrued)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "Accrued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccrued is a log parse operation binding the contract event 0x3303facd24627943a92e9dc87cfbb34b15c49b726eec3ad3487c16be9ab8efe8.
//
// Solidity: event Accrued(address indexed asset, address indexed reward, address indexed user, uint256 assetIndex, uint256 userIndex, uint256 rewardsAccrued)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseAccrued(log types.Log) (*AaveV3RewardControllerAccrued, error) {
	event := new(AaveV3RewardControllerAccrued)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "Accrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3RewardControllerAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerAssetConfigUpdatedIterator struct {
	Event *AaveV3RewardControllerAssetConfigUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerAssetConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerAssetConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerAssetConfigUpdated represents a AssetConfigUpdated event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerAssetConfigUpdated struct {
	Asset              common.Address
	Reward             common.Address
	OldEmission        *big.Int
	NewEmission        *big.Int
	OldDistributionEnd *big.Int
	NewDistributionEnd *big.Int
	AssetIndex         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0xac1777479f07f3e7c34da8402139d54027a6a260caaae168bdee825ca5580dc5.
//
// Solidity: event AssetConfigUpdated(address indexed asset, address indexed reward, uint256 oldEmission, uint256 newEmission, uint256 oldDistributionEnd, uint256 newDistributionEnd, uint256 assetIndex)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, asset []common.Address, reward []common.Address) (*AaveV3RewardControllerAssetConfigUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "AssetConfigUpdated", assetRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerAssetConfigUpdatedIterator{contract: _AaveV3RewardController.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0xac1777479f07f3e7c34da8402139d54027a6a260caaae168bdee825ca5580dc5.
//
// Solidity: event AssetConfigUpdated(address indexed asset, address indexed reward, uint256 oldEmission, uint256 newEmission, uint256 oldDistributionEnd, uint256 newDistributionEnd, uint256 assetIndex)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerAssetConfigUpdated, asset []common.Address, reward []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "AssetConfigUpdated", assetRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerAssetConfigUpdated)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetConfigUpdated is a log parse operation binding the contract event 0xac1777479f07f3e7c34da8402139d54027a6a260caaae168bdee825ca5580dc5.
//
// Solidity: event AssetConfigUpdated(address indexed asset, address indexed reward, uint256 oldEmission, uint256 newEmission, uint256 oldDistributionEnd, uint256 newDistributionEnd, uint256 assetIndex)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseAssetConfigUpdated(log types.Log) (*AaveV3RewardControllerAssetConfigUpdated, error) {
	event := new(AaveV3RewardControllerAssetConfigUpdated)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3RewardControllerClaimerSetIterator is returned from FilterClaimerSet and is used to iterate over the raw logs and unpacked data for ClaimerSet events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerClaimerSetIterator struct {
	Event *AaveV3RewardControllerClaimerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerClaimerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerClaimerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerClaimerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerClaimerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerClaimerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerClaimerSet represents a ClaimerSet event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerClaimerSet struct {
	User    common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimerSet is a free log retrieval operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterClaimerSet(opts *bind.FilterOpts, user []common.Address, claimer []common.Address) (*AaveV3RewardControllerClaimerSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerClaimerSetIterator{contract: _AaveV3RewardController.contract, event: "ClaimerSet", logs: logs, sub: sub}, nil
}

// WatchClaimerSet is a free log subscription operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchClaimerSet(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerClaimerSet, user []common.Address, claimer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerClaimerSet)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimerSet is a log parse operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseClaimerSet(log types.Log) (*AaveV3RewardControllerClaimerSet, error) {
	event := new(AaveV3RewardControllerClaimerSet)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3RewardControllerEmissionManagerUpdatedIterator is returned from FilterEmissionManagerUpdated and is used to iterate over the raw logs and unpacked data for EmissionManagerUpdated events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerEmissionManagerUpdatedIterator struct {
	Event *AaveV3RewardControllerEmissionManagerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerEmissionManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerEmissionManagerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerEmissionManagerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerEmissionManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerEmissionManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerEmissionManagerUpdated represents a EmissionManagerUpdated event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerEmissionManagerUpdated struct {
	OldEmissionManager common.Address
	NewEmissionManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEmissionManagerUpdated is a free log retrieval operation binding the contract event 0x53271355c244f99d37f622c90fe574eb7c35c7b3548ea276cf9b5b11c601605e.
//
// Solidity: event EmissionManagerUpdated(address indexed oldEmissionManager, address indexed newEmissionManager)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterEmissionManagerUpdated(opts *bind.FilterOpts, oldEmissionManager []common.Address, newEmissionManager []common.Address) (*AaveV3RewardControllerEmissionManagerUpdatedIterator, error) {

	var oldEmissionManagerRule []interface{}
	for _, oldEmissionManagerItem := range oldEmissionManager {
		oldEmissionManagerRule = append(oldEmissionManagerRule, oldEmissionManagerItem)
	}
	var newEmissionManagerRule []interface{}
	for _, newEmissionManagerItem := range newEmissionManager {
		newEmissionManagerRule = append(newEmissionManagerRule, newEmissionManagerItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "EmissionManagerUpdated", oldEmissionManagerRule, newEmissionManagerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerEmissionManagerUpdatedIterator{contract: _AaveV3RewardController.contract, event: "EmissionManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchEmissionManagerUpdated is a free log subscription operation binding the contract event 0x53271355c244f99d37f622c90fe574eb7c35c7b3548ea276cf9b5b11c601605e.
//
// Solidity: event EmissionManagerUpdated(address indexed oldEmissionManager, address indexed newEmissionManager)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchEmissionManagerUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerEmissionManagerUpdated, oldEmissionManager []common.Address, newEmissionManager []common.Address) (event.Subscription, error) {

	var oldEmissionManagerRule []interface{}
	for _, oldEmissionManagerItem := range oldEmissionManager {
		oldEmissionManagerRule = append(oldEmissionManagerRule, oldEmissionManagerItem)
	}
	var newEmissionManagerRule []interface{}
	for _, newEmissionManagerItem := range newEmissionManager {
		newEmissionManagerRule = append(newEmissionManagerRule, newEmissionManagerItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "EmissionManagerUpdated", oldEmissionManagerRule, newEmissionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerEmissionManagerUpdated)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "EmissionManagerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmissionManagerUpdated is a log parse operation binding the contract event 0x53271355c244f99d37f622c90fe574eb7c35c7b3548ea276cf9b5b11c601605e.
//
// Solidity: event EmissionManagerUpdated(address indexed oldEmissionManager, address indexed newEmissionManager)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseEmissionManagerUpdated(log types.Log) (*AaveV3RewardControllerEmissionManagerUpdated, error) {
	event := new(AaveV3RewardControllerEmissionManagerUpdated)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "EmissionManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3RewardControllerRewardOracleUpdatedIterator is returned from FilterRewardOracleUpdated and is used to iterate over the raw logs and unpacked data for RewardOracleUpdated events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerRewardOracleUpdatedIterator struct {
	Event *AaveV3RewardControllerRewardOracleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerRewardOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerRewardOracleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerRewardOracleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerRewardOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerRewardOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerRewardOracleUpdated represents a RewardOracleUpdated event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerRewardOracleUpdated struct {
	Reward       common.Address
	RewardOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardOracleUpdated is a free log retrieval operation binding the contract event 0x1a1cd5483e52e60b9ff7f3b9d1db3bbd9e9d21c6324ad3a8c79dba9b75e62f4d.
//
// Solidity: event RewardOracleUpdated(address indexed reward, address indexed rewardOracle)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterRewardOracleUpdated(opts *bind.FilterOpts, reward []common.Address, rewardOracle []common.Address) (*AaveV3RewardControllerRewardOracleUpdatedIterator, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var rewardOracleRule []interface{}
	for _, rewardOracleItem := range rewardOracle {
		rewardOracleRule = append(rewardOracleRule, rewardOracleItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "RewardOracleUpdated", rewardRule, rewardOracleRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerRewardOracleUpdatedIterator{contract: _AaveV3RewardController.contract, event: "RewardOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardOracleUpdated is a free log subscription operation binding the contract event 0x1a1cd5483e52e60b9ff7f3b9d1db3bbd9e9d21c6324ad3a8c79dba9b75e62f4d.
//
// Solidity: event RewardOracleUpdated(address indexed reward, address indexed rewardOracle)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchRewardOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerRewardOracleUpdated, reward []common.Address, rewardOracle []common.Address) (event.Subscription, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var rewardOracleRule []interface{}
	for _, rewardOracleItem := range rewardOracle {
		rewardOracleRule = append(rewardOracleRule, rewardOracleItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "RewardOracleUpdated", rewardRule, rewardOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerRewardOracleUpdated)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "RewardOracleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardOracleUpdated is a log parse operation binding the contract event 0x1a1cd5483e52e60b9ff7f3b9d1db3bbd9e9d21c6324ad3a8c79dba9b75e62f4d.
//
// Solidity: event RewardOracleUpdated(address indexed reward, address indexed rewardOracle)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseRewardOracleUpdated(log types.Log) (*AaveV3RewardControllerRewardOracleUpdated, error) {
	event := new(AaveV3RewardControllerRewardOracleUpdated)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "RewardOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3RewardControllerRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerRewardsClaimedIterator struct {
	Event *AaveV3RewardControllerRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerRewardsClaimed represents a RewardsClaimed event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerRewardsClaimed struct {
	User    common.Address
	Reward  common.Address
	To      common.Address
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xc052130bc4ef84580db505783484b067ea8b71b3bca78a7e12db7aea8658f004.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed reward, address indexed to, address claimer, uint256 amount)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, reward []common.Address, to []common.Address) (*AaveV3RewardControllerRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "RewardsClaimed", userRule, rewardRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerRewardsClaimedIterator{contract: _AaveV3RewardController.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xc052130bc4ef84580db505783484b067ea8b71b3bca78a7e12db7aea8658f004.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed reward, address indexed to, address claimer, uint256 amount)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerRewardsClaimed, user []common.Address, reward []common.Address, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "RewardsClaimed", userRule, rewardRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerRewardsClaimed)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0xc052130bc4ef84580db505783484b067ea8b71b3bca78a7e12db7aea8658f004.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed reward, address indexed to, address claimer, uint256 amount)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseRewardsClaimed(log types.Log) (*AaveV3RewardControllerRewardsClaimed, error) {
	event := new(AaveV3RewardControllerRewardsClaimed)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3RewardControllerTransferStrategyInstalledIterator is returned from FilterTransferStrategyInstalled and is used to iterate over the raw logs and unpacked data for TransferStrategyInstalled events raised by the AaveV3RewardController contract.
type AaveV3RewardControllerTransferStrategyInstalledIterator struct {
	Event *AaveV3RewardControllerTransferStrategyInstalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3RewardControllerTransferStrategyInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3RewardControllerTransferStrategyInstalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3RewardControllerTransferStrategyInstalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3RewardControllerTransferStrategyInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3RewardControllerTransferStrategyInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3RewardControllerTransferStrategyInstalled represents a TransferStrategyInstalled event raised by the AaveV3RewardController contract.
type AaveV3RewardControllerTransferStrategyInstalled struct {
	Reward           common.Address
	TransferStrategy common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTransferStrategyInstalled is a free log retrieval operation binding the contract event 0x8ca1d928f1d72493a6b78c4f74aabde976bc37ffe2570f2a1ce5a8abd3dde0aa.
//
// Solidity: event TransferStrategyInstalled(address indexed reward, address indexed transferStrategy)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) FilterTransferStrategyInstalled(opts *bind.FilterOpts, reward []common.Address, transferStrategy []common.Address) (*AaveV3RewardControllerTransferStrategyInstalledIterator, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var transferStrategyRule []interface{}
	for _, transferStrategyItem := range transferStrategy {
		transferStrategyRule = append(transferStrategyRule, transferStrategyItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.FilterLogs(opts, "TransferStrategyInstalled", rewardRule, transferStrategyRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3RewardControllerTransferStrategyInstalledIterator{contract: _AaveV3RewardController.contract, event: "TransferStrategyInstalled", logs: logs, sub: sub}, nil
}

// WatchTransferStrategyInstalled is a free log subscription operation binding the contract event 0x8ca1d928f1d72493a6b78c4f74aabde976bc37ffe2570f2a1ce5a8abd3dde0aa.
//
// Solidity: event TransferStrategyInstalled(address indexed reward, address indexed transferStrategy)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) WatchTransferStrategyInstalled(opts *bind.WatchOpts, sink chan<- *AaveV3RewardControllerTransferStrategyInstalled, reward []common.Address, transferStrategy []common.Address) (event.Subscription, error) {

	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}
	var transferStrategyRule []interface{}
	for _, transferStrategyItem := range transferStrategy {
		transferStrategyRule = append(transferStrategyRule, transferStrategyItem)
	}

	logs, sub, err := _AaveV3RewardController.contract.WatchLogs(opts, "TransferStrategyInstalled", rewardRule, transferStrategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3RewardControllerTransferStrategyInstalled)
				if err := _AaveV3RewardController.contract.UnpackLog(event, "TransferStrategyInstalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferStrategyInstalled is a log parse operation binding the contract event 0x8ca1d928f1d72493a6b78c4f74aabde976bc37ffe2570f2a1ce5a8abd3dde0aa.
//
// Solidity: event TransferStrategyInstalled(address indexed reward, address indexed transferStrategy)
func (_AaveV3RewardController *AaveV3RewardControllerFilterer) ParseTransferStrategyInstalled(log types.Log) (*AaveV3RewardControllerTransferStrategyInstalled, error) {
	event := new(AaveV3RewardControllerTransferStrategyInstalled)
	if err := _AaveV3RewardController.contract.UnpackLog(event, "TransferStrategyInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
