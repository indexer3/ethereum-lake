// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maple

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

// MapleXmplMetaData contains all meta data concerning the MapleXmpl contract.
var MapleXmplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"precision_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freeAssets_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issuanceRate_\",\"type\":\"uint256\"}],\"name\":\"IssuanceParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MigrationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAsset_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"MigrationPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAsset_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAsset_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"migrator_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"migrationTime_\",\"type\":\"uint256\"}],\"name\":\"MigrationScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"OwnershipAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner_\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vestingPeriodFinish_\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_MIGRATION_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"balanceOfAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balanceOfAssets_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedAmount_\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"}],\"name\":\"depositWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedAmount_\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuanceRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAssets_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares_\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxAssets_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAssets_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"}],\"name\":\"mintWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAsset_\",\"type\":\"address\"}],\"name\":\"scheduleMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduledMigrationTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduledMigrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduledNewAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner_\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalManagedAssets_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vestingPeriod_\",\"type\":\"uint256\"}],\"name\":\"updateVestingSchedule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"issuanceRate_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeAssets_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vestingPeriodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MapleXmplABI is the input ABI used to generate the binding from.
// Deprecated: Use MapleXmplMetaData.ABI instead.
var MapleXmplABI = MapleXmplMetaData.ABI

// MapleXmpl is an auto generated Go binding around an Ethereum contract.
type MapleXmpl struct {
	MapleXmplCaller     // Read-only binding to the contract
	MapleXmplTransactor // Write-only binding to the contract
	MapleXmplFilterer   // Log filterer for contract events
}

// MapleXmplCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapleXmplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapleXmplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapleXmplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapleXmplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapleXmplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapleXmplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapleXmplSession struct {
	Contract     *MapleXmpl        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapleXmplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapleXmplCallerSession struct {
	Contract *MapleXmplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MapleXmplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapleXmplTransactorSession struct {
	Contract     *MapleXmplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MapleXmplRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapleXmplRaw struct {
	Contract *MapleXmpl // Generic contract binding to access the raw methods on
}

// MapleXmplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapleXmplCallerRaw struct {
	Contract *MapleXmplCaller // Generic read-only contract binding to access the raw methods on
}

// MapleXmplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapleXmplTransactorRaw struct {
	Contract *MapleXmplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapleXmpl creates a new instance of MapleXmpl, bound to a specific deployed contract.
func NewMapleXmpl(address common.Address, backend bind.ContractBackend) (*MapleXmpl, error) {
	contract, err := bindMapleXmpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MapleXmpl{MapleXmplCaller: MapleXmplCaller{contract: contract}, MapleXmplTransactor: MapleXmplTransactor{contract: contract}, MapleXmplFilterer: MapleXmplFilterer{contract: contract}}, nil
}

// NewMapleXmplCaller creates a new read-only instance of MapleXmpl, bound to a specific deployed contract.
func NewMapleXmplCaller(address common.Address, caller bind.ContractCaller) (*MapleXmplCaller, error) {
	contract, err := bindMapleXmpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapleXmplCaller{contract: contract}, nil
}

// NewMapleXmplTransactor creates a new write-only instance of MapleXmpl, bound to a specific deployed contract.
func NewMapleXmplTransactor(address common.Address, transactor bind.ContractTransactor) (*MapleXmplTransactor, error) {
	contract, err := bindMapleXmpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapleXmplTransactor{contract: contract}, nil
}

// NewMapleXmplFilterer creates a new log filterer instance of MapleXmpl, bound to a specific deployed contract.
func NewMapleXmplFilterer(address common.Address, filterer bind.ContractFilterer) (*MapleXmplFilterer, error) {
	contract, err := bindMapleXmpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapleXmplFilterer{contract: contract}, nil
}

// bindMapleXmpl binds a generic wrapper to an already deployed contract.
func bindMapleXmpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MapleXmplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapleXmpl *MapleXmplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapleXmpl.Contract.MapleXmplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapleXmpl *MapleXmplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleXmpl.Contract.MapleXmplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapleXmpl *MapleXmplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapleXmpl.Contract.MapleXmplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapleXmpl *MapleXmplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapleXmpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapleXmpl *MapleXmplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleXmpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapleXmpl *MapleXmplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapleXmpl.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator_)
func (_MapleXmpl *MapleXmplCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator_)
func (_MapleXmpl *MapleXmplSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MapleXmpl.Contract.DOMAINSEPARATOR(&_MapleXmpl.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 domainSeparator_)
func (_MapleXmpl *MapleXmplCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MapleXmpl.Contract.DOMAINSEPARATOR(&_MapleXmpl.CallOpts)
}

// MINIMUMMIGRATIONDELAY is a free data retrieval call binding the contract method 0x6e2f2aca.
//
// Solidity: function MINIMUM_MIGRATION_DELAY() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) MINIMUMMIGRATIONDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "MINIMUM_MIGRATION_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMMIGRATIONDELAY is a free data retrieval call binding the contract method 0x6e2f2aca.
//
// Solidity: function MINIMUM_MIGRATION_DELAY() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) MINIMUMMIGRATIONDELAY() (*big.Int, error) {
	return _MapleXmpl.Contract.MINIMUMMIGRATIONDELAY(&_MapleXmpl.CallOpts)
}

// MINIMUMMIGRATIONDELAY is a free data retrieval call binding the contract method 0x6e2f2aca.
//
// Solidity: function MINIMUM_MIGRATION_DELAY() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) MINIMUMMIGRATIONDELAY() (*big.Int, error) {
	return _MapleXmpl.Contract.MINIMUMMIGRATIONDELAY(&_MapleXmpl.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MapleXmpl *MapleXmplCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MapleXmpl *MapleXmplSession) PERMITTYPEHASH() ([32]byte, error) {
	return _MapleXmpl.Contract.PERMITTYPEHASH(&_MapleXmpl.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_MapleXmpl *MapleXmplCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _MapleXmpl.Contract.PERMITTYPEHASH(&_MapleXmpl.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_MapleXmpl *MapleXmplSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.Allowance(&_MapleXmpl.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.Allowance(&_MapleXmpl.CallOpts, arg0, arg1)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MapleXmpl *MapleXmplCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MapleXmpl *MapleXmplSession) Asset() (common.Address, error) {
	return _MapleXmpl.Contract.Asset(&_MapleXmpl.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MapleXmpl *MapleXmplCallerSession) Asset() (common.Address, error) {
	return _MapleXmpl.Contract.Asset(&_MapleXmpl.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MapleXmpl *MapleXmplSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.BalanceOf(&_MapleXmpl.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.BalanceOf(&_MapleXmpl.CallOpts, arg0)
}

// BalanceOfAssets is a free data retrieval call binding the contract method 0x9159b206.
//
// Solidity: function balanceOfAssets(address account_) view returns(uint256 balanceOfAssets_)
func (_MapleXmpl *MapleXmplCaller) BalanceOfAssets(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "balanceOfAssets", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAssets is a free data retrieval call binding the contract method 0x9159b206.
//
// Solidity: function balanceOfAssets(address account_) view returns(uint256 balanceOfAssets_)
func (_MapleXmpl *MapleXmplSession) BalanceOfAssets(account_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.BalanceOfAssets(&_MapleXmpl.CallOpts, account_)
}

// BalanceOfAssets is a free data retrieval call binding the contract method 0x9159b206.
//
// Solidity: function balanceOfAssets(address account_) view returns(uint256 balanceOfAssets_)
func (_MapleXmpl *MapleXmplCallerSession) BalanceOfAssets(account_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.BalanceOfAssets(&_MapleXmpl.CallOpts, account_)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplCaller) ConvertToAssets(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "convertToAssets", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplSession) ConvertToAssets(shares_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.ConvertToAssets(&_MapleXmpl.CallOpts, shares_)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplCallerSession) ConvertToAssets(shares_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.ConvertToAssets(&_MapleXmpl.CallOpts, shares_)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplCaller) ConvertToShares(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "convertToShares", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplSession) ConvertToShares(assets_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.ConvertToShares(&_MapleXmpl.CallOpts, assets_)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplCallerSession) ConvertToShares(assets_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.ConvertToShares(&_MapleXmpl.CallOpts, assets_)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MapleXmpl *MapleXmplCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MapleXmpl *MapleXmplSession) Decimals() (uint8, error) {
	return _MapleXmpl.Contract.Decimals(&_MapleXmpl.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MapleXmpl *MapleXmplCallerSession) Decimals() (uint8, error) {
	return _MapleXmpl.Contract.Decimals(&_MapleXmpl.CallOpts)
}

// FreeAssets is a free data retrieval call binding the contract method 0x11f240ac.
//
// Solidity: function freeAssets() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) FreeAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "freeAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeAssets is a free data retrieval call binding the contract method 0x11f240ac.
//
// Solidity: function freeAssets() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) FreeAssets() (*big.Int, error) {
	return _MapleXmpl.Contract.FreeAssets(&_MapleXmpl.CallOpts)
}

// FreeAssets is a free data retrieval call binding the contract method 0x11f240ac.
//
// Solidity: function freeAssets() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) FreeAssets() (*big.Int, error) {
	return _MapleXmpl.Contract.FreeAssets(&_MapleXmpl.CallOpts)
}

// IssuanceRate is a free data retrieval call binding the contract method 0x3c9ae2ba.
//
// Solidity: function issuanceRate() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) IssuanceRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "issuanceRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssuanceRate is a free data retrieval call binding the contract method 0x3c9ae2ba.
//
// Solidity: function issuanceRate() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) IssuanceRate() (*big.Int, error) {
	return _MapleXmpl.Contract.IssuanceRate(&_MapleXmpl.CallOpts)
}

// IssuanceRate is a free data retrieval call binding the contract method 0x3c9ae2ba.
//
// Solidity: function issuanceRate() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) IssuanceRate() (*big.Int, error) {
	return _MapleXmpl.Contract.IssuanceRate(&_MapleXmpl.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) LastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "lastUpdated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) LastUpdated() (*big.Int, error) {
	return _MapleXmpl.Contract.LastUpdated(&_MapleXmpl.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0xd0b06f5d.
//
// Solidity: function lastUpdated() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) LastUpdated() (*big.Int, error) {
	return _MapleXmpl.Contract.LastUpdated(&_MapleXmpl.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver_) pure returns(uint256 maxAssets_)
func (_MapleXmpl *MapleXmplCaller) MaxDeposit(opts *bind.CallOpts, receiver_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "maxDeposit", receiver_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver_) pure returns(uint256 maxAssets_)
func (_MapleXmpl *MapleXmplSession) MaxDeposit(receiver_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxDeposit(&_MapleXmpl.CallOpts, receiver_)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver_) pure returns(uint256 maxAssets_)
func (_MapleXmpl *MapleXmplCallerSession) MaxDeposit(receiver_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxDeposit(&_MapleXmpl.CallOpts, receiver_)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver_) pure returns(uint256 maxShares_)
func (_MapleXmpl *MapleXmplCaller) MaxMint(opts *bind.CallOpts, receiver_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "maxMint", receiver_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver_) pure returns(uint256 maxShares_)
func (_MapleXmpl *MapleXmplSession) MaxMint(receiver_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxMint(&_MapleXmpl.CallOpts, receiver_)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver_) pure returns(uint256 maxShares_)
func (_MapleXmpl *MapleXmplCallerSession) MaxMint(receiver_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxMint(&_MapleXmpl.CallOpts, receiver_)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner_) view returns(uint256 maxShares_)
func (_MapleXmpl *MapleXmplCaller) MaxRedeem(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "maxRedeem", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner_) view returns(uint256 maxShares_)
func (_MapleXmpl *MapleXmplSession) MaxRedeem(owner_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxRedeem(&_MapleXmpl.CallOpts, owner_)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner_) view returns(uint256 maxShares_)
func (_MapleXmpl *MapleXmplCallerSession) MaxRedeem(owner_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxRedeem(&_MapleXmpl.CallOpts, owner_)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner_) view returns(uint256 maxAssets_)
func (_MapleXmpl *MapleXmplCaller) MaxWithdraw(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "maxWithdraw", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner_) view returns(uint256 maxAssets_)
func (_MapleXmpl *MapleXmplSession) MaxWithdraw(owner_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxWithdraw(&_MapleXmpl.CallOpts, owner_)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner_) view returns(uint256 maxAssets_)
func (_MapleXmpl *MapleXmplCallerSession) MaxWithdraw(owner_ common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.MaxWithdraw(&_MapleXmpl.CallOpts, owner_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MapleXmpl *MapleXmplCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MapleXmpl *MapleXmplSession) Name() (string, error) {
	return _MapleXmpl.Contract.Name(&_MapleXmpl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MapleXmpl *MapleXmplCallerSession) Name() (string, error) {
	return _MapleXmpl.Contract.Name(&_MapleXmpl.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MapleXmpl *MapleXmplSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.Nonces(&_MapleXmpl.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MapleXmpl.Contract.Nonces(&_MapleXmpl.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MapleXmpl *MapleXmplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MapleXmpl *MapleXmplSession) Owner() (common.Address, error) {
	return _MapleXmpl.Contract.Owner(&_MapleXmpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MapleXmpl *MapleXmplCallerSession) Owner() (common.Address, error) {
	return _MapleXmpl.Contract.Owner(&_MapleXmpl.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MapleXmpl *MapleXmplCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MapleXmpl *MapleXmplSession) PendingOwner() (common.Address, error) {
	return _MapleXmpl.Contract.PendingOwner(&_MapleXmpl.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MapleXmpl *MapleXmplCallerSession) PendingOwner() (common.Address, error) {
	return _MapleXmpl.Contract.PendingOwner(&_MapleXmpl.CallOpts)
}

// Precision is a free data retrieval call binding the contract method 0xd3b5dc3b.
//
// Solidity: function precision() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) Precision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "precision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Precision is a free data retrieval call binding the contract method 0xd3b5dc3b.
//
// Solidity: function precision() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) Precision() (*big.Int, error) {
	return _MapleXmpl.Contract.Precision(&_MapleXmpl.CallOpts)
}

// Precision is a free data retrieval call binding the contract method 0xd3b5dc3b.
//
// Solidity: function precision() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) Precision() (*big.Int, error) {
	return _MapleXmpl.Contract.Precision(&_MapleXmpl.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplCaller) PreviewDeposit(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "previewDeposit", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplSession) PreviewDeposit(assets_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewDeposit(&_MapleXmpl.CallOpts, assets_)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplCallerSession) PreviewDeposit(assets_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewDeposit(&_MapleXmpl.CallOpts, assets_)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplCaller) PreviewMint(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "previewMint", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplSession) PreviewMint(shares_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewMint(&_MapleXmpl.CallOpts, shares_)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplCallerSession) PreviewMint(shares_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewMint(&_MapleXmpl.CallOpts, shares_)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplCaller) PreviewRedeem(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "previewRedeem", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplSession) PreviewRedeem(shares_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewRedeem(&_MapleXmpl.CallOpts, shares_)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares_) view returns(uint256 assets_)
func (_MapleXmpl *MapleXmplCallerSession) PreviewRedeem(shares_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewRedeem(&_MapleXmpl.CallOpts, shares_)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplCaller) PreviewWithdraw(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "previewWithdraw", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplSession) PreviewWithdraw(assets_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewWithdraw(&_MapleXmpl.CallOpts, assets_)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets_) view returns(uint256 shares_)
func (_MapleXmpl *MapleXmplCallerSession) PreviewWithdraw(assets_ *big.Int) (*big.Int, error) {
	return _MapleXmpl.Contract.PreviewWithdraw(&_MapleXmpl.CallOpts, assets_)
}

// ScheduledMigrationTimestamp is a free data retrieval call binding the contract method 0x899d0c01.
//
// Solidity: function scheduledMigrationTimestamp() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) ScheduledMigrationTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "scheduledMigrationTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScheduledMigrationTimestamp is a free data retrieval call binding the contract method 0x899d0c01.
//
// Solidity: function scheduledMigrationTimestamp() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) ScheduledMigrationTimestamp() (*big.Int, error) {
	return _MapleXmpl.Contract.ScheduledMigrationTimestamp(&_MapleXmpl.CallOpts)
}

// ScheduledMigrationTimestamp is a free data retrieval call binding the contract method 0x899d0c01.
//
// Solidity: function scheduledMigrationTimestamp() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) ScheduledMigrationTimestamp() (*big.Int, error) {
	return _MapleXmpl.Contract.ScheduledMigrationTimestamp(&_MapleXmpl.CallOpts)
}

// ScheduledMigrator is a free data retrieval call binding the contract method 0x7bb002d6.
//
// Solidity: function scheduledMigrator() view returns(address)
func (_MapleXmpl *MapleXmplCaller) ScheduledMigrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "scheduledMigrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScheduledMigrator is a free data retrieval call binding the contract method 0x7bb002d6.
//
// Solidity: function scheduledMigrator() view returns(address)
func (_MapleXmpl *MapleXmplSession) ScheduledMigrator() (common.Address, error) {
	return _MapleXmpl.Contract.ScheduledMigrator(&_MapleXmpl.CallOpts)
}

// ScheduledMigrator is a free data retrieval call binding the contract method 0x7bb002d6.
//
// Solidity: function scheduledMigrator() view returns(address)
func (_MapleXmpl *MapleXmplCallerSession) ScheduledMigrator() (common.Address, error) {
	return _MapleXmpl.Contract.ScheduledMigrator(&_MapleXmpl.CallOpts)
}

// ScheduledNewAsset is a free data retrieval call binding the contract method 0x36f8d3f1.
//
// Solidity: function scheduledNewAsset() view returns(address)
func (_MapleXmpl *MapleXmplCaller) ScheduledNewAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "scheduledNewAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScheduledNewAsset is a free data retrieval call binding the contract method 0x36f8d3f1.
//
// Solidity: function scheduledNewAsset() view returns(address)
func (_MapleXmpl *MapleXmplSession) ScheduledNewAsset() (common.Address, error) {
	return _MapleXmpl.Contract.ScheduledNewAsset(&_MapleXmpl.CallOpts)
}

// ScheduledNewAsset is a free data retrieval call binding the contract method 0x36f8d3f1.
//
// Solidity: function scheduledNewAsset() view returns(address)
func (_MapleXmpl *MapleXmplCallerSession) ScheduledNewAsset() (common.Address, error) {
	return _MapleXmpl.Contract.ScheduledNewAsset(&_MapleXmpl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MapleXmpl *MapleXmplCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MapleXmpl *MapleXmplSession) Symbol() (string, error) {
	return _MapleXmpl.Contract.Symbol(&_MapleXmpl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MapleXmpl *MapleXmplCallerSession) Symbol() (string, error) {
	return _MapleXmpl.Contract.Symbol(&_MapleXmpl.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 totalManagedAssets_)
func (_MapleXmpl *MapleXmplCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 totalManagedAssets_)
func (_MapleXmpl *MapleXmplSession) TotalAssets() (*big.Int, error) {
	return _MapleXmpl.Contract.TotalAssets(&_MapleXmpl.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 totalManagedAssets_)
func (_MapleXmpl *MapleXmplCallerSession) TotalAssets() (*big.Int, error) {
	return _MapleXmpl.Contract.TotalAssets(&_MapleXmpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) TotalSupply() (*big.Int, error) {
	return _MapleXmpl.Contract.TotalSupply(&_MapleXmpl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) TotalSupply() (*big.Int, error) {
	return _MapleXmpl.Contract.TotalSupply(&_MapleXmpl.CallOpts)
}

// VestingPeriodFinish is a free data retrieval call binding the contract method 0x3c2f7773.
//
// Solidity: function vestingPeriodFinish() view returns(uint256)
func (_MapleXmpl *MapleXmplCaller) VestingPeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleXmpl.contract.Call(opts, &out, "vestingPeriodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingPeriodFinish is a free data retrieval call binding the contract method 0x3c2f7773.
//
// Solidity: function vestingPeriodFinish() view returns(uint256)
func (_MapleXmpl *MapleXmplSession) VestingPeriodFinish() (*big.Int, error) {
	return _MapleXmpl.Contract.VestingPeriodFinish(&_MapleXmpl.CallOpts)
}

// VestingPeriodFinish is a free data retrieval call binding the contract method 0x3c2f7773.
//
// Solidity: function vestingPeriodFinish() view returns(uint256)
func (_MapleXmpl *MapleXmplCallerSession) VestingPeriodFinish() (*big.Int, error) {
	return _MapleXmpl.Contract.VestingPeriodFinish(&_MapleXmpl.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MapleXmpl *MapleXmplTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MapleXmpl *MapleXmplSession) AcceptOwnership() (*types.Transaction, error) {
	return _MapleXmpl.Contract.AcceptOwnership(&_MapleXmpl.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MapleXmpl *MapleXmplTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MapleXmpl.Contract.AcceptOwnership(&_MapleXmpl.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactor) Approve(opts *bind.TransactOpts, spender_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "approve", spender_, amount_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplSession) Approve(spender_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Approve(&_MapleXmpl.TransactOpts, spender_, amount_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactorSession) Approve(spender_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Approve(&_MapleXmpl.TransactOpts, spender_, amount_)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns()
func (_MapleXmpl *MapleXmplTransactor) CancelMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "cancelMigration")
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns()
func (_MapleXmpl *MapleXmplSession) CancelMigration() (*types.Transaction, error) {
	return _MapleXmpl.Contract.CancelMigration(&_MapleXmpl.TransactOpts)
}

// CancelMigration is a paid mutator transaction binding the contract method 0x10639ea0.
//
// Solidity: function cancelMigration() returns()
func (_MapleXmpl *MapleXmplTransactorSession) CancelMigration() (*types.Transaction, error) {
	return _MapleXmpl.Contract.CancelMigration(&_MapleXmpl.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender_, uint256 subtractedAmount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender_ common.Address, subtractedAmount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "decreaseAllowance", spender_, subtractedAmount_)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender_, uint256 subtractedAmount_) returns(bool success_)
func (_MapleXmpl *MapleXmplSession) DecreaseAllowance(spender_ common.Address, subtractedAmount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.DecreaseAllowance(&_MapleXmpl.TransactOpts, spender_, subtractedAmount_)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender_, uint256 subtractedAmount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactorSession) DecreaseAllowance(spender_ common.Address, subtractedAmount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.DecreaseAllowance(&_MapleXmpl.TransactOpts, spender_, subtractedAmount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplTransactor) Deposit(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "deposit", assets_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplSession) Deposit(assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Deposit(&_MapleXmpl.TransactOpts, assets_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplTransactorSession) Deposit(assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Deposit(&_MapleXmpl.TransactOpts, assets_, receiver_)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x50921b23.
//
// Solidity: function depositWithPermit(uint256 assets_, address receiver_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplTransactor) DepositWithPermit(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "depositWithPermit", assets_, receiver_, deadline_, v_, r_, s_)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x50921b23.
//
// Solidity: function depositWithPermit(uint256 assets_, address receiver_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplSession) DepositWithPermit(assets_ *big.Int, receiver_ common.Address, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.Contract.DepositWithPermit(&_MapleXmpl.TransactOpts, assets_, receiver_, deadline_, v_, r_, s_)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x50921b23.
//
// Solidity: function depositWithPermit(uint256 assets_, address receiver_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplTransactorSession) DepositWithPermit(assets_ *big.Int, receiver_ common.Address, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.Contract.DepositWithPermit(&_MapleXmpl.TransactOpts, assets_, receiver_, deadline_, v_, r_, s_)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender_, uint256 addedAmount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender_ common.Address, addedAmount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "increaseAllowance", spender_, addedAmount_)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender_, uint256 addedAmount_) returns(bool success_)
func (_MapleXmpl *MapleXmplSession) IncreaseAllowance(spender_ common.Address, addedAmount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.IncreaseAllowance(&_MapleXmpl.TransactOpts, spender_, addedAmount_)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender_, uint256 addedAmount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactorSession) IncreaseAllowance(spender_ common.Address, addedAmount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.IncreaseAllowance(&_MapleXmpl.TransactOpts, spender_, addedAmount_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplTransactor) Mint(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "mint", shares_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplSession) Mint(shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Mint(&_MapleXmpl.TransactOpts, shares_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplTransactorSession) Mint(shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Mint(&_MapleXmpl.TransactOpts, shares_, receiver_)
}

// MintWithPermit is a paid mutator transaction binding the contract method 0x60dd37d9.
//
// Solidity: function mintWithPermit(uint256 shares_, address receiver_, uint256 maxAssets_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplTransactor) MintWithPermit(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address, maxAssets_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "mintWithPermit", shares_, receiver_, maxAssets_, deadline_, v_, r_, s_)
}

// MintWithPermit is a paid mutator transaction binding the contract method 0x60dd37d9.
//
// Solidity: function mintWithPermit(uint256 shares_, address receiver_, uint256 maxAssets_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplSession) MintWithPermit(shares_ *big.Int, receiver_ common.Address, maxAssets_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.Contract.MintWithPermit(&_MapleXmpl.TransactOpts, shares_, receiver_, maxAssets_, deadline_, v_, r_, s_)
}

// MintWithPermit is a paid mutator transaction binding the contract method 0x60dd37d9.
//
// Solidity: function mintWithPermit(uint256 shares_, address receiver_, uint256 maxAssets_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplTransactorSession) MintWithPermit(shares_ *big.Int, receiver_ common.Address, maxAssets_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.Contract.MintWithPermit(&_MapleXmpl.TransactOpts, shares_, receiver_, maxAssets_, deadline_, v_, r_, s_)
}

// PerformMigration is a paid mutator transaction binding the contract method 0x6109ff4b.
//
// Solidity: function performMigration() returns()
func (_MapleXmpl *MapleXmplTransactor) PerformMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "performMigration")
}

// PerformMigration is a paid mutator transaction binding the contract method 0x6109ff4b.
//
// Solidity: function performMigration() returns()
func (_MapleXmpl *MapleXmplSession) PerformMigration() (*types.Transaction, error) {
	return _MapleXmpl.Contract.PerformMigration(&_MapleXmpl.TransactOpts)
}

// PerformMigration is a paid mutator transaction binding the contract method 0x6109ff4b.
//
// Solidity: function performMigration() returns()
func (_MapleXmpl *MapleXmplTransactorSession) PerformMigration() (*types.Transaction, error) {
	return _MapleXmpl.Contract.PerformMigration(&_MapleXmpl.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner_, address spender_, uint256 amount_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_MapleXmpl *MapleXmplTransactor) Permit(opts *bind.TransactOpts, owner_ common.Address, spender_ common.Address, amount_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "permit", owner_, spender_, amount_, deadline_, v_, r_, s_)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner_, address spender_, uint256 amount_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_MapleXmpl *MapleXmplSession) Permit(owner_ common.Address, spender_ common.Address, amount_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Permit(&_MapleXmpl.TransactOpts, owner_, spender_, amount_, deadline_, v_, r_, s_)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner_, address spender_, uint256 amount_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_MapleXmpl *MapleXmplTransactorSession) Permit(owner_ common.Address, spender_ common.Address, amount_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Permit(&_MapleXmpl.TransactOpts, owner_, spender_, amount_, deadline_, v_, r_, s_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplTransactor) Redeem(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "redeem", shares_, receiver_, owner_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplSession) Redeem(shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Redeem(&_MapleXmpl.TransactOpts, shares_, receiver_, owner_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256 assets_)
func (_MapleXmpl *MapleXmplTransactorSession) Redeem(shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Redeem(&_MapleXmpl.TransactOpts, shares_, receiver_, owner_)
}

// ScheduleMigration is a paid mutator transaction binding the contract method 0xab0b0fd3.
//
// Solidity: function scheduleMigration(address migrator_, address newAsset_) returns()
func (_MapleXmpl *MapleXmplTransactor) ScheduleMigration(opts *bind.TransactOpts, migrator_ common.Address, newAsset_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "scheduleMigration", migrator_, newAsset_)
}

// ScheduleMigration is a paid mutator transaction binding the contract method 0xab0b0fd3.
//
// Solidity: function scheduleMigration(address migrator_, address newAsset_) returns()
func (_MapleXmpl *MapleXmplSession) ScheduleMigration(migrator_ common.Address, newAsset_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.ScheduleMigration(&_MapleXmpl.TransactOpts, migrator_, newAsset_)
}

// ScheduleMigration is a paid mutator transaction binding the contract method 0xab0b0fd3.
//
// Solidity: function scheduleMigration(address migrator_, address newAsset_) returns()
func (_MapleXmpl *MapleXmplTransactorSession) ScheduleMigration(migrator_ common.Address, newAsset_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.ScheduleMigration(&_MapleXmpl.TransactOpts, migrator_, newAsset_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_MapleXmpl *MapleXmplTransactor) SetPendingOwner(opts *bind.TransactOpts, pendingOwner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "setPendingOwner", pendingOwner_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_MapleXmpl *MapleXmplSession) SetPendingOwner(pendingOwner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.SetPendingOwner(&_MapleXmpl.TransactOpts, pendingOwner_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_MapleXmpl *MapleXmplTransactorSession) SetPendingOwner(pendingOwner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.SetPendingOwner(&_MapleXmpl.TransactOpts, pendingOwner_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactor) Transfer(opts *bind.TransactOpts, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "transfer", recipient_, amount_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplSession) Transfer(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Transfer(&_MapleXmpl.TransactOpts, recipient_, amount_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactorSession) Transfer(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Transfer(&_MapleXmpl.TransactOpts, recipient_, amount_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner_, address recipient_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactor) TransferFrom(opts *bind.TransactOpts, owner_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "transferFrom", owner_, recipient_, amount_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner_, address recipient_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplSession) TransferFrom(owner_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.TransferFrom(&_MapleXmpl.TransactOpts, owner_, recipient_, amount_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner_, address recipient_, uint256 amount_) returns(bool success_)
func (_MapleXmpl *MapleXmplTransactorSession) TransferFrom(owner_ common.Address, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.TransferFrom(&_MapleXmpl.TransactOpts, owner_, recipient_, amount_)
}

// UpdateVestingSchedule is a paid mutator transaction binding the contract method 0xe13aa990.
//
// Solidity: function updateVestingSchedule(uint256 vestingPeriod_) returns(uint256 issuanceRate_, uint256 freeAssets_)
func (_MapleXmpl *MapleXmplTransactor) UpdateVestingSchedule(opts *bind.TransactOpts, vestingPeriod_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "updateVestingSchedule", vestingPeriod_)
}

// UpdateVestingSchedule is a paid mutator transaction binding the contract method 0xe13aa990.
//
// Solidity: function updateVestingSchedule(uint256 vestingPeriod_) returns(uint256 issuanceRate_, uint256 freeAssets_)
func (_MapleXmpl *MapleXmplSession) UpdateVestingSchedule(vestingPeriod_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.UpdateVestingSchedule(&_MapleXmpl.TransactOpts, vestingPeriod_)
}

// UpdateVestingSchedule is a paid mutator transaction binding the contract method 0xe13aa990.
//
// Solidity: function updateVestingSchedule(uint256 vestingPeriod_) returns(uint256 issuanceRate_, uint256 freeAssets_)
func (_MapleXmpl *MapleXmplTransactorSession) UpdateVestingSchedule(vestingPeriod_ *big.Int) (*types.Transaction, error) {
	return _MapleXmpl.Contract.UpdateVestingSchedule(&_MapleXmpl.TransactOpts, vestingPeriod_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplTransactor) Withdraw(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.contract.Transact(opts, "withdraw", assets_, receiver_, owner_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplSession) Withdraw(assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Withdraw(&_MapleXmpl.TransactOpts, assets_, receiver_, owner_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256 shares_)
func (_MapleXmpl *MapleXmplTransactorSession) Withdraw(assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _MapleXmpl.Contract.Withdraw(&_MapleXmpl.TransactOpts, assets_, receiver_, owner_)
}

// MapleXmplApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MapleXmpl contract.
type MapleXmplApprovalIterator struct {
	Event *MapleXmplApproval // Event containing the contract specifics and raw log

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
func (it *MapleXmplApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplApproval)
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
		it.Event = new(MapleXmplApproval)
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
func (it *MapleXmplApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplApproval represents a Approval event raised by the MapleXmpl contract.
type MapleXmplApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner_, address indexed spender_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) FilterApproval(opts *bind.FilterOpts, owner_ []common.Address, spender_ []common.Address) (*MapleXmplApprovalIterator, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}
	var spender_Rule []interface{}
	for _, spender_Item := range spender_ {
		spender_Rule = append(spender_Rule, spender_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "Approval", owner_Rule, spender_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplApprovalIterator{contract: _MapleXmpl.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner_, address indexed spender_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MapleXmplApproval, owner_ []common.Address, spender_ []common.Address) (event.Subscription, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}
	var spender_Rule []interface{}
	for _, spender_Item := range spender_ {
		spender_Rule = append(spender_Rule, spender_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "Approval", owner_Rule, spender_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplApproval)
				if err := _MapleXmpl.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner_, address indexed spender_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) ParseApproval(log types.Log) (*MapleXmplApproval, error) {
	event := new(MapleXmplApproval)
	if err := _MapleXmpl.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MapleXmpl contract.
type MapleXmplDepositIterator struct {
	Event *MapleXmplDeposit // Event containing the contract specifics and raw log

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
func (it *MapleXmplDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplDeposit)
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
		it.Event = new(MapleXmplDeposit)
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
func (it *MapleXmplDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplDeposit represents a Deposit event raised by the MapleXmpl contract.
type MapleXmplDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller_, address indexed owner_, uint256 assets_, uint256 shares_)
func (_MapleXmpl *MapleXmplFilterer) FilterDeposit(opts *bind.FilterOpts, caller_ []common.Address, owner_ []common.Address) (*MapleXmplDepositIterator, error) {

	var caller_Rule []interface{}
	for _, caller_Item := range caller_ {
		caller_Rule = append(caller_Rule, caller_Item)
	}
	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "Deposit", caller_Rule, owner_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplDepositIterator{contract: _MapleXmpl.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller_, address indexed owner_, uint256 assets_, uint256 shares_)
func (_MapleXmpl *MapleXmplFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MapleXmplDeposit, caller_ []common.Address, owner_ []common.Address) (event.Subscription, error) {

	var caller_Rule []interface{}
	for _, caller_Item := range caller_ {
		caller_Rule = append(caller_Rule, caller_Item)
	}
	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "Deposit", caller_Rule, owner_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplDeposit)
				if err := _MapleXmpl.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller_, address indexed owner_, uint256 assets_, uint256 shares_)
func (_MapleXmpl *MapleXmplFilterer) ParseDeposit(log types.Log) (*MapleXmplDeposit, error) {
	event := new(MapleXmplDeposit)
	if err := _MapleXmpl.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplIssuanceParamsUpdatedIterator is returned from FilterIssuanceParamsUpdated and is used to iterate over the raw logs and unpacked data for IssuanceParamsUpdated events raised by the MapleXmpl contract.
type MapleXmplIssuanceParamsUpdatedIterator struct {
	Event *MapleXmplIssuanceParamsUpdated // Event containing the contract specifics and raw log

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
func (it *MapleXmplIssuanceParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplIssuanceParamsUpdated)
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
		it.Event = new(MapleXmplIssuanceParamsUpdated)
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
func (it *MapleXmplIssuanceParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplIssuanceParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplIssuanceParamsUpdated represents a IssuanceParamsUpdated event raised by the MapleXmpl contract.
type MapleXmplIssuanceParamsUpdated struct {
	FreeAssets   *big.Int
	IssuanceRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIssuanceParamsUpdated is a free log retrieval operation binding the contract event 0x68b521a89bf844ff03e484d89fd64ce292a698ec522170f0dad7ecd11c2dc8fa.
//
// Solidity: event IssuanceParamsUpdated(uint256 freeAssets_, uint256 issuanceRate_)
func (_MapleXmpl *MapleXmplFilterer) FilterIssuanceParamsUpdated(opts *bind.FilterOpts) (*MapleXmplIssuanceParamsUpdatedIterator, error) {

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "IssuanceParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &MapleXmplIssuanceParamsUpdatedIterator{contract: _MapleXmpl.contract, event: "IssuanceParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchIssuanceParamsUpdated is a free log subscription operation binding the contract event 0x68b521a89bf844ff03e484d89fd64ce292a698ec522170f0dad7ecd11c2dc8fa.
//
// Solidity: event IssuanceParamsUpdated(uint256 freeAssets_, uint256 issuanceRate_)
func (_MapleXmpl *MapleXmplFilterer) WatchIssuanceParamsUpdated(opts *bind.WatchOpts, sink chan<- *MapleXmplIssuanceParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "IssuanceParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplIssuanceParamsUpdated)
				if err := _MapleXmpl.contract.UnpackLog(event, "IssuanceParamsUpdated", log); err != nil {
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

// ParseIssuanceParamsUpdated is a log parse operation binding the contract event 0x68b521a89bf844ff03e484d89fd64ce292a698ec522170f0dad7ecd11c2dc8fa.
//
// Solidity: event IssuanceParamsUpdated(uint256 freeAssets_, uint256 issuanceRate_)
func (_MapleXmpl *MapleXmplFilterer) ParseIssuanceParamsUpdated(log types.Log) (*MapleXmplIssuanceParamsUpdated, error) {
	event := new(MapleXmplIssuanceParamsUpdated)
	if err := _MapleXmpl.contract.UnpackLog(event, "IssuanceParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplMigrationCancelledIterator is returned from FilterMigrationCancelled and is used to iterate over the raw logs and unpacked data for MigrationCancelled events raised by the MapleXmpl contract.
type MapleXmplMigrationCancelledIterator struct {
	Event *MapleXmplMigrationCancelled // Event containing the contract specifics and raw log

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
func (it *MapleXmplMigrationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplMigrationCancelled)
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
		it.Event = new(MapleXmplMigrationCancelled)
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
func (it *MapleXmplMigrationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplMigrationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplMigrationCancelled represents a MigrationCancelled event raised by the MapleXmpl contract.
type MapleXmplMigrationCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMigrationCancelled is a free log retrieval operation binding the contract event 0x29183829efed7c15f37ba982c36934966b558232f770f334108ac30e4301333a.
//
// Solidity: event MigrationCancelled()
func (_MapleXmpl *MapleXmplFilterer) FilterMigrationCancelled(opts *bind.FilterOpts) (*MapleXmplMigrationCancelledIterator, error) {

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "MigrationCancelled")
	if err != nil {
		return nil, err
	}
	return &MapleXmplMigrationCancelledIterator{contract: _MapleXmpl.contract, event: "MigrationCancelled", logs: logs, sub: sub}, nil
}

// WatchMigrationCancelled is a free log subscription operation binding the contract event 0x29183829efed7c15f37ba982c36934966b558232f770f334108ac30e4301333a.
//
// Solidity: event MigrationCancelled()
func (_MapleXmpl *MapleXmplFilterer) WatchMigrationCancelled(opts *bind.WatchOpts, sink chan<- *MapleXmplMigrationCancelled) (event.Subscription, error) {

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "MigrationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplMigrationCancelled)
				if err := _MapleXmpl.contract.UnpackLog(event, "MigrationCancelled", log); err != nil {
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

// ParseMigrationCancelled is a log parse operation binding the contract event 0x29183829efed7c15f37ba982c36934966b558232f770f334108ac30e4301333a.
//
// Solidity: event MigrationCancelled()
func (_MapleXmpl *MapleXmplFilterer) ParseMigrationCancelled(log types.Log) (*MapleXmplMigrationCancelled, error) {
	event := new(MapleXmplMigrationCancelled)
	if err := _MapleXmpl.contract.UnpackLog(event, "MigrationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplMigrationPerformedIterator is returned from FilterMigrationPerformed and is used to iterate over the raw logs and unpacked data for MigrationPerformed events raised by the MapleXmpl contract.
type MapleXmplMigrationPerformedIterator struct {
	Event *MapleXmplMigrationPerformed // Event containing the contract specifics and raw log

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
func (it *MapleXmplMigrationPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplMigrationPerformed)
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
		it.Event = new(MapleXmplMigrationPerformed)
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
func (it *MapleXmplMigrationPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplMigrationPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplMigrationPerformed represents a MigrationPerformed event raised by the MapleXmpl contract.
type MapleXmplMigrationPerformed struct {
	FromAsset common.Address
	ToAsset   common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMigrationPerformed is a free log retrieval operation binding the contract event 0xd2b4bfa00cffc450ec3f60db3c511bebc1e88911c9c7b25ef8b294c934fa116c.
//
// Solidity: event MigrationPerformed(address indexed fromAsset_, address indexed toAsset_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) FilterMigrationPerformed(opts *bind.FilterOpts, fromAsset_ []common.Address, toAsset_ []common.Address) (*MapleXmplMigrationPerformedIterator, error) {

	var fromAsset_Rule []interface{}
	for _, fromAsset_Item := range fromAsset_ {
		fromAsset_Rule = append(fromAsset_Rule, fromAsset_Item)
	}
	var toAsset_Rule []interface{}
	for _, toAsset_Item := range toAsset_ {
		toAsset_Rule = append(toAsset_Rule, toAsset_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "MigrationPerformed", fromAsset_Rule, toAsset_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplMigrationPerformedIterator{contract: _MapleXmpl.contract, event: "MigrationPerformed", logs: logs, sub: sub}, nil
}

// WatchMigrationPerformed is a free log subscription operation binding the contract event 0xd2b4bfa00cffc450ec3f60db3c511bebc1e88911c9c7b25ef8b294c934fa116c.
//
// Solidity: event MigrationPerformed(address indexed fromAsset_, address indexed toAsset_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) WatchMigrationPerformed(opts *bind.WatchOpts, sink chan<- *MapleXmplMigrationPerformed, fromAsset_ []common.Address, toAsset_ []common.Address) (event.Subscription, error) {

	var fromAsset_Rule []interface{}
	for _, fromAsset_Item := range fromAsset_ {
		fromAsset_Rule = append(fromAsset_Rule, fromAsset_Item)
	}
	var toAsset_Rule []interface{}
	for _, toAsset_Item := range toAsset_ {
		toAsset_Rule = append(toAsset_Rule, toAsset_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "MigrationPerformed", fromAsset_Rule, toAsset_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplMigrationPerformed)
				if err := _MapleXmpl.contract.UnpackLog(event, "MigrationPerformed", log); err != nil {
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

// ParseMigrationPerformed is a log parse operation binding the contract event 0xd2b4bfa00cffc450ec3f60db3c511bebc1e88911c9c7b25ef8b294c934fa116c.
//
// Solidity: event MigrationPerformed(address indexed fromAsset_, address indexed toAsset_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) ParseMigrationPerformed(log types.Log) (*MapleXmplMigrationPerformed, error) {
	event := new(MapleXmplMigrationPerformed)
	if err := _MapleXmpl.contract.UnpackLog(event, "MigrationPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplMigrationScheduledIterator is returned from FilterMigrationScheduled and is used to iterate over the raw logs and unpacked data for MigrationScheduled events raised by the MapleXmpl contract.
type MapleXmplMigrationScheduledIterator struct {
	Event *MapleXmplMigrationScheduled // Event containing the contract specifics and raw log

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
func (it *MapleXmplMigrationScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplMigrationScheduled)
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
		it.Event = new(MapleXmplMigrationScheduled)
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
func (it *MapleXmplMigrationScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplMigrationScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplMigrationScheduled represents a MigrationScheduled event raised by the MapleXmpl contract.
type MapleXmplMigrationScheduled struct {
	FromAsset     common.Address
	ToAsset       common.Address
	Migrator      common.Address
	MigrationTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMigrationScheduled is a free log retrieval operation binding the contract event 0xfc384cf40fadcb79796c4f6b26dab65f122020fbffaa19c2ad28bf4da410cded.
//
// Solidity: event MigrationScheduled(address indexed fromAsset_, address indexed toAsset_, address indexed migrator_, uint256 migrationTime_)
func (_MapleXmpl *MapleXmplFilterer) FilterMigrationScheduled(opts *bind.FilterOpts, fromAsset_ []common.Address, toAsset_ []common.Address, migrator_ []common.Address) (*MapleXmplMigrationScheduledIterator, error) {

	var fromAsset_Rule []interface{}
	for _, fromAsset_Item := range fromAsset_ {
		fromAsset_Rule = append(fromAsset_Rule, fromAsset_Item)
	}
	var toAsset_Rule []interface{}
	for _, toAsset_Item := range toAsset_ {
		toAsset_Rule = append(toAsset_Rule, toAsset_Item)
	}
	var migrator_Rule []interface{}
	for _, migrator_Item := range migrator_ {
		migrator_Rule = append(migrator_Rule, migrator_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "MigrationScheduled", fromAsset_Rule, toAsset_Rule, migrator_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplMigrationScheduledIterator{contract: _MapleXmpl.contract, event: "MigrationScheduled", logs: logs, sub: sub}, nil
}

// WatchMigrationScheduled is a free log subscription operation binding the contract event 0xfc384cf40fadcb79796c4f6b26dab65f122020fbffaa19c2ad28bf4da410cded.
//
// Solidity: event MigrationScheduled(address indexed fromAsset_, address indexed toAsset_, address indexed migrator_, uint256 migrationTime_)
func (_MapleXmpl *MapleXmplFilterer) WatchMigrationScheduled(opts *bind.WatchOpts, sink chan<- *MapleXmplMigrationScheduled, fromAsset_ []common.Address, toAsset_ []common.Address, migrator_ []common.Address) (event.Subscription, error) {

	var fromAsset_Rule []interface{}
	for _, fromAsset_Item := range fromAsset_ {
		fromAsset_Rule = append(fromAsset_Rule, fromAsset_Item)
	}
	var toAsset_Rule []interface{}
	for _, toAsset_Item := range toAsset_ {
		toAsset_Rule = append(toAsset_Rule, toAsset_Item)
	}
	var migrator_Rule []interface{}
	for _, migrator_Item := range migrator_ {
		migrator_Rule = append(migrator_Rule, migrator_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "MigrationScheduled", fromAsset_Rule, toAsset_Rule, migrator_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplMigrationScheduled)
				if err := _MapleXmpl.contract.UnpackLog(event, "MigrationScheduled", log); err != nil {
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

// ParseMigrationScheduled is a log parse operation binding the contract event 0xfc384cf40fadcb79796c4f6b26dab65f122020fbffaa19c2ad28bf4da410cded.
//
// Solidity: event MigrationScheduled(address indexed fromAsset_, address indexed toAsset_, address indexed migrator_, uint256 migrationTime_)
func (_MapleXmpl *MapleXmplFilterer) ParseMigrationScheduled(log types.Log) (*MapleXmplMigrationScheduled, error) {
	event := new(MapleXmplMigrationScheduled)
	if err := _MapleXmpl.contract.UnpackLog(event, "MigrationScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplOwnershipAcceptedIterator is returned from FilterOwnershipAccepted and is used to iterate over the raw logs and unpacked data for OwnershipAccepted events raised by the MapleXmpl contract.
type MapleXmplOwnershipAcceptedIterator struct {
	Event *MapleXmplOwnershipAccepted // Event containing the contract specifics and raw log

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
func (it *MapleXmplOwnershipAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplOwnershipAccepted)
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
		it.Event = new(MapleXmplOwnershipAccepted)
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
func (it *MapleXmplOwnershipAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplOwnershipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplOwnershipAccepted represents a OwnershipAccepted event raised by the MapleXmpl contract.
type MapleXmplOwnershipAccepted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipAccepted is a free log retrieval operation binding the contract event 0x357bdeb5828fa83945f38a88510ce5cd7d628dafb346d767efbc693149fdd97c.
//
// Solidity: event OwnershipAccepted(address indexed previousOwner_, address indexed newOwner_)
func (_MapleXmpl *MapleXmplFilterer) FilterOwnershipAccepted(opts *bind.FilterOpts, previousOwner_ []common.Address, newOwner_ []common.Address) (*MapleXmplOwnershipAcceptedIterator, error) {

	var previousOwner_Rule []interface{}
	for _, previousOwner_Item := range previousOwner_ {
		previousOwner_Rule = append(previousOwner_Rule, previousOwner_Item)
	}
	var newOwner_Rule []interface{}
	for _, newOwner_Item := range newOwner_ {
		newOwner_Rule = append(newOwner_Rule, newOwner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "OwnershipAccepted", previousOwner_Rule, newOwner_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplOwnershipAcceptedIterator{contract: _MapleXmpl.contract, event: "OwnershipAccepted", logs: logs, sub: sub}, nil
}

// WatchOwnershipAccepted is a free log subscription operation binding the contract event 0x357bdeb5828fa83945f38a88510ce5cd7d628dafb346d767efbc693149fdd97c.
//
// Solidity: event OwnershipAccepted(address indexed previousOwner_, address indexed newOwner_)
func (_MapleXmpl *MapleXmplFilterer) WatchOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *MapleXmplOwnershipAccepted, previousOwner_ []common.Address, newOwner_ []common.Address) (event.Subscription, error) {

	var previousOwner_Rule []interface{}
	for _, previousOwner_Item := range previousOwner_ {
		previousOwner_Rule = append(previousOwner_Rule, previousOwner_Item)
	}
	var newOwner_Rule []interface{}
	for _, newOwner_Item := range newOwner_ {
		newOwner_Rule = append(newOwner_Rule, newOwner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "OwnershipAccepted", previousOwner_Rule, newOwner_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplOwnershipAccepted)
				if err := _MapleXmpl.contract.UnpackLog(event, "OwnershipAccepted", log); err != nil {
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

// ParseOwnershipAccepted is a log parse operation binding the contract event 0x357bdeb5828fa83945f38a88510ce5cd7d628dafb346d767efbc693149fdd97c.
//
// Solidity: event OwnershipAccepted(address indexed previousOwner_, address indexed newOwner_)
func (_MapleXmpl *MapleXmplFilterer) ParseOwnershipAccepted(log types.Log) (*MapleXmplOwnershipAccepted, error) {
	event := new(MapleXmplOwnershipAccepted)
	if err := _MapleXmpl.contract.UnpackLog(event, "OwnershipAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplPendingOwnerSetIterator is returned from FilterPendingOwnerSet and is used to iterate over the raw logs and unpacked data for PendingOwnerSet events raised by the MapleXmpl contract.
type MapleXmplPendingOwnerSetIterator struct {
	Event *MapleXmplPendingOwnerSet // Event containing the contract specifics and raw log

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
func (it *MapleXmplPendingOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplPendingOwnerSet)
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
		it.Event = new(MapleXmplPendingOwnerSet)
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
func (it *MapleXmplPendingOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplPendingOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplPendingOwnerSet represents a PendingOwnerSet event raised by the MapleXmpl contract.
type MapleXmplPendingOwnerSet struct {
	Owner        common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingOwnerSet is a free log retrieval operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed owner_, address indexed pendingOwner_)
func (_MapleXmpl *MapleXmplFilterer) FilterPendingOwnerSet(opts *bind.FilterOpts, owner_ []common.Address, pendingOwner_ []common.Address) (*MapleXmplPendingOwnerSetIterator, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}
	var pendingOwner_Rule []interface{}
	for _, pendingOwner_Item := range pendingOwner_ {
		pendingOwner_Rule = append(pendingOwner_Rule, pendingOwner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "PendingOwnerSet", owner_Rule, pendingOwner_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplPendingOwnerSetIterator{contract: _MapleXmpl.contract, event: "PendingOwnerSet", logs: logs, sub: sub}, nil
}

// WatchPendingOwnerSet is a free log subscription operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed owner_, address indexed pendingOwner_)
func (_MapleXmpl *MapleXmplFilterer) WatchPendingOwnerSet(opts *bind.WatchOpts, sink chan<- *MapleXmplPendingOwnerSet, owner_ []common.Address, pendingOwner_ []common.Address) (event.Subscription, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}
	var pendingOwner_Rule []interface{}
	for _, pendingOwner_Item := range pendingOwner_ {
		pendingOwner_Rule = append(pendingOwner_Rule, pendingOwner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "PendingOwnerSet", owner_Rule, pendingOwner_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplPendingOwnerSet)
				if err := _MapleXmpl.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
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

// ParsePendingOwnerSet is a log parse operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed owner_, address indexed pendingOwner_)
func (_MapleXmpl *MapleXmplFilterer) ParsePendingOwnerSet(log types.Log) (*MapleXmplPendingOwnerSet, error) {
	event := new(MapleXmplPendingOwnerSet)
	if err := _MapleXmpl.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MapleXmpl contract.
type MapleXmplTransferIterator struct {
	Event *MapleXmplTransfer // Event containing the contract specifics and raw log

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
func (it *MapleXmplTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplTransfer)
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
		it.Event = new(MapleXmplTransfer)
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
func (it *MapleXmplTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplTransfer represents a Transfer event raised by the MapleXmpl contract.
type MapleXmplTransfer struct {
	Owner     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed owner_, address indexed recipient_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) FilterTransfer(opts *bind.FilterOpts, owner_ []common.Address, recipient_ []common.Address) (*MapleXmplTransferIterator, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}
	var recipient_Rule []interface{}
	for _, recipient_Item := range recipient_ {
		recipient_Rule = append(recipient_Rule, recipient_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "Transfer", owner_Rule, recipient_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplTransferIterator{contract: _MapleXmpl.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed owner_, address indexed recipient_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MapleXmplTransfer, owner_ []common.Address, recipient_ []common.Address) (event.Subscription, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}
	var recipient_Rule []interface{}
	for _, recipient_Item := range recipient_ {
		recipient_Rule = append(recipient_Rule, recipient_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "Transfer", owner_Rule, recipient_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplTransfer)
				if err := _MapleXmpl.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed owner_, address indexed recipient_, uint256 amount_)
func (_MapleXmpl *MapleXmplFilterer) ParseTransfer(log types.Log) (*MapleXmplTransfer, error) {
	event := new(MapleXmplTransfer)
	if err := _MapleXmpl.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplVestingScheduleUpdatedIterator is returned from FilterVestingScheduleUpdated and is used to iterate over the raw logs and unpacked data for VestingScheduleUpdated events raised by the MapleXmpl contract.
type MapleXmplVestingScheduleUpdatedIterator struct {
	Event *MapleXmplVestingScheduleUpdated // Event containing the contract specifics and raw log

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
func (it *MapleXmplVestingScheduleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplVestingScheduleUpdated)
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
		it.Event = new(MapleXmplVestingScheduleUpdated)
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
func (it *MapleXmplVestingScheduleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplVestingScheduleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplVestingScheduleUpdated represents a VestingScheduleUpdated event raised by the MapleXmpl contract.
type MapleXmplVestingScheduleUpdated struct {
	Owner               common.Address
	VestingPeriodFinish *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleUpdated is a free log retrieval operation binding the contract event 0x8c84e3b4df93f5b7c8d4ab6647708f5b14cacc124e22908187e30695ec54bab3.
//
// Solidity: event VestingScheduleUpdated(address indexed owner_, uint256 vestingPeriodFinish_)
func (_MapleXmpl *MapleXmplFilterer) FilterVestingScheduleUpdated(opts *bind.FilterOpts, owner_ []common.Address) (*MapleXmplVestingScheduleUpdatedIterator, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "VestingScheduleUpdated", owner_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplVestingScheduleUpdatedIterator{contract: _MapleXmpl.contract, event: "VestingScheduleUpdated", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleUpdated is a free log subscription operation binding the contract event 0x8c84e3b4df93f5b7c8d4ab6647708f5b14cacc124e22908187e30695ec54bab3.
//
// Solidity: event VestingScheduleUpdated(address indexed owner_, uint256 vestingPeriodFinish_)
func (_MapleXmpl *MapleXmplFilterer) WatchVestingScheduleUpdated(opts *bind.WatchOpts, sink chan<- *MapleXmplVestingScheduleUpdated, owner_ []common.Address) (event.Subscription, error) {

	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "VestingScheduleUpdated", owner_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplVestingScheduleUpdated)
				if err := _MapleXmpl.contract.UnpackLog(event, "VestingScheduleUpdated", log); err != nil {
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

// ParseVestingScheduleUpdated is a log parse operation binding the contract event 0x8c84e3b4df93f5b7c8d4ab6647708f5b14cacc124e22908187e30695ec54bab3.
//
// Solidity: event VestingScheduleUpdated(address indexed owner_, uint256 vestingPeriodFinish_)
func (_MapleXmpl *MapleXmplFilterer) ParseVestingScheduleUpdated(log types.Log) (*MapleXmplVestingScheduleUpdated, error) {
	event := new(MapleXmplVestingScheduleUpdated)
	if err := _MapleXmpl.contract.UnpackLog(event, "VestingScheduleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleXmplWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MapleXmpl contract.
type MapleXmplWithdrawIterator struct {
	Event *MapleXmplWithdraw // Event containing the contract specifics and raw log

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
func (it *MapleXmplWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleXmplWithdraw)
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
		it.Event = new(MapleXmplWithdraw)
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
func (it *MapleXmplWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleXmplWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleXmplWithdraw represents a Withdraw event raised by the MapleXmpl contract.
type MapleXmplWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller_, address indexed receiver_, address indexed owner_, uint256 assets_, uint256 shares_)
func (_MapleXmpl *MapleXmplFilterer) FilterWithdraw(opts *bind.FilterOpts, caller_ []common.Address, receiver_ []common.Address, owner_ []common.Address) (*MapleXmplWithdrawIterator, error) {

	var caller_Rule []interface{}
	for _, caller_Item := range caller_ {
		caller_Rule = append(caller_Rule, caller_Item)
	}
	var receiver_Rule []interface{}
	for _, receiver_Item := range receiver_ {
		receiver_Rule = append(receiver_Rule, receiver_Item)
	}
	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.FilterLogs(opts, "Withdraw", caller_Rule, receiver_Rule, owner_Rule)
	if err != nil {
		return nil, err
	}
	return &MapleXmplWithdrawIterator{contract: _MapleXmpl.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller_, address indexed receiver_, address indexed owner_, uint256 assets_, uint256 shares_)
func (_MapleXmpl *MapleXmplFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MapleXmplWithdraw, caller_ []common.Address, receiver_ []common.Address, owner_ []common.Address) (event.Subscription, error) {

	var caller_Rule []interface{}
	for _, caller_Item := range caller_ {
		caller_Rule = append(caller_Rule, caller_Item)
	}
	var receiver_Rule []interface{}
	for _, receiver_Item := range receiver_ {
		receiver_Rule = append(receiver_Rule, receiver_Item)
	}
	var owner_Rule []interface{}
	for _, owner_Item := range owner_ {
		owner_Rule = append(owner_Rule, owner_Item)
	}

	logs, sub, err := _MapleXmpl.contract.WatchLogs(opts, "Withdraw", caller_Rule, receiver_Rule, owner_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleXmplWithdraw)
				if err := _MapleXmpl.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller_, address indexed receiver_, address indexed owner_, uint256 assets_, uint256 shares_)
func (_MapleXmpl *MapleXmplFilterer) ParseWithdraw(log types.Log) (*MapleXmplWithdraw, error) {
	event := new(MapleXmplWithdraw)
	if err := _MapleXmpl.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
