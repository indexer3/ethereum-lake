// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn

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

// RegisteryAdapterV2VaultsAdapterInfo is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAdapterInfo struct {
	Id         common.Address
	TypeId     string
	CategoryId string
}

// RegisteryAdapterV2VaultsAdapterPosition is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAdapterPosition struct {
	BalanceUsdc *big.Int
}

// RegisteryAdapterV2VaultsAllowance is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAllowance struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
}

// RegisteryAdapterV2VaultsAssetDynamic is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAssetDynamic struct {
	Id                     common.Address
	TypeId                 string
	TokenId                common.Address
	UnderlyingTokenBalance RegisteryAdapterV2VaultsTokenAmount
	Metadata               RegisteryAdapterV2VaultsAssetMetadata
}

// RegisteryAdapterV2VaultsAssetMetadata is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAssetMetadata struct {
	PricePerShare      *big.Int
	MigrationAvailable bool
	LatestVaultAddress common.Address
	DepositLimit       *big.Int
	EmergencyShutdown  bool
}

// RegisteryAdapterV2VaultsAssetStatic is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAssetStatic struct {
	Id       common.Address
	TypeId   string
	TokenId  common.Address
	Name     string
	Version  string
	Symbol   string
	Decimals uint8
}

// RegisteryAdapterV2VaultsAssetUserMetadata is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsAssetUserMetadata struct {
	DepositBalance common.Address
}

// RegisteryAdapterV2VaultsPosition is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsPosition struct {
	AssetId                common.Address
	TokenId                common.Address
	TypeId                 string
	Balance                *big.Int
	UnderlyingTokenBalance RegisteryAdapterV2VaultsTokenAmount
	TokenAllowances        []RegisteryAdapterV2VaultsAllowance
	AssetAllowances        []RegisteryAdapterV2VaultsAllowance
}

// RegisteryAdapterV2VaultsTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type RegisteryAdapterV2VaultsTokenAmount struct {
	Amount     *big.Int
	AmountUsdc *big.Int
}

// RegisterAdapterV2VaultMetaData contains all meta data concerning the RegisterAdapterV2Vault contract.
var RegisterAdapterV2VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_helperAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressesGeneratorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"adapterInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"categoryId\",\"type\":\"string\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AdapterInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"adapterPositionOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balanceUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AdapterPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressesGeneratorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"assetAllowances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"assetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"assetDynamic\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.TokenAmount\",\"name\":\"underlyingTokenBalance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"migrationAvailable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"latestVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"emergencyShutdown\",\"type\":\"bool\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetDynamic\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"assetPositionsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.TokenAmount\",\"name\":\"underlyingTokenBalance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"tokenAllowances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"assetAllowances\",\"type\":\"tuple[]\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"assetStatic\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetStatic\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"assetUnderlyingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"assetUserMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"depositBalance\",\"type\":\"address\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetUserMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsDynamic\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.TokenAmount\",\"name\":\"underlyingTokenBalance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"migrationAvailable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"latestVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"emergencyShutdown\",\"type\":\"bool\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetDynamic[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assetsAddresses\",\"type\":\"address[]\"}],\"name\":\"assetsDynamic\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.TokenAmount\",\"name\":\"underlyingTokenBalance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"migrationAvailable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"latestVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"emergencyShutdown\",\"type\":\"bool\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetDynamic[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_assetsAddresses\",\"type\":\"address[]\"}],\"name\":\"assetsPositionsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.TokenAmount\",\"name\":\"underlyingTokenBalance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"tokenAllowances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"assetAllowances\",\"type\":\"tuple[]\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"assetsPositionsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdc\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.TokenAmount\",\"name\":\"underlyingTokenBalance\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"tokenAllowances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"assetAllowances\",\"type\":\"tuple[]\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Position[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsStatic\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetStatic[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assetsAddresses\",\"type\":\"address[]\"}],\"name\":\"assetsStatic\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetStatic[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsTokensAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"assetsUserMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"depositBalance\",\"type\":\"address\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.AssetUserMetadata[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionsAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"helperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newExtensionsAddresses\",\"type\":\"address[]\"}],\"name\":\"setExtensionsAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportedPositions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"}],\"name\":\"tokenAllowances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structRegisteryAdapterV2Vaults.Allowance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateSlot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegisterAdapterV2VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use RegisterAdapterV2VaultMetaData.ABI instead.
var RegisterAdapterV2VaultABI = RegisterAdapterV2VaultMetaData.ABI

// RegisterAdapterV2Vault is an auto generated Go binding around an Ethereum contract.
type RegisterAdapterV2Vault struct {
	RegisterAdapterV2VaultCaller     // Read-only binding to the contract
	RegisterAdapterV2VaultTransactor // Write-only binding to the contract
	RegisterAdapterV2VaultFilterer   // Log filterer for contract events
}

// RegisterAdapterV2VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegisterAdapterV2VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterAdapterV2VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegisterAdapterV2VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterAdapterV2VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegisterAdapterV2VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterAdapterV2VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegisterAdapterV2VaultSession struct {
	Contract     *RegisterAdapterV2Vault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegisterAdapterV2VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegisterAdapterV2VaultCallerSession struct {
	Contract *RegisterAdapterV2VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RegisterAdapterV2VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegisterAdapterV2VaultTransactorSession struct {
	Contract     *RegisterAdapterV2VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RegisterAdapterV2VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegisterAdapterV2VaultRaw struct {
	Contract *RegisterAdapterV2Vault // Generic contract binding to access the raw methods on
}

// RegisterAdapterV2VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegisterAdapterV2VaultCallerRaw struct {
	Contract *RegisterAdapterV2VaultCaller // Generic read-only contract binding to access the raw methods on
}

// RegisterAdapterV2VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegisterAdapterV2VaultTransactorRaw struct {
	Contract *RegisterAdapterV2VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegisterAdapterV2Vault creates a new instance of RegisterAdapterV2Vault, bound to a specific deployed contract.
func NewRegisterAdapterV2Vault(address common.Address, backend bind.ContractBackend) (*RegisterAdapterV2Vault, error) {
	contract, err := bindRegisterAdapterV2Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegisterAdapterV2Vault{RegisterAdapterV2VaultCaller: RegisterAdapterV2VaultCaller{contract: contract}, RegisterAdapterV2VaultTransactor: RegisterAdapterV2VaultTransactor{contract: contract}, RegisterAdapterV2VaultFilterer: RegisterAdapterV2VaultFilterer{contract: contract}}, nil
}

// NewRegisterAdapterV2VaultCaller creates a new read-only instance of RegisterAdapterV2Vault, bound to a specific deployed contract.
func NewRegisterAdapterV2VaultCaller(address common.Address, caller bind.ContractCaller) (*RegisterAdapterV2VaultCaller, error) {
	contract, err := bindRegisterAdapterV2Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterAdapterV2VaultCaller{contract: contract}, nil
}

// NewRegisterAdapterV2VaultTransactor creates a new write-only instance of RegisterAdapterV2Vault, bound to a specific deployed contract.
func NewRegisterAdapterV2VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*RegisterAdapterV2VaultTransactor, error) {
	contract, err := bindRegisterAdapterV2Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterAdapterV2VaultTransactor{contract: contract}, nil
}

// NewRegisterAdapterV2VaultFilterer creates a new log filterer instance of RegisterAdapterV2Vault, bound to a specific deployed contract.
func NewRegisterAdapterV2VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*RegisterAdapterV2VaultFilterer, error) {
	contract, err := bindRegisterAdapterV2Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegisterAdapterV2VaultFilterer{contract: contract}, nil
}

// bindRegisterAdapterV2Vault binds a generic wrapper to an already deployed contract.
func bindRegisterAdapterV2Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegisterAdapterV2VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegisterAdapterV2Vault.Contract.RegisterAdapterV2VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.RegisterAdapterV2VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.RegisterAdapterV2VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegisterAdapterV2Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.contract.Transact(opts, method, params...)
}

// AdapterInfo is a free data retrieval call binding the contract method 0xc10e0eeb.
//
// Solidity: function adapterInfo() view returns((address,string,string))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AdapterInfo(opts *bind.CallOpts) (RegisteryAdapterV2VaultsAdapterInfo, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "adapterInfo")

	if err != nil {
		return *new(RegisteryAdapterV2VaultsAdapterInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RegisteryAdapterV2VaultsAdapterInfo)).(*RegisteryAdapterV2VaultsAdapterInfo)

	return out0, err

}

// AdapterInfo is a free data retrieval call binding the contract method 0xc10e0eeb.
//
// Solidity: function adapterInfo() view returns((address,string,string))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AdapterInfo() (RegisteryAdapterV2VaultsAdapterInfo, error) {
	return _RegisterAdapterV2Vault.Contract.AdapterInfo(&_RegisterAdapterV2Vault.CallOpts)
}

// AdapterInfo is a free data retrieval call binding the contract method 0xc10e0eeb.
//
// Solidity: function adapterInfo() view returns((address,string,string))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AdapterInfo() (RegisteryAdapterV2VaultsAdapterInfo, error) {
	return _RegisterAdapterV2Vault.Contract.AdapterInfo(&_RegisterAdapterV2Vault.CallOpts)
}

// AdapterPositionOf is a free data retrieval call binding the contract method 0xb7cc58c1.
//
// Solidity: function adapterPositionOf(address accountAddress) view returns((uint256))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AdapterPositionOf(opts *bind.CallOpts, accountAddress common.Address) (RegisteryAdapterV2VaultsAdapterPosition, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "adapterPositionOf", accountAddress)

	if err != nil {
		return *new(RegisteryAdapterV2VaultsAdapterPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(RegisteryAdapterV2VaultsAdapterPosition)).(*RegisteryAdapterV2VaultsAdapterPosition)

	return out0, err

}

// AdapterPositionOf is a free data retrieval call binding the contract method 0xb7cc58c1.
//
// Solidity: function adapterPositionOf(address accountAddress) view returns((uint256))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AdapterPositionOf(accountAddress common.Address) (RegisteryAdapterV2VaultsAdapterPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AdapterPositionOf(&_RegisterAdapterV2Vault.CallOpts, accountAddress)
}

// AdapterPositionOf is a free data retrieval call binding the contract method 0xb7cc58c1.
//
// Solidity: function adapterPositionOf(address accountAddress) view returns((uint256))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AdapterPositionOf(accountAddress common.Address) (RegisteryAdapterV2VaultsAdapterPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AdapterPositionOf(&_RegisterAdapterV2Vault.CallOpts, accountAddress)
}

// AddressesGeneratorAddress is a free data retrieval call binding the contract method 0x59bd3909.
//
// Solidity: function addressesGeneratorAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AddressesGeneratorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "addressesGeneratorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesGeneratorAddress is a free data retrieval call binding the contract method 0x59bd3909.
//
// Solidity: function addressesGeneratorAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AddressesGeneratorAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AddressesGeneratorAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// AddressesGeneratorAddress is a free data retrieval call binding the contract method 0x59bd3909.
//
// Solidity: function addressesGeneratorAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AddressesGeneratorAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AddressesGeneratorAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetAllowances is a free data retrieval call binding the contract method 0xc6d0dc8b.
//
// Solidity: function assetAllowances(address accountAddress, address assetAddress) view returns((address,address,uint256)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetAllowances(opts *bind.CallOpts, accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsAllowance, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetAllowances", accountAddress, assetAddress)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAllowance), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAllowance)).(*[]RegisteryAdapterV2VaultsAllowance)

	return out0, err

}

// AssetAllowances is a free data retrieval call binding the contract method 0xc6d0dc8b.
//
// Solidity: function assetAllowances(address accountAddress, address assetAddress) view returns((address,address,uint256)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetAllowances(accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsAllowance, error) {
	return _RegisterAdapterV2Vault.Contract.AssetAllowances(&_RegisterAdapterV2Vault.CallOpts, accountAddress, assetAddress)
}

// AssetAllowances is a free data retrieval call binding the contract method 0xc6d0dc8b.
//
// Solidity: function assetAllowances(address accountAddress, address assetAddress) view returns((address,address,uint256)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetAllowances(accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsAllowance, error) {
	return _RegisterAdapterV2Vault.Contract.AssetAllowances(&_RegisterAdapterV2Vault.CallOpts, accountAddress, assetAddress)
}

// AssetBalance is a free data retrieval call binding the contract method 0xcd88e558.
//
// Solidity: function assetBalance(address assetAddress) view returns(uint256)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetBalance(opts *bind.CallOpts, assetAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetBalance", assetAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetBalance is a free data retrieval call binding the contract method 0xcd88e558.
//
// Solidity: function assetBalance(address assetAddress) view returns(uint256)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetBalance(assetAddress common.Address) (*big.Int, error) {
	return _RegisterAdapterV2Vault.Contract.AssetBalance(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetBalance is a free data retrieval call binding the contract method 0xcd88e558.
//
// Solidity: function assetBalance(address assetAddress) view returns(uint256)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetBalance(assetAddress common.Address) (*big.Int, error) {
	return _RegisterAdapterV2Vault.Contract.AssetBalance(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetDynamic is a free data retrieval call binding the contract method 0x69706fed.
//
// Solidity: function assetDynamic(address assetAddress) view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool)))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetDynamic(opts *bind.CallOpts, assetAddress common.Address) (RegisteryAdapterV2VaultsAssetDynamic, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetDynamic", assetAddress)

	if err != nil {
		return *new(RegisteryAdapterV2VaultsAssetDynamic), err
	}

	out0 := *abi.ConvertType(out[0], new(RegisteryAdapterV2VaultsAssetDynamic)).(*RegisteryAdapterV2VaultsAssetDynamic)

	return out0, err

}

// AssetDynamic is a free data retrieval call binding the contract method 0x69706fed.
//
// Solidity: function assetDynamic(address assetAddress) view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool)))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetDynamic(assetAddress common.Address) (RegisteryAdapterV2VaultsAssetDynamic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetDynamic(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetDynamic is a free data retrieval call binding the contract method 0x69706fed.
//
// Solidity: function assetDynamic(address assetAddress) view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool)))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetDynamic(assetAddress common.Address) (RegisteryAdapterV2VaultsAssetDynamic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetDynamic(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetPositionsOf is a free data retrieval call binding the contract method 0xe258f16a.
//
// Solidity: function assetPositionsOf(address accountAddress, address assetAddress) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetPositionsOf(opts *bind.CallOpts, accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetPositionsOf", accountAddress, assetAddress)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsPosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsPosition)).(*[]RegisteryAdapterV2VaultsPosition)

	return out0, err

}

// AssetPositionsOf is a free data retrieval call binding the contract method 0xe258f16a.
//
// Solidity: function assetPositionsOf(address accountAddress, address assetAddress) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetPositionsOf(accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AssetPositionsOf(&_RegisterAdapterV2Vault.CallOpts, accountAddress, assetAddress)
}

// AssetPositionsOf is a free data retrieval call binding the contract method 0xe258f16a.
//
// Solidity: function assetPositionsOf(address accountAddress, address assetAddress) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetPositionsOf(accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AssetPositionsOf(&_RegisterAdapterV2Vault.CallOpts, accountAddress, assetAddress)
}

// AssetStatic is a free data retrieval call binding the contract method 0x3d90e2c8.
//
// Solidity: function assetStatic(address assetAddress) view returns((address,string,address,string,string,string,uint8))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetStatic(opts *bind.CallOpts, assetAddress common.Address) (RegisteryAdapterV2VaultsAssetStatic, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetStatic", assetAddress)

	if err != nil {
		return *new(RegisteryAdapterV2VaultsAssetStatic), err
	}

	out0 := *abi.ConvertType(out[0], new(RegisteryAdapterV2VaultsAssetStatic)).(*RegisteryAdapterV2VaultsAssetStatic)

	return out0, err

}

// AssetStatic is a free data retrieval call binding the contract method 0x3d90e2c8.
//
// Solidity: function assetStatic(address assetAddress) view returns((address,string,address,string,string,string,uint8))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetStatic(assetAddress common.Address) (RegisteryAdapterV2VaultsAssetStatic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetStatic(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetStatic is a free data retrieval call binding the contract method 0x3d90e2c8.
//
// Solidity: function assetStatic(address assetAddress) view returns((address,string,address,string,string,string,uint8))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetStatic(assetAddress common.Address) (RegisteryAdapterV2VaultsAssetStatic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetStatic(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetUnderlyingTokenAddress is a free data retrieval call binding the contract method 0x532f4273.
//
// Solidity: function assetUnderlyingTokenAddress(address assetAddress) view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetUnderlyingTokenAddress(opts *bind.CallOpts, assetAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetUnderlyingTokenAddress", assetAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetUnderlyingTokenAddress is a free data retrieval call binding the contract method 0x532f4273.
//
// Solidity: function assetUnderlyingTokenAddress(address assetAddress) view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetUnderlyingTokenAddress(assetAddress common.Address) (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AssetUnderlyingTokenAddress(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetUnderlyingTokenAddress is a free data retrieval call binding the contract method 0x532f4273.
//
// Solidity: function assetUnderlyingTokenAddress(address assetAddress) view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetUnderlyingTokenAddress(assetAddress common.Address) (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AssetUnderlyingTokenAddress(&_RegisterAdapterV2Vault.CallOpts, assetAddress)
}

// AssetUserMetadata is a free data retrieval call binding the contract method 0xfdc9af8b.
//
// Solidity: function assetUserMetadata(address assetAddress, address accountAddress) view returns((address))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetUserMetadata(opts *bind.CallOpts, assetAddress common.Address, accountAddress common.Address) (RegisteryAdapterV2VaultsAssetUserMetadata, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetUserMetadata", assetAddress, accountAddress)

	if err != nil {
		return *new(RegisteryAdapterV2VaultsAssetUserMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(RegisteryAdapterV2VaultsAssetUserMetadata)).(*RegisteryAdapterV2VaultsAssetUserMetadata)

	return out0, err

}

// AssetUserMetadata is a free data retrieval call binding the contract method 0xfdc9af8b.
//
// Solidity: function assetUserMetadata(address assetAddress, address accountAddress) view returns((address))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetUserMetadata(assetAddress common.Address, accountAddress common.Address) (RegisteryAdapterV2VaultsAssetUserMetadata, error) {
	return _RegisterAdapterV2Vault.Contract.AssetUserMetadata(&_RegisterAdapterV2Vault.CallOpts, assetAddress, accountAddress)
}

// AssetUserMetadata is a free data retrieval call binding the contract method 0xfdc9af8b.
//
// Solidity: function assetUserMetadata(address assetAddress, address accountAddress) view returns((address))
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetUserMetadata(assetAddress common.Address, accountAddress common.Address) (RegisteryAdapterV2VaultsAssetUserMetadata, error) {
	return _RegisterAdapterV2Vault.Contract.AssetUserMetadata(&_RegisterAdapterV2Vault.CallOpts, assetAddress, accountAddress)
}

// AssetsAddresses is a free data retrieval call binding the contract method 0xa31091c7.
//
// Solidity: function assetsAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AssetsAddresses is a free data retrieval call binding the contract method 0xa31091c7.
//
// Solidity: function assetsAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsAddresses() ([]common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsAddresses(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsAddresses is a free data retrieval call binding the contract method 0xa31091c7.
//
// Solidity: function assetsAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsAddresses() ([]common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsAddresses(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsDynamic is a free data retrieval call binding the contract method 0x57d02836.
//
// Solidity: function assetsDynamic() view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool))[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsDynamic(opts *bind.CallOpts) ([]RegisteryAdapterV2VaultsAssetDynamic, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsDynamic")

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAssetDynamic), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAssetDynamic)).(*[]RegisteryAdapterV2VaultsAssetDynamic)

	return out0, err

}

// AssetsDynamic is a free data retrieval call binding the contract method 0x57d02836.
//
// Solidity: function assetsDynamic() view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool))[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsDynamic() ([]RegisteryAdapterV2VaultsAssetDynamic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsDynamic(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsDynamic is a free data retrieval call binding the contract method 0x57d02836.
//
// Solidity: function assetsDynamic() view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool))[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsDynamic() ([]RegisteryAdapterV2VaultsAssetDynamic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsDynamic(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsDynamic0 is a free data retrieval call binding the contract method 0x87920845.
//
// Solidity: function assetsDynamic(address[] _assetsAddresses) view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool))[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsDynamic0(opts *bind.CallOpts, _assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsAssetDynamic, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsDynamic0", _assetsAddresses)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAssetDynamic), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAssetDynamic)).(*[]RegisteryAdapterV2VaultsAssetDynamic)

	return out0, err

}

// AssetsDynamic0 is a free data retrieval call binding the contract method 0x87920845.
//
// Solidity: function assetsDynamic(address[] _assetsAddresses) view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool))[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsDynamic0(_assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsAssetDynamic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsDynamic0(&_RegisterAdapterV2Vault.CallOpts, _assetsAddresses)
}

// AssetsDynamic0 is a free data retrieval call binding the contract method 0x87920845.
//
// Solidity: function assetsDynamic(address[] _assetsAddresses) view returns((address,string,address,(uint256,uint256),(uint256,bool,address,uint256,bool))[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsDynamic0(_assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsAssetDynamic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsDynamic0(&_RegisterAdapterV2Vault.CallOpts, _assetsAddresses)
}

// AssetsLength is a free data retrieval call binding the contract method 0xf50477a2.
//
// Solidity: function assetsLength() view returns(uint256)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetsLength is a free data retrieval call binding the contract method 0xf50477a2.
//
// Solidity: function assetsLength() view returns(uint256)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsLength() (*big.Int, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsLength(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsLength is a free data retrieval call binding the contract method 0xf50477a2.
//
// Solidity: function assetsLength() view returns(uint256)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsLength() (*big.Int, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsLength(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsPositionsOf is a free data retrieval call binding the contract method 0xd33c39d2.
//
// Solidity: function assetsPositionsOf(address accountAddress, address[] _assetsAddresses) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsPositionsOf(opts *bind.CallOpts, accountAddress common.Address, _assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsPositionsOf", accountAddress, _assetsAddresses)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsPosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsPosition)).(*[]RegisteryAdapterV2VaultsPosition)

	return out0, err

}

// AssetsPositionsOf is a free data retrieval call binding the contract method 0xd33c39d2.
//
// Solidity: function assetsPositionsOf(address accountAddress, address[] _assetsAddresses) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsPositionsOf(accountAddress common.Address, _assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsPositionsOf(&_RegisterAdapterV2Vault.CallOpts, accountAddress, _assetsAddresses)
}

// AssetsPositionsOf is a free data retrieval call binding the contract method 0xd33c39d2.
//
// Solidity: function assetsPositionsOf(address accountAddress, address[] _assetsAddresses) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsPositionsOf(accountAddress common.Address, _assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsPositionsOf(&_RegisterAdapterV2Vault.CallOpts, accountAddress, _assetsAddresses)
}

// AssetsPositionsOf0 is a free data retrieval call binding the contract method 0xe23121b1.
//
// Solidity: function assetsPositionsOf(address accountAddress) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsPositionsOf0(opts *bind.CallOpts, accountAddress common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsPositionsOf0", accountAddress)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsPosition), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsPosition)).(*[]RegisteryAdapterV2VaultsPosition)

	return out0, err

}

// AssetsPositionsOf0 is a free data retrieval call binding the contract method 0xe23121b1.
//
// Solidity: function assetsPositionsOf(address accountAddress) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsPositionsOf0(accountAddress common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsPositionsOf0(&_RegisterAdapterV2Vault.CallOpts, accountAddress)
}

// AssetsPositionsOf0 is a free data retrieval call binding the contract method 0xe23121b1.
//
// Solidity: function assetsPositionsOf(address accountAddress) view returns((address,address,string,uint256,(uint256,uint256),(address,address,uint256)[],(address,address,uint256)[])[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsPositionsOf0(accountAddress common.Address) ([]RegisteryAdapterV2VaultsPosition, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsPositionsOf0(&_RegisterAdapterV2Vault.CallOpts, accountAddress)
}

// AssetsStatic is a free data retrieval call binding the contract method 0x9adbba59.
//
// Solidity: function assetsStatic() view returns((address,string,address,string,string,string,uint8)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsStatic(opts *bind.CallOpts) ([]RegisteryAdapterV2VaultsAssetStatic, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsStatic")

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAssetStatic), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAssetStatic)).(*[]RegisteryAdapterV2VaultsAssetStatic)

	return out0, err

}

// AssetsStatic is a free data retrieval call binding the contract method 0x9adbba59.
//
// Solidity: function assetsStatic() view returns((address,string,address,string,string,string,uint8)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsStatic() ([]RegisteryAdapterV2VaultsAssetStatic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsStatic(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsStatic is a free data retrieval call binding the contract method 0x9adbba59.
//
// Solidity: function assetsStatic() view returns((address,string,address,string,string,string,uint8)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsStatic() ([]RegisteryAdapterV2VaultsAssetStatic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsStatic(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsStatic0 is a free data retrieval call binding the contract method 0xd36ec1cf.
//
// Solidity: function assetsStatic(address[] _assetsAddresses) view returns((address,string,address,string,string,string,uint8)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsStatic0(opts *bind.CallOpts, _assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsAssetStatic, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsStatic0", _assetsAddresses)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAssetStatic), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAssetStatic)).(*[]RegisteryAdapterV2VaultsAssetStatic)

	return out0, err

}

// AssetsStatic0 is a free data retrieval call binding the contract method 0xd36ec1cf.
//
// Solidity: function assetsStatic(address[] _assetsAddresses) view returns((address,string,address,string,string,string,uint8)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsStatic0(_assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsAssetStatic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsStatic0(&_RegisterAdapterV2Vault.CallOpts, _assetsAddresses)
}

// AssetsStatic0 is a free data retrieval call binding the contract method 0xd36ec1cf.
//
// Solidity: function assetsStatic(address[] _assetsAddresses) view returns((address,string,address,string,string,string,uint8)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsStatic0(_assetsAddresses []common.Address) ([]RegisteryAdapterV2VaultsAssetStatic, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsStatic0(&_RegisterAdapterV2Vault.CallOpts, _assetsAddresses)
}

// AssetsTokensAddresses is a free data retrieval call binding the contract method 0xa2f93565.
//
// Solidity: function assetsTokensAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsTokensAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsTokensAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AssetsTokensAddresses is a free data retrieval call binding the contract method 0xa2f93565.
//
// Solidity: function assetsTokensAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsTokensAddresses() ([]common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsTokensAddresses(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsTokensAddresses is a free data retrieval call binding the contract method 0xa2f93565.
//
// Solidity: function assetsTokensAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsTokensAddresses() ([]common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsTokensAddresses(&_RegisterAdapterV2Vault.CallOpts)
}

// AssetsUserMetadata is a free data retrieval call binding the contract method 0x8b185f36.
//
// Solidity: function assetsUserMetadata(address accountAddress) view returns((address)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) AssetsUserMetadata(opts *bind.CallOpts, accountAddress common.Address) ([]RegisteryAdapterV2VaultsAssetUserMetadata, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "assetsUserMetadata", accountAddress)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAssetUserMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAssetUserMetadata)).(*[]RegisteryAdapterV2VaultsAssetUserMetadata)

	return out0, err

}

// AssetsUserMetadata is a free data retrieval call binding the contract method 0x8b185f36.
//
// Solidity: function assetsUserMetadata(address accountAddress) view returns((address)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) AssetsUserMetadata(accountAddress common.Address) ([]RegisteryAdapterV2VaultsAssetUserMetadata, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsUserMetadata(&_RegisterAdapterV2Vault.CallOpts, accountAddress)
}

// AssetsUserMetadata is a free data retrieval call binding the contract method 0x8b185f36.
//
// Solidity: function assetsUserMetadata(address accountAddress) view returns((address)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) AssetsUserMetadata(accountAddress common.Address) ([]RegisteryAdapterV2VaultsAssetUserMetadata, error) {
	return _RegisterAdapterV2Vault.Contract.AssetsUserMetadata(&_RegisterAdapterV2Vault.CallOpts, accountAddress)
}

// ExtensionsAddresses is a free data retrieval call binding the contract method 0xb618e5c3.
//
// Solidity: function extensionsAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) ExtensionsAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "extensionsAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExtensionsAddresses is a free data retrieval call binding the contract method 0xb618e5c3.
//
// Solidity: function extensionsAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) ExtensionsAddresses() ([]common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.ExtensionsAddresses(&_RegisterAdapterV2Vault.CallOpts)
}

// ExtensionsAddresses is a free data retrieval call binding the contract method 0xb618e5c3.
//
// Solidity: function extensionsAddresses() view returns(address[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) ExtensionsAddresses() ([]common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.ExtensionsAddresses(&_RegisterAdapterV2Vault.CallOpts)
}

// HelperAddress is a free data retrieval call binding the contract method 0x7974db4b.
//
// Solidity: function helperAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) HelperAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "helperAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HelperAddress is a free data retrieval call binding the contract method 0x7974db4b.
//
// Solidity: function helperAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) HelperAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.HelperAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// HelperAddress is a free data retrieval call binding the contract method 0x7974db4b.
//
// Solidity: function helperAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) HelperAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.HelperAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) OracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "oracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) OracleAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.OracleAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) OracleAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.OracleAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) Owner() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.Owner(&_RegisterAdapterV2Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) Owner() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.Owner(&_RegisterAdapterV2Vault.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) RegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "registryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) RegistryAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.RegistryAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() view returns(address)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) RegistryAddress() (common.Address, error) {
	return _RegisterAdapterV2Vault.Contract.RegistryAddress(&_RegisterAdapterV2Vault.CallOpts)
}

// SupportedPositions is a free data retrieval call binding the contract method 0x294e9131.
//
// Solidity: function supportedPositions(uint256 ) view returns(string)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) SupportedPositions(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "supportedPositions", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SupportedPositions is a free data retrieval call binding the contract method 0x294e9131.
//
// Solidity: function supportedPositions(uint256 ) view returns(string)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) SupportedPositions(arg0 *big.Int) (string, error) {
	return _RegisterAdapterV2Vault.Contract.SupportedPositions(&_RegisterAdapterV2Vault.CallOpts, arg0)
}

// SupportedPositions is a free data retrieval call binding the contract method 0x294e9131.
//
// Solidity: function supportedPositions(uint256 ) view returns(string)
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) SupportedPositions(arg0 *big.Int) (string, error) {
	return _RegisterAdapterV2Vault.Contract.SupportedPositions(&_RegisterAdapterV2Vault.CallOpts, arg0)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xd68bda7c.
//
// Solidity: function tokenAllowances(address accountAddress, address assetAddress) view returns((address,address,uint256)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCaller) TokenAllowances(opts *bind.CallOpts, accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsAllowance, error) {
	var out []interface{}
	err := _RegisterAdapterV2Vault.contract.Call(opts, &out, "tokenAllowances", accountAddress, assetAddress)

	if err != nil {
		return *new([]RegisteryAdapterV2VaultsAllowance), err
	}

	out0 := *abi.ConvertType(out[0], new([]RegisteryAdapterV2VaultsAllowance)).(*[]RegisteryAdapterV2VaultsAllowance)

	return out0, err

}

// TokenAllowances is a free data retrieval call binding the contract method 0xd68bda7c.
//
// Solidity: function tokenAllowances(address accountAddress, address assetAddress) view returns((address,address,uint256)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) TokenAllowances(accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsAllowance, error) {
	return _RegisterAdapterV2Vault.Contract.TokenAllowances(&_RegisterAdapterV2Vault.CallOpts, accountAddress, assetAddress)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xd68bda7c.
//
// Solidity: function tokenAllowances(address accountAddress, address assetAddress) view returns((address,address,uint256)[])
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultCallerSession) TokenAllowances(accountAddress common.Address, assetAddress common.Address) ([]RegisteryAdapterV2VaultsAllowance, error) {
	return _RegisterAdapterV2Vault.Contract.TokenAllowances(&_RegisterAdapterV2Vault.CallOpts, accountAddress, assetAddress)
}

// SetExtensionsAddresses is a paid mutator transaction binding the contract method 0x77bcb01f.
//
// Solidity: function setExtensionsAddresses(address[] _newExtensionsAddresses) returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactor) SetExtensionsAddresses(opts *bind.TransactOpts, _newExtensionsAddresses []common.Address) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.contract.Transact(opts, "setExtensionsAddresses", _newExtensionsAddresses)
}

// SetExtensionsAddresses is a paid mutator transaction binding the contract method 0x77bcb01f.
//
// Solidity: function setExtensionsAddresses(address[] _newExtensionsAddresses) returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) SetExtensionsAddresses(_newExtensionsAddresses []common.Address) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.SetExtensionsAddresses(&_RegisterAdapterV2Vault.TransactOpts, _newExtensionsAddresses)
}

// SetExtensionsAddresses is a paid mutator transaction binding the contract method 0x77bcb01f.
//
// Solidity: function setExtensionsAddresses(address[] _newExtensionsAddresses) returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactorSession) SetExtensionsAddresses(_newExtensionsAddresses []common.Address) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.SetExtensionsAddresses(&_RegisterAdapterV2Vault.TransactOpts, _newExtensionsAddresses)
}

// UpdateSlot is a paid mutator transaction binding the contract method 0x91ea83e8.
//
// Solidity: function updateSlot(bytes32 slot, bytes32 value) returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactor) UpdateSlot(opts *bind.TransactOpts, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.contract.Transact(opts, "updateSlot", slot, value)
}

// UpdateSlot is a paid mutator transaction binding the contract method 0x91ea83e8.
//
// Solidity: function updateSlot(bytes32 slot, bytes32 value) returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) UpdateSlot(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.UpdateSlot(&_RegisterAdapterV2Vault.TransactOpts, slot, value)
}

// UpdateSlot is a paid mutator transaction binding the contract method 0x91ea83e8.
//
// Solidity: function updateSlot(bytes32 slot, bytes32 value) returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactorSession) UpdateSlot(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.UpdateSlot(&_RegisterAdapterV2Vault.TransactOpts, slot, value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.Fallback(&_RegisterAdapterV2Vault.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_RegisterAdapterV2Vault *RegisterAdapterV2VaultTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RegisterAdapterV2Vault.Contract.Fallback(&_RegisterAdapterV2Vault.TransactOpts, calldata)
}
