// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package makerdao

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

// MakerDaoCDPSMetaData contains all meta data concerning the MakerDaoCDPS contract.
var MakerDaoCDPSMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"getCdpsAsc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"urns\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ilks\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"getCdpsDesc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"urns\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ilks\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MakerDaoCDPSABI is the input ABI used to generate the binding from.
// Deprecated: Use MakerDaoCDPSMetaData.ABI instead.
var MakerDaoCDPSABI = MakerDaoCDPSMetaData.ABI

// MakerDaoCDPS is an auto generated Go binding around an Ethereum contract.
type MakerDaoCDPS struct {
	MakerDaoCDPSCaller     // Read-only binding to the contract
	MakerDaoCDPSTransactor // Write-only binding to the contract
	MakerDaoCDPSFilterer   // Log filterer for contract events
}

// MakerDaoCDPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type MakerDaoCDPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoCDPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MakerDaoCDPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoCDPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MakerDaoCDPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoCDPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MakerDaoCDPSSession struct {
	Contract     *MakerDaoCDPS     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MakerDaoCDPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MakerDaoCDPSCallerSession struct {
	Contract *MakerDaoCDPSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MakerDaoCDPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MakerDaoCDPSTransactorSession struct {
	Contract     *MakerDaoCDPSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MakerDaoCDPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type MakerDaoCDPSRaw struct {
	Contract *MakerDaoCDPS // Generic contract binding to access the raw methods on
}

// MakerDaoCDPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MakerDaoCDPSCallerRaw struct {
	Contract *MakerDaoCDPSCaller // Generic read-only contract binding to access the raw methods on
}

// MakerDaoCDPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MakerDaoCDPSTransactorRaw struct {
	Contract *MakerDaoCDPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMakerDaoCDPS creates a new instance of MakerDaoCDPS, bound to a specific deployed contract.
func NewMakerDaoCDPS(address common.Address, backend bind.ContractBackend) (*MakerDaoCDPS, error) {
	contract, err := bindMakerDaoCDPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MakerDaoCDPS{MakerDaoCDPSCaller: MakerDaoCDPSCaller{contract: contract}, MakerDaoCDPSTransactor: MakerDaoCDPSTransactor{contract: contract}, MakerDaoCDPSFilterer: MakerDaoCDPSFilterer{contract: contract}}, nil
}

// NewMakerDaoCDPSCaller creates a new read-only instance of MakerDaoCDPS, bound to a specific deployed contract.
func NewMakerDaoCDPSCaller(address common.Address, caller bind.ContractCaller) (*MakerDaoCDPSCaller, error) {
	contract, err := bindMakerDaoCDPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoCDPSCaller{contract: contract}, nil
}

// NewMakerDaoCDPSTransactor creates a new write-only instance of MakerDaoCDPS, bound to a specific deployed contract.
func NewMakerDaoCDPSTransactor(address common.Address, transactor bind.ContractTransactor) (*MakerDaoCDPSTransactor, error) {
	contract, err := bindMakerDaoCDPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoCDPSTransactor{contract: contract}, nil
}

// NewMakerDaoCDPSFilterer creates a new log filterer instance of MakerDaoCDPS, bound to a specific deployed contract.
func NewMakerDaoCDPSFilterer(address common.Address, filterer bind.ContractFilterer) (*MakerDaoCDPSFilterer, error) {
	contract, err := bindMakerDaoCDPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MakerDaoCDPSFilterer{contract: contract}, nil
}

// bindMakerDaoCDPS binds a generic wrapper to an already deployed contract.
func bindMakerDaoCDPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MakerDaoCDPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoCDPS *MakerDaoCDPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoCDPS.Contract.MakerDaoCDPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoCDPS *MakerDaoCDPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoCDPS.Contract.MakerDaoCDPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoCDPS *MakerDaoCDPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoCDPS.Contract.MakerDaoCDPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoCDPS *MakerDaoCDPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoCDPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoCDPS *MakerDaoCDPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoCDPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoCDPS *MakerDaoCDPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoCDPS.Contract.contract.Transact(opts, method, params...)
}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_MakerDaoCDPS *MakerDaoCDPSCaller) GetCdpsAsc(opts *bind.CallOpts, manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	var out []interface{}
	err := _MakerDaoCDPS.contract.Call(opts, &out, "getCdpsAsc", manager, guy)

	outstruct := new(struct {
		Ids  []*big.Int
		Urns []common.Address
		Ilks [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urns = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Ilks = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_MakerDaoCDPS *MakerDaoCDPSSession) GetCdpsAsc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _MakerDaoCDPS.Contract.GetCdpsAsc(&_MakerDaoCDPS.CallOpts, manager, guy)
}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_MakerDaoCDPS *MakerDaoCDPSCallerSession) GetCdpsAsc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _MakerDaoCDPS.Contract.GetCdpsAsc(&_MakerDaoCDPS.CallOpts, manager, guy)
}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_MakerDaoCDPS *MakerDaoCDPSCaller) GetCdpsDesc(opts *bind.CallOpts, manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	var out []interface{}
	err := _MakerDaoCDPS.contract.Call(opts, &out, "getCdpsDesc", manager, guy)

	outstruct := new(struct {
		Ids  []*big.Int
		Urns []common.Address
		Ilks [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urns = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Ilks = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_MakerDaoCDPS *MakerDaoCDPSSession) GetCdpsDesc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _MakerDaoCDPS.Contract.GetCdpsDesc(&_MakerDaoCDPS.CallOpts, manager, guy)
}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_MakerDaoCDPS *MakerDaoCDPSCallerSession) GetCdpsDesc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _MakerDaoCDPS.Contract.GetCdpsDesc(&_MakerDaoCDPS.CallOpts, manager, guy)
}
