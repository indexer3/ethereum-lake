// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package animal_farm

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

// DogPoundAutoPoolUserInfo is an auto generated low-level Go binding around an user-defined struct.
type DogPoundAutoPoolUserInfo struct {
	Amount             *big.Int
	LpMask             *big.Int
	PigsClaimedTotal   *big.Int
	LastRmsClaimed     *big.Int
	LpDebt             *big.Int
	TotalLPCollected   *big.Int
	TotalPigsCollected *big.Int
}

// AnimalFarmMetaData contains all meta data concerning the AnimalFarm contract.
var AnimalFarmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BnbLiquidateThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOGS_BNB_MC_PID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DogPoundManger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DogsExchangeHelper\",\"outputs\":[{\"internalType\":\"contractIDogsExchangeHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DogsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Dogs_BNB_LpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MClocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MClockedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MasterchefPigs\",\"outputs\":[{\"internalType\":\"contractIMasterchefPigs\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PancakeRouter\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PigsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ad\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"addInitAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"allowCompound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"busdCurrencyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"changeUpdateInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimLpTokensAndPigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"claimPigsHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dogsBnbPath\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historyInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rms\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quant\",\"type\":\"uint256\"}],\"name\":\"increasePigsBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"initCompounders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initMCStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"temp1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temp2\",\"type\":\"uint256\"}],\"name\":\"initStakeMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pigsClaimedTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRmsClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLPCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPigsCollected\",\"type\":\"uint256\"}],\"internalType\":\"structDogPoundAutoPool.UserInfo[]\",\"name\":\"_info\",\"type\":\"tuple[]\"}],\"name\":\"initializeMd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDogPoundAutoPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"initializeU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeUnpaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDogPoundAutoPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"histlen\",\"type\":\"uint256\"}],\"name\":\"initializeVariables\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPigsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDogPoundManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpRoundMask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpRoundMasktemp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerNotLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseInitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"pendingLpRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingLp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"pendingLpRewardsInternal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingLp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingPigsRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"pendingPigsRewardsHelper\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_pendingPigs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pendingLp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setDogPoundManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeSinceLastCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDogsStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLPCollected\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLPstakedTemp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLpStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThrehshold\",\"type\":\"uint256\"}],\"name\":\"updateBnbLiqThreshhold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateDogsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressDogs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressLpBNB\",\"type\":\"address\"}],\"name\":\"updateDogsAndLPAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPid\",\"type\":\"uint256\"}],\"name\":\"updateDogsBnBPID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateDogsExchanceHelperAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateMasterchefPigsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updatePigsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pigsClaimedTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRmsClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLPCollected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPigsCollected\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wbnbCurrencyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AnimalFarmABI is the input ABI used to generate the binding from.
// Deprecated: Use AnimalFarmMetaData.ABI instead.
var AnimalFarmABI = AnimalFarmMetaData.ABI

// AnimalFarm is an auto generated Go binding around an Ethereum contract.
type AnimalFarm struct {
	AnimalFarmCaller     // Read-only binding to the contract
	AnimalFarmTransactor // Write-only binding to the contract
	AnimalFarmFilterer   // Log filterer for contract events
}

// AnimalFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnimalFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnimalFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnimalFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnimalFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnimalFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnimalFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnimalFarmSession struct {
	Contract     *AnimalFarm       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnimalFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnimalFarmCallerSession struct {
	Contract *AnimalFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AnimalFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnimalFarmTransactorSession struct {
	Contract     *AnimalFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AnimalFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnimalFarmRaw struct {
	Contract *AnimalFarm // Generic contract binding to access the raw methods on
}

// AnimalFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnimalFarmCallerRaw struct {
	Contract *AnimalFarmCaller // Generic read-only contract binding to access the raw methods on
}

// AnimalFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnimalFarmTransactorRaw struct {
	Contract *AnimalFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnimalFarm creates a new instance of AnimalFarm, bound to a specific deployed contract.
func NewAnimalFarm(address common.Address, backend bind.ContractBackend) (*AnimalFarm, error) {
	contract, err := bindAnimalFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnimalFarm{AnimalFarmCaller: AnimalFarmCaller{contract: contract}, AnimalFarmTransactor: AnimalFarmTransactor{contract: contract}, AnimalFarmFilterer: AnimalFarmFilterer{contract: contract}}, nil
}

// NewAnimalFarmCaller creates a new read-only instance of AnimalFarm, bound to a specific deployed contract.
func NewAnimalFarmCaller(address common.Address, caller bind.ContractCaller) (*AnimalFarmCaller, error) {
	contract, err := bindAnimalFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnimalFarmCaller{contract: contract}, nil
}

// NewAnimalFarmTransactor creates a new write-only instance of AnimalFarm, bound to a specific deployed contract.
func NewAnimalFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*AnimalFarmTransactor, error) {
	contract, err := bindAnimalFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnimalFarmTransactor{contract: contract}, nil
}

// NewAnimalFarmFilterer creates a new log filterer instance of AnimalFarm, bound to a specific deployed contract.
func NewAnimalFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*AnimalFarmFilterer, error) {
	contract, err := bindAnimalFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnimalFarmFilterer{contract: contract}, nil
}

// bindAnimalFarm binds a generic wrapper to an already deployed contract.
func bindAnimalFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnimalFarmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnimalFarm *AnimalFarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnimalFarm.Contract.AnimalFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnimalFarm *AnimalFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.Contract.AnimalFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnimalFarm *AnimalFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnimalFarm.Contract.AnimalFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnimalFarm *AnimalFarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnimalFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnimalFarm *AnimalFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnimalFarm *AnimalFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnimalFarm.Contract.contract.Transact(opts, method, params...)
}

// BnbLiquidateThreshold is a free data retrieval call binding the contract method 0x453978fd.
//
// Solidity: function BnbLiquidateThreshold() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) BnbLiquidateThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "BnbLiquidateThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BnbLiquidateThreshold is a free data retrieval call binding the contract method 0x453978fd.
//
// Solidity: function BnbLiquidateThreshold() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) BnbLiquidateThreshold() (*big.Int, error) {
	return _AnimalFarm.Contract.BnbLiquidateThreshold(&_AnimalFarm.CallOpts)
}

// BnbLiquidateThreshold is a free data retrieval call binding the contract method 0x453978fd.
//
// Solidity: function BnbLiquidateThreshold() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) BnbLiquidateThreshold() (*big.Int, error) {
	return _AnimalFarm.Contract.BnbLiquidateThreshold(&_AnimalFarm.CallOpts)
}

// DOGSBNBMCPID is a free data retrieval call binding the contract method 0x61042ec4.
//
// Solidity: function DOGS_BNB_MC_PID() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) DOGSBNBMCPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "DOGS_BNB_MC_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOGSBNBMCPID is a free data retrieval call binding the contract method 0x61042ec4.
//
// Solidity: function DOGS_BNB_MC_PID() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) DOGSBNBMCPID() (*big.Int, error) {
	return _AnimalFarm.Contract.DOGSBNBMCPID(&_AnimalFarm.CallOpts)
}

// DOGSBNBMCPID is a free data retrieval call binding the contract method 0x61042ec4.
//
// Solidity: function DOGS_BNB_MC_PID() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) DOGSBNBMCPID() (*big.Int, error) {
	return _AnimalFarm.Contract.DOGSBNBMCPID(&_AnimalFarm.CallOpts)
}

// DogPoundManger is a free data retrieval call binding the contract method 0x7443b671.
//
// Solidity: function DogPoundManger() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) DogPoundManger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "DogPoundManger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogPoundManger is a free data retrieval call binding the contract method 0x7443b671.
//
// Solidity: function DogPoundManger() view returns(address)
func (_AnimalFarm *AnimalFarmSession) DogPoundManger() (common.Address, error) {
	return _AnimalFarm.Contract.DogPoundManger(&_AnimalFarm.CallOpts)
}

// DogPoundManger is a free data retrieval call binding the contract method 0x7443b671.
//
// Solidity: function DogPoundManger() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) DogPoundManger() (common.Address, error) {
	return _AnimalFarm.Contract.DogPoundManger(&_AnimalFarm.CallOpts)
}

// DogsExchangeHelper is a free data retrieval call binding the contract method 0x91948063.
//
// Solidity: function DogsExchangeHelper() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) DogsExchangeHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "DogsExchangeHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogsExchangeHelper is a free data retrieval call binding the contract method 0x91948063.
//
// Solidity: function DogsExchangeHelper() view returns(address)
func (_AnimalFarm *AnimalFarmSession) DogsExchangeHelper() (common.Address, error) {
	return _AnimalFarm.Contract.DogsExchangeHelper(&_AnimalFarm.CallOpts)
}

// DogsExchangeHelper is a free data retrieval call binding the contract method 0x91948063.
//
// Solidity: function DogsExchangeHelper() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) DogsExchangeHelper() (common.Address, error) {
	return _AnimalFarm.Contract.DogsExchangeHelper(&_AnimalFarm.CallOpts)
}

// DogsToken is a free data retrieval call binding the contract method 0x7e6151c9.
//
// Solidity: function DogsToken() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) DogsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "DogsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogsToken is a free data retrieval call binding the contract method 0x7e6151c9.
//
// Solidity: function DogsToken() view returns(address)
func (_AnimalFarm *AnimalFarmSession) DogsToken() (common.Address, error) {
	return _AnimalFarm.Contract.DogsToken(&_AnimalFarm.CallOpts)
}

// DogsToken is a free data retrieval call binding the contract method 0x7e6151c9.
//
// Solidity: function DogsToken() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) DogsToken() (common.Address, error) {
	return _AnimalFarm.Contract.DogsToken(&_AnimalFarm.CallOpts)
}

// DogsBNBLpToken is a free data retrieval call binding the contract method 0x3a80f3a6.
//
// Solidity: function Dogs_BNB_LpToken() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) DogsBNBLpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "Dogs_BNB_LpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogsBNBLpToken is a free data retrieval call binding the contract method 0x3a80f3a6.
//
// Solidity: function Dogs_BNB_LpToken() view returns(address)
func (_AnimalFarm *AnimalFarmSession) DogsBNBLpToken() (common.Address, error) {
	return _AnimalFarm.Contract.DogsBNBLpToken(&_AnimalFarm.CallOpts)
}

// DogsBNBLpToken is a free data retrieval call binding the contract method 0x3a80f3a6.
//
// Solidity: function Dogs_BNB_LpToken() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) DogsBNBLpToken() (common.Address, error) {
	return _AnimalFarm.Contract.DogsBNBLpToken(&_AnimalFarm.CallOpts)
}

// MClocked is a free data retrieval call binding the contract method 0x0e6e6dec.
//
// Solidity: function MClocked() view returns(bool)
func (_AnimalFarm *AnimalFarmCaller) MClocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "MClocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MClocked is a free data retrieval call binding the contract method 0x0e6e6dec.
//
// Solidity: function MClocked() view returns(bool)
func (_AnimalFarm *AnimalFarmSession) MClocked() (bool, error) {
	return _AnimalFarm.Contract.MClocked(&_AnimalFarm.CallOpts)
}

// MClocked is a free data retrieval call binding the contract method 0x0e6e6dec.
//
// Solidity: function MClocked() view returns(bool)
func (_AnimalFarm *AnimalFarmCallerSession) MClocked() (bool, error) {
	return _AnimalFarm.Contract.MClocked(&_AnimalFarm.CallOpts)
}

// MasterchefPigs is a free data retrieval call binding the contract method 0xa4364b03.
//
// Solidity: function MasterchefPigs() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) MasterchefPigs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "MasterchefPigs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterchefPigs is a free data retrieval call binding the contract method 0xa4364b03.
//
// Solidity: function MasterchefPigs() view returns(address)
func (_AnimalFarm *AnimalFarmSession) MasterchefPigs() (common.Address, error) {
	return _AnimalFarm.Contract.MasterchefPigs(&_AnimalFarm.CallOpts)
}

// MasterchefPigs is a free data retrieval call binding the contract method 0xa4364b03.
//
// Solidity: function MasterchefPigs() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) MasterchefPigs() (common.Address, error) {
	return _AnimalFarm.Contract.MasterchefPigs(&_AnimalFarm.CallOpts)
}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) PancakeRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "PancakeRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_AnimalFarm *AnimalFarmSession) PancakeRouter() (common.Address, error) {
	return _AnimalFarm.Contract.PancakeRouter(&_AnimalFarm.CallOpts)
}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) PancakeRouter() (common.Address, error) {
	return _AnimalFarm.Contract.PancakeRouter(&_AnimalFarm.CallOpts)
}

// PigsToken is a free data retrieval call binding the contract method 0x0e3d6bfa.
//
// Solidity: function PigsToken() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) PigsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "PigsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PigsToken is a free data retrieval call binding the contract method 0x0e3d6bfa.
//
// Solidity: function PigsToken() view returns(address)
func (_AnimalFarm *AnimalFarmSession) PigsToken() (common.Address, error) {
	return _AnimalFarm.Contract.PigsToken(&_AnimalFarm.CallOpts)
}

// PigsToken is a free data retrieval call binding the contract method 0x0e3d6bfa.
//
// Solidity: function PigsToken() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) PigsToken() (common.Address, error) {
	return _AnimalFarm.Contract.PigsToken(&_AnimalFarm.CallOpts)
}

// BusdCurrencyAddress is a free data retrieval call binding the contract method 0x03ba7f66.
//
// Solidity: function busdCurrencyAddress() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) BusdCurrencyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "busdCurrencyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BusdCurrencyAddress is a free data retrieval call binding the contract method 0x03ba7f66.
//
// Solidity: function busdCurrencyAddress() view returns(address)
func (_AnimalFarm *AnimalFarmSession) BusdCurrencyAddress() (common.Address, error) {
	return _AnimalFarm.Contract.BusdCurrencyAddress(&_AnimalFarm.CallOpts)
}

// BusdCurrencyAddress is a free data retrieval call binding the contract method 0x03ba7f66.
//
// Solidity: function busdCurrencyAddress() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) BusdCurrencyAddress() (common.Address, error) {
	return _AnimalFarm.Contract.BusdCurrencyAddress(&_AnimalFarm.CallOpts)
}

// DogsBnbPath is a free data retrieval call binding the contract method 0xd5b4127f.
//
// Solidity: function dogsBnbPath(uint256 ) view returns(address)
func (_AnimalFarm *AnimalFarmCaller) DogsBnbPath(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "dogsBnbPath", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogsBnbPath is a free data retrieval call binding the contract method 0xd5b4127f.
//
// Solidity: function dogsBnbPath(uint256 ) view returns(address)
func (_AnimalFarm *AnimalFarmSession) DogsBnbPath(arg0 *big.Int) (common.Address, error) {
	return _AnimalFarm.Contract.DogsBnbPath(&_AnimalFarm.CallOpts, arg0)
}

// DogsBnbPath is a free data retrieval call binding the contract method 0xd5b4127f.
//
// Solidity: function dogsBnbPath(uint256 ) view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) DogsBnbPath(arg0 *big.Int) (common.Address, error) {
	return _AnimalFarm.Contract.DogsBnbPath(&_AnimalFarm.CallOpts, arg0)
}

// HistoryInfo is a free data retrieval call binding the contract method 0x8fa6a2ff.
//
// Solidity: function historyInfo(uint256 ) view returns(uint256 pps, uint256 rms)
func (_AnimalFarm *AnimalFarmCaller) HistoryInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Pps *big.Int
	Rms *big.Int
}, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "historyInfo", arg0)

	outstruct := new(struct {
		Pps *big.Int
		Rms *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pps = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rms = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HistoryInfo is a free data retrieval call binding the contract method 0x8fa6a2ff.
//
// Solidity: function historyInfo(uint256 ) view returns(uint256 pps, uint256 rms)
func (_AnimalFarm *AnimalFarmSession) HistoryInfo(arg0 *big.Int) (struct {
	Pps *big.Int
	Rms *big.Int
}, error) {
	return _AnimalFarm.Contract.HistoryInfo(&_AnimalFarm.CallOpts, arg0)
}

// HistoryInfo is a free data retrieval call binding the contract method 0x8fa6a2ff.
//
// Solidity: function historyInfo(uint256 ) view returns(uint256 pps, uint256 rms)
func (_AnimalFarm *AnimalFarmCallerSession) HistoryInfo(arg0 *big.Int) (struct {
	Pps *big.Int
	Rms *big.Int
}, error) {
	return _AnimalFarm.Contract.HistoryInfo(&_AnimalFarm.CallOpts, arg0)
}

// InitializeUnpaused is a free data retrieval call binding the contract method 0x17fe9f65.
//
// Solidity: function initializeUnpaused() view returns(bool)
func (_AnimalFarm *AnimalFarmCaller) InitializeUnpaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "initializeUnpaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InitializeUnpaused is a free data retrieval call binding the contract method 0x17fe9f65.
//
// Solidity: function initializeUnpaused() view returns(bool)
func (_AnimalFarm *AnimalFarmSession) InitializeUnpaused() (bool, error) {
	return _AnimalFarm.Contract.InitializeUnpaused(&_AnimalFarm.CallOpts)
}

// InitializeUnpaused is a free data retrieval call binding the contract method 0x17fe9f65.
//
// Solidity: function initializeUnpaused() view returns(bool)
func (_AnimalFarm *AnimalFarmCallerSession) InitializeUnpaused() (bool, error) {
	return _AnimalFarm.Contract.InitializeUnpaused(&_AnimalFarm.CallOpts)
}

// LastPigsBalance is a free data retrieval call binding the contract method 0xb1904c05.
//
// Solidity: function lastPigsBalance() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) LastPigsBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "lastPigsBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPigsBalance is a free data retrieval call binding the contract method 0xb1904c05.
//
// Solidity: function lastPigsBalance() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) LastPigsBalance() (*big.Int, error) {
	return _AnimalFarm.Contract.LastPigsBalance(&_AnimalFarm.CallOpts)
}

// LastPigsBalance is a free data retrieval call binding the contract method 0xb1904c05.
//
// Solidity: function lastPigsBalance() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) LastPigsBalance() (*big.Int, error) {
	return _AnimalFarm.Contract.LastPigsBalance(&_AnimalFarm.CallOpts)
}

// LpRoundMask is a free data retrieval call binding the contract method 0x9d395a1a.
//
// Solidity: function lpRoundMask() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) LpRoundMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "lpRoundMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpRoundMask is a free data retrieval call binding the contract method 0x9d395a1a.
//
// Solidity: function lpRoundMask() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) LpRoundMask() (*big.Int, error) {
	return _AnimalFarm.Contract.LpRoundMask(&_AnimalFarm.CallOpts)
}

// LpRoundMask is a free data retrieval call binding the contract method 0x9d395a1a.
//
// Solidity: function lpRoundMask() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) LpRoundMask() (*big.Int, error) {
	return _AnimalFarm.Contract.LpRoundMask(&_AnimalFarm.CallOpts)
}

// LpRoundMasktemp is a free data retrieval call binding the contract method 0x136f60ee.
//
// Solidity: function lpRoundMasktemp() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) LpRoundMasktemp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "lpRoundMasktemp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpRoundMasktemp is a free data retrieval call binding the contract method 0x136f60ee.
//
// Solidity: function lpRoundMasktemp() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) LpRoundMasktemp() (*big.Int, error) {
	return _AnimalFarm.Contract.LpRoundMasktemp(&_AnimalFarm.CallOpts)
}

// LpRoundMasktemp is a free data retrieval call binding the contract method 0x136f60ee.
//
// Solidity: function lpRoundMasktemp() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) LpRoundMasktemp() (*big.Int, error) {
	return _AnimalFarm.Contract.LpRoundMasktemp(&_AnimalFarm.CallOpts)
}

// ManagerNotLocked is a free data retrieval call binding the contract method 0x20a7a900.
//
// Solidity: function managerNotLocked() view returns(bool)
func (_AnimalFarm *AnimalFarmCaller) ManagerNotLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "managerNotLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ManagerNotLocked is a free data retrieval call binding the contract method 0x20a7a900.
//
// Solidity: function managerNotLocked() view returns(bool)
func (_AnimalFarm *AnimalFarmSession) ManagerNotLocked() (bool, error) {
	return _AnimalFarm.Contract.ManagerNotLocked(&_AnimalFarm.CallOpts)
}

// ManagerNotLocked is a free data retrieval call binding the contract method 0x20a7a900.
//
// Solidity: function managerNotLocked() view returns(bool)
func (_AnimalFarm *AnimalFarmCallerSession) ManagerNotLocked() (bool, error) {
	return _AnimalFarm.Contract.ManagerNotLocked(&_AnimalFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnimalFarm *AnimalFarmSession) Owner() (common.Address, error) {
	return _AnimalFarm.Contract.Owner(&_AnimalFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) Owner() (common.Address, error) {
	return _AnimalFarm.Contract.Owner(&_AnimalFarm.CallOpts)
}

// PendingLpRewards is a free data retrieval call binding the contract method 0x53475567.
//
// Solidity: function pendingLpRewards(address _userAddress) view returns(uint256 pendingLp)
func (_AnimalFarm *AnimalFarmCaller) PendingLpRewards(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "pendingLpRewards", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingLpRewards is a free data retrieval call binding the contract method 0x53475567.
//
// Solidity: function pendingLpRewards(address _userAddress) view returns(uint256 pendingLp)
func (_AnimalFarm *AnimalFarmSession) PendingLpRewards(_userAddress common.Address) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingLpRewards(&_AnimalFarm.CallOpts, _userAddress)
}

// PendingLpRewards is a free data retrieval call binding the contract method 0x53475567.
//
// Solidity: function pendingLpRewards(address _userAddress) view returns(uint256 pendingLp)
func (_AnimalFarm *AnimalFarmCallerSession) PendingLpRewards(_userAddress common.Address) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingLpRewards(&_AnimalFarm.CallOpts, _userAddress)
}

// PendingLpRewardsInternal is a free data retrieval call binding the contract method 0x0198cf2b.
//
// Solidity: function pendingLpRewardsInternal(address _userAddress) view returns(uint256 pendingLp)
func (_AnimalFarm *AnimalFarmCaller) PendingLpRewardsInternal(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "pendingLpRewardsInternal", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingLpRewardsInternal is a free data retrieval call binding the contract method 0x0198cf2b.
//
// Solidity: function pendingLpRewardsInternal(address _userAddress) view returns(uint256 pendingLp)
func (_AnimalFarm *AnimalFarmSession) PendingLpRewardsInternal(_userAddress common.Address) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingLpRewardsInternal(&_AnimalFarm.CallOpts, _userAddress)
}

// PendingLpRewardsInternal is a free data retrieval call binding the contract method 0x0198cf2b.
//
// Solidity: function pendingLpRewardsInternal(address _userAddress) view returns(uint256 pendingLp)
func (_AnimalFarm *AnimalFarmCallerSession) PendingLpRewardsInternal(_userAddress common.Address) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingLpRewardsInternal(&_AnimalFarm.CallOpts, _userAddress)
}

// PendingPigsRewards is a free data retrieval call binding the contract method 0x68c7d51d.
//
// Solidity: function pendingPigsRewards(address _user) view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) PendingPigsRewards(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "pendingPigsRewards", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingPigsRewards is a free data retrieval call binding the contract method 0x68c7d51d.
//
// Solidity: function pendingPigsRewards(address _user) view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) PendingPigsRewards(_user common.Address) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingPigsRewards(&_AnimalFarm.CallOpts, _user)
}

// PendingPigsRewards is a free data retrieval call binding the contract method 0x68c7d51d.
//
// Solidity: function pendingPigsRewards(address _user) view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) PendingPigsRewards(_user common.Address) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingPigsRewards(&_AnimalFarm.CallOpts, _user)
}

// PendingPigsRewardsHelper is a free data retrieval call binding the contract method 0xf846bcd6.
//
// Solidity: function pendingPigsRewardsHelper(address _user, uint256 startIndex) view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) PendingPigsRewardsHelper(opts *bind.CallOpts, _user common.Address, startIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "pendingPigsRewardsHelper", _user, startIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingPigsRewardsHelper is a free data retrieval call binding the contract method 0xf846bcd6.
//
// Solidity: function pendingPigsRewardsHelper(address _user, uint256 startIndex) view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) PendingPigsRewardsHelper(_user common.Address, startIndex *big.Int) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingPigsRewardsHelper(&_AnimalFarm.CallOpts, _user, startIndex)
}

// PendingPigsRewardsHelper is a free data retrieval call binding the contract method 0xf846bcd6.
//
// Solidity: function pendingPigsRewardsHelper(address _user, uint256 startIndex) view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) PendingPigsRewardsHelper(_user common.Address, startIndex *big.Int) (*big.Int, error) {
	return _AnimalFarm.Contract.PendingPigsRewardsHelper(&_AnimalFarm.CallOpts, _user, startIndex)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address _userAddress) view returns(uint256 _pendingPigs, uint256 _pendingLp)
func (_AnimalFarm *AnimalFarmCaller) PendingRewards(opts *bind.CallOpts, _userAddress common.Address) (struct {
	PendingPigs *big.Int
	PendingLp   *big.Int
}, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "pendingRewards", _userAddress)

	outstruct := new(struct {
		PendingPigs *big.Int
		PendingLp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingPigs = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingLp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address _userAddress) view returns(uint256 _pendingPigs, uint256 _pendingLp)
func (_AnimalFarm *AnimalFarmSession) PendingRewards(_userAddress common.Address) (struct {
	PendingPigs *big.Int
	PendingLp   *big.Int
}, error) {
	return _AnimalFarm.Contract.PendingRewards(&_AnimalFarm.CallOpts, _userAddress)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address _userAddress) view returns(uint256 _pendingPigs, uint256 _pendingLp)
func (_AnimalFarm *AnimalFarmCallerSession) PendingRewards(_userAddress common.Address) (struct {
	PendingPigs *big.Int
	PendingLp   *big.Int
}, error) {
	return _AnimalFarm.Contract.PendingRewards(&_AnimalFarm.CallOpts, _userAddress)
}

// TimeSinceLastCall is a free data retrieval call binding the contract method 0x3b39dfd4.
//
// Solidity: function timeSinceLastCall() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) TimeSinceLastCall(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "timeSinceLastCall")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeSinceLastCall is a free data retrieval call binding the contract method 0x3b39dfd4.
//
// Solidity: function timeSinceLastCall() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) TimeSinceLastCall() (*big.Int, error) {
	return _AnimalFarm.Contract.TimeSinceLastCall(&_AnimalFarm.CallOpts)
}

// TimeSinceLastCall is a free data retrieval call binding the contract method 0x3b39dfd4.
//
// Solidity: function timeSinceLastCall() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) TimeSinceLastCall() (*big.Int, error) {
	return _AnimalFarm.Contract.TimeSinceLastCall(&_AnimalFarm.CallOpts)
}

// TotalDogsStaked is a free data retrieval call binding the contract method 0x6eef45ab.
//
// Solidity: function totalDogsStaked() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) TotalDogsStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "totalDogsStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDogsStaked is a free data retrieval call binding the contract method 0x6eef45ab.
//
// Solidity: function totalDogsStaked() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) TotalDogsStaked() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalDogsStaked(&_AnimalFarm.CallOpts)
}

// TotalDogsStaked is a free data retrieval call binding the contract method 0x6eef45ab.
//
// Solidity: function totalDogsStaked() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) TotalDogsStaked() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalDogsStaked(&_AnimalFarm.CallOpts)
}

// TotalLPCollected is a free data retrieval call binding the contract method 0x94bf8f40.
//
// Solidity: function totalLPCollected() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) TotalLPCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "totalLPCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLPCollected is a free data retrieval call binding the contract method 0x94bf8f40.
//
// Solidity: function totalLPCollected() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) TotalLPCollected() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalLPCollected(&_AnimalFarm.CallOpts)
}

// TotalLPCollected is a free data retrieval call binding the contract method 0x94bf8f40.
//
// Solidity: function totalLPCollected() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) TotalLPCollected() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalLPCollected(&_AnimalFarm.CallOpts)
}

// TotalLPstakedTemp is a free data retrieval call binding the contract method 0x2c343581.
//
// Solidity: function totalLPstakedTemp() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) TotalLPstakedTemp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "totalLPstakedTemp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLPstakedTemp is a free data retrieval call binding the contract method 0x2c343581.
//
// Solidity: function totalLPstakedTemp() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) TotalLPstakedTemp() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalLPstakedTemp(&_AnimalFarm.CallOpts)
}

// TotalLPstakedTemp is a free data retrieval call binding the contract method 0x2c343581.
//
// Solidity: function totalLPstakedTemp() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) TotalLPstakedTemp() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalLPstakedTemp(&_AnimalFarm.CallOpts)
}

// TotalLpStaked is a free data retrieval call binding the contract method 0x983d90b4.
//
// Solidity: function totalLpStaked() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) TotalLpStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "totalLpStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLpStaked is a free data retrieval call binding the contract method 0x983d90b4.
//
// Solidity: function totalLpStaked() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) TotalLpStaked() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalLpStaked(&_AnimalFarm.CallOpts)
}

// TotalLpStaked is a free data retrieval call binding the contract method 0x983d90b4.
//
// Solidity: function totalLpStaked() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) TotalLpStaked() (*big.Int, error) {
	return _AnimalFarm.Contract.TotalLpStaked(&_AnimalFarm.CallOpts)
}

// UpdateInterval is a free data retrieval call binding the contract method 0xfd2c80ae.
//
// Solidity: function updateInterval() view returns(uint256)
func (_AnimalFarm *AnimalFarmCaller) UpdateInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "updateInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateInterval is a free data retrieval call binding the contract method 0xfd2c80ae.
//
// Solidity: function updateInterval() view returns(uint256)
func (_AnimalFarm *AnimalFarmSession) UpdateInterval() (*big.Int, error) {
	return _AnimalFarm.Contract.UpdateInterval(&_AnimalFarm.CallOpts)
}

// UpdateInterval is a free data retrieval call binding the contract method 0xfd2c80ae.
//
// Solidity: function updateInterval() view returns(uint256)
func (_AnimalFarm *AnimalFarmCallerSession) UpdateInterval() (*big.Int, error) {
	return _AnimalFarm.Contract.UpdateInterval(&_AnimalFarm.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 lpMask, uint256 pigsClaimedTotal, uint256 lastRmsClaimed, uint256 lpDebt, uint256 totalLPCollected, uint256 totalPigsCollected)
func (_AnimalFarm *AnimalFarmCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount             *big.Int
	LpMask             *big.Int
	PigsClaimedTotal   *big.Int
	LastRmsClaimed     *big.Int
	LpDebt             *big.Int
	TotalLPCollected   *big.Int
	TotalPigsCollected *big.Int
}, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount             *big.Int
		LpMask             *big.Int
		PigsClaimedTotal   *big.Int
		LastRmsClaimed     *big.Int
		LpDebt             *big.Int
		TotalLPCollected   *big.Int
		TotalPigsCollected *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LpMask = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PigsClaimedTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastRmsClaimed = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LpDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalLPCollected = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalPigsCollected = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 lpMask, uint256 pigsClaimedTotal, uint256 lastRmsClaimed, uint256 lpDebt, uint256 totalLPCollected, uint256 totalPigsCollected)
func (_AnimalFarm *AnimalFarmSession) UserInfo(arg0 common.Address) (struct {
	Amount             *big.Int
	LpMask             *big.Int
	PigsClaimedTotal   *big.Int
	LastRmsClaimed     *big.Int
	LpDebt             *big.Int
	TotalLPCollected   *big.Int
	TotalPigsCollected *big.Int
}, error) {
	return _AnimalFarm.Contract.UserInfo(&_AnimalFarm.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 lpMask, uint256 pigsClaimedTotal, uint256 lastRmsClaimed, uint256 lpDebt, uint256 totalLPCollected, uint256 totalPigsCollected)
func (_AnimalFarm *AnimalFarmCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount             *big.Int
	LpMask             *big.Int
	PigsClaimedTotal   *big.Int
	LastRmsClaimed     *big.Int
	LpDebt             *big.Int
	TotalLPCollected   *big.Int
	TotalPigsCollected *big.Int
}, error) {
	return _AnimalFarm.Contract.UserInfo(&_AnimalFarm.CallOpts, arg0)
}

// WbnbCurrencyAddress is a free data retrieval call binding the contract method 0x9b231eae.
//
// Solidity: function wbnbCurrencyAddress() view returns(address)
func (_AnimalFarm *AnimalFarmCaller) WbnbCurrencyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnimalFarm.contract.Call(opts, &out, "wbnbCurrencyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WbnbCurrencyAddress is a free data retrieval call binding the contract method 0x9b231eae.
//
// Solidity: function wbnbCurrencyAddress() view returns(address)
func (_AnimalFarm *AnimalFarmSession) WbnbCurrencyAddress() (common.Address, error) {
	return _AnimalFarm.Contract.WbnbCurrencyAddress(&_AnimalFarm.CallOpts)
}

// WbnbCurrencyAddress is a free data retrieval call binding the contract method 0x9b231eae.
//
// Solidity: function wbnbCurrencyAddress() view returns(address)
func (_AnimalFarm *AnimalFarmCallerSession) WbnbCurrencyAddress() (common.Address, error) {
	return _AnimalFarm.Contract.WbnbCurrencyAddress(&_AnimalFarm.CallOpts)
}

// MClockedAddress is a paid mutator transaction binding the contract method 0x45c9cc96.
//
// Solidity: function MClockedAddress() returns()
func (_AnimalFarm *AnimalFarmTransactor) MClockedAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "MClockedAddress")
}

// MClockedAddress is a paid mutator transaction binding the contract method 0x45c9cc96.
//
// Solidity: function MClockedAddress() returns()
func (_AnimalFarm *AnimalFarmSession) MClockedAddress() (*types.Transaction, error) {
	return _AnimalFarm.Contract.MClockedAddress(&_AnimalFarm.TransactOpts)
}

// MClockedAddress is a paid mutator transaction binding the contract method 0x45c9cc96.
//
// Solidity: function MClockedAddress() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) MClockedAddress() (*types.Transaction, error) {
	return _AnimalFarm.Contract.MClockedAddress(&_AnimalFarm.TransactOpts)
}

// AddInitAllowed is a paid mutator transaction binding the contract method 0x4cc308b2.
//
// Solidity: function addInitAllowed(address _ad, bool _bool) returns()
func (_AnimalFarm *AnimalFarmTransactor) AddInitAllowed(opts *bind.TransactOpts, _ad common.Address, _bool bool) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "addInitAllowed", _ad, _bool)
}

// AddInitAllowed is a paid mutator transaction binding the contract method 0x4cc308b2.
//
// Solidity: function addInitAllowed(address _ad, bool _bool) returns()
func (_AnimalFarm *AnimalFarmSession) AddInitAllowed(_ad common.Address, _bool bool) (*types.Transaction, error) {
	return _AnimalFarm.Contract.AddInitAllowed(&_AnimalFarm.TransactOpts, _ad, _bool)
}

// AddInitAllowed is a paid mutator transaction binding the contract method 0x4cc308b2.
//
// Solidity: function addInitAllowed(address _ad, bool _bool) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) AddInitAllowed(_ad common.Address, _bool bool) (*types.Transaction, error) {
	return _AnimalFarm.Contract.AddInitAllowed(&_AnimalFarm.TransactOpts, _ad, _bool)
}

// AllowCompound is a paid mutator transaction binding the contract method 0xab47eaa4.
//
// Solidity: function allowCompound(uint256 _time) returns()
func (_AnimalFarm *AnimalFarmTransactor) AllowCompound(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "allowCompound", _time)
}

// AllowCompound is a paid mutator transaction binding the contract method 0xab47eaa4.
//
// Solidity: function allowCompound(uint256 _time) returns()
func (_AnimalFarm *AnimalFarmSession) AllowCompound(_time *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.AllowCompound(&_AnimalFarm.TransactOpts, _time)
}

// AllowCompound is a paid mutator transaction binding the contract method 0xab47eaa4.
//
// Solidity: function allowCompound(uint256 _time) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) AllowCompound(_time *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.AllowCompound(&_AnimalFarm.TransactOpts, _time)
}

// ChangeUpdateInterval is a paid mutator transaction binding the contract method 0x12cf6366.
//
// Solidity: function changeUpdateInterval(uint256 _time) returns()
func (_AnimalFarm *AnimalFarmTransactor) ChangeUpdateInterval(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "changeUpdateInterval", _time)
}

// ChangeUpdateInterval is a paid mutator transaction binding the contract method 0x12cf6366.
//
// Solidity: function changeUpdateInterval(uint256 _time) returns()
func (_AnimalFarm *AnimalFarmSession) ChangeUpdateInterval(_time *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.ChangeUpdateInterval(&_AnimalFarm.TransactOpts, _time)
}

// ChangeUpdateInterval is a paid mutator transaction binding the contract method 0x12cf6366.
//
// Solidity: function changeUpdateInterval(uint256 _time) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) ChangeUpdateInterval(_time *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.ChangeUpdateInterval(&_AnimalFarm.TransactOpts, _time)
}

// ClaimLpTokensAndPigs is a paid mutator transaction binding the contract method 0xd5d0ed45.
//
// Solidity: function claimLpTokensAndPigs() returns()
func (_AnimalFarm *AnimalFarmTransactor) ClaimLpTokensAndPigs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "claimLpTokensAndPigs")
}

// ClaimLpTokensAndPigs is a paid mutator transaction binding the contract method 0xd5d0ed45.
//
// Solidity: function claimLpTokensAndPigs() returns()
func (_AnimalFarm *AnimalFarmSession) ClaimLpTokensAndPigs() (*types.Transaction, error) {
	return _AnimalFarm.Contract.ClaimLpTokensAndPigs(&_AnimalFarm.TransactOpts)
}

// ClaimLpTokensAndPigs is a paid mutator transaction binding the contract method 0xd5d0ed45.
//
// Solidity: function claimLpTokensAndPigs() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) ClaimLpTokensAndPigs() (*types.Transaction, error) {
	return _AnimalFarm.Contract.ClaimLpTokensAndPigs(&_AnimalFarm.TransactOpts)
}

// ClaimPigs is a paid mutator transaction binding the contract method 0x7284d47d.
//
// Solidity: function claimPigs() returns()
func (_AnimalFarm *AnimalFarmTransactor) ClaimPigs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "claimPigs")
}

// ClaimPigs is a paid mutator transaction binding the contract method 0x7284d47d.
//
// Solidity: function claimPigs() returns()
func (_AnimalFarm *AnimalFarmSession) ClaimPigs() (*types.Transaction, error) {
	return _AnimalFarm.Contract.ClaimPigs(&_AnimalFarm.TransactOpts)
}

// ClaimPigs is a paid mutator transaction binding the contract method 0x7284d47d.
//
// Solidity: function claimPigs() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) ClaimPigs() (*types.Transaction, error) {
	return _AnimalFarm.Contract.ClaimPigs(&_AnimalFarm.TransactOpts)
}

// ClaimPigsHelper is a paid mutator transaction binding the contract method 0x196faea3.
//
// Solidity: function claimPigsHelper(uint256 startIndex) returns()
func (_AnimalFarm *AnimalFarmTransactor) ClaimPigsHelper(opts *bind.TransactOpts, startIndex *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "claimPigsHelper", startIndex)
}

// ClaimPigsHelper is a paid mutator transaction binding the contract method 0x196faea3.
//
// Solidity: function claimPigsHelper(uint256 startIndex) returns()
func (_AnimalFarm *AnimalFarmSession) ClaimPigsHelper(startIndex *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.ClaimPigsHelper(&_AnimalFarm.TransactOpts, startIndex)
}

// ClaimPigsHelper is a paid mutator transaction binding the contract method 0x196faea3.
//
// Solidity: function claimPigsHelper(uint256 startIndex) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) ClaimPigsHelper(startIndex *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.ClaimPigsHelper(&_AnimalFarm.TransactOpts, startIndex)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_AnimalFarm *AnimalFarmTransactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_AnimalFarm *AnimalFarmSession) Compound() (*types.Transaction, error) {
	return _AnimalFarm.Contract.Compound(&_AnimalFarm.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) Compound() (*types.Transaction, error) {
	return _AnimalFarm.Contract.Compound(&_AnimalFarm.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _user, uint256 _amount) returns()
func (_AnimalFarm *AnimalFarmTransactor) Deposit(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "deposit", _user, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _user, uint256 _amount) returns()
func (_AnimalFarm *AnimalFarmSession) Deposit(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.Deposit(&_AnimalFarm.TransactOpts, _user, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _user, uint256 _amount) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) Deposit(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.Deposit(&_AnimalFarm.TransactOpts, _user, _amount)
}

// IncreasePigsBuffer is a paid mutator transaction binding the contract method 0xf8acae4b.
//
// Solidity: function increasePigsBuffer(uint256 quant) returns()
func (_AnimalFarm *AnimalFarmTransactor) IncreasePigsBuffer(opts *bind.TransactOpts, quant *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "increasePigsBuffer", quant)
}

// IncreasePigsBuffer is a paid mutator transaction binding the contract method 0xf8acae4b.
//
// Solidity: function increasePigsBuffer(uint256 quant) returns()
func (_AnimalFarm *AnimalFarmSession) IncreasePigsBuffer(quant *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.IncreasePigsBuffer(&_AnimalFarm.TransactOpts, quant)
}

// IncreasePigsBuffer is a paid mutator transaction binding the contract method 0xf8acae4b.
//
// Solidity: function increasePigsBuffer(uint256 quant) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) IncreasePigsBuffer(quant *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.IncreasePigsBuffer(&_AnimalFarm.TransactOpts, quant)
}

// InitCompounders is a paid mutator transaction binding the contract method 0x4827d640.
//
// Solidity: function initCompounders(address[] _users) returns()
func (_AnimalFarm *AnimalFarmTransactor) InitCompounders(opts *bind.TransactOpts, _users []common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "initCompounders", _users)
}

// InitCompounders is a paid mutator transaction binding the contract method 0x4827d640.
//
// Solidity: function initCompounders(address[] _users) returns()
func (_AnimalFarm *AnimalFarmSession) InitCompounders(_users []common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitCompounders(&_AnimalFarm.TransactOpts, _users)
}

// InitCompounders is a paid mutator transaction binding the contract method 0x4827d640.
//
// Solidity: function initCompounders(address[] _users) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) InitCompounders(_users []common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitCompounders(&_AnimalFarm.TransactOpts, _users)
}

// InitMCStake is a paid mutator transaction binding the contract method 0x3ab14b6d.
//
// Solidity: function initMCStake() returns()
func (_AnimalFarm *AnimalFarmTransactor) InitMCStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "initMCStake")
}

// InitMCStake is a paid mutator transaction binding the contract method 0x3ab14b6d.
//
// Solidity: function initMCStake() returns()
func (_AnimalFarm *AnimalFarmSession) InitMCStake() (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitMCStake(&_AnimalFarm.TransactOpts)
}

// InitMCStake is a paid mutator transaction binding the contract method 0x3ab14b6d.
//
// Solidity: function initMCStake() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) InitMCStake() (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitMCStake(&_AnimalFarm.TransactOpts)
}

// InitStakeMult is a paid mutator transaction binding the contract method 0x66994c16.
//
// Solidity: function initStakeMult(uint256 temp1, uint256 temp2) returns()
func (_AnimalFarm *AnimalFarmTransactor) InitStakeMult(opts *bind.TransactOpts, temp1 *big.Int, temp2 *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "initStakeMult", temp1, temp2)
}

// InitStakeMult is a paid mutator transaction binding the contract method 0x66994c16.
//
// Solidity: function initStakeMult(uint256 temp1, uint256 temp2) returns()
func (_AnimalFarm *AnimalFarmSession) InitStakeMult(temp1 *big.Int, temp2 *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitStakeMult(&_AnimalFarm.TransactOpts, temp1, temp2)
}

// InitStakeMult is a paid mutator transaction binding the contract method 0x66994c16.
//
// Solidity: function initStakeMult(uint256 temp1, uint256 temp2) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) InitStakeMult(temp1 *big.Int, temp2 *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitStakeMult(&_AnimalFarm.TransactOpts, temp1, temp2)
}

// InitializeMd is a paid mutator transaction binding the contract method 0x9d75ff33.
//
// Solidity: function initializeMd(address[] _users, (uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _info) returns()
func (_AnimalFarm *AnimalFarmTransactor) InitializeMd(opts *bind.TransactOpts, _users []common.Address, _info []DogPoundAutoPoolUserInfo) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "initializeMd", _users, _info)
}

// InitializeMd is a paid mutator transaction binding the contract method 0x9d75ff33.
//
// Solidity: function initializeMd(address[] _users, (uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _info) returns()
func (_AnimalFarm *AnimalFarmSession) InitializeMd(_users []common.Address, _info []DogPoundAutoPoolUserInfo) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitializeMd(&_AnimalFarm.TransactOpts, _users, _info)
}

// InitializeMd is a paid mutator transaction binding the contract method 0x9d75ff33.
//
// Solidity: function initializeMd(address[] _users, (uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _info) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) InitializeMd(_users []common.Address, _info []DogPoundAutoPoolUserInfo) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitializeMd(&_AnimalFarm.TransactOpts, _users, _info)
}

// InitializeU is a paid mutator transaction binding the contract method 0xaa57ce4e.
//
// Solidity: function initializeU(address _pool, address[] _users) returns()
func (_AnimalFarm *AnimalFarmTransactor) InitializeU(opts *bind.TransactOpts, _pool common.Address, _users []common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "initializeU", _pool, _users)
}

// InitializeU is a paid mutator transaction binding the contract method 0xaa57ce4e.
//
// Solidity: function initializeU(address _pool, address[] _users) returns()
func (_AnimalFarm *AnimalFarmSession) InitializeU(_pool common.Address, _users []common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitializeU(&_AnimalFarm.TransactOpts, _pool, _users)
}

// InitializeU is a paid mutator transaction binding the contract method 0xaa57ce4e.
//
// Solidity: function initializeU(address _pool, address[] _users) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) InitializeU(_pool common.Address, _users []common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitializeU(&_AnimalFarm.TransactOpts, _pool, _users)
}

// InitializeVariables is a paid mutator transaction binding the contract method 0x0992d596.
//
// Solidity: function initializeVariables(address _pool, uint256 histlen) returns()
func (_AnimalFarm *AnimalFarmTransactor) InitializeVariables(opts *bind.TransactOpts, _pool common.Address, histlen *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "initializeVariables", _pool, histlen)
}

// InitializeVariables is a paid mutator transaction binding the contract method 0x0992d596.
//
// Solidity: function initializeVariables(address _pool, uint256 histlen) returns()
func (_AnimalFarm *AnimalFarmSession) InitializeVariables(_pool common.Address, histlen *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitializeVariables(&_AnimalFarm.TransactOpts, _pool, histlen)
}

// InitializeVariables is a paid mutator transaction binding the contract method 0x0992d596.
//
// Solidity: function initializeVariables(address _pool, uint256 histlen) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) InitializeVariables(_pool common.Address, histlen *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.InitializeVariables(&_AnimalFarm.TransactOpts, _pool, histlen)
}

// LockDogPoundManager is a paid mutator transaction binding the contract method 0x7a838833.
//
// Solidity: function lockDogPoundManager() returns()
func (_AnimalFarm *AnimalFarmTransactor) LockDogPoundManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "lockDogPoundManager")
}

// LockDogPoundManager is a paid mutator transaction binding the contract method 0x7a838833.
//
// Solidity: function lockDogPoundManager() returns()
func (_AnimalFarm *AnimalFarmSession) LockDogPoundManager() (*types.Transaction, error) {
	return _AnimalFarm.Contract.LockDogPoundManager(&_AnimalFarm.TransactOpts)
}

// LockDogPoundManager is a paid mutator transaction binding the contract method 0x7a838833.
//
// Solidity: function lockDogPoundManager() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) LockDogPoundManager() (*types.Transaction, error) {
	return _AnimalFarm.Contract.LockDogPoundManager(&_AnimalFarm.TransactOpts)
}

// PauseInitialize is a paid mutator transaction binding the contract method 0x35e867b1.
//
// Solidity: function pauseInitialize() returns()
func (_AnimalFarm *AnimalFarmTransactor) PauseInitialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "pauseInitialize")
}

// PauseInitialize is a paid mutator transaction binding the contract method 0x35e867b1.
//
// Solidity: function pauseInitialize() returns()
func (_AnimalFarm *AnimalFarmSession) PauseInitialize() (*types.Transaction, error) {
	return _AnimalFarm.Contract.PauseInitialize(&_AnimalFarm.TransactOpts)
}

// PauseInitialize is a paid mutator transaction binding the contract method 0x35e867b1.
//
// Solidity: function pauseInitialize() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) PauseInitialize() (*types.Transaction, error) {
	return _AnimalFarm.Contract.PauseInitialize(&_AnimalFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnimalFarm *AnimalFarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnimalFarm *AnimalFarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _AnimalFarm.Contract.RenounceOwnership(&_AnimalFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnimalFarm *AnimalFarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AnimalFarm.Contract.RenounceOwnership(&_AnimalFarm.TransactOpts)
}

// SetDogPoundManager is a paid mutator transaction binding the contract method 0x09101772.
//
// Solidity: function setDogPoundManager(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactor) SetDogPoundManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "setDogPoundManager", _address)
}

// SetDogPoundManager is a paid mutator transaction binding the contract method 0x09101772.
//
// Solidity: function setDogPoundManager(address _address) returns()
func (_AnimalFarm *AnimalFarmSession) SetDogPoundManager(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.SetDogPoundManager(&_AnimalFarm.TransactOpts, _address)
}

// SetDogPoundManager is a paid mutator transaction binding the contract method 0x09101772.
//
// Solidity: function setDogPoundManager(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) SetDogPoundManager(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.SetDogPoundManager(&_AnimalFarm.TransactOpts, _address)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnimalFarm *AnimalFarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnimalFarm *AnimalFarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.TransferOwnership(&_AnimalFarm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.TransferOwnership(&_AnimalFarm.TransactOpts, newOwner)
}

// UpdateBnbLiqThreshhold is a paid mutator transaction binding the contract method 0x4e38cc6a.
//
// Solidity: function updateBnbLiqThreshhold(uint256 newThrehshold) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdateBnbLiqThreshhold(opts *bind.TransactOpts, newThrehshold *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updateBnbLiqThreshhold", newThrehshold)
}

// UpdateBnbLiqThreshhold is a paid mutator transaction binding the contract method 0x4e38cc6a.
//
// Solidity: function updateBnbLiqThreshhold(uint256 newThrehshold) returns()
func (_AnimalFarm *AnimalFarmSession) UpdateBnbLiqThreshhold(newThrehshold *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateBnbLiqThreshhold(&_AnimalFarm.TransactOpts, newThrehshold)
}

// UpdateBnbLiqThreshhold is a paid mutator transaction binding the contract method 0x4e38cc6a.
//
// Solidity: function updateBnbLiqThreshhold(uint256 newThrehshold) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdateBnbLiqThreshhold(newThrehshold *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateBnbLiqThreshhold(&_AnimalFarm.TransactOpts, newThrehshold)
}

// UpdateDogsAddress is a paid mutator transaction binding the contract method 0x01cc969c.
//
// Solidity: function updateDogsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdateDogsAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updateDogsAddress", _address)
}

// UpdateDogsAddress is a paid mutator transaction binding the contract method 0x01cc969c.
//
// Solidity: function updateDogsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmSession) UpdateDogsAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdateDogsAddress is a paid mutator transaction binding the contract method 0x01cc969c.
//
// Solidity: function updateDogsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdateDogsAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdateDogsAndLPAddress is a paid mutator transaction binding the contract method 0x9a9598bc.
//
// Solidity: function updateDogsAndLPAddress(address _addressDogs, address _addressLpBNB) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdateDogsAndLPAddress(opts *bind.TransactOpts, _addressDogs common.Address, _addressLpBNB common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updateDogsAndLPAddress", _addressDogs, _addressLpBNB)
}

// UpdateDogsAndLPAddress is a paid mutator transaction binding the contract method 0x9a9598bc.
//
// Solidity: function updateDogsAndLPAddress(address _addressDogs, address _addressLpBNB) returns()
func (_AnimalFarm *AnimalFarmSession) UpdateDogsAndLPAddress(_addressDogs common.Address, _addressLpBNB common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsAndLPAddress(&_AnimalFarm.TransactOpts, _addressDogs, _addressLpBNB)
}

// UpdateDogsAndLPAddress is a paid mutator transaction binding the contract method 0x9a9598bc.
//
// Solidity: function updateDogsAndLPAddress(address _addressDogs, address _addressLpBNB) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdateDogsAndLPAddress(_addressDogs common.Address, _addressLpBNB common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsAndLPAddress(&_AnimalFarm.TransactOpts, _addressDogs, _addressLpBNB)
}

// UpdateDogsBnBPID is a paid mutator transaction binding the contract method 0x38472ccf.
//
// Solidity: function updateDogsBnBPID(uint256 newPid) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdateDogsBnBPID(opts *bind.TransactOpts, newPid *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updateDogsBnBPID", newPid)
}

// UpdateDogsBnBPID is a paid mutator transaction binding the contract method 0x38472ccf.
//
// Solidity: function updateDogsBnBPID(uint256 newPid) returns()
func (_AnimalFarm *AnimalFarmSession) UpdateDogsBnBPID(newPid *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsBnBPID(&_AnimalFarm.TransactOpts, newPid)
}

// UpdateDogsBnBPID is a paid mutator transaction binding the contract method 0x38472ccf.
//
// Solidity: function updateDogsBnBPID(uint256 newPid) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdateDogsBnBPID(newPid *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsBnBPID(&_AnimalFarm.TransactOpts, newPid)
}

// UpdateDogsExchanceHelperAddress is a paid mutator transaction binding the contract method 0xd74c8515.
//
// Solidity: function updateDogsExchanceHelperAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdateDogsExchanceHelperAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updateDogsExchanceHelperAddress", _address)
}

// UpdateDogsExchanceHelperAddress is a paid mutator transaction binding the contract method 0xd74c8515.
//
// Solidity: function updateDogsExchanceHelperAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmSession) UpdateDogsExchanceHelperAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsExchanceHelperAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdateDogsExchanceHelperAddress is a paid mutator transaction binding the contract method 0xd74c8515.
//
// Solidity: function updateDogsExchanceHelperAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdateDogsExchanceHelperAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateDogsExchanceHelperAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdateMasterchefPigsAddress is a paid mutator transaction binding the contract method 0x410ee17a.
//
// Solidity: function updateMasterchefPigsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdateMasterchefPigsAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updateMasterchefPigsAddress", _address)
}

// UpdateMasterchefPigsAddress is a paid mutator transaction binding the contract method 0x410ee17a.
//
// Solidity: function updateMasterchefPigsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmSession) UpdateMasterchefPigsAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateMasterchefPigsAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdateMasterchefPigsAddress is a paid mutator transaction binding the contract method 0x410ee17a.
//
// Solidity: function updateMasterchefPigsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdateMasterchefPigsAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdateMasterchefPigsAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdatePigsAddress is a paid mutator transaction binding the contract method 0x395c4438.
//
// Solidity: function updatePigsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactor) UpdatePigsAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "updatePigsAddress", _address)
}

// UpdatePigsAddress is a paid mutator transaction binding the contract method 0x395c4438.
//
// Solidity: function updatePigsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmSession) UpdatePigsAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdatePigsAddress(&_AnimalFarm.TransactOpts, _address)
}

// UpdatePigsAddress is a paid mutator transaction binding the contract method 0x395c4438.
//
// Solidity: function updatePigsAddress(address _address) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) UpdatePigsAddress(_address common.Address) (*types.Transaction, error) {
	return _AnimalFarm.Contract.UpdatePigsAddress(&_AnimalFarm.TransactOpts, _address)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _user, uint256 _amount) returns()
func (_AnimalFarm *AnimalFarmTransactor) Withdraw(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.contract.Transact(opts, "withdraw", _user, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _user, uint256 _amount) returns()
func (_AnimalFarm *AnimalFarmSession) Withdraw(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.Withdraw(&_AnimalFarm.TransactOpts, _user, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _user, uint256 _amount) returns()
func (_AnimalFarm *AnimalFarmTransactorSession) Withdraw(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AnimalFarm.Contract.Withdraw(&_AnimalFarm.TransactOpts, _user, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnimalFarm *AnimalFarmTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnimalFarm.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnimalFarm *AnimalFarmSession) Receive() (*types.Transaction, error) {
	return _AnimalFarm.Contract.Receive(&_AnimalFarm.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnimalFarm *AnimalFarmTransactorSession) Receive() (*types.Transaction, error) {
	return _AnimalFarm.Contract.Receive(&_AnimalFarm.TransactOpts)
}

// AnimalFarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AnimalFarm contract.
type AnimalFarmOwnershipTransferredIterator struct {
	Event *AnimalFarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AnimalFarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnimalFarmOwnershipTransferred)
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
		it.Event = new(AnimalFarmOwnershipTransferred)
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
func (it *AnimalFarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnimalFarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnimalFarmOwnershipTransferred represents a OwnershipTransferred event raised by the AnimalFarm contract.
type AnimalFarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnimalFarm *AnimalFarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AnimalFarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AnimalFarm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AnimalFarmOwnershipTransferredIterator{contract: _AnimalFarm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnimalFarm *AnimalFarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AnimalFarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AnimalFarm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnimalFarmOwnershipTransferred)
				if err := _AnimalFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnimalFarm *AnimalFarmFilterer) ParseOwnershipTransferred(log types.Log) (*AnimalFarmOwnershipTransferred, error) {
	event := new(AnimalFarmOwnershipTransferred)
	if err := _AnimalFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
