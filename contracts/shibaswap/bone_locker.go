// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shibaswap

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

// ShibaswapBoneLockerMetaData contains all meta data concerning the ShibaswapBoneLocker contract.
var ShibaswapBoneLockerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_boneToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockingPeriodInDays\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_devLockingPeriodInDays\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLockingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDevLockingPeriod\",\"type\":\"uint256\"}],\"name\":\"LockingPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"claimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimAllForUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devLockingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"emergencyWithdrawOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getLeftRightCounters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestCounterByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isDev\",\"type\":\"bool\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockInfoByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isDev\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAddr\",\"type\":\"address\"}],\"name\":\"setEmergencyAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_emergencyFlag\",\"type\":\"bool\"}],\"name\":\"setEmergencyFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLockingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newDevLockingPeriod\",\"type\":\"uint256\"}],\"name\":\"setLockingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unclaimedTokensByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ShibaswapBoneLockerABI is the input ABI used to generate the binding from.
// Deprecated: Use ShibaswapBoneLockerMetaData.ABI instead.
var ShibaswapBoneLockerABI = ShibaswapBoneLockerMetaData.ABI

// ShibaswapBoneLocker is an auto generated Go binding around an Ethereum contract.
type ShibaswapBoneLocker struct {
	ShibaswapBoneLockerCaller     // Read-only binding to the contract
	ShibaswapBoneLockerTransactor // Write-only binding to the contract
	ShibaswapBoneLockerFilterer   // Log filterer for contract events
}

// ShibaswapBoneLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShibaswapBoneLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibaswapBoneLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShibaswapBoneLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibaswapBoneLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShibaswapBoneLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibaswapBoneLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShibaswapBoneLockerSession struct {
	Contract     *ShibaswapBoneLocker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ShibaswapBoneLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShibaswapBoneLockerCallerSession struct {
	Contract *ShibaswapBoneLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ShibaswapBoneLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShibaswapBoneLockerTransactorSession struct {
	Contract     *ShibaswapBoneLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ShibaswapBoneLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShibaswapBoneLockerRaw struct {
	Contract *ShibaswapBoneLocker // Generic contract binding to access the raw methods on
}

// ShibaswapBoneLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShibaswapBoneLockerCallerRaw struct {
	Contract *ShibaswapBoneLockerCaller // Generic read-only contract binding to access the raw methods on
}

// ShibaswapBoneLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShibaswapBoneLockerTransactorRaw struct {
	Contract *ShibaswapBoneLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShibaswapBoneLocker creates a new instance of ShibaswapBoneLocker, bound to a specific deployed contract.
func NewShibaswapBoneLocker(address common.Address, backend bind.ContractBackend) (*ShibaswapBoneLocker, error) {
	contract, err := bindShibaswapBoneLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShibaswapBoneLocker{ShibaswapBoneLockerCaller: ShibaswapBoneLockerCaller{contract: contract}, ShibaswapBoneLockerTransactor: ShibaswapBoneLockerTransactor{contract: contract}, ShibaswapBoneLockerFilterer: ShibaswapBoneLockerFilterer{contract: contract}}, nil
}

// NewShibaswapBoneLockerCaller creates a new read-only instance of ShibaswapBoneLocker, bound to a specific deployed contract.
func NewShibaswapBoneLockerCaller(address common.Address, caller bind.ContractCaller) (*ShibaswapBoneLockerCaller, error) {
	contract, err := bindShibaswapBoneLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShibaswapBoneLockerCaller{contract: contract}, nil
}

// NewShibaswapBoneLockerTransactor creates a new write-only instance of ShibaswapBoneLocker, bound to a specific deployed contract.
func NewShibaswapBoneLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*ShibaswapBoneLockerTransactor, error) {
	contract, err := bindShibaswapBoneLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShibaswapBoneLockerTransactor{contract: contract}, nil
}

// NewShibaswapBoneLockerFilterer creates a new log filterer instance of ShibaswapBoneLocker, bound to a specific deployed contract.
func NewShibaswapBoneLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*ShibaswapBoneLockerFilterer, error) {
	contract, err := bindShibaswapBoneLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShibaswapBoneLockerFilterer{contract: contract}, nil
}

// bindShibaswapBoneLocker binds a generic wrapper to an already deployed contract.
func bindShibaswapBoneLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShibaswapBoneLockerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShibaswapBoneLocker *ShibaswapBoneLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShibaswapBoneLocker.Contract.ShibaswapBoneLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShibaswapBoneLocker *ShibaswapBoneLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.ShibaswapBoneLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShibaswapBoneLocker *ShibaswapBoneLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.ShibaswapBoneLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShibaswapBoneLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.contract.Transact(opts, method, params...)
}

// DevLockingPeriod is a free data retrieval call binding the contract method 0xbcf78da4.
//
// Solidity: function devLockingPeriod() view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) DevLockingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "devLockingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevLockingPeriod is a free data retrieval call binding the contract method 0xbcf78da4.
//
// Solidity: function devLockingPeriod() view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) DevLockingPeriod() (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.DevLockingPeriod(&_ShibaswapBoneLocker.CallOpts)
}

// DevLockingPeriod is a free data retrieval call binding the contract method 0xbcf78da4.
//
// Solidity: function devLockingPeriod() view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) DevLockingPeriod() (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.DevLockingPeriod(&_ShibaswapBoneLocker.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address _user) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) GetClaimableAmount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "getClaimableAmount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address _user) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) GetClaimableAmount(_user common.Address) (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.GetClaimableAmount(&_ShibaswapBoneLocker.CallOpts, _user)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address _user) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) GetClaimableAmount(_user common.Address) (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.GetClaimableAmount(&_ShibaswapBoneLocker.CallOpts, _user)
}

// GetLeftRightCounters is a free data retrieval call binding the contract method 0xc3ca2d80.
//
// Solidity: function getLeftRightCounters(address _user) view returns(uint256, uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) GetLeftRightCounters(opts *bind.CallOpts, _user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "getLeftRightCounters", _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLeftRightCounters is a free data retrieval call binding the contract method 0xc3ca2d80.
//
// Solidity: function getLeftRightCounters(address _user) view returns(uint256, uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) GetLeftRightCounters(_user common.Address) (*big.Int, *big.Int, error) {
	return _ShibaswapBoneLocker.Contract.GetLeftRightCounters(&_ShibaswapBoneLocker.CallOpts, _user)
}

// GetLeftRightCounters is a free data retrieval call binding the contract method 0xc3ca2d80.
//
// Solidity: function getLeftRightCounters(address _user) view returns(uint256, uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) GetLeftRightCounters(_user common.Address) (*big.Int, *big.Int, error) {
	return _ShibaswapBoneLocker.Contract.GetLeftRightCounters(&_ShibaswapBoneLocker.CallOpts, _user)
}

// LatestCounterByUser is a free data retrieval call binding the contract method 0x758a23e0.
//
// Solidity: function latestCounterByUser(address ) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) LatestCounterByUser(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "latestCounterByUser", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestCounterByUser is a free data retrieval call binding the contract method 0x758a23e0.
//
// Solidity: function latestCounterByUser(address ) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) LatestCounterByUser(arg0 common.Address) (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.LatestCounterByUser(&_ShibaswapBoneLocker.CallOpts, arg0)
}

// LatestCounterByUser is a free data retrieval call binding the contract method 0x758a23e0.
//
// Solidity: function latestCounterByUser(address ) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) LatestCounterByUser(arg0 common.Address) (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.LatestCounterByUser(&_ShibaswapBoneLocker.CallOpts, arg0)
}

// LockInfoByUser is a free data retrieval call binding the contract method 0x230a3f7f.
//
// Solidity: function lockInfoByUser(address , uint256 ) view returns(uint256 _amount, uint256 _timestamp, bool _isDev)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) LockInfoByUser(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
	IsDev     bool
}, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "lockInfoByUser", arg0, arg1)

	outstruct := new(struct {
		Amount    *big.Int
		Timestamp *big.Int
		IsDev     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsDev = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// LockInfoByUser is a free data retrieval call binding the contract method 0x230a3f7f.
//
// Solidity: function lockInfoByUser(address , uint256 ) view returns(uint256 _amount, uint256 _timestamp, bool _isDev)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) LockInfoByUser(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
	IsDev     bool
}, error) {
	return _ShibaswapBoneLocker.Contract.LockInfoByUser(&_ShibaswapBoneLocker.CallOpts, arg0, arg1)
}

// LockInfoByUser is a free data retrieval call binding the contract method 0x230a3f7f.
//
// Solidity: function lockInfoByUser(address , uint256 ) view returns(uint256 _amount, uint256 _timestamp, bool _isDev)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) LockInfoByUser(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
	IsDev     bool
}, error) {
	return _ShibaswapBoneLocker.Contract.LockInfoByUser(&_ShibaswapBoneLocker.CallOpts, arg0, arg1)
}

// LockingPeriod is a free data retrieval call binding the contract method 0x550066d5.
//
// Solidity: function lockingPeriod() view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) LockingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "lockingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockingPeriod is a free data retrieval call binding the contract method 0x550066d5.
//
// Solidity: function lockingPeriod() view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) LockingPeriod() (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.LockingPeriod(&_ShibaswapBoneLocker.CallOpts)
}

// LockingPeriod is a free data retrieval call binding the contract method 0x550066d5.
//
// Solidity: function lockingPeriod() view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) LockingPeriod() (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.LockingPeriod(&_ShibaswapBoneLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) Owner() (common.Address, error) {
	return _ShibaswapBoneLocker.Contract.Owner(&_ShibaswapBoneLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) Owner() (common.Address, error) {
	return _ShibaswapBoneLocker.Contract.Owner(&_ShibaswapBoneLocker.CallOpts)
}

// UnclaimedTokensByUser is a free data retrieval call binding the contract method 0x750320f1.
//
// Solidity: function unclaimedTokensByUser(address ) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCaller) UnclaimedTokensByUser(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapBoneLocker.contract.Call(opts, &out, "unclaimedTokensByUser", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedTokensByUser is a free data retrieval call binding the contract method 0x750320f1.
//
// Solidity: function unclaimedTokensByUser(address ) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) UnclaimedTokensByUser(arg0 common.Address) (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.UnclaimedTokensByUser(&_ShibaswapBoneLocker.CallOpts, arg0)
}

// UnclaimedTokensByUser is a free data retrieval call binding the contract method 0x750320f1.
//
// Solidity: function unclaimedTokensByUser(address ) view returns(uint256)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerCallerSession) UnclaimedTokensByUser(arg0 common.Address) (*big.Int, error) {
	return _ShibaswapBoneLocker.Contract.UnclaimedTokensByUser(&_ShibaswapBoneLocker.CallOpts, arg0)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x6930fd2a.
//
// Solidity: function claimAll(uint256 r) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) ClaimAll(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "claimAll", r)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x6930fd2a.
//
// Solidity: function claimAll(uint256 r) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) ClaimAll(r *big.Int) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.ClaimAll(&_ShibaswapBoneLocker.TransactOpts, r)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x6930fd2a.
//
// Solidity: function claimAll(uint256 r) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) ClaimAll(r *big.Int) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.ClaimAll(&_ShibaswapBoneLocker.TransactOpts, r)
}

// ClaimAllForUser is a paid mutator transaction binding the contract method 0x4e60ac82.
//
// Solidity: function claimAllForUser(uint256 r, address user) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) ClaimAllForUser(opts *bind.TransactOpts, r *big.Int, user common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "claimAllForUser", r, user)
}

// ClaimAllForUser is a paid mutator transaction binding the contract method 0x4e60ac82.
//
// Solidity: function claimAllForUser(uint256 r, address user) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) ClaimAllForUser(r *big.Int, user common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.ClaimAllForUser(&_ShibaswapBoneLocker.TransactOpts, r, user)
}

// ClaimAllForUser is a paid mutator transaction binding the contract method 0x4e60ac82.
//
// Solidity: function claimAllForUser(uint256 r, address user) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) ClaimAllForUser(r *big.Int, user common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.ClaimAllForUser(&_ShibaswapBoneLocker.TransactOpts, r, user)
}

// EmergencyWithdrawOwner is a paid mutator transaction binding the contract method 0x62bfcb06.
//
// Solidity: function emergencyWithdrawOwner(address _to) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) EmergencyWithdrawOwner(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "emergencyWithdrawOwner", _to)
}

// EmergencyWithdrawOwner is a paid mutator transaction binding the contract method 0x62bfcb06.
//
// Solidity: function emergencyWithdrawOwner(address _to) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) EmergencyWithdrawOwner(_to common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.EmergencyWithdrawOwner(&_ShibaswapBoneLocker.TransactOpts, _to)
}

// EmergencyWithdrawOwner is a paid mutator transaction binding the contract method 0x62bfcb06.
//
// Solidity: function emergencyWithdrawOwner(address _to) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) EmergencyWithdrawOwner(_to common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.EmergencyWithdrawOwner(&_ShibaswapBoneLocker.TransactOpts, _to)
}

// Lock is a paid mutator transaction binding the contract method 0x507bb14d.
//
// Solidity: function lock(address _holder, uint256 _amount, bool _isDev) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) Lock(opts *bind.TransactOpts, _holder common.Address, _amount *big.Int, _isDev bool) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "lock", _holder, _amount, _isDev)
}

// Lock is a paid mutator transaction binding the contract method 0x507bb14d.
//
// Solidity: function lock(address _holder, uint256 _amount, bool _isDev) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) Lock(_holder common.Address, _amount *big.Int, _isDev bool) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.Lock(&_ShibaswapBoneLocker.TransactOpts, _holder, _amount, _isDev)
}

// Lock is a paid mutator transaction binding the contract method 0x507bb14d.
//
// Solidity: function lock(address _holder, uint256 _amount, bool _isDev) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) Lock(_holder common.Address, _amount *big.Int, _isDev bool) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.Lock(&_ShibaswapBoneLocker.TransactOpts, _holder, _amount, _isDev)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.RenounceOwnership(&_ShibaswapBoneLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.RenounceOwnership(&_ShibaswapBoneLocker.TransactOpts)
}

// SetEmergencyAddr is a paid mutator transaction binding the contract method 0xf92139df.
//
// Solidity: function setEmergencyAddr(address _newAddr) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) SetEmergencyAddr(opts *bind.TransactOpts, _newAddr common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "setEmergencyAddr", _newAddr)
}

// SetEmergencyAddr is a paid mutator transaction binding the contract method 0xf92139df.
//
// Solidity: function setEmergencyAddr(address _newAddr) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) SetEmergencyAddr(_newAddr common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.SetEmergencyAddr(&_ShibaswapBoneLocker.TransactOpts, _newAddr)
}

// SetEmergencyAddr is a paid mutator transaction binding the contract method 0xf92139df.
//
// Solidity: function setEmergencyAddr(address _newAddr) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) SetEmergencyAddr(_newAddr common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.SetEmergencyAddr(&_ShibaswapBoneLocker.TransactOpts, _newAddr)
}

// SetEmergencyFlag is a paid mutator transaction binding the contract method 0xc697444b.
//
// Solidity: function setEmergencyFlag(bool _emergencyFlag) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) SetEmergencyFlag(opts *bind.TransactOpts, _emergencyFlag bool) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "setEmergencyFlag", _emergencyFlag)
}

// SetEmergencyFlag is a paid mutator transaction binding the contract method 0xc697444b.
//
// Solidity: function setEmergencyFlag(bool _emergencyFlag) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) SetEmergencyFlag(_emergencyFlag bool) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.SetEmergencyFlag(&_ShibaswapBoneLocker.TransactOpts, _emergencyFlag)
}

// SetEmergencyFlag is a paid mutator transaction binding the contract method 0xc697444b.
//
// Solidity: function setEmergencyFlag(bool _emergencyFlag) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) SetEmergencyFlag(_emergencyFlag bool) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.SetEmergencyFlag(&_ShibaswapBoneLocker.TransactOpts, _emergencyFlag)
}

// SetLockingPeriod is a paid mutator transaction binding the contract method 0xbff783ae.
//
// Solidity: function setLockingPeriod(uint256 _newLockingPeriod, uint256 _newDevLockingPeriod) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) SetLockingPeriod(opts *bind.TransactOpts, _newLockingPeriod *big.Int, _newDevLockingPeriod *big.Int) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "setLockingPeriod", _newLockingPeriod, _newDevLockingPeriod)
}

// SetLockingPeriod is a paid mutator transaction binding the contract method 0xbff783ae.
//
// Solidity: function setLockingPeriod(uint256 _newLockingPeriod, uint256 _newDevLockingPeriod) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) SetLockingPeriod(_newLockingPeriod *big.Int, _newDevLockingPeriod *big.Int) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.SetLockingPeriod(&_ShibaswapBoneLocker.TransactOpts, _newLockingPeriod, _newDevLockingPeriod)
}

// SetLockingPeriod is a paid mutator transaction binding the contract method 0xbff783ae.
//
// Solidity: function setLockingPeriod(uint256 _newLockingPeriod, uint256 _newDevLockingPeriod) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) SetLockingPeriod(_newLockingPeriod *big.Int, _newDevLockingPeriod *big.Int) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.SetLockingPeriod(&_ShibaswapBoneLocker.TransactOpts, _newLockingPeriod, _newDevLockingPeriod)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.TransferOwnership(&_ShibaswapBoneLocker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibaswapBoneLocker *ShibaswapBoneLockerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShibaswapBoneLocker.Contract.TransferOwnership(&_ShibaswapBoneLocker.TransactOpts, newOwner)
}

// ShibaswapBoneLockerLockingPeriodIterator is returned from FilterLockingPeriod and is used to iterate over the raw logs and unpacked data for LockingPeriod events raised by the ShibaswapBoneLocker contract.
type ShibaswapBoneLockerLockingPeriodIterator struct {
	Event *ShibaswapBoneLockerLockingPeriod // Event containing the contract specifics and raw log

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
func (it *ShibaswapBoneLockerLockingPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapBoneLockerLockingPeriod)
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
		it.Event = new(ShibaswapBoneLockerLockingPeriod)
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
func (it *ShibaswapBoneLockerLockingPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapBoneLockerLockingPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapBoneLockerLockingPeriod represents a LockingPeriod event raised by the ShibaswapBoneLocker contract.
type ShibaswapBoneLockerLockingPeriod struct {
	User                common.Address
	NewLockingPeriod    *big.Int
	NewDevLockingPeriod *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLockingPeriod is a free log retrieval operation binding the contract event 0x2d8847c3ea1c4b6f02609bff2ac1776bc3663d31a747102d5722bdffcc2e3721.
//
// Solidity: event LockingPeriod(address indexed user, uint256 newLockingPeriod, uint256 newDevLockingPeriod)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerFilterer) FilterLockingPeriod(opts *bind.FilterOpts, user []common.Address) (*ShibaswapBoneLockerLockingPeriodIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapBoneLocker.contract.FilterLogs(opts, "LockingPeriod", userRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapBoneLockerLockingPeriodIterator{contract: _ShibaswapBoneLocker.contract, event: "LockingPeriod", logs: logs, sub: sub}, nil
}

// WatchLockingPeriod is a free log subscription operation binding the contract event 0x2d8847c3ea1c4b6f02609bff2ac1776bc3663d31a747102d5722bdffcc2e3721.
//
// Solidity: event LockingPeriod(address indexed user, uint256 newLockingPeriod, uint256 newDevLockingPeriod)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerFilterer) WatchLockingPeriod(opts *bind.WatchOpts, sink chan<- *ShibaswapBoneLockerLockingPeriod, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapBoneLocker.contract.WatchLogs(opts, "LockingPeriod", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapBoneLockerLockingPeriod)
				if err := _ShibaswapBoneLocker.contract.UnpackLog(event, "LockingPeriod", log); err != nil {
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

// ParseLockingPeriod is a log parse operation binding the contract event 0x2d8847c3ea1c4b6f02609bff2ac1776bc3663d31a747102d5722bdffcc2e3721.
//
// Solidity: event LockingPeriod(address indexed user, uint256 newLockingPeriod, uint256 newDevLockingPeriod)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerFilterer) ParseLockingPeriod(log types.Log) (*ShibaswapBoneLockerLockingPeriod, error) {
	event := new(ShibaswapBoneLockerLockingPeriod)
	if err := _ShibaswapBoneLocker.contract.UnpackLog(event, "LockingPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapBoneLockerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShibaswapBoneLocker contract.
type ShibaswapBoneLockerOwnershipTransferredIterator struct {
	Event *ShibaswapBoneLockerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShibaswapBoneLockerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapBoneLockerOwnershipTransferred)
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
		it.Event = new(ShibaswapBoneLockerOwnershipTransferred)
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
func (it *ShibaswapBoneLockerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapBoneLockerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapBoneLockerOwnershipTransferred represents a OwnershipTransferred event raised by the ShibaswapBoneLocker contract.
type ShibaswapBoneLockerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShibaswapBoneLockerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShibaswapBoneLocker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapBoneLockerOwnershipTransferredIterator{contract: _ShibaswapBoneLocker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShibaswapBoneLocker *ShibaswapBoneLockerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShibaswapBoneLockerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShibaswapBoneLocker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapBoneLockerOwnershipTransferred)
				if err := _ShibaswapBoneLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShibaswapBoneLocker *ShibaswapBoneLockerFilterer) ParseOwnershipTransferred(log types.Log) (*ShibaswapBoneLockerOwnershipTransferred, error) {
	event := new(ShibaswapBoneLockerOwnershipTransferred)
	if err := _ShibaswapBoneLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
