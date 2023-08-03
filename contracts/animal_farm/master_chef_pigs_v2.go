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

// MasterChefPigsV2MetaData contains all meta data concerning the MasterChefPigsV2 contract.
var MasterChefPigsV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platform\",\"type\":\"address\"},{\"internalType\":\"contractIFounderStakerV2\",\"name\":\"_founder\",\"type\":\"address\"},{\"internalType\":\"contractIDDSCA\",\"name\":\"_ddsca\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestInterval\",\"type\":\"uint256\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dogPoundAutoPool\",\"type\":\"address\"}],\"name\":\"DogPoundAutoPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"govAddress\",\"type\":\"address\"}],\"name\":\"GovUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLockedUp\",\"type\":\"uint256\"}],\"name\":\"RewardLockedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"founder\",\"type\":\"address\"}],\"name\":\"SetFounder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerReward\",\"type\":\"uint256\"}],\"name\":\"SetOwnersRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetPlatformAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestInterval\",\"type\":\"uint256\"}],\"name\":\"SetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DogPoundAutoPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOUNDER\",\"outputs\":[{\"internalType\":\"contractIFounderStakerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_HARVEST_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PLATFORM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PancakeRouter\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PigsV2Token\",\"outputs\":[{\"internalType\":\"contractIPigsToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestInterval\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"busdCurrencyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"canHarvest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dripBusdPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dripTaxVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getPigsMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increasePigsSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerPigsReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingPigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accPigsPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositFeeBP\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestInterval\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDDSCA\",\"name\":\"_ddsca\",\"type\":\"address\"}],\"name\":\"setDDSCAAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dogPoundAutoPool\",\"type\":\"address\"}],\"name\":\"setDogPoundAutoPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartBlock\",\"type\":\"uint256\"}],\"name\":\"setFarmStartBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFounderStakerV2\",\"name\":\"_founder\",\"type\":\"address\"}],\"name\":\"setFoundersAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardsAmount\",\"type\":\"uint256\"}],\"name\":\"setFoundersRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_govAddress\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformAddress\",\"type\":\"address\"}],\"name\":\"setPlatformAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedUpRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInCents\",\"type\":\"uint256\"}],\"name\":\"updateEmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"updateMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pigsRewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardLockedUp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextHarvestUntil\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterChefPigsV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterChefPigsV2MetaData.ABI instead.
var MasterChefPigsV2ABI = MasterChefPigsV2MetaData.ABI

// MasterChefPigsV2 is an auto generated Go binding around an Ethereum contract.
type MasterChefPigsV2 struct {
	MasterChefPigsV2Caller     // Read-only binding to the contract
	MasterChefPigsV2Transactor // Write-only binding to the contract
	MasterChefPigsV2Filterer   // Log filterer for contract events
}

// MasterChefPigsV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasterChefPigsV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefPigsV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterChefPigsV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefPigsV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterChefPigsV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefPigsV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterChefPigsV2Session struct {
	Contract     *MasterChefPigsV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterChefPigsV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterChefPigsV2CallerSession struct {
	Contract *MasterChefPigsV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MasterChefPigsV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterChefPigsV2TransactorSession struct {
	Contract     *MasterChefPigsV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MasterChefPigsV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasterChefPigsV2Raw struct {
	Contract *MasterChefPigsV2 // Generic contract binding to access the raw methods on
}

// MasterChefPigsV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterChefPigsV2CallerRaw struct {
	Contract *MasterChefPigsV2Caller // Generic read-only contract binding to access the raw methods on
}

// MasterChefPigsV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterChefPigsV2TransactorRaw struct {
	Contract *MasterChefPigsV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterChefPigsV2 creates a new instance of MasterChefPigsV2, bound to a specific deployed contract.
func NewMasterChefPigsV2(address common.Address, backend bind.ContractBackend) (*MasterChefPigsV2, error) {
	contract, err := bindMasterChefPigsV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2{MasterChefPigsV2Caller: MasterChefPigsV2Caller{contract: contract}, MasterChefPigsV2Transactor: MasterChefPigsV2Transactor{contract: contract}, MasterChefPigsV2Filterer: MasterChefPigsV2Filterer{contract: contract}}, nil
}

// NewMasterChefPigsV2Caller creates a new read-only instance of MasterChefPigsV2, bound to a specific deployed contract.
func NewMasterChefPigsV2Caller(address common.Address, caller bind.ContractCaller) (*MasterChefPigsV2Caller, error) {
	contract, err := bindMasterChefPigsV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2Caller{contract: contract}, nil
}

// NewMasterChefPigsV2Transactor creates a new write-only instance of MasterChefPigsV2, bound to a specific deployed contract.
func NewMasterChefPigsV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MasterChefPigsV2Transactor, error) {
	contract, err := bindMasterChefPigsV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2Transactor{contract: contract}, nil
}

// NewMasterChefPigsV2Filterer creates a new log filterer instance of MasterChefPigsV2, bound to a specific deployed contract.
func NewMasterChefPigsV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MasterChefPigsV2Filterer, error) {
	contract, err := bindMasterChefPigsV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2Filterer{contract: contract}, nil
}

// bindMasterChefPigsV2 binds a generic wrapper to an already deployed contract.
func bindMasterChefPigsV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterChefPigsV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChefPigsV2 *MasterChefPigsV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChefPigsV2.Contract.MasterChefPigsV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChefPigsV2 *MasterChefPigsV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.MasterChefPigsV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChefPigsV2 *MasterChefPigsV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.MasterChefPigsV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChefPigsV2 *MasterChefPigsV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChefPigsV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.contract.Transact(opts, method, params...)
}

// DogPoundAutoPool is a free data retrieval call binding the contract method 0xf974dada.
//
// Solidity: function DogPoundAutoPool() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) DogPoundAutoPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "DogPoundAutoPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogPoundAutoPool is a free data retrieval call binding the contract method 0xf974dada.
//
// Solidity: function DogPoundAutoPool() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) DogPoundAutoPool() (common.Address, error) {
	return _MasterChefPigsV2.Contract.DogPoundAutoPool(&_MasterChefPigsV2.CallOpts)
}

// DogPoundAutoPool is a free data retrieval call binding the contract method 0xf974dada.
//
// Solidity: function DogPoundAutoPool() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) DogPoundAutoPool() (common.Address, error) {
	return _MasterChefPigsV2.Contract.DogPoundAutoPool(&_MasterChefPigsV2.CallOpts)
}

// FOUNDER is a free data retrieval call binding the contract method 0xca58f2d3.
//
// Solidity: function FOUNDER() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) FOUNDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "FOUNDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FOUNDER is a free data retrieval call binding the contract method 0xca58f2d3.
//
// Solidity: function FOUNDER() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) FOUNDER() (common.Address, error) {
	return _MasterChefPigsV2.Contract.FOUNDER(&_MasterChefPigsV2.CallOpts)
}

// FOUNDER is a free data retrieval call binding the contract method 0xca58f2d3.
//
// Solidity: function FOUNDER() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) FOUNDER() (common.Address, error) {
	return _MasterChefPigsV2.Contract.FOUNDER(&_MasterChefPigsV2.CallOpts)
}

// MAXIMUMHARVESTINTERVAL is a free data retrieval call binding the contract method 0xde73149d.
//
// Solidity: function MAXIMUM_HARVEST_INTERVAL() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) MAXIMUMHARVESTINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "MAXIMUM_HARVEST_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMHARVESTINTERVAL is a free data retrieval call binding the contract method 0xde73149d.
//
// Solidity: function MAXIMUM_HARVEST_INTERVAL() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) MAXIMUMHARVESTINTERVAL() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.MAXIMUMHARVESTINTERVAL(&_MasterChefPigsV2.CallOpts)
}

// MAXIMUMHARVESTINTERVAL is a free data retrieval call binding the contract method 0xde73149d.
//
// Solidity: function MAXIMUM_HARVEST_INTERVAL() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) MAXIMUMHARVESTINTERVAL() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.MAXIMUMHARVESTINTERVAL(&_MasterChefPigsV2.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0xc2894b08.
//
// Solidity: function Migrator() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "Migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0xc2894b08.
//
// Solidity: function Migrator() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) Migrator() (common.Address, error) {
	return _MasterChefPigsV2.Contract.Migrator(&_MasterChefPigsV2.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0xc2894b08.
//
// Solidity: function Migrator() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) Migrator() (common.Address, error) {
	return _MasterChefPigsV2.Contract.Migrator(&_MasterChefPigsV2.CallOpts)
}

// PLATFORMADDRESS is a free data retrieval call binding the contract method 0xbbc28489.
//
// Solidity: function PLATFORM_ADDRESS() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PLATFORMADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "PLATFORM_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PLATFORMADDRESS is a free data retrieval call binding the contract method 0xbbc28489.
//
// Solidity: function PLATFORM_ADDRESS() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PLATFORMADDRESS() (common.Address, error) {
	return _MasterChefPigsV2.Contract.PLATFORMADDRESS(&_MasterChefPigsV2.CallOpts)
}

// PLATFORMADDRESS is a free data retrieval call binding the contract method 0xbbc28489.
//
// Solidity: function PLATFORM_ADDRESS() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PLATFORMADDRESS() (common.Address, error) {
	return _MasterChefPigsV2.Contract.PLATFORMADDRESS(&_MasterChefPigsV2.CallOpts)
}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PancakeRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "PancakeRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PancakeRouter() (common.Address, error) {
	return _MasterChefPigsV2.Contract.PancakeRouter(&_MasterChefPigsV2.CallOpts)
}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PancakeRouter() (common.Address, error) {
	return _MasterChefPigsV2.Contract.PancakeRouter(&_MasterChefPigsV2.CallOpts)
}

// PigsV2Token is a free data retrieval call binding the contract method 0x115af8f6.
//
// Solidity: function PigsV2Token() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PigsV2Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "PigsV2Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PigsV2Token is a free data retrieval call binding the contract method 0x115af8f6.
//
// Solidity: function PigsV2Token() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PigsV2Token() (common.Address, error) {
	return _MasterChefPigsV2.Contract.PigsV2Token(&_MasterChefPigsV2.CallOpts)
}

// PigsV2Token is a free data retrieval call binding the contract method 0x115af8f6.
//
// Solidity: function PigsV2Token() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PigsV2Token() (common.Address, error) {
	return _MasterChefPigsV2.Contract.PigsV2Token(&_MasterChefPigsV2.CallOpts)
}

// BusdCurrencyAddress is a free data retrieval call binding the contract method 0x03ba7f66.
//
// Solidity: function busdCurrencyAddress() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) BusdCurrencyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "busdCurrencyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BusdCurrencyAddress is a free data retrieval call binding the contract method 0x03ba7f66.
//
// Solidity: function busdCurrencyAddress() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) BusdCurrencyAddress() (common.Address, error) {
	return _MasterChefPigsV2.Contract.BusdCurrencyAddress(&_MasterChefPigsV2.CallOpts)
}

// BusdCurrencyAddress is a free data retrieval call binding the contract method 0x03ba7f66.
//
// Solidity: function busdCurrencyAddress() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) BusdCurrencyAddress() (common.Address, error) {
	return _MasterChefPigsV2.Contract.BusdCurrencyAddress(&_MasterChefPigsV2.CallOpts)
}

// CanHarvest is a free data retrieval call binding the contract method 0x2e6c998d.
//
// Solidity: function canHarvest(uint256 _pid, address _user) view returns(bool)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) CanHarvest(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (bool, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "canHarvest", _pid, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanHarvest is a free data retrieval call binding the contract method 0x2e6c998d.
//
// Solidity: function canHarvest(uint256 _pid, address _user) view returns(bool)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) CanHarvest(_pid *big.Int, _user common.Address) (bool, error) {
	return _MasterChefPigsV2.Contract.CanHarvest(&_MasterChefPigsV2.CallOpts, _pid, _user)
}

// CanHarvest is a free data retrieval call binding the contract method 0x2e6c998d.
//
// Solidity: function canHarvest(uint256 _pid, address _user) view returns(bool)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) CanHarvest(_pid *big.Int, _user common.Address) (bool, error) {
	return _MasterChefPigsV2.Contract.CanHarvest(&_MasterChefPigsV2.CallOpts, _pid, _user)
}

// DripBusdPid is a free data retrieval call binding the contract method 0x01e71db9.
//
// Solidity: function dripBusdPid() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) DripBusdPid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "dripBusdPid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DripBusdPid is a free data retrieval call binding the contract method 0x01e71db9.
//
// Solidity: function dripBusdPid() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) DripBusdPid() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.DripBusdPid(&_MasterChefPigsV2.CallOpts)
}

// DripBusdPid is a free data retrieval call binding the contract method 0x01e71db9.
//
// Solidity: function dripBusdPid() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) DripBusdPid() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.DripBusdPid(&_MasterChefPigsV2.CallOpts)
}

// DripTaxVault is a free data retrieval call binding the contract method 0x32b2c4bb.
//
// Solidity: function dripTaxVault() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) DripTaxVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "dripTaxVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DripTaxVault is a free data retrieval call binding the contract method 0x32b2c4bb.
//
// Solidity: function dripTaxVault() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) DripTaxVault() (common.Address, error) {
	return _MasterChefPigsV2.Contract.DripTaxVault(&_MasterChefPigsV2.CallOpts)
}

// DripTaxVault is a free data retrieval call binding the contract method 0x32b2c4bb.
//
// Solidity: function dripTaxVault() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) DripTaxVault() (common.Address, error) {
	return _MasterChefPigsV2.Contract.DripTaxVault(&_MasterChefPigsV2.CallOpts)
}

// GetPigsMultiplier is a free data retrieval call binding the contract method 0xe6eb5a8f.
//
// Solidity: function getPigsMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) GetPigsMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "getPigsMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPigsMultiplier is a free data retrieval call binding the contract method 0xe6eb5a8f.
//
// Solidity: function getPigsMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) GetPigsMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _MasterChefPigsV2.Contract.GetPigsMultiplier(&_MasterChefPigsV2.CallOpts, _from, _to)
}

// GetPigsMultiplier is a free data retrieval call binding the contract method 0xe6eb5a8f.
//
// Solidity: function getPigsMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) GetPigsMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _MasterChefPigsV2.Contract.GetPigsMultiplier(&_MasterChefPigsV2.CallOpts, _from, _to)
}

// GovAddress is a free data retrieval call binding the contract method 0x46008a07.
//
// Solidity: function govAddress() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) GovAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "govAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovAddress is a free data retrieval call binding the contract method 0x46008a07.
//
// Solidity: function govAddress() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) GovAddress() (common.Address, error) {
	return _MasterChefPigsV2.Contract.GovAddress(&_MasterChefPigsV2.CallOpts)
}

// GovAddress is a free data retrieval call binding the contract method 0x46008a07.
//
// Solidity: function govAddress() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) GovAddress() (common.Address, error) {
	return _MasterChefPigsV2.Contract.GovAddress(&_MasterChefPigsV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) Owner() (common.Address, error) {
	return _MasterChefPigsV2.Contract.Owner(&_MasterChefPigsV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) Owner() (common.Address, error) {
	return _MasterChefPigsV2.Contract.Owner(&_MasterChefPigsV2.CallOpts)
}

// OwnerPigsReward is a free data retrieval call binding the contract method 0x5cc80bd5.
//
// Solidity: function ownerPigsReward() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) OwnerPigsReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "ownerPigsReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerPigsReward is a free data retrieval call binding the contract method 0x5cc80bd5.
//
// Solidity: function ownerPigsReward() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) OwnerPigsReward() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.OwnerPigsReward(&_MasterChefPigsV2.CallOpts)
}

// OwnerPigsReward is a free data retrieval call binding the contract method 0x5cc80bd5.
//
// Solidity: function ownerPigsReward() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) OwnerPigsReward() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.OwnerPigsReward(&_MasterChefPigsV2.CallOpts)
}

// PendingPigs is a free data retrieval call binding the contract method 0xbb46ba69.
//
// Solidity: function pendingPigs(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PendingPigs(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "pendingPigs", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingPigs is a free data retrieval call binding the contract method 0xbb46ba69.
//
// Solidity: function pendingPigs(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PendingPigs(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _MasterChefPigsV2.Contract.PendingPigs(&_MasterChefPigsV2.CallOpts, _pid, _user)
}

// PendingPigs is a free data retrieval call binding the contract method 0xbb46ba69.
//
// Solidity: function pendingPigs(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PendingPigs(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _MasterChefPigsV2.Contract.PendingPigs(&_MasterChefPigsV2.CallOpts, _pid, _user)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PoolExistence(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "poolExistence", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PoolExistence(arg0 common.Address) (bool, error) {
	return _MasterChefPigsV2.Contract.PoolExistence(&_MasterChefPigsV2.CallOpts, arg0)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PoolExistence(arg0 common.Address) (bool, error) {
	return _MasterChefPigsV2.Contract.PoolExistence(&_MasterChefPigsV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accPigsPerShare, uint256 lpSupply, uint256 harvestInterval, uint256 depositFeeBP)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccPigsPerShare *big.Int
	LpSupply        *big.Int
	HarvestInterval *big.Int
	DepositFeeBP    *big.Int
}, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccPigsPerShare *big.Int
		LpSupply        *big.Int
		HarvestInterval *big.Int
		DepositFeeBP    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccPigsPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LpSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HarvestInterval = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DepositFeeBP = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accPigsPerShare, uint256 lpSupply, uint256 harvestInterval, uint256 depositFeeBP)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccPigsPerShare *big.Int
	LpSupply        *big.Int
	HarvestInterval *big.Int
	DepositFeeBP    *big.Int
}, error) {
	return _MasterChefPigsV2.Contract.PoolInfo(&_MasterChefPigsV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accPigsPerShare, uint256 lpSupply, uint256 harvestInterval, uint256 depositFeeBP)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccPigsPerShare *big.Int
	LpSupply        *big.Int
	HarvestInterval *big.Int
	DepositFeeBP    *big.Int
}, error) {
	return _MasterChefPigsV2.Contract.PoolInfo(&_MasterChefPigsV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) PoolLength() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.PoolLength(&_MasterChefPigsV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) PoolLength() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.PoolLength(&_MasterChefPigsV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) TotalAllocPoint() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.TotalAllocPoint(&_MasterChefPigsV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.TotalAllocPoint(&_MasterChefPigsV2.CallOpts)
}

// TotalLockedUpRewards is a free data retrieval call binding the contract method 0x474fa630.
//
// Solidity: function totalLockedUpRewards() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) TotalLockedUpRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "totalLockedUpRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedUpRewards is a free data retrieval call binding the contract method 0x474fa630.
//
// Solidity: function totalLockedUpRewards() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) TotalLockedUpRewards() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.TotalLockedUpRewards(&_MasterChefPigsV2.CallOpts)
}

// TotalLockedUpRewards is a free data retrieval call binding the contract method 0x474fa630.
//
// Solidity: function totalLockedUpRewards() view returns(uint256)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) TotalLockedUpRewards() (*big.Int, error) {
	return _MasterChefPigsV2.Contract.TotalLockedUpRewards(&_MasterChefPigsV2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 pigsRewardDebt, uint256 rewardLockedUp, uint256 nextHarvestUntil)
func (_MasterChefPigsV2 *MasterChefPigsV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	PigsRewardDebt   *big.Int
	RewardLockedUp   *big.Int
	NextHarvestUntil *big.Int
}, error) {
	var out []interface{}
	err := _MasterChefPigsV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount           *big.Int
		PigsRewardDebt   *big.Int
		RewardLockedUp   *big.Int
		NextHarvestUntil *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PigsRewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardLockedUp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NextHarvestUntil = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 pigsRewardDebt, uint256 rewardLockedUp, uint256 nextHarvestUntil)
func (_MasterChefPigsV2 *MasterChefPigsV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	PigsRewardDebt   *big.Int
	RewardLockedUp   *big.Int
	NextHarvestUntil *big.Int
}, error) {
	return _MasterChefPigsV2.Contract.UserInfo(&_MasterChefPigsV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 pigsRewardDebt, uint256 rewardLockedUp, uint256 nextHarvestUntil)
func (_MasterChefPigsV2 *MasterChefPigsV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	PigsRewardDebt   *big.Int
	RewardLockedUp   *big.Int
	NextHarvestUntil *big.Int
}, error) {
	return _MasterChefPigsV2.Contract.UserInfo(&_MasterChefPigsV2.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xf6e6a0e6.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, uint256 _depositFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _depositFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "add", _allocPoint, _lpToken, _depositFeeBP, _harvestInterval, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0xf6e6a0e6.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, uint256 _depositFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) Add(_allocPoint *big.Int, _lpToken common.Address, _depositFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Add(&_MasterChefPigsV2.TransactOpts, _allocPoint, _lpToken, _depositFeeBP, _harvestInterval, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0xf6e6a0e6.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, uint256 _depositFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _depositFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Add(&_MasterChefPigsV2.TransactOpts, _allocPoint, _lpToken, _depositFeeBP, _harvestInterval, _withUpdate)
}

// BurnMint is a paid mutator transaction binding the contract method 0xc6fe62b0.
//
// Solidity: function burnMint() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) BurnMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "burnMint")
}

// BurnMint is a paid mutator transaction binding the contract method 0xc6fe62b0.
//
// Solidity: function burnMint() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) BurnMint() (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.BurnMint(&_MasterChefPigsV2.TransactOpts)
}

// BurnMint is a paid mutator transaction binding the contract method 0xc6fe62b0.
//
// Solidity: function burnMint() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) BurnMint() (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.BurnMint(&_MasterChefPigsV2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Deposit(&_MasterChefPigsV2.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Deposit(&_MasterChefPigsV2.TransactOpts, _pid, _amount)
}

// DepositMigrator is a paid mutator transaction binding the contract method 0x8bf316a9.
//
// Solidity: function depositMigrator(address _userAddress, uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) DepositMigrator(opts *bind.TransactOpts, _userAddress common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "depositMigrator", _userAddress, _pid, _amount)
}

// DepositMigrator is a paid mutator transaction binding the contract method 0x8bf316a9.
//
// Solidity: function depositMigrator(address _userAddress, uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) DepositMigrator(_userAddress common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.DepositMigrator(&_MasterChefPigsV2.TransactOpts, _userAddress, _pid, _amount)
}

// DepositMigrator is a paid mutator transaction binding the contract method 0x8bf316a9.
//
// Solidity: function depositMigrator(address _userAddress, uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) DepositMigrator(_userAddress common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.DepositMigrator(&_MasterChefPigsV2.TransactOpts, _userAddress, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.EmergencyWithdraw(&_MasterChefPigsV2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.EmergencyWithdraw(&_MasterChefPigsV2.TransactOpts, _pid)
}

// IncreasePigsSupply is a paid mutator transaction binding the contract method 0xae2e3f31.
//
// Solidity: function increasePigsSupply(uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) IncreasePigsSupply(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "increasePigsSupply", _amount)
}

// IncreasePigsSupply is a paid mutator transaction binding the contract method 0xae2e3f31.
//
// Solidity: function increasePigsSupply(uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) IncreasePigsSupply(_amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.IncreasePigsSupply(&_MasterChefPigsV2.TransactOpts, _amount)
}

// IncreasePigsSupply is a paid mutator transaction binding the contract method 0xae2e3f31.
//
// Solidity: function increasePigsSupply(uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) IncreasePigsSupply(_amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.IncreasePigsSupply(&_MasterChefPigsV2.TransactOpts, _amount)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.MassUpdatePools(&_MasterChefPigsV2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.MassUpdatePools(&_MasterChefPigsV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.RenounceOwnership(&_MasterChefPigsV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.RenounceOwnership(&_MasterChefPigsV2.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60e4c4ce.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, uint256 _depositFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _depositFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "set", _pid, _allocPoint, _depositFeeBP, _harvestInterval, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x60e4c4ce.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, uint256 _depositFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) Set(_pid *big.Int, _allocPoint *big.Int, _depositFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Set(&_MasterChefPigsV2.TransactOpts, _pid, _allocPoint, _depositFeeBP, _harvestInterval, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x60e4c4ce.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, uint256 _depositFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _depositFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Set(&_MasterChefPigsV2.TransactOpts, _pid, _allocPoint, _depositFeeBP, _harvestInterval, _withUpdate)
}

// SetDDSCAAddress is a paid mutator transaction binding the contract method 0xdeaa63d8.
//
// Solidity: function setDDSCAAddress(address _ddsca) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetDDSCAAddress(opts *bind.TransactOpts, _ddsca common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setDDSCAAddress", _ddsca)
}

// SetDDSCAAddress is a paid mutator transaction binding the contract method 0xdeaa63d8.
//
// Solidity: function setDDSCAAddress(address _ddsca) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetDDSCAAddress(_ddsca common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetDDSCAAddress(&_MasterChefPigsV2.TransactOpts, _ddsca)
}

// SetDDSCAAddress is a paid mutator transaction binding the contract method 0xdeaa63d8.
//
// Solidity: function setDDSCAAddress(address _ddsca) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetDDSCAAddress(_ddsca common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetDDSCAAddress(&_MasterChefPigsV2.TransactOpts, _ddsca)
}

// SetDogPoundAutoPool is a paid mutator transaction binding the contract method 0x9299e044.
//
// Solidity: function setDogPoundAutoPool(address _dogPoundAutoPool) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetDogPoundAutoPool(opts *bind.TransactOpts, _dogPoundAutoPool common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setDogPoundAutoPool", _dogPoundAutoPool)
}

// SetDogPoundAutoPool is a paid mutator transaction binding the contract method 0x9299e044.
//
// Solidity: function setDogPoundAutoPool(address _dogPoundAutoPool) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetDogPoundAutoPool(_dogPoundAutoPool common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetDogPoundAutoPool(&_MasterChefPigsV2.TransactOpts, _dogPoundAutoPool)
}

// SetDogPoundAutoPool is a paid mutator transaction binding the contract method 0x9299e044.
//
// Solidity: function setDogPoundAutoPool(address _dogPoundAutoPool) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetDogPoundAutoPool(_dogPoundAutoPool common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetDogPoundAutoPool(&_MasterChefPigsV2.TransactOpts, _dogPoundAutoPool)
}

// SetFarmStartBlock is a paid mutator transaction binding the contract method 0xf678fc18.
//
// Solidity: function setFarmStartBlock(uint256 _newStartBlock) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetFarmStartBlock(opts *bind.TransactOpts, _newStartBlock *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setFarmStartBlock", _newStartBlock)
}

// SetFarmStartBlock is a paid mutator transaction binding the contract method 0xf678fc18.
//
// Solidity: function setFarmStartBlock(uint256 _newStartBlock) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetFarmStartBlock(_newStartBlock *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetFarmStartBlock(&_MasterChefPigsV2.TransactOpts, _newStartBlock)
}

// SetFarmStartBlock is a paid mutator transaction binding the contract method 0xf678fc18.
//
// Solidity: function setFarmStartBlock(uint256 _newStartBlock) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetFarmStartBlock(_newStartBlock *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetFarmStartBlock(&_MasterChefPigsV2.TransactOpts, _newStartBlock)
}

// SetFoundersAddresses is a paid mutator transaction binding the contract method 0x003be807.
//
// Solidity: function setFoundersAddresses(address _founder) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetFoundersAddresses(opts *bind.TransactOpts, _founder common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setFoundersAddresses", _founder)
}

// SetFoundersAddresses is a paid mutator transaction binding the contract method 0x003be807.
//
// Solidity: function setFoundersAddresses(address _founder) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetFoundersAddresses(_founder common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetFoundersAddresses(&_MasterChefPigsV2.TransactOpts, _founder)
}

// SetFoundersAddresses is a paid mutator transaction binding the contract method 0x003be807.
//
// Solidity: function setFoundersAddresses(address _founder) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetFoundersAddresses(_founder common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetFoundersAddresses(&_MasterChefPigsV2.TransactOpts, _founder)
}

// SetFoundersRewards is a paid mutator transaction binding the contract method 0x33b95b7f.
//
// Solidity: function setFoundersRewards(uint256 _newRewardsAmount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetFoundersRewards(opts *bind.TransactOpts, _newRewardsAmount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setFoundersRewards", _newRewardsAmount)
}

// SetFoundersRewards is a paid mutator transaction binding the contract method 0x33b95b7f.
//
// Solidity: function setFoundersRewards(uint256 _newRewardsAmount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetFoundersRewards(_newRewardsAmount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetFoundersRewards(&_MasterChefPigsV2.TransactOpts, _newRewardsAmount)
}

// SetFoundersRewards is a paid mutator transaction binding the contract method 0x33b95b7f.
//
// Solidity: function setFoundersRewards(uint256 _newRewardsAmount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetFoundersRewards(_newRewardsAmount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetFoundersRewards(&_MasterChefPigsV2.TransactOpts, _newRewardsAmount)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _govAddress) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetGov(opts *bind.TransactOpts, _govAddress common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setGov", _govAddress)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _govAddress) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetGov(_govAddress common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetGov(&_MasterChefPigsV2.TransactOpts, _govAddress)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _govAddress) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetGov(_govAddress common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetGov(&_MasterChefPigsV2.TransactOpts, _govAddress)
}

// SetPlatformAddress is a paid mutator transaction binding the contract method 0xcc03c342.
//
// Solidity: function setPlatformAddress(address _platformAddress) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) SetPlatformAddress(opts *bind.TransactOpts, _platformAddress common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "setPlatformAddress", _platformAddress)
}

// SetPlatformAddress is a paid mutator transaction binding the contract method 0xcc03c342.
//
// Solidity: function setPlatformAddress(address _platformAddress) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) SetPlatformAddress(_platformAddress common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetPlatformAddress(&_MasterChefPigsV2.TransactOpts, _platformAddress)
}

// SetPlatformAddress is a paid mutator transaction binding the contract method 0xcc03c342.
//
// Solidity: function setPlatformAddress(address _platformAddress) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) SetPlatformAddress(_platformAddress common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.SetPlatformAddress(&_MasterChefPigsV2.TransactOpts, _platformAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.TransferOwnership(&_MasterChefPigsV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.TransferOwnership(&_MasterChefPigsV2.TransactOpts, newOwner)
}

// UpdateEmissions is a paid mutator transaction binding the contract method 0xe39f9481.
//
// Solidity: function updateEmissions(uint256 priceInCents) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) UpdateEmissions(opts *bind.TransactOpts, priceInCents *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "updateEmissions", priceInCents)
}

// UpdateEmissions is a paid mutator transaction binding the contract method 0xe39f9481.
//
// Solidity: function updateEmissions(uint256 priceInCents) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) UpdateEmissions(priceInCents *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.UpdateEmissions(&_MasterChefPigsV2.TransactOpts, priceInCents)
}

// UpdateEmissions is a paid mutator transaction binding the contract method 0xe39f9481.
//
// Solidity: function updateEmissions(uint256 priceInCents) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) UpdateEmissions(priceInCents *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.UpdateEmissions(&_MasterChefPigsV2.TransactOpts, priceInCents)
}

// UpdateMigrator is a paid mutator transaction binding the contract method 0x28e8696a.
//
// Solidity: function updateMigrator(address _migrator) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) UpdateMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "updateMigrator", _migrator)
}

// UpdateMigrator is a paid mutator transaction binding the contract method 0x28e8696a.
//
// Solidity: function updateMigrator(address _migrator) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) UpdateMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.UpdateMigrator(&_MasterChefPigsV2.TransactOpts, _migrator)
}

// UpdateMigrator is a paid mutator transaction binding the contract method 0x28e8696a.
//
// Solidity: function updateMigrator(address _migrator) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) UpdateMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.UpdateMigrator(&_MasterChefPigsV2.TransactOpts, _migrator)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.UpdatePool(&_MasterChefPigsV2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.UpdatePool(&_MasterChefPigsV2.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Withdraw(&_MasterChefPigsV2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefPigsV2 *MasterChefPigsV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefPigsV2.Contract.Withdraw(&_MasterChefPigsV2.TransactOpts, _pid, _amount)
}

// MasterChefPigsV2AddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2AddPoolIterator struct {
	Event *MasterChefPigsV2AddPool // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2AddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2AddPool)
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
		it.Event = new(MasterChefPigsV2AddPool)
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
func (it *MasterChefPigsV2AddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2AddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2AddPool represents a AddPool event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2AddPool struct {
	Pid             *big.Int
	AllocPoint      *big.Int
	LpToken         common.Address
	DepositFeeBP    *big.Int
	HarvestInterval *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0xcb295bb60abeeb687f819f8557e33080c20eb264bce90585dd400e011adaceaf.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address lpToken, uint256 depositFeeBP, uint256 harvestInterval)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterAddPool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefPigsV2AddPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "AddPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2AddPoolIterator{contract: _MasterChefPigsV2.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0xcb295bb60abeeb687f819f8557e33080c20eb264bce90585dd400e011adaceaf.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address lpToken, uint256 depositFeeBP, uint256 harvestInterval)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2AddPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "AddPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2AddPool)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0xcb295bb60abeeb687f819f8557e33080c20eb264bce90585dd400e011adaceaf.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address lpToken, uint256 depositFeeBP, uint256 harvestInterval)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseAddPool(log types.Log) (*MasterChefPigsV2AddPool, error) {
	event := new(MasterChefPigsV2AddPool)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2DepositIterator struct {
	Event *MasterChefPigsV2Deposit // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2Deposit)
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
		it.Event = new(MasterChefPigsV2Deposit)
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
func (it *MasterChefPigsV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2Deposit represents a Deposit event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefPigsV2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2DepositIterator{contract: _MasterChefPigsV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2Deposit)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseDeposit(log types.Log) (*MasterChefPigsV2Deposit, error) {
	event := new(MasterChefPigsV2Deposit)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2DogPoundAutoPoolUpdatedIterator is returned from FilterDogPoundAutoPoolUpdated and is used to iterate over the raw logs and unpacked data for DogPoundAutoPoolUpdated events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2DogPoundAutoPoolUpdatedIterator struct {
	Event *MasterChefPigsV2DogPoundAutoPoolUpdated // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2DogPoundAutoPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2DogPoundAutoPoolUpdated)
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
		it.Event = new(MasterChefPigsV2DogPoundAutoPoolUpdated)
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
func (it *MasterChefPigsV2DogPoundAutoPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2DogPoundAutoPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2DogPoundAutoPoolUpdated represents a DogPoundAutoPoolUpdated event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2DogPoundAutoPoolUpdated struct {
	DogPoundAutoPool common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDogPoundAutoPoolUpdated is a free log retrieval operation binding the contract event 0x320f84dfd5710962958edb0276cf219923be57326bd10346dbc8b9beb5fba45c.
//
// Solidity: event DogPoundAutoPoolUpdated(address dogPoundAutoPool)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterDogPoundAutoPoolUpdated(opts *bind.FilterOpts) (*MasterChefPigsV2DogPoundAutoPoolUpdatedIterator, error) {

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "DogPoundAutoPoolUpdated")
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2DogPoundAutoPoolUpdatedIterator{contract: _MasterChefPigsV2.contract, event: "DogPoundAutoPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchDogPoundAutoPoolUpdated is a free log subscription operation binding the contract event 0x320f84dfd5710962958edb0276cf219923be57326bd10346dbc8b9beb5fba45c.
//
// Solidity: event DogPoundAutoPoolUpdated(address dogPoundAutoPool)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchDogPoundAutoPoolUpdated(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2DogPoundAutoPoolUpdated) (event.Subscription, error) {

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "DogPoundAutoPoolUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2DogPoundAutoPoolUpdated)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "DogPoundAutoPoolUpdated", log); err != nil {
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

// ParseDogPoundAutoPoolUpdated is a log parse operation binding the contract event 0x320f84dfd5710962958edb0276cf219923be57326bd10346dbc8b9beb5fba45c.
//
// Solidity: event DogPoundAutoPoolUpdated(address dogPoundAutoPool)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseDogPoundAutoPoolUpdated(log types.Log) (*MasterChefPigsV2DogPoundAutoPoolUpdated, error) {
	event := new(MasterChefPigsV2DogPoundAutoPoolUpdated)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "DogPoundAutoPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2EmergencyWithdrawIterator struct {
	Event *MasterChefPigsV2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2EmergencyWithdraw)
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
		it.Event = new(MasterChefPigsV2EmergencyWithdraw)
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
func (it *MasterChefPigsV2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2EmergencyWithdraw represents a EmergencyWithdraw event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefPigsV2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2EmergencyWithdrawIterator{contract: _MasterChefPigsV2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2EmergencyWithdraw)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseEmergencyWithdraw(log types.Log) (*MasterChefPigsV2EmergencyWithdraw, error) {
	event := new(MasterChefPigsV2EmergencyWithdraw)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2GovUpdatedIterator is returned from FilterGovUpdated and is used to iterate over the raw logs and unpacked data for GovUpdated events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2GovUpdatedIterator struct {
	Event *MasterChefPigsV2GovUpdated // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2GovUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2GovUpdated)
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
		it.Event = new(MasterChefPigsV2GovUpdated)
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
func (it *MasterChefPigsV2GovUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2GovUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2GovUpdated represents a GovUpdated event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2GovUpdated struct {
	GovAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGovUpdated is a free log retrieval operation binding the contract event 0xb49a738bf7c18189b8cd3a6cd9d2b44045ca8070a4fa7e9db46ecfcbedee6bd9.
//
// Solidity: event GovUpdated(address govAddress)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterGovUpdated(opts *bind.FilterOpts) (*MasterChefPigsV2GovUpdatedIterator, error) {

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "GovUpdated")
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2GovUpdatedIterator{contract: _MasterChefPigsV2.contract, event: "GovUpdated", logs: logs, sub: sub}, nil
}

// WatchGovUpdated is a free log subscription operation binding the contract event 0xb49a738bf7c18189b8cd3a6cd9d2b44045ca8070a4fa7e9db46ecfcbedee6bd9.
//
// Solidity: event GovUpdated(address govAddress)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchGovUpdated(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2GovUpdated) (event.Subscription, error) {

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "GovUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2GovUpdated)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "GovUpdated", log); err != nil {
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

// ParseGovUpdated is a log parse operation binding the contract event 0xb49a738bf7c18189b8cd3a6cd9d2b44045ca8070a4fa7e9db46ecfcbedee6bd9.
//
// Solidity: event GovUpdated(address govAddress)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseGovUpdated(log types.Log) (*MasterChefPigsV2GovUpdated, error) {
	event := new(MasterChefPigsV2GovUpdated)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "GovUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2OwnershipTransferredIterator struct {
	Event *MasterChefPigsV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2OwnershipTransferred)
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
		it.Event = new(MasterChefPigsV2OwnershipTransferred)
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
func (it *MasterChefPigsV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2OwnershipTransferred represents a OwnershipTransferred event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterChefPigsV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2OwnershipTransferredIterator{contract: _MasterChefPigsV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2OwnershipTransferred)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseOwnershipTransferred(log types.Log) (*MasterChefPigsV2OwnershipTransferred, error) {
	event := new(MasterChefPigsV2OwnershipTransferred)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2RewardLockedUpIterator is returned from FilterRewardLockedUp and is used to iterate over the raw logs and unpacked data for RewardLockedUp events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2RewardLockedUpIterator struct {
	Event *MasterChefPigsV2RewardLockedUp // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2RewardLockedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2RewardLockedUp)
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
		it.Event = new(MasterChefPigsV2RewardLockedUp)
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
func (it *MasterChefPigsV2RewardLockedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2RewardLockedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2RewardLockedUp represents a RewardLockedUp event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2RewardLockedUp struct {
	User           common.Address
	Pid            *big.Int
	AmountLockedUp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardLockedUp is a free log retrieval operation binding the contract event 0xee470483107f579a55c754fa00613c45a9a3b617a418b39cb0be97e5381ba7c1.
//
// Solidity: event RewardLockedUp(address indexed user, uint256 indexed pid, uint256 amountLockedUp)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterRewardLockedUp(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefPigsV2RewardLockedUpIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "RewardLockedUp", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2RewardLockedUpIterator{contract: _MasterChefPigsV2.contract, event: "RewardLockedUp", logs: logs, sub: sub}, nil
}

// WatchRewardLockedUp is a free log subscription operation binding the contract event 0xee470483107f579a55c754fa00613c45a9a3b617a418b39cb0be97e5381ba7c1.
//
// Solidity: event RewardLockedUp(address indexed user, uint256 indexed pid, uint256 amountLockedUp)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchRewardLockedUp(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2RewardLockedUp, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "RewardLockedUp", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2RewardLockedUp)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "RewardLockedUp", log); err != nil {
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

// ParseRewardLockedUp is a log parse operation binding the contract event 0xee470483107f579a55c754fa00613c45a9a3b617a418b39cb0be97e5381ba7c1.
//
// Solidity: event RewardLockedUp(address indexed user, uint256 indexed pid, uint256 amountLockedUp)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseRewardLockedUp(log types.Log) (*MasterChefPigsV2RewardLockedUp, error) {
	event := new(MasterChefPigsV2RewardLockedUp)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "RewardLockedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2SetFounderIterator is returned from FilterSetFounder and is used to iterate over the raw logs and unpacked data for SetFounder events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetFounderIterator struct {
	Event *MasterChefPigsV2SetFounder // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2SetFounderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2SetFounder)
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
		it.Event = new(MasterChefPigsV2SetFounder)
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
func (it *MasterChefPigsV2SetFounderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2SetFounderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2SetFounder represents a SetFounder event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetFounder struct {
	Founder common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFounder is a free log retrieval operation binding the contract event 0xc5c86d65a14ecb5e8ef453185dbb744f2c56de556e089660a50245b96d893b2c.
//
// Solidity: event SetFounder(address founder)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterSetFounder(opts *bind.FilterOpts) (*MasterChefPigsV2SetFounderIterator, error) {

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "SetFounder")
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2SetFounderIterator{contract: _MasterChefPigsV2.contract, event: "SetFounder", logs: logs, sub: sub}, nil
}

// WatchSetFounder is a free log subscription operation binding the contract event 0xc5c86d65a14ecb5e8ef453185dbb744f2c56de556e089660a50245b96d893b2c.
//
// Solidity: event SetFounder(address founder)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchSetFounder(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2SetFounder) (event.Subscription, error) {

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "SetFounder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2SetFounder)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetFounder", log); err != nil {
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

// ParseSetFounder is a log parse operation binding the contract event 0xc5c86d65a14ecb5e8ef453185dbb744f2c56de556e089660a50245b96d893b2c.
//
// Solidity: event SetFounder(address founder)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseSetFounder(log types.Log) (*MasterChefPigsV2SetFounder, error) {
	event := new(MasterChefPigsV2SetFounder)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetFounder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2SetOwnersRewardsIterator is returned from FilterSetOwnersRewards and is used to iterate over the raw logs and unpacked data for SetOwnersRewards events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetOwnersRewardsIterator struct {
	Event *MasterChefPigsV2SetOwnersRewards // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2SetOwnersRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2SetOwnersRewards)
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
		it.Event = new(MasterChefPigsV2SetOwnersRewards)
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
func (it *MasterChefPigsV2SetOwnersRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2SetOwnersRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2SetOwnersRewards represents a SetOwnersRewards event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetOwnersRewards struct {
	OwnerReward *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetOwnersRewards is a free log retrieval operation binding the contract event 0x5a361f311b5ba9aa1cec0efbcc3da6efff4fa398700201ca14194deedb673464.
//
// Solidity: event SetOwnersRewards(uint256 ownerReward)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterSetOwnersRewards(opts *bind.FilterOpts) (*MasterChefPigsV2SetOwnersRewardsIterator, error) {

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "SetOwnersRewards")
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2SetOwnersRewardsIterator{contract: _MasterChefPigsV2.contract, event: "SetOwnersRewards", logs: logs, sub: sub}, nil
}

// WatchSetOwnersRewards is a free log subscription operation binding the contract event 0x5a361f311b5ba9aa1cec0efbcc3da6efff4fa398700201ca14194deedb673464.
//
// Solidity: event SetOwnersRewards(uint256 ownerReward)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchSetOwnersRewards(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2SetOwnersRewards) (event.Subscription, error) {

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "SetOwnersRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2SetOwnersRewards)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetOwnersRewards", log); err != nil {
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

// ParseSetOwnersRewards is a log parse operation binding the contract event 0x5a361f311b5ba9aa1cec0efbcc3da6efff4fa398700201ca14194deedb673464.
//
// Solidity: event SetOwnersRewards(uint256 ownerReward)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseSetOwnersRewards(log types.Log) (*MasterChefPigsV2SetOwnersRewards, error) {
	event := new(MasterChefPigsV2SetOwnersRewards)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetOwnersRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2SetPlatformAddressIterator is returned from FilterSetPlatformAddress and is used to iterate over the raw logs and unpacked data for SetPlatformAddress events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetPlatformAddressIterator struct {
	Event *MasterChefPigsV2SetPlatformAddress // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2SetPlatformAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2SetPlatformAddress)
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
		it.Event = new(MasterChefPigsV2SetPlatformAddress)
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
func (it *MasterChefPigsV2SetPlatformAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2SetPlatformAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2SetPlatformAddress represents a SetPlatformAddress event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetPlatformAddress struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPlatformAddress is a free log retrieval operation binding the contract event 0xd2991d7f16266fc3594a7a085f0da1060fde4dba356999af15ced38362eee36c.
//
// Solidity: event SetPlatformAddress(address indexed newAddress)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterSetPlatformAddress(opts *bind.FilterOpts, newAddress []common.Address) (*MasterChefPigsV2SetPlatformAddressIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "SetPlatformAddress", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2SetPlatformAddressIterator{contract: _MasterChefPigsV2.contract, event: "SetPlatformAddress", logs: logs, sub: sub}, nil
}

// WatchSetPlatformAddress is a free log subscription operation binding the contract event 0xd2991d7f16266fc3594a7a085f0da1060fde4dba356999af15ced38362eee36c.
//
// Solidity: event SetPlatformAddress(address indexed newAddress)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchSetPlatformAddress(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2SetPlatformAddress, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "SetPlatformAddress", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2SetPlatformAddress)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetPlatformAddress", log); err != nil {
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

// ParseSetPlatformAddress is a log parse operation binding the contract event 0xd2991d7f16266fc3594a7a085f0da1060fde4dba356999af15ced38362eee36c.
//
// Solidity: event SetPlatformAddress(address indexed newAddress)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseSetPlatformAddress(log types.Log) (*MasterChefPigsV2SetPlatformAddress, error) {
	event := new(MasterChefPigsV2SetPlatformAddress)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetPlatformAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2SetPoolIterator is returned from FilterSetPool and is used to iterate over the raw logs and unpacked data for SetPool events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetPoolIterator struct {
	Event *MasterChefPigsV2SetPool // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2SetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2SetPool)
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
		it.Event = new(MasterChefPigsV2SetPool)
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
func (it *MasterChefPigsV2SetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2SetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2SetPool represents a SetPool event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2SetPool struct {
	Pid             *big.Int
	AllocPoint      *big.Int
	DepositFeeBP    *big.Int
	HarvestInterval *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetPool is a free log retrieval operation binding the contract event 0x638f84b26ab0fed05e944aa4ad0d38a62a9fc8234984bbb9e0091a676bf305f4.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint, uint256 depositFeeBP, uint256 harvestInterval)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterSetPool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefPigsV2SetPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2SetPoolIterator{contract: _MasterChefPigsV2.contract, event: "SetPool", logs: logs, sub: sub}, nil
}

// WatchSetPool is a free log subscription operation binding the contract event 0x638f84b26ab0fed05e944aa4ad0d38a62a9fc8234984bbb9e0091a676bf305f4.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint, uint256 depositFeeBP, uint256 harvestInterval)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchSetPool(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2SetPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2SetPool)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetPool", log); err != nil {
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

// ParseSetPool is a log parse operation binding the contract event 0x638f84b26ab0fed05e944aa4ad0d38a62a9fc8234984bbb9e0091a676bf305f4.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint, uint256 depositFeeBP, uint256 harvestInterval)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseSetPool(log types.Log) (*MasterChefPigsV2SetPool, error) {
	event := new(MasterChefPigsV2SetPool)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "SetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefPigsV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2WithdrawIterator struct {
	Event *MasterChefPigsV2Withdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefPigsV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefPigsV2Withdraw)
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
		it.Event = new(MasterChefPigsV2Withdraw)
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
func (it *MasterChefPigsV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefPigsV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefPigsV2Withdraw represents a Withdraw event raised by the MasterChefPigsV2 contract.
type MasterChefPigsV2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefPigsV2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefPigsV2WithdrawIterator{contract: _MasterChefPigsV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefPigsV2Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefPigsV2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefPigsV2Withdraw)
				if err := _MasterChefPigsV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefPigsV2 *MasterChefPigsV2Filterer) ParseWithdraw(log types.Log) (*MasterChefPigsV2Withdraw, error) {
	event := new(MasterChefPigsV2Withdraw)
	if err := _MasterChefPigsV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
