// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package synthetix

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

// SynthetixLiquidatorRewardsMetaData contains all meta data concerning the SynthetixLiquidatorRewards contract.
var SynthetixLiquidatorRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"CacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accumulatedRewardsPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"claimable\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"entryAccumulatedRewards\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebuildCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"updateEntry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SynthetixLiquidatorRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynthetixLiquidatorRewardsMetaData.ABI instead.
var SynthetixLiquidatorRewardsABI = SynthetixLiquidatorRewardsMetaData.ABI

// SynthetixLiquidatorRewards is an auto generated Go binding around an Ethereum contract.
type SynthetixLiquidatorRewards struct {
	SynthetixLiquidatorRewardsCaller     // Read-only binding to the contract
	SynthetixLiquidatorRewardsTransactor // Write-only binding to the contract
	SynthetixLiquidatorRewardsFilterer   // Log filterer for contract events
}

// SynthetixLiquidatorRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynthetixLiquidatorRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixLiquidatorRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynthetixLiquidatorRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixLiquidatorRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynthetixLiquidatorRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixLiquidatorRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynthetixLiquidatorRewardsSession struct {
	Contract     *SynthetixLiquidatorRewards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SynthetixLiquidatorRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynthetixLiquidatorRewardsCallerSession struct {
	Contract *SynthetixLiquidatorRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// SynthetixLiquidatorRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynthetixLiquidatorRewardsTransactorSession struct {
	Contract     *SynthetixLiquidatorRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// SynthetixLiquidatorRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynthetixLiquidatorRewardsRaw struct {
	Contract *SynthetixLiquidatorRewards // Generic contract binding to access the raw methods on
}

// SynthetixLiquidatorRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynthetixLiquidatorRewardsCallerRaw struct {
	Contract *SynthetixLiquidatorRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// SynthetixLiquidatorRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynthetixLiquidatorRewardsTransactorRaw struct {
	Contract *SynthetixLiquidatorRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynthetixLiquidatorRewards creates a new instance of SynthetixLiquidatorRewards, bound to a specific deployed contract.
func NewSynthetixLiquidatorRewards(address common.Address, backend bind.ContractBackend) (*SynthetixLiquidatorRewards, error) {
	contract, err := bindSynthetixLiquidatorRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewards{SynthetixLiquidatorRewardsCaller: SynthetixLiquidatorRewardsCaller{contract: contract}, SynthetixLiquidatorRewardsTransactor: SynthetixLiquidatorRewardsTransactor{contract: contract}, SynthetixLiquidatorRewardsFilterer: SynthetixLiquidatorRewardsFilterer{contract: contract}}, nil
}

// NewSynthetixLiquidatorRewardsCaller creates a new read-only instance of SynthetixLiquidatorRewards, bound to a specific deployed contract.
func NewSynthetixLiquidatorRewardsCaller(address common.Address, caller bind.ContractCaller) (*SynthetixLiquidatorRewardsCaller, error) {
	contract, err := bindSynthetixLiquidatorRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsCaller{contract: contract}, nil
}

// NewSynthetixLiquidatorRewardsTransactor creates a new write-only instance of SynthetixLiquidatorRewards, bound to a specific deployed contract.
func NewSynthetixLiquidatorRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynthetixLiquidatorRewardsTransactor, error) {
	contract, err := bindSynthetixLiquidatorRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsTransactor{contract: contract}, nil
}

// NewSynthetixLiquidatorRewardsFilterer creates a new log filterer instance of SynthetixLiquidatorRewards, bound to a specific deployed contract.
func NewSynthetixLiquidatorRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynthetixLiquidatorRewardsFilterer, error) {
	contract, err := bindSynthetixLiquidatorRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsFilterer{contract: contract}, nil
}

// bindSynthetixLiquidatorRewards binds a generic wrapper to an already deployed contract.
func bindSynthetixLiquidatorRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynthetixLiquidatorRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixLiquidatorRewards.Contract.SynthetixLiquidatorRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.SynthetixLiquidatorRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.SynthetixLiquidatorRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixLiquidatorRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) CONTRACTNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "CONTRACT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixLiquidatorRewards.Contract.CONTRACTNAME(&_SynthetixLiquidatorRewards.CallOpts)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixLiquidatorRewards.Contract.CONTRACTNAME(&_SynthetixLiquidatorRewards.CallOpts)
}

// AccumulatedRewardsPerShare is a free data retrieval call binding the contract method 0x061960aa.
//
// Solidity: function accumulatedRewardsPerShare() view returns(uint256)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) AccumulatedRewardsPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "accumulatedRewardsPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewardsPerShare is a free data retrieval call binding the contract method 0x061960aa.
//
// Solidity: function accumulatedRewardsPerShare() view returns(uint256)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) AccumulatedRewardsPerShare() (*big.Int, error) {
	return _SynthetixLiquidatorRewards.Contract.AccumulatedRewardsPerShare(&_SynthetixLiquidatorRewards.CallOpts)
}

// AccumulatedRewardsPerShare is a free data retrieval call binding the contract method 0x061960aa.
//
// Solidity: function accumulatedRewardsPerShare() view returns(uint256)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) AccumulatedRewardsPerShare() (*big.Int, error) {
	return _SynthetixLiquidatorRewards.Contract.AccumulatedRewardsPerShare(&_SynthetixLiquidatorRewards.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _SynthetixLiquidatorRewards.Contract.Earned(&_SynthetixLiquidatorRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _SynthetixLiquidatorRewards.Contract.Earned(&_SynthetixLiquidatorRewards.CallOpts, account)
}

// Entries is a free data retrieval call binding the contract method 0xf29ee125.
//
// Solidity: function entries(address ) view returns(uint128 claimable, uint128 entryAccumulatedRewards)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) Entries(opts *bind.CallOpts, arg0 common.Address) (struct {
	Claimable               *big.Int
	EntryAccumulatedRewards *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Claimable               *big.Int
		EntryAccumulatedRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Claimable = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EntryAccumulatedRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xf29ee125.
//
// Solidity: function entries(address ) view returns(uint128 claimable, uint128 entryAccumulatedRewards)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) Entries(arg0 common.Address) (struct {
	Claimable               *big.Int
	EntryAccumulatedRewards *big.Int
}, error) {
	return _SynthetixLiquidatorRewards.Contract.Entries(&_SynthetixLiquidatorRewards.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xf29ee125.
//
// Solidity: function entries(address ) view returns(uint128 claimable, uint128 entryAccumulatedRewards)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) Entries(arg0 common.Address) (struct {
	Claimable               *big.Int
	EntryAccumulatedRewards *big.Int
}, error) {
	return _SynthetixLiquidatorRewards.Contract.Entries(&_SynthetixLiquidatorRewards.CallOpts, arg0)
}

// Initiated is a free data retrieval call binding the contract method 0x049939f3.
//
// Solidity: function initiated(address ) view returns(bool)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) Initiated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "initiated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initiated is a free data retrieval call binding the contract method 0x049939f3.
//
// Solidity: function initiated(address ) view returns(bool)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) Initiated(arg0 common.Address) (bool, error) {
	return _SynthetixLiquidatorRewards.Contract.Initiated(&_SynthetixLiquidatorRewards.CallOpts, arg0)
}

// Initiated is a free data retrieval call binding the contract method 0x049939f3.
//
// Solidity: function initiated(address ) view returns(bool)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) Initiated(arg0 common.Address) (bool, error) {
	return _SynthetixLiquidatorRewards.Contract.Initiated(&_SynthetixLiquidatorRewards.CallOpts, arg0)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) IsResolverCached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "isResolverCached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) IsResolverCached() (bool, error) {
	return _SynthetixLiquidatorRewards.Contract.IsResolverCached(&_SynthetixLiquidatorRewards.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) IsResolverCached() (bool, error) {
	return _SynthetixLiquidatorRewards.Contract.IsResolverCached(&_SynthetixLiquidatorRewards.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) NominatedOwner() (common.Address, error) {
	return _SynthetixLiquidatorRewards.Contract.NominatedOwner(&_SynthetixLiquidatorRewards.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) NominatedOwner() (common.Address, error) {
	return _SynthetixLiquidatorRewards.Contract.NominatedOwner(&_SynthetixLiquidatorRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) Owner() (common.Address, error) {
	return _SynthetixLiquidatorRewards.Contract.Owner(&_SynthetixLiquidatorRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) Owner() (common.Address, error) {
	return _SynthetixLiquidatorRewards.Contract.Owner(&_SynthetixLiquidatorRewards.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) Resolver() (common.Address, error) {
	return _SynthetixLiquidatorRewards.Contract.Resolver(&_SynthetixLiquidatorRewards.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) Resolver() (common.Address, error) {
	return _SynthetixLiquidatorRewards.Contract.Resolver(&_SynthetixLiquidatorRewards.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCaller) ResolverAddressesRequired(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SynthetixLiquidatorRewards.contract.Call(opts, &out, "resolverAddressesRequired")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixLiquidatorRewards.Contract.ResolverAddressesRequired(&_SynthetixLiquidatorRewards.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsCallerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixLiquidatorRewards.Contract.ResolverAddressesRequired(&_SynthetixLiquidatorRewards.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.AcceptOwnership(&_SynthetixLiquidatorRewards.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.AcceptOwnership(&_SynthetixLiquidatorRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactor) GetReward(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.contract.Transact(opts, "getReward", account)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) GetReward(account common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.GetReward(&_SynthetixLiquidatorRewards.TransactOpts, account)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address account) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorSession) GetReward(account common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.GetReward(&_SynthetixLiquidatorRewards.TransactOpts, account)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.NominateNewOwner(&_SynthetixLiquidatorRewards.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.NominateNewOwner(&_SynthetixLiquidatorRewards.TransactOpts, _owner)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.NotifyRewardAmount(&_SynthetixLiquidatorRewards.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.NotifyRewardAmount(&_SynthetixLiquidatorRewards.TransactOpts, reward)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactor) RebuildCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.contract.Transact(opts, "rebuildCache")
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.RebuildCache(&_SynthetixLiquidatorRewards.TransactOpts)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.RebuildCache(&_SynthetixLiquidatorRewards.TransactOpts)
}

// UpdateEntry is a paid mutator transaction binding the contract method 0x270fb338.
//
// Solidity: function updateEntry(address account) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactor) UpdateEntry(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.contract.Transact(opts, "updateEntry", account)
}

// UpdateEntry is a paid mutator transaction binding the contract method 0x270fb338.
//
// Solidity: function updateEntry(address account) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsSession) UpdateEntry(account common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.UpdateEntry(&_SynthetixLiquidatorRewards.TransactOpts, account)
}

// UpdateEntry is a paid mutator transaction binding the contract method 0x270fb338.
//
// Solidity: function updateEntry(address account) returns()
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsTransactorSession) UpdateEntry(account common.Address) (*types.Transaction, error) {
	return _SynthetixLiquidatorRewards.Contract.UpdateEntry(&_SynthetixLiquidatorRewards.TransactOpts, account)
}

// SynthetixLiquidatorRewardsCacheUpdatedIterator is returned from FilterCacheUpdated and is used to iterate over the raw logs and unpacked data for CacheUpdated events raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsCacheUpdatedIterator struct {
	Event *SynthetixLiquidatorRewardsCacheUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixLiquidatorRewardsCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixLiquidatorRewardsCacheUpdated)
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
		it.Event = new(SynthetixLiquidatorRewardsCacheUpdated)
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
func (it *SynthetixLiquidatorRewardsCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixLiquidatorRewardsCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixLiquidatorRewardsCacheUpdated represents a CacheUpdated event raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsCacheUpdated struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCacheUpdated is a free log retrieval operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) FilterCacheUpdated(opts *bind.FilterOpts) (*SynthetixLiquidatorRewardsCacheUpdatedIterator, error) {

	logs, sub, err := _SynthetixLiquidatorRewards.contract.FilterLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsCacheUpdatedIterator{contract: _SynthetixLiquidatorRewards.contract, event: "CacheUpdated", logs: logs, sub: sub}, nil
}

// WatchCacheUpdated is a free log subscription operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) WatchCacheUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixLiquidatorRewardsCacheUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixLiquidatorRewards.contract.WatchLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixLiquidatorRewardsCacheUpdated)
				if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
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

// ParseCacheUpdated is a log parse operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) ParseCacheUpdated(log types.Log) (*SynthetixLiquidatorRewardsCacheUpdated, error) {
	event := new(SynthetixLiquidatorRewardsCacheUpdated)
	if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixLiquidatorRewardsOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsOwnerChangedIterator struct {
	Event *SynthetixLiquidatorRewardsOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SynthetixLiquidatorRewardsOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixLiquidatorRewardsOwnerChanged)
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
		it.Event = new(SynthetixLiquidatorRewardsOwnerChanged)
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
func (it *SynthetixLiquidatorRewardsOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixLiquidatorRewardsOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixLiquidatorRewardsOwnerChanged represents a OwnerChanged event raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SynthetixLiquidatorRewardsOwnerChangedIterator, error) {

	logs, sub, err := _SynthetixLiquidatorRewards.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsOwnerChangedIterator{contract: _SynthetixLiquidatorRewards.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SynthetixLiquidatorRewardsOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SynthetixLiquidatorRewards.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixLiquidatorRewardsOwnerChanged)
				if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) ParseOwnerChanged(log types.Log) (*SynthetixLiquidatorRewardsOwnerChanged, error) {
	event := new(SynthetixLiquidatorRewardsOwnerChanged)
	if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixLiquidatorRewardsOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsOwnerNominatedIterator struct {
	Event *SynthetixLiquidatorRewardsOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SynthetixLiquidatorRewardsOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixLiquidatorRewardsOwnerNominated)
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
		it.Event = new(SynthetixLiquidatorRewardsOwnerNominated)
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
func (it *SynthetixLiquidatorRewardsOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixLiquidatorRewardsOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixLiquidatorRewardsOwnerNominated represents a OwnerNominated event raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SynthetixLiquidatorRewardsOwnerNominatedIterator, error) {

	logs, sub, err := _SynthetixLiquidatorRewards.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsOwnerNominatedIterator{contract: _SynthetixLiquidatorRewards.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SynthetixLiquidatorRewardsOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SynthetixLiquidatorRewards.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixLiquidatorRewardsOwnerNominated)
				if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) ParseOwnerNominated(log types.Log) (*SynthetixLiquidatorRewardsOwnerNominated, error) {
	event := new(SynthetixLiquidatorRewardsOwnerNominated)
	if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixLiquidatorRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsRewardPaidIterator struct {
	Event *SynthetixLiquidatorRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *SynthetixLiquidatorRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixLiquidatorRewardsRewardPaid)
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
		it.Event = new(SynthetixLiquidatorRewardsRewardPaid)
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
func (it *SynthetixLiquidatorRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixLiquidatorRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixLiquidatorRewardsRewardPaid represents a RewardPaid event raised by the SynthetixLiquidatorRewards contract.
type SynthetixLiquidatorRewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*SynthetixLiquidatorRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SynthetixLiquidatorRewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixLiquidatorRewardsRewardPaidIterator{contract: _SynthetixLiquidatorRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *SynthetixLiquidatorRewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SynthetixLiquidatorRewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixLiquidatorRewardsRewardPaid)
				if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_SynthetixLiquidatorRewards *SynthetixLiquidatorRewardsFilterer) ParseRewardPaid(log types.Log) (*SynthetixLiquidatorRewardsRewardPaid, error) {
	event := new(SynthetixLiquidatorRewardsRewardPaid)
	if err := _SynthetixLiquidatorRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
