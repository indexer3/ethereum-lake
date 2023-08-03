// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package double

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

// DoubleDoNFTMetaData contains all meta data concerning the DoubleDoNFT contract.
var DoubleDoNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AddBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"RemoveBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"UpdateUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"UpgradedNFT\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MINT_WHITELIST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"addBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"team\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"teamMint\",\"type\":\"uint256[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"airdropUsers\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"expireDays\",\"type\":\"uint64[]\"}],\"name\":\"airdropTrialPass\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpiredAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getPassLevel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hiddenMetadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageURIExtension\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUniqueImage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWhitelistMintActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newImageURIExtension\",\"type\":\"string\"}],\"name\":\"setImageExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_upgradeAmount\",\"type\":\"uint256\"}],\"name\":\"setUpgradeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsUniqueImage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleReveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleWhitelistMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trialPassIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"upgradeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DoubleDoNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use DoubleDoNFTMetaData.ABI instead.
var DoubleDoNFTABI = DoubleDoNFTMetaData.ABI

// DoubleDoNFT is an auto generated Go binding around an Ethereum contract.
type DoubleDoNFT struct {
	DoubleDoNFTCaller     // Read-only binding to the contract
	DoubleDoNFTTransactor // Write-only binding to the contract
	DoubleDoNFTFilterer   // Log filterer for contract events
}

// DoubleDoNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoubleDoNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleDoNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoubleDoNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleDoNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoubleDoNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleDoNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoubleDoNFTSession struct {
	Contract     *DoubleDoNFT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoubleDoNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoubleDoNFTCallerSession struct {
	Contract *DoubleDoNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DoubleDoNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoubleDoNFTTransactorSession struct {
	Contract     *DoubleDoNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DoubleDoNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoubleDoNFTRaw struct {
	Contract *DoubleDoNFT // Generic contract binding to access the raw methods on
}

// DoubleDoNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoubleDoNFTCallerRaw struct {
	Contract *DoubleDoNFTCaller // Generic read-only contract binding to access the raw methods on
}

// DoubleDoNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoubleDoNFTTransactorRaw struct {
	Contract *DoubleDoNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoubleDoNFT creates a new instance of DoubleDoNFT, bound to a specific deployed contract.
func NewDoubleDoNFT(address common.Address, backend bind.ContractBackend) (*DoubleDoNFT, error) {
	contract, err := bindDoubleDoNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFT{DoubleDoNFTCaller: DoubleDoNFTCaller{contract: contract}, DoubleDoNFTTransactor: DoubleDoNFTTransactor{contract: contract}, DoubleDoNFTFilterer: DoubleDoNFTFilterer{contract: contract}}, nil
}

// NewDoubleDoNFTCaller creates a new read-only instance of DoubleDoNFT, bound to a specific deployed contract.
func NewDoubleDoNFTCaller(address common.Address, caller bind.ContractCaller) (*DoubleDoNFTCaller, error) {
	contract, err := bindDoubleDoNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTCaller{contract: contract}, nil
}

// NewDoubleDoNFTTransactor creates a new write-only instance of DoubleDoNFT, bound to a specific deployed contract.
func NewDoubleDoNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*DoubleDoNFTTransactor, error) {
	contract, err := bindDoubleDoNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTTransactor{contract: contract}, nil
}

// NewDoubleDoNFTFilterer creates a new log filterer instance of DoubleDoNFT, bound to a specific deployed contract.
func NewDoubleDoNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*DoubleDoNFTFilterer, error) {
	contract, err := bindDoubleDoNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTFilterer{contract: contract}, nil
}

// bindDoubleDoNFT binds a generic wrapper to an already deployed contract.
func bindDoubleDoNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DoubleDoNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleDoNFT *DoubleDoNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleDoNFT.Contract.DoubleDoNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleDoNFT *DoubleDoNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.DoubleDoNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleDoNFT *DoubleDoNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.DoubleDoNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleDoNFT *DoubleDoNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleDoNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleDoNFT *DoubleDoNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleDoNFT *DoubleDoNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DoubleDoNFT.Contract.DEFAULTADMINROLE(&_DoubleDoNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DoubleDoNFT.Contract.DEFAULTADMINROLE(&_DoubleDoNFT.CallOpts)
}

// MAXMINTWHITELIST is a free data retrieval call binding the contract method 0x58941a4d.
//
// Solidity: function MAX_MINT_WHITELIST() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) MAXMINTWHITELIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "MAX_MINT_WHITELIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMINTWHITELIST is a free data retrieval call binding the contract method 0x58941a4d.
//
// Solidity: function MAX_MINT_WHITELIST() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) MAXMINTWHITELIST() (*big.Int, error) {
	return _DoubleDoNFT.Contract.MAXMINTWHITELIST(&_DoubleDoNFT.CallOpts)
}

// MAXMINTWHITELIST is a free data retrieval call binding the contract method 0x58941a4d.
//
// Solidity: function MAX_MINT_WHITELIST() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) MAXMINTWHITELIST() (*big.Int, error) {
	return _DoubleDoNFT.Contract.MAXMINTWHITELIST(&_DoubleDoNFT.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) MAXSUPPLY() (*big.Int, error) {
	return _DoubleDoNFT.Contract.MAXSUPPLY(&_DoubleDoNFT.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _DoubleDoNFT.Contract.MAXSUPPLY(&_DoubleDoNFT.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTSession) OPERATORROLE() ([32]byte, error) {
	return _DoubleDoNFT.Contract.OPERATORROLE(&_DoubleDoNFT.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) OPERATORROLE() ([32]byte, error) {
	return _DoubleDoNFT.Contract.OPERATORROLE(&_DoubleDoNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DoubleDoNFT.Contract.BalanceOf(&_DoubleDoNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DoubleDoNFT.Contract.BalanceOf(&_DoubleDoNFT.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) BaseURI() (string, error) {
	return _DoubleDoNFT.Contract.BaseURI(&_DoubleDoNFT.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) BaseURI() (string, error) {
	return _DoubleDoNFT.Contract.BaseURI(&_DoubleDoNFT.CallOpts)
}

// Blacklist is a free data retrieval call binding the contract method 0x0f2a0919.
//
// Solidity: function blacklist(uint256 ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) Blacklist(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "blacklist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0x0f2a0919.
//
// Solidity: function blacklist(uint256 ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) Blacklist(arg0 *big.Int) (*big.Int, error) {
	return _DoubleDoNFT.Contract.Blacklist(&_DoubleDoNFT.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0x0f2a0919.
//
// Solidity: function blacklist(uint256 ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) Blacklist(arg0 *big.Int) (*big.Int, error) {
	return _DoubleDoNFT.Contract.Blacklist(&_DoubleDoNFT.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DoubleDoNFT.Contract.GetApproved(&_DoubleDoNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DoubleDoNFT.Contract.GetApproved(&_DoubleDoNFT.CallOpts, tokenId)
}

// GetExpiredAmount is a free data retrieval call binding the contract method 0x662d62f1.
//
// Solidity: function getExpiredAmount() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) GetExpiredAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "getExpiredAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpiredAmount is a free data retrieval call binding the contract method 0x662d62f1.
//
// Solidity: function getExpiredAmount() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) GetExpiredAmount() (*big.Int, error) {
	return _DoubleDoNFT.Contract.GetExpiredAmount(&_DoubleDoNFT.CallOpts)
}

// GetExpiredAmount is a free data retrieval call binding the contract method 0x662d62f1.
//
// Solidity: function getExpiredAmount() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) GetExpiredAmount() (*big.Int, error) {
	return _DoubleDoNFT.Contract.GetExpiredAmount(&_DoubleDoNFT.CallOpts)
}

// GetPassLevel is a free data retrieval call binding the contract method 0x4fe19878.
//
// Solidity: function getPassLevel(uint256 _tokenId) view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) GetPassLevel(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "getPassLevel", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPassLevel is a free data retrieval call binding the contract method 0x4fe19878.
//
// Solidity: function getPassLevel(uint256 _tokenId) view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) GetPassLevel(_tokenId *big.Int) (string, error) {
	return _DoubleDoNFT.Contract.GetPassLevel(&_DoubleDoNFT.CallOpts, _tokenId)
}

// GetPassLevel is a free data retrieval call binding the contract method 0x4fe19878.
//
// Solidity: function getPassLevel(uint256 _tokenId) view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) GetPassLevel(_tokenId *big.Int) (string, error) {
	return _DoubleDoNFT.Contract.GetPassLevel(&_DoubleDoNFT.CallOpts, _tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DoubleDoNFT.Contract.GetRoleAdmin(&_DoubleDoNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DoubleDoNFT.Contract.GetRoleAdmin(&_DoubleDoNFT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DoubleDoNFT.Contract.HasRole(&_DoubleDoNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DoubleDoNFT.Contract.HasRole(&_DoubleDoNFT.CallOpts, role, account)
}

// HiddenMetadataURI is a free data retrieval call binding the contract method 0xba9e12f7.
//
// Solidity: function hiddenMetadataURI() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) HiddenMetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "hiddenMetadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HiddenMetadataURI is a free data retrieval call binding the contract method 0xba9e12f7.
//
// Solidity: function hiddenMetadataURI() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) HiddenMetadataURI() (string, error) {
	return _DoubleDoNFT.Contract.HiddenMetadataURI(&_DoubleDoNFT.CallOpts)
}

// HiddenMetadataURI is a free data retrieval call binding the contract method 0xba9e12f7.
//
// Solidity: function hiddenMetadataURI() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) HiddenMetadataURI() (string, error) {
	return _DoubleDoNFT.Contract.HiddenMetadataURI(&_DoubleDoNFT.CallOpts)
}

// ImageURIExtension is a free data retrieval call binding the contract method 0x0e481416.
//
// Solidity: function imageURIExtension() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) ImageURIExtension(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "imageURIExtension")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImageURIExtension is a free data retrieval call binding the contract method 0x0e481416.
//
// Solidity: function imageURIExtension() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) ImageURIExtension() (string, error) {
	return _DoubleDoNFT.Contract.ImageURIExtension(&_DoubleDoNFT.CallOpts)
}

// ImageURIExtension is a free data retrieval call binding the contract method 0x0e481416.
//
// Solidity: function imageURIExtension() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) ImageURIExtension() (string, error) {
	return _DoubleDoNFT.Contract.ImageURIExtension(&_DoubleDoNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DoubleDoNFT.Contract.IsApprovedForAll(&_DoubleDoNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DoubleDoNFT.Contract.IsApprovedForAll(&_DoubleDoNFT.CallOpts, owner, operator)
}

// IsBlacklist is a free data retrieval call binding the contract method 0xf026788b.
//
// Solidity: function isBlacklist(uint256 tokenId) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) IsBlacklist(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "isBlacklist", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklist is a free data retrieval call binding the contract method 0xf026788b.
//
// Solidity: function isBlacklist(uint256 tokenId) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) IsBlacklist(tokenId *big.Int) (bool, error) {
	return _DoubleDoNFT.Contract.IsBlacklist(&_DoubleDoNFT.CallOpts, tokenId)
}

// IsBlacklist is a free data retrieval call binding the contract method 0xf026788b.
//
// Solidity: function isBlacklist(uint256 tokenId) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) IsBlacklist(tokenId *big.Int) (bool, error) {
	return _DoubleDoNFT.Contract.IsBlacklist(&_DoubleDoNFT.CallOpts, tokenId)
}

// IsUniqueImage is a free data retrieval call binding the contract method 0xb22e3374.
//
// Solidity: function isUniqueImage() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) IsUniqueImage(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "isUniqueImage")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUniqueImage is a free data retrieval call binding the contract method 0xb22e3374.
//
// Solidity: function isUniqueImage() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) IsUniqueImage() (bool, error) {
	return _DoubleDoNFT.Contract.IsUniqueImage(&_DoubleDoNFT.CallOpts)
}

// IsUniqueImage is a free data retrieval call binding the contract method 0xb22e3374.
//
// Solidity: function isUniqueImage() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) IsUniqueImage() (bool, error) {
	return _DoubleDoNFT.Contract.IsUniqueImage(&_DoubleDoNFT.CallOpts)
}

// IsWhitelistMintActive is a free data retrieval call binding the contract method 0xfabd1d2d.
//
// Solidity: function isWhitelistMintActive() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) IsWhitelistMintActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "isWhitelistMintActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistMintActive is a free data retrieval call binding the contract method 0xfabd1d2d.
//
// Solidity: function isWhitelistMintActive() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) IsWhitelistMintActive() (bool, error) {
	return _DoubleDoNFT.Contract.IsWhitelistMintActive(&_DoubleDoNFT.CallOpts)
}

// IsWhitelistMintActive is a free data retrieval call binding the contract method 0xfabd1d2d.
//
// Solidity: function isWhitelistMintActive() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) IsWhitelistMintActive() (bool, error) {
	return _DoubleDoNFT.Contract.IsWhitelistMintActive(&_DoubleDoNFT.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTSession) MerkleRoot() ([32]byte, error) {
	return _DoubleDoNFT.Contract.MerkleRoot(&_DoubleDoNFT.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) MerkleRoot() ([32]byte, error) {
	return _DoubleDoNFT.Contract.MerkleRoot(&_DoubleDoNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) Name() (string, error) {
	return _DoubleDoNFT.Contract.Name(&_DoubleDoNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) Name() (string, error) {
	return _DoubleDoNFT.Contract.Name(&_DoubleDoNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoubleDoNFT *DoubleDoNFTSession) Owner() (common.Address, error) {
	return _DoubleDoNFT.Contract.Owner(&_DoubleDoNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) Owner() (common.Address, error) {
	return _DoubleDoNFT.Contract.Owner(&_DoubleDoNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DoubleDoNFT.Contract.OwnerOf(&_DoubleDoNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DoubleDoNFT.Contract.OwnerOf(&_DoubleDoNFT.CallOpts, tokenId)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) Revealed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "revealed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) Revealed() (bool, error) {
	return _DoubleDoNFT.Contract.Revealed(&_DoubleDoNFT.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) Revealed() (bool, error) {
	return _DoubleDoNFT.Contract.Revealed(&_DoubleDoNFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DoubleDoNFT.Contract.SupportsInterface(&_DoubleDoNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DoubleDoNFT.Contract.SupportsInterface(&_DoubleDoNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) Symbol() (string, error) {
	return _DoubleDoNFT.Contract.Symbol(&_DoubleDoNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) Symbol() (string, error) {
	return _DoubleDoNFT.Contract.Symbol(&_DoubleDoNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DoubleDoNFT *DoubleDoNFTSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DoubleDoNFT.Contract.TokenURI(&_DoubleDoNFT.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DoubleDoNFT.Contract.TokenURI(&_DoubleDoNFT.CallOpts, _tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0x003d4790.
//
// Solidity: function totalMinted(address ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) TotalMinted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "totalMinted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0x003d4790.
//
// Solidity: function totalMinted(address ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) TotalMinted(arg0 common.Address) (*big.Int, error) {
	return _DoubleDoNFT.Contract.TotalMinted(&_DoubleDoNFT.CallOpts, arg0)
}

// TotalMinted is a free data retrieval call binding the contract method 0x003d4790.
//
// Solidity: function totalMinted(address ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) TotalMinted(arg0 common.Address) (*big.Int, error) {
	return _DoubleDoNFT.Contract.TotalMinted(&_DoubleDoNFT.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) TotalSupply() (*big.Int, error) {
	return _DoubleDoNFT.Contract.TotalSupply(&_DoubleDoNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _DoubleDoNFT.Contract.TotalSupply(&_DoubleDoNFT.CallOpts)
}

// TrialPassIds is a free data retrieval call binding the contract method 0xdaa46640.
//
// Solidity: function trialPassIds(uint256 ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) TrialPassIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "trialPassIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrialPassIds is a free data retrieval call binding the contract method 0xdaa46640.
//
// Solidity: function trialPassIds(uint256 ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) TrialPassIds(arg0 *big.Int) (*big.Int, error) {
	return _DoubleDoNFT.Contract.TrialPassIds(&_DoubleDoNFT.CallOpts, arg0)
}

// TrialPassIds is a free data retrieval call binding the contract method 0xdaa46640.
//
// Solidity: function trialPassIds(uint256 ) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) TrialPassIds(arg0 *big.Int) (*big.Int, error) {
	return _DoubleDoNFT.Contract.TrialPassIds(&_DoubleDoNFT.CallOpts, arg0)
}

// UpgradeAmount is a free data retrieval call binding the contract method 0xce264e1f.
//
// Solidity: function upgradeAmount() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) UpgradeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "upgradeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpgradeAmount is a free data retrieval call binding the contract method 0xce264e1f.
//
// Solidity: function upgradeAmount() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) UpgradeAmount() (*big.Int, error) {
	return _DoubleDoNFT.Contract.UpgradeAmount(&_DoubleDoNFT.CallOpts)
}

// UpgradeAmount is a free data retrieval call binding the contract method 0xce264e1f.
//
// Solidity: function upgradeAmount() view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) UpgradeAmount() (*big.Int, error) {
	return _DoubleDoNFT.Contract.UpgradeAmount(&_DoubleDoNFT.CallOpts)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCaller) UserExpires(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "userExpires", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _DoubleDoNFT.Contract.UserExpires(&_DoubleDoNFT.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _DoubleDoNFT.Contract.UserExpires(&_DoubleDoNFT.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCaller) UserOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "userOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _DoubleDoNFT.Contract.UserOf(&_DoubleDoNFT.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _DoubleDoNFT.Contract.UserOf(&_DoubleDoNFT.CallOpts, tokenId)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DoubleDoNFT.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_DoubleDoNFT *DoubleDoNFTSession) VaultAddress() (common.Address, error) {
	return _DoubleDoNFT.Contract.VaultAddress(&_DoubleDoNFT.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_DoubleDoNFT *DoubleDoNFTCallerSession) VaultAddress() (common.Address, error) {
	return _DoubleDoNFT.Contract.VaultAddress(&_DoubleDoNFT.CallOpts)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x6113015f.
//
// Solidity: function addBlacklist(uint256 tokenId) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) AddBlacklist(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "addBlacklist", tokenId)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x6113015f.
//
// Solidity: function addBlacklist(uint256 tokenId) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) AddBlacklist(tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.AddBlacklist(&_DoubleDoNFT.TransactOpts, tokenId)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x6113015f.
//
// Solidity: function addBlacklist(uint256 tokenId) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) AddBlacklist(tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.AddBlacklist(&_DoubleDoNFT.TransactOpts, tokenId)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] team, uint256[] teamMint) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) Airdrop(opts *bind.TransactOpts, team []common.Address, teamMint []*big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "airdrop", team, teamMint)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] team, uint256[] teamMint) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) Airdrop(team []common.Address, teamMint []*big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.Airdrop(&_DoubleDoNFT.TransactOpts, team, teamMint)
}

// Airdrop is a paid mutator transaction binding the contract method 0x67243482.
//
// Solidity: function airdrop(address[] team, uint256[] teamMint) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) Airdrop(team []common.Address, teamMint []*big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.Airdrop(&_DoubleDoNFT.TransactOpts, team, teamMint)
}

// AirdropTrialPass is a paid mutator transaction binding the contract method 0xa3b53fbd.
//
// Solidity: function airdropTrialPass(address[] airdropUsers, uint64[] expireDays) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) AirdropTrialPass(opts *bind.TransactOpts, airdropUsers []common.Address, expireDays []uint64) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "airdropTrialPass", airdropUsers, expireDays)
}

// AirdropTrialPass is a paid mutator transaction binding the contract method 0xa3b53fbd.
//
// Solidity: function airdropTrialPass(address[] airdropUsers, uint64[] expireDays) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) AirdropTrialPass(airdropUsers []common.Address, expireDays []uint64) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.AirdropTrialPass(&_DoubleDoNFT.TransactOpts, airdropUsers, expireDays)
}

// AirdropTrialPass is a paid mutator transaction binding the contract method 0xa3b53fbd.
//
// Solidity: function airdropTrialPass(address[] airdropUsers, uint64[] expireDays) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) AirdropTrialPass(airdropUsers []common.Address, expireDays []uint64) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.AirdropTrialPass(&_DoubleDoNFT.TransactOpts, airdropUsers, expireDays)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.Approve(&_DoubleDoNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.Approve(&_DoubleDoNFT.TransactOpts, to, tokenId)
}

// CloseMint is a paid mutator transaction binding the contract method 0x64f101f0.
//
// Solidity: function closeMint() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) CloseMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "closeMint")
}

// CloseMint is a paid mutator transaction binding the contract method 0x64f101f0.
//
// Solidity: function closeMint() returns()
func (_DoubleDoNFT *DoubleDoNFTSession) CloseMint() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.CloseMint(&_DoubleDoNFT.TransactOpts)
}

// CloseMint is a paid mutator transaction binding the contract method 0x64f101f0.
//
// Solidity: function closeMint() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) CloseMint() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.CloseMint(&_DoubleDoNFT.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.GrantRole(&_DoubleDoNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.GrantRole(&_DoubleDoNFT.TransactOpts, role, account)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DoubleDoNFT *DoubleDoNFTTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DoubleDoNFT *DoubleDoNFTSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.OnERC721Received(&_DoubleDoNFT.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.OnERC721Received(&_DoubleDoNFT.TransactOpts, arg0, arg1, arg2, arg3)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0x7f994945.
//
// Solidity: function removeBlacklist(uint256 tokenId) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) RemoveBlacklist(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "removeBlacklist", tokenId)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0x7f994945.
//
// Solidity: function removeBlacklist(uint256 tokenId) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) RemoveBlacklist(tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RemoveBlacklist(&_DoubleDoNFT.TransactOpts, tokenId)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0x7f994945.
//
// Solidity: function removeBlacklist(uint256 tokenId) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) RemoveBlacklist(tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RemoveBlacklist(&_DoubleDoNFT.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoubleDoNFT *DoubleDoNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RenounceOwnership(&_DoubleDoNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RenounceOwnership(&_DoubleDoNFT.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RenounceRole(&_DoubleDoNFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RenounceRole(&_DoubleDoNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RevokeRole(&_DoubleDoNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.RevokeRole(&_DoubleDoNFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SafeTransferFrom(&_DoubleDoNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SafeTransferFrom(&_DoubleDoNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SafeTransferFrom0(&_DoubleDoNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SafeTransferFrom0(&_DoubleDoNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetApprovalForAll(&_DoubleDoNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetApprovalForAll(&_DoubleDoNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newURI) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetBaseURI(opts *bind.TransactOpts, _newURI string) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setBaseURI", _newURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newURI) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetBaseURI(_newURI string) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetBaseURI(&_DoubleDoNFT.TransactOpts, _newURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newURI) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetBaseURI(_newURI string) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetBaseURI(&_DoubleDoNFT.TransactOpts, _newURI)
}

// SetImageExtension is a paid mutator transaction binding the contract method 0x39e3aa7e.
//
// Solidity: function setImageExtension(string _newImageURIExtension) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetImageExtension(opts *bind.TransactOpts, _newImageURIExtension string) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setImageExtension", _newImageURIExtension)
}

// SetImageExtension is a paid mutator transaction binding the contract method 0x39e3aa7e.
//
// Solidity: function setImageExtension(string _newImageURIExtension) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetImageExtension(_newImageURIExtension string) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetImageExtension(&_DoubleDoNFT.TransactOpts, _newImageURIExtension)
}

// SetImageExtension is a paid mutator transaction binding the contract method 0x39e3aa7e.
//
// Solidity: function setImageExtension(string _newImageURIExtension) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetImageExtension(_newImageURIExtension string) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetImageExtension(&_DoubleDoNFT.TransactOpts, _newImageURIExtension)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setMerkleRoot", _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetMerkleRoot(&_DoubleDoNFT.TransactOpts, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _merkleRoot) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetMerkleRoot(&_DoubleDoNFT.TransactOpts, _merkleRoot)
}

// SetUpgradeAmount is a paid mutator transaction binding the contract method 0xedd66c63.
//
// Solidity: function setUpgradeAmount(uint256 _upgradeAmount) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetUpgradeAmount(opts *bind.TransactOpts, _upgradeAmount *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setUpgradeAmount", _upgradeAmount)
}

// SetUpgradeAmount is a paid mutator transaction binding the contract method 0xedd66c63.
//
// Solidity: function setUpgradeAmount(uint256 _upgradeAmount) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetUpgradeAmount(_upgradeAmount *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetUpgradeAmount(&_DoubleDoNFT.TransactOpts, _upgradeAmount)
}

// SetUpgradeAmount is a paid mutator transaction binding the contract method 0xedd66c63.
//
// Solidity: function setUpgradeAmount(uint256 _upgradeAmount) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetUpgradeAmount(_upgradeAmount *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetUpgradeAmount(&_DoubleDoNFT.TransactOpts, _upgradeAmount)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetUser(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setUser", tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetUser(&_DoubleDoNFT.TransactOpts, tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetUser(&_DoubleDoNFT.TransactOpts, tokenId, user, expires)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetVaultAddress(&_DoubleDoNFT.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.SetVaultAddress(&_DoubleDoNFT.TransactOpts, _vaultAddress)
}

// ToggleIsUniqueImage is a paid mutator transaction binding the contract method 0x48f84209.
//
// Solidity: function toggleIsUniqueImage() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) ToggleIsUniqueImage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "toggleIsUniqueImage")
}

// ToggleIsUniqueImage is a paid mutator transaction binding the contract method 0x48f84209.
//
// Solidity: function toggleIsUniqueImage() returns()
func (_DoubleDoNFT *DoubleDoNFTSession) ToggleIsUniqueImage() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.ToggleIsUniqueImage(&_DoubleDoNFT.TransactOpts)
}

// ToggleIsUniqueImage is a paid mutator transaction binding the contract method 0x48f84209.
//
// Solidity: function toggleIsUniqueImage() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) ToggleIsUniqueImage() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.ToggleIsUniqueImage(&_DoubleDoNFT.TransactOpts)
}

// ToggleReveal is a paid mutator transaction binding the contract method 0x5b8ad429.
//
// Solidity: function toggleReveal() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) ToggleReveal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "toggleReveal")
}

// ToggleReveal is a paid mutator transaction binding the contract method 0x5b8ad429.
//
// Solidity: function toggleReveal() returns()
func (_DoubleDoNFT *DoubleDoNFTSession) ToggleReveal() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.ToggleReveal(&_DoubleDoNFT.TransactOpts)
}

// ToggleReveal is a paid mutator transaction binding the contract method 0x5b8ad429.
//
// Solidity: function toggleReveal() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) ToggleReveal() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.ToggleReveal(&_DoubleDoNFT.TransactOpts)
}

// ToggleWhitelistMint is a paid mutator transaction binding the contract method 0x6f63b60a.
//
// Solidity: function toggleWhitelistMint() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) ToggleWhitelistMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "toggleWhitelistMint")
}

// ToggleWhitelistMint is a paid mutator transaction binding the contract method 0x6f63b60a.
//
// Solidity: function toggleWhitelistMint() returns()
func (_DoubleDoNFT *DoubleDoNFTSession) ToggleWhitelistMint() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.ToggleWhitelistMint(&_DoubleDoNFT.TransactOpts)
}

// ToggleWhitelistMint is a paid mutator transaction binding the contract method 0x6f63b60a.
//
// Solidity: function toggleWhitelistMint() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) ToggleWhitelistMint() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.ToggleWhitelistMint(&_DoubleDoNFT.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.TransferFrom(&_DoubleDoNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.TransferFrom(&_DoubleDoNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.TransferOwnership(&_DoubleDoNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.TransferOwnership(&_DoubleDoNFT.TransactOpts, newOwner)
}

// UpgradeNFT is a paid mutator transaction binding the contract method 0x64e5d4fe.
//
// Solidity: function upgradeNFT(uint256[] tokenIds) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) UpgradeNFT(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "upgradeNFT", tokenIds)
}

// UpgradeNFT is a paid mutator transaction binding the contract method 0x64e5d4fe.
//
// Solidity: function upgradeNFT(uint256[] tokenIds) returns()
func (_DoubleDoNFT *DoubleDoNFTSession) UpgradeNFT(tokenIds []*big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.UpgradeNFT(&_DoubleDoNFT.TransactOpts, tokenIds)
}

// UpgradeNFT is a paid mutator transaction binding the contract method 0x64e5d4fe.
//
// Solidity: function upgradeNFT(uint256[] tokenIds) returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) UpgradeNFT(tokenIds []*big.Int) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.UpgradeNFT(&_DoubleDoNFT.TransactOpts, tokenIds)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xd2cab056.
//
// Solidity: function whitelistMint(uint256 _quantity, bytes32[] _proof) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) WhitelistMint(opts *bind.TransactOpts, _quantity *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "whitelistMint", _quantity, _proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xd2cab056.
//
// Solidity: function whitelistMint(uint256 _quantity, bytes32[] _proof) payable returns()
func (_DoubleDoNFT *DoubleDoNFTSession) WhitelistMint(_quantity *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.WhitelistMint(&_DoubleDoNFT.TransactOpts, _quantity, _proof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0xd2cab056.
//
// Solidity: function whitelistMint(uint256 _quantity, bytes32[] _proof) payable returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) WhitelistMint(_quantity *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.WhitelistMint(&_DoubleDoNFT.TransactOpts, _quantity, _proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleDoNFT.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DoubleDoNFT *DoubleDoNFTSession) Withdraw() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.Withdraw(&_DoubleDoNFT.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_DoubleDoNFT *DoubleDoNFTTransactorSession) Withdraw() (*types.Transaction, error) {
	return _DoubleDoNFT.Contract.Withdraw(&_DoubleDoNFT.TransactOpts)
}

// DoubleDoNFTAddBlacklistIterator is returned from FilterAddBlacklist and is used to iterate over the raw logs and unpacked data for AddBlacklist events raised by the DoubleDoNFT contract.
type DoubleDoNFTAddBlacklistIterator struct {
	Event *DoubleDoNFTAddBlacklist // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTAddBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTAddBlacklist)
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
		it.Event = new(DoubleDoNFTAddBlacklist)
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
func (it *DoubleDoNFTAddBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTAddBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTAddBlacklist represents a AddBlacklist event raised by the DoubleDoNFT contract.
type DoubleDoNFTAddBlacklist struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddBlacklist is a free log retrieval operation binding the contract event 0x3b2e1a1bab46e5ef778a0d07131df8c28114f93a2efe62547a1425abf5aba0a3.
//
// Solidity: event AddBlacklist(uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterAddBlacklist(opts *bind.FilterOpts) (*DoubleDoNFTAddBlacklistIterator, error) {

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "AddBlacklist")
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTAddBlacklistIterator{contract: _DoubleDoNFT.contract, event: "AddBlacklist", logs: logs, sub: sub}, nil
}

// WatchAddBlacklist is a free log subscription operation binding the contract event 0x3b2e1a1bab46e5ef778a0d07131df8c28114f93a2efe62547a1425abf5aba0a3.
//
// Solidity: event AddBlacklist(uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchAddBlacklist(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTAddBlacklist) (event.Subscription, error) {

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "AddBlacklist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTAddBlacklist)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "AddBlacklist", log); err != nil {
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

// ParseAddBlacklist is a log parse operation binding the contract event 0x3b2e1a1bab46e5ef778a0d07131df8c28114f93a2efe62547a1425abf5aba0a3.
//
// Solidity: event AddBlacklist(uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseAddBlacklist(log types.Log) (*DoubleDoNFTAddBlacklist, error) {
	event := new(DoubleDoNFTAddBlacklist)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "AddBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DoubleDoNFT contract.
type DoubleDoNFTApprovalIterator struct {
	Event *DoubleDoNFTApproval // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTApproval)
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
		it.Event = new(DoubleDoNFTApproval)
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
func (it *DoubleDoNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTApproval represents a Approval event raised by the DoubleDoNFT contract.
type DoubleDoNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DoubleDoNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTApprovalIterator{contract: _DoubleDoNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTApproval)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseApproval(log types.Log) (*DoubleDoNFTApproval, error) {
	event := new(DoubleDoNFTApproval)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DoubleDoNFT contract.
type DoubleDoNFTApprovalForAllIterator struct {
	Event *DoubleDoNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTApprovalForAll)
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
		it.Event = new(DoubleDoNFTApprovalForAll)
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
func (it *DoubleDoNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTApprovalForAll represents a ApprovalForAll event raised by the DoubleDoNFT contract.
type DoubleDoNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DoubleDoNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTApprovalForAllIterator{contract: _DoubleDoNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTApprovalForAll)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseApprovalForAll(log types.Log) (*DoubleDoNFTApprovalForAll, error) {
	event := new(DoubleDoNFTApprovalForAll)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the DoubleDoNFT contract.
type DoubleDoNFTConsecutiveTransferIterator struct {
	Event *DoubleDoNFTConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTConsecutiveTransfer)
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
		it.Event = new(DoubleDoNFTConsecutiveTransfer)
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
func (it *DoubleDoNFTConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTConsecutiveTransfer represents a ConsecutiveTransfer event raised by the DoubleDoNFT contract.
type DoubleDoNFTConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*DoubleDoNFTConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTConsecutiveTransferIterator{contract: _DoubleDoNFT.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTConsecutiveTransfer)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseConsecutiveTransfer(log types.Log) (*DoubleDoNFTConsecutiveTransfer, error) {
	event := new(DoubleDoNFTConsecutiveTransfer)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DoubleDoNFT contract.
type DoubleDoNFTOwnershipTransferredIterator struct {
	Event *DoubleDoNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTOwnershipTransferred)
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
		it.Event = new(DoubleDoNFTOwnershipTransferred)
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
func (it *DoubleDoNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTOwnershipTransferred represents a OwnershipTransferred event raised by the DoubleDoNFT contract.
type DoubleDoNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DoubleDoNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTOwnershipTransferredIterator{contract: _DoubleDoNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTOwnershipTransferred)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseOwnershipTransferred(log types.Log) (*DoubleDoNFTOwnershipTransferred, error) {
	event := new(DoubleDoNFTOwnershipTransferred)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTRemoveBlacklistIterator is returned from FilterRemoveBlacklist and is used to iterate over the raw logs and unpacked data for RemoveBlacklist events raised by the DoubleDoNFT contract.
type DoubleDoNFTRemoveBlacklistIterator struct {
	Event *DoubleDoNFTRemoveBlacklist // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTRemoveBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTRemoveBlacklist)
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
		it.Event = new(DoubleDoNFTRemoveBlacklist)
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
func (it *DoubleDoNFTRemoveBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTRemoveBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTRemoveBlacklist represents a RemoveBlacklist event raised by the DoubleDoNFT contract.
type DoubleDoNFTRemoveBlacklist struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlacklist is a free log retrieval operation binding the contract event 0x182aec2693af7aa18cf80584fe9ea15104a308a2bb5f8bd5242e5dce2e9aa849.
//
// Solidity: event RemoveBlacklist(uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterRemoveBlacklist(opts *bind.FilterOpts) (*DoubleDoNFTRemoveBlacklistIterator, error) {

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "RemoveBlacklist")
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTRemoveBlacklistIterator{contract: _DoubleDoNFT.contract, event: "RemoveBlacklist", logs: logs, sub: sub}, nil
}

// WatchRemoveBlacklist is a free log subscription operation binding the contract event 0x182aec2693af7aa18cf80584fe9ea15104a308a2bb5f8bd5242e5dce2e9aa849.
//
// Solidity: event RemoveBlacklist(uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchRemoveBlacklist(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTRemoveBlacklist) (event.Subscription, error) {

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "RemoveBlacklist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTRemoveBlacklist)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "RemoveBlacklist", log); err != nil {
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

// ParseRemoveBlacklist is a log parse operation binding the contract event 0x182aec2693af7aa18cf80584fe9ea15104a308a2bb5f8bd5242e5dce2e9aa849.
//
// Solidity: event RemoveBlacklist(uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseRemoveBlacklist(log types.Log) (*DoubleDoNFTRemoveBlacklist, error) {
	event := new(DoubleDoNFTRemoveBlacklist)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "RemoveBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DoubleDoNFT contract.
type DoubleDoNFTRoleAdminChangedIterator struct {
	Event *DoubleDoNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTRoleAdminChanged)
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
		it.Event = new(DoubleDoNFTRoleAdminChanged)
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
func (it *DoubleDoNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTRoleAdminChanged represents a RoleAdminChanged event raised by the DoubleDoNFT contract.
type DoubleDoNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DoubleDoNFTRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTRoleAdminChangedIterator{contract: _DoubleDoNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTRoleAdminChanged)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseRoleAdminChanged(log types.Log) (*DoubleDoNFTRoleAdminChanged, error) {
	event := new(DoubleDoNFTRoleAdminChanged)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DoubleDoNFT contract.
type DoubleDoNFTRoleGrantedIterator struct {
	Event *DoubleDoNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTRoleGranted)
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
		it.Event = new(DoubleDoNFTRoleGranted)
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
func (it *DoubleDoNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTRoleGranted represents a RoleGranted event raised by the DoubleDoNFT contract.
type DoubleDoNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DoubleDoNFTRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTRoleGrantedIterator{contract: _DoubleDoNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTRoleGranted)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseRoleGranted(log types.Log) (*DoubleDoNFTRoleGranted, error) {
	event := new(DoubleDoNFTRoleGranted)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DoubleDoNFT contract.
type DoubleDoNFTRoleRevokedIterator struct {
	Event *DoubleDoNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTRoleRevoked)
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
		it.Event = new(DoubleDoNFTRoleRevoked)
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
func (it *DoubleDoNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTRoleRevoked represents a RoleRevoked event raised by the DoubleDoNFT contract.
type DoubleDoNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DoubleDoNFTRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTRoleRevokedIterator{contract: _DoubleDoNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTRoleRevoked)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseRoleRevoked(log types.Log) (*DoubleDoNFTRoleRevoked, error) {
	event := new(DoubleDoNFTRoleRevoked)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DoubleDoNFT contract.
type DoubleDoNFTTransferIterator struct {
	Event *DoubleDoNFTTransfer // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTTransfer)
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
		it.Event = new(DoubleDoNFTTransfer)
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
func (it *DoubleDoNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTTransfer represents a Transfer event raised by the DoubleDoNFT contract.
type DoubleDoNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DoubleDoNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTTransferIterator{contract: _DoubleDoNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTTransfer)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseTransfer(log types.Log) (*DoubleDoNFTTransfer, error) {
	event := new(DoubleDoNFTTransfer)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTUpdateUserIterator is returned from FilterUpdateUser and is used to iterate over the raw logs and unpacked data for UpdateUser events raised by the DoubleDoNFT contract.
type DoubleDoNFTUpdateUserIterator struct {
	Event *DoubleDoNFTUpdateUser // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTUpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTUpdateUser)
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
		it.Event = new(DoubleDoNFTUpdateUser)
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
func (it *DoubleDoNFTUpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTUpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTUpdateUser represents a UpdateUser event raised by the DoubleDoNFT contract.
type DoubleDoNFTUpdateUser struct {
	TokenId *big.Int
	User    common.Address
	Expires uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateUser is a free log retrieval operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterUpdateUser(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*DoubleDoNFTUpdateUserIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTUpdateUserIterator{contract: _DoubleDoNFT.contract, event: "UpdateUser", logs: logs, sub: sub}, nil
}

// WatchUpdateUser is a free log subscription operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchUpdateUser(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTUpdateUser, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTUpdateUser)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "UpdateUser", log); err != nil {
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

// ParseUpdateUser is a log parse operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseUpdateUser(log types.Log) (*DoubleDoNFTUpdateUser, error) {
	event := new(DoubleDoNFTUpdateUser)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "UpdateUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleDoNFTUpgradedNFTIterator is returned from FilterUpgradedNFT and is used to iterate over the raw logs and unpacked data for UpgradedNFT events raised by the DoubleDoNFT contract.
type DoubleDoNFTUpgradedNFTIterator struct {
	Event *DoubleDoNFTUpgradedNFT // Event containing the contract specifics and raw log

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
func (it *DoubleDoNFTUpgradedNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleDoNFTUpgradedNFT)
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
		it.Event = new(DoubleDoNFTUpgradedNFT)
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
func (it *DoubleDoNFTUpgradedNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleDoNFTUpgradedNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleDoNFTUpgradedNFT represents a UpgradedNFT event raised by the DoubleDoNFT contract.
type DoubleDoNFTUpgradedNFT struct {
	Owner    common.Address
	TokenIds []*big.Int
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpgradedNFT is a free log retrieval operation binding the contract event 0xa2eba574039caa31054b9444b0b704df00127d630d957317032fd1e3bbe55753.
//
// Solidity: event UpgradedNFT(address owner, uint256[] tokenIds, uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) FilterUpgradedNFT(opts *bind.FilterOpts) (*DoubleDoNFTUpgradedNFTIterator, error) {

	logs, sub, err := _DoubleDoNFT.contract.FilterLogs(opts, "UpgradedNFT")
	if err != nil {
		return nil, err
	}
	return &DoubleDoNFTUpgradedNFTIterator{contract: _DoubleDoNFT.contract, event: "UpgradedNFT", logs: logs, sub: sub}, nil
}

// WatchUpgradedNFT is a free log subscription operation binding the contract event 0xa2eba574039caa31054b9444b0b704df00127d630d957317032fd1e3bbe55753.
//
// Solidity: event UpgradedNFT(address owner, uint256[] tokenIds, uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) WatchUpgradedNFT(opts *bind.WatchOpts, sink chan<- *DoubleDoNFTUpgradedNFT) (event.Subscription, error) {

	logs, sub, err := _DoubleDoNFT.contract.WatchLogs(opts, "UpgradedNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleDoNFTUpgradedNFT)
				if err := _DoubleDoNFT.contract.UnpackLog(event, "UpgradedNFT", log); err != nil {
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

// ParseUpgradedNFT is a log parse operation binding the contract event 0xa2eba574039caa31054b9444b0b704df00127d630d957317032fd1e3bbe55753.
//
// Solidity: event UpgradedNFT(address owner, uint256[] tokenIds, uint256 tokenId)
func (_DoubleDoNFT *DoubleDoNFTFilterer) ParseUpgradedNFT(log types.Log) (*DoubleDoNFTUpgradedNFT, error) {
	event := new(DoubleDoNFTUpgradedNFT)
	if err := _DoubleDoNFT.contract.UnpackLog(event, "UpgradedNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
