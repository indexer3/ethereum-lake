// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package convex

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

// ConvexMetaData contains all meta data concerning the Convex contract.
var ConvexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"}],\"name\":\"AddUserVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"}],\"name\":\"PoolDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"RewardActiveOnCreationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"RewardImplementationChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addUserVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"createNewPoolRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"deactivatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"active\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"poolVaultLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolVaultList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsStartActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_op\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setRewardActiveOnCreation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_imp\",\"type\":\"address\"}],\"name\":\"setRewardImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConvexABI is the input ABI used to generate the binding from.
// Deprecated: Use ConvexMetaData.ABI instead.
var ConvexABI = ConvexMetaData.ABI

// Convex is an auto generated Go binding around an Ethereum contract.
type Convex struct {
	ConvexCaller     // Read-only binding to the contract
	ConvexTransactor // Write-only binding to the contract
	ConvexFilterer   // Log filterer for contract events
}

// ConvexCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvexSession struct {
	Contract     *Convex           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvexCallerSession struct {
	Contract *ConvexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConvexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvexTransactorSession struct {
	Contract     *ConvexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvexRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvexRaw struct {
	Contract *Convex // Generic contract binding to access the raw methods on
}

// ConvexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvexCallerRaw struct {
	Contract *ConvexCaller // Generic read-only contract binding to access the raw methods on
}

// ConvexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvexTransactorRaw struct {
	Contract *ConvexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvex creates a new instance of Convex, bound to a specific deployed contract.
func NewConvex(address common.Address, backend bind.ContractBackend) (*Convex, error) {
	contract, err := bindConvex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Convex{ConvexCaller: ConvexCaller{contract: contract}, ConvexTransactor: ConvexTransactor{contract: contract}, ConvexFilterer: ConvexFilterer{contract: contract}}, nil
}

// NewConvexCaller creates a new read-only instance of Convex, bound to a specific deployed contract.
func NewConvexCaller(address common.Address, caller bind.ContractCaller) (*ConvexCaller, error) {
	contract, err := bindConvex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexCaller{contract: contract}, nil
}

// NewConvexTransactor creates a new write-only instance of Convex, bound to a specific deployed contract.
func NewConvexTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvexTransactor, error) {
	contract, err := bindConvex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexTransactor{contract: contract}, nil
}

// NewConvexFilterer creates a new log filterer instance of Convex, bound to a specific deployed contract.
func NewConvexFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvexFilterer, error) {
	contract, err := bindConvex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvexFilterer{contract: contract}, nil
}

// bindConvex binds a generic wrapper to an already deployed contract.
func bindConvex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConvexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Convex *ConvexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Convex.Contract.ConvexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Convex *ConvexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Convex.Contract.ConvexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Convex *ConvexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Convex.Contract.ConvexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Convex *ConvexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Convex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Convex *ConvexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Convex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Convex *ConvexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Convex.Contract.contract.Transact(opts, method, params...)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Convex *ConvexCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Convex *ConvexSession) Operator() (common.Address, error) {
	return _Convex.Contract.Operator(&_Convex.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Convex *ConvexCallerSession) Operator() (common.Address, error) {
	return _Convex.Contract.Operator(&_Convex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Convex *ConvexCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Convex *ConvexSession) Owner() (common.Address, error) {
	return _Convex.Contract.Owner(&_Convex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Convex *ConvexCallerSession) Owner() (common.Address, error) {
	return _Convex.Contract.Owner(&_Convex.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address implementation, address stakingAddress, address stakingToken, address rewardsAddress, uint8 active)
func (_Convex *ConvexCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Implementation common.Address
	StakingAddress common.Address
	StakingToken   common.Address
	RewardsAddress common.Address
	Active         uint8
}, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Implementation common.Address
		StakingAddress common.Address
		StakingToken   common.Address
		RewardsAddress common.Address
		Active         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Implementation = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakingAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakingToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.RewardsAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address implementation, address stakingAddress, address stakingToken, address rewardsAddress, uint8 active)
func (_Convex *ConvexSession) PoolInfo(arg0 *big.Int) (struct {
	Implementation common.Address
	StakingAddress common.Address
	StakingToken   common.Address
	RewardsAddress common.Address
	Active         uint8
}, error) {
	return _Convex.Contract.PoolInfo(&_Convex.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address implementation, address stakingAddress, address stakingToken, address rewardsAddress, uint8 active)
func (_Convex *ConvexCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Implementation common.Address
	StakingAddress common.Address
	StakingToken   common.Address
	RewardsAddress common.Address
	Active         uint8
}, error) {
	return _Convex.Contract.PoolInfo(&_Convex.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Convex *ConvexCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Convex *ConvexSession) PoolLength() (*big.Int, error) {
	return _Convex.Contract.PoolLength(&_Convex.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Convex *ConvexCallerSession) PoolLength() (*big.Int, error) {
	return _Convex.Contract.PoolLength(&_Convex.CallOpts)
}

// PoolVaultLength is a free data retrieval call binding the contract method 0x8b47168a.
//
// Solidity: function poolVaultLength(uint256 _pid) view returns(uint256)
func (_Convex *ConvexCaller) PoolVaultLength(opts *bind.CallOpts, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "poolVaultLength", _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolVaultLength is a free data retrieval call binding the contract method 0x8b47168a.
//
// Solidity: function poolVaultLength(uint256 _pid) view returns(uint256)
func (_Convex *ConvexSession) PoolVaultLength(_pid *big.Int) (*big.Int, error) {
	return _Convex.Contract.PoolVaultLength(&_Convex.CallOpts, _pid)
}

// PoolVaultLength is a free data retrieval call binding the contract method 0x8b47168a.
//
// Solidity: function poolVaultLength(uint256 _pid) view returns(uint256)
func (_Convex *ConvexCallerSession) PoolVaultLength(_pid *big.Int) (*big.Int, error) {
	return _Convex.Contract.PoolVaultLength(&_Convex.CallOpts, _pid)
}

// PoolVaultList is a free data retrieval call binding the contract method 0xafe2b66d.
//
// Solidity: function poolVaultList(uint256 , uint256 ) view returns(address)
func (_Convex *ConvexCaller) PoolVaultList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "poolVaultList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolVaultList is a free data retrieval call binding the contract method 0xafe2b66d.
//
// Solidity: function poolVaultList(uint256 , uint256 ) view returns(address)
func (_Convex *ConvexSession) PoolVaultList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Convex.Contract.PoolVaultList(&_Convex.CallOpts, arg0, arg1)
}

// PoolVaultList is a free data retrieval call binding the contract method 0xafe2b66d.
//
// Solidity: function poolVaultList(uint256 , uint256 ) view returns(address)
func (_Convex *ConvexCallerSession) PoolVaultList(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Convex.Contract.PoolVaultList(&_Convex.CallOpts, arg0, arg1)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_Convex *ConvexCaller) ProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "proxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_Convex *ConvexSession) ProxyFactory() (common.Address, error) {
	return _Convex.Contract.ProxyFactory(&_Convex.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_Convex *ConvexCallerSession) ProxyFactory() (common.Address, error) {
	return _Convex.Contract.ProxyFactory(&_Convex.CallOpts)
}

// RewardImplementation is a free data retrieval call binding the contract method 0xddc72f94.
//
// Solidity: function rewardImplementation() view returns(address)
func (_Convex *ConvexCaller) RewardImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "rewardImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardImplementation is a free data retrieval call binding the contract method 0xddc72f94.
//
// Solidity: function rewardImplementation() view returns(address)
func (_Convex *ConvexSession) RewardImplementation() (common.Address, error) {
	return _Convex.Contract.RewardImplementation(&_Convex.CallOpts)
}

// RewardImplementation is a free data retrieval call binding the contract method 0xddc72f94.
//
// Solidity: function rewardImplementation() view returns(address)
func (_Convex *ConvexCallerSession) RewardImplementation() (common.Address, error) {
	return _Convex.Contract.RewardImplementation(&_Convex.CallOpts)
}

// RewardsStartActive is a free data retrieval call binding the contract method 0x51cecf77.
//
// Solidity: function rewardsStartActive() view returns(bool)
func (_Convex *ConvexCaller) RewardsStartActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "rewardsStartActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsStartActive is a free data retrieval call binding the contract method 0x51cecf77.
//
// Solidity: function rewardsStartActive() view returns(bool)
func (_Convex *ConvexSession) RewardsStartActive() (bool, error) {
	return _Convex.Contract.RewardsStartActive(&_Convex.CallOpts)
}

// RewardsStartActive is a free data retrieval call binding the contract method 0x51cecf77.
//
// Solidity: function rewardsStartActive() view returns(bool)
func (_Convex *ConvexCallerSession) RewardsStartActive() (bool, error) {
	return _Convex.Contract.RewardsStartActive(&_Convex.CallOpts)
}

// VaultMap is a free data retrieval call binding the contract method 0xbfff576f.
//
// Solidity: function vaultMap(uint256 , address ) view returns(address)
func (_Convex *ConvexCaller) VaultMap(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Convex.contract.Call(opts, &out, "vaultMap", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultMap is a free data retrieval call binding the contract method 0xbfff576f.
//
// Solidity: function vaultMap(uint256 , address ) view returns(address)
func (_Convex *ConvexSession) VaultMap(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _Convex.Contract.VaultMap(&_Convex.CallOpts, arg0, arg1)
}

// VaultMap is a free data retrieval call binding the contract method 0xbfff576f.
//
// Solidity: function vaultMap(uint256 , address ) view returns(address)
func (_Convex *ConvexCallerSession) VaultMap(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _Convex.Contract.VaultMap(&_Convex.CallOpts, arg0, arg1)
}

// AddPool is a paid mutator transaction binding the contract method 0xaba65098.
//
// Solidity: function addPool(address _implementation, address _stakingAddress, address _stakingToken) returns()
func (_Convex *ConvexTransactor) AddPool(opts *bind.TransactOpts, _implementation common.Address, _stakingAddress common.Address, _stakingToken common.Address) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "addPool", _implementation, _stakingAddress, _stakingToken)
}

// AddPool is a paid mutator transaction binding the contract method 0xaba65098.
//
// Solidity: function addPool(address _implementation, address _stakingAddress, address _stakingToken) returns()
func (_Convex *ConvexSession) AddPool(_implementation common.Address, _stakingAddress common.Address, _stakingToken common.Address) (*types.Transaction, error) {
	return _Convex.Contract.AddPool(&_Convex.TransactOpts, _implementation, _stakingAddress, _stakingToken)
}

// AddPool is a paid mutator transaction binding the contract method 0xaba65098.
//
// Solidity: function addPool(address _implementation, address _stakingAddress, address _stakingToken) returns()
func (_Convex *ConvexTransactorSession) AddPool(_implementation common.Address, _stakingAddress common.Address, _stakingToken common.Address) (*types.Transaction, error) {
	return _Convex.Contract.AddPool(&_Convex.TransactOpts, _implementation, _stakingAddress, _stakingToken)
}

// AddUserVault is a paid mutator transaction binding the contract method 0xf2bc788b.
//
// Solidity: function addUserVault(uint256 _pid, address _user) returns(address vault, address stakingAddress, address stakingToken, address rewards)
func (_Convex *ConvexTransactor) AddUserVault(opts *bind.TransactOpts, _pid *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "addUserVault", _pid, _user)
}

// AddUserVault is a paid mutator transaction binding the contract method 0xf2bc788b.
//
// Solidity: function addUserVault(uint256 _pid, address _user) returns(address vault, address stakingAddress, address stakingToken, address rewards)
func (_Convex *ConvexSession) AddUserVault(_pid *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Convex.Contract.AddUserVault(&_Convex.TransactOpts, _pid, _user)
}

// AddUserVault is a paid mutator transaction binding the contract method 0xf2bc788b.
//
// Solidity: function addUserVault(uint256 _pid, address _user) returns(address vault, address stakingAddress, address stakingToken, address rewards)
func (_Convex *ConvexTransactorSession) AddUserVault(_pid *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Convex.Contract.AddUserVault(&_Convex.TransactOpts, _pid, _user)
}

// CreateNewPoolRewards is a paid mutator transaction binding the contract method 0xe57197cf.
//
// Solidity: function createNewPoolRewards(uint256 _pid) returns()
func (_Convex *ConvexTransactor) CreateNewPoolRewards(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "createNewPoolRewards", _pid)
}

// CreateNewPoolRewards is a paid mutator transaction binding the contract method 0xe57197cf.
//
// Solidity: function createNewPoolRewards(uint256 _pid) returns()
func (_Convex *ConvexSession) CreateNewPoolRewards(_pid *big.Int) (*types.Transaction, error) {
	return _Convex.Contract.CreateNewPoolRewards(&_Convex.TransactOpts, _pid)
}

// CreateNewPoolRewards is a paid mutator transaction binding the contract method 0xe57197cf.
//
// Solidity: function createNewPoolRewards(uint256 _pid) returns()
func (_Convex *ConvexTransactorSession) CreateNewPoolRewards(_pid *big.Int) (*types.Transaction, error) {
	return _Convex.Contract.CreateNewPoolRewards(&_Convex.TransactOpts, _pid)
}

// DeactivatePool is a paid mutator transaction binding the contract method 0x9abd9b05.
//
// Solidity: function deactivatePool(uint256 _pid) returns()
func (_Convex *ConvexTransactor) DeactivatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "deactivatePool", _pid)
}

// DeactivatePool is a paid mutator transaction binding the contract method 0x9abd9b05.
//
// Solidity: function deactivatePool(uint256 _pid) returns()
func (_Convex *ConvexSession) DeactivatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Convex.Contract.DeactivatePool(&_Convex.TransactOpts, _pid)
}

// DeactivatePool is a paid mutator transaction binding the contract method 0x9abd9b05.
//
// Solidity: function deactivatePool(uint256 _pid) returns()
func (_Convex *ConvexTransactorSession) DeactivatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Convex.Contract.DeactivatePool(&_Convex.TransactOpts, _pid)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _op) returns()
func (_Convex *ConvexTransactor) SetOperator(opts *bind.TransactOpts, _op common.Address) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "setOperator", _op)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _op) returns()
func (_Convex *ConvexSession) SetOperator(_op common.Address) (*types.Transaction, error) {
	return _Convex.Contract.SetOperator(&_Convex.TransactOpts, _op)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _op) returns()
func (_Convex *ConvexTransactorSession) SetOperator(_op common.Address) (*types.Transaction, error) {
	return _Convex.Contract.SetOperator(&_Convex.TransactOpts, _op)
}

// SetRewardActiveOnCreation is a paid mutator transaction binding the contract method 0x3fd46786.
//
// Solidity: function setRewardActiveOnCreation(bool _active) returns()
func (_Convex *ConvexTransactor) SetRewardActiveOnCreation(opts *bind.TransactOpts, _active bool) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "setRewardActiveOnCreation", _active)
}

// SetRewardActiveOnCreation is a paid mutator transaction binding the contract method 0x3fd46786.
//
// Solidity: function setRewardActiveOnCreation(bool _active) returns()
func (_Convex *ConvexSession) SetRewardActiveOnCreation(_active bool) (*types.Transaction, error) {
	return _Convex.Contract.SetRewardActiveOnCreation(&_Convex.TransactOpts, _active)
}

// SetRewardActiveOnCreation is a paid mutator transaction binding the contract method 0x3fd46786.
//
// Solidity: function setRewardActiveOnCreation(bool _active) returns()
func (_Convex *ConvexTransactorSession) SetRewardActiveOnCreation(_active bool) (*types.Transaction, error) {
	return _Convex.Contract.SetRewardActiveOnCreation(&_Convex.TransactOpts, _active)
}

// SetRewardImplementation is a paid mutator transaction binding the contract method 0xfa58ee02.
//
// Solidity: function setRewardImplementation(address _imp) returns()
func (_Convex *ConvexTransactor) SetRewardImplementation(opts *bind.TransactOpts, _imp common.Address) (*types.Transaction, error) {
	return _Convex.contract.Transact(opts, "setRewardImplementation", _imp)
}

// SetRewardImplementation is a paid mutator transaction binding the contract method 0xfa58ee02.
//
// Solidity: function setRewardImplementation(address _imp) returns()
func (_Convex *ConvexSession) SetRewardImplementation(_imp common.Address) (*types.Transaction, error) {
	return _Convex.Contract.SetRewardImplementation(&_Convex.TransactOpts, _imp)
}

// SetRewardImplementation is a paid mutator transaction binding the contract method 0xfa58ee02.
//
// Solidity: function setRewardImplementation(address _imp) returns()
func (_Convex *ConvexTransactorSession) SetRewardImplementation(_imp common.Address) (*types.Transaction, error) {
	return _Convex.Contract.SetRewardImplementation(&_Convex.TransactOpts, _imp)
}

// ConvexAddUserVaultIterator is returned from FilterAddUserVault and is used to iterate over the raw logs and unpacked data for AddUserVault events raised by the Convex contract.
type ConvexAddUserVaultIterator struct {
	Event *ConvexAddUserVault // Event containing the contract specifics and raw log

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
func (it *ConvexAddUserVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexAddUserVault)
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
		it.Event = new(ConvexAddUserVault)
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
func (it *ConvexAddUserVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexAddUserVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexAddUserVault represents a AddUserVault event raised by the Convex contract.
type ConvexAddUserVault struct {
	User   common.Address
	Poolid *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddUserVault is a free log retrieval operation binding the contract event 0xc3a719ac2c66bb292413ff9bb5cc91f486266e1b70bf1b394f666fc761ec64a3.
//
// Solidity: event AddUserVault(address indexed user, uint256 indexed poolid)
func (_Convex *ConvexFilterer) FilterAddUserVault(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*ConvexAddUserVaultIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Convex.contract.FilterLogs(opts, "AddUserVault", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &ConvexAddUserVaultIterator{contract: _Convex.contract, event: "AddUserVault", logs: logs, sub: sub}, nil
}

// WatchAddUserVault is a free log subscription operation binding the contract event 0xc3a719ac2c66bb292413ff9bb5cc91f486266e1b70bf1b394f666fc761ec64a3.
//
// Solidity: event AddUserVault(address indexed user, uint256 indexed poolid)
func (_Convex *ConvexFilterer) WatchAddUserVault(opts *bind.WatchOpts, sink chan<- *ConvexAddUserVault, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Convex.contract.WatchLogs(opts, "AddUserVault", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexAddUserVault)
				if err := _Convex.contract.UnpackLog(event, "AddUserVault", log); err != nil {
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

// ParseAddUserVault is a log parse operation binding the contract event 0xc3a719ac2c66bb292413ff9bb5cc91f486266e1b70bf1b394f666fc761ec64a3.
//
// Solidity: event AddUserVault(address indexed user, uint256 indexed poolid)
func (_Convex *ConvexFilterer) ParseAddUserVault(log types.Log) (*ConvexAddUserVault, error) {
	event := new(ConvexAddUserVault)
	if err := _Convex.contract.UnpackLog(event, "AddUserVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the Convex contract.
type ConvexOperatorChangedIterator struct {
	Event *ConvexOperatorChanged // Event containing the contract specifics and raw log

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
func (it *ConvexOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexOperatorChanged)
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
		it.Event = new(ConvexOperatorChanged)
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
func (it *ConvexOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexOperatorChanged represents a OperatorChanged event raised by the Convex contract.
type ConvexOperatorChanged struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address indexed account)
func (_Convex *ConvexFilterer) FilterOperatorChanged(opts *bind.FilterOpts, account []common.Address) (*ConvexOperatorChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Convex.contract.FilterLogs(opts, "OperatorChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &ConvexOperatorChangedIterator{contract: _Convex.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address indexed account)
func (_Convex *ConvexFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *ConvexOperatorChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Convex.contract.WatchLogs(opts, "OperatorChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexOperatorChanged)
				if err := _Convex.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address indexed account)
func (_Convex *ConvexFilterer) ParseOperatorChanged(log types.Log) (*ConvexOperatorChanged, error) {
	event := new(ConvexOperatorChanged)
	if err := _Convex.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Convex contract.
type ConvexPoolCreatedIterator struct {
	Event *ConvexPoolCreated // Event containing the contract specifics and raw log

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
func (it *ConvexPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexPoolCreated)
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
		it.Event = new(ConvexPoolCreated)
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
func (it *ConvexPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexPoolCreated represents a PoolCreated event raised by the Convex contract.
type ConvexPoolCreated struct {
	Poolid         *big.Int
	Implementation common.Address
	StakingAddress common.Address
	StakingToken   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x070da9921e58d11c5e49283ff7930c3c05239838f04b056bb98b767be6955879.
//
// Solidity: event PoolCreated(uint256 indexed poolid, address indexed implementation, address stakingAddress, address stakingToken)
func (_Convex *ConvexFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolid []*big.Int, implementation []common.Address) (*ConvexPoolCreatedIterator, error) {

	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Convex.contract.FilterLogs(opts, "PoolCreated", poolidRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return &ConvexPoolCreatedIterator{contract: _Convex.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x070da9921e58d11c5e49283ff7930c3c05239838f04b056bb98b767be6955879.
//
// Solidity: event PoolCreated(uint256 indexed poolid, address indexed implementation, address stakingAddress, address stakingToken)
func (_Convex *ConvexFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *ConvexPoolCreated, poolid []*big.Int, implementation []common.Address) (event.Subscription, error) {

	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Convex.contract.WatchLogs(opts, "PoolCreated", poolidRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexPoolCreated)
				if err := _Convex.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x070da9921e58d11c5e49283ff7930c3c05239838f04b056bb98b767be6955879.
//
// Solidity: event PoolCreated(uint256 indexed poolid, address indexed implementation, address stakingAddress, address stakingToken)
func (_Convex *ConvexFilterer) ParsePoolCreated(log types.Log) (*ConvexPoolCreated, error) {
	event := new(ConvexPoolCreated)
	if err := _Convex.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexPoolDeactivatedIterator is returned from FilterPoolDeactivated and is used to iterate over the raw logs and unpacked data for PoolDeactivated events raised by the Convex contract.
type ConvexPoolDeactivatedIterator struct {
	Event *ConvexPoolDeactivated // Event containing the contract specifics and raw log

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
func (it *ConvexPoolDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexPoolDeactivated)
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
		it.Event = new(ConvexPoolDeactivated)
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
func (it *ConvexPoolDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexPoolDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexPoolDeactivated represents a PoolDeactivated event raised by the Convex contract.
type ConvexPoolDeactivated struct {
	Poolid *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolDeactivated is a free log retrieval operation binding the contract event 0xf924dd04accfc1837d0eeddb10bc7732e2057f5d916c0b5a21e8372393b244a2.
//
// Solidity: event PoolDeactivated(uint256 indexed poolid)
func (_Convex *ConvexFilterer) FilterPoolDeactivated(opts *bind.FilterOpts, poolid []*big.Int) (*ConvexPoolDeactivatedIterator, error) {

	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Convex.contract.FilterLogs(opts, "PoolDeactivated", poolidRule)
	if err != nil {
		return nil, err
	}
	return &ConvexPoolDeactivatedIterator{contract: _Convex.contract, event: "PoolDeactivated", logs: logs, sub: sub}, nil
}

// WatchPoolDeactivated is a free log subscription operation binding the contract event 0xf924dd04accfc1837d0eeddb10bc7732e2057f5d916c0b5a21e8372393b244a2.
//
// Solidity: event PoolDeactivated(uint256 indexed poolid)
func (_Convex *ConvexFilterer) WatchPoolDeactivated(opts *bind.WatchOpts, sink chan<- *ConvexPoolDeactivated, poolid []*big.Int) (event.Subscription, error) {

	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _Convex.contract.WatchLogs(opts, "PoolDeactivated", poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexPoolDeactivated)
				if err := _Convex.contract.UnpackLog(event, "PoolDeactivated", log); err != nil {
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

// ParsePoolDeactivated is a log parse operation binding the contract event 0xf924dd04accfc1837d0eeddb10bc7732e2057f5d916c0b5a21e8372393b244a2.
//
// Solidity: event PoolDeactivated(uint256 indexed poolid)
func (_Convex *ConvexFilterer) ParsePoolDeactivated(log types.Log) (*ConvexPoolDeactivated, error) {
	event := new(ConvexPoolDeactivated)
	if err := _Convex.contract.UnpackLog(event, "PoolDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexRewardActiveOnCreationChangedIterator is returned from FilterRewardActiveOnCreationChanged and is used to iterate over the raw logs and unpacked data for RewardActiveOnCreationChanged events raised by the Convex contract.
type ConvexRewardActiveOnCreationChangedIterator struct {
	Event *ConvexRewardActiveOnCreationChanged // Event containing the contract specifics and raw log

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
func (it *ConvexRewardActiveOnCreationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexRewardActiveOnCreationChanged)
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
		it.Event = new(ConvexRewardActiveOnCreationChanged)
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
func (it *ConvexRewardActiveOnCreationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexRewardActiveOnCreationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexRewardActiveOnCreationChanged represents a RewardActiveOnCreationChanged event raised by the Convex contract.
type ConvexRewardActiveOnCreationChanged struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardActiveOnCreationChanged is a free log retrieval operation binding the contract event 0x7466d7918c00b6a6e62641ace256dadfc65e824de0a0d0740eda888901c2b63b.
//
// Solidity: event RewardActiveOnCreationChanged(bool value)
func (_Convex *ConvexFilterer) FilterRewardActiveOnCreationChanged(opts *bind.FilterOpts) (*ConvexRewardActiveOnCreationChangedIterator, error) {

	logs, sub, err := _Convex.contract.FilterLogs(opts, "RewardActiveOnCreationChanged")
	if err != nil {
		return nil, err
	}
	return &ConvexRewardActiveOnCreationChangedIterator{contract: _Convex.contract, event: "RewardActiveOnCreationChanged", logs: logs, sub: sub}, nil
}

// WatchRewardActiveOnCreationChanged is a free log subscription operation binding the contract event 0x7466d7918c00b6a6e62641ace256dadfc65e824de0a0d0740eda888901c2b63b.
//
// Solidity: event RewardActiveOnCreationChanged(bool value)
func (_Convex *ConvexFilterer) WatchRewardActiveOnCreationChanged(opts *bind.WatchOpts, sink chan<- *ConvexRewardActiveOnCreationChanged) (event.Subscription, error) {

	logs, sub, err := _Convex.contract.WatchLogs(opts, "RewardActiveOnCreationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexRewardActiveOnCreationChanged)
				if err := _Convex.contract.UnpackLog(event, "RewardActiveOnCreationChanged", log); err != nil {
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

// ParseRewardActiveOnCreationChanged is a log parse operation binding the contract event 0x7466d7918c00b6a6e62641ace256dadfc65e824de0a0d0740eda888901c2b63b.
//
// Solidity: event RewardActiveOnCreationChanged(bool value)
func (_Convex *ConvexFilterer) ParseRewardActiveOnCreationChanged(log types.Log) (*ConvexRewardActiveOnCreationChanged, error) {
	event := new(ConvexRewardActiveOnCreationChanged)
	if err := _Convex.contract.UnpackLog(event, "RewardActiveOnCreationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexRewardImplementationChangedIterator is returned from FilterRewardImplementationChanged and is used to iterate over the raw logs and unpacked data for RewardImplementationChanged events raised by the Convex contract.
type ConvexRewardImplementationChangedIterator struct {
	Event *ConvexRewardImplementationChanged // Event containing the contract specifics and raw log

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
func (it *ConvexRewardImplementationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexRewardImplementationChanged)
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
		it.Event = new(ConvexRewardImplementationChanged)
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
func (it *ConvexRewardImplementationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexRewardImplementationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexRewardImplementationChanged represents a RewardImplementationChanged event raised by the Convex contract.
type ConvexRewardImplementationChanged struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardImplementationChanged is a free log retrieval operation binding the contract event 0xfbeb0fbc478d221bbef0241e46aa78721db7379c47cc2dddf9fc562ac071486c.
//
// Solidity: event RewardImplementationChanged(address indexed implementation)
func (_Convex *ConvexFilterer) FilterRewardImplementationChanged(opts *bind.FilterOpts, implementation []common.Address) (*ConvexRewardImplementationChangedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Convex.contract.FilterLogs(opts, "RewardImplementationChanged", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ConvexRewardImplementationChangedIterator{contract: _Convex.contract, event: "RewardImplementationChanged", logs: logs, sub: sub}, nil
}

// WatchRewardImplementationChanged is a free log subscription operation binding the contract event 0xfbeb0fbc478d221bbef0241e46aa78721db7379c47cc2dddf9fc562ac071486c.
//
// Solidity: event RewardImplementationChanged(address indexed implementation)
func (_Convex *ConvexFilterer) WatchRewardImplementationChanged(opts *bind.WatchOpts, sink chan<- *ConvexRewardImplementationChanged, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Convex.contract.WatchLogs(opts, "RewardImplementationChanged", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexRewardImplementationChanged)
				if err := _Convex.contract.UnpackLog(event, "RewardImplementationChanged", log); err != nil {
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

// ParseRewardImplementationChanged is a log parse operation binding the contract event 0xfbeb0fbc478d221bbef0241e46aa78721db7379c47cc2dddf9fc562ac071486c.
//
// Solidity: event RewardImplementationChanged(address indexed implementation)
func (_Convex *ConvexFilterer) ParseRewardImplementationChanged(log types.Log) (*ConvexRewardImplementationChanged, error) {
	event := new(ConvexRewardImplementationChanged)
	if err := _Convex.contract.UnpackLog(event, "RewardImplementationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
