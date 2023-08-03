// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package insta_dapp

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

// AccountResolverAccountData is an auto generated low-level Go binding around an user-defined struct.
type AccountResolverAccountData struct {
	ID          *big.Int
	Account     common.Address
	Version     *big.Int
	Authorities []common.Address
}

// AccountResolverAuthType is an auto generated low-level Go binding around an user-defined struct.
type AccountResolverAuthType struct {
	Owner    common.Address
	AuthType *big.Int
}

// AccountResolverAuthorityData is an auto generated low-level Go binding around an user-defined struct.
type AccountResolverAuthorityData struct {
	IDs      []uint64
	Accounts []common.Address
	Versions []*big.Int
}

// ConnectorsResolverConnectorsData is an auto generated low-level Go binding around an user-defined struct.
type ConnectorsResolverConnectorsData struct {
	Connector   common.Address
	ConnectorID *big.Int
	Name        string
}

// InstaDappDSAResolverMetaData contains all meta data concerning the InstaDappDSAResolver contract.
var InstaDappDSAResolverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_index\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_gnosisFactory\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountAuthorities\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountAuthoritiesTypes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"authType\",\"type\":\"uint256\"}],\"internalType\":\"structAccountResolver.AuthType[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"authorities\",\"type\":\"address[]\"}],\"internalType\":\"structAccountResolver.AccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAccountIdDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"authorities\",\"type\":\"address[]\"}],\"internalType\":\"structAccountResolver.AccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"getAccountVersions\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"getAuthorityAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"getAuthorityDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"IDs\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"versions\",\"type\":\"uint256[]\"}],\"internalType\":\"structAccountResolver.AuthorityData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"getAuthorityIDs\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"authorities\",\"type\":\"address[]\"}],\"name\":\"getAuthorityTypes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"authType\",\"type\":\"uint256\"}],\"internalType\":\"structAccountResolver.AuthType[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getContractCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"o_code\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnabledConnectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnabledConnectorsData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"connectorID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structConnectorsResolver.ConnectorsData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getIDAuthorities\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConnectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConnectorsData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"connectorID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structConnectorsResolver.ConnectorsData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gnosisFactoryContracts\",\"outputs\":[{\"internalType\":\"contractGnosisFactoryInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"index\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isShield\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"shield\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InstaDappDSAResolverABI is the input ABI used to generate the binding from.
// Deprecated: Use InstaDappDSAResolverMetaData.ABI instead.
var InstaDappDSAResolverABI = InstaDappDSAResolverMetaData.ABI

// InstaDappDSAResolver is an auto generated Go binding around an Ethereum contract.
type InstaDappDSAResolver struct {
	InstaDappDSAResolverCaller     // Read-only binding to the contract
	InstaDappDSAResolverTransactor // Write-only binding to the contract
	InstaDappDSAResolverFilterer   // Log filterer for contract events
}

// InstaDappDSAResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstaDappDSAResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappDSAResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstaDappDSAResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappDSAResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstaDappDSAResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappDSAResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstaDappDSAResolverSession struct {
	Contract     *InstaDappDSAResolver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InstaDappDSAResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstaDappDSAResolverCallerSession struct {
	Contract *InstaDappDSAResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InstaDappDSAResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstaDappDSAResolverTransactorSession struct {
	Contract     *InstaDappDSAResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InstaDappDSAResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstaDappDSAResolverRaw struct {
	Contract *InstaDappDSAResolver // Generic contract binding to access the raw methods on
}

// InstaDappDSAResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstaDappDSAResolverCallerRaw struct {
	Contract *InstaDappDSAResolverCaller // Generic read-only contract binding to access the raw methods on
}

// InstaDappDSAResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstaDappDSAResolverTransactorRaw struct {
	Contract *InstaDappDSAResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstaDappDSAResolver creates a new instance of InstaDappDSAResolver, bound to a specific deployed contract.
func NewInstaDappDSAResolver(address common.Address, backend bind.ContractBackend) (*InstaDappDSAResolver, error) {
	contract, err := bindInstaDappDSAResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InstaDappDSAResolver{InstaDappDSAResolverCaller: InstaDappDSAResolverCaller{contract: contract}, InstaDappDSAResolverTransactor: InstaDappDSAResolverTransactor{contract: contract}, InstaDappDSAResolverFilterer: InstaDappDSAResolverFilterer{contract: contract}}, nil
}

// NewInstaDappDSAResolverCaller creates a new read-only instance of InstaDappDSAResolver, bound to a specific deployed contract.
func NewInstaDappDSAResolverCaller(address common.Address, caller bind.ContractCaller) (*InstaDappDSAResolverCaller, error) {
	contract, err := bindInstaDappDSAResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstaDappDSAResolverCaller{contract: contract}, nil
}

// NewInstaDappDSAResolverTransactor creates a new write-only instance of InstaDappDSAResolver, bound to a specific deployed contract.
func NewInstaDappDSAResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*InstaDappDSAResolverTransactor, error) {
	contract, err := bindInstaDappDSAResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstaDappDSAResolverTransactor{contract: contract}, nil
}

// NewInstaDappDSAResolverFilterer creates a new log filterer instance of InstaDappDSAResolver, bound to a specific deployed contract.
func NewInstaDappDSAResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*InstaDappDSAResolverFilterer, error) {
	contract, err := bindInstaDappDSAResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstaDappDSAResolverFilterer{contract: contract}, nil
}

// bindInstaDappDSAResolver binds a generic wrapper to an already deployed contract.
func bindInstaDappDSAResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InstaDappDSAResolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstaDappDSAResolver *InstaDappDSAResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstaDappDSAResolver.Contract.InstaDappDSAResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstaDappDSAResolver *InstaDappDSAResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappDSAResolver.Contract.InstaDappDSAResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstaDappDSAResolver *InstaDappDSAResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstaDappDSAResolver.Contract.InstaDappDSAResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstaDappDSAResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstaDappDSAResolver *InstaDappDSAResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappDSAResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstaDappDSAResolver *InstaDappDSAResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstaDappDSAResolver.Contract.contract.Transact(opts, method, params...)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) Connectors(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "connectors")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) Connectors() (common.Address, error) {
	return _InstaDappDSAResolver.Contract.Connectors(&_InstaDappDSAResolver.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) Connectors() (common.Address, error) {
	return _InstaDappDSAResolver.Contract.Connectors(&_InstaDappDSAResolver.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0x77ae211a.
//
// Solidity: function getAccount(uint64 id) view returns(address account)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAccount(opts *bind.CallOpts, id uint64) (common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAccount", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0x77ae211a.
//
// Solidity: function getAccount(uint64 id) view returns(address account)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAccount(id uint64) (common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetAccount(&_InstaDappDSAResolver.CallOpts, id)
}

// GetAccount is a free data retrieval call binding the contract method 0x77ae211a.
//
// Solidity: function getAccount(uint64 id) view returns(address account)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAccount(id uint64) (common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetAccount(&_InstaDappDSAResolver.CallOpts, id)
}

// GetAccountAuthorities is a free data retrieval call binding the contract method 0xb05088c4.
//
// Solidity: function getAccountAuthorities(address account) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAccountAuthorities(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAccountAuthorities", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAccountAuthorities is a free data retrieval call binding the contract method 0xb05088c4.
//
// Solidity: function getAccountAuthorities(address account) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAccountAuthorities(account common.Address) ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetAccountAuthorities(&_InstaDappDSAResolver.CallOpts, account)
}

// GetAccountAuthorities is a free data retrieval call binding the contract method 0xb05088c4.
//
// Solidity: function getAccountAuthorities(address account) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAccountAuthorities(account common.Address) ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetAccountAuthorities(&_InstaDappDSAResolver.CallOpts, account)
}

// GetAccountAuthoritiesTypes is a free data retrieval call binding the contract method 0x62ec17bd.
//
// Solidity: function getAccountAuthoritiesTypes(address account) view returns((address,uint256)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAccountAuthoritiesTypes(opts *bind.CallOpts, account common.Address) ([]AccountResolverAuthType, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAccountAuthoritiesTypes", account)

	if err != nil {
		return *new([]AccountResolverAuthType), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccountResolverAuthType)).(*[]AccountResolverAuthType)

	return out0, err

}

// GetAccountAuthoritiesTypes is a free data retrieval call binding the contract method 0x62ec17bd.
//
// Solidity: function getAccountAuthoritiesTypes(address account) view returns((address,uint256)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAccountAuthoritiesTypes(account common.Address) ([]AccountResolverAuthType, error) {
	return _InstaDappDSAResolver.Contract.GetAccountAuthoritiesTypes(&_InstaDappDSAResolver.CallOpts, account)
}

// GetAccountAuthoritiesTypes is a free data retrieval call binding the contract method 0x62ec17bd.
//
// Solidity: function getAccountAuthoritiesTypes(address account) view returns((address,uint256)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAccountAuthoritiesTypes(account common.Address) ([]AccountResolverAuthType, error) {
	return _InstaDappDSAResolver.Contract.GetAccountAuthoritiesTypes(&_InstaDappDSAResolver.CallOpts, account)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address account) view returns((uint256,address,uint256,address[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAccountDetails(opts *bind.CallOpts, account common.Address) (AccountResolverAccountData, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAccountDetails", account)

	if err != nil {
		return *new(AccountResolverAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountResolverAccountData)).(*AccountResolverAccountData)

	return out0, err

}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address account) view returns((uint256,address,uint256,address[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAccountDetails(account common.Address) (AccountResolverAccountData, error) {
	return _InstaDappDSAResolver.Contract.GetAccountDetails(&_InstaDappDSAResolver.CallOpts, account)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address account) view returns((uint256,address,uint256,address[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAccountDetails(account common.Address) (AccountResolverAccountData, error) {
	return _InstaDappDSAResolver.Contract.GetAccountDetails(&_InstaDappDSAResolver.CallOpts, account)
}

// GetAccountIdDetails is a free data retrieval call binding the contract method 0x385d93ec.
//
// Solidity: function getAccountIdDetails(uint256 id) view returns((uint256,address,uint256,address[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAccountIdDetails(opts *bind.CallOpts, id *big.Int) (AccountResolverAccountData, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAccountIdDetails", id)

	if err != nil {
		return *new(AccountResolverAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountResolverAccountData)).(*AccountResolverAccountData)

	return out0, err

}

// GetAccountIdDetails is a free data retrieval call binding the contract method 0x385d93ec.
//
// Solidity: function getAccountIdDetails(uint256 id) view returns((uint256,address,uint256,address[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAccountIdDetails(id *big.Int) (AccountResolverAccountData, error) {
	return _InstaDappDSAResolver.Contract.GetAccountIdDetails(&_InstaDappDSAResolver.CallOpts, id)
}

// GetAccountIdDetails is a free data retrieval call binding the contract method 0x385d93ec.
//
// Solidity: function getAccountIdDetails(uint256 id) view returns((uint256,address,uint256,address[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAccountIdDetails(id *big.Int) (AccountResolverAccountData, error) {
	return _InstaDappDSAResolver.Contract.GetAccountIdDetails(&_InstaDappDSAResolver.CallOpts, id)
}

// GetAccountVersions is a free data retrieval call binding the contract method 0xe8f4880d.
//
// Solidity: function getAccountVersions(address[] accounts) view returns(uint256[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAccountVersions(opts *bind.CallOpts, accounts []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAccountVersions", accounts)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAccountVersions is a free data retrieval call binding the contract method 0xe8f4880d.
//
// Solidity: function getAccountVersions(address[] accounts) view returns(uint256[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAccountVersions(accounts []common.Address) ([]*big.Int, error) {
	return _InstaDappDSAResolver.Contract.GetAccountVersions(&_InstaDappDSAResolver.CallOpts, accounts)
}

// GetAccountVersions is a free data retrieval call binding the contract method 0xe8f4880d.
//
// Solidity: function getAccountVersions(address[] accounts) view returns(uint256[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAccountVersions(accounts []common.Address) ([]*big.Int, error) {
	return _InstaDappDSAResolver.Contract.GetAccountVersions(&_InstaDappDSAResolver.CallOpts, accounts)
}

// GetAuthorityAccounts is a free data retrieval call binding the contract method 0x757957c4.
//
// Solidity: function getAuthorityAccounts(address authority) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAuthorityAccounts(opts *bind.CallOpts, authority common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAuthorityAccounts", authority)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorityAccounts is a free data retrieval call binding the contract method 0x757957c4.
//
// Solidity: function getAuthorityAccounts(address authority) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAuthorityAccounts(authority common.Address) ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityAccounts(&_InstaDappDSAResolver.CallOpts, authority)
}

// GetAuthorityAccounts is a free data retrieval call binding the contract method 0x757957c4.
//
// Solidity: function getAuthorityAccounts(address authority) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAuthorityAccounts(authority common.Address) ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityAccounts(&_InstaDappDSAResolver.CallOpts, authority)
}

// GetAuthorityDetails is a free data retrieval call binding the contract method 0x93865244.
//
// Solidity: function getAuthorityDetails(address authority) view returns((uint64[],address[],uint256[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAuthorityDetails(opts *bind.CallOpts, authority common.Address) (AccountResolverAuthorityData, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAuthorityDetails", authority)

	if err != nil {
		return *new(AccountResolverAuthorityData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountResolverAuthorityData)).(*AccountResolverAuthorityData)

	return out0, err

}

// GetAuthorityDetails is a free data retrieval call binding the contract method 0x93865244.
//
// Solidity: function getAuthorityDetails(address authority) view returns((uint64[],address[],uint256[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAuthorityDetails(authority common.Address) (AccountResolverAuthorityData, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityDetails(&_InstaDappDSAResolver.CallOpts, authority)
}

// GetAuthorityDetails is a free data retrieval call binding the contract method 0x93865244.
//
// Solidity: function getAuthorityDetails(address authority) view returns((uint64[],address[],uint256[]))
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAuthorityDetails(authority common.Address) (AccountResolverAuthorityData, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityDetails(&_InstaDappDSAResolver.CallOpts, authority)
}

// GetAuthorityIDs is a free data retrieval call binding the contract method 0x003e1a36.
//
// Solidity: function getAuthorityIDs(address authority) view returns(uint64[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAuthorityIDs(opts *bind.CallOpts, authority common.Address) ([]uint64, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAuthorityIDs", authority)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetAuthorityIDs is a free data retrieval call binding the contract method 0x003e1a36.
//
// Solidity: function getAuthorityIDs(address authority) view returns(uint64[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAuthorityIDs(authority common.Address) ([]uint64, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityIDs(&_InstaDappDSAResolver.CallOpts, authority)
}

// GetAuthorityIDs is a free data retrieval call binding the contract method 0x003e1a36.
//
// Solidity: function getAuthorityIDs(address authority) view returns(uint64[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAuthorityIDs(authority common.Address) ([]uint64, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityIDs(&_InstaDappDSAResolver.CallOpts, authority)
}

// GetAuthorityTypes is a free data retrieval call binding the contract method 0xfdca3798.
//
// Solidity: function getAuthorityTypes(address[] authorities) view returns((address,uint256)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetAuthorityTypes(opts *bind.CallOpts, authorities []common.Address) ([]AccountResolverAuthType, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getAuthorityTypes", authorities)

	if err != nil {
		return *new([]AccountResolverAuthType), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccountResolverAuthType)).(*[]AccountResolverAuthType)

	return out0, err

}

// GetAuthorityTypes is a free data retrieval call binding the contract method 0xfdca3798.
//
// Solidity: function getAuthorityTypes(address[] authorities) view returns((address,uint256)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetAuthorityTypes(authorities []common.Address) ([]AccountResolverAuthType, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityTypes(&_InstaDappDSAResolver.CallOpts, authorities)
}

// GetAuthorityTypes is a free data retrieval call binding the contract method 0xfdca3798.
//
// Solidity: function getAuthorityTypes(address[] authorities) view returns((address,uint256)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetAuthorityTypes(authorities []common.Address) ([]AccountResolverAuthType, error) {
	return _InstaDappDSAResolver.Contract.GetAuthorityTypes(&_InstaDappDSAResolver.CallOpts, authorities)
}

// GetContractCode is a free data retrieval call binding the contract method 0x28249577.
//
// Solidity: function getContractCode(address _addr) view returns(bytes o_code)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetContractCode(opts *bind.CallOpts, _addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getContractCode", _addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetContractCode is a free data retrieval call binding the contract method 0x28249577.
//
// Solidity: function getContractCode(address _addr) view returns(bytes o_code)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetContractCode(_addr common.Address) ([]byte, error) {
	return _InstaDappDSAResolver.Contract.GetContractCode(&_InstaDappDSAResolver.CallOpts, _addr)
}

// GetContractCode is a free data retrieval call binding the contract method 0x28249577.
//
// Solidity: function getContractCode(address _addr) view returns(bytes o_code)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetContractCode(_addr common.Address) ([]byte, error) {
	return _InstaDappDSAResolver.Contract.GetContractCode(&_InstaDappDSAResolver.CallOpts, _addr)
}

// GetEnabledConnectors is a free data retrieval call binding the contract method 0x6b1feeca.
//
// Solidity: function getEnabledConnectors() view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetEnabledConnectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getEnabledConnectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEnabledConnectors is a free data retrieval call binding the contract method 0x6b1feeca.
//
// Solidity: function getEnabledConnectors() view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetEnabledConnectors() ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetEnabledConnectors(&_InstaDappDSAResolver.CallOpts)
}

// GetEnabledConnectors is a free data retrieval call binding the contract method 0x6b1feeca.
//
// Solidity: function getEnabledConnectors() view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetEnabledConnectors() ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetEnabledConnectors(&_InstaDappDSAResolver.CallOpts)
}

// GetEnabledConnectorsData is a free data retrieval call binding the contract method 0x94176485.
//
// Solidity: function getEnabledConnectorsData() view returns((address,uint256,string)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetEnabledConnectorsData(opts *bind.CallOpts) ([]ConnectorsResolverConnectorsData, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getEnabledConnectorsData")

	if err != nil {
		return *new([]ConnectorsResolverConnectorsData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectorsResolverConnectorsData)).(*[]ConnectorsResolverConnectorsData)

	return out0, err

}

// GetEnabledConnectorsData is a free data retrieval call binding the contract method 0x94176485.
//
// Solidity: function getEnabledConnectorsData() view returns((address,uint256,string)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetEnabledConnectorsData() ([]ConnectorsResolverConnectorsData, error) {
	return _InstaDappDSAResolver.Contract.GetEnabledConnectorsData(&_InstaDappDSAResolver.CallOpts)
}

// GetEnabledConnectorsData is a free data retrieval call binding the contract method 0x94176485.
//
// Solidity: function getEnabledConnectorsData() view returns((address,uint256,string)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetEnabledConnectorsData() ([]ConnectorsResolverConnectorsData, error) {
	return _InstaDappDSAResolver.Contract.GetEnabledConnectorsData(&_InstaDappDSAResolver.CallOpts)
}

// GetID is a free data retrieval call binding the contract method 0x99f826a5.
//
// Solidity: function getID(address account) view returns(uint256 id)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetID(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getID", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetID is a free data retrieval call binding the contract method 0x99f826a5.
//
// Solidity: function getID(address account) view returns(uint256 id)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetID(account common.Address) (*big.Int, error) {
	return _InstaDappDSAResolver.Contract.GetID(&_InstaDappDSAResolver.CallOpts, account)
}

// GetID is a free data retrieval call binding the contract method 0x99f826a5.
//
// Solidity: function getID(address account) view returns(uint256 id)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetID(account common.Address) (*big.Int, error) {
	return _InstaDappDSAResolver.Contract.GetID(&_InstaDappDSAResolver.CallOpts, account)
}

// GetIDAuthorities is a free data retrieval call binding the contract method 0xc916d968.
//
// Solidity: function getIDAuthorities(uint256 id) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetIDAuthorities(opts *bind.CallOpts, id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getIDAuthorities", id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetIDAuthorities is a free data retrieval call binding the contract method 0xc916d968.
//
// Solidity: function getIDAuthorities(uint256 id) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetIDAuthorities(id *big.Int) ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetIDAuthorities(&_InstaDappDSAResolver.CallOpts, id)
}

// GetIDAuthorities is a free data retrieval call binding the contract method 0xc916d968.
//
// Solidity: function getIDAuthorities(uint256 id) view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetIDAuthorities(id *big.Int) ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetIDAuthorities(&_InstaDappDSAResolver.CallOpts, id)
}

// GetStaticConnectors is a free data retrieval call binding the contract method 0x49fef8dd.
//
// Solidity: function getStaticConnectors() view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetStaticConnectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getStaticConnectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStaticConnectors is a free data retrieval call binding the contract method 0x49fef8dd.
//
// Solidity: function getStaticConnectors() view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetStaticConnectors() ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetStaticConnectors(&_InstaDappDSAResolver.CallOpts)
}

// GetStaticConnectors is a free data retrieval call binding the contract method 0x49fef8dd.
//
// Solidity: function getStaticConnectors() view returns(address[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetStaticConnectors() ([]common.Address, error) {
	return _InstaDappDSAResolver.Contract.GetStaticConnectors(&_InstaDappDSAResolver.CallOpts)
}

// GetStaticConnectorsData is a free data retrieval call binding the contract method 0xaca3735e.
//
// Solidity: function getStaticConnectorsData() view returns((address,uint256,string)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GetStaticConnectorsData(opts *bind.CallOpts) ([]ConnectorsResolverConnectorsData, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "getStaticConnectorsData")

	if err != nil {
		return *new([]ConnectorsResolverConnectorsData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectorsResolverConnectorsData)).(*[]ConnectorsResolverConnectorsData)

	return out0, err

}

// GetStaticConnectorsData is a free data retrieval call binding the contract method 0xaca3735e.
//
// Solidity: function getStaticConnectorsData() view returns((address,uint256,string)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GetStaticConnectorsData() ([]ConnectorsResolverConnectorsData, error) {
	return _InstaDappDSAResolver.Contract.GetStaticConnectorsData(&_InstaDappDSAResolver.CallOpts)
}

// GetStaticConnectorsData is a free data retrieval call binding the contract method 0xaca3735e.
//
// Solidity: function getStaticConnectorsData() view returns((address,uint256,string)[])
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GetStaticConnectorsData() ([]ConnectorsResolverConnectorsData, error) {
	return _InstaDappDSAResolver.Contract.GetStaticConnectorsData(&_InstaDappDSAResolver.CallOpts)
}

// GnosisFactoryContracts is a free data retrieval call binding the contract method 0x67bbd907.
//
// Solidity: function gnosisFactoryContracts(uint256 ) view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) GnosisFactoryContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "gnosisFactoryContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GnosisFactoryContracts is a free data retrieval call binding the contract method 0x67bbd907.
//
// Solidity: function gnosisFactoryContracts(uint256 ) view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) GnosisFactoryContracts(arg0 *big.Int) (common.Address, error) {
	return _InstaDappDSAResolver.Contract.GnosisFactoryContracts(&_InstaDappDSAResolver.CallOpts, arg0)
}

// GnosisFactoryContracts is a free data retrieval call binding the contract method 0x67bbd907.
//
// Solidity: function gnosisFactoryContracts(uint256 ) view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) GnosisFactoryContracts(arg0 *big.Int) (common.Address, error) {
	return _InstaDappDSAResolver.Contract.GnosisFactoryContracts(&_InstaDappDSAResolver.CallOpts, arg0)
}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) Index(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "index")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) Index() (common.Address, error) {
	return _InstaDappDSAResolver.Contract.Index(&_InstaDappDSAResolver.CallOpts)
}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) Index() (common.Address, error) {
	return _InstaDappDSAResolver.Contract.Index(&_InstaDappDSAResolver.CallOpts)
}

// IsShield is a free data retrieval call binding the contract method 0x295c0d98.
//
// Solidity: function isShield(address account) view returns(bool shield)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) IsShield(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "isShield", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShield is a free data retrieval call binding the contract method 0x295c0d98.
//
// Solidity: function isShield(address account) view returns(bool shield)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) IsShield(account common.Address) (bool, error) {
	return _InstaDappDSAResolver.Contract.IsShield(&_InstaDappDSAResolver.CallOpts, account)
}

// IsShield is a free data retrieval call binding the contract method 0x295c0d98.
//
// Solidity: function isShield(address account) view returns(bool shield)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) IsShield(account common.Address) (bool, error) {
	return _InstaDappDSAResolver.Contract.IsShield(&_InstaDappDSAResolver.CallOpts, account)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) List(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "list")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) List() (common.Address, error) {
	return _InstaDappDSAResolver.Contract.List(&_InstaDappDSAResolver.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(address)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) List() (common.Address, error) {
	return _InstaDappDSAResolver.Contract.List(&_InstaDappDSAResolver.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) Name() (string, error) {
	return _InstaDappDSAResolver.Contract.Name(&_InstaDappDSAResolver.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) Name() (string, error) {
	return _InstaDappDSAResolver.Contract.Name(&_InstaDappDSAResolver.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_InstaDappDSAResolver *InstaDappDSAResolverCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappDSAResolver.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_InstaDappDSAResolver *InstaDappDSAResolverSession) Version() (*big.Int, error) {
	return _InstaDappDSAResolver.Contract.Version(&_InstaDappDSAResolver.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_InstaDappDSAResolver *InstaDappDSAResolverCallerSession) Version() (*big.Int, error) {
	return _InstaDappDSAResolver.Contract.Version(&_InstaDappDSAResolver.CallOpts)
}
