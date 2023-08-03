// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unicrypt

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

// TokenVestingLockParams is an auto generated low-level Go binding around an user-defined struct.
type TokenVestingLockParams struct {
	Owner         common.Address
	Amount        *big.Int
	StartEmission *big.Int
	EndEmission   *big.Int
	Condition     common.Address
}

// UnicryptTokenVestingMetaData contains all meta data concerning the UnicryptTokenVesting contract.
var UnicryptTokenVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIUnicryptAdmin\",\"name\":\"_uncxAdmins\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountInTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startEmission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endEmission\",\"type\":\"uint256\"}],\"name\":\"onLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountInTokens\",\"type\":\"uint256\"}],\"name\":\"onMigrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"name\":\"onRelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromLockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toLockID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountInTokens\",\"type\":\"uint256\"}],\"name\":\"onSplitLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockIDFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockIDto\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"onTransferLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountInTokens\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLACKLIST\",\"outputs\":[{\"internalType\":\"contractITokenBlacklist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeLockingFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"freeLockingToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LOCKS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sharesDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEmission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endEmission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"condition\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIGRATOR\",\"outputs\":[{\"internalType\":\"contractIMigrator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NONCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SHARES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_add\",\"type\":\"bool\"}],\"name\":\"adminSetWhitelister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertSharesToTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"convertTokensToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_add\",\"type\":\"bool\"}],\"name\":\"editZeroFeeWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"}],\"name\":\"getLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumLockedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTokenAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTokenLockIDAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTokenLocksLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTokenWhitelisterAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenWhitelisterLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getTokenWhitelisterStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserLockIDForTokenAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserLockedTokenAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLockedTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getUserLocksForTokenLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"}],\"name\":\"getWithdrawableShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"}],\"name\":\"getWithdrawableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getZeroFeeTokenAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZeroFeeTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEmission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endEmission\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"condition\",\"type\":\"address\"}],\"internalType\":\"structTokenVesting.LockParams[]\",\"name\":\"_lock_params\",\"type\":\"tuple[]\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_option\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"payForFreeTokenLocks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_date\",\"type\":\"uint256\"}],\"name\":\"relock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"}],\"name\":\"revokeCondition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITokenBlacklist\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setBlacklistContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_freeLockingFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_freeLockingToken\",\"type\":\"address\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigrator\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"condition\",\"type\":\"address\"}],\"name\":\"testCondition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"tokenOnZeroFeeWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferLockOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnicryptTokenVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use UnicryptTokenVestingMetaData.ABI instead.
var UnicryptTokenVestingABI = UnicryptTokenVestingMetaData.ABI

// UnicryptTokenVesting is an auto generated Go binding around an Ethereum contract.
type UnicryptTokenVesting struct {
	UnicryptTokenVestingCaller     // Read-only binding to the contract
	UnicryptTokenVestingTransactor // Write-only binding to the contract
	UnicryptTokenVestingFilterer   // Log filterer for contract events
}

// UnicryptTokenVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnicryptTokenVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicryptTokenVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnicryptTokenVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicryptTokenVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnicryptTokenVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicryptTokenVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnicryptTokenVestingSession struct {
	Contract     *UnicryptTokenVesting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnicryptTokenVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnicryptTokenVestingCallerSession struct {
	Contract *UnicryptTokenVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// UnicryptTokenVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnicryptTokenVestingTransactorSession struct {
	Contract     *UnicryptTokenVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// UnicryptTokenVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnicryptTokenVestingRaw struct {
	Contract *UnicryptTokenVesting // Generic contract binding to access the raw methods on
}

// UnicryptTokenVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnicryptTokenVestingCallerRaw struct {
	Contract *UnicryptTokenVestingCaller // Generic read-only contract binding to access the raw methods on
}

// UnicryptTokenVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnicryptTokenVestingTransactorRaw struct {
	Contract *UnicryptTokenVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnicryptTokenVesting creates a new instance of UnicryptTokenVesting, bound to a specific deployed contract.
func NewUnicryptTokenVesting(address common.Address, backend bind.ContractBackend) (*UnicryptTokenVesting, error) {
	contract, err := bindUnicryptTokenVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVesting{UnicryptTokenVestingCaller: UnicryptTokenVestingCaller{contract: contract}, UnicryptTokenVestingTransactor: UnicryptTokenVestingTransactor{contract: contract}, UnicryptTokenVestingFilterer: UnicryptTokenVestingFilterer{contract: contract}}, nil
}

// NewUnicryptTokenVestingCaller creates a new read-only instance of UnicryptTokenVesting, bound to a specific deployed contract.
func NewUnicryptTokenVestingCaller(address common.Address, caller bind.ContractCaller) (*UnicryptTokenVestingCaller, error) {
	contract, err := bindUnicryptTokenVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingCaller{contract: contract}, nil
}

// NewUnicryptTokenVestingTransactor creates a new write-only instance of UnicryptTokenVesting, bound to a specific deployed contract.
func NewUnicryptTokenVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*UnicryptTokenVestingTransactor, error) {
	contract, err := bindUnicryptTokenVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingTransactor{contract: contract}, nil
}

// NewUnicryptTokenVestingFilterer creates a new log filterer instance of UnicryptTokenVesting, bound to a specific deployed contract.
func NewUnicryptTokenVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*UnicryptTokenVestingFilterer, error) {
	contract, err := bindUnicryptTokenVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingFilterer{contract: contract}, nil
}

// bindUnicryptTokenVesting binds a generic wrapper to an already deployed contract.
func bindUnicryptTokenVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnicryptTokenVestingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnicryptTokenVesting *UnicryptTokenVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnicryptTokenVesting.Contract.UnicryptTokenVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnicryptTokenVesting *UnicryptTokenVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.UnicryptTokenVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnicryptTokenVesting *UnicryptTokenVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.UnicryptTokenVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnicryptTokenVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.contract.Transact(opts, method, params...)
}

// BLACKLIST is a free data retrieval call binding the contract method 0xc8b0cf68.
//
// Solidity: function BLACKLIST() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) BLACKLIST(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "BLACKLIST")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BLACKLIST is a free data retrieval call binding the contract method 0xc8b0cf68.
//
// Solidity: function BLACKLIST() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) BLACKLIST() (common.Address, error) {
	return _UnicryptTokenVesting.Contract.BLACKLIST(&_UnicryptTokenVesting.CallOpts)
}

// BLACKLIST is a free data retrieval call binding the contract method 0xc8b0cf68.
//
// Solidity: function BLACKLIST() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) BLACKLIST() (common.Address, error) {
	return _UnicryptTokenVesting.Contract.BLACKLIST(&_UnicryptTokenVesting.CallOpts)
}

// FEES is a free data retrieval call binding the contract method 0x8b7b23ee.
//
// Solidity: function FEES() view returns(uint256 tokenFee, uint256 freeLockingFee, address feeAddress, address freeLockingToken)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) FEES(opts *bind.CallOpts) (struct {
	TokenFee         *big.Int
	FreeLockingFee   *big.Int
	FeeAddress       common.Address
	FreeLockingToken common.Address
}, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "FEES")

	outstruct := new(struct {
		TokenFee         *big.Int
		FreeLockingFee   *big.Int
		FeeAddress       common.Address
		FreeLockingToken common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FreeLockingFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.FreeLockingToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FEES is a free data retrieval call binding the contract method 0x8b7b23ee.
//
// Solidity: function FEES() view returns(uint256 tokenFee, uint256 freeLockingFee, address feeAddress, address freeLockingToken)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) FEES() (struct {
	TokenFee         *big.Int
	FreeLockingFee   *big.Int
	FeeAddress       common.Address
	FreeLockingToken common.Address
}, error) {
	return _UnicryptTokenVesting.Contract.FEES(&_UnicryptTokenVesting.CallOpts)
}

// FEES is a free data retrieval call binding the contract method 0x8b7b23ee.
//
// Solidity: function FEES() view returns(uint256 tokenFee, uint256 freeLockingFee, address feeAddress, address freeLockingToken)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) FEES() (struct {
	TokenFee         *big.Int
	FreeLockingFee   *big.Int
	FeeAddress       common.Address
	FreeLockingToken common.Address
}, error) {
	return _UnicryptTokenVesting.Contract.FEES(&_UnicryptTokenVesting.CallOpts)
}

// LOCKS is a free data retrieval call binding the contract method 0xcf0d5af3.
//
// Solidity: function LOCKS(uint256 ) view returns(address tokenAddress, uint256 sharesDeposited, uint256 sharesWithdrawn, uint256 startEmission, uint256 endEmission, uint256 lockID, address owner, address condition)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) LOCKS(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenAddress    common.Address
	SharesDeposited *big.Int
	SharesWithdrawn *big.Int
	StartEmission   *big.Int
	EndEmission     *big.Int
	LockID          *big.Int
	Owner           common.Address
	Condition       common.Address
}, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "LOCKS", arg0)

	outstruct := new(struct {
		TokenAddress    common.Address
		SharesDeposited *big.Int
		SharesWithdrawn *big.Int
		StartEmission   *big.Int
		EndEmission     *big.Int
		LockID          *big.Int
		Owner           common.Address
		Condition       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SharesDeposited = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SharesWithdrawn = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartEmission = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndEmission = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LockID = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Condition = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LOCKS is a free data retrieval call binding the contract method 0xcf0d5af3.
//
// Solidity: function LOCKS(uint256 ) view returns(address tokenAddress, uint256 sharesDeposited, uint256 sharesWithdrawn, uint256 startEmission, uint256 endEmission, uint256 lockID, address owner, address condition)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) LOCKS(arg0 *big.Int) (struct {
	TokenAddress    common.Address
	SharesDeposited *big.Int
	SharesWithdrawn *big.Int
	StartEmission   *big.Int
	EndEmission     *big.Int
	LockID          *big.Int
	Owner           common.Address
	Condition       common.Address
}, error) {
	return _UnicryptTokenVesting.Contract.LOCKS(&_UnicryptTokenVesting.CallOpts, arg0)
}

// LOCKS is a free data retrieval call binding the contract method 0xcf0d5af3.
//
// Solidity: function LOCKS(uint256 ) view returns(address tokenAddress, uint256 sharesDeposited, uint256 sharesWithdrawn, uint256 startEmission, uint256 endEmission, uint256 lockID, address owner, address condition)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) LOCKS(arg0 *big.Int) (struct {
	TokenAddress    common.Address
	SharesDeposited *big.Int
	SharesWithdrawn *big.Int
	StartEmission   *big.Int
	EndEmission     *big.Int
	LockID          *big.Int
	Owner           common.Address
	Condition       common.Address
}, error) {
	return _UnicryptTokenVesting.Contract.LOCKS(&_UnicryptTokenVesting.CallOpts, arg0)
}

// MIGRATOR is a free data retrieval call binding the contract method 0x9ecd7472.
//
// Solidity: function MIGRATOR() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) MIGRATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "MIGRATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MIGRATOR is a free data retrieval call binding the contract method 0x9ecd7472.
//
// Solidity: function MIGRATOR() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) MIGRATOR() (common.Address, error) {
	return _UnicryptTokenVesting.Contract.MIGRATOR(&_UnicryptTokenVesting.CallOpts)
}

// MIGRATOR is a free data retrieval call binding the contract method 0x9ecd7472.
//
// Solidity: function MIGRATOR() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) MIGRATOR() (common.Address, error) {
	return _UnicryptTokenVesting.Contract.MIGRATOR(&_UnicryptTokenVesting.CallOpts)
}

// MINIMUMDEPOSIT is a free data retrieval call binding the contract method 0xf19451d8.
//
// Solidity: function MINIMUM_DEPOSIT() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) MINIMUMDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "MINIMUM_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMDEPOSIT is a free data retrieval call binding the contract method 0xf19451d8.
//
// Solidity: function MINIMUM_DEPOSIT() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) MINIMUMDEPOSIT() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.MINIMUMDEPOSIT(&_UnicryptTokenVesting.CallOpts)
}

// MINIMUMDEPOSIT is a free data retrieval call binding the contract method 0xf19451d8.
//
// Solidity: function MINIMUM_DEPOSIT() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) MINIMUMDEPOSIT() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.MINIMUMDEPOSIT(&_UnicryptTokenVesting.CallOpts)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) NONCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "NONCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) NONCE() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.NONCE(&_UnicryptTokenVesting.CallOpts)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) NONCE() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.NONCE(&_UnicryptTokenVesting.CallOpts)
}

// SHARES is a free data retrieval call binding the contract method 0xc368803a.
//
// Solidity: function SHARES(address ) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) SHARES(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "SHARES", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHARES is a free data retrieval call binding the contract method 0xc368803a.
//
// Solidity: function SHARES(address ) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) SHARES(arg0 common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.SHARES(&_UnicryptTokenVesting.CallOpts, arg0)
}

// SHARES is a free data retrieval call binding the contract method 0xc368803a.
//
// Solidity: function SHARES(address ) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) SHARES(arg0 common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.SHARES(&_UnicryptTokenVesting.CallOpts, arg0)
}

// ConvertSharesToTokens is a free data retrieval call binding the contract method 0x0cdebc9e.
//
// Solidity: function convertSharesToTokens(address _token, uint256 _shares) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) ConvertSharesToTokens(opts *bind.CallOpts, _token common.Address, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "convertSharesToTokens", _token, _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSharesToTokens is a free data retrieval call binding the contract method 0x0cdebc9e.
//
// Solidity: function convertSharesToTokens(address _token, uint256 _shares) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) ConvertSharesToTokens(_token common.Address, _shares *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.ConvertSharesToTokens(&_UnicryptTokenVesting.CallOpts, _token, _shares)
}

// ConvertSharesToTokens is a free data retrieval call binding the contract method 0x0cdebc9e.
//
// Solidity: function convertSharesToTokens(address _token, uint256 _shares) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) ConvertSharesToTokens(_token common.Address, _shares *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.ConvertSharesToTokens(&_UnicryptTokenVesting.CallOpts, _token, _shares)
}

// ConvertTokensToShares is a free data retrieval call binding the contract method 0xca5cc0c2.
//
// Solidity: function convertTokensToShares(address _token, uint256 _tokens) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) ConvertTokensToShares(opts *bind.CallOpts, _token common.Address, _tokens *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "convertTokensToShares", _token, _tokens)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertTokensToShares is a free data retrieval call binding the contract method 0xca5cc0c2.
//
// Solidity: function convertTokensToShares(address _token, uint256 _tokens) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) ConvertTokensToShares(_token common.Address, _tokens *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.ConvertTokensToShares(&_UnicryptTokenVesting.CallOpts, _token, _tokens)
}

// ConvertTokensToShares is a free data retrieval call binding the contract method 0xca5cc0c2.
//
// Solidity: function convertTokensToShares(address _token, uint256 _tokens) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) ConvertTokensToShares(_token common.Address, _tokens *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.ConvertTokensToShares(&_UnicryptTokenVesting.CallOpts, _token, _tokens)
}

// GetLock is a free data retrieval call binding the contract method 0xd68f4dd1.
//
// Solidity: function getLock(uint256 _lockID) view returns(uint256, address, uint256, uint256, uint256, uint256, uint256, uint256, address, address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetLock(opts *bind.CallOpts, _lockID *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getLock", _lockID)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	out9 := *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, err

}

// GetLock is a free data retrieval call binding the contract method 0xd68f4dd1.
//
// Solidity: function getLock(uint256 _lockID) view returns(uint256, address, uint256, uint256, uint256, uint256, uint256, uint256, address, address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetLock(_lockID *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetLock(&_UnicryptTokenVesting.CallOpts, _lockID)
}

// GetLock is a free data retrieval call binding the contract method 0xd68f4dd1.
//
// Solidity: function getLock(uint256 _lockID) view returns(uint256, address, uint256, uint256, uint256, uint256, uint256, uint256, address, address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetLock(_lockID *big.Int) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetLock(&_UnicryptTokenVesting.CallOpts, _lockID)
}

// GetNumLockedTokens is a free data retrieval call binding the contract method 0x783451e8.
//
// Solidity: function getNumLockedTokens() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetNumLockedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getNumLockedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumLockedTokens is a free data retrieval call binding the contract method 0x783451e8.
//
// Solidity: function getNumLockedTokens() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetNumLockedTokens() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetNumLockedTokens(&_UnicryptTokenVesting.CallOpts)
}

// GetNumLockedTokens is a free data retrieval call binding the contract method 0x783451e8.
//
// Solidity: function getNumLockedTokens() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetNumLockedTokens() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetNumLockedTokens(&_UnicryptTokenVesting.CallOpts)
}

// GetTokenAtIndex is a free data retrieval call binding the contract method 0x97988dce.
//
// Solidity: function getTokenAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetTokenAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getTokenAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAtIndex is a free data retrieval call binding the contract method 0x97988dce.
//
// Solidity: function getTokenAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetTokenAtIndex(_index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _index)
}

// GetTokenAtIndex is a free data retrieval call binding the contract method 0x97988dce.
//
// Solidity: function getTokenAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetTokenAtIndex(_index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _index)
}

// GetTokenLockIDAtIndex is a free data retrieval call binding the contract method 0x8ba74f17.
//
// Solidity: function getTokenLockIDAtIndex(address _token, uint256 _index) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetTokenLockIDAtIndex(opts *bind.CallOpts, _token common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getTokenLockIDAtIndex", _token, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenLockIDAtIndex is a free data retrieval call binding the contract method 0x8ba74f17.
//
// Solidity: function getTokenLockIDAtIndex(address _token, uint256 _index) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetTokenLockIDAtIndex(_token common.Address, _index *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetTokenLockIDAtIndex(&_UnicryptTokenVesting.CallOpts, _token, _index)
}

// GetTokenLockIDAtIndex is a free data retrieval call binding the contract method 0x8ba74f17.
//
// Solidity: function getTokenLockIDAtIndex(address _token, uint256 _index) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetTokenLockIDAtIndex(_token common.Address, _index *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetTokenLockIDAtIndex(&_UnicryptTokenVesting.CallOpts, _token, _index)
}

// GetTokenLocksLength is a free data retrieval call binding the contract method 0xd060e175.
//
// Solidity: function getTokenLocksLength(address _token) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetTokenLocksLength(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getTokenLocksLength", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenLocksLength is a free data retrieval call binding the contract method 0xd060e175.
//
// Solidity: function getTokenLocksLength(address _token) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetTokenLocksLength(_token common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetTokenLocksLength(&_UnicryptTokenVesting.CallOpts, _token)
}

// GetTokenLocksLength is a free data retrieval call binding the contract method 0xd060e175.
//
// Solidity: function getTokenLocksLength(address _token) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetTokenLocksLength(_token common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetTokenLocksLength(&_UnicryptTokenVesting.CallOpts, _token)
}

// GetTokenWhitelisterAtIndex is a free data retrieval call binding the contract method 0xe52c4b7a.
//
// Solidity: function getTokenWhitelisterAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetTokenWhitelisterAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getTokenWhitelisterAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWhitelisterAtIndex is a free data retrieval call binding the contract method 0xe52c4b7a.
//
// Solidity: function getTokenWhitelisterAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetTokenWhitelisterAtIndex(_index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetTokenWhitelisterAtIndex(&_UnicryptTokenVesting.CallOpts, _index)
}

// GetTokenWhitelisterAtIndex is a free data retrieval call binding the contract method 0xe52c4b7a.
//
// Solidity: function getTokenWhitelisterAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetTokenWhitelisterAtIndex(_index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetTokenWhitelisterAtIndex(&_UnicryptTokenVesting.CallOpts, _index)
}

// GetTokenWhitelisterLength is a free data retrieval call binding the contract method 0x00623ae3.
//
// Solidity: function getTokenWhitelisterLength() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetTokenWhitelisterLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getTokenWhitelisterLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenWhitelisterLength is a free data retrieval call binding the contract method 0x00623ae3.
//
// Solidity: function getTokenWhitelisterLength() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetTokenWhitelisterLength() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetTokenWhitelisterLength(&_UnicryptTokenVesting.CallOpts)
}

// GetTokenWhitelisterLength is a free data retrieval call binding the contract method 0x00623ae3.
//
// Solidity: function getTokenWhitelisterLength() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetTokenWhitelisterLength() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetTokenWhitelisterLength(&_UnicryptTokenVesting.CallOpts)
}

// GetTokenWhitelisterStatus is a free data retrieval call binding the contract method 0xcf6dde4a.
//
// Solidity: function getTokenWhitelisterStatus(address _user) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetTokenWhitelisterStatus(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getTokenWhitelisterStatus", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTokenWhitelisterStatus is a free data retrieval call binding the contract method 0xcf6dde4a.
//
// Solidity: function getTokenWhitelisterStatus(address _user) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetTokenWhitelisterStatus(_user common.Address) (bool, error) {
	return _UnicryptTokenVesting.Contract.GetTokenWhitelisterStatus(&_UnicryptTokenVesting.CallOpts, _user)
}

// GetTokenWhitelisterStatus is a free data retrieval call binding the contract method 0xcf6dde4a.
//
// Solidity: function getTokenWhitelisterStatus(address _user) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetTokenWhitelisterStatus(_user common.Address) (bool, error) {
	return _UnicryptTokenVesting.Contract.GetTokenWhitelisterStatus(&_UnicryptTokenVesting.CallOpts, _user)
}

// GetUserLockIDForTokenAtIndex is a free data retrieval call binding the contract method 0x1c30ffb1.
//
// Solidity: function getUserLockIDForTokenAtIndex(address _user, address _token, uint256 _index) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetUserLockIDForTokenAtIndex(opts *bind.CallOpts, _user common.Address, _token common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getUserLockIDForTokenAtIndex", _user, _token, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLockIDForTokenAtIndex is a free data retrieval call binding the contract method 0x1c30ffb1.
//
// Solidity: function getUserLockIDForTokenAtIndex(address _user, address _token, uint256 _index) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetUserLockIDForTokenAtIndex(_user common.Address, _token common.Address, _index *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetUserLockIDForTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _user, _token, _index)
}

// GetUserLockIDForTokenAtIndex is a free data retrieval call binding the contract method 0x1c30ffb1.
//
// Solidity: function getUserLockIDForTokenAtIndex(address _user, address _token, uint256 _index) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetUserLockIDForTokenAtIndex(_user common.Address, _token common.Address, _index *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetUserLockIDForTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _user, _token, _index)
}

// GetUserLockedTokenAtIndex is a free data retrieval call binding the contract method 0x903df806.
//
// Solidity: function getUserLockedTokenAtIndex(address _user, uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetUserLockedTokenAtIndex(opts *bind.CallOpts, _user common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getUserLockedTokenAtIndex", _user, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserLockedTokenAtIndex is a free data retrieval call binding the contract method 0x903df806.
//
// Solidity: function getUserLockedTokenAtIndex(address _user, uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetUserLockedTokenAtIndex(_user common.Address, _index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetUserLockedTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _user, _index)
}

// GetUserLockedTokenAtIndex is a free data retrieval call binding the contract method 0x903df806.
//
// Solidity: function getUserLockedTokenAtIndex(address _user, uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetUserLockedTokenAtIndex(_user common.Address, _index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetUserLockedTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _user, _index)
}

// GetUserLockedTokensLength is a free data retrieval call binding the contract method 0xa6dcb8de.
//
// Solidity: function getUserLockedTokensLength(address _user) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetUserLockedTokensLength(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getUserLockedTokensLength", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLockedTokensLength is a free data retrieval call binding the contract method 0xa6dcb8de.
//
// Solidity: function getUserLockedTokensLength(address _user) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetUserLockedTokensLength(_user common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetUserLockedTokensLength(&_UnicryptTokenVesting.CallOpts, _user)
}

// GetUserLockedTokensLength is a free data retrieval call binding the contract method 0xa6dcb8de.
//
// Solidity: function getUserLockedTokensLength(address _user) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetUserLockedTokensLength(_user common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetUserLockedTokensLength(&_UnicryptTokenVesting.CallOpts, _user)
}

// GetUserLocksForTokenLength is a free data retrieval call binding the contract method 0xbb5ee001.
//
// Solidity: function getUserLocksForTokenLength(address _user, address _token) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetUserLocksForTokenLength(opts *bind.CallOpts, _user common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getUserLocksForTokenLength", _user, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLocksForTokenLength is a free data retrieval call binding the contract method 0xbb5ee001.
//
// Solidity: function getUserLocksForTokenLength(address _user, address _token) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetUserLocksForTokenLength(_user common.Address, _token common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetUserLocksForTokenLength(&_UnicryptTokenVesting.CallOpts, _user, _token)
}

// GetUserLocksForTokenLength is a free data retrieval call binding the contract method 0xbb5ee001.
//
// Solidity: function getUserLocksForTokenLength(address _user, address _token) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetUserLocksForTokenLength(_user common.Address, _token common.Address) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetUserLocksForTokenLength(&_UnicryptTokenVesting.CallOpts, _user, _token)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xe04ab139.
//
// Solidity: function getWithdrawableShares(uint256 _lockID) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetWithdrawableShares(opts *bind.CallOpts, _lockID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getWithdrawableShares", _lockID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xe04ab139.
//
// Solidity: function getWithdrawableShares(uint256 _lockID) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetWithdrawableShares(_lockID *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetWithdrawableShares(&_UnicryptTokenVesting.CallOpts, _lockID)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xe04ab139.
//
// Solidity: function getWithdrawableShares(uint256 _lockID) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetWithdrawableShares(_lockID *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetWithdrawableShares(&_UnicryptTokenVesting.CallOpts, _lockID)
}

// GetWithdrawableTokens is a free data retrieval call binding the contract method 0xd323cdbf.
//
// Solidity: function getWithdrawableTokens(uint256 _lockID) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetWithdrawableTokens(opts *bind.CallOpts, _lockID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getWithdrawableTokens", _lockID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableTokens is a free data retrieval call binding the contract method 0xd323cdbf.
//
// Solidity: function getWithdrawableTokens(uint256 _lockID) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetWithdrawableTokens(_lockID *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetWithdrawableTokens(&_UnicryptTokenVesting.CallOpts, _lockID)
}

// GetWithdrawableTokens is a free data retrieval call binding the contract method 0xd323cdbf.
//
// Solidity: function getWithdrawableTokens(uint256 _lockID) view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetWithdrawableTokens(_lockID *big.Int) (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetWithdrawableTokens(&_UnicryptTokenVesting.CallOpts, _lockID)
}

// GetZeroFeeTokenAtIndex is a free data retrieval call binding the contract method 0xb1d7655c.
//
// Solidity: function getZeroFeeTokenAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetZeroFeeTokenAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getZeroFeeTokenAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZeroFeeTokenAtIndex is a free data retrieval call binding the contract method 0xb1d7655c.
//
// Solidity: function getZeroFeeTokenAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetZeroFeeTokenAtIndex(_index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetZeroFeeTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _index)
}

// GetZeroFeeTokenAtIndex is a free data retrieval call binding the contract method 0xb1d7655c.
//
// Solidity: function getZeroFeeTokenAtIndex(uint256 _index) view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetZeroFeeTokenAtIndex(_index *big.Int) (common.Address, error) {
	return _UnicryptTokenVesting.Contract.GetZeroFeeTokenAtIndex(&_UnicryptTokenVesting.CallOpts, _index)
}

// GetZeroFeeTokensLength is a free data retrieval call binding the contract method 0x03e1cdf4.
//
// Solidity: function getZeroFeeTokensLength() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) GetZeroFeeTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "getZeroFeeTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZeroFeeTokensLength is a free data retrieval call binding the contract method 0x03e1cdf4.
//
// Solidity: function getZeroFeeTokensLength() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) GetZeroFeeTokensLength() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetZeroFeeTokensLength(&_UnicryptTokenVesting.CallOpts)
}

// GetZeroFeeTokensLength is a free data retrieval call binding the contract method 0x03e1cdf4.
//
// Solidity: function getZeroFeeTokensLength() view returns(uint256)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) GetZeroFeeTokensLength() (*big.Int, error) {
	return _UnicryptTokenVesting.Contract.GetZeroFeeTokensLength(&_UnicryptTokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) Owner() (common.Address, error) {
	return _UnicryptTokenVesting.Contract.Owner(&_UnicryptTokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) Owner() (common.Address, error) {
	return _UnicryptTokenVesting.Contract.Owner(&_UnicryptTokenVesting.CallOpts)
}

// TestCondition is a free data retrieval call binding the contract method 0x675187a3.
//
// Solidity: function testCondition(address condition) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) TestCondition(opts *bind.CallOpts, condition common.Address) (bool, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "testCondition", condition)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestCondition is a free data retrieval call binding the contract method 0x675187a3.
//
// Solidity: function testCondition(address condition) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) TestCondition(condition common.Address) (bool, error) {
	return _UnicryptTokenVesting.Contract.TestCondition(&_UnicryptTokenVesting.CallOpts, condition)
}

// TestCondition is a free data retrieval call binding the contract method 0x675187a3.
//
// Solidity: function testCondition(address condition) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) TestCondition(condition common.Address) (bool, error) {
	return _UnicryptTokenVesting.Contract.TestCondition(&_UnicryptTokenVesting.CallOpts, condition)
}

// TokenOnZeroFeeWhitelist is a free data retrieval call binding the contract method 0x1d7065db.
//
// Solidity: function tokenOnZeroFeeWhitelist(address _token) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingCaller) TokenOnZeroFeeWhitelist(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _UnicryptTokenVesting.contract.Call(opts, &out, "tokenOnZeroFeeWhitelist", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenOnZeroFeeWhitelist is a free data retrieval call binding the contract method 0x1d7065db.
//
// Solidity: function tokenOnZeroFeeWhitelist(address _token) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) TokenOnZeroFeeWhitelist(_token common.Address) (bool, error) {
	return _UnicryptTokenVesting.Contract.TokenOnZeroFeeWhitelist(&_UnicryptTokenVesting.CallOpts, _token)
}

// TokenOnZeroFeeWhitelist is a free data retrieval call binding the contract method 0x1d7065db.
//
// Solidity: function tokenOnZeroFeeWhitelist(address _token) view returns(bool)
func (_UnicryptTokenVesting *UnicryptTokenVestingCallerSession) TokenOnZeroFeeWhitelist(_token common.Address) (bool, error) {
	return _UnicryptTokenVesting.Contract.TokenOnZeroFeeWhitelist(&_UnicryptTokenVesting.CallOpts, _token)
}

// AdminSetWhitelister is a paid mutator transaction binding the contract method 0x86de2fcb.
//
// Solidity: function adminSetWhitelister(address _user, bool _add) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) AdminSetWhitelister(opts *bind.TransactOpts, _user common.Address, _add bool) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "adminSetWhitelister", _user, _add)
}

// AdminSetWhitelister is a paid mutator transaction binding the contract method 0x86de2fcb.
//
// Solidity: function adminSetWhitelister(address _user, bool _add) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) AdminSetWhitelister(_user common.Address, _add bool) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.AdminSetWhitelister(&_UnicryptTokenVesting.TransactOpts, _user, _add)
}

// AdminSetWhitelister is a paid mutator transaction binding the contract method 0x86de2fcb.
//
// Solidity: function adminSetWhitelister(address _user, bool _add) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) AdminSetWhitelister(_user common.Address, _add bool) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.AdminSetWhitelister(&_UnicryptTokenVesting.TransactOpts, _user, _add)
}

// EditZeroFeeWhitelist is a paid mutator transaction binding the contract method 0xb4540fa7.
//
// Solidity: function editZeroFeeWhitelist(address _token, bool _add) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) EditZeroFeeWhitelist(opts *bind.TransactOpts, _token common.Address, _add bool) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "editZeroFeeWhitelist", _token, _add)
}

// EditZeroFeeWhitelist is a paid mutator transaction binding the contract method 0xb4540fa7.
//
// Solidity: function editZeroFeeWhitelist(address _token, bool _add) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) EditZeroFeeWhitelist(_token common.Address, _add bool) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.EditZeroFeeWhitelist(&_UnicryptTokenVesting.TransactOpts, _token, _add)
}

// EditZeroFeeWhitelist is a paid mutator transaction binding the contract method 0xb4540fa7.
//
// Solidity: function editZeroFeeWhitelist(address _token, bool _add) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) EditZeroFeeWhitelist(_token common.Address, _add bool) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.EditZeroFeeWhitelist(&_UnicryptTokenVesting.TransactOpts, _token, _add)
}

// IncrementLock is a paid mutator transaction binding the contract method 0x3717dee7.
//
// Solidity: function incrementLock(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) IncrementLock(opts *bind.TransactOpts, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "incrementLock", _lockID, _amount)
}

// IncrementLock is a paid mutator transaction binding the contract method 0x3717dee7.
//
// Solidity: function incrementLock(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) IncrementLock(_lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.IncrementLock(&_UnicryptTokenVesting.TransactOpts, _lockID, _amount)
}

// IncrementLock is a paid mutator transaction binding the contract method 0x3717dee7.
//
// Solidity: function incrementLock(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) IncrementLock(_lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.IncrementLock(&_UnicryptTokenVesting.TransactOpts, _lockID, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x13ef2b1b.
//
// Solidity: function lock(address _token, (address,uint256,uint256,uint256,address)[] _lock_params) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) Lock(opts *bind.TransactOpts, _token common.Address, _lock_params []TokenVestingLockParams) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "lock", _token, _lock_params)
}

// Lock is a paid mutator transaction binding the contract method 0x13ef2b1b.
//
// Solidity: function lock(address _token, (address,uint256,uint256,uint256,address)[] _lock_params) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) Lock(_token common.Address, _lock_params []TokenVestingLockParams) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Lock(&_UnicryptTokenVesting.TransactOpts, _token, _lock_params)
}

// Lock is a paid mutator transaction binding the contract method 0x13ef2b1b.
//
// Solidity: function lock(address _token, (address,uint256,uint256,uint256,address)[] _lock_params) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) Lock(_token common.Address, _lock_params []TokenVestingLockParams) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Lock(&_UnicryptTokenVesting.TransactOpts, _token, _lock_params)
}

// Migrate is a paid mutator transaction binding the contract method 0x3e54bacb.
//
// Solidity: function migrate(uint256 _lockID, uint256 _option) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) Migrate(opts *bind.TransactOpts, _lockID *big.Int, _option *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "migrate", _lockID, _option)
}

// Migrate is a paid mutator transaction binding the contract method 0x3e54bacb.
//
// Solidity: function migrate(uint256 _lockID, uint256 _option) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) Migrate(_lockID *big.Int, _option *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Migrate(&_UnicryptTokenVesting.TransactOpts, _lockID, _option)
}

// Migrate is a paid mutator transaction binding the contract method 0x3e54bacb.
//
// Solidity: function migrate(uint256 _lockID, uint256 _option) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) Migrate(_lockID *big.Int, _option *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Migrate(&_UnicryptTokenVesting.TransactOpts, _lockID, _option)
}

// PayForFreeTokenLocks is a paid mutator transaction binding the contract method 0xa34df14f.
//
// Solidity: function payForFreeTokenLocks(address _token) payable returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) PayForFreeTokenLocks(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "payForFreeTokenLocks", _token)
}

// PayForFreeTokenLocks is a paid mutator transaction binding the contract method 0xa34df14f.
//
// Solidity: function payForFreeTokenLocks(address _token) payable returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) PayForFreeTokenLocks(_token common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.PayForFreeTokenLocks(&_UnicryptTokenVesting.TransactOpts, _token)
}

// PayForFreeTokenLocks is a paid mutator transaction binding the contract method 0xa34df14f.
//
// Solidity: function payForFreeTokenLocks(address _token) payable returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) PayForFreeTokenLocks(_token common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.PayForFreeTokenLocks(&_UnicryptTokenVesting.TransactOpts, _token)
}

// Relock is a paid mutator transaction binding the contract method 0xb2fb30cb.
//
// Solidity: function relock(uint256 _lockID, uint256 _unlock_date) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) Relock(opts *bind.TransactOpts, _lockID *big.Int, _unlock_date *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "relock", _lockID, _unlock_date)
}

// Relock is a paid mutator transaction binding the contract method 0xb2fb30cb.
//
// Solidity: function relock(uint256 _lockID, uint256 _unlock_date) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) Relock(_lockID *big.Int, _unlock_date *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Relock(&_UnicryptTokenVesting.TransactOpts, _lockID, _unlock_date)
}

// Relock is a paid mutator transaction binding the contract method 0xb2fb30cb.
//
// Solidity: function relock(uint256 _lockID, uint256 _unlock_date) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) Relock(_lockID *big.Int, _unlock_date *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Relock(&_UnicryptTokenVesting.TransactOpts, _lockID, _unlock_date)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.RenounceOwnership(&_UnicryptTokenVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.RenounceOwnership(&_UnicryptTokenVesting.TransactOpts)
}

// RevokeCondition is a paid mutator transaction binding the contract method 0x5a5b8d9e.
//
// Solidity: function revokeCondition(uint256 _lockID) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) RevokeCondition(opts *bind.TransactOpts, _lockID *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "revokeCondition", _lockID)
}

// RevokeCondition is a paid mutator transaction binding the contract method 0x5a5b8d9e.
//
// Solidity: function revokeCondition(uint256 _lockID) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) RevokeCondition(_lockID *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.RevokeCondition(&_UnicryptTokenVesting.TransactOpts, _lockID)
}

// RevokeCondition is a paid mutator transaction binding the contract method 0x5a5b8d9e.
//
// Solidity: function revokeCondition(uint256 _lockID) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) RevokeCondition(_lockID *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.RevokeCondition(&_UnicryptTokenVesting.TransactOpts, _lockID)
}

// SetBlacklistContract is a paid mutator transaction binding the contract method 0x741af17e.
//
// Solidity: function setBlacklistContract(address _contract) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) SetBlacklistContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "setBlacklistContract", _contract)
}

// SetBlacklistContract is a paid mutator transaction binding the contract method 0x741af17e.
//
// Solidity: function setBlacklistContract(address _contract) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) SetBlacklistContract(_contract common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SetBlacklistContract(&_UnicryptTokenVesting.TransactOpts, _contract)
}

// SetBlacklistContract is a paid mutator transaction binding the contract method 0x741af17e.
//
// Solidity: function setBlacklistContract(address _contract) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) SetBlacklistContract(_contract common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SetBlacklistContract(&_UnicryptTokenVesting.TransactOpts, _contract)
}

// SetFees is a paid mutator transaction binding the contract method 0xfc0633d0.
//
// Solidity: function setFees(uint256 _tokenFee, uint256 _freeLockingFee, address _feeAddress, address _freeLockingToken) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) SetFees(opts *bind.TransactOpts, _tokenFee *big.Int, _freeLockingFee *big.Int, _feeAddress common.Address, _freeLockingToken common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "setFees", _tokenFee, _freeLockingFee, _feeAddress, _freeLockingToken)
}

// SetFees is a paid mutator transaction binding the contract method 0xfc0633d0.
//
// Solidity: function setFees(uint256 _tokenFee, uint256 _freeLockingFee, address _feeAddress, address _freeLockingToken) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) SetFees(_tokenFee *big.Int, _freeLockingFee *big.Int, _feeAddress common.Address, _freeLockingToken common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SetFees(&_UnicryptTokenVesting.TransactOpts, _tokenFee, _freeLockingFee, _feeAddress, _freeLockingToken)
}

// SetFees is a paid mutator transaction binding the contract method 0xfc0633d0.
//
// Solidity: function setFees(uint256 _tokenFee, uint256 _freeLockingFee, address _feeAddress, address _freeLockingToken) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) SetFees(_tokenFee *big.Int, _freeLockingFee *big.Int, _feeAddress common.Address, _freeLockingToken common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SetFees(&_UnicryptTokenVesting.TransactOpts, _tokenFee, _freeLockingFee, _feeAddress, _freeLockingToken)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SetMigrator(&_UnicryptTokenVesting.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SetMigrator(&_UnicryptTokenVesting.TransactOpts, _migrator)
}

// SplitLock is a paid mutator transaction binding the contract method 0x6588fc03.
//
// Solidity: function splitLock(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) SplitLock(opts *bind.TransactOpts, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "splitLock", _lockID, _amount)
}

// SplitLock is a paid mutator transaction binding the contract method 0x6588fc03.
//
// Solidity: function splitLock(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) SplitLock(_lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SplitLock(&_UnicryptTokenVesting.TransactOpts, _lockID, _amount)
}

// SplitLock is a paid mutator transaction binding the contract method 0x6588fc03.
//
// Solidity: function splitLock(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) SplitLock(_lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.SplitLock(&_UnicryptTokenVesting.TransactOpts, _lockID, _amount)
}

// TransferLockOwnership is a paid mutator transaction binding the contract method 0x5a04fb69.
//
// Solidity: function transferLockOwnership(uint256 _lockID, address _newOwner) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) TransferLockOwnership(opts *bind.TransactOpts, _lockID *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "transferLockOwnership", _lockID, _newOwner)
}

// TransferLockOwnership is a paid mutator transaction binding the contract method 0x5a04fb69.
//
// Solidity: function transferLockOwnership(uint256 _lockID, address _newOwner) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) TransferLockOwnership(_lockID *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.TransferLockOwnership(&_UnicryptTokenVesting.TransactOpts, _lockID, _newOwner)
}

// TransferLockOwnership is a paid mutator transaction binding the contract method 0x5a04fb69.
//
// Solidity: function transferLockOwnership(uint256 _lockID, address _newOwner) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) TransferLockOwnership(_lockID *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.TransferLockOwnership(&_UnicryptTokenVesting.TransactOpts, _lockID, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.TransferOwnership(&_UnicryptTokenVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.TransferOwnership(&_UnicryptTokenVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactor) Withdraw(opts *bind.TransactOpts, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.contract.Transact(opts, "withdraw", _lockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingSession) Withdraw(_lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Withdraw(&_UnicryptTokenVesting.TransactOpts, _lockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _lockID, uint256 _amount) returns()
func (_UnicryptTokenVesting *UnicryptTokenVestingTransactorSession) Withdraw(_lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _UnicryptTokenVesting.Contract.Withdraw(&_UnicryptTokenVesting.TransactOpts, _lockID, _amount)
}

// UnicryptTokenVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOwnershipTransferredIterator struct {
	Event *UnicryptTokenVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOwnershipTransferred)
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
		it.Event = new(UnicryptTokenVestingOwnershipTransferred)
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
func (it *UnicryptTokenVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOwnershipTransferred represents a OwnershipTransferred event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnicryptTokenVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOwnershipTransferredIterator{contract: _UnicryptTokenVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOwnershipTransferred)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOwnershipTransferred(log types.Log) (*UnicryptTokenVestingOwnershipTransferred, error) {
	event := new(UnicryptTokenVestingOwnershipTransferred)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptTokenVestingOnLockIterator is returned from FilterOnLock and is used to iterate over the raw logs and unpacked data for OnLock events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnLockIterator struct {
	Event *UnicryptTokenVestingOnLock // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOnLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOnLock)
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
		it.Event = new(UnicryptTokenVestingOnLock)
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
func (it *UnicryptTokenVestingOnLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOnLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOnLock represents a OnLock event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnLock struct {
	LockID         *big.Int
	Token          common.Address
	Owner          common.Address
	AmountInTokens *big.Int
	StartEmission  *big.Int
	EndEmission    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnLock is a free log retrieval operation binding the contract event 0x4d0f9887048089b093c92bec885ab529000723bce1a79f4e81a5990619910b96.
//
// Solidity: event onLock(uint256 lockID, address token, address owner, uint256 amountInTokens, uint256 startEmission, uint256 endEmission)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOnLock(opts *bind.FilterOpts) (*UnicryptTokenVestingOnLockIterator, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "onLock")
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOnLockIterator{contract: _UnicryptTokenVesting.contract, event: "onLock", logs: logs, sub: sub}, nil
}

// WatchOnLock is a free log subscription operation binding the contract event 0x4d0f9887048089b093c92bec885ab529000723bce1a79f4e81a5990619910b96.
//
// Solidity: event onLock(uint256 lockID, address token, address owner, uint256 amountInTokens, uint256 startEmission, uint256 endEmission)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOnLock(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOnLock) (event.Subscription, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "onLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOnLock)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onLock", log); err != nil {
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

// ParseOnLock is a log parse operation binding the contract event 0x4d0f9887048089b093c92bec885ab529000723bce1a79f4e81a5990619910b96.
//
// Solidity: event onLock(uint256 lockID, address token, address owner, uint256 amountInTokens, uint256 startEmission, uint256 endEmission)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOnLock(log types.Log) (*UnicryptTokenVestingOnLock, error) {
	event := new(UnicryptTokenVestingOnLock)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptTokenVestingOnMigrateIterator is returned from FilterOnMigrate and is used to iterate over the raw logs and unpacked data for OnMigrate events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnMigrateIterator struct {
	Event *UnicryptTokenVestingOnMigrate // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOnMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOnMigrate)
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
		it.Event = new(UnicryptTokenVestingOnMigrate)
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
func (it *UnicryptTokenVestingOnMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOnMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOnMigrate represents a OnMigrate event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnMigrate struct {
	LockID         *big.Int
	AmountInTokens *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnMigrate is a free log retrieval operation binding the contract event 0x8b1497abd425402b9139208bcb7f83d61a7545712a632d73a1ae68039cd593ab.
//
// Solidity: event onMigrate(uint256 lockID, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOnMigrate(opts *bind.FilterOpts) (*UnicryptTokenVestingOnMigrateIterator, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "onMigrate")
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOnMigrateIterator{contract: _UnicryptTokenVesting.contract, event: "onMigrate", logs: logs, sub: sub}, nil
}

// WatchOnMigrate is a free log subscription operation binding the contract event 0x8b1497abd425402b9139208bcb7f83d61a7545712a632d73a1ae68039cd593ab.
//
// Solidity: event onMigrate(uint256 lockID, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOnMigrate(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOnMigrate) (event.Subscription, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "onMigrate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOnMigrate)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onMigrate", log); err != nil {
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

// ParseOnMigrate is a log parse operation binding the contract event 0x8b1497abd425402b9139208bcb7f83d61a7545712a632d73a1ae68039cd593ab.
//
// Solidity: event onMigrate(uint256 lockID, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOnMigrate(log types.Log) (*UnicryptTokenVestingOnMigrate, error) {
	event := new(UnicryptTokenVestingOnMigrate)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onMigrate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptTokenVestingOnRelockIterator is returned from FilterOnRelock and is used to iterate over the raw logs and unpacked data for OnRelock events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnRelockIterator struct {
	Event *UnicryptTokenVestingOnRelock // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOnRelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOnRelock)
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
		it.Event = new(UnicryptTokenVestingOnRelock)
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
func (it *UnicryptTokenVestingOnRelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOnRelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOnRelock represents a OnRelock event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnRelock struct {
	LockID     *big.Int
	UnlockDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnRelock is a free log retrieval operation binding the contract event 0xefaff1b90138281d215452c67f793017f52e456f65c28ac63f5309a89a059b47.
//
// Solidity: event onRelock(uint256 lockID, uint256 unlockDate)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOnRelock(opts *bind.FilterOpts) (*UnicryptTokenVestingOnRelockIterator, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "onRelock")
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOnRelockIterator{contract: _UnicryptTokenVesting.contract, event: "onRelock", logs: logs, sub: sub}, nil
}

// WatchOnRelock is a free log subscription operation binding the contract event 0xefaff1b90138281d215452c67f793017f52e456f65c28ac63f5309a89a059b47.
//
// Solidity: event onRelock(uint256 lockID, uint256 unlockDate)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOnRelock(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOnRelock) (event.Subscription, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "onRelock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOnRelock)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onRelock", log); err != nil {
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

// ParseOnRelock is a log parse operation binding the contract event 0xefaff1b90138281d215452c67f793017f52e456f65c28ac63f5309a89a059b47.
//
// Solidity: event onRelock(uint256 lockID, uint256 unlockDate)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOnRelock(log types.Log) (*UnicryptTokenVestingOnRelock, error) {
	event := new(UnicryptTokenVestingOnRelock)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onRelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptTokenVestingOnSplitLockIterator is returned from FilterOnSplitLock and is used to iterate over the raw logs and unpacked data for OnSplitLock events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnSplitLockIterator struct {
	Event *UnicryptTokenVestingOnSplitLock // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOnSplitLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOnSplitLock)
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
		it.Event = new(UnicryptTokenVestingOnSplitLock)
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
func (it *UnicryptTokenVestingOnSplitLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOnSplitLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOnSplitLock represents a OnSplitLock event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnSplitLock struct {
	FromLockID     *big.Int
	ToLockID       *big.Int
	AmountInTokens *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnSplitLock is a free log retrieval operation binding the contract event 0xbfc02f4250f5e54a31a72371b1932061a0245db2afb5124b547306f6ebffd8ad.
//
// Solidity: event onSplitLock(uint256 fromLockID, uint256 toLockID, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOnSplitLock(opts *bind.FilterOpts) (*UnicryptTokenVestingOnSplitLockIterator, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "onSplitLock")
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOnSplitLockIterator{contract: _UnicryptTokenVesting.contract, event: "onSplitLock", logs: logs, sub: sub}, nil
}

// WatchOnSplitLock is a free log subscription operation binding the contract event 0xbfc02f4250f5e54a31a72371b1932061a0245db2afb5124b547306f6ebffd8ad.
//
// Solidity: event onSplitLock(uint256 fromLockID, uint256 toLockID, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOnSplitLock(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOnSplitLock) (event.Subscription, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "onSplitLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOnSplitLock)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onSplitLock", log); err != nil {
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

// ParseOnSplitLock is a log parse operation binding the contract event 0xbfc02f4250f5e54a31a72371b1932061a0245db2afb5124b547306f6ebffd8ad.
//
// Solidity: event onSplitLock(uint256 fromLockID, uint256 toLockID, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOnSplitLock(log types.Log) (*UnicryptTokenVestingOnSplitLock, error) {
	event := new(UnicryptTokenVestingOnSplitLock)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onSplitLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptTokenVestingOnTransferLockIterator is returned from FilterOnTransferLock and is used to iterate over the raw logs and unpacked data for OnTransferLock events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnTransferLockIterator struct {
	Event *UnicryptTokenVestingOnTransferLock // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOnTransferLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOnTransferLock)
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
		it.Event = new(UnicryptTokenVestingOnTransferLock)
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
func (it *UnicryptTokenVestingOnTransferLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOnTransferLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOnTransferLock represents a OnTransferLock event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnTransferLock struct {
	LockIDFrom *big.Int
	LockIDto   *big.Int
	OldOwner   common.Address
	NewOwner   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnTransferLock is a free log retrieval operation binding the contract event 0xc1329eea12893a6eff780df43cf4ec708fdc6ce7fbacf84ee8cf9355a9af1dd8.
//
// Solidity: event onTransferLock(uint256 lockIDFrom, uint256 lockIDto, address oldOwner, address newOwner)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOnTransferLock(opts *bind.FilterOpts) (*UnicryptTokenVestingOnTransferLockIterator, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "onTransferLock")
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOnTransferLockIterator{contract: _UnicryptTokenVesting.contract, event: "onTransferLock", logs: logs, sub: sub}, nil
}

// WatchOnTransferLock is a free log subscription operation binding the contract event 0xc1329eea12893a6eff780df43cf4ec708fdc6ce7fbacf84ee8cf9355a9af1dd8.
//
// Solidity: event onTransferLock(uint256 lockIDFrom, uint256 lockIDto, address oldOwner, address newOwner)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOnTransferLock(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOnTransferLock) (event.Subscription, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "onTransferLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOnTransferLock)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onTransferLock", log); err != nil {
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

// ParseOnTransferLock is a log parse operation binding the contract event 0xc1329eea12893a6eff780df43cf4ec708fdc6ce7fbacf84ee8cf9355a9af1dd8.
//
// Solidity: event onTransferLock(uint256 lockIDFrom, uint256 lockIDto, address oldOwner, address newOwner)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOnTransferLock(log types.Log) (*UnicryptTokenVestingOnTransferLock, error) {
	event := new(UnicryptTokenVestingOnTransferLock)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onTransferLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptTokenVestingOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnWithdrawIterator struct {
	Event *UnicryptTokenVestingOnWithdraw // Event containing the contract specifics and raw log

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
func (it *UnicryptTokenVestingOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptTokenVestingOnWithdraw)
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
		it.Event = new(UnicryptTokenVestingOnWithdraw)
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
func (it *UnicryptTokenVestingOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptTokenVestingOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptTokenVestingOnWithdraw represents a OnWithdraw event raised by the UnicryptTokenVesting contract.
type UnicryptTokenVestingOnWithdraw struct {
	LpToken        common.Address
	AmountInTokens *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(address lpToken, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) FilterOnWithdraw(opts *bind.FilterOpts) (*UnicryptTokenVestingOnWithdrawIterator, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.FilterLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return &UnicryptTokenVestingOnWithdrawIterator{contract: _UnicryptTokenVesting.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(address lpToken, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *UnicryptTokenVestingOnWithdraw) (event.Subscription, error) {

	logs, sub, err := _UnicryptTokenVesting.contract.WatchLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptTokenVestingOnWithdraw)
				if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// ParseOnWithdraw is a log parse operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(address lpToken, uint256 amountInTokens)
func (_UnicryptTokenVesting *UnicryptTokenVestingFilterer) ParseOnWithdraw(log types.Log) (*UnicryptTokenVestingOnWithdraw, error) {
	event := new(UnicryptTokenVestingOnWithdraw)
	if err := _UnicryptTokenVesting.contract.UnpackLog(event, "onWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
