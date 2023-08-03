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

// MasterChefDogsV2migrationInfo is an auto generated low-level Go binding around an user-defined struct.
type MasterChefDogsV2migrationInfo struct {
	LpToken      common.Address
	AmountStaked *big.Int
}

// MasterChefDogsV2MetaData contains all meta data concerning the MasterChefDogsV2 contract.
var MasterChefDogsV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDogsTokenV2\",\"name\":\"_dogsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_platform\",\"type\":\"address\"},{\"internalType\":\"contractIDDSCA\",\"name\":\"_ddsca\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLPToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestInterval\",\"type\":\"uint256\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"govAddress\",\"type\":\"address\"}],\"name\":\"GovUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLockedUp\",\"type\":\"uint256\"}],\"name\":\"RewardLockedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dogsAddress\",\"type\":\"address\"}],\"name\":\"SetDogsReferral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetPlatformAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestInterval\",\"type\":\"uint256\"}],\"name\":\"SetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAXIMUM_HARVEST_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PLATFORM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLPToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestInterval\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_usersStakeData\",\"type\":\"uint256[]\"}],\"name\":\"addPoolUserData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"canHarvest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"canMigrate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"canMigratePools\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"}],\"internalType\":\"structMasterChefDogsV2.migrationInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"depositMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dogsToken\",\"outputs\":[{\"internalType\":\"contractDogsTokenV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"contractIFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getDogsMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseDogsSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBurned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingDogs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformnotLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accDogsPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"harvestInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLPToken\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referralCommissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_harvestInterval\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDDSCA\",\"name\":\"_ddsca\",\"type\":\"address\"}],\"name\":\"setDDSCAAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartBlock\",\"type\":\"uint256\"}],\"name\":\"setFarmStartBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_govAddress\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformAddress\",\"type\":\"address\"}],\"name\":\"setPlatformAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIReferralSystem\",\"name\":\"_dogsReferral\",\"type\":\"address\"}],\"name\":\"setReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"toggleMigrationEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedUpRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInCents\",\"type\":\"uint256\"}],\"name\":\"updateEmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFeeManager\",\"name\":\"_feeManagerAddress\",\"type\":\"address\"}],\"name\":\"updateFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dogsRewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardLockedUp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextHarvestUntil\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMigrationInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterChefDogsV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterChefDogsV2MetaData.ABI instead.
var MasterChefDogsV2ABI = MasterChefDogsV2MetaData.ABI

// MasterChefDogsV2 is an auto generated Go binding around an Ethereum contract.
type MasterChefDogsV2 struct {
	MasterChefDogsV2Caller     // Read-only binding to the contract
	MasterChefDogsV2Transactor // Write-only binding to the contract
	MasterChefDogsV2Filterer   // Log filterer for contract events
}

// MasterChefDogsV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasterChefDogsV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefDogsV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterChefDogsV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefDogsV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterChefDogsV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefDogsV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterChefDogsV2Session struct {
	Contract     *MasterChefDogsV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterChefDogsV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterChefDogsV2CallerSession struct {
	Contract *MasterChefDogsV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MasterChefDogsV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterChefDogsV2TransactorSession struct {
	Contract     *MasterChefDogsV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MasterChefDogsV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasterChefDogsV2Raw struct {
	Contract *MasterChefDogsV2 // Generic contract binding to access the raw methods on
}

// MasterChefDogsV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterChefDogsV2CallerRaw struct {
	Contract *MasterChefDogsV2Caller // Generic read-only contract binding to access the raw methods on
}

// MasterChefDogsV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterChefDogsV2TransactorRaw struct {
	Contract *MasterChefDogsV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterChefDogsV2 creates a new instance of MasterChefDogsV2, bound to a specific deployed contract.
func NewMasterChefDogsV2(address common.Address, backend bind.ContractBackend) (*MasterChefDogsV2, error) {
	contract, err := bindMasterChefDogsV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2{MasterChefDogsV2Caller: MasterChefDogsV2Caller{contract: contract}, MasterChefDogsV2Transactor: MasterChefDogsV2Transactor{contract: contract}, MasterChefDogsV2Filterer: MasterChefDogsV2Filterer{contract: contract}}, nil
}

// NewMasterChefDogsV2Caller creates a new read-only instance of MasterChefDogsV2, bound to a specific deployed contract.
func NewMasterChefDogsV2Caller(address common.Address, caller bind.ContractCaller) (*MasterChefDogsV2Caller, error) {
	contract, err := bindMasterChefDogsV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2Caller{contract: contract}, nil
}

// NewMasterChefDogsV2Transactor creates a new write-only instance of MasterChefDogsV2, bound to a specific deployed contract.
func NewMasterChefDogsV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MasterChefDogsV2Transactor, error) {
	contract, err := bindMasterChefDogsV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2Transactor{contract: contract}, nil
}

// NewMasterChefDogsV2Filterer creates a new log filterer instance of MasterChefDogsV2, bound to a specific deployed contract.
func NewMasterChefDogsV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MasterChefDogsV2Filterer, error) {
	contract, err := bindMasterChefDogsV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2Filterer{contract: contract}, nil
}

// bindMasterChefDogsV2 binds a generic wrapper to an already deployed contract.
func bindMasterChefDogsV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterChefDogsV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChefDogsV2 *MasterChefDogsV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChefDogsV2.Contract.MasterChefDogsV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChefDogsV2 *MasterChefDogsV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.MasterChefDogsV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChefDogsV2 *MasterChefDogsV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.MasterChefDogsV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChefDogsV2 *MasterChefDogsV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChefDogsV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.contract.Transact(opts, method, params...)
}

// MAXIMUMHARVESTINTERVAL is a free data retrieval call binding the contract method 0xde73149d.
//
// Solidity: function MAXIMUM_HARVEST_INTERVAL() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) MAXIMUMHARVESTINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "MAXIMUM_HARVEST_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMHARVESTINTERVAL is a free data retrieval call binding the contract method 0xde73149d.
//
// Solidity: function MAXIMUM_HARVEST_INTERVAL() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) MAXIMUMHARVESTINTERVAL() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.MAXIMUMHARVESTINTERVAL(&_MasterChefDogsV2.CallOpts)
}

// MAXIMUMHARVESTINTERVAL is a free data retrieval call binding the contract method 0xde73149d.
//
// Solidity: function MAXIMUM_HARVEST_INTERVAL() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) MAXIMUMHARVESTINTERVAL() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.MAXIMUMHARVESTINTERVAL(&_MasterChefDogsV2.CallOpts)
}

// PLATFORMADDRESS is a free data retrieval call binding the contract method 0xbbc28489.
//
// Solidity: function PLATFORM_ADDRESS() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) PLATFORMADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "PLATFORM_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PLATFORMADDRESS is a free data retrieval call binding the contract method 0xbbc28489.
//
// Solidity: function PLATFORM_ADDRESS() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) PLATFORMADDRESS() (common.Address, error) {
	return _MasterChefDogsV2.Contract.PLATFORMADDRESS(&_MasterChefDogsV2.CallOpts)
}

// PLATFORMADDRESS is a free data retrieval call binding the contract method 0xbbc28489.
//
// Solidity: function PLATFORM_ADDRESS() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) PLATFORMADDRESS() (common.Address, error) {
	return _MasterChefDogsV2.Contract.PLATFORMADDRESS(&_MasterChefDogsV2.CallOpts)
}

// CanHarvest is a free data retrieval call binding the contract method 0x2e6c998d.
//
// Solidity: function canHarvest(uint256 _pid, address _user) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) CanHarvest(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (bool, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "canHarvest", _pid, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanHarvest is a free data retrieval call binding the contract method 0x2e6c998d.
//
// Solidity: function canHarvest(uint256 _pid, address _user) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) CanHarvest(_pid *big.Int, _user common.Address) (bool, error) {
	return _MasterChefDogsV2.Contract.CanHarvest(&_MasterChefDogsV2.CallOpts, _pid, _user)
}

// CanHarvest is a free data retrieval call binding the contract method 0x2e6c998d.
//
// Solidity: function canHarvest(uint256 _pid, address _user) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) CanHarvest(_pid *big.Int, _user common.Address) (bool, error) {
	return _MasterChefDogsV2.Contract.CanHarvest(&_MasterChefDogsV2.CallOpts, _pid, _user)
}

// CanMigrate is a free data retrieval call binding the contract method 0x7de07cea.
//
// Solidity: function canMigrate(address _address) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) CanMigrate(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "canMigrate", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanMigrate is a free data retrieval call binding the contract method 0x7de07cea.
//
// Solidity: function canMigrate(address _address) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) CanMigrate(_address common.Address) (bool, error) {
	return _MasterChefDogsV2.Contract.CanMigrate(&_MasterChefDogsV2.CallOpts, _address)
}

// CanMigrate is a free data retrieval call binding the contract method 0x7de07cea.
//
// Solidity: function canMigrate(address _address) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) CanMigrate(_address common.Address) (bool, error) {
	return _MasterChefDogsV2.Contract.CanMigrate(&_MasterChefDogsV2.CallOpts, _address)
}

// CanMigratePools is a free data retrieval call binding the contract method 0xfcbc5eee.
//
// Solidity: function canMigratePools(address _address) view returns((address,uint256)[])
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) CanMigratePools(opts *bind.CallOpts, _address common.Address) ([]MasterChefDogsV2migrationInfo, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "canMigratePools", _address)

	if err != nil {
		return *new([]MasterChefDogsV2migrationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]MasterChefDogsV2migrationInfo)).(*[]MasterChefDogsV2migrationInfo)

	return out0, err

}

// CanMigratePools is a free data retrieval call binding the contract method 0xfcbc5eee.
//
// Solidity: function canMigratePools(address _address) view returns((address,uint256)[])
func (_MasterChefDogsV2 *MasterChefDogsV2Session) CanMigratePools(_address common.Address) ([]MasterChefDogsV2migrationInfo, error) {
	return _MasterChefDogsV2.Contract.CanMigratePools(&_MasterChefDogsV2.CallOpts, _address)
}

// CanMigratePools is a free data retrieval call binding the contract method 0xfcbc5eee.
//
// Solidity: function canMigratePools(address _address) view returns((address,uint256)[])
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) CanMigratePools(_address common.Address) ([]MasterChefDogsV2migrationInfo, error) {
	return _MasterChefDogsV2.Contract.CanMigratePools(&_MasterChefDogsV2.CallOpts, _address)
}

// DogsToken is a free data retrieval call binding the contract method 0x5e2eeb35.
//
// Solidity: function dogsToken() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) DogsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "dogsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogsToken is a free data retrieval call binding the contract method 0x5e2eeb35.
//
// Solidity: function dogsToken() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) DogsToken() (common.Address, error) {
	return _MasterChefDogsV2.Contract.DogsToken(&_MasterChefDogsV2.CallOpts)
}

// DogsToken is a free data retrieval call binding the contract method 0x5e2eeb35.
//
// Solidity: function dogsToken() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) DogsToken() (common.Address, error) {
	return _MasterChefDogsV2.Contract.DogsToken(&_MasterChefDogsV2.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) FeeManager() (common.Address, error) {
	return _MasterChefDogsV2.Contract.FeeManager(&_MasterChefDogsV2.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) FeeManager() (common.Address, error) {
	return _MasterChefDogsV2.Contract.FeeManager(&_MasterChefDogsV2.CallOpts)
}

// GetDogsMultiplier is a free data retrieval call binding the contract method 0x36a3bced.
//
// Solidity: function getDogsMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) GetDogsMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "getDogsMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDogsMultiplier is a free data retrieval call binding the contract method 0x36a3bced.
//
// Solidity: function getDogsMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) GetDogsMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _MasterChefDogsV2.Contract.GetDogsMultiplier(&_MasterChefDogsV2.CallOpts, _from, _to)
}

// GetDogsMultiplier is a free data retrieval call binding the contract method 0x36a3bced.
//
// Solidity: function getDogsMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) GetDogsMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _MasterChefDogsV2.Contract.GetDogsMultiplier(&_MasterChefDogsV2.CallOpts, _from, _to)
}

// GovAddress is a free data retrieval call binding the contract method 0x46008a07.
//
// Solidity: function govAddress() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) GovAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "govAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovAddress is a free data retrieval call binding the contract method 0x46008a07.
//
// Solidity: function govAddress() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) GovAddress() (common.Address, error) {
	return _MasterChefDogsV2.Contract.GovAddress(&_MasterChefDogsV2.CallOpts)
}

// GovAddress is a free data retrieval call binding the contract method 0x46008a07.
//
// Solidity: function govAddress() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) GovAddress() (common.Address, error) {
	return _MasterChefDogsV2.Contract.GovAddress(&_MasterChefDogsV2.CallOpts)
}

// MigrationEnabled is a free data retrieval call binding the contract method 0x35b944bf.
//
// Solidity: function migrationEnabled() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) MigrationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "migrationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MigrationEnabled is a free data retrieval call binding the contract method 0x35b944bf.
//
// Solidity: function migrationEnabled() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) MigrationEnabled() (bool, error) {
	return _MasterChefDogsV2.Contract.MigrationEnabled(&_MasterChefDogsV2.CallOpts)
}

// MigrationEnabled is a free data retrieval call binding the contract method 0x35b944bf.
//
// Solidity: function migrationEnabled() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) MigrationEnabled() (bool, error) {
	return _MasterChefDogsV2.Contract.MigrationEnabled(&_MasterChefDogsV2.CallOpts)
}

// MintBurned is a free data retrieval call binding the contract method 0xc3643025.
//
// Solidity: function mintBurned() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) MintBurned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "mintBurned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintBurned is a free data retrieval call binding the contract method 0xc3643025.
//
// Solidity: function mintBurned() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) MintBurned() (bool, error) {
	return _MasterChefDogsV2.Contract.MintBurned(&_MasterChefDogsV2.CallOpts)
}

// MintBurned is a free data retrieval call binding the contract method 0xc3643025.
//
// Solidity: function mintBurned() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) MintBurned() (bool, error) {
	return _MasterChefDogsV2.Contract.MintBurned(&_MasterChefDogsV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) Owner() (common.Address, error) {
	return _MasterChefDogsV2.Contract.Owner(&_MasterChefDogsV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) Owner() (common.Address, error) {
	return _MasterChefDogsV2.Contract.Owner(&_MasterChefDogsV2.CallOpts)
}

// PendingDogs is a free data retrieval call binding the contract method 0x48f31825.
//
// Solidity: function pendingDogs(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) PendingDogs(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "pendingDogs", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingDogs is a free data retrieval call binding the contract method 0x48f31825.
//
// Solidity: function pendingDogs(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) PendingDogs(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _MasterChefDogsV2.Contract.PendingDogs(&_MasterChefDogsV2.CallOpts, _pid, _user)
}

// PendingDogs is a free data retrieval call binding the contract method 0x48f31825.
//
// Solidity: function pendingDogs(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) PendingDogs(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _MasterChefDogsV2.Contract.PendingDogs(&_MasterChefDogsV2.CallOpts, _pid, _user)
}

// PlatformnotLocked is a free data retrieval call binding the contract method 0xee13e62b.
//
// Solidity: function platformnotLocked() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) PlatformnotLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "platformnotLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PlatformnotLocked is a free data retrieval call binding the contract method 0xee13e62b.
//
// Solidity: function platformnotLocked() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) PlatformnotLocked() (bool, error) {
	return _MasterChefDogsV2.Contract.PlatformnotLocked(&_MasterChefDogsV2.CallOpts)
}

// PlatformnotLocked is a free data retrieval call binding the contract method 0xee13e62b.
//
// Solidity: function platformnotLocked() view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) PlatformnotLocked() (bool, error) {
	return _MasterChefDogsV2.Contract.PlatformnotLocked(&_MasterChefDogsV2.CallOpts)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) PoolExistence(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "poolExistence", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) PoolExistence(arg0 common.Address) (bool, error) {
	return _MasterChefDogsV2.Contract.PoolExistence(&_MasterChefDogsV2.CallOpts, arg0)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) PoolExistence(arg0 common.Address) (bool, error) {
	return _MasterChefDogsV2.Contract.PoolExistence(&_MasterChefDogsV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, address strategy, uint256 allocPoint, uint256 lastRewardBlock, uint256 accDogsPerShare, uint256 lpSupply, uint256 harvestInterval, uint256 depositFeeBP, uint256 withdrawFeeBP, bool isLPToken)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	Strategy        common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccDogsPerShare *big.Int
	LpSupply        *big.Int
	HarvestInterval *big.Int
	DepositFeeBP    *big.Int
	WithdrawFeeBP   *big.Int
	IsLPToken       bool
}, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		Strategy        common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccDogsPerShare *big.Int
		LpSupply        *big.Int
		HarvestInterval *big.Int
		DepositFeeBP    *big.Int
		WithdrawFeeBP   *big.Int
		IsLPToken       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Strategy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccDogsPerShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LpSupply = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.HarvestInterval = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DepositFeeBP = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.WithdrawFeeBP = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.IsLPToken = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, address strategy, uint256 allocPoint, uint256 lastRewardBlock, uint256 accDogsPerShare, uint256 lpSupply, uint256 harvestInterval, uint256 depositFeeBP, uint256 withdrawFeeBP, bool isLPToken)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	Strategy        common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccDogsPerShare *big.Int
	LpSupply        *big.Int
	HarvestInterval *big.Int
	DepositFeeBP    *big.Int
	WithdrawFeeBP   *big.Int
	IsLPToken       bool
}, error) {
	return _MasterChefDogsV2.Contract.PoolInfo(&_MasterChefDogsV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, address strategy, uint256 allocPoint, uint256 lastRewardBlock, uint256 accDogsPerShare, uint256 lpSupply, uint256 harvestInterval, uint256 depositFeeBP, uint256 withdrawFeeBP, bool isLPToken)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	Strategy        common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccDogsPerShare *big.Int
	LpSupply        *big.Int
	HarvestInterval *big.Int
	DepositFeeBP    *big.Int
	WithdrawFeeBP   *big.Int
	IsLPToken       bool
}, error) {
	return _MasterChefDogsV2.Contract.PoolInfo(&_MasterChefDogsV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) PoolLength() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.PoolLength(&_MasterChefDogsV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) PoolLength() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.PoolLength(&_MasterChefDogsV2.CallOpts)
}

// ReferralCommissionRate is a free data retrieval call binding the contract method 0xd30ef61b.
//
// Solidity: function referralCommissionRate() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) ReferralCommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "referralCommissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferralCommissionRate is a free data retrieval call binding the contract method 0xd30ef61b.
//
// Solidity: function referralCommissionRate() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) ReferralCommissionRate() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.ReferralCommissionRate(&_MasterChefDogsV2.CallOpts)
}

// ReferralCommissionRate is a free data retrieval call binding the contract method 0xd30ef61b.
//
// Solidity: function referralCommissionRate() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) ReferralCommissionRate() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.ReferralCommissionRate(&_MasterChefDogsV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) TotalAllocPoint() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.TotalAllocPoint(&_MasterChefDogsV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.TotalAllocPoint(&_MasterChefDogsV2.CallOpts)
}

// TotalLockedUpRewards is a free data retrieval call binding the contract method 0x474fa630.
//
// Solidity: function totalLockedUpRewards() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) TotalLockedUpRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "totalLockedUpRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedUpRewards is a free data retrieval call binding the contract method 0x474fa630.
//
// Solidity: function totalLockedUpRewards() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) TotalLockedUpRewards() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.TotalLockedUpRewards(&_MasterChefDogsV2.CallOpts)
}

// TotalLockedUpRewards is a free data retrieval call binding the contract method 0x474fa630.
//
// Solidity: function totalLockedUpRewards() view returns(uint256)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) TotalLockedUpRewards() (*big.Int, error) {
	return _MasterChefDogsV2.Contract.TotalLockedUpRewards(&_MasterChefDogsV2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 dogsRewardDebt, uint256 rewardLockedUp, uint256 nextHarvestUntil)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	DogsRewardDebt   *big.Int
	RewardLockedUp   *big.Int
	NextHarvestUntil *big.Int
}, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount           *big.Int
		DogsRewardDebt   *big.Int
		RewardLockedUp   *big.Int
		NextHarvestUntil *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DogsRewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardLockedUp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NextHarvestUntil = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 dogsRewardDebt, uint256 rewardLockedUp, uint256 nextHarvestUntil)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	DogsRewardDebt   *big.Int
	RewardLockedUp   *big.Int
	NextHarvestUntil *big.Int
}, error) {
	return _MasterChefDogsV2.Contract.UserInfo(&_MasterChefDogsV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 dogsRewardDebt, uint256 rewardLockedUp, uint256 nextHarvestUntil)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	DogsRewardDebt   *big.Int
	RewardLockedUp   *big.Int
	NextHarvestUntil *big.Int
}, error) {
	return _MasterChefDogsV2.Contract.UserInfo(&_MasterChefDogsV2.CallOpts, arg0, arg1)
}

// UserMigrationInfo is a free data retrieval call binding the contract method 0xef1d7542.
//
// Solidity: function userMigrationInfo(uint256 , address ) view returns(uint256 amountStaked)
func (_MasterChefDogsV2 *MasterChefDogsV2Caller) UserMigrationInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefDogsV2.contract.Call(opts, &out, "userMigrationInfo", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMigrationInfo is a free data retrieval call binding the contract method 0xef1d7542.
//
// Solidity: function userMigrationInfo(uint256 , address ) view returns(uint256 amountStaked)
func (_MasterChefDogsV2 *MasterChefDogsV2Session) UserMigrationInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _MasterChefDogsV2.Contract.UserMigrationInfo(&_MasterChefDogsV2.CallOpts, arg0, arg1)
}

// UserMigrationInfo is a free data retrieval call binding the contract method 0xef1d7542.
//
// Solidity: function userMigrationInfo(uint256 , address ) view returns(uint256 amountStaked)
func (_MasterChefDogsV2 *MasterChefDogsV2CallerSession) UserMigrationInfo(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _MasterChefDogsV2.Contract.UserMigrationInfo(&_MasterChefDogsV2.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x748e1370.
//
// Solidity: function add(bool _isLPToken, uint256 _allocPoint, address _lpToken, address _strategy, uint256 _depositFeeBP, uint256 _withdrawFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) Add(opts *bind.TransactOpts, _isLPToken bool, _allocPoint *big.Int, _lpToken common.Address, _strategy common.Address, _depositFeeBP *big.Int, _withdrawFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "add", _isLPToken, _allocPoint, _lpToken, _strategy, _depositFeeBP, _withdrawFeeBP, _harvestInterval, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x748e1370.
//
// Solidity: function add(bool _isLPToken, uint256 _allocPoint, address _lpToken, address _strategy, uint256 _depositFeeBP, uint256 _withdrawFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) Add(_isLPToken bool, _allocPoint *big.Int, _lpToken common.Address, _strategy common.Address, _depositFeeBP *big.Int, _withdrawFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Add(&_MasterChefDogsV2.TransactOpts, _isLPToken, _allocPoint, _lpToken, _strategy, _depositFeeBP, _withdrawFeeBP, _harvestInterval, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x748e1370.
//
// Solidity: function add(bool _isLPToken, uint256 _allocPoint, address _lpToken, address _strategy, uint256 _depositFeeBP, uint256 _withdrawFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) Add(_isLPToken bool, _allocPoint *big.Int, _lpToken common.Address, _strategy common.Address, _depositFeeBP *big.Int, _withdrawFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Add(&_MasterChefDogsV2.TransactOpts, _isLPToken, _allocPoint, _lpToken, _strategy, _depositFeeBP, _withdrawFeeBP, _harvestInterval, _withUpdate)
}

// AddPoolUserData is a paid mutator transaction binding the contract method 0xc194dfab.
//
// Solidity: function addPoolUserData(uint256 _poolIndex, address[] _users, uint256[] _usersStakeData) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) AddPoolUserData(opts *bind.TransactOpts, _poolIndex *big.Int, _users []common.Address, _usersStakeData []*big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "addPoolUserData", _poolIndex, _users, _usersStakeData)
}

// AddPoolUserData is a paid mutator transaction binding the contract method 0xc194dfab.
//
// Solidity: function addPoolUserData(uint256 _poolIndex, address[] _users, uint256[] _usersStakeData) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) AddPoolUserData(_poolIndex *big.Int, _users []common.Address, _usersStakeData []*big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.AddPoolUserData(&_MasterChefDogsV2.TransactOpts, _poolIndex, _users, _usersStakeData)
}

// AddPoolUserData is a paid mutator transaction binding the contract method 0xc194dfab.
//
// Solidity: function addPoolUserData(uint256 _poolIndex, address[] _users, uint256[] _usersStakeData) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) AddPoolUserData(_poolIndex *big.Int, _users []common.Address, _usersStakeData []*big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.AddPoolUserData(&_MasterChefDogsV2.TransactOpts, _poolIndex, _users, _usersStakeData)
}

// BurnMint is a paid mutator transaction binding the contract method 0xc6fe62b0.
//
// Solidity: function burnMint() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) BurnMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "burnMint")
}

// BurnMint is a paid mutator transaction binding the contract method 0xc6fe62b0.
//
// Solidity: function burnMint() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) BurnMint() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.BurnMint(&_MasterChefDogsV2.TransactOpts)
}

// BurnMint is a paid mutator transaction binding the contract method 0xc6fe62b0.
//
// Solidity: function burnMint() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) BurnMint() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.BurnMint(&_MasterChefDogsV2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, address _referrer) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "deposit", _pid, _amount, _referrer)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, address _referrer) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) Deposit(_pid *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Deposit(&_MasterChefDogsV2.TransactOpts, _pid, _amount, _referrer)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, address _referrer) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Deposit(&_MasterChefDogsV2.TransactOpts, _pid, _amount, _referrer)
}

// DepositMigrator is a paid mutator transaction binding the contract method 0x21943779.
//
// Solidity: function depositMigrator(uint256 _pid, uint256 _amount, address _referrer) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) DepositMigrator(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "depositMigrator", _pid, _amount, _referrer)
}

// DepositMigrator is a paid mutator transaction binding the contract method 0x21943779.
//
// Solidity: function depositMigrator(uint256 _pid, uint256 _amount, address _referrer) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) DepositMigrator(_pid *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.DepositMigrator(&_MasterChefDogsV2.TransactOpts, _pid, _amount, _referrer)
}

// DepositMigrator is a paid mutator transaction binding the contract method 0x21943779.
//
// Solidity: function depositMigrator(uint256 _pid, uint256 _amount, address _referrer) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) DepositMigrator(_pid *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.DepositMigrator(&_MasterChefDogsV2.TransactOpts, _pid, _amount, _referrer)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.EmergencyWithdraw(&_MasterChefDogsV2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.EmergencyWithdraw(&_MasterChefDogsV2.TransactOpts, _pid)
}

// IncreaseDogsSupply is a paid mutator transaction binding the contract method 0x1b4c495d.
//
// Solidity: function increaseDogsSupply(uint256 _amount) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) IncreaseDogsSupply(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "increaseDogsSupply", _amount)
}

// IncreaseDogsSupply is a paid mutator transaction binding the contract method 0x1b4c495d.
//
// Solidity: function increaseDogsSupply(uint256 _amount) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) IncreaseDogsSupply(_amount *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.IncreaseDogsSupply(&_MasterChefDogsV2.TransactOpts, _amount)
}

// IncreaseDogsSupply is a paid mutator transaction binding the contract method 0x1b4c495d.
//
// Solidity: function increaseDogsSupply(uint256 _amount) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) IncreaseDogsSupply(_amount *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.IncreaseDogsSupply(&_MasterChefDogsV2.TransactOpts, _amount)
}

// LockPlatform is a paid mutator transaction binding the contract method 0x35d70427.
//
// Solidity: function lockPlatform() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) LockPlatform(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "lockPlatform")
}

// LockPlatform is a paid mutator transaction binding the contract method 0x35d70427.
//
// Solidity: function lockPlatform() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) LockPlatform() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.LockPlatform(&_MasterChefDogsV2.TransactOpts)
}

// LockPlatform is a paid mutator transaction binding the contract method 0x35d70427.
//
// Solidity: function lockPlatform() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) LockPlatform() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.LockPlatform(&_MasterChefDogsV2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.MassUpdatePools(&_MasterChefDogsV2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.MassUpdatePools(&_MasterChefDogsV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.RenounceOwnership(&_MasterChefDogsV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.RenounceOwnership(&_MasterChefDogsV2.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x77200848.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, uint256 _depositFeeBP, uint256 _withdrawFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _depositFeeBP *big.Int, _withdrawFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "set", _pid, _allocPoint, _depositFeeBP, _withdrawFeeBP, _harvestInterval, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x77200848.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, uint256 _depositFeeBP, uint256 _withdrawFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) Set(_pid *big.Int, _allocPoint *big.Int, _depositFeeBP *big.Int, _withdrawFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Set(&_MasterChefDogsV2.TransactOpts, _pid, _allocPoint, _depositFeeBP, _withdrawFeeBP, _harvestInterval, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x77200848.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, uint256 _depositFeeBP, uint256 _withdrawFeeBP, uint256 _harvestInterval, bool _withUpdate) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _depositFeeBP *big.Int, _withdrawFeeBP *big.Int, _harvestInterval *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Set(&_MasterChefDogsV2.TransactOpts, _pid, _allocPoint, _depositFeeBP, _withdrawFeeBP, _harvestInterval, _withUpdate)
}

// SetDDSCAAddress is a paid mutator transaction binding the contract method 0xdeaa63d8.
//
// Solidity: function setDDSCAAddress(address _ddsca) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) SetDDSCAAddress(opts *bind.TransactOpts, _ddsca common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "setDDSCAAddress", _ddsca)
}

// SetDDSCAAddress is a paid mutator transaction binding the contract method 0xdeaa63d8.
//
// Solidity: function setDDSCAAddress(address _ddsca) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) SetDDSCAAddress(_ddsca common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetDDSCAAddress(&_MasterChefDogsV2.TransactOpts, _ddsca)
}

// SetDDSCAAddress is a paid mutator transaction binding the contract method 0xdeaa63d8.
//
// Solidity: function setDDSCAAddress(address _ddsca) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) SetDDSCAAddress(_ddsca common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetDDSCAAddress(&_MasterChefDogsV2.TransactOpts, _ddsca)
}

// SetFarmStartBlock is a paid mutator transaction binding the contract method 0xf678fc18.
//
// Solidity: function setFarmStartBlock(uint256 _newStartBlock) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) SetFarmStartBlock(opts *bind.TransactOpts, _newStartBlock *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "setFarmStartBlock", _newStartBlock)
}

// SetFarmStartBlock is a paid mutator transaction binding the contract method 0xf678fc18.
//
// Solidity: function setFarmStartBlock(uint256 _newStartBlock) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) SetFarmStartBlock(_newStartBlock *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetFarmStartBlock(&_MasterChefDogsV2.TransactOpts, _newStartBlock)
}

// SetFarmStartBlock is a paid mutator transaction binding the contract method 0xf678fc18.
//
// Solidity: function setFarmStartBlock(uint256 _newStartBlock) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) SetFarmStartBlock(_newStartBlock *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetFarmStartBlock(&_MasterChefDogsV2.TransactOpts, _newStartBlock)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _govAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) SetGov(opts *bind.TransactOpts, _govAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "setGov", _govAddress)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _govAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) SetGov(_govAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetGov(&_MasterChefDogsV2.TransactOpts, _govAddress)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _govAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) SetGov(_govAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetGov(&_MasterChefDogsV2.TransactOpts, _govAddress)
}

// SetPlatformAddress is a paid mutator transaction binding the contract method 0xcc03c342.
//
// Solidity: function setPlatformAddress(address _platformAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) SetPlatformAddress(opts *bind.TransactOpts, _platformAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "setPlatformAddress", _platformAddress)
}

// SetPlatformAddress is a paid mutator transaction binding the contract method 0xcc03c342.
//
// Solidity: function setPlatformAddress(address _platformAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) SetPlatformAddress(_platformAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetPlatformAddress(&_MasterChefDogsV2.TransactOpts, _platformAddress)
}

// SetPlatformAddress is a paid mutator transaction binding the contract method 0xcc03c342.
//
// Solidity: function setPlatformAddress(address _platformAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) SetPlatformAddress(_platformAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetPlatformAddress(&_MasterChefDogsV2.TransactOpts, _platformAddress)
}

// SetReferral is a paid mutator transaction binding the contract method 0x9e5914da.
//
// Solidity: function setReferral(address _dogsReferral) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) SetReferral(opts *bind.TransactOpts, _dogsReferral common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "setReferral", _dogsReferral)
}

// SetReferral is a paid mutator transaction binding the contract method 0x9e5914da.
//
// Solidity: function setReferral(address _dogsReferral) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) SetReferral(_dogsReferral common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetReferral(&_MasterChefDogsV2.TransactOpts, _dogsReferral)
}

// SetReferral is a paid mutator transaction binding the contract method 0x9e5914da.
//
// Solidity: function setReferral(address _dogsReferral) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) SetReferral(_dogsReferral common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.SetReferral(&_MasterChefDogsV2.TransactOpts, _dogsReferral)
}

// ToggleMigrationEnabled is a paid mutator transaction binding the contract method 0xc151e28b.
//
// Solidity: function toggleMigrationEnabled(bool _state) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) ToggleMigrationEnabled(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "toggleMigrationEnabled", _state)
}

// ToggleMigrationEnabled is a paid mutator transaction binding the contract method 0xc151e28b.
//
// Solidity: function toggleMigrationEnabled(bool _state) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) ToggleMigrationEnabled(_state bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.ToggleMigrationEnabled(&_MasterChefDogsV2.TransactOpts, _state)
}

// ToggleMigrationEnabled is a paid mutator transaction binding the contract method 0xc151e28b.
//
// Solidity: function toggleMigrationEnabled(bool _state) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) ToggleMigrationEnabled(_state bool) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.ToggleMigrationEnabled(&_MasterChefDogsV2.TransactOpts, _state)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.TransferOwnership(&_MasterChefDogsV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.TransferOwnership(&_MasterChefDogsV2.TransactOpts, newOwner)
}

// UpdateEmissions is a paid mutator transaction binding the contract method 0xe39f9481.
//
// Solidity: function updateEmissions(uint256 priceInCents) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) UpdateEmissions(opts *bind.TransactOpts, priceInCents *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "updateEmissions", priceInCents)
}

// UpdateEmissions is a paid mutator transaction binding the contract method 0xe39f9481.
//
// Solidity: function updateEmissions(uint256 priceInCents) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) UpdateEmissions(priceInCents *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.UpdateEmissions(&_MasterChefDogsV2.TransactOpts, priceInCents)
}

// UpdateEmissions is a paid mutator transaction binding the contract method 0xe39f9481.
//
// Solidity: function updateEmissions(uint256 priceInCents) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) UpdateEmissions(priceInCents *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.UpdateEmissions(&_MasterChefDogsV2.TransactOpts, priceInCents)
}

// UpdateFeeManager is a paid mutator transaction binding the contract method 0xc6acd274.
//
// Solidity: function updateFeeManager(address _feeManagerAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) UpdateFeeManager(opts *bind.TransactOpts, _feeManagerAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "updateFeeManager", _feeManagerAddress)
}

// UpdateFeeManager is a paid mutator transaction binding the contract method 0xc6acd274.
//
// Solidity: function updateFeeManager(address _feeManagerAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) UpdateFeeManager(_feeManagerAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.UpdateFeeManager(&_MasterChefDogsV2.TransactOpts, _feeManagerAddress)
}

// UpdateFeeManager is a paid mutator transaction binding the contract method 0xc6acd274.
//
// Solidity: function updateFeeManager(address _feeManagerAddress) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) UpdateFeeManager(_feeManagerAddress common.Address) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.UpdateFeeManager(&_MasterChefDogsV2.TransactOpts, _feeManagerAddress)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.UpdatePool(&_MasterChefDogsV2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.UpdatePool(&_MasterChefDogsV2.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Withdraw(&_MasterChefDogsV2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefDogsV2 *MasterChefDogsV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefDogsV2.Contract.Withdraw(&_MasterChefDogsV2.TransactOpts, _pid, _amount)
}

// MasterChefDogsV2AddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2AddPoolIterator struct {
	Event *MasterChefDogsV2AddPool // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2AddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2AddPool)
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
		it.Event = new(MasterChefDogsV2AddPool)
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
func (it *MasterChefDogsV2AddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2AddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2AddPool represents a AddPool event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2AddPool struct {
	Pid             *big.Int
	IsLPToken       bool
	AllocPoint      *big.Int
	LpToken         common.Address
	DepositFeeBP    *big.Int
	WithdrawFeeBP   *big.Int
	HarvestInterval *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x699e37f4ba4eef3b223583fe8b9f4d74b199f24e5bfba8237739307dfd445d56.
//
// Solidity: event AddPool(uint256 indexed pid, bool isLPToken, uint256 allocPoint, address lpToken, uint256 depositFeeBP, uint256 _withdrawFeeBP, uint256 harvestInterval)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterAddPool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefDogsV2AddPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "AddPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2AddPoolIterator{contract: _MasterChefDogsV2.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x699e37f4ba4eef3b223583fe8b9f4d74b199f24e5bfba8237739307dfd445d56.
//
// Solidity: event AddPool(uint256 indexed pid, bool isLPToken, uint256 allocPoint, address lpToken, uint256 depositFeeBP, uint256 _withdrawFeeBP, uint256 harvestInterval)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2AddPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "AddPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2AddPool)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0x699e37f4ba4eef3b223583fe8b9f4d74b199f24e5bfba8237739307dfd445d56.
//
// Solidity: event AddPool(uint256 indexed pid, bool isLPToken, uint256 allocPoint, address lpToken, uint256 depositFeeBP, uint256 _withdrawFeeBP, uint256 harvestInterval)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseAddPool(log types.Log) (*MasterChefDogsV2AddPool, error) {
	event := new(MasterChefDogsV2AddPool)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2DepositIterator struct {
	Event *MasterChefDogsV2Deposit // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2Deposit)
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
		it.Event = new(MasterChefDogsV2Deposit)
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
func (it *MasterChefDogsV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2Deposit represents a Deposit event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefDogsV2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2DepositIterator{contract: _MasterChefDogsV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2Deposit)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseDeposit(log types.Log) (*MasterChefDogsV2Deposit, error) {
	event := new(MasterChefDogsV2Deposit)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2EmergencyWithdrawIterator struct {
	Event *MasterChefDogsV2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2EmergencyWithdraw)
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
		it.Event = new(MasterChefDogsV2EmergencyWithdraw)
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
func (it *MasterChefDogsV2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2EmergencyWithdraw represents a EmergencyWithdraw event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefDogsV2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2EmergencyWithdrawIterator{contract: _MasterChefDogsV2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2EmergencyWithdraw)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseEmergencyWithdraw(log types.Log) (*MasterChefDogsV2EmergencyWithdraw, error) {
	event := new(MasterChefDogsV2EmergencyWithdraw)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2GovUpdatedIterator is returned from FilterGovUpdated and is used to iterate over the raw logs and unpacked data for GovUpdated events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2GovUpdatedIterator struct {
	Event *MasterChefDogsV2GovUpdated // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2GovUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2GovUpdated)
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
		it.Event = new(MasterChefDogsV2GovUpdated)
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
func (it *MasterChefDogsV2GovUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2GovUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2GovUpdated represents a GovUpdated event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2GovUpdated struct {
	GovAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGovUpdated is a free log retrieval operation binding the contract event 0xb49a738bf7c18189b8cd3a6cd9d2b44045ca8070a4fa7e9db46ecfcbedee6bd9.
//
// Solidity: event GovUpdated(address govAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterGovUpdated(opts *bind.FilterOpts) (*MasterChefDogsV2GovUpdatedIterator, error) {

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "GovUpdated")
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2GovUpdatedIterator{contract: _MasterChefDogsV2.contract, event: "GovUpdated", logs: logs, sub: sub}, nil
}

// WatchGovUpdated is a free log subscription operation binding the contract event 0xb49a738bf7c18189b8cd3a6cd9d2b44045ca8070a4fa7e9db46ecfcbedee6bd9.
//
// Solidity: event GovUpdated(address govAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchGovUpdated(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2GovUpdated) (event.Subscription, error) {

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "GovUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2GovUpdated)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "GovUpdated", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseGovUpdated(log types.Log) (*MasterChefDogsV2GovUpdated, error) {
	event := new(MasterChefDogsV2GovUpdated)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "GovUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2OwnershipTransferredIterator struct {
	Event *MasterChefDogsV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2OwnershipTransferred)
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
		it.Event = new(MasterChefDogsV2OwnershipTransferred)
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
func (it *MasterChefDogsV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2OwnershipTransferred represents a OwnershipTransferred event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterChefDogsV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2OwnershipTransferredIterator{contract: _MasterChefDogsV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2OwnershipTransferred)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseOwnershipTransferred(log types.Log) (*MasterChefDogsV2OwnershipTransferred, error) {
	event := new(MasterChefDogsV2OwnershipTransferred)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2RewardLockedUpIterator is returned from FilterRewardLockedUp and is used to iterate over the raw logs and unpacked data for RewardLockedUp events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2RewardLockedUpIterator struct {
	Event *MasterChefDogsV2RewardLockedUp // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2RewardLockedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2RewardLockedUp)
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
		it.Event = new(MasterChefDogsV2RewardLockedUp)
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
func (it *MasterChefDogsV2RewardLockedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2RewardLockedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2RewardLockedUp represents a RewardLockedUp event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2RewardLockedUp struct {
	User           common.Address
	Pid            *big.Int
	AmountLockedUp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardLockedUp is a free log retrieval operation binding the contract event 0xee470483107f579a55c754fa00613c45a9a3b617a418b39cb0be97e5381ba7c1.
//
// Solidity: event RewardLockedUp(address indexed user, uint256 indexed pid, uint256 amountLockedUp)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterRewardLockedUp(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefDogsV2RewardLockedUpIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "RewardLockedUp", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2RewardLockedUpIterator{contract: _MasterChefDogsV2.contract, event: "RewardLockedUp", logs: logs, sub: sub}, nil
}

// WatchRewardLockedUp is a free log subscription operation binding the contract event 0xee470483107f579a55c754fa00613c45a9a3b617a418b39cb0be97e5381ba7c1.
//
// Solidity: event RewardLockedUp(address indexed user, uint256 indexed pid, uint256 amountLockedUp)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchRewardLockedUp(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2RewardLockedUp, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "RewardLockedUp", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2RewardLockedUp)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "RewardLockedUp", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseRewardLockedUp(log types.Log) (*MasterChefDogsV2RewardLockedUp, error) {
	event := new(MasterChefDogsV2RewardLockedUp)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "RewardLockedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2SetDogsReferralIterator is returned from FilterSetDogsReferral and is used to iterate over the raw logs and unpacked data for SetDogsReferral events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2SetDogsReferralIterator struct {
	Event *MasterChefDogsV2SetDogsReferral // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2SetDogsReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2SetDogsReferral)
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
		it.Event = new(MasterChefDogsV2SetDogsReferral)
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
func (it *MasterChefDogsV2SetDogsReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2SetDogsReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2SetDogsReferral represents a SetDogsReferral event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2SetDogsReferral struct {
	DogsAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetDogsReferral is a free log retrieval operation binding the contract event 0xaac494b245dbe2139de7414ae68e68155297a6ceeba1f74fb1a515da18634342.
//
// Solidity: event SetDogsReferral(address dogsAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterSetDogsReferral(opts *bind.FilterOpts) (*MasterChefDogsV2SetDogsReferralIterator, error) {

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "SetDogsReferral")
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2SetDogsReferralIterator{contract: _MasterChefDogsV2.contract, event: "SetDogsReferral", logs: logs, sub: sub}, nil
}

// WatchSetDogsReferral is a free log subscription operation binding the contract event 0xaac494b245dbe2139de7414ae68e68155297a6ceeba1f74fb1a515da18634342.
//
// Solidity: event SetDogsReferral(address dogsAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchSetDogsReferral(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2SetDogsReferral) (event.Subscription, error) {

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "SetDogsReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2SetDogsReferral)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "SetDogsReferral", log); err != nil {
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

// ParseSetDogsReferral is a log parse operation binding the contract event 0xaac494b245dbe2139de7414ae68e68155297a6ceeba1f74fb1a515da18634342.
//
// Solidity: event SetDogsReferral(address dogsAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseSetDogsReferral(log types.Log) (*MasterChefDogsV2SetDogsReferral, error) {
	event := new(MasterChefDogsV2SetDogsReferral)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "SetDogsReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2SetPlatformAddressIterator is returned from FilterSetPlatformAddress and is used to iterate over the raw logs and unpacked data for SetPlatformAddress events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2SetPlatformAddressIterator struct {
	Event *MasterChefDogsV2SetPlatformAddress // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2SetPlatformAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2SetPlatformAddress)
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
		it.Event = new(MasterChefDogsV2SetPlatformAddress)
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
func (it *MasterChefDogsV2SetPlatformAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2SetPlatformAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2SetPlatformAddress represents a SetPlatformAddress event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2SetPlatformAddress struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPlatformAddress is a free log retrieval operation binding the contract event 0xd2991d7f16266fc3594a7a085f0da1060fde4dba356999af15ced38362eee36c.
//
// Solidity: event SetPlatformAddress(address indexed newAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterSetPlatformAddress(opts *bind.FilterOpts, newAddress []common.Address) (*MasterChefDogsV2SetPlatformAddressIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "SetPlatformAddress", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2SetPlatformAddressIterator{contract: _MasterChefDogsV2.contract, event: "SetPlatformAddress", logs: logs, sub: sub}, nil
}

// WatchSetPlatformAddress is a free log subscription operation binding the contract event 0xd2991d7f16266fc3594a7a085f0da1060fde4dba356999af15ced38362eee36c.
//
// Solidity: event SetPlatformAddress(address indexed newAddress)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchSetPlatformAddress(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2SetPlatformAddress, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "SetPlatformAddress", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2SetPlatformAddress)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "SetPlatformAddress", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseSetPlatformAddress(log types.Log) (*MasterChefDogsV2SetPlatformAddress, error) {
	event := new(MasterChefDogsV2SetPlatformAddress)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "SetPlatformAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2SetPoolIterator is returned from FilterSetPool and is used to iterate over the raw logs and unpacked data for SetPool events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2SetPoolIterator struct {
	Event *MasterChefDogsV2SetPool // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2SetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2SetPool)
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
		it.Event = new(MasterChefDogsV2SetPool)
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
func (it *MasterChefDogsV2SetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2SetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2SetPool represents a SetPool event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2SetPool struct {
	Pid             *big.Int
	AllocPoint      *big.Int
	DepositFeeBP    *big.Int
	WithdrawFeeBP   *big.Int
	HarvestInterval *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetPool is a free log retrieval operation binding the contract event 0xcd68802350e451beb059b0e629041bf0bd1328eae134e107b7c1be7f4ebde195.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint, uint256 depositFeeBP, uint256 _withdrawFeeBP, uint256 harvestInterval)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterSetPool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefDogsV2SetPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2SetPoolIterator{contract: _MasterChefDogsV2.contract, event: "SetPool", logs: logs, sub: sub}, nil
}

// WatchSetPool is a free log subscription operation binding the contract event 0xcd68802350e451beb059b0e629041bf0bd1328eae134e107b7c1be7f4ebde195.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint, uint256 depositFeeBP, uint256 _withdrawFeeBP, uint256 harvestInterval)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchSetPool(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2SetPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2SetPool)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "SetPool", log); err != nil {
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

// ParseSetPool is a log parse operation binding the contract event 0xcd68802350e451beb059b0e629041bf0bd1328eae134e107b7c1be7f4ebde195.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint, uint256 depositFeeBP, uint256 _withdrawFeeBP, uint256 harvestInterval)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseSetPool(log types.Log) (*MasterChefDogsV2SetPool, error) {
	event := new(MasterChefDogsV2SetPool)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "SetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefDogsV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2WithdrawIterator struct {
	Event *MasterChefDogsV2Withdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefDogsV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefDogsV2Withdraw)
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
		it.Event = new(MasterChefDogsV2Withdraw)
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
func (it *MasterChefDogsV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefDogsV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefDogsV2Withdraw represents a Withdraw event raised by the MasterChefDogsV2 contract.
type MasterChefDogsV2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefDogsV2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefDogsV2WithdrawIterator{contract: _MasterChefDogsV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefDogsV2Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefDogsV2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefDogsV2Withdraw)
				if err := _MasterChefDogsV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_MasterChefDogsV2 *MasterChefDogsV2Filterer) ParseWithdraw(log types.Log) (*MasterChefDogsV2Withdraw, error) {
	event := new(MasterChefDogsV2Withdraw)
	if err := _MasterChefDogsV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
