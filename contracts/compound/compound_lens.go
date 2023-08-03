// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compound

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

// CompoundLensAccountLimits is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensAccountLimits struct {
	Markets   []common.Address
	Liquidity *big.Int
	Shortfall *big.Int
}

// CompoundLensCTokenBalances is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCTokenBalances struct {
	CToken               common.Address
	BalanceOf            *big.Int
	BorrowBalanceCurrent *big.Int
	BalanceOfUnderlying  *big.Int
	TokenBalance         *big.Int
	TokenAllowance       *big.Int
}

// CompoundLensCTokenMetadata is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCTokenMetadata struct {
	CToken                   common.Address
	ExchangeRateCurrent      *big.Int
	SupplyRatePerBlock       *big.Int
	BorrowRatePerBlock       *big.Int
	ReserveFactorMantissa    *big.Int
	TotalBorrows             *big.Int
	TotalReserves            *big.Int
	TotalSupply              *big.Int
	TotalCash                *big.Int
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	UnderlyingAssetAddress   common.Address
	CTokenDecimals           *big.Int
	UnderlyingDecimals       *big.Int
	CompSupplySpeed          *big.Int
	CompBorrowSpeed          *big.Int
	BorrowCap                *big.Int
}

// CompoundLensCTokenUnderlyingPrice is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCTokenUnderlyingPrice struct {
	CToken          common.Address
	UnderlyingPrice *big.Int
}

// CompoundLensCompBalanceMetadata is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCompBalanceMetadata struct {
	Balance  *big.Int
	Votes    *big.Int
	Delegate common.Address
}

// CompoundLensCompBalanceMetadataExt is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCompBalanceMetadataExt struct {
	Balance   *big.Int
	Votes     *big.Int
	Delegate  common.Address
	Allocated *big.Int
}

// CompoundLensCompVotes is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCompVotes struct {
	BlockNumber *big.Int
	Votes       *big.Int
}

// CompoundLensGovBravoProposal is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovBravoProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Eta          *big.Int
	Targets      []common.Address
	Values       []*big.Int
	Signatures   []string
	Calldatas    [][]byte
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	AbstainVotes *big.Int
	Canceled     bool
	Executed     bool
}

// CompoundLensGovBravoReceipt is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovBravoReceipt struct {
	ProposalId *big.Int
	HasVoted   bool
	Support    uint8
	Votes      *big.Int
}

// CompoundLensGovProposal is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Eta          *big.Int
	Targets      []common.Address
	Values       []*big.Int
	Signatures   []string
	Calldatas    [][]byte
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
}

// CompoundLensGovReceipt is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovReceipt struct {
	ProposalId *big.Int
	HasVoted   bool
	Support    bool
	Votes      *big.Int
}

// CompoundLensMetaData contains all meta data concerning the CompoundLens contract.
var CompoundLensMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cTokenBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalanceCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAllowance\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenBalances\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cTokenBalancesAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalanceCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAllowance\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenBalances[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"cTokenMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRateCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compSupplySpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compBorrowSpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"cTokenMetadataAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRateCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compSupplySpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compBorrowSpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenMetadata[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"cTokenUnderlyingPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenUnderlyingPrice\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"cTokenUnderlyingPriceAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenUnderlyingPrice[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLimits\",\"outputs\":[{\"components\":[{\"internalType\":\"contractCToken[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfall\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.AccountLimits\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractComp\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCompBalanceMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structCompoundLens.CompBalanceMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComp\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCompBalanceMetadataExt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CompBalanceMetadataExt\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractComp\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"blockNumbers\",\"type\":\"uint32[]\"}],\"name\":\"getCompVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CompVotes[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractGovernorBravoInterface\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovBravoProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structCompoundLens.GovBravoProposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractGovernorBravoInterface\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovBravoReceipts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structCompoundLens.GovBravoReceipt[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractGovernorAlpha\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structCompoundLens.GovProposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractGovernorAlpha\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovReceipts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structCompoundLens.GovReceipt[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CompoundLensABI is the input ABI used to generate the binding from.
// Deprecated: Use CompoundLensMetaData.ABI instead.
var CompoundLensABI = CompoundLensMetaData.ABI

// CompoundLens is an auto generated Go binding around an Ethereum contract.
type CompoundLens struct {
	CompoundLensCaller     // Read-only binding to the contract
	CompoundLensTransactor // Write-only binding to the contract
	CompoundLensFilterer   // Log filterer for contract events
}

// CompoundLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundLensSession struct {
	Contract     *CompoundLens     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundLensCallerSession struct {
	Contract *CompoundLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CompoundLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundLensTransactorSession struct {
	Contract     *CompoundLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CompoundLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundLensRaw struct {
	Contract *CompoundLens // Generic contract binding to access the raw methods on
}

// CompoundLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundLensCallerRaw struct {
	Contract *CompoundLensCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundLensTransactorRaw struct {
	Contract *CompoundLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundLens creates a new instance of CompoundLens, bound to a specific deployed contract.
func NewCompoundLens(address common.Address, backend bind.ContractBackend) (*CompoundLens, error) {
	contract, err := bindCompoundLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundLens{CompoundLensCaller: CompoundLensCaller{contract: contract}, CompoundLensTransactor: CompoundLensTransactor{contract: contract}, CompoundLensFilterer: CompoundLensFilterer{contract: contract}}, nil
}

// NewCompoundLensCaller creates a new read-only instance of CompoundLens, bound to a specific deployed contract.
func NewCompoundLensCaller(address common.Address, caller bind.ContractCaller) (*CompoundLensCaller, error) {
	contract, err := bindCompoundLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundLensCaller{contract: contract}, nil
}

// NewCompoundLensTransactor creates a new write-only instance of CompoundLens, bound to a specific deployed contract.
func NewCompoundLensTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundLensTransactor, error) {
	contract, err := bindCompoundLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundLensTransactor{contract: contract}, nil
}

// NewCompoundLensFilterer creates a new log filterer instance of CompoundLens, bound to a specific deployed contract.
func NewCompoundLensFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundLensFilterer, error) {
	contract, err := bindCompoundLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundLensFilterer{contract: contract}, nil
}

// bindCompoundLens binds a generic wrapper to an already deployed contract.
func bindCompoundLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompoundLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundLens *CompoundLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundLens.Contract.CompoundLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundLens *CompoundLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundLens.Contract.CompoundLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundLens *CompoundLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundLens.Contract.CompoundLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundLens *CompoundLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundLens *CompoundLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundLens *CompoundLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundLens.Contract.contract.Transact(opts, method, params...)
}

// GetCompBalanceMetadata is a free data retrieval call binding the contract method 0x416405d7.
//
// Solidity: function getCompBalanceMetadata(address comp, address account) view returns((uint256,uint256,address))
func (_CompoundLens *CompoundLensCaller) GetCompBalanceMetadata(opts *bind.CallOpts, comp common.Address, account common.Address) (CompoundLensCompBalanceMetadata, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getCompBalanceMetadata", comp, account)

	if err != nil {
		return *new(CompoundLensCompBalanceMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(CompoundLensCompBalanceMetadata)).(*CompoundLensCompBalanceMetadata)

	return out0, err

}

// GetCompBalanceMetadata is a free data retrieval call binding the contract method 0x416405d7.
//
// Solidity: function getCompBalanceMetadata(address comp, address account) view returns((uint256,uint256,address))
func (_CompoundLens *CompoundLensSession) GetCompBalanceMetadata(comp common.Address, account common.Address) (CompoundLensCompBalanceMetadata, error) {
	return _CompoundLens.Contract.GetCompBalanceMetadata(&_CompoundLens.CallOpts, comp, account)
}

// GetCompBalanceMetadata is a free data retrieval call binding the contract method 0x416405d7.
//
// Solidity: function getCompBalanceMetadata(address comp, address account) view returns((uint256,uint256,address))
func (_CompoundLens *CompoundLensCallerSession) GetCompBalanceMetadata(comp common.Address, account common.Address) (CompoundLensCompBalanceMetadata, error) {
	return _CompoundLens.Contract.GetCompBalanceMetadata(&_CompoundLens.CallOpts, comp, account)
}

// GetCompBalanceMetadataExt is a free data retrieval call binding the contract method 0x1ea63741.
//
// Solidity: function getCompBalanceMetadataExt(address comp, address comptroller, address account) view returns((uint256,uint256,address,uint256))
func (_CompoundLens *CompoundLensCaller) GetCompBalanceMetadataExt(opts *bind.CallOpts, comp common.Address, comptroller common.Address, account common.Address) (CompoundLensCompBalanceMetadataExt, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getCompBalanceMetadataExt", comp, comptroller, account)

	if err != nil {
		return *new(CompoundLensCompBalanceMetadataExt), err
	}

	out0 := *abi.ConvertType(out[0], new(CompoundLensCompBalanceMetadataExt)).(*CompoundLensCompBalanceMetadataExt)

	return out0, err

}

// GetCompBalanceMetadataExt is a free data retrieval call binding the contract method 0x1ea63741.
//
// Solidity: function getCompBalanceMetadataExt(address comp, address comptroller, address account) view returns((uint256,uint256,address,uint256))
func (_CompoundLens *CompoundLensSession) GetCompBalanceMetadataExt(comp common.Address, comptroller common.Address, account common.Address) (CompoundLensCompBalanceMetadataExt, error) {
	return _CompoundLens.Contract.GetCompBalanceMetadataExt(&_CompoundLens.CallOpts, comp, comptroller, account)
}

// GetCompBalanceMetadataExt is a free data retrieval call binding the contract method 0x1ea63741.
//
// Solidity: function getCompBalanceMetadataExt(address comp, address comptroller, address account) view returns((uint256,uint256,address,uint256))
func (_CompoundLens *CompoundLensCallerSession) GetCompBalanceMetadataExt(comp common.Address, comptroller common.Address, account common.Address) (CompoundLensCompBalanceMetadataExt, error) {
	return _CompoundLens.Contract.GetCompBalanceMetadataExt(&_CompoundLens.CallOpts, comp, comptroller, account)
}

// GetCompVotes is a free data retrieval call binding the contract method 0x59564219.
//
// Solidity: function getCompVotes(address comp, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_CompoundLens *CompoundLensCaller) GetCompVotes(opts *bind.CallOpts, comp common.Address, account common.Address, blockNumbers []uint32) ([]CompoundLensCompVotes, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getCompVotes", comp, account, blockNumbers)

	if err != nil {
		return *new([]CompoundLensCompVotes), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensCompVotes)).(*[]CompoundLensCompVotes)

	return out0, err

}

// GetCompVotes is a free data retrieval call binding the contract method 0x59564219.
//
// Solidity: function getCompVotes(address comp, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_CompoundLens *CompoundLensSession) GetCompVotes(comp common.Address, account common.Address, blockNumbers []uint32) ([]CompoundLensCompVotes, error) {
	return _CompoundLens.Contract.GetCompVotes(&_CompoundLens.CallOpts, comp, account, blockNumbers)
}

// GetCompVotes is a free data retrieval call binding the contract method 0x59564219.
//
// Solidity: function getCompVotes(address comp, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_CompoundLens *CompoundLensCallerSession) GetCompVotes(comp common.Address, account common.Address, blockNumbers []uint32) ([]CompoundLensCompVotes, error) {
	return _CompoundLens.Contract.GetCompVotes(&_CompoundLens.CallOpts, comp, account, blockNumbers)
}

// GetGovBravoProposals is a free data retrieval call binding the contract method 0x75d80e90.
//
// Solidity: function getGovBravoProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,uint256,bool,bool)[])
func (_CompoundLens *CompoundLensCaller) GetGovBravoProposals(opts *bind.CallOpts, governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoProposal, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getGovBravoProposals", governor, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovBravoProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovBravoProposal)).(*[]CompoundLensGovBravoProposal)

	return out0, err

}

// GetGovBravoProposals is a free data retrieval call binding the contract method 0x75d80e90.
//
// Solidity: function getGovBravoProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,uint256,bool,bool)[])
func (_CompoundLens *CompoundLensSession) GetGovBravoProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoProposal, error) {
	return _CompoundLens.Contract.GetGovBravoProposals(&_CompoundLens.CallOpts, governor, proposalIds)
}

// GetGovBravoProposals is a free data retrieval call binding the contract method 0x75d80e90.
//
// Solidity: function getGovBravoProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,uint256,bool,bool)[])
func (_CompoundLens *CompoundLensCallerSession) GetGovBravoProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoProposal, error) {
	return _CompoundLens.Contract.GetGovBravoProposals(&_CompoundLens.CallOpts, governor, proposalIds)
}

// GetGovBravoReceipts is a free data retrieval call binding the contract method 0x43c811cc.
//
// Solidity: function getGovBravoReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,uint8,uint96)[])
func (_CompoundLens *CompoundLensCaller) GetGovBravoReceipts(opts *bind.CallOpts, governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoReceipt, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getGovBravoReceipts", governor, voter, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovBravoReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovBravoReceipt)).(*[]CompoundLensGovBravoReceipt)

	return out0, err

}

// GetGovBravoReceipts is a free data retrieval call binding the contract method 0x43c811cc.
//
// Solidity: function getGovBravoReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,uint8,uint96)[])
func (_CompoundLens *CompoundLensSession) GetGovBravoReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoReceipt, error) {
	return _CompoundLens.Contract.GetGovBravoReceipts(&_CompoundLens.CallOpts, governor, voter, proposalIds)
}

// GetGovBravoReceipts is a free data retrieval call binding the contract method 0x43c811cc.
//
// Solidity: function getGovBravoReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,uint8,uint96)[])
func (_CompoundLens *CompoundLensCallerSession) GetGovBravoReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoReceipt, error) {
	return _CompoundLens.Contract.GetGovBravoReceipts(&_CompoundLens.CallOpts, governor, voter, proposalIds)
}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_CompoundLens *CompoundLensCaller) GetGovProposals(opts *bind.CallOpts, governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovProposal, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getGovProposals", governor, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovProposal)).(*[]CompoundLensGovProposal)

	return out0, err

}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_CompoundLens *CompoundLensSession) GetGovProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovProposal, error) {
	return _CompoundLens.Contract.GetGovProposals(&_CompoundLens.CallOpts, governor, proposalIds)
}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_CompoundLens *CompoundLensCallerSession) GetGovProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovProposal, error) {
	return _CompoundLens.Contract.GetGovProposals(&_CompoundLens.CallOpts, governor, proposalIds)
}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_CompoundLens *CompoundLensCaller) GetGovReceipts(opts *bind.CallOpts, governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovReceipt, error) {
	var out []interface{}
	err := _CompoundLens.contract.Call(opts, &out, "getGovReceipts", governor, voter, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovReceipt)).(*[]CompoundLensGovReceipt)

	return out0, err

}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_CompoundLens *CompoundLensSession) GetGovReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovReceipt, error) {
	return _CompoundLens.Contract.GetGovReceipts(&_CompoundLens.CallOpts, governor, voter, proposalIds)
}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_CompoundLens *CompoundLensCallerSession) GetGovReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovReceipt, error) {
	return _CompoundLens.Contract.GetGovReceipts(&_CompoundLens.CallOpts, governor, voter, proposalIds)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0xbdf950c9.
//
// Solidity: function cTokenBalances(address cToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_CompoundLens *CompoundLensTransactor) CTokenBalances(opts *bind.TransactOpts, cToken common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "cTokenBalances", cToken, account)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0xbdf950c9.
//
// Solidity: function cTokenBalances(address cToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_CompoundLens *CompoundLensSession) CTokenBalances(cToken common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenBalances(&_CompoundLens.TransactOpts, cToken, account)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0xbdf950c9.
//
// Solidity: function cTokenBalances(address cToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_CompoundLens *CompoundLensTransactorSession) CTokenBalances(cToken common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenBalances(&_CompoundLens.TransactOpts, cToken, account)
}

// CTokenBalancesAll is a paid mutator transaction binding the contract method 0x0972bf8b.
//
// Solidity: function cTokenBalancesAll(address[] cTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_CompoundLens *CompoundLensTransactor) CTokenBalancesAll(opts *bind.TransactOpts, cTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "cTokenBalancesAll", cTokens, account)
}

// CTokenBalancesAll is a paid mutator transaction binding the contract method 0x0972bf8b.
//
// Solidity: function cTokenBalancesAll(address[] cTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_CompoundLens *CompoundLensSession) CTokenBalancesAll(cTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenBalancesAll(&_CompoundLens.TransactOpts, cTokens, account)
}

// CTokenBalancesAll is a paid mutator transaction binding the contract method 0x0972bf8b.
//
// Solidity: function cTokenBalancesAll(address[] cTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_CompoundLens *CompoundLensTransactorSession) CTokenBalancesAll(cTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenBalancesAll(&_CompoundLens.TransactOpts, cTokens, account)
}

// CTokenMetadata is a paid mutator transaction binding the contract method 0x158eca8b.
//
// Solidity: function cTokenMetadata(address cToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_CompoundLens *CompoundLensTransactor) CTokenMetadata(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "cTokenMetadata", cToken)
}

// CTokenMetadata is a paid mutator transaction binding the contract method 0x158eca8b.
//
// Solidity: function cTokenMetadata(address cToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_CompoundLens *CompoundLensSession) CTokenMetadata(cToken common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenMetadata(&_CompoundLens.TransactOpts, cToken)
}

// CTokenMetadata is a paid mutator transaction binding the contract method 0x158eca8b.
//
// Solidity: function cTokenMetadata(address cToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_CompoundLens *CompoundLensTransactorSession) CTokenMetadata(cToken common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenMetadata(&_CompoundLens.TransactOpts, cToken)
}

// CTokenMetadataAll is a paid mutator transaction binding the contract method 0x4b70d84b.
//
// Solidity: function cTokenMetadataAll(address[] cTokens) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256)[])
func (_CompoundLens *CompoundLensTransactor) CTokenMetadataAll(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "cTokenMetadataAll", cTokens)
}

// CTokenMetadataAll is a paid mutator transaction binding the contract method 0x4b70d84b.
//
// Solidity: function cTokenMetadataAll(address[] cTokens) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256)[])
func (_CompoundLens *CompoundLensSession) CTokenMetadataAll(cTokens []common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenMetadataAll(&_CompoundLens.TransactOpts, cTokens)
}

// CTokenMetadataAll is a paid mutator transaction binding the contract method 0x4b70d84b.
//
// Solidity: function cTokenMetadataAll(address[] cTokens) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256)[])
func (_CompoundLens *CompoundLensTransactorSession) CTokenMetadataAll(cTokens []common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenMetadataAll(&_CompoundLens.TransactOpts, cTokens)
}

// CTokenUnderlyingPrice is a paid mutator transaction binding the contract method 0xc5ae5934.
//
// Solidity: function cTokenUnderlyingPrice(address cToken) returns((address,uint256))
func (_CompoundLens *CompoundLensTransactor) CTokenUnderlyingPrice(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "cTokenUnderlyingPrice", cToken)
}

// CTokenUnderlyingPrice is a paid mutator transaction binding the contract method 0xc5ae5934.
//
// Solidity: function cTokenUnderlyingPrice(address cToken) returns((address,uint256))
func (_CompoundLens *CompoundLensSession) CTokenUnderlyingPrice(cToken common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenUnderlyingPrice(&_CompoundLens.TransactOpts, cToken)
}

// CTokenUnderlyingPrice is a paid mutator transaction binding the contract method 0xc5ae5934.
//
// Solidity: function cTokenUnderlyingPrice(address cToken) returns((address,uint256))
func (_CompoundLens *CompoundLensTransactorSession) CTokenUnderlyingPrice(cToken common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenUnderlyingPrice(&_CompoundLens.TransactOpts, cToken)
}

// CTokenUnderlyingPriceAll is a paid mutator transaction binding the contract method 0x2b2d5ed6.
//
// Solidity: function cTokenUnderlyingPriceAll(address[] cTokens) returns((address,uint256)[])
func (_CompoundLens *CompoundLensTransactor) CTokenUnderlyingPriceAll(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "cTokenUnderlyingPriceAll", cTokens)
}

// CTokenUnderlyingPriceAll is a paid mutator transaction binding the contract method 0x2b2d5ed6.
//
// Solidity: function cTokenUnderlyingPriceAll(address[] cTokens) returns((address,uint256)[])
func (_CompoundLens *CompoundLensSession) CTokenUnderlyingPriceAll(cTokens []common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenUnderlyingPriceAll(&_CompoundLens.TransactOpts, cTokens)
}

// CTokenUnderlyingPriceAll is a paid mutator transaction binding the contract method 0x2b2d5ed6.
//
// Solidity: function cTokenUnderlyingPriceAll(address[] cTokens) returns((address,uint256)[])
func (_CompoundLens *CompoundLensTransactorSession) CTokenUnderlyingPriceAll(cTokens []common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.CTokenUnderlyingPriceAll(&_CompoundLens.TransactOpts, cTokens)
}

// GetAccountLimits is a paid mutator transaction binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) returns((address[],uint256,uint256))
func (_CompoundLens *CompoundLensTransactor) GetAccountLimits(opts *bind.TransactOpts, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.contract.Transact(opts, "getAccountLimits", comptroller, account)
}

// GetAccountLimits is a paid mutator transaction binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) returns((address[],uint256,uint256))
func (_CompoundLens *CompoundLensSession) GetAccountLimits(comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.GetAccountLimits(&_CompoundLens.TransactOpts, comptroller, account)
}

// GetAccountLimits is a paid mutator transaction binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) returns((address[],uint256,uint256))
func (_CompoundLens *CompoundLensTransactorSession) GetAccountLimits(comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _CompoundLens.Contract.GetAccountLimits(&_CompoundLens.TransactOpts, comptroller, account)
}
