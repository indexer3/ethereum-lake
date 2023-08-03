// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venus

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

// VenusComprollerMetaData contains all meta data concerning the VenusComproller contract.
var VenusComprollerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPausedMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"ActionProtocolPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusVAIMintIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIMinterVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIVaultVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldComptrollerLens\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newComptrollerLens\",\"type\":\"address\"}],\"name\":\"NewComptrollerLens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLiquidatorContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLiquidatorContract\",\"type\":\"address\"}],\"name\":\"NewLiquidatorContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"oldVAIController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"newVAIController\",\"type\":\"address\"}],\"name\":\"NewVAIController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVAIMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"NewVAIMintRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseInterval_\",\"type\":\"uint256\"}],\"name\":\"NewVAIVaultInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIVaultRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIVaultRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIVaultRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"VenusGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"VenusSpeedUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_grantXVS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptrollerLens_\",\"type\":\"address\"}],\"name\":\"_setComptrollerLens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLiquidatorContract_\",\"type\":\"address\"}],\"name\":\"_setLiquidatorContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setProtocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"vaiController_\",\"type\":\"address\"}],\"name\":\"_setVAIController\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"_setVAIMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReleaseAmount_\",\"type\":\"uint256\"}],\"name\":\"_setVAIVaultInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"venusSpeed\",\"type\":\"uint256\"}],\"name\":\"_setVenusSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIVaultRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIVaultRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"collateral\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenusAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerLens\",\"outputs\":[{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getXVSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getXVSVTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContributorBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateVAICalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVenus\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minReleaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedVAIs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repayVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintedVAIOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiController\",\"outputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusContributorSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIVaultRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VenusComprollerABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusComprollerMetaData.ABI instead.
var VenusComprollerABI = VenusComprollerMetaData.ABI

// VenusComproller is an auto generated Go binding around an Ethereum contract.
type VenusComproller struct {
	VenusComprollerCaller     // Read-only binding to the contract
	VenusComprollerTransactor // Write-only binding to the contract
	VenusComprollerFilterer   // Log filterer for contract events
}

// VenusComprollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusComprollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusComprollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusComprollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusComprollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusComprollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusComprollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusComprollerSession struct {
	Contract     *VenusComproller  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenusComprollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusComprollerCallerSession struct {
	Contract *VenusComprollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VenusComprollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusComprollerTransactorSession struct {
	Contract     *VenusComprollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VenusComprollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusComprollerRaw struct {
	Contract *VenusComproller // Generic contract binding to access the raw methods on
}

// VenusComprollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusComprollerCallerRaw struct {
	Contract *VenusComprollerCaller // Generic read-only contract binding to access the raw methods on
}

// VenusComprollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusComprollerTransactorRaw struct {
	Contract *VenusComprollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusComproller creates a new instance of VenusComproller, bound to a specific deployed contract.
func NewVenusComproller(address common.Address, backend bind.ContractBackend) (*VenusComproller, error) {
	contract, err := bindVenusComproller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusComproller{VenusComprollerCaller: VenusComprollerCaller{contract: contract}, VenusComprollerTransactor: VenusComprollerTransactor{contract: contract}, VenusComprollerFilterer: VenusComprollerFilterer{contract: contract}}, nil
}

// NewVenusComprollerCaller creates a new read-only instance of VenusComproller, bound to a specific deployed contract.
func NewVenusComprollerCaller(address common.Address, caller bind.ContractCaller) (*VenusComprollerCaller, error) {
	contract, err := bindVenusComproller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerCaller{contract: contract}, nil
}

// NewVenusComprollerTransactor creates a new write-only instance of VenusComproller, bound to a specific deployed contract.
func NewVenusComprollerTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusComprollerTransactor, error) {
	contract, err := bindVenusComproller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerTransactor{contract: contract}, nil
}

// NewVenusComprollerFilterer creates a new log filterer instance of VenusComproller, bound to a specific deployed contract.
func NewVenusComprollerFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusComprollerFilterer, error) {
	contract, err := bindVenusComproller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerFilterer{contract: contract}, nil
}

// bindVenusComproller binds a generic wrapper to an already deployed contract.
func bindVenusComproller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VenusComprollerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusComproller *VenusComprollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusComproller.Contract.VenusComprollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusComproller *VenusComprollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusComproller.Contract.VenusComprollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusComproller *VenusComprollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusComproller.Contract.VenusComprollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusComproller *VenusComprollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusComproller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusComproller *VenusComprollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusComproller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusComproller *VenusComprollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusComproller.Contract.contract.Transact(opts, method, params...)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) BorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerSession) BorrowGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.BorrowGuardianPaused(&_VenusComproller.CallOpts)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) BorrowGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.BorrowGuardianPaused(&_VenusComproller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_VenusComproller *VenusComprollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_VenusComproller *VenusComprollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _VenusComproller.Contract.AccountAssets(&_VenusComproller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _VenusComproller.Contract.AccountAssets(&_VenusComproller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusComproller *VenusComprollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusComproller *VenusComprollerSession) Admin() (common.Address, error) {
	return _VenusComproller.Contract.Admin(&_VenusComproller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) Admin() (common.Address, error) {
	return _VenusComproller.Contract.Admin(&_VenusComproller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_VenusComproller *VenusComprollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_VenusComproller *VenusComprollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _VenusComproller.Contract.AllMarkets(&_VenusComproller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _VenusComproller.Contract.AllMarkets(&_VenusComproller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_VenusComproller *VenusComprollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_VenusComproller *VenusComprollerSession) BorrowCapGuardian() (common.Address, error) {
	return _VenusComproller.Contract.BorrowCapGuardian(&_VenusComproller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _VenusComproller.Contract.BorrowCapGuardian(&_VenusComproller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.BorrowCaps(&_VenusComproller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.BorrowCaps(&_VenusComproller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_VenusComproller *VenusComprollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, vToken common.Address) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "checkMembership", account, vToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_VenusComproller *VenusComprollerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _VenusComproller.Contract.CheckMembership(&_VenusComproller.CallOpts, account, vToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _VenusComproller.Contract.CheckMembership(&_VenusComproller.CallOpts, account, vToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _VenusComproller.Contract.CloseFactorMantissa(&_VenusComproller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _VenusComproller.Contract.CloseFactorMantissa(&_VenusComproller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_VenusComproller *VenusComprollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_VenusComproller *VenusComprollerSession) ComptrollerImplementation() (common.Address, error) {
	return _VenusComproller.Contract.ComptrollerImplementation(&_VenusComproller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _VenusComproller.Contract.ComptrollerImplementation(&_VenusComproller.CallOpts)
}

// ComptrollerLens is a free data retrieval call binding the contract method 0xd3270f99.
//
// Solidity: function comptrollerLens() view returns(address)
func (_VenusComproller *VenusComprollerCaller) ComptrollerLens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "comptrollerLens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerLens is a free data retrieval call binding the contract method 0xd3270f99.
//
// Solidity: function comptrollerLens() view returns(address)
func (_VenusComproller *VenusComprollerSession) ComptrollerLens() (common.Address, error) {
	return _VenusComproller.Contract.ComptrollerLens(&_VenusComproller.CallOpts)
}

// ComptrollerLens is a free data retrieval call binding the contract method 0xd3270f99.
//
// Solidity: function comptrollerLens() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) ComptrollerLens() (common.Address, error) {
	return _VenusComproller.Contract.ComptrollerLens(&_VenusComproller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_VenusComproller *VenusComprollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_VenusComproller *VenusComprollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusComproller.Contract.GetAccountLiquidity(&_VenusComproller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_VenusComproller *VenusComprollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusComproller.Contract.GetAccountLiquidity(&_VenusComproller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_VenusComproller *VenusComprollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_VenusComproller *VenusComprollerSession) GetAllMarkets() ([]common.Address, error) {
	return _VenusComproller.Contract.GetAllMarkets(&_VenusComproller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_VenusComproller *VenusComprollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _VenusComproller.Contract.GetAllMarkets(&_VenusComproller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_VenusComproller *VenusComprollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_VenusComproller *VenusComprollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _VenusComproller.Contract.GetAssetsIn(&_VenusComproller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_VenusComproller *VenusComprollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _VenusComproller.Contract.GetAssetsIn(&_VenusComproller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) GetBlockNumber() (*big.Int, error) {
	return _VenusComproller.Contract.GetBlockNumber(&_VenusComproller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _VenusComproller.Contract.GetBlockNumber(&_VenusComproller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_VenusComproller *VenusComprollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, vTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_VenusComproller *VenusComprollerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusComproller.Contract.GetHypotheticalAccountLiquidity(&_VenusComproller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_VenusComproller *VenusComprollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _VenusComproller.Contract.GetHypotheticalAccountLiquidity(&_VenusComproller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_VenusComproller *VenusComprollerCaller) GetXVSAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getXVSAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_VenusComproller *VenusComprollerSession) GetXVSAddress() (common.Address, error) {
	return _VenusComproller.Contract.GetXVSAddress(&_VenusComproller.CallOpts)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) GetXVSAddress() (common.Address, error) {
	return _VenusComproller.Contract.GetXVSAddress(&_VenusComproller.CallOpts)
}

// GetXVSVTokenAddress is a free data retrieval call binding the contract method 0xededbae6.
//
// Solidity: function getXVSVTokenAddress() view returns(address)
func (_VenusComproller *VenusComprollerCaller) GetXVSVTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "getXVSVTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetXVSVTokenAddress is a free data retrieval call binding the contract method 0xededbae6.
//
// Solidity: function getXVSVTokenAddress() view returns(address)
func (_VenusComproller *VenusComprollerSession) GetXVSVTokenAddress() (common.Address, error) {
	return _VenusComproller.Contract.GetXVSVTokenAddress(&_VenusComproller.CallOpts)
}

// GetXVSVTokenAddress is a free data retrieval call binding the contract method 0xededbae6.
//
// Solidity: function getXVSVTokenAddress() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) GetXVSVTokenAddress() (common.Address, error) {
	return _VenusComproller.Contract.GetXVSVTokenAddress(&_VenusComproller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_VenusComproller *VenusComprollerSession) IsComptroller() (bool, error) {
	return _VenusComproller.Contract.IsComptroller(&_VenusComproller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) IsComptroller() (bool, error) {
	return _VenusComproller.Contract.IsComptroller(&_VenusComproller.CallOpts)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) LastContributorBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "lastContributorBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.LastContributorBlock(&_VenusComproller.CallOpts, arg0)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.LastContributorBlock(&_VenusComproller.CallOpts, arg0)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusComproller *VenusComprollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", vTokenBorrowed, vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusComproller *VenusComprollerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusComproller.Contract.LiquidateCalculateSeizeTokens(&_VenusComproller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusComproller *VenusComprollerCallerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusComproller.Contract.LiquidateCalculateSeizeTokens(&_VenusComproller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusComproller *VenusComprollerCaller) LiquidateVAICalculateSeizeTokens(opts *bind.CallOpts, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "liquidateVAICalculateSeizeTokens", vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusComproller *VenusComprollerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusComproller.Contract.LiquidateVAICalculateSeizeTokens(&_VenusComproller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_VenusComproller *VenusComprollerCallerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _VenusComproller.Contract.LiquidateVAICalculateSeizeTokens(&_VenusComproller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _VenusComproller.Contract.LiquidationIncentiveMantissa(&_VenusComproller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _VenusComproller.Contract.LiquidationIncentiveMantissa(&_VenusComproller.CallOpts)
}

// LiquidatorContract is a free data retrieval call binding the contract method 0x9bb27d62.
//
// Solidity: function liquidatorContract() view returns(address)
func (_VenusComproller *VenusComprollerCaller) LiquidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "liquidatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidatorContract is a free data retrieval call binding the contract method 0x9bb27d62.
//
// Solidity: function liquidatorContract() view returns(address)
func (_VenusComproller *VenusComprollerSession) LiquidatorContract() (common.Address, error) {
	return _VenusComproller.Contract.LiquidatorContract(&_VenusComproller.CallOpts)
}

// LiquidatorContract is a free data retrieval call binding the contract method 0x9bb27d62.
//
// Solidity: function liquidatorContract() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) LiquidatorContract() (common.Address, error) {
	return _VenusComproller.Contract.LiquidatorContract(&_VenusComproller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_VenusComproller *VenusComprollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsVenus                  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsVenus = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_VenusComproller *VenusComprollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _VenusComproller.Contract.Markets(&_VenusComproller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_VenusComproller *VenusComprollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _VenusComproller.Contract.Markets(&_VenusComproller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) MaxAssets() (*big.Int, error) {
	return _VenusComproller.Contract.MaxAssets(&_VenusComproller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) MaxAssets() (*big.Int, error) {
	return _VenusComproller.Contract.MaxAssets(&_VenusComproller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) MinReleaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "minReleaseAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) MinReleaseAmount() (*big.Int, error) {
	return _VenusComproller.Contract.MinReleaseAmount(&_VenusComproller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) MinReleaseAmount() (*big.Int, error) {
	return _VenusComproller.Contract.MinReleaseAmount(&_VenusComproller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_VenusComproller *VenusComprollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_VenusComproller *VenusComprollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _VenusComproller.Contract.MintGuardianPaused(&_VenusComproller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _VenusComproller.Contract.MintGuardianPaused(&_VenusComproller.CallOpts, arg0)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) MintVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "mintVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerSession) MintVAIGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.MintVAIGuardianPaused(&_VenusComproller.CallOpts)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) MintVAIGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.MintVAIGuardianPaused(&_VenusComproller.CallOpts)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) MintedVAIs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "mintedVAIs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.MintedVAIs(&_VenusComproller.CallOpts, arg0)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.MintedVAIs(&_VenusComproller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_VenusComproller *VenusComprollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_VenusComproller *VenusComprollerSession) Oracle() (common.Address, error) {
	return _VenusComproller.Contract.Oracle(&_VenusComproller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) Oracle() (common.Address, error) {
	return _VenusComproller.Contract.Oracle(&_VenusComproller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_VenusComproller *VenusComprollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_VenusComproller *VenusComprollerSession) PauseGuardian() (common.Address, error) {
	return _VenusComproller.Contract.PauseGuardian(&_VenusComproller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) PauseGuardian() (common.Address, error) {
	return _VenusComproller.Contract.PauseGuardian(&_VenusComproller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusComproller *VenusComprollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusComproller *VenusComprollerSession) PendingAdmin() (common.Address, error) {
	return _VenusComproller.Contract.PendingAdmin(&_VenusComproller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) PendingAdmin() (common.Address, error) {
	return _VenusComproller.Contract.PendingAdmin(&_VenusComproller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_VenusComproller *VenusComprollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_VenusComproller *VenusComprollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _VenusComproller.Contract.PendingComptrollerImplementation(&_VenusComproller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _VenusComproller.Contract.PendingComptrollerImplementation(&_VenusComproller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) ProtocolPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "protocolPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_VenusComproller *VenusComprollerSession) ProtocolPaused() (bool, error) {
	return _VenusComproller.Contract.ProtocolPaused(&_VenusComproller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) ProtocolPaused() (bool, error) {
	return _VenusComproller.Contract.ProtocolPaused(&_VenusComproller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) ReleaseStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "releaseStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) ReleaseStartBlock() (*big.Int, error) {
	return _VenusComproller.Contract.ReleaseStartBlock(&_VenusComproller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) ReleaseStartBlock() (*big.Int, error) {
	return _VenusComproller.Contract.ReleaseStartBlock(&_VenusComproller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) RepayVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "repayVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerSession) RepayVAIGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.RepayVAIGuardianPaused(&_VenusComproller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) RepayVAIGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.RepayVAIGuardianPaused(&_VenusComproller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerSession) SeizeGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.SeizeGuardianPaused(&_VenusComproller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.SeizeGuardianPaused(&_VenusComproller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerSession) TransferGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.TransferGuardianPaused(&_VenusComproller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_VenusComproller *VenusComprollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _VenusComproller.Contract.TransferGuardianPaused(&_VenusComproller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VenusComproller *VenusComprollerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VenusComproller *VenusComprollerSession) TreasuryAddress() (common.Address, error) {
	return _VenusComproller.Contract.TreasuryAddress(&_VenusComproller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) TreasuryAddress() (common.Address, error) {
	return _VenusComproller.Contract.TreasuryAddress(&_VenusComproller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VenusComproller *VenusComprollerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "treasuryGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VenusComproller *VenusComprollerSession) TreasuryGuardian() (common.Address, error) {
	return _VenusComproller.Contract.TreasuryGuardian(&_VenusComproller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _VenusComproller.Contract.TreasuryGuardian(&_VenusComproller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) TreasuryPercent() (*big.Int, error) {
	return _VenusComproller.Contract.TreasuryPercent(&_VenusComproller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _VenusComproller.Contract.TreasuryPercent(&_VenusComproller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_VenusComproller *VenusComprollerCaller) VaiController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "vaiController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_VenusComproller *VenusComprollerSession) VaiController() (common.Address, error) {
	return _VenusComproller.Contract.VaiController(&_VenusComproller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) VaiController() (common.Address, error) {
	return _VenusComproller.Contract.VaiController(&_VenusComproller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VaiMintRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "vaiMintRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VaiMintRate() (*big.Int, error) {
	return _VenusComproller.Contract.VaiMintRate(&_VenusComproller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VaiMintRate() (*big.Int, error) {
	return _VenusComproller.Contract.VaiMintRate(&_VenusComproller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_VenusComproller *VenusComprollerCaller) VaiVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "vaiVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_VenusComproller *VenusComprollerSession) VaiVaultAddress() (common.Address, error) {
	return _VenusComproller.Contract.VaiVaultAddress(&_VenusComproller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_VenusComproller *VenusComprollerCallerSession) VaiVaultAddress() (common.Address, error) {
	return _VenusComproller.Contract.VaiVaultAddress(&_VenusComproller.CallOpts)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusAccrued(&_VenusComproller.CallOpts, arg0)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusAccrued(&_VenusComproller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_VenusComproller *VenusComprollerCaller) VenusBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_VenusComproller *VenusComprollerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusComproller.Contract.VenusBorrowState(&_VenusComproller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_VenusComproller *VenusComprollerCallerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusComproller.Contract.VenusBorrowState(&_VenusComproller.CallOpts, arg0)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusBorrowerIndex(&_VenusComproller.CallOpts, arg0, arg1)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusBorrowerIndex(&_VenusComproller.CallOpts, arg0, arg1)
}

// VenusContributorSpeeds is a free data retrieval call binding the contract method 0xa9046134.
//
// Solidity: function venusContributorSpeeds(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusContributorSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusContributorSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusContributorSpeeds is a free data retrieval call binding the contract method 0xa9046134.
//
// Solidity: function venusContributorSpeeds(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusContributorSpeeds(&_VenusComproller.CallOpts, arg0)
}

// VenusContributorSpeeds is a free data retrieval call binding the contract method 0xa9046134.
//
// Solidity: function venusContributorSpeeds(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusContributorSpeeds(&_VenusComproller.CallOpts, arg0)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VenusComproller *VenusComprollerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VenusComproller *VenusComprollerSession) VenusInitialIndex() (*big.Int, error) {
	return _VenusComproller.Contract.VenusInitialIndex(&_VenusComproller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_VenusComproller *VenusComprollerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _VenusComproller.Contract.VenusInitialIndex(&_VenusComproller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusRate() (*big.Int, error) {
	return _VenusComproller.Contract.VenusRate(&_VenusComproller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusRate() (*big.Int, error) {
	return _VenusComproller.Contract.VenusRate(&_VenusComproller.CallOpts)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusSpeeds(&_VenusComproller.CallOpts, arg0)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusSpeeds(&_VenusComproller.CallOpts, arg0)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusSupplierIndex(&_VenusComproller.CallOpts, arg0, arg1)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusComproller.Contract.VenusSupplierIndex(&_VenusComproller.CallOpts, arg0, arg1)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_VenusComproller *VenusComprollerCaller) VenusSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_VenusComproller *VenusComprollerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusComproller.Contract.VenusSupplyState(&_VenusComproller.CallOpts, arg0)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_VenusComproller *VenusComprollerCallerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _VenusComproller.Contract.VenusSupplyState(&_VenusComproller.CallOpts, arg0)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusVAIRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusVAIRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusVAIRate() (*big.Int, error) {
	return _VenusComproller.Contract.VenusVAIRate(&_VenusComproller.CallOpts)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusVAIRate() (*big.Int, error) {
	return _VenusComproller.Contract.VenusVAIRate(&_VenusComproller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCaller) VenusVAIVaultRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusComproller.contract.Call(opts, &out, "venusVAIVaultRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_VenusComproller *VenusComprollerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _VenusComproller.Contract.VenusVAIVaultRate(&_VenusComproller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_VenusComproller *VenusComprollerCallerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _VenusComproller.Contract.VenusVAIVaultRate(&_VenusComproller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VenusComproller *VenusComprollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VenusComproller *VenusComprollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.Become(&_VenusComproller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_VenusComproller *VenusComprollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.Become(&_VenusComproller.TransactOpts, unitroller)
}

// GrantXVS is a paid mutator transaction binding the contract method 0xa7604b41.
//
// Solidity: function _grantXVS(address recipient, uint256 amount) returns()
func (_VenusComproller *VenusComprollerTransactor) GrantXVS(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_grantXVS", recipient, amount)
}

// GrantXVS is a paid mutator transaction binding the contract method 0xa7604b41.
//
// Solidity: function _grantXVS(address recipient, uint256 amount) returns()
func (_VenusComproller *VenusComprollerSession) GrantXVS(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.GrantXVS(&_VenusComproller.TransactOpts, recipient, amount)
}

// GrantXVS is a paid mutator transaction binding the contract method 0xa7604b41.
//
// Solidity: function _grantXVS(address recipient, uint256 amount) returns()
func (_VenusComproller *VenusComprollerTransactorSession) GrantXVS(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.GrantXVS(&_VenusComproller.TransactOpts, recipient, amount)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_VenusComproller *VenusComprollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_VenusComproller *VenusComprollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetBorrowCapGuardian(&_VenusComproller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetBorrowCapGuardian(&_VenusComproller.TransactOpts, newBorrowCapGuardian)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetCloseFactor(&_VenusComproller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetCloseFactor(&_VenusComproller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setCollateralFactor", vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetCollateralFactor(&_VenusComproller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetCollateralFactor(&_VenusComproller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetComptrollerLens is a paid mutator transaction binding the contract method 0x9bf34cbb.
//
// Solidity: function _setComptrollerLens(address comptrollerLens_) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetComptrollerLens(opts *bind.TransactOpts, comptrollerLens_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setComptrollerLens", comptrollerLens_)
}

// SetComptrollerLens is a paid mutator transaction binding the contract method 0x9bf34cbb.
//
// Solidity: function _setComptrollerLens(address comptrollerLens_) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetComptrollerLens(comptrollerLens_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetComptrollerLens(&_VenusComproller.TransactOpts, comptrollerLens_)
}

// SetComptrollerLens is a paid mutator transaction binding the contract method 0x9bf34cbb.
//
// Solidity: function _setComptrollerLens(address comptrollerLens_) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetComptrollerLens(comptrollerLens_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetComptrollerLens(&_VenusComproller.TransactOpts, comptrollerLens_)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetLiquidationIncentive(&_VenusComproller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetLiquidationIncentive(&_VenusComproller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidatorContract is a paid mutator transaction binding the contract method 0xbb857450.
//
// Solidity: function _setLiquidatorContract(address newLiquidatorContract_) returns()
func (_VenusComproller *VenusComprollerTransactor) SetLiquidatorContract(opts *bind.TransactOpts, newLiquidatorContract_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setLiquidatorContract", newLiquidatorContract_)
}

// SetLiquidatorContract is a paid mutator transaction binding the contract method 0xbb857450.
//
// Solidity: function _setLiquidatorContract(address newLiquidatorContract_) returns()
func (_VenusComproller *VenusComprollerSession) SetLiquidatorContract(newLiquidatorContract_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetLiquidatorContract(&_VenusComproller.TransactOpts, newLiquidatorContract_)
}

// SetLiquidatorContract is a paid mutator transaction binding the contract method 0xbb857450.
//
// Solidity: function _setLiquidatorContract(address newLiquidatorContract_) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SetLiquidatorContract(newLiquidatorContract_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetLiquidatorContract(&_VenusComproller.TransactOpts, newLiquidatorContract_)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_VenusComproller *VenusComprollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setMarketBorrowCaps", vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_VenusComproller *VenusComprollerSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetMarketBorrowCaps(&_VenusComproller.TransactOpts, vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetMarketBorrowCaps(&_VenusComproller.TransactOpts, vTokens, newBorrowCaps)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetPauseGuardian(&_VenusComproller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetPauseGuardian(&_VenusComproller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetPriceOracle(&_VenusComproller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetPriceOracle(&_VenusComproller.TransactOpts, newOracle)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_VenusComproller *VenusComprollerTransactor) SetProtocolPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setProtocolPaused", state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_VenusComproller *VenusComprollerSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetProtocolPaused(&_VenusComproller.TransactOpts, state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_VenusComproller *VenusComprollerTransactorSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetProtocolPaused(&_VenusComproller.TransactOpts, state)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetTreasuryData(&_VenusComproller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetTreasuryData(&_VenusComproller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetVAIController(opts *bind.TransactOpts, vaiController_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setVAIController", vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVAIController(&_VenusComproller.TransactOpts, vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVAIController(&_VenusComproller.TransactOpts, vaiController_)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetVAIMintRate(opts *bind.TransactOpts, newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setVAIMintRate", newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVAIMintRate(&_VenusComproller.TransactOpts, newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVAIMintRate(&_VenusComproller.TransactOpts, newVAIMintRate)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_VenusComproller *VenusComprollerTransactor) SetVAIVaultInfo(opts *bind.TransactOpts, vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setVAIVaultInfo", vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_VenusComproller *VenusComprollerSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVAIVaultInfo(&_VenusComproller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVAIVaultInfo(&_VenusComproller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_VenusComproller *VenusComprollerTransactor) SetVenusSpeed(opts *bind.TransactOpts, vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setVenusSpeed", vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_VenusComproller *VenusComprollerSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVenusSpeed(&_VenusComproller.TransactOpts, vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVenusSpeed(&_VenusComproller.TransactOpts, vToken, venusSpeed)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_VenusComproller *VenusComprollerTransactor) SetVenusVAIVaultRate(opts *bind.TransactOpts, venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_setVenusVAIVaultRate", venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_VenusComproller *VenusComprollerSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVenusVAIVaultRate(&_VenusComproller.TransactOpts, venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetVenusVAIVaultRate(&_VenusComproller.TransactOpts, venusVAIVaultRate_)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SupportMarket(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "_supportMarket", vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SupportMarket(&_VenusComproller.TransactOpts, vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.SupportMarket(&_VenusComproller.TransactOpts, vToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) BorrowAllowed(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "borrowAllowed", vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_VenusComproller *VenusComprollerSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.BorrowAllowed(&_VenusComproller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.BorrowAllowed(&_VenusComproller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_VenusComproller *VenusComprollerTransactor) BorrowVerify(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "borrowVerify", vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_VenusComproller *VenusComprollerSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.BorrowVerify(&_VenusComproller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_VenusComproller *VenusComprollerTransactorSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.BorrowVerify(&_VenusComproller.TransactOpts, vToken, borrower, borrowAmount)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x70bf66f0.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers, bool collateral) returns()
func (_VenusComproller *VenusComprollerTransactor) ClaimVenus(opts *bind.TransactOpts, holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool, collateral bool) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "claimVenus", holders, vTokens, borrowers, suppliers, collateral)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x70bf66f0.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers, bool collateral) returns()
func (_VenusComproller *VenusComprollerSession) ClaimVenus(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool, collateral bool) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus(&_VenusComproller.TransactOpts, holders, vTokens, borrowers, suppliers, collateral)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x70bf66f0.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers, bool collateral) returns()
func (_VenusComproller *VenusComprollerTransactorSession) ClaimVenus(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool, collateral bool) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus(&_VenusComproller.TransactOpts, holders, vTokens, borrowers, suppliers, collateral)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_VenusComproller *VenusComprollerTransactor) ClaimVenus0(opts *bind.TransactOpts, holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "claimVenus0", holder, vTokens)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_VenusComproller *VenusComprollerSession) ClaimVenus0(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus0(&_VenusComproller.TransactOpts, holder, vTokens)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_VenusComproller *VenusComprollerTransactorSession) ClaimVenus0(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus0(&_VenusComproller.TransactOpts, holder, vTokens)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_VenusComproller *VenusComprollerTransactor) ClaimVenus1(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "claimVenus1", holder)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_VenusComproller *VenusComprollerSession) ClaimVenus1(holder common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus1(&_VenusComproller.TransactOpts, holder)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_VenusComproller *VenusComprollerTransactorSession) ClaimVenus1(holder common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus1(&_VenusComproller.TransactOpts, holder)
}

// ClaimVenus2 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_VenusComproller *VenusComprollerTransactor) ClaimVenus2(opts *bind.TransactOpts, holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "claimVenus2", holders, vTokens, borrowers, suppliers)
}

// ClaimVenus2 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_VenusComproller *VenusComprollerSession) ClaimVenus2(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus2(&_VenusComproller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// ClaimVenus2 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_VenusComproller *VenusComprollerTransactorSession) ClaimVenus2(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenus2(&_VenusComproller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// ClaimVenusAsCollateral is a paid mutator transaction binding the contract method 0x7858524d.
//
// Solidity: function claimVenusAsCollateral(address holder) returns()
func (_VenusComproller *VenusComprollerTransactor) ClaimVenusAsCollateral(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "claimVenusAsCollateral", holder)
}

// ClaimVenusAsCollateral is a paid mutator transaction binding the contract method 0x7858524d.
//
// Solidity: function claimVenusAsCollateral(address holder) returns()
func (_VenusComproller *VenusComprollerSession) ClaimVenusAsCollateral(holder common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenusAsCollateral(&_VenusComproller.TransactOpts, holder)
}

// ClaimVenusAsCollateral is a paid mutator transaction binding the contract method 0x7858524d.
//
// Solidity: function claimVenusAsCollateral(address holder) returns()
func (_VenusComproller *VenusComprollerTransactorSession) ClaimVenusAsCollateral(holder common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ClaimVenusAsCollateral(&_VenusComproller.TransactOpts, holder)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_VenusComproller *VenusComprollerTransactor) EnterMarkets(opts *bind.TransactOpts, vTokens []common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "enterMarkets", vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_VenusComproller *VenusComprollerSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.EnterMarkets(&_VenusComproller.TransactOpts, vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_VenusComproller *VenusComprollerTransactorSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.EnterMarkets(&_VenusComproller.TransactOpts, vTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) ExitMarket(opts *bind.TransactOpts, vTokenAddress common.Address) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "exitMarket", vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_VenusComproller *VenusComprollerSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ExitMarket(&_VenusComproller.TransactOpts, vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _VenusComproller.Contract.ExitMarket(&_VenusComproller.TransactOpts, vTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "liquidateBorrowAllowed", vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusComproller *VenusComprollerSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.LiquidateBorrowAllowed(&_VenusComproller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.LiquidateBorrowAllowed(&_VenusComproller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_VenusComproller *VenusComprollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "liquidateBorrowVerify", vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_VenusComproller *VenusComprollerSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.LiquidateBorrowVerify(&_VenusComproller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_VenusComproller *VenusComprollerTransactorSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.LiquidateBorrowVerify(&_VenusComproller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) MintAllowed(opts *bind.TransactOpts, vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "mintAllowed", vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_VenusComproller *VenusComprollerSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.MintAllowed(&_VenusComproller.TransactOpts, vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.MintAllowed(&_VenusComproller.TransactOpts, vToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_VenusComproller *VenusComprollerTransactor) MintVerify(opts *bind.TransactOpts, vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "mintVerify", vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_VenusComproller *VenusComprollerSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.MintVerify(&_VenusComproller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_VenusComproller *VenusComprollerTransactorSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.MintVerify(&_VenusComproller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) RedeemAllowed(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "redeemAllowed", vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_VenusComproller *VenusComprollerSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RedeemAllowed(&_VenusComproller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RedeemAllowed(&_VenusComproller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_VenusComproller *VenusComprollerTransactor) RedeemVerify(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "redeemVerify", vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_VenusComproller *VenusComprollerSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RedeemVerify(&_VenusComproller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_VenusComproller *VenusComprollerTransactorSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RedeemVerify(&_VenusComproller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_VenusComproller *VenusComprollerTransactor) ReleaseToVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "releaseToVault")
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_VenusComproller *VenusComprollerSession) ReleaseToVault() (*types.Transaction, error) {
	return _VenusComproller.Contract.ReleaseToVault(&_VenusComproller.TransactOpts)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_VenusComproller *VenusComprollerTransactorSession) ReleaseToVault() (*types.Transaction, error) {
	return _VenusComproller.Contract.ReleaseToVault(&_VenusComproller.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "repayBorrowAllowed", vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusComproller *VenusComprollerSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RepayBorrowAllowed(&_VenusComproller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RepayBorrowAllowed(&_VenusComproller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_VenusComproller *VenusComprollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "repayBorrowVerify", vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_VenusComproller *VenusComprollerSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RepayBorrowVerify(&_VenusComproller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_VenusComproller *VenusComprollerTransactorSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.RepayBorrowVerify(&_VenusComproller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SeizeAllowed(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "seizeAllowed", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SeizeAllowed(&_VenusComproller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SeizeAllowed(&_VenusComproller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_VenusComproller *VenusComprollerTransactor) SeizeVerify(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "seizeVerify", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_VenusComproller *VenusComprollerSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SeizeVerify(&_VenusComproller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_VenusComproller *VenusComprollerTransactorSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SeizeVerify(&_VenusComproller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) SetMintedVAIOf(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "setMintedVAIOf", owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_VenusComproller *VenusComprollerSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetMintedVAIOf(&_VenusComproller.TransactOpts, owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.SetMintedVAIOf(&_VenusComproller.TransactOpts, owner, amount)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_VenusComproller *VenusComprollerTransactor) TransferAllowed(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "transferAllowed", vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_VenusComproller *VenusComprollerSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.TransferAllowed(&_VenusComproller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_VenusComproller *VenusComprollerTransactorSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.TransferAllowed(&_VenusComproller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_VenusComproller *VenusComprollerTransactor) TransferVerify(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.contract.Transact(opts, "transferVerify", vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_VenusComproller *VenusComprollerSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.TransferVerify(&_VenusComproller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_VenusComproller *VenusComprollerTransactorSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _VenusComproller.Contract.TransferVerify(&_VenusComproller.TransactOpts, vToken, src, dst, transferTokens)
}

// VenusComprollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the VenusComproller contract.
type VenusComprollerActionPausedIterator struct {
	Event *VenusComprollerActionPaused // Event containing the contract specifics and raw log

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
func (it *VenusComprollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerActionPaused)
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
		it.Event = new(VenusComprollerActionPaused)
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
func (it *VenusComprollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerActionPaused represents a ActionPaused event raised by the VenusComproller contract.
type VenusComprollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_VenusComproller *VenusComprollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*VenusComprollerActionPausedIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerActionPausedIterator{contract: _VenusComproller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_VenusComproller *VenusComprollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *VenusComprollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerActionPaused)
				if err := _VenusComproller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
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

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_VenusComproller *VenusComprollerFilterer) ParseActionPaused(log types.Log) (*VenusComprollerActionPaused, error) {
	event := new(VenusComprollerActionPaused)
	if err := _VenusComproller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerActionPausedMarketIterator is returned from FilterActionPausedMarket and is used to iterate over the raw logs and unpacked data for ActionPausedMarket events raised by the VenusComproller contract.
type VenusComprollerActionPausedMarketIterator struct {
	Event *VenusComprollerActionPausedMarket // Event containing the contract specifics and raw log

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
func (it *VenusComprollerActionPausedMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerActionPausedMarket)
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
		it.Event = new(VenusComprollerActionPausedMarket)
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
func (it *VenusComprollerActionPausedMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerActionPausedMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerActionPausedMarket represents a ActionPausedMarket event raised by the VenusComproller contract.
type VenusComprollerActionPausedMarket struct {
	VToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPausedMarket is a free log retrieval operation binding the contract event 0xae9bd19d6b7dcb722457462ed82420131e5740c4283e4f5b0c8332e96becd556.
//
// Solidity: event ActionPausedMarket(address vToken, string action, bool pauseState)
func (_VenusComproller *VenusComprollerFilterer) FilterActionPausedMarket(opts *bind.FilterOpts) (*VenusComprollerActionPausedMarketIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "ActionPausedMarket")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerActionPausedMarketIterator{contract: _VenusComproller.contract, event: "ActionPausedMarket", logs: logs, sub: sub}, nil
}

// WatchActionPausedMarket is a free log subscription operation binding the contract event 0xae9bd19d6b7dcb722457462ed82420131e5740c4283e4f5b0c8332e96becd556.
//
// Solidity: event ActionPausedMarket(address vToken, string action, bool pauseState)
func (_VenusComproller *VenusComprollerFilterer) WatchActionPausedMarket(opts *bind.WatchOpts, sink chan<- *VenusComprollerActionPausedMarket) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "ActionPausedMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerActionPausedMarket)
				if err := _VenusComproller.contract.UnpackLog(event, "ActionPausedMarket", log); err != nil {
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

// ParseActionPausedMarket is a log parse operation binding the contract event 0xae9bd19d6b7dcb722457462ed82420131e5740c4283e4f5b0c8332e96becd556.
//
// Solidity: event ActionPausedMarket(address vToken, string action, bool pauseState)
func (_VenusComproller *VenusComprollerFilterer) ParseActionPausedMarket(log types.Log) (*VenusComprollerActionPausedMarket, error) {
	event := new(VenusComprollerActionPausedMarket)
	if err := _VenusComproller.contract.UnpackLog(event, "ActionPausedMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerActionProtocolPausedIterator is returned from FilterActionProtocolPaused and is used to iterate over the raw logs and unpacked data for ActionProtocolPaused events raised by the VenusComproller contract.
type VenusComprollerActionProtocolPausedIterator struct {
	Event *VenusComprollerActionProtocolPaused // Event containing the contract specifics and raw log

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
func (it *VenusComprollerActionProtocolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerActionProtocolPaused)
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
		it.Event = new(VenusComprollerActionProtocolPaused)
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
func (it *VenusComprollerActionProtocolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerActionProtocolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerActionProtocolPaused represents a ActionProtocolPaused event raised by the VenusComproller contract.
type VenusComprollerActionProtocolPaused struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterActionProtocolPaused is a free log retrieval operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_VenusComproller *VenusComprollerFilterer) FilterActionProtocolPaused(opts *bind.FilterOpts) (*VenusComprollerActionProtocolPausedIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerActionProtocolPausedIterator{contract: _VenusComproller.contract, event: "ActionProtocolPaused", logs: logs, sub: sub}, nil
}

// WatchActionProtocolPaused is a free log subscription operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_VenusComproller *VenusComprollerFilterer) WatchActionProtocolPaused(opts *bind.WatchOpts, sink chan<- *VenusComprollerActionProtocolPaused) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerActionProtocolPaused)
				if err := _VenusComproller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
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

// ParseActionProtocolPaused is a log parse operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_VenusComproller *VenusComprollerFilterer) ParseActionProtocolPaused(log types.Log) (*VenusComprollerActionProtocolPaused, error) {
	event := new(VenusComprollerActionProtocolPaused)
	if err := _VenusComproller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerDistributedBorrowerVenusIterator is returned from FilterDistributedBorrowerVenus and is used to iterate over the raw logs and unpacked data for DistributedBorrowerVenus events raised by the VenusComproller contract.
type VenusComprollerDistributedBorrowerVenusIterator struct {
	Event *VenusComprollerDistributedBorrowerVenus // Event containing the contract specifics and raw log

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
func (it *VenusComprollerDistributedBorrowerVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerDistributedBorrowerVenus)
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
		it.Event = new(VenusComprollerDistributedBorrowerVenus)
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
func (it *VenusComprollerDistributedBorrowerVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerDistributedBorrowerVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerDistributedBorrowerVenus represents a DistributedBorrowerVenus event raised by the VenusComproller contract.
type VenusComprollerDistributedBorrowerVenus struct {
	VToken           common.Address
	Borrower         common.Address
	VenusDelta       *big.Int
	VenusBorrowIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerVenus is a free log retrieval operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_VenusComproller *VenusComprollerFilterer) FilterDistributedBorrowerVenus(opts *bind.FilterOpts, vToken []common.Address, borrower []common.Address) (*VenusComprollerDistributedBorrowerVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerDistributedBorrowerVenusIterator{contract: _VenusComproller.contract, event: "DistributedBorrowerVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerVenus is a free log subscription operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_VenusComproller *VenusComprollerFilterer) WatchDistributedBorrowerVenus(opts *bind.WatchOpts, sink chan<- *VenusComprollerDistributedBorrowerVenus, vToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerDistributedBorrowerVenus)
				if err := _VenusComproller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
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

// ParseDistributedBorrowerVenus is a log parse operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_VenusComproller *VenusComprollerFilterer) ParseDistributedBorrowerVenus(log types.Log) (*VenusComprollerDistributedBorrowerVenus, error) {
	event := new(VenusComprollerDistributedBorrowerVenus)
	if err := _VenusComproller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerDistributedSupplierVenusIterator is returned from FilterDistributedSupplierVenus and is used to iterate over the raw logs and unpacked data for DistributedSupplierVenus events raised by the VenusComproller contract.
type VenusComprollerDistributedSupplierVenusIterator struct {
	Event *VenusComprollerDistributedSupplierVenus // Event containing the contract specifics and raw log

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
func (it *VenusComprollerDistributedSupplierVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerDistributedSupplierVenus)
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
		it.Event = new(VenusComprollerDistributedSupplierVenus)
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
func (it *VenusComprollerDistributedSupplierVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerDistributedSupplierVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerDistributedSupplierVenus represents a DistributedSupplierVenus event raised by the VenusComproller contract.
type VenusComprollerDistributedSupplierVenus struct {
	VToken           common.Address
	Supplier         common.Address
	VenusDelta       *big.Int
	VenusSupplyIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierVenus is a free log retrieval operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_VenusComproller *VenusComprollerFilterer) FilterDistributedSupplierVenus(opts *bind.FilterOpts, vToken []common.Address, supplier []common.Address) (*VenusComprollerDistributedSupplierVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerDistributedSupplierVenusIterator{contract: _VenusComproller.contract, event: "DistributedSupplierVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierVenus is a free log subscription operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_VenusComproller *VenusComprollerFilterer) WatchDistributedSupplierVenus(opts *bind.WatchOpts, sink chan<- *VenusComprollerDistributedSupplierVenus, vToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerDistributedSupplierVenus)
				if err := _VenusComproller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
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

// ParseDistributedSupplierVenus is a log parse operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_VenusComproller *VenusComprollerFilterer) ParseDistributedSupplierVenus(log types.Log) (*VenusComprollerDistributedSupplierVenus, error) {
	event := new(VenusComprollerDistributedSupplierVenus)
	if err := _VenusComproller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerDistributedVAIMinterVenusIterator is returned from FilterDistributedVAIMinterVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIMinterVenus events raised by the VenusComproller contract.
type VenusComprollerDistributedVAIMinterVenusIterator struct {
	Event *VenusComprollerDistributedVAIMinterVenus // Event containing the contract specifics and raw log

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
func (it *VenusComprollerDistributedVAIMinterVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerDistributedVAIMinterVenus)
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
		it.Event = new(VenusComprollerDistributedVAIMinterVenus)
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
func (it *VenusComprollerDistributedVAIMinterVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerDistributedVAIMinterVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerDistributedVAIMinterVenus represents a DistributedVAIMinterVenus event raised by the VenusComproller contract.
type VenusComprollerDistributedVAIMinterVenus struct {
	VaiMinter         common.Address
	VenusDelta        *big.Int
	VenusVAIMintIndex *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIMinterVenus is a free log retrieval operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_VenusComproller *VenusComprollerFilterer) FilterDistributedVAIMinterVenus(opts *bind.FilterOpts, vaiMinter []common.Address) (*VenusComprollerDistributedVAIMinterVenusIterator, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerDistributedVAIMinterVenusIterator{contract: _VenusComproller.contract, event: "DistributedVAIMinterVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIMinterVenus is a free log subscription operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_VenusComproller *VenusComprollerFilterer) WatchDistributedVAIMinterVenus(opts *bind.WatchOpts, sink chan<- *VenusComprollerDistributedVAIMinterVenus, vaiMinter []common.Address) (event.Subscription, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerDistributedVAIMinterVenus)
				if err := _VenusComproller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
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

// ParseDistributedVAIMinterVenus is a log parse operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_VenusComproller *VenusComprollerFilterer) ParseDistributedVAIMinterVenus(log types.Log) (*VenusComprollerDistributedVAIMinterVenus, error) {
	event := new(VenusComprollerDistributedVAIMinterVenus)
	if err := _VenusComproller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerDistributedVAIVaultVenusIterator is returned from FilterDistributedVAIVaultVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIVaultVenus events raised by the VenusComproller contract.
type VenusComprollerDistributedVAIVaultVenusIterator struct {
	Event *VenusComprollerDistributedVAIVaultVenus // Event containing the contract specifics and raw log

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
func (it *VenusComprollerDistributedVAIVaultVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerDistributedVAIVaultVenus)
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
		it.Event = new(VenusComprollerDistributedVAIVaultVenus)
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
func (it *VenusComprollerDistributedVAIVaultVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerDistributedVAIVaultVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerDistributedVAIVaultVenus represents a DistributedVAIVaultVenus event raised by the VenusComproller contract.
type VenusComprollerDistributedVAIVaultVenus struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIVaultVenus is a free log retrieval operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_VenusComproller *VenusComprollerFilterer) FilterDistributedVAIVaultVenus(opts *bind.FilterOpts) (*VenusComprollerDistributedVAIVaultVenusIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerDistributedVAIVaultVenusIterator{contract: _VenusComproller.contract, event: "DistributedVAIVaultVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIVaultVenus is a free log subscription operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_VenusComproller *VenusComprollerFilterer) WatchDistributedVAIVaultVenus(opts *bind.WatchOpts, sink chan<- *VenusComprollerDistributedVAIVaultVenus) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerDistributedVAIVaultVenus)
				if err := _VenusComproller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
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

// ParseDistributedVAIVaultVenus is a log parse operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_VenusComproller *VenusComprollerFilterer) ParseDistributedVAIVaultVenus(log types.Log) (*VenusComprollerDistributedVAIVaultVenus, error) {
	event := new(VenusComprollerDistributedVAIVaultVenus)
	if err := _VenusComproller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the VenusComproller contract.
type VenusComprollerFailureIterator struct {
	Event *VenusComprollerFailure // Event containing the contract specifics and raw log

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
func (it *VenusComprollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerFailure)
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
		it.Event = new(VenusComprollerFailure)
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
func (it *VenusComprollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerFailure represents a Failure event raised by the VenusComproller contract.
type VenusComprollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VenusComproller *VenusComprollerFilterer) FilterFailure(opts *bind.FilterOpts) (*VenusComprollerFailureIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerFailureIterator{contract: _VenusComproller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VenusComproller *VenusComprollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *VenusComprollerFailure) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerFailure)
				if err := _VenusComproller.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_VenusComproller *VenusComprollerFilterer) ParseFailure(log types.Log) (*VenusComprollerFailure, error) {
	event := new(VenusComprollerFailure)
	if err := _VenusComproller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the VenusComproller contract.
type VenusComprollerMarketEnteredIterator struct {
	Event *VenusComprollerMarketEntered // Event containing the contract specifics and raw log

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
func (it *VenusComprollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerMarketEntered)
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
		it.Event = new(VenusComprollerMarketEntered)
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
func (it *VenusComprollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerMarketEntered represents a MarketEntered event raised by the VenusComproller contract.
type VenusComprollerMarketEntered struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_VenusComproller *VenusComprollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*VenusComprollerMarketEnteredIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerMarketEnteredIterator{contract: _VenusComproller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_VenusComproller *VenusComprollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *VenusComprollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerMarketEntered)
				if err := _VenusComproller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_VenusComproller *VenusComprollerFilterer) ParseMarketEntered(log types.Log) (*VenusComprollerMarketEntered, error) {
	event := new(VenusComprollerMarketEntered)
	if err := _VenusComproller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the VenusComproller contract.
type VenusComprollerMarketExitedIterator struct {
	Event *VenusComprollerMarketExited // Event containing the contract specifics and raw log

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
func (it *VenusComprollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerMarketExited)
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
		it.Event = new(VenusComprollerMarketExited)
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
func (it *VenusComprollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerMarketExited represents a MarketExited event raised by the VenusComproller contract.
type VenusComprollerMarketExited struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_VenusComproller *VenusComprollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*VenusComprollerMarketExitedIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerMarketExitedIterator{contract: _VenusComproller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_VenusComproller *VenusComprollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *VenusComprollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerMarketExited)
				if err := _VenusComproller.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_VenusComproller *VenusComprollerFilterer) ParseMarketExited(log types.Log) (*VenusComprollerMarketExited, error) {
	event := new(VenusComprollerMarketExited)
	if err := _VenusComproller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the VenusComproller contract.
type VenusComprollerMarketListedIterator struct {
	Event *VenusComprollerMarketListed // Event containing the contract specifics and raw log

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
func (it *VenusComprollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerMarketListed)
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
		it.Event = new(VenusComprollerMarketListed)
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
func (it *VenusComprollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerMarketListed represents a MarketListed event raised by the VenusComproller contract.
type VenusComprollerMarketListed struct {
	VToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_VenusComproller *VenusComprollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*VenusComprollerMarketListedIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerMarketListedIterator{contract: _VenusComproller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_VenusComproller *VenusComprollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *VenusComprollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerMarketListed)
				if err := _VenusComproller.contract.UnpackLog(event, "MarketListed", log); err != nil {
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

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_VenusComproller *VenusComprollerFilterer) ParseMarketListed(log types.Log) (*VenusComprollerMarketListed, error) {
	event := new(VenusComprollerMarketListed)
	if err := _VenusComproller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the VenusComproller contract.
type VenusComprollerNewBorrowCapIterator struct {
	Event *VenusComprollerNewBorrowCap // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewBorrowCap)
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
		it.Event = new(VenusComprollerNewBorrowCap)
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
func (it *VenusComprollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewBorrowCap represents a NewBorrowCap event raised by the VenusComproller contract.
type VenusComprollerNewBorrowCap struct {
	VToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_VenusComproller *VenusComprollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, vToken []common.Address) (*VenusComprollerNewBorrowCapIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewBorrowCapIterator{contract: _VenusComproller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_VenusComproller *VenusComprollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewBorrowCap, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewBorrowCap)
				if err := _VenusComproller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
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

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_VenusComproller *VenusComprollerFilterer) ParseNewBorrowCap(log types.Log) (*VenusComprollerNewBorrowCap, error) {
	event := new(VenusComprollerNewBorrowCap)
	if err := _VenusComproller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the VenusComproller contract.
type VenusComprollerNewBorrowCapGuardianIterator struct {
	Event *VenusComprollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewBorrowCapGuardian)
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
		it.Event = new(VenusComprollerNewBorrowCapGuardian)
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
func (it *VenusComprollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the VenusComproller contract.
type VenusComprollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_VenusComproller *VenusComprollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*VenusComprollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewBorrowCapGuardianIterator{contract: _VenusComproller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_VenusComproller *VenusComprollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewBorrowCapGuardian)
				if err := _VenusComproller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
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

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_VenusComproller *VenusComprollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*VenusComprollerNewBorrowCapGuardian, error) {
	event := new(VenusComprollerNewBorrowCapGuardian)
	if err := _VenusComproller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the VenusComproller contract.
type VenusComprollerNewCloseFactorIterator struct {
	Event *VenusComprollerNewCloseFactor // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewCloseFactor)
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
		it.Event = new(VenusComprollerNewCloseFactor)
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
func (it *VenusComprollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewCloseFactor represents a NewCloseFactor event raised by the VenusComproller contract.
type VenusComprollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_VenusComproller *VenusComprollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*VenusComprollerNewCloseFactorIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewCloseFactorIterator{contract: _VenusComproller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_VenusComproller *VenusComprollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewCloseFactor)
				if err := _VenusComproller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_VenusComproller *VenusComprollerFilterer) ParseNewCloseFactor(log types.Log) (*VenusComprollerNewCloseFactor, error) {
	event := new(VenusComprollerNewCloseFactor)
	if err := _VenusComproller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the VenusComproller contract.
type VenusComprollerNewCollateralFactorIterator struct {
	Event *VenusComprollerNewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewCollateralFactor)
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
		it.Event = new(VenusComprollerNewCollateralFactor)
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
func (it *VenusComprollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewCollateralFactor represents a NewCollateralFactor event raised by the VenusComproller contract.
type VenusComprollerNewCollateralFactor struct {
	VToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_VenusComproller *VenusComprollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*VenusComprollerNewCollateralFactorIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewCollateralFactorIterator{contract: _VenusComproller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_VenusComproller *VenusComprollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewCollateralFactor)
				if err := _VenusComproller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_VenusComproller *VenusComprollerFilterer) ParseNewCollateralFactor(log types.Log) (*VenusComprollerNewCollateralFactor, error) {
	event := new(VenusComprollerNewCollateralFactor)
	if err := _VenusComproller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewComptrollerLensIterator is returned from FilterNewComptrollerLens and is used to iterate over the raw logs and unpacked data for NewComptrollerLens events raised by the VenusComproller contract.
type VenusComprollerNewComptrollerLensIterator struct {
	Event *VenusComprollerNewComptrollerLens // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewComptrollerLensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewComptrollerLens)
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
		it.Event = new(VenusComprollerNewComptrollerLens)
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
func (it *VenusComprollerNewComptrollerLensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewComptrollerLensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewComptrollerLens represents a NewComptrollerLens event raised by the VenusComproller contract.
type VenusComprollerNewComptrollerLens struct {
	OldComptrollerLens common.Address
	NewComptrollerLens common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewComptrollerLens is a free log retrieval operation binding the contract event 0x0f7eb572d1b3053a0a3a33c04151364cf88d163182a5e4e1088cb8e52321e08a.
//
// Solidity: event NewComptrollerLens(address oldComptrollerLens, address newComptrollerLens)
func (_VenusComproller *VenusComprollerFilterer) FilterNewComptrollerLens(opts *bind.FilterOpts) (*VenusComprollerNewComptrollerLensIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewComptrollerLens")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewComptrollerLensIterator{contract: _VenusComproller.contract, event: "NewComptrollerLens", logs: logs, sub: sub}, nil
}

// WatchNewComptrollerLens is a free log subscription operation binding the contract event 0x0f7eb572d1b3053a0a3a33c04151364cf88d163182a5e4e1088cb8e52321e08a.
//
// Solidity: event NewComptrollerLens(address oldComptrollerLens, address newComptrollerLens)
func (_VenusComproller *VenusComprollerFilterer) WatchNewComptrollerLens(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewComptrollerLens) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewComptrollerLens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewComptrollerLens)
				if err := _VenusComproller.contract.UnpackLog(event, "NewComptrollerLens", log); err != nil {
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

// ParseNewComptrollerLens is a log parse operation binding the contract event 0x0f7eb572d1b3053a0a3a33c04151364cf88d163182a5e4e1088cb8e52321e08a.
//
// Solidity: event NewComptrollerLens(address oldComptrollerLens, address newComptrollerLens)
func (_VenusComproller *VenusComprollerFilterer) ParseNewComptrollerLens(log types.Log) (*VenusComprollerNewComptrollerLens, error) {
	event := new(VenusComprollerNewComptrollerLens)
	if err := _VenusComproller.contract.UnpackLog(event, "NewComptrollerLens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the VenusComproller contract.
type VenusComprollerNewLiquidationIncentiveIterator struct {
	Event *VenusComprollerNewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewLiquidationIncentive)
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
		it.Event = new(VenusComprollerNewLiquidationIncentive)
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
func (it *VenusComprollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the VenusComproller contract.
type VenusComprollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_VenusComproller *VenusComprollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*VenusComprollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewLiquidationIncentiveIterator{contract: _VenusComproller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_VenusComproller *VenusComprollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewLiquidationIncentive)
				if err := _VenusComproller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_VenusComproller *VenusComprollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*VenusComprollerNewLiquidationIncentive, error) {
	event := new(VenusComprollerNewLiquidationIncentive)
	if err := _VenusComproller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewLiquidatorContractIterator is returned from FilterNewLiquidatorContract and is used to iterate over the raw logs and unpacked data for NewLiquidatorContract events raised by the VenusComproller contract.
type VenusComprollerNewLiquidatorContractIterator struct {
	Event *VenusComprollerNewLiquidatorContract // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewLiquidatorContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewLiquidatorContract)
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
		it.Event = new(VenusComprollerNewLiquidatorContract)
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
func (it *VenusComprollerNewLiquidatorContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewLiquidatorContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewLiquidatorContract represents a NewLiquidatorContract event raised by the VenusComproller contract.
type VenusComprollerNewLiquidatorContract struct {
	OldLiquidatorContract common.Address
	NewLiquidatorContract common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidatorContract is a free log retrieval operation binding the contract event 0xa5ea5616d5f6dbbaa62fdfcf3856723216eed485c394d9e51ec8e6d40e1d1ac1.
//
// Solidity: event NewLiquidatorContract(address oldLiquidatorContract, address newLiquidatorContract)
func (_VenusComproller *VenusComprollerFilterer) FilterNewLiquidatorContract(opts *bind.FilterOpts) (*VenusComprollerNewLiquidatorContractIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewLiquidatorContract")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewLiquidatorContractIterator{contract: _VenusComproller.contract, event: "NewLiquidatorContract", logs: logs, sub: sub}, nil
}

// WatchNewLiquidatorContract is a free log subscription operation binding the contract event 0xa5ea5616d5f6dbbaa62fdfcf3856723216eed485c394d9e51ec8e6d40e1d1ac1.
//
// Solidity: event NewLiquidatorContract(address oldLiquidatorContract, address newLiquidatorContract)
func (_VenusComproller *VenusComprollerFilterer) WatchNewLiquidatorContract(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewLiquidatorContract) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewLiquidatorContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewLiquidatorContract)
				if err := _VenusComproller.contract.UnpackLog(event, "NewLiquidatorContract", log); err != nil {
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

// ParseNewLiquidatorContract is a log parse operation binding the contract event 0xa5ea5616d5f6dbbaa62fdfcf3856723216eed485c394d9e51ec8e6d40e1d1ac1.
//
// Solidity: event NewLiquidatorContract(address oldLiquidatorContract, address newLiquidatorContract)
func (_VenusComproller *VenusComprollerFilterer) ParseNewLiquidatorContract(log types.Log) (*VenusComprollerNewLiquidatorContract, error) {
	event := new(VenusComprollerNewLiquidatorContract)
	if err := _VenusComproller.contract.UnpackLog(event, "NewLiquidatorContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the VenusComproller contract.
type VenusComprollerNewPauseGuardianIterator struct {
	Event *VenusComprollerNewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewPauseGuardian)
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
		it.Event = new(VenusComprollerNewPauseGuardian)
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
func (it *VenusComprollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewPauseGuardian represents a NewPauseGuardian event raised by the VenusComproller contract.
type VenusComprollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_VenusComproller *VenusComprollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*VenusComprollerNewPauseGuardianIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewPauseGuardianIterator{contract: _VenusComproller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_VenusComproller *VenusComprollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewPauseGuardian)
				if err := _VenusComproller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_VenusComproller *VenusComprollerFilterer) ParseNewPauseGuardian(log types.Log) (*VenusComprollerNewPauseGuardian, error) {
	event := new(VenusComprollerNewPauseGuardian)
	if err := _VenusComproller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the VenusComproller contract.
type VenusComprollerNewPriceOracleIterator struct {
	Event *VenusComprollerNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewPriceOracle)
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
		it.Event = new(VenusComprollerNewPriceOracle)
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
func (it *VenusComprollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewPriceOracle represents a NewPriceOracle event raised by the VenusComproller contract.
type VenusComprollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_VenusComproller *VenusComprollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*VenusComprollerNewPriceOracleIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewPriceOracleIterator{contract: _VenusComproller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_VenusComproller *VenusComprollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewPriceOracle)
				if err := _VenusComproller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_VenusComproller *VenusComprollerFilterer) ParseNewPriceOracle(log types.Log) (*VenusComprollerNewPriceOracle, error) {
	event := new(VenusComprollerNewPriceOracle)
	if err := _VenusComproller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the VenusComproller contract.
type VenusComprollerNewTreasuryAddressIterator struct {
	Event *VenusComprollerNewTreasuryAddress // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewTreasuryAddress)
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
		it.Event = new(VenusComprollerNewTreasuryAddress)
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
func (it *VenusComprollerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewTreasuryAddress represents a NewTreasuryAddress event raised by the VenusComproller contract.
type VenusComprollerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VenusComproller *VenusComprollerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*VenusComprollerNewTreasuryAddressIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewTreasuryAddressIterator{contract: _VenusComproller.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VenusComproller *VenusComprollerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewTreasuryAddress)
				if err := _VenusComproller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
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

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_VenusComproller *VenusComprollerFilterer) ParseNewTreasuryAddress(log types.Log) (*VenusComprollerNewTreasuryAddress, error) {
	event := new(VenusComprollerNewTreasuryAddress)
	if err := _VenusComproller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the VenusComproller contract.
type VenusComprollerNewTreasuryGuardianIterator struct {
	Event *VenusComprollerNewTreasuryGuardian // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewTreasuryGuardian)
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
		it.Event = new(VenusComprollerNewTreasuryGuardian)
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
func (it *VenusComprollerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the VenusComproller contract.
type VenusComprollerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VenusComproller *VenusComprollerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*VenusComprollerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewTreasuryGuardianIterator{contract: _VenusComproller.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VenusComproller *VenusComprollerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewTreasuryGuardian)
				if err := _VenusComproller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
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

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_VenusComproller *VenusComprollerFilterer) ParseNewTreasuryGuardian(log types.Log) (*VenusComprollerNewTreasuryGuardian, error) {
	event := new(VenusComprollerNewTreasuryGuardian)
	if err := _VenusComproller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the VenusComproller contract.
type VenusComprollerNewTreasuryPercentIterator struct {
	Event *VenusComprollerNewTreasuryPercent // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewTreasuryPercent)
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
		it.Event = new(VenusComprollerNewTreasuryPercent)
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
func (it *VenusComprollerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewTreasuryPercent represents a NewTreasuryPercent event raised by the VenusComproller contract.
type VenusComprollerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VenusComproller *VenusComprollerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*VenusComprollerNewTreasuryPercentIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewTreasuryPercentIterator{contract: _VenusComproller.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VenusComproller *VenusComprollerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewTreasuryPercent)
				if err := _VenusComproller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
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

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_VenusComproller *VenusComprollerFilterer) ParseNewTreasuryPercent(log types.Log) (*VenusComprollerNewTreasuryPercent, error) {
	event := new(VenusComprollerNewTreasuryPercent)
	if err := _VenusComproller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewVAIControllerIterator is returned from FilterNewVAIController and is used to iterate over the raw logs and unpacked data for NewVAIController events raised by the VenusComproller contract.
type VenusComprollerNewVAIControllerIterator struct {
	Event *VenusComprollerNewVAIController // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewVAIControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewVAIController)
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
		it.Event = new(VenusComprollerNewVAIController)
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
func (it *VenusComprollerNewVAIControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewVAIControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewVAIController represents a NewVAIController event raised by the VenusComproller contract.
type VenusComprollerNewVAIController struct {
	OldVAIController common.Address
	NewVAIController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewVAIController is a free log retrieval operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_VenusComproller *VenusComprollerFilterer) FilterNewVAIController(opts *bind.FilterOpts) (*VenusComprollerNewVAIControllerIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewVAIControllerIterator{contract: _VenusComproller.contract, event: "NewVAIController", logs: logs, sub: sub}, nil
}

// WatchNewVAIController is a free log subscription operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_VenusComproller *VenusComprollerFilterer) WatchNewVAIController(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewVAIController) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewVAIController)
				if err := _VenusComproller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
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

// ParseNewVAIController is a log parse operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_VenusComproller *VenusComprollerFilterer) ParseNewVAIController(log types.Log) (*VenusComprollerNewVAIController, error) {
	event := new(VenusComprollerNewVAIController)
	if err := _VenusComproller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewVAIMintRateIterator is returned from FilterNewVAIMintRate and is used to iterate over the raw logs and unpacked data for NewVAIMintRate events raised by the VenusComproller contract.
type VenusComprollerNewVAIMintRateIterator struct {
	Event *VenusComprollerNewVAIMintRate // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewVAIMintRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewVAIMintRate)
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
		it.Event = new(VenusComprollerNewVAIMintRate)
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
func (it *VenusComprollerNewVAIMintRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewVAIMintRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewVAIMintRate represents a NewVAIMintRate event raised by the VenusComproller contract.
type VenusComprollerNewVAIMintRate struct {
	OldVAIMintRate *big.Int
	NewVAIMintRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewVAIMintRate is a free log retrieval operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_VenusComproller *VenusComprollerFilterer) FilterNewVAIMintRate(opts *bind.FilterOpts) (*VenusComprollerNewVAIMintRateIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewVAIMintRateIterator{contract: _VenusComproller.contract, event: "NewVAIMintRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIMintRate is a free log subscription operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_VenusComproller *VenusComprollerFilterer) WatchNewVAIMintRate(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewVAIMintRate) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewVAIMintRate)
				if err := _VenusComproller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
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

// ParseNewVAIMintRate is a log parse operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_VenusComproller *VenusComprollerFilterer) ParseNewVAIMintRate(log types.Log) (*VenusComprollerNewVAIMintRate, error) {
	event := new(VenusComprollerNewVAIMintRate)
	if err := _VenusComproller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewVAIVaultInfoIterator is returned from FilterNewVAIVaultInfo and is used to iterate over the raw logs and unpacked data for NewVAIVaultInfo events raised by the VenusComproller contract.
type VenusComprollerNewVAIVaultInfoIterator struct {
	Event *VenusComprollerNewVAIVaultInfo // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewVAIVaultInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewVAIVaultInfo)
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
		it.Event = new(VenusComprollerNewVAIVaultInfo)
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
func (it *VenusComprollerNewVAIVaultInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewVAIVaultInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewVAIVaultInfo represents a NewVAIVaultInfo event raised by the VenusComproller contract.
type VenusComprollerNewVAIVaultInfo struct {
	Vault             common.Address
	ReleaseStartBlock *big.Int
	ReleaseInterval   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewVAIVaultInfo is a free log retrieval operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_VenusComproller *VenusComprollerFilterer) FilterNewVAIVaultInfo(opts *bind.FilterOpts) (*VenusComprollerNewVAIVaultInfoIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewVAIVaultInfoIterator{contract: _VenusComproller.contract, event: "NewVAIVaultInfo", logs: logs, sub: sub}, nil
}

// WatchNewVAIVaultInfo is a free log subscription operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_VenusComproller *VenusComprollerFilterer) WatchNewVAIVaultInfo(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewVAIVaultInfo) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewVAIVaultInfo)
				if err := _VenusComproller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
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

// ParseNewVAIVaultInfo is a log parse operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_VenusComproller *VenusComprollerFilterer) ParseNewVAIVaultInfo(log types.Log) (*VenusComprollerNewVAIVaultInfo, error) {
	event := new(VenusComprollerNewVAIVaultInfo)
	if err := _VenusComproller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewVenusVAIRateIterator is returned from FilterNewVenusVAIRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIRate events raised by the VenusComproller contract.
type VenusComprollerNewVenusVAIRateIterator struct {
	Event *VenusComprollerNewVenusVAIRate // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewVenusVAIRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewVenusVAIRate)
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
		it.Event = new(VenusComprollerNewVenusVAIRate)
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
func (it *VenusComprollerNewVenusVAIRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewVenusVAIRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewVenusVAIRate represents a NewVenusVAIRate event raised by the VenusComproller contract.
type VenusComprollerNewVenusVAIRate struct {
	OldVenusVAIRate *big.Int
	NewVenusVAIRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIRate is a free log retrieval operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_VenusComproller *VenusComprollerFilterer) FilterNewVenusVAIRate(opts *bind.FilterOpts) (*VenusComprollerNewVenusVAIRateIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewVenusVAIRateIterator{contract: _VenusComproller.contract, event: "NewVenusVAIRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIRate is a free log subscription operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_VenusComproller *VenusComprollerFilterer) WatchNewVenusVAIRate(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewVenusVAIRate) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewVenusVAIRate)
				if err := _VenusComproller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
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

// ParseNewVenusVAIRate is a log parse operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_VenusComproller *VenusComprollerFilterer) ParseNewVenusVAIRate(log types.Log) (*VenusComprollerNewVenusVAIRate, error) {
	event := new(VenusComprollerNewVenusVAIRate)
	if err := _VenusComproller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerNewVenusVAIVaultRateIterator is returned from FilterNewVenusVAIVaultRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIVaultRate events raised by the VenusComproller contract.
type VenusComprollerNewVenusVAIVaultRateIterator struct {
	Event *VenusComprollerNewVenusVAIVaultRate // Event containing the contract specifics and raw log

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
func (it *VenusComprollerNewVenusVAIVaultRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerNewVenusVAIVaultRate)
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
		it.Event = new(VenusComprollerNewVenusVAIVaultRate)
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
func (it *VenusComprollerNewVenusVAIVaultRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerNewVenusVAIVaultRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerNewVenusVAIVaultRate represents a NewVenusVAIVaultRate event raised by the VenusComproller contract.
type VenusComprollerNewVenusVAIVaultRate struct {
	OldVenusVAIVaultRate *big.Int
	NewVenusVAIVaultRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIVaultRate is a free log retrieval operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_VenusComproller *VenusComprollerFilterer) FilterNewVenusVAIVaultRate(opts *bind.FilterOpts) (*VenusComprollerNewVenusVAIVaultRateIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerNewVenusVAIVaultRateIterator{contract: _VenusComproller.contract, event: "NewVenusVAIVaultRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIVaultRate is a free log subscription operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_VenusComproller *VenusComprollerFilterer) WatchNewVenusVAIVaultRate(opts *bind.WatchOpts, sink chan<- *VenusComprollerNewVenusVAIVaultRate) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerNewVenusVAIVaultRate)
				if err := _VenusComproller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
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

// ParseNewVenusVAIVaultRate is a log parse operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_VenusComproller *VenusComprollerFilterer) ParseNewVenusVAIVaultRate(log types.Log) (*VenusComprollerNewVenusVAIVaultRate, error) {
	event := new(VenusComprollerNewVenusVAIVaultRate)
	if err := _VenusComproller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerVenusGrantedIterator is returned from FilterVenusGranted and is used to iterate over the raw logs and unpacked data for VenusGranted events raised by the VenusComproller contract.
type VenusComprollerVenusGrantedIterator struct {
	Event *VenusComprollerVenusGranted // Event containing the contract specifics and raw log

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
func (it *VenusComprollerVenusGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerVenusGranted)
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
		it.Event = new(VenusComprollerVenusGranted)
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
func (it *VenusComprollerVenusGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerVenusGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerVenusGranted represents a VenusGranted event raised by the VenusComproller contract.
type VenusComprollerVenusGranted struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVenusGranted is a free log retrieval operation binding the contract event 0xd7fe674cac9eee3998fe3cbd7a6f93c3bc70509d97ec1550a59364be6438147e.
//
// Solidity: event VenusGranted(address recipient, uint256 amount)
func (_VenusComproller *VenusComprollerFilterer) FilterVenusGranted(opts *bind.FilterOpts) (*VenusComprollerVenusGrantedIterator, error) {

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "VenusGranted")
	if err != nil {
		return nil, err
	}
	return &VenusComprollerVenusGrantedIterator{contract: _VenusComproller.contract, event: "VenusGranted", logs: logs, sub: sub}, nil
}

// WatchVenusGranted is a free log subscription operation binding the contract event 0xd7fe674cac9eee3998fe3cbd7a6f93c3bc70509d97ec1550a59364be6438147e.
//
// Solidity: event VenusGranted(address recipient, uint256 amount)
func (_VenusComproller *VenusComprollerFilterer) WatchVenusGranted(opts *bind.WatchOpts, sink chan<- *VenusComprollerVenusGranted) (event.Subscription, error) {

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "VenusGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerVenusGranted)
				if err := _VenusComproller.contract.UnpackLog(event, "VenusGranted", log); err != nil {
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

// ParseVenusGranted is a log parse operation binding the contract event 0xd7fe674cac9eee3998fe3cbd7a6f93c3bc70509d97ec1550a59364be6438147e.
//
// Solidity: event VenusGranted(address recipient, uint256 amount)
func (_VenusComproller *VenusComprollerFilterer) ParseVenusGranted(log types.Log) (*VenusComprollerVenusGranted, error) {
	event := new(VenusComprollerVenusGranted)
	if err := _VenusComproller.contract.UnpackLog(event, "VenusGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusComprollerVenusSpeedUpdatedIterator is returned from FilterVenusSpeedUpdated and is used to iterate over the raw logs and unpacked data for VenusSpeedUpdated events raised by the VenusComproller contract.
type VenusComprollerVenusSpeedUpdatedIterator struct {
	Event *VenusComprollerVenusSpeedUpdated // Event containing the contract specifics and raw log

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
func (it *VenusComprollerVenusSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusComprollerVenusSpeedUpdated)
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
		it.Event = new(VenusComprollerVenusSpeedUpdated)
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
func (it *VenusComprollerVenusSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusComprollerVenusSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusComprollerVenusSpeedUpdated represents a VenusSpeedUpdated event raised by the VenusComproller contract.
type VenusComprollerVenusSpeedUpdated struct {
	VToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVenusSpeedUpdated is a free log retrieval operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_VenusComproller *VenusComprollerFilterer) FilterVenusSpeedUpdated(opts *bind.FilterOpts, vToken []common.Address) (*VenusComprollerVenusSpeedUpdatedIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusComproller.contract.FilterLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenusComprollerVenusSpeedUpdatedIterator{contract: _VenusComproller.contract, event: "VenusSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchVenusSpeedUpdated is a free log subscription operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_VenusComproller *VenusComprollerFilterer) WatchVenusSpeedUpdated(opts *bind.WatchOpts, sink chan<- *VenusComprollerVenusSpeedUpdated, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _VenusComproller.contract.WatchLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusComprollerVenusSpeedUpdated)
				if err := _VenusComproller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
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

// ParseVenusSpeedUpdated is a log parse operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_VenusComproller *VenusComprollerFilterer) ParseVenusSpeedUpdated(log types.Log) (*VenusComprollerVenusSpeedUpdated, error) {
	event := new(VenusComprollerVenusSpeedUpdated)
	if err := _VenusComproller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
