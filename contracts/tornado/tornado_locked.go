// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tornado

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

// GovernanceReceipt is an auto generated low-level Go binding around an user-defined struct.
type GovernanceReceipt struct {
	HasVoted bool
	Support  bool
	Votes    *big.Int
}

// TornadoLockedMetaData contains all meta data concerning the TornadoLocked contract.
var TornadoLockedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingRewardsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasCompLogic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userVaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"RewardUpdateFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RewardUpdateSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLOSING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTION_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTION_EXPIRATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSAL_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUORUM_VOTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Staking\",\"outputs\":[{\"internalType\":\"contractTornadoStakingRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_EXTEND_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTING_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"domains\",\"type\":\"bytes32[]\"}],\"name\":\"bulkResolve\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canWithdrawAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"from\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"castDelegatedVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"checkIfQuorumReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatedTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasCompensationVault\",\"outputs\":[{\"internalType\":\"contractIGasCompensationVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structGovernance.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAccountVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_torn\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestProposalIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockWithApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"proposeByDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnMultisigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"closingPeriod\",\"type\":\"uint256\"}],\"name\":\"setClosingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionDelay\",\"type\":\"uint256\"}],\"name\":\"setExecutionDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionExpiration\",\"type\":\"uint256\"}],\"name\":\"setExecutionExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasCompensationsLimit\",\"type\":\"uint256\"}],\"name\":\"setGasCompensations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumVotes\",\"type\":\"uint256\"}],\"name\":\"setQuorumVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"voteExtendTime\",\"type\":\"uint256\"}],\"name\":\"setVoteExtendTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingDelay\",\"type\":\"uint256\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumGovernance.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"torn\",\"outputs\":[{\"internalType\":\"contractTORN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userVault\",\"outputs\":[{\"internalType\":\"contractITornadoVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TornadoLockedABI is the input ABI used to generate the binding from.
// Deprecated: Use TornadoLockedMetaData.ABI instead.
var TornadoLockedABI = TornadoLockedMetaData.ABI

// TornadoLocked is an auto generated Go binding around an Ethereum contract.
type TornadoLocked struct {
	TornadoLockedCaller     // Read-only binding to the contract
	TornadoLockedTransactor // Write-only binding to the contract
	TornadoLockedFilterer   // Log filterer for contract events
}

// TornadoLockedCaller is an auto generated read-only Go binding around an Ethereum contract.
type TornadoLockedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TornadoLockedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TornadoLockedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TornadoLockedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TornadoLockedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TornadoLockedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TornadoLockedSession struct {
	Contract     *TornadoLocked    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TornadoLockedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TornadoLockedCallerSession struct {
	Contract *TornadoLockedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TornadoLockedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TornadoLockedTransactorSession struct {
	Contract     *TornadoLockedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TornadoLockedRaw is an auto generated low-level Go binding around an Ethereum contract.
type TornadoLockedRaw struct {
	Contract *TornadoLocked // Generic contract binding to access the raw methods on
}

// TornadoLockedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TornadoLockedCallerRaw struct {
	Contract *TornadoLockedCaller // Generic read-only contract binding to access the raw methods on
}

// TornadoLockedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TornadoLockedTransactorRaw struct {
	Contract *TornadoLockedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTornadoLocked creates a new instance of TornadoLocked, bound to a specific deployed contract.
func NewTornadoLocked(address common.Address, backend bind.ContractBackend) (*TornadoLocked, error) {
	contract, err := bindTornadoLocked(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TornadoLocked{TornadoLockedCaller: TornadoLockedCaller{contract: contract}, TornadoLockedTransactor: TornadoLockedTransactor{contract: contract}, TornadoLockedFilterer: TornadoLockedFilterer{contract: contract}}, nil
}

// NewTornadoLockedCaller creates a new read-only instance of TornadoLocked, bound to a specific deployed contract.
func NewTornadoLockedCaller(address common.Address, caller bind.ContractCaller) (*TornadoLockedCaller, error) {
	contract, err := bindTornadoLocked(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedCaller{contract: contract}, nil
}

// NewTornadoLockedTransactor creates a new write-only instance of TornadoLocked, bound to a specific deployed contract.
func NewTornadoLockedTransactor(address common.Address, transactor bind.ContractTransactor) (*TornadoLockedTransactor, error) {
	contract, err := bindTornadoLocked(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedTransactor{contract: contract}, nil
}

// NewTornadoLockedFilterer creates a new log filterer instance of TornadoLocked, bound to a specific deployed contract.
func NewTornadoLockedFilterer(address common.Address, filterer bind.ContractFilterer) (*TornadoLockedFilterer, error) {
	contract, err := bindTornadoLocked(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedFilterer{contract: contract}, nil
}

// bindTornadoLocked binds a generic wrapper to an already deployed contract.
func bindTornadoLocked(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TornadoLockedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TornadoLocked *TornadoLockedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TornadoLocked.Contract.TornadoLockedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TornadoLocked *TornadoLockedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoLocked.Contract.TornadoLockedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TornadoLocked *TornadoLockedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TornadoLocked.Contract.TornadoLockedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TornadoLocked *TornadoLockedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TornadoLocked.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TornadoLocked *TornadoLockedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoLocked.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TornadoLocked *TornadoLockedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TornadoLocked.Contract.contract.Transact(opts, method, params...)
}

// CLOSINGPERIOD is a free data retrieval call binding the contract method 0xce25d71c.
//
// Solidity: function CLOSING_PERIOD() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) CLOSINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "CLOSING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLOSINGPERIOD is a free data retrieval call binding the contract method 0xce25d71c.
//
// Solidity: function CLOSING_PERIOD() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) CLOSINGPERIOD() (*big.Int, error) {
	return _TornadoLocked.Contract.CLOSINGPERIOD(&_TornadoLocked.CallOpts)
}

// CLOSINGPERIOD is a free data retrieval call binding the contract method 0xce25d71c.
//
// Solidity: function CLOSING_PERIOD() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) CLOSINGPERIOD() (*big.Int, error) {
	return _TornadoLocked.Contract.CLOSINGPERIOD(&_TornadoLocked.CallOpts)
}

// EXECUTIONDELAY is a free data retrieval call binding the contract method 0x37f135d7.
//
// Solidity: function EXECUTION_DELAY() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) EXECUTIONDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "EXECUTION_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXECUTIONDELAY is a free data retrieval call binding the contract method 0x37f135d7.
//
// Solidity: function EXECUTION_DELAY() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) EXECUTIONDELAY() (*big.Int, error) {
	return _TornadoLocked.Contract.EXECUTIONDELAY(&_TornadoLocked.CallOpts)
}

// EXECUTIONDELAY is a free data retrieval call binding the contract method 0x37f135d7.
//
// Solidity: function EXECUTION_DELAY() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) EXECUTIONDELAY() (*big.Int, error) {
	return _TornadoLocked.Contract.EXECUTIONDELAY(&_TornadoLocked.CallOpts)
}

// EXECUTIONEXPIRATION is a free data retrieval call binding the contract method 0x6a661755.
//
// Solidity: function EXECUTION_EXPIRATION() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) EXECUTIONEXPIRATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "EXECUTION_EXPIRATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXECUTIONEXPIRATION is a free data retrieval call binding the contract method 0x6a661755.
//
// Solidity: function EXECUTION_EXPIRATION() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) EXECUTIONEXPIRATION() (*big.Int, error) {
	return _TornadoLocked.Contract.EXECUTIONEXPIRATION(&_TornadoLocked.CallOpts)
}

// EXECUTIONEXPIRATION is a free data retrieval call binding the contract method 0x6a661755.
//
// Solidity: function EXECUTION_EXPIRATION() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) EXECUTIONEXPIRATION() (*big.Int, error) {
	return _TornadoLocked.Contract.EXECUTIONEXPIRATION(&_TornadoLocked.CallOpts)
}

// PROPOSALTHRESHOLD is a free data retrieval call binding the contract method 0xa6c26603.
//
// Solidity: function PROPOSAL_THRESHOLD() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) PROPOSALTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "PROPOSAL_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROPOSALTHRESHOLD is a free data retrieval call binding the contract method 0xa6c26603.
//
// Solidity: function PROPOSAL_THRESHOLD() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) PROPOSALTHRESHOLD() (*big.Int, error) {
	return _TornadoLocked.Contract.PROPOSALTHRESHOLD(&_TornadoLocked.CallOpts)
}

// PROPOSALTHRESHOLD is a free data retrieval call binding the contract method 0xa6c26603.
//
// Solidity: function PROPOSAL_THRESHOLD() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) PROPOSALTHRESHOLD() (*big.Int, error) {
	return _TornadoLocked.Contract.PROPOSALTHRESHOLD(&_TornadoLocked.CallOpts)
}

// QUORUMVOTES is a free data retrieval call binding the contract method 0x671dd275.
//
// Solidity: function QUORUM_VOTES() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) QUORUMVOTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "QUORUM_VOTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUORUMVOTES is a free data retrieval call binding the contract method 0x671dd275.
//
// Solidity: function QUORUM_VOTES() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) QUORUMVOTES() (*big.Int, error) {
	return _TornadoLocked.Contract.QUORUMVOTES(&_TornadoLocked.CallOpts)
}

// QUORUMVOTES is a free data retrieval call binding the contract method 0x671dd275.
//
// Solidity: function QUORUM_VOTES() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) QUORUMVOTES() (*big.Int, error) {
	return _TornadoLocked.Contract.QUORUMVOTES(&_TornadoLocked.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0xf57df22e.
//
// Solidity: function Staking() view returns(address)
func (_TornadoLocked *TornadoLockedCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "Staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0xf57df22e.
//
// Solidity: function Staking() view returns(address)
func (_TornadoLocked *TornadoLockedSession) Staking() (common.Address, error) {
	return _TornadoLocked.Contract.Staking(&_TornadoLocked.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0xf57df22e.
//
// Solidity: function Staking() view returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) Staking() (common.Address, error) {
	return _TornadoLocked.Contract.Staking(&_TornadoLocked.CallOpts)
}

// VOTEEXTENDTIME is a free data retrieval call binding the contract method 0x587a6ecb.
//
// Solidity: function VOTE_EXTEND_TIME() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) VOTEEXTENDTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "VOTE_EXTEND_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VOTEEXTENDTIME is a free data retrieval call binding the contract method 0x587a6ecb.
//
// Solidity: function VOTE_EXTEND_TIME() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) VOTEEXTENDTIME() (*big.Int, error) {
	return _TornadoLocked.Contract.VOTEEXTENDTIME(&_TornadoLocked.CallOpts)
}

// VOTEEXTENDTIME is a free data retrieval call binding the contract method 0x587a6ecb.
//
// Solidity: function VOTE_EXTEND_TIME() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) VOTEEXTENDTIME() (*big.Int, error) {
	return _TornadoLocked.Contract.VOTEEXTENDTIME(&_TornadoLocked.CallOpts)
}

// VOTINGDELAY is a free data retrieval call binding the contract method 0xd6159fe5.
//
// Solidity: function VOTING_DELAY() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) VOTINGDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "VOTING_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VOTINGDELAY is a free data retrieval call binding the contract method 0xd6159fe5.
//
// Solidity: function VOTING_DELAY() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) VOTINGDELAY() (*big.Int, error) {
	return _TornadoLocked.Contract.VOTINGDELAY(&_TornadoLocked.CallOpts)
}

// VOTINGDELAY is a free data retrieval call binding the contract method 0xd6159fe5.
//
// Solidity: function VOTING_DELAY() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) VOTINGDELAY() (*big.Int, error) {
	return _TornadoLocked.Contract.VOTINGDELAY(&_TornadoLocked.CallOpts)
}

// VOTINGPERIOD is a free data retrieval call binding the contract method 0xb1610d7e.
//
// Solidity: function VOTING_PERIOD() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) VOTINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "VOTING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VOTINGPERIOD is a free data retrieval call binding the contract method 0xb1610d7e.
//
// Solidity: function VOTING_PERIOD() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) VOTINGPERIOD() (*big.Int, error) {
	return _TornadoLocked.Contract.VOTINGPERIOD(&_TornadoLocked.CallOpts)
}

// VOTINGPERIOD is a free data retrieval call binding the contract method 0xb1610d7e.
//
// Solidity: function VOTING_PERIOD() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) VOTINGPERIOD() (*big.Int, error) {
	return _TornadoLocked.Contract.VOTINGPERIOD(&_TornadoLocked.CallOpts)
}

// BulkResolve is a free data retrieval call binding the contract method 0xf9e54234.
//
// Solidity: function bulkResolve(bytes32[] domains) view returns(address[] result)
func (_TornadoLocked *TornadoLockedCaller) BulkResolve(opts *bind.CallOpts, domains [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "bulkResolve", domains)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// BulkResolve is a free data retrieval call binding the contract method 0xf9e54234.
//
// Solidity: function bulkResolve(bytes32[] domains) view returns(address[] result)
func (_TornadoLocked *TornadoLockedSession) BulkResolve(domains [][32]byte) ([]common.Address, error) {
	return _TornadoLocked.Contract.BulkResolve(&_TornadoLocked.CallOpts, domains)
}

// BulkResolve is a free data retrieval call binding the contract method 0xf9e54234.
//
// Solidity: function bulkResolve(bytes32[] domains) view returns(address[] result)
func (_TornadoLocked *TornadoLockedCallerSession) BulkResolve(domains [][32]byte) ([]common.Address, error) {
	return _TornadoLocked.Contract.BulkResolve(&_TornadoLocked.CallOpts, domains)
}

// CanWithdrawAfter is a free data retrieval call binding the contract method 0xa72edda3.
//
// Solidity: function canWithdrawAfter(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) CanWithdrawAfter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "canWithdrawAfter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CanWithdrawAfter is a free data retrieval call binding the contract method 0xa72edda3.
//
// Solidity: function canWithdrawAfter(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) CanWithdrawAfter(arg0 common.Address) (*big.Int, error) {
	return _TornadoLocked.Contract.CanWithdrawAfter(&_TornadoLocked.CallOpts, arg0)
}

// CanWithdrawAfter is a free data retrieval call binding the contract method 0xa72edda3.
//
// Solidity: function canWithdrawAfter(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) CanWithdrawAfter(arg0 common.Address) (*big.Int, error) {
	return _TornadoLocked.Contract.CanWithdrawAfter(&_TornadoLocked.CallOpts, arg0)
}

// CheckIfQuorumReached is a free data retrieval call binding the contract method 0x32687ec1.
//
// Solidity: function checkIfQuorumReached(uint256 proposalId) view returns(bool)
func (_TornadoLocked *TornadoLockedCaller) CheckIfQuorumReached(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "checkIfQuorumReached", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfQuorumReached is a free data retrieval call binding the contract method 0x32687ec1.
//
// Solidity: function checkIfQuorumReached(uint256 proposalId) view returns(bool)
func (_TornadoLocked *TornadoLockedSession) CheckIfQuorumReached(proposalId *big.Int) (bool, error) {
	return _TornadoLocked.Contract.CheckIfQuorumReached(&_TornadoLocked.CallOpts, proposalId)
}

// CheckIfQuorumReached is a free data retrieval call binding the contract method 0x32687ec1.
//
// Solidity: function checkIfQuorumReached(uint256 proposalId) view returns(bool)
func (_TornadoLocked *TornadoLockedCallerSession) CheckIfQuorumReached(proposalId *big.Int) (bool, error) {
	return _TornadoLocked.Contract.CheckIfQuorumReached(&_TornadoLocked.CallOpts, proposalId)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_TornadoLocked *TornadoLockedCaller) DelegatedTo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "delegatedTo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_TornadoLocked *TornadoLockedSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _TornadoLocked.Contract.DelegatedTo(&_TornadoLocked.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _TornadoLocked.Contract.DelegatedTo(&_TornadoLocked.CallOpts, arg0)
}

// GasCompensationVault is a free data retrieval call binding the contract method 0x8b34a960.
//
// Solidity: function gasCompensationVault() view returns(address)
func (_TornadoLocked *TornadoLockedCaller) GasCompensationVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "gasCompensationVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasCompensationVault is a free data retrieval call binding the contract method 0x8b34a960.
//
// Solidity: function gasCompensationVault() view returns(address)
func (_TornadoLocked *TornadoLockedSession) GasCompensationVault() (common.Address, error) {
	return _TornadoLocked.Contract.GasCompensationVault(&_TornadoLocked.CallOpts)
}

// GasCompensationVault is a free data retrieval call binding the contract method 0x8b34a960.
//
// Solidity: function gasCompensationVault() view returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) GasCompensationVault() (common.Address, error) {
	return _TornadoLocked.Contract.GasCompensationVault(&_TornadoLocked.CallOpts)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,bool,uint256))
func (_TornadoLocked *TornadoLockedCaller) GetReceipt(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "getReceipt", proposalId, voter)

	if err != nil {
		return *new(GovernanceReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(GovernanceReceipt)).(*GovernanceReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,bool,uint256))
func (_TornadoLocked *TornadoLockedSession) GetReceipt(proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	return _TornadoLocked.Contract.GetReceipt(&_TornadoLocked.CallOpts, proposalId, voter)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,bool,uint256))
func (_TornadoLocked *TornadoLockedCallerSession) GetReceipt(proposalId *big.Int, voter common.Address) (GovernanceReceipt, error) {
	return _TornadoLocked.Contract.GetReceipt(&_TornadoLocked.CallOpts, proposalId, voter)
}

// HasAccountVoted is a free data retrieval call binding the contract method 0xe525aa08.
//
// Solidity: function hasAccountVoted(uint256 proposalId, address account) view returns(bool)
func (_TornadoLocked *TornadoLockedCaller) HasAccountVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "hasAccountVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAccountVoted is a free data retrieval call binding the contract method 0xe525aa08.
//
// Solidity: function hasAccountVoted(uint256 proposalId, address account) view returns(bool)
func (_TornadoLocked *TornadoLockedSession) HasAccountVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _TornadoLocked.Contract.HasAccountVoted(&_TornadoLocked.CallOpts, proposalId, account)
}

// HasAccountVoted is a free data retrieval call binding the contract method 0xe525aa08.
//
// Solidity: function hasAccountVoted(uint256 proposalId, address account) view returns(bool)
func (_TornadoLocked *TornadoLockedCallerSession) HasAccountVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _TornadoLocked.Contract.HasAccountVoted(&_TornadoLocked.CallOpts, proposalId, account)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) LatestProposalIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "latestProposalIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _TornadoLocked.Contract.LatestProposalIds(&_TornadoLocked.CallOpts, arg0)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _TornadoLocked.Contract.LatestProposalIds(&_TornadoLocked.CallOpts, arg0)
}

// LockedBalance is a free data retrieval call binding the contract method 0x9ae697bf.
//
// Solidity: function lockedBalance(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) LockedBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "lockedBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalance is a free data retrieval call binding the contract method 0x9ae697bf.
//
// Solidity: function lockedBalance(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) LockedBalance(arg0 common.Address) (*big.Int, error) {
	return _TornadoLocked.Contract.LockedBalance(&_TornadoLocked.CallOpts, arg0)
}

// LockedBalance is a free data retrieval call binding the contract method 0x9ae697bf.
//
// Solidity: function lockedBalance(address ) view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) LockedBalance(arg0 common.Address) (*big.Int, error) {
	return _TornadoLocked.Contract.LockedBalance(&_TornadoLocked.CallOpts, arg0)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_TornadoLocked *TornadoLockedCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_TornadoLocked *TornadoLockedSession) ProposalCount() (*big.Int, error) {
	return _TornadoLocked.Contract.ProposalCount(&_TornadoLocked.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_TornadoLocked *TornadoLockedCallerSession) ProposalCount() (*big.Int, error) {
	return _TornadoLocked.Contract.ProposalCount(&_TornadoLocked.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address proposer, address target, uint256 startTime, uint256 endTime, uint256 forVotes, uint256 againstVotes, bool executed, bool extended)
func (_TornadoLocked *TornadoLockedCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Target       common.Address
	StartTime    *big.Int
	EndTime      *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Executed     bool
	Extended     bool
}, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		Target       common.Address
		StartTime    *big.Int
		EndTime      *big.Int
		ForVotes     *big.Int
		AgainstVotes *big.Int
		Executed     bool
		Extended     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Target = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AgainstVotes = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Extended = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address proposer, address target, uint256 startTime, uint256 endTime, uint256 forVotes, uint256 againstVotes, bool executed, bool extended)
func (_TornadoLocked *TornadoLockedSession) Proposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Target       common.Address
	StartTime    *big.Int
	EndTime      *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Executed     bool
	Extended     bool
}, error) {
	return _TornadoLocked.Contract.Proposals(&_TornadoLocked.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address proposer, address target, uint256 startTime, uint256 endTime, uint256 forVotes, uint256 againstVotes, bool executed, bool extended)
func (_TornadoLocked *TornadoLockedCallerSession) Proposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Target       common.Address
	StartTime    *big.Int
	EndTime      *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Executed     bool
	Extended     bool
}, error) {
	return _TornadoLocked.Contract.Proposals(&_TornadoLocked.CallOpts, arg0)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_TornadoLocked *TornadoLockedCaller) Resolve(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "resolve", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_TornadoLocked *TornadoLockedSession) Resolve(node [32]byte) (common.Address, error) {
	return _TornadoLocked.Contract.Resolve(&_TornadoLocked.CallOpts, node)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) Resolve(node [32]byte) (common.Address, error) {
	return _TornadoLocked.Contract.Resolve(&_TornadoLocked.CallOpts, node)
}

// ReturnMultisigAddress is a free data retrieval call binding the contract method 0x24b0435f.
//
// Solidity: function returnMultisigAddress() pure returns(address)
func (_TornadoLocked *TornadoLockedCaller) ReturnMultisigAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "returnMultisigAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReturnMultisigAddress is a free data retrieval call binding the contract method 0x24b0435f.
//
// Solidity: function returnMultisigAddress() pure returns(address)
func (_TornadoLocked *TornadoLockedSession) ReturnMultisigAddress() (common.Address, error) {
	return _TornadoLocked.Contract.ReturnMultisigAddress(&_TornadoLocked.CallOpts)
}

// ReturnMultisigAddress is a free data retrieval call binding the contract method 0x24b0435f.
//
// Solidity: function returnMultisigAddress() pure returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) ReturnMultisigAddress() (common.Address, error) {
	return _TornadoLocked.Contract.ReturnMultisigAddress(&_TornadoLocked.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_TornadoLocked *TornadoLockedCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_TornadoLocked *TornadoLockedSession) State(proposalId *big.Int) (uint8, error) {
	return _TornadoLocked.Contract.State(&_TornadoLocked.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_TornadoLocked *TornadoLockedCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _TornadoLocked.Contract.State(&_TornadoLocked.CallOpts, proposalId)
}

// Torn is a free data retrieval call binding the contract method 0xadf898a4.
//
// Solidity: function torn() view returns(address)
func (_TornadoLocked *TornadoLockedCaller) Torn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "torn")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Torn is a free data retrieval call binding the contract method 0xadf898a4.
//
// Solidity: function torn() view returns(address)
func (_TornadoLocked *TornadoLockedSession) Torn() (common.Address, error) {
	return _TornadoLocked.Contract.Torn(&_TornadoLocked.CallOpts)
}

// Torn is a free data retrieval call binding the contract method 0xadf898a4.
//
// Solidity: function torn() view returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) Torn() (common.Address, error) {
	return _TornadoLocked.Contract.Torn(&_TornadoLocked.CallOpts)
}

// UserVault is a free data retrieval call binding the contract method 0x9daafec7.
//
// Solidity: function userVault() view returns(address)
func (_TornadoLocked *TornadoLockedCaller) UserVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "userVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserVault is a free data retrieval call binding the contract method 0x9daafec7.
//
// Solidity: function userVault() view returns(address)
func (_TornadoLocked *TornadoLockedSession) UserVault() (common.Address, error) {
	return _TornadoLocked.Contract.UserVault(&_TornadoLocked.CallOpts)
}

// UserVault is a free data retrieval call binding the contract method 0x9daafec7.
//
// Solidity: function userVault() view returns(address)
func (_TornadoLocked *TornadoLockedCallerSession) UserVault() (common.Address, error) {
	return _TornadoLocked.Contract.UserVault(&_TornadoLocked.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_TornadoLocked *TornadoLockedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TornadoLocked.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_TornadoLocked *TornadoLockedSession) Version() (string, error) {
	return _TornadoLocked.Contract.Version(&_TornadoLocked.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_TornadoLocked *TornadoLockedCallerSession) Version() (string, error) {
	return _TornadoLocked.Contract.Version(&_TornadoLocked.CallOpts)
}

// CastDelegatedVote is a paid mutator transaction binding the contract method 0xb859f11b.
//
// Solidity: function castDelegatedVote(address[] from, uint256 proposalId, bool support) returns()
func (_TornadoLocked *TornadoLockedTransactor) CastDelegatedVote(opts *bind.TransactOpts, from []common.Address, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "castDelegatedVote", from, proposalId, support)
}

// CastDelegatedVote is a paid mutator transaction binding the contract method 0xb859f11b.
//
// Solidity: function castDelegatedVote(address[] from, uint256 proposalId, bool support) returns()
func (_TornadoLocked *TornadoLockedSession) CastDelegatedVote(from []common.Address, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _TornadoLocked.Contract.CastDelegatedVote(&_TornadoLocked.TransactOpts, from, proposalId, support)
}

// CastDelegatedVote is a paid mutator transaction binding the contract method 0xb859f11b.
//
// Solidity: function castDelegatedVote(address[] from, uint256 proposalId, bool support) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) CastDelegatedVote(from []common.Address, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _TornadoLocked.Contract.CastDelegatedVote(&_TornadoLocked.TransactOpts, from, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_TornadoLocked *TornadoLockedTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_TornadoLocked *TornadoLockedSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _TornadoLocked.Contract.CastVote(&_TornadoLocked.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x15373e3d.
//
// Solidity: function castVote(uint256 proposalId, bool support) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) CastVote(proposalId *big.Int, support bool) (*types.Transaction, error) {
	return _TornadoLocked.Contract.CastVote(&_TornadoLocked.TransactOpts, proposalId, support)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_TornadoLocked *TornadoLockedTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_TornadoLocked *TornadoLockedSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Delegate(&_TornadoLocked.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Delegate(&_TornadoLocked.TransactOpts, to)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_TornadoLocked *TornadoLockedTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_TornadoLocked *TornadoLockedSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Execute(&_TornadoLocked.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Execute(&_TornadoLocked.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 _torn) returns()
func (_TornadoLocked *TornadoLockedTransactor) Initialize(opts *bind.TransactOpts, _torn [32]byte) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "initialize", _torn)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 _torn) returns()
func (_TornadoLocked *TornadoLockedSession) Initialize(_torn [32]byte) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Initialize(&_TornadoLocked.TransactOpts, _torn)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 _torn) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Initialize(_torn [32]byte) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Initialize(&_TornadoLocked.TransactOpts, _torn)
}

// Lock is a paid mutator transaction binding the contract method 0xf0b76892.
//
// Solidity: function lock(address owner, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TornadoLocked *TornadoLockedTransactor) Lock(opts *bind.TransactOpts, owner common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "lock", owner, amount, deadline, v, r, s)
}

// Lock is a paid mutator transaction binding the contract method 0xf0b76892.
//
// Solidity: function lock(address owner, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TornadoLocked *TornadoLockedSession) Lock(owner common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Lock(&_TornadoLocked.TransactOpts, owner, amount, deadline, v, r, s)
}

// Lock is a paid mutator transaction binding the contract method 0xf0b76892.
//
// Solidity: function lock(address owner, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Lock(owner common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Lock(&_TornadoLocked.TransactOpts, owner, amount, deadline, v, r, s)
}

// LockWithApproval is a paid mutator transaction binding the contract method 0xb54426c8.
//
// Solidity: function lockWithApproval(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedTransactor) LockWithApproval(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "lockWithApproval", amount)
}

// LockWithApproval is a paid mutator transaction binding the contract method 0xb54426c8.
//
// Solidity: function lockWithApproval(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedSession) LockWithApproval(amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.LockWithApproval(&_TornadoLocked.TransactOpts, amount)
}

// LockWithApproval is a paid mutator transaction binding the contract method 0xb54426c8.
//
// Solidity: function lockWithApproval(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) LockWithApproval(amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.LockWithApproval(&_TornadoLocked.TransactOpts, amount)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
//
// Solidity: function propose(address target, string description) returns(uint256)
func (_TornadoLocked *TornadoLockedTransactor) Propose(opts *bind.TransactOpts, target common.Address, description string) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "propose", target, description)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
//
// Solidity: function propose(address target, string description) returns(uint256)
func (_TornadoLocked *TornadoLockedSession) Propose(target common.Address, description string) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Propose(&_TornadoLocked.TransactOpts, target, description)
}

// Propose is a paid mutator transaction binding the contract method 0xd6f0948c.
//
// Solidity: function propose(address target, string description) returns(uint256)
func (_TornadoLocked *TornadoLockedTransactorSession) Propose(target common.Address, description string) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Propose(&_TornadoLocked.TransactOpts, target, description)
}

// ProposeByDelegate is a paid mutator transaction binding the contract method 0x58e9fff0.
//
// Solidity: function proposeByDelegate(address from, address target, string description) returns(uint256)
func (_TornadoLocked *TornadoLockedTransactor) ProposeByDelegate(opts *bind.TransactOpts, from common.Address, target common.Address, description string) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "proposeByDelegate", from, target, description)
}

// ProposeByDelegate is a paid mutator transaction binding the contract method 0x58e9fff0.
//
// Solidity: function proposeByDelegate(address from, address target, string description) returns(uint256)
func (_TornadoLocked *TornadoLockedSession) ProposeByDelegate(from common.Address, target common.Address, description string) (*types.Transaction, error) {
	return _TornadoLocked.Contract.ProposeByDelegate(&_TornadoLocked.TransactOpts, from, target, description)
}

// ProposeByDelegate is a paid mutator transaction binding the contract method 0x58e9fff0.
//
// Solidity: function proposeByDelegate(address from, address target, string description) returns(uint256)
func (_TornadoLocked *TornadoLockedTransactorSession) ProposeByDelegate(from common.Address, target common.Address, description string) (*types.Transaction, error) {
	return _TornadoLocked.Contract.ProposeByDelegate(&_TornadoLocked.TransactOpts, from, target, description)
}

// SetClosingPeriod is a paid mutator transaction binding the contract method 0xc0c0e820.
//
// Solidity: function setClosingPeriod(uint256 closingPeriod) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetClosingPeriod(opts *bind.TransactOpts, closingPeriod *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setClosingPeriod", closingPeriod)
}

// SetClosingPeriod is a paid mutator transaction binding the contract method 0xc0c0e820.
//
// Solidity: function setClosingPeriod(uint256 closingPeriod) returns()
func (_TornadoLocked *TornadoLockedSession) SetClosingPeriod(closingPeriod *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetClosingPeriod(&_TornadoLocked.TransactOpts, closingPeriod)
}

// SetClosingPeriod is a paid mutator transaction binding the contract method 0xc0c0e820.
//
// Solidity: function setClosingPeriod(uint256 closingPeriod) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetClosingPeriod(closingPeriod *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetClosingPeriod(&_TornadoLocked.TransactOpts, closingPeriod)
}

// SetExecutionDelay is a paid mutator transaction binding the contract method 0xe4917d9f.
//
// Solidity: function setExecutionDelay(uint256 executionDelay) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetExecutionDelay(opts *bind.TransactOpts, executionDelay *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setExecutionDelay", executionDelay)
}

// SetExecutionDelay is a paid mutator transaction binding the contract method 0xe4917d9f.
//
// Solidity: function setExecutionDelay(uint256 executionDelay) returns()
func (_TornadoLocked *TornadoLockedSession) SetExecutionDelay(executionDelay *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetExecutionDelay(&_TornadoLocked.TransactOpts, executionDelay)
}

// SetExecutionDelay is a paid mutator transaction binding the contract method 0xe4917d9f.
//
// Solidity: function setExecutionDelay(uint256 executionDelay) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetExecutionDelay(executionDelay *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetExecutionDelay(&_TornadoLocked.TransactOpts, executionDelay)
}

// SetExecutionExpiration is a paid mutator transaction binding the contract method 0x9a9e3b6e.
//
// Solidity: function setExecutionExpiration(uint256 executionExpiration) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetExecutionExpiration(opts *bind.TransactOpts, executionExpiration *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setExecutionExpiration", executionExpiration)
}

// SetExecutionExpiration is a paid mutator transaction binding the contract method 0x9a9e3b6e.
//
// Solidity: function setExecutionExpiration(uint256 executionExpiration) returns()
func (_TornadoLocked *TornadoLockedSession) SetExecutionExpiration(executionExpiration *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetExecutionExpiration(&_TornadoLocked.TransactOpts, executionExpiration)
}

// SetExecutionExpiration is a paid mutator transaction binding the contract method 0x9a9e3b6e.
//
// Solidity: function setExecutionExpiration(uint256 executionExpiration) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetExecutionExpiration(executionExpiration *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetExecutionExpiration(&_TornadoLocked.TransactOpts, executionExpiration)
}

// SetGasCompensations is a paid mutator transaction binding the contract method 0xef3f8bb1.
//
// Solidity: function setGasCompensations(uint256 gasCompensationsLimit) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetGasCompensations(opts *bind.TransactOpts, gasCompensationsLimit *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setGasCompensations", gasCompensationsLimit)
}

// SetGasCompensations is a paid mutator transaction binding the contract method 0xef3f8bb1.
//
// Solidity: function setGasCompensations(uint256 gasCompensationsLimit) returns()
func (_TornadoLocked *TornadoLockedSession) SetGasCompensations(gasCompensationsLimit *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetGasCompensations(&_TornadoLocked.TransactOpts, gasCompensationsLimit)
}

// SetGasCompensations is a paid mutator transaction binding the contract method 0xef3f8bb1.
//
// Solidity: function setGasCompensations(uint256 gasCompensationsLimit) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetGasCompensations(gasCompensationsLimit *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetGasCompensations(&_TornadoLocked.TransactOpts, gasCompensationsLimit)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 proposalThreshold) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetProposalThreshold(opts *bind.TransactOpts, proposalThreshold *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setProposalThreshold", proposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 proposalThreshold) returns()
func (_TornadoLocked *TornadoLockedSession) SetProposalThreshold(proposalThreshold *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetProposalThreshold(&_TornadoLocked.TransactOpts, proposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 proposalThreshold) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetProposalThreshold(proposalThreshold *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetProposalThreshold(&_TornadoLocked.TransactOpts, proposalThreshold)
}

// SetQuorumVotes is a paid mutator transaction binding the contract method 0x02ec8f9e.
//
// Solidity: function setQuorumVotes(uint256 quorumVotes) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetQuorumVotes(opts *bind.TransactOpts, quorumVotes *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setQuorumVotes", quorumVotes)
}

// SetQuorumVotes is a paid mutator transaction binding the contract method 0x02ec8f9e.
//
// Solidity: function setQuorumVotes(uint256 quorumVotes) returns()
func (_TornadoLocked *TornadoLockedSession) SetQuorumVotes(quorumVotes *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetQuorumVotes(&_TornadoLocked.TransactOpts, quorumVotes)
}

// SetQuorumVotes is a paid mutator transaction binding the contract method 0x02ec8f9e.
//
// Solidity: function setQuorumVotes(uint256 quorumVotes) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetQuorumVotes(quorumVotes *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetQuorumVotes(&_TornadoLocked.TransactOpts, quorumVotes)
}

// SetVoteExtendTime is a paid mutator transaction binding the contract method 0x6dc2dc6c.
//
// Solidity: function setVoteExtendTime(uint256 voteExtendTime) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetVoteExtendTime(opts *bind.TransactOpts, voteExtendTime *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setVoteExtendTime", voteExtendTime)
}

// SetVoteExtendTime is a paid mutator transaction binding the contract method 0x6dc2dc6c.
//
// Solidity: function setVoteExtendTime(uint256 voteExtendTime) returns()
func (_TornadoLocked *TornadoLockedSession) SetVoteExtendTime(voteExtendTime *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetVoteExtendTime(&_TornadoLocked.TransactOpts, voteExtendTime)
}

// SetVoteExtendTime is a paid mutator transaction binding the contract method 0x6dc2dc6c.
//
// Solidity: function setVoteExtendTime(uint256 voteExtendTime) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetVoteExtendTime(voteExtendTime *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetVoteExtendTime(&_TornadoLocked.TransactOpts, voteExtendTime)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 votingDelay) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetVotingDelay(opts *bind.TransactOpts, votingDelay *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setVotingDelay", votingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 votingDelay) returns()
func (_TornadoLocked *TornadoLockedSession) SetVotingDelay(votingDelay *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetVotingDelay(&_TornadoLocked.TransactOpts, votingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 votingDelay) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetVotingDelay(votingDelay *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetVotingDelay(&_TornadoLocked.TransactOpts, votingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 votingPeriod) returns()
func (_TornadoLocked *TornadoLockedTransactor) SetVotingPeriod(opts *bind.TransactOpts, votingPeriod *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "setVotingPeriod", votingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 votingPeriod) returns()
func (_TornadoLocked *TornadoLockedSession) SetVotingPeriod(votingPeriod *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetVotingPeriod(&_TornadoLocked.TransactOpts, votingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 votingPeriod) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) SetVotingPeriod(votingPeriod *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.SetVotingPeriod(&_TornadoLocked.TransactOpts, votingPeriod)
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns()
func (_TornadoLocked *TornadoLockedTransactor) Undelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "undelegate")
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns()
func (_TornadoLocked *TornadoLockedSession) Undelegate() (*types.Transaction, error) {
	return _TornadoLocked.Contract.Undelegate(&_TornadoLocked.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Undelegate() (*types.Transaction, error) {
	return _TornadoLocked.Contract.Undelegate(&_TornadoLocked.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedTransactor) Unlock(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "unlock", amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedSession) Unlock(amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Unlock(&_TornadoLocked.TransactOpts, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Unlock(amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.Unlock(&_TornadoLocked.TransactOpts, amount)
}

// WithdrawFromHelper is a paid mutator transaction binding the contract method 0x932d5157.
//
// Solidity: function withdrawFromHelper(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedTransactor) WithdrawFromHelper(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.contract.Transact(opts, "withdrawFromHelper", amount)
}

// WithdrawFromHelper is a paid mutator transaction binding the contract method 0x932d5157.
//
// Solidity: function withdrawFromHelper(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedSession) WithdrawFromHelper(amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.WithdrawFromHelper(&_TornadoLocked.TransactOpts, amount)
}

// WithdrawFromHelper is a paid mutator transaction binding the contract method 0x932d5157.
//
// Solidity: function withdrawFromHelper(uint256 amount) returns()
func (_TornadoLocked *TornadoLockedTransactorSession) WithdrawFromHelper(amount *big.Int) (*types.Transaction, error) {
	return _TornadoLocked.Contract.WithdrawFromHelper(&_TornadoLocked.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TornadoLocked *TornadoLockedTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoLocked.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TornadoLocked *TornadoLockedSession) Receive() (*types.Transaction, error) {
	return _TornadoLocked.Contract.Receive(&_TornadoLocked.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TornadoLocked *TornadoLockedTransactorSession) Receive() (*types.Transaction, error) {
	return _TornadoLocked.Contract.Receive(&_TornadoLocked.TransactOpts)
}

// TornadoLockedDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the TornadoLocked contract.
type TornadoLockedDelegatedIterator struct {
	Event *TornadoLockedDelegated // Event containing the contract specifics and raw log

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
func (it *TornadoLockedDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedDelegated)
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
		it.Event = new(TornadoLockedDelegated)
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
func (it *TornadoLockedDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedDelegated represents a Delegated event raised by the TornadoLocked contract.
type TornadoLockedDelegated struct {
	Account common.Address
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x4bc154dd35d6a5cb9206482ecb473cdbf2473006d6bce728b9cc0741bcc59ea2.
//
// Solidity: event Delegated(address indexed account, address indexed to)
func (_TornadoLocked *TornadoLockedFilterer) FilterDelegated(opts *bind.FilterOpts, account []common.Address, to []common.Address) (*TornadoLockedDelegatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "Delegated", accountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedDelegatedIterator{contract: _TornadoLocked.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x4bc154dd35d6a5cb9206482ecb473cdbf2473006d6bce728b9cc0741bcc59ea2.
//
// Solidity: event Delegated(address indexed account, address indexed to)
func (_TornadoLocked *TornadoLockedFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *TornadoLockedDelegated, account []common.Address, to []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "Delegated", accountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedDelegated)
				if err := _TornadoLocked.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x4bc154dd35d6a5cb9206482ecb473cdbf2473006d6bce728b9cc0741bcc59ea2.
//
// Solidity: event Delegated(address indexed account, address indexed to)
func (_TornadoLocked *TornadoLockedFilterer) ParseDelegated(log types.Log) (*TornadoLockedDelegated, error) {
	event := new(TornadoLockedDelegated)
	if err := _TornadoLocked.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoLockedProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the TornadoLocked contract.
type TornadoLockedProposalCreatedIterator struct {
	Event *TornadoLockedProposalCreated // Event containing the contract specifics and raw log

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
func (it *TornadoLockedProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedProposalCreated)
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
		it.Event = new(TornadoLockedProposalCreated)
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
func (it *TornadoLockedProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedProposalCreated represents a ProposalCreated event raised by the TornadoLocked contract.
type TornadoLockedProposalCreated struct {
	Id          *big.Int
	Proposer    common.Address
	Target      common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x90ec05050aa23d54ba425e926fe646c318e85825bc400b13a46010abe86eb2f0.
//
// Solidity: event ProposalCreated(uint256 indexed id, address indexed proposer, address target, uint256 startTime, uint256 endTime, string description)
func (_TornadoLocked *TornadoLockedFilterer) FilterProposalCreated(opts *bind.FilterOpts, id []*big.Int, proposer []common.Address) (*TornadoLockedProposalCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "ProposalCreated", idRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedProposalCreatedIterator{contract: _TornadoLocked.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x90ec05050aa23d54ba425e926fe646c318e85825bc400b13a46010abe86eb2f0.
//
// Solidity: event ProposalCreated(uint256 indexed id, address indexed proposer, address target, uint256 startTime, uint256 endTime, string description)
func (_TornadoLocked *TornadoLockedFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *TornadoLockedProposalCreated, id []*big.Int, proposer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "ProposalCreated", idRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedProposalCreated)
				if err := _TornadoLocked.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x90ec05050aa23d54ba425e926fe646c318e85825bc400b13a46010abe86eb2f0.
//
// Solidity: event ProposalCreated(uint256 indexed id, address indexed proposer, address target, uint256 startTime, uint256 endTime, string description)
func (_TornadoLocked *TornadoLockedFilterer) ParseProposalCreated(log types.Log) (*TornadoLockedProposalCreated, error) {
	event := new(TornadoLockedProposalCreated)
	if err := _TornadoLocked.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoLockedProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the TornadoLocked contract.
type TornadoLockedProposalExecutedIterator struct {
	Event *TornadoLockedProposalExecuted // Event containing the contract specifics and raw log

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
func (it *TornadoLockedProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedProposalExecuted)
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
		it.Event = new(TornadoLockedProposalExecuted)
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
func (it *TornadoLockedProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedProposalExecuted represents a ProposalExecuted event raised by the TornadoLocked contract.
type TornadoLockedProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_TornadoLocked *TornadoLockedFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*TornadoLockedProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedProposalExecutedIterator{contract: _TornadoLocked.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_TornadoLocked *TornadoLockedFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *TornadoLockedProposalExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedProposalExecuted)
				if err := _TornadoLocked.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_TornadoLocked *TornadoLockedFilterer) ParseProposalExecuted(log types.Log) (*TornadoLockedProposalExecuted, error) {
	event := new(TornadoLockedProposalExecuted)
	if err := _TornadoLocked.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoLockedRewardUpdateFailedIterator is returned from FilterRewardUpdateFailed and is used to iterate over the raw logs and unpacked data for RewardUpdateFailed events raised by the TornadoLocked contract.
type TornadoLockedRewardUpdateFailedIterator struct {
	Event *TornadoLockedRewardUpdateFailed // Event containing the contract specifics and raw log

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
func (it *TornadoLockedRewardUpdateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedRewardUpdateFailed)
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
		it.Event = new(TornadoLockedRewardUpdateFailed)
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
func (it *TornadoLockedRewardUpdateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedRewardUpdateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedRewardUpdateFailed represents a RewardUpdateFailed event raised by the TornadoLocked contract.
type TornadoLockedRewardUpdateFailed struct {
	Account   common.Address
	ErrorData common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdateFailed is a free log retrieval operation binding the contract event 0x5a6216e80d86159dc87dcebfe519205477a94005b7d9d6bd313606450a5344f6.
//
// Solidity: event RewardUpdateFailed(address indexed account, bytes indexed errorData)
func (_TornadoLocked *TornadoLockedFilterer) FilterRewardUpdateFailed(opts *bind.FilterOpts, account []common.Address, errorData [][]byte) (*TornadoLockedRewardUpdateFailedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var errorDataRule []interface{}
	for _, errorDataItem := range errorData {
		errorDataRule = append(errorDataRule, errorDataItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "RewardUpdateFailed", accountRule, errorDataRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedRewardUpdateFailedIterator{contract: _TornadoLocked.contract, event: "RewardUpdateFailed", logs: logs, sub: sub}, nil
}

// WatchRewardUpdateFailed is a free log subscription operation binding the contract event 0x5a6216e80d86159dc87dcebfe519205477a94005b7d9d6bd313606450a5344f6.
//
// Solidity: event RewardUpdateFailed(address indexed account, bytes indexed errorData)
func (_TornadoLocked *TornadoLockedFilterer) WatchRewardUpdateFailed(opts *bind.WatchOpts, sink chan<- *TornadoLockedRewardUpdateFailed, account []common.Address, errorData [][]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var errorDataRule []interface{}
	for _, errorDataItem := range errorData {
		errorDataRule = append(errorDataRule, errorDataItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "RewardUpdateFailed", accountRule, errorDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedRewardUpdateFailed)
				if err := _TornadoLocked.contract.UnpackLog(event, "RewardUpdateFailed", log); err != nil {
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

// ParseRewardUpdateFailed is a log parse operation binding the contract event 0x5a6216e80d86159dc87dcebfe519205477a94005b7d9d6bd313606450a5344f6.
//
// Solidity: event RewardUpdateFailed(address indexed account, bytes indexed errorData)
func (_TornadoLocked *TornadoLockedFilterer) ParseRewardUpdateFailed(log types.Log) (*TornadoLockedRewardUpdateFailed, error) {
	event := new(TornadoLockedRewardUpdateFailed)
	if err := _TornadoLocked.contract.UnpackLog(event, "RewardUpdateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoLockedRewardUpdateSuccessfulIterator is returned from FilterRewardUpdateSuccessful and is used to iterate over the raw logs and unpacked data for RewardUpdateSuccessful events raised by the TornadoLocked contract.
type TornadoLockedRewardUpdateSuccessfulIterator struct {
	Event *TornadoLockedRewardUpdateSuccessful // Event containing the contract specifics and raw log

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
func (it *TornadoLockedRewardUpdateSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedRewardUpdateSuccessful)
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
		it.Event = new(TornadoLockedRewardUpdateSuccessful)
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
func (it *TornadoLockedRewardUpdateSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedRewardUpdateSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedRewardUpdateSuccessful represents a RewardUpdateSuccessful event raised by the TornadoLocked contract.
type TornadoLockedRewardUpdateSuccessful struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdateSuccessful is a free log retrieval operation binding the contract event 0x9b227b0c1ae308b34f72d4fdf9a1943fa769ff4814933595e7bc5230a117698b.
//
// Solidity: event RewardUpdateSuccessful(address indexed account)
func (_TornadoLocked *TornadoLockedFilterer) FilterRewardUpdateSuccessful(opts *bind.FilterOpts, account []common.Address) (*TornadoLockedRewardUpdateSuccessfulIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "RewardUpdateSuccessful", accountRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedRewardUpdateSuccessfulIterator{contract: _TornadoLocked.contract, event: "RewardUpdateSuccessful", logs: logs, sub: sub}, nil
}

// WatchRewardUpdateSuccessful is a free log subscription operation binding the contract event 0x9b227b0c1ae308b34f72d4fdf9a1943fa769ff4814933595e7bc5230a117698b.
//
// Solidity: event RewardUpdateSuccessful(address indexed account)
func (_TornadoLocked *TornadoLockedFilterer) WatchRewardUpdateSuccessful(opts *bind.WatchOpts, sink chan<- *TornadoLockedRewardUpdateSuccessful, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "RewardUpdateSuccessful", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedRewardUpdateSuccessful)
				if err := _TornadoLocked.contract.UnpackLog(event, "RewardUpdateSuccessful", log); err != nil {
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

// ParseRewardUpdateSuccessful is a log parse operation binding the contract event 0x9b227b0c1ae308b34f72d4fdf9a1943fa769ff4814933595e7bc5230a117698b.
//
// Solidity: event RewardUpdateSuccessful(address indexed account)
func (_TornadoLocked *TornadoLockedFilterer) ParseRewardUpdateSuccessful(log types.Log) (*TornadoLockedRewardUpdateSuccessful, error) {
	event := new(TornadoLockedRewardUpdateSuccessful)
	if err := _TornadoLocked.contract.UnpackLog(event, "RewardUpdateSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoLockedUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the TornadoLocked contract.
type TornadoLockedUndelegatedIterator struct {
	Event *TornadoLockedUndelegated // Event containing the contract specifics and raw log

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
func (it *TornadoLockedUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedUndelegated)
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
		it.Event = new(TornadoLockedUndelegated)
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
func (it *TornadoLockedUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedUndelegated represents a Undelegated event raised by the TornadoLocked contract.
type TornadoLockedUndelegated struct {
	Account common.Address
	From    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x1af5b1c85495b3618ea659a1ba256c8b8974b437297d3b914e321e086a28da72.
//
// Solidity: event Undelegated(address indexed account, address indexed from)
func (_TornadoLocked *TornadoLockedFilterer) FilterUndelegated(opts *bind.FilterOpts, account []common.Address, from []common.Address) (*TornadoLockedUndelegatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "Undelegated", accountRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedUndelegatedIterator{contract: _TornadoLocked.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x1af5b1c85495b3618ea659a1ba256c8b8974b437297d3b914e321e086a28da72.
//
// Solidity: event Undelegated(address indexed account, address indexed from)
func (_TornadoLocked *TornadoLockedFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *TornadoLockedUndelegated, account []common.Address, from []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "Undelegated", accountRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedUndelegated)
				if err := _TornadoLocked.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x1af5b1c85495b3618ea659a1ba256c8b8974b437297d3b914e321e086a28da72.
//
// Solidity: event Undelegated(address indexed account, address indexed from)
func (_TornadoLocked *TornadoLockedFilterer) ParseUndelegated(log types.Log) (*TornadoLockedUndelegated, error) {
	event := new(TornadoLockedUndelegated)
	if err := _TornadoLocked.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoLockedVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the TornadoLocked contract.
type TornadoLockedVotedIterator struct {
	Event *TornadoLockedVoted // Event containing the contract specifics and raw log

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
func (it *TornadoLockedVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoLockedVoted)
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
		it.Event = new(TornadoLockedVoted)
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
func (it *TornadoLockedVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoLockedVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoLockedVoted represents a Voted event raised by the TornadoLocked contract.
type TornadoLockedVoted struct {
	ProposalId *big.Int
	Voter      common.Address
	Support    bool
	Votes      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x7c2de587c00d75474a0c6c6fa96fd3b45dc974cd4e8a75f712bb84c950dce1b5.
//
// Solidity: event Voted(uint256 indexed proposalId, address indexed voter, bool indexed support, uint256 votes)
func (_TornadoLocked *TornadoLockedFilterer) FilterVoted(opts *bind.FilterOpts, proposalId []*big.Int, voter []common.Address, support []bool) (*TornadoLockedVotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var supportRule []interface{}
	for _, supportItem := range support {
		supportRule = append(supportRule, supportItem)
	}

	logs, sub, err := _TornadoLocked.contract.FilterLogs(opts, "Voted", proposalIdRule, voterRule, supportRule)
	if err != nil {
		return nil, err
	}
	return &TornadoLockedVotedIterator{contract: _TornadoLocked.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x7c2de587c00d75474a0c6c6fa96fd3b45dc974cd4e8a75f712bb84c950dce1b5.
//
// Solidity: event Voted(uint256 indexed proposalId, address indexed voter, bool indexed support, uint256 votes)
func (_TornadoLocked *TornadoLockedFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TornadoLockedVoted, proposalId []*big.Int, voter []common.Address, support []bool) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var supportRule []interface{}
	for _, supportItem := range support {
		supportRule = append(supportRule, supportItem)
	}

	logs, sub, err := _TornadoLocked.contract.WatchLogs(opts, "Voted", proposalIdRule, voterRule, supportRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoLockedVoted)
				if err := _TornadoLocked.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x7c2de587c00d75474a0c6c6fa96fd3b45dc974cd4e8a75f712bb84c950dce1b5.
//
// Solidity: event Voted(uint256 indexed proposalId, address indexed voter, bool indexed support, uint256 votes)
func (_TornadoLocked *TornadoLockedFilterer) ParseVoted(log types.Log) (*TornadoLockedVoted, error) {
	event := new(TornadoLockedVoted)
	if err := _TornadoLocked.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
