// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeswap

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
)

// PancakePredictionV2BetInfo is an auto generated low-level Go binding around an user-defined struct.
type PancakePredictionV2BetInfo struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}

// PredictionMetaData contains all meta data concerning the Prediction contract.
var PredictionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_intervalSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bufferSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oracleUpdateAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetBear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetBull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"EndRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"LockRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdminAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bufferSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"intervalSeconds\",\"type\":\"uint256\"}],\"name\":\"NewBufferAndIntervalSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBetAmount\",\"type\":\"uint256\"}],\"name\":\"NewMinBetAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NewOperatorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oracleUpdateAllowance\",\"type\":\"uint256\"}],\"name\":\"NewOracleUpdateAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardBaseCalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsCalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"StartRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TreasuryClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TREASURY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"betBear\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"betBull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bufferSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisLockOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisLockRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisStartOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisStartRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getUserRounds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"enumPancakePredictionV2.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"internalType\":\"structPancakePredictionV2.BetInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserRoundsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intervalSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"enumPancakePredictionV2.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleLatestRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleUpdateAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"refundable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"lockPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"closePrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lockOracleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closeOracleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bullAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bearAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBaseCalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleCalled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bufferSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_intervalSeconds\",\"type\":\"uint256\"}],\"name\":\"setBufferAndIntervalSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBetAmount\",\"type\":\"uint256\"}],\"name\":\"setMinBetAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_oracleUpdateAllowance\",\"type\":\"uint256\"}],\"name\":\"setOracleUpdateAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userRounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PredictionABI is the input ABI used to generate the binding from.
// Deprecated: Use PredictionMetaData.ABI instead.
var PredictionABI = PredictionMetaData.ABI

// Prediction is an auto generated Go binding around an Ethereum contract.
type Prediction struct {
	PredictionCaller     // Read-only binding to the contract
	PredictionTransactor // Write-only binding to the contract
	PredictionFilterer   // Log filterer for contract events
}

// PredictionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictionSession struct {
	Contract     *Prediction       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PredictionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictionCallerSession struct {
	Contract *PredictionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PredictionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictionTransactorSession struct {
	Contract     *PredictionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PredictionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictionRaw struct {
	Contract *Prediction // Generic contract binding to access the raw methods on
}

// PredictionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictionCallerRaw struct {
	Contract *PredictionCaller // Generic read-only contract binding to access the raw methods on
}

// PredictionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictionTransactorRaw struct {
	Contract *PredictionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrediction creates a new instance of Prediction, bound to a specific deployed contract.
func NewPrediction(address common.Address, backend bind.ContractBackend) (*Prediction, error) {
	contract, err := bindPrediction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Prediction{PredictionCaller: PredictionCaller{contract: contract}, PredictionTransactor: PredictionTransactor{contract: contract}, PredictionFilterer: PredictionFilterer{contract: contract}}, nil
}

// NewPredictionCaller creates a new read-only instance of Prediction, bound to a specific deployed contract.
func NewPredictionCaller(address common.Address, caller bind.ContractCaller) (*PredictionCaller, error) {
	contract, err := bindPrediction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionCaller{contract: contract}, nil
}

// NewPredictionTransactor creates a new write-only instance of Prediction, bound to a specific deployed contract.
func NewPredictionTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictionTransactor, error) {
	contract, err := bindPrediction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionTransactor{contract: contract}, nil
}

// NewPredictionFilterer creates a new log filterer instance of Prediction, bound to a specific deployed contract.
func NewPredictionFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictionFilterer, error) {
	contract, err := bindPrediction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictionFilterer{contract: contract}, nil
}

// bindPrediction binds a generic wrapper to an already deployed contract.
func bindPrediction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PredictionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prediction *PredictionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prediction.Contract.PredictionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prediction *PredictionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.Contract.PredictionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prediction *PredictionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prediction.Contract.PredictionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prediction *PredictionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Prediction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Prediction *PredictionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Prediction *PredictionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Prediction.Contract.contract.Transact(opts, method, params...)
}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_Prediction *PredictionCaller) MAXTREASURYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "MAX_TREASURY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_Prediction *PredictionSession) MAXTREASURYFEE() (*big.Int, error) {
	return _Prediction.Contract.MAXTREASURYFEE(&_Prediction.CallOpts)
}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_Prediction *PredictionCallerSession) MAXTREASURYFEE() (*big.Int, error) {
	return _Prediction.Contract.MAXTREASURYFEE(&_Prediction.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Prediction *PredictionCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "adminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Prediction *PredictionSession) AdminAddress() (common.Address, error) {
	return _Prediction.Contract.AdminAddress(&_Prediction.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Prediction *PredictionCallerSession) AdminAddress() (common.Address, error) {
	return _Prediction.Contract.AdminAddress(&_Prediction.CallOpts)
}

// BufferSeconds is a free data retrieval call binding the contract method 0xeaba2361.
//
// Solidity: function bufferSeconds() view returns(uint256)
func (_Prediction *PredictionCaller) BufferSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "bufferSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferSeconds is a free data retrieval call binding the contract method 0xeaba2361.
//
// Solidity: function bufferSeconds() view returns(uint256)
func (_Prediction *PredictionSession) BufferSeconds() (*big.Int, error) {
	return _Prediction.Contract.BufferSeconds(&_Prediction.CallOpts)
}

// BufferSeconds is a free data retrieval call binding the contract method 0xeaba2361.
//
// Solidity: function bufferSeconds() view returns(uint256)
func (_Prediction *PredictionCallerSession) BufferSeconds() (*big.Int, error) {
	return _Prediction.Contract.BufferSeconds(&_Prediction.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_Prediction *PredictionCaller) Claimable(opts *bind.CallOpts, epoch *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "claimable", epoch, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_Prediction *PredictionSession) Claimable(epoch *big.Int, user common.Address) (bool, error) {
	return _Prediction.Contract.Claimable(&_Prediction.CallOpts, epoch, user)
}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_Prediction *PredictionCallerSession) Claimable(epoch *big.Int, user common.Address) (bool, error) {
	return _Prediction.Contract.Claimable(&_Prediction.CallOpts, epoch, user)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Prediction *PredictionCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Prediction *PredictionSession) CurrentEpoch() (*big.Int, error) {
	return _Prediction.Contract.CurrentEpoch(&_Prediction.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Prediction *PredictionCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Prediction.Contract.CurrentEpoch(&_Prediction.CallOpts)
}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_Prediction *PredictionCaller) GenesisLockOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "genesisLockOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_Prediction *PredictionSession) GenesisLockOnce() (bool, error) {
	return _Prediction.Contract.GenesisLockOnce(&_Prediction.CallOpts)
}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_Prediction *PredictionCallerSession) GenesisLockOnce() (bool, error) {
	return _Prediction.Contract.GenesisLockOnce(&_Prediction.CallOpts)
}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_Prediction *PredictionCaller) GenesisStartOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "genesisStartOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_Prediction *PredictionSession) GenesisStartOnce() (bool, error) {
	return _Prediction.Contract.GenesisStartOnce(&_Prediction.CallOpts)
}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_Prediction *PredictionCallerSession) GenesisStartOnce() (bool, error) {
	return _Prediction.Contract.GenesisStartOnce(&_Prediction.CallOpts)
}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], (uint8,uint256,bool)[], uint256)
func (_Prediction *PredictionCaller) GetUserRounds(opts *bind.CallOpts, user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, []PancakePredictionV2BetInfo, *big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "getUserRounds", user, cursor, size)

	if err != nil {
		return *new([]*big.Int), *new([]PancakePredictionV2BetInfo), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]PancakePredictionV2BetInfo)).(*[]PancakePredictionV2BetInfo)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], (uint8,uint256,bool)[], uint256)
func (_Prediction *PredictionSession) GetUserRounds(user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, []PancakePredictionV2BetInfo, *big.Int, error) {
	return _Prediction.Contract.GetUserRounds(&_Prediction.CallOpts, user, cursor, size)
}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], (uint8,uint256,bool)[], uint256)
func (_Prediction *PredictionCallerSession) GetUserRounds(user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, []PancakePredictionV2BetInfo, *big.Int, error) {
	return _Prediction.Contract.GetUserRounds(&_Prediction.CallOpts, user, cursor, size)
}

// GetUserRoundsLength is a free data retrieval call binding the contract method 0x273867d4.
//
// Solidity: function getUserRoundsLength(address user) view returns(uint256)
func (_Prediction *PredictionCaller) GetUserRoundsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "getUserRoundsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRoundsLength is a free data retrieval call binding the contract method 0x273867d4.
//
// Solidity: function getUserRoundsLength(address user) view returns(uint256)
func (_Prediction *PredictionSession) GetUserRoundsLength(user common.Address) (*big.Int, error) {
	return _Prediction.Contract.GetUserRoundsLength(&_Prediction.CallOpts, user)
}

// GetUserRoundsLength is a free data retrieval call binding the contract method 0x273867d4.
//
// Solidity: function getUserRoundsLength(address user) view returns(uint256)
func (_Prediction *PredictionCallerSession) GetUserRoundsLength(user common.Address) (*big.Int, error) {
	return _Prediction.Contract.GetUserRoundsLength(&_Prediction.CallOpts, user)
}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_Prediction *PredictionCaller) IntervalSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "intervalSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_Prediction *PredictionSession) IntervalSeconds() (*big.Int, error) {
	return _Prediction.Contract.IntervalSeconds(&_Prediction.CallOpts)
}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_Prediction *PredictionCallerSession) IntervalSeconds() (*big.Int, error) {
	return _Prediction.Contract.IntervalSeconds(&_Prediction.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_Prediction *PredictionCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "ledger", arg0, arg1)

	outstruct := new(struct {
		Position uint8
		Amount   *big.Int
		Claimed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Position = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_Prediction *PredictionSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	return _Prediction.Contract.Ledger(&_Prediction.CallOpts, arg0, arg1)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_Prediction *PredictionCallerSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	return _Prediction.Contract.Ledger(&_Prediction.CallOpts, arg0, arg1)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Prediction *PredictionCaller) MinBetAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "minBetAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Prediction *PredictionSession) MinBetAmount() (*big.Int, error) {
	return _Prediction.Contract.MinBetAmount(&_Prediction.CallOpts)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Prediction *PredictionCallerSession) MinBetAmount() (*big.Int, error) {
	return _Prediction.Contract.MinBetAmount(&_Prediction.CallOpts)
}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_Prediction *PredictionCaller) OperatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "operatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_Prediction *PredictionSession) OperatorAddress() (common.Address, error) {
	return _Prediction.Contract.OperatorAddress(&_Prediction.CallOpts)
}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_Prediction *PredictionCallerSession) OperatorAddress() (common.Address, error) {
	return _Prediction.Contract.OperatorAddress(&_Prediction.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Prediction *PredictionCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Prediction *PredictionSession) Oracle() (common.Address, error) {
	return _Prediction.Contract.Oracle(&_Prediction.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Prediction *PredictionCallerSession) Oracle() (common.Address, error) {
	return _Prediction.Contract.Oracle(&_Prediction.CallOpts)
}

// OracleLatestRoundId is a free data retrieval call binding the contract method 0xec324703.
//
// Solidity: function oracleLatestRoundId() view returns(uint256)
func (_Prediction *PredictionCaller) OracleLatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "oracleLatestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleLatestRoundId is a free data retrieval call binding the contract method 0xec324703.
//
// Solidity: function oracleLatestRoundId() view returns(uint256)
func (_Prediction *PredictionSession) OracleLatestRoundId() (*big.Int, error) {
	return _Prediction.Contract.OracleLatestRoundId(&_Prediction.CallOpts)
}

// OracleLatestRoundId is a free data retrieval call binding the contract method 0xec324703.
//
// Solidity: function oracleLatestRoundId() view returns(uint256)
func (_Prediction *PredictionCallerSession) OracleLatestRoundId() (*big.Int, error) {
	return _Prediction.Contract.OracleLatestRoundId(&_Prediction.CallOpts)
}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_Prediction *PredictionCaller) OracleUpdateAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "oracleUpdateAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_Prediction *PredictionSession) OracleUpdateAllowance() (*big.Int, error) {
	return _Prediction.Contract.OracleUpdateAllowance(&_Prediction.CallOpts)
}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_Prediction *PredictionCallerSession) OracleUpdateAllowance() (*big.Int, error) {
	return _Prediction.Contract.OracleUpdateAllowance(&_Prediction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prediction *PredictionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prediction *PredictionSession) Owner() (common.Address, error) {
	return _Prediction.Contract.Owner(&_Prediction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Prediction *PredictionCallerSession) Owner() (common.Address, error) {
	return _Prediction.Contract.Owner(&_Prediction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Prediction *PredictionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Prediction *PredictionSession) Paused() (bool, error) {
	return _Prediction.Contract.Paused(&_Prediction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Prediction *PredictionCallerSession) Paused() (bool, error) {
	return _Prediction.Contract.Paused(&_Prediction.CallOpts)
}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_Prediction *PredictionCaller) Refundable(opts *bind.CallOpts, epoch *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "refundable", epoch, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_Prediction *PredictionSession) Refundable(epoch *big.Int, user common.Address) (bool, error) {
	return _Prediction.Contract.Refundable(&_Prediction.CallOpts, epoch, user)
}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_Prediction *PredictionCallerSession) Refundable(epoch *big.Int, user common.Address) (bool, error) {
	return _Prediction.Contract.Refundable(&_Prediction.CallOpts, epoch, user)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 lockOracleId, uint256 closeOracleId, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_Prediction *PredictionCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		Epoch               *big.Int
		StartTimestamp      *big.Int
		LockTimestamp       *big.Int
		CloseTimestamp      *big.Int
		LockPrice           *big.Int
		ClosePrice          *big.Int
		LockOracleId        *big.Int
		CloseOracleId       *big.Int
		TotalAmount         *big.Int
		BullAmount          *big.Int
		BearAmount          *big.Int
		RewardBaseCalAmount *big.Int
		RewardAmount        *big.Int
		OracleCalled        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CloseTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ClosePrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LockOracleId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CloseOracleId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BullAmount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BearAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.RewardBaseCalAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.RewardAmount = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.OracleCalled = *abi.ConvertType(out[13], new(bool)).(*bool)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 lockOracleId, uint256 closeOracleId, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_Prediction *PredictionSession) Rounds(arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	return _Prediction.Contract.Rounds(&_Prediction.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 lockOracleId, uint256 closeOracleId, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_Prediction *PredictionCallerSession) Rounds(arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	return _Prediction.Contract.Rounds(&_Prediction.CallOpts, arg0)
}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_Prediction *PredictionCaller) TreasuryAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "treasuryAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_Prediction *PredictionSession) TreasuryAmount() (*big.Int, error) {
	return _Prediction.Contract.TreasuryAmount(&_Prediction.CallOpts)
}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_Prediction *PredictionCallerSession) TreasuryAmount() (*big.Int, error) {
	return _Prediction.Contract.TreasuryAmount(&_Prediction.CallOpts)
}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_Prediction *PredictionCaller) TreasuryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "treasuryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_Prediction *PredictionSession) TreasuryFee() (*big.Int, error) {
	return _Prediction.Contract.TreasuryFee(&_Prediction.CallOpts)
}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_Prediction *PredictionCallerSession) TreasuryFee() (*big.Int, error) {
	return _Prediction.Contract.TreasuryFee(&_Prediction.CallOpts)
}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_Prediction *PredictionCaller) UserRounds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Prediction.contract.Call(opts, &out, "userRounds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_Prediction *PredictionSession) UserRounds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Prediction.Contract.UserRounds(&_Prediction.CallOpts, arg0, arg1)
}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_Prediction *PredictionCallerSession) UserRounds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Prediction.Contract.UserRounds(&_Prediction.CallOpts, arg0, arg1)
}

// BetBear is a paid mutator transaction binding the contract method 0xaa6b873a.
//
// Solidity: function betBear(uint256 epoch) payable returns()
func (_Prediction *PredictionTransactor) BetBear(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "betBear", epoch)
}

// BetBear is a paid mutator transaction binding the contract method 0xaa6b873a.
//
// Solidity: function betBear(uint256 epoch) payable returns()
func (_Prediction *PredictionSession) BetBear(epoch *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.BetBear(&_Prediction.TransactOpts, epoch)
}

// BetBear is a paid mutator transaction binding the contract method 0xaa6b873a.
//
// Solidity: function betBear(uint256 epoch) payable returns()
func (_Prediction *PredictionTransactorSession) BetBear(epoch *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.BetBear(&_Prediction.TransactOpts, epoch)
}

// BetBull is a paid mutator transaction binding the contract method 0x57fb096f.
//
// Solidity: function betBull(uint256 epoch) payable returns()
func (_Prediction *PredictionTransactor) BetBull(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "betBull", epoch)
}

// BetBull is a paid mutator transaction binding the contract method 0x57fb096f.
//
// Solidity: function betBull(uint256 epoch) payable returns()
func (_Prediction *PredictionSession) BetBull(epoch *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.BetBull(&_Prediction.TransactOpts, epoch)
}

// BetBull is a paid mutator transaction binding the contract method 0x57fb096f.
//
// Solidity: function betBull(uint256 epoch) payable returns()
func (_Prediction *PredictionTransactorSession) BetBull(epoch *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.BetBull(&_Prediction.TransactOpts, epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x6ba4c138.
//
// Solidity: function claim(uint256[] epochs) returns()
func (_Prediction *PredictionTransactor) Claim(opts *bind.TransactOpts, epochs []*big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "claim", epochs)
}

// Claim is a paid mutator transaction binding the contract method 0x6ba4c138.
//
// Solidity: function claim(uint256[] epochs) returns()
func (_Prediction *PredictionSession) Claim(epochs []*big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.Claim(&_Prediction.TransactOpts, epochs)
}

// Claim is a paid mutator transaction binding the contract method 0x6ba4c138.
//
// Solidity: function claim(uint256[] epochs) returns()
func (_Prediction *PredictionTransactorSession) Claim(epochs []*big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.Claim(&_Prediction.TransactOpts, epochs)
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_Prediction *PredictionTransactor) ClaimTreasury(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "claimTreasury")
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_Prediction *PredictionSession) ClaimTreasury() (*types.Transaction, error) {
	return _Prediction.Contract.ClaimTreasury(&_Prediction.TransactOpts)
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_Prediction *PredictionTransactorSession) ClaimTreasury() (*types.Transaction, error) {
	return _Prediction.Contract.ClaimTreasury(&_Prediction.TransactOpts)
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_Prediction *PredictionTransactor) ExecuteRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "executeRound")
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_Prediction *PredictionSession) ExecuteRound() (*types.Transaction, error) {
	return _Prediction.Contract.ExecuteRound(&_Prediction.TransactOpts)
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_Prediction *PredictionTransactorSession) ExecuteRound() (*types.Transaction, error) {
	return _Prediction.Contract.ExecuteRound(&_Prediction.TransactOpts)
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_Prediction *PredictionTransactor) GenesisLockRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "genesisLockRound")
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_Prediction *PredictionSession) GenesisLockRound() (*types.Transaction, error) {
	return _Prediction.Contract.GenesisLockRound(&_Prediction.TransactOpts)
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_Prediction *PredictionTransactorSession) GenesisLockRound() (*types.Transaction, error) {
	return _Prediction.Contract.GenesisLockRound(&_Prediction.TransactOpts)
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_Prediction *PredictionTransactor) GenesisStartRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "genesisStartRound")
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_Prediction *PredictionSession) GenesisStartRound() (*types.Transaction, error) {
	return _Prediction.Contract.GenesisStartRound(&_Prediction.TransactOpts)
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_Prediction *PredictionTransactorSession) GenesisStartRound() (*types.Transaction, error) {
	return _Prediction.Contract.GenesisStartRound(&_Prediction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Prediction *PredictionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Prediction *PredictionSession) Pause() (*types.Transaction, error) {
	return _Prediction.Contract.Pause(&_Prediction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Prediction *PredictionTransactorSession) Pause() (*types.Transaction, error) {
	return _Prediction.Contract.Pause(&_Prediction.TransactOpts)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xb29a8140.
//
// Solidity: function recoverToken(address _token, uint256 _amount) returns()
func (_Prediction *PredictionTransactor) RecoverToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "recoverToken", _token, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xb29a8140.
//
// Solidity: function recoverToken(address _token, uint256 _amount) returns()
func (_Prediction *PredictionSession) RecoverToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.RecoverToken(&_Prediction.TransactOpts, _token, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xb29a8140.
//
// Solidity: function recoverToken(address _token, uint256 _amount) returns()
func (_Prediction *PredictionTransactorSession) RecoverToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.RecoverToken(&_Prediction.TransactOpts, _token, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prediction *PredictionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prediction *PredictionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prediction.Contract.RenounceOwnership(&_Prediction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Prediction *PredictionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Prediction.Contract.RenounceOwnership(&_Prediction.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_Prediction *PredictionTransactor) SetAdmin(opts *bind.TransactOpts, _adminAddress common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setAdmin", _adminAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_Prediction *PredictionSession) SetAdmin(_adminAddress common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.SetAdmin(&_Prediction.TransactOpts, _adminAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_Prediction *PredictionTransactorSession) SetAdmin(_adminAddress common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.SetAdmin(&_Prediction.TransactOpts, _adminAddress)
}

// SetBufferAndIntervalSeconds is a paid mutator transaction binding the contract method 0x890dc766.
//
// Solidity: function setBufferAndIntervalSeconds(uint256 _bufferSeconds, uint256 _intervalSeconds) returns()
func (_Prediction *PredictionTransactor) SetBufferAndIntervalSeconds(opts *bind.TransactOpts, _bufferSeconds *big.Int, _intervalSeconds *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setBufferAndIntervalSeconds", _bufferSeconds, _intervalSeconds)
}

// SetBufferAndIntervalSeconds is a paid mutator transaction binding the contract method 0x890dc766.
//
// Solidity: function setBufferAndIntervalSeconds(uint256 _bufferSeconds, uint256 _intervalSeconds) returns()
func (_Prediction *PredictionSession) SetBufferAndIntervalSeconds(_bufferSeconds *big.Int, _intervalSeconds *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetBufferAndIntervalSeconds(&_Prediction.TransactOpts, _bufferSeconds, _intervalSeconds)
}

// SetBufferAndIntervalSeconds is a paid mutator transaction binding the contract method 0x890dc766.
//
// Solidity: function setBufferAndIntervalSeconds(uint256 _bufferSeconds, uint256 _intervalSeconds) returns()
func (_Prediction *PredictionTransactorSession) SetBufferAndIntervalSeconds(_bufferSeconds *big.Int, _intervalSeconds *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetBufferAndIntervalSeconds(&_Prediction.TransactOpts, _bufferSeconds, _intervalSeconds)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_Prediction *PredictionTransactor) SetMinBetAmount(opts *bind.TransactOpts, _minBetAmount *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setMinBetAmount", _minBetAmount)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_Prediction *PredictionSession) SetMinBetAmount(_minBetAmount *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetMinBetAmount(&_Prediction.TransactOpts, _minBetAmount)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_Prediction *PredictionTransactorSession) SetMinBetAmount(_minBetAmount *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetMinBetAmount(&_Prediction.TransactOpts, _minBetAmount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_Prediction *PredictionTransactor) SetOperator(opts *bind.TransactOpts, _operatorAddress common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setOperator", _operatorAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_Prediction *PredictionSession) SetOperator(_operatorAddress common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.SetOperator(&_Prediction.TransactOpts, _operatorAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_Prediction *PredictionTransactorSession) SetOperator(_operatorAddress common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.SetOperator(&_Prediction.TransactOpts, _operatorAddress)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Prediction *PredictionTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Prediction *PredictionSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.SetOracle(&_Prediction.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Prediction *PredictionTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.SetOracle(&_Prediction.TransactOpts, _oracle)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_Prediction *PredictionTransactor) SetOracleUpdateAllowance(opts *bind.TransactOpts, _oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setOracleUpdateAllowance", _oracleUpdateAllowance)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_Prediction *PredictionSession) SetOracleUpdateAllowance(_oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetOracleUpdateAllowance(&_Prediction.TransactOpts, _oracleUpdateAllowance)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_Prediction *PredictionTransactorSession) SetOracleUpdateAllowance(_oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetOracleUpdateAllowance(&_Prediction.TransactOpts, _oracleUpdateAllowance)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_Prediction *PredictionTransactor) SetTreasuryFee(opts *bind.TransactOpts, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "setTreasuryFee", _treasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_Prediction *PredictionSession) SetTreasuryFee(_treasuryFee *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetTreasuryFee(&_Prediction.TransactOpts, _treasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_Prediction *PredictionTransactorSession) SetTreasuryFee(_treasuryFee *big.Int) (*types.Transaction, error) {
	return _Prediction.Contract.SetTreasuryFee(&_Prediction.TransactOpts, _treasuryFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prediction *PredictionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prediction *PredictionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.TransferOwnership(&_Prediction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Prediction *PredictionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Prediction.Contract.TransferOwnership(&_Prediction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Prediction *PredictionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Prediction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Prediction *PredictionSession) Unpause() (*types.Transaction, error) {
	return _Prediction.Contract.Unpause(&_Prediction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Prediction *PredictionTransactorSession) Unpause() (*types.Transaction, error) {
	return _Prediction.Contract.Unpause(&_Prediction.TransactOpts)
}

// PredictionBetBearIterator is returned from FilterBetBear and is used to iterate over the raw logs and unpacked data for BetBear events raised by the Prediction contract.
type PredictionBetBearIterator struct {
	Event *PredictionBetBear // Event containing the contract specifics and raw log

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
func (it *PredictionBetBearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionBetBear)
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
		it.Event = new(PredictionBetBear)
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
func (it *PredictionBetBearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionBetBearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionBetBear represents a BetBear event raised by the Prediction contract.
type PredictionBetBear struct {
	Sender common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetBear is a free log retrieval operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) FilterBetBear(opts *bind.FilterOpts, sender []common.Address, epoch []*big.Int) (*PredictionBetBearIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "BetBear", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionBetBearIterator{contract: _Prediction.contract, event: "BetBear", logs: logs, sub: sub}, nil
}

// WatchBetBear is a free log subscription operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) WatchBetBear(opts *bind.WatchOpts, sink chan<- *PredictionBetBear, sender []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "BetBear", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionBetBear)
				if err := _Prediction.contract.UnpackLog(event, "BetBear", log); err != nil {
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

// ParseBetBear is a log parse operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) ParseBetBear(log types.Log) (*PredictionBetBear, error) {
	event := new(PredictionBetBear)
	if err := _Prediction.contract.UnpackLog(event, "BetBear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionBetBullIterator is returned from FilterBetBull and is used to iterate over the raw logs and unpacked data for BetBull events raised by the Prediction contract.
type PredictionBetBullIterator struct {
	Event *PredictionBetBull // Event containing the contract specifics and raw log

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
func (it *PredictionBetBullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionBetBull)
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
		it.Event = new(PredictionBetBull)
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
func (it *PredictionBetBullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionBetBullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionBetBull represents a BetBull event raised by the Prediction contract.
type PredictionBetBull struct {
	Sender common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetBull is a free log retrieval operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) FilterBetBull(opts *bind.FilterOpts, sender []common.Address, epoch []*big.Int) (*PredictionBetBullIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "BetBull", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionBetBullIterator{contract: _Prediction.contract, event: "BetBull", logs: logs, sub: sub}, nil
}

// WatchBetBull is a free log subscription operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) WatchBetBull(opts *bind.WatchOpts, sink chan<- *PredictionBetBull, sender []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "BetBull", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionBetBull)
				if err := _Prediction.contract.UnpackLog(event, "BetBull", log); err != nil {
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

// ParseBetBull is a log parse operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) ParseBetBull(log types.Log) (*PredictionBetBull, error) {
	event := new(PredictionBetBull)
	if err := _Prediction.contract.UnpackLog(event, "BetBull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Prediction contract.
type PredictionClaimIterator struct {
	Event *PredictionClaim // Event containing the contract specifics and raw log

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
func (it *PredictionClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionClaim)
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
		it.Event = new(PredictionClaim)
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
func (it *PredictionClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionClaim represents a Claim event raised by the Prediction contract.
type PredictionClaim struct {
	Sender common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) FilterClaim(opts *bind.FilterOpts, sender []common.Address, epoch []*big.Int) (*PredictionClaimIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Claim", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionClaimIterator{contract: _Prediction.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *PredictionClaim, sender []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Claim", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionClaim)
				if err := _Prediction.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_Prediction *PredictionFilterer) ParseClaim(log types.Log) (*PredictionClaim, error) {
	event := new(PredictionClaim)
	if err := _Prediction.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionEndRoundIterator is returned from FilterEndRound and is used to iterate over the raw logs and unpacked data for EndRound events raised by the Prediction contract.
type PredictionEndRoundIterator struct {
	Event *PredictionEndRound // Event containing the contract specifics and raw log

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
func (it *PredictionEndRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionEndRound)
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
		it.Event = new(PredictionEndRound)
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
func (it *PredictionEndRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionEndRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionEndRound represents a EndRound event raised by the Prediction contract.
type PredictionEndRound struct {
	Epoch   *big.Int
	RoundId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEndRound is a free log retrieval operation binding the contract event 0xb6ff1fe915db84788cbbbc017f0d2bef9485fad9fd0bd8ce9340fde0d8410dd8.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_Prediction *PredictionFilterer) FilterEndRound(opts *bind.FilterOpts, epoch []*big.Int, roundId []*big.Int) (*PredictionEndRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "EndRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionEndRoundIterator{contract: _Prediction.contract, event: "EndRound", logs: logs, sub: sub}, nil
}

// WatchEndRound is a free log subscription operation binding the contract event 0xb6ff1fe915db84788cbbbc017f0d2bef9485fad9fd0bd8ce9340fde0d8410dd8.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_Prediction *PredictionFilterer) WatchEndRound(opts *bind.WatchOpts, sink chan<- *PredictionEndRound, epoch []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "EndRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionEndRound)
				if err := _Prediction.contract.UnpackLog(event, "EndRound", log); err != nil {
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

// ParseEndRound is a log parse operation binding the contract event 0xb6ff1fe915db84788cbbbc017f0d2bef9485fad9fd0bd8ce9340fde0d8410dd8.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_Prediction *PredictionFilterer) ParseEndRound(log types.Log) (*PredictionEndRound, error) {
	event := new(PredictionEndRound)
	if err := _Prediction.contract.UnpackLog(event, "EndRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionLockRoundIterator is returned from FilterLockRound and is used to iterate over the raw logs and unpacked data for LockRound events raised by the Prediction contract.
type PredictionLockRoundIterator struct {
	Event *PredictionLockRound // Event containing the contract specifics and raw log

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
func (it *PredictionLockRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionLockRound)
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
		it.Event = new(PredictionLockRound)
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
func (it *PredictionLockRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionLockRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionLockRound represents a LockRound event raised by the Prediction contract.
type PredictionLockRound struct {
	Epoch   *big.Int
	RoundId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockRound is a free log retrieval operation binding the contract event 0x482e76a65b448a42deef26e99e58fb20c85e26f075defff8df6aa80459b39006.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_Prediction *PredictionFilterer) FilterLockRound(opts *bind.FilterOpts, epoch []*big.Int, roundId []*big.Int) (*PredictionLockRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "LockRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionLockRoundIterator{contract: _Prediction.contract, event: "LockRound", logs: logs, sub: sub}, nil
}

// WatchLockRound is a free log subscription operation binding the contract event 0x482e76a65b448a42deef26e99e58fb20c85e26f075defff8df6aa80459b39006.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_Prediction *PredictionFilterer) WatchLockRound(opts *bind.WatchOpts, sink chan<- *PredictionLockRound, epoch []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "LockRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionLockRound)
				if err := _Prediction.contract.UnpackLog(event, "LockRound", log); err != nil {
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

// ParseLockRound is a log parse operation binding the contract event 0x482e76a65b448a42deef26e99e58fb20c85e26f075defff8df6aa80459b39006.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_Prediction *PredictionFilterer) ParseLockRound(log types.Log) (*PredictionLockRound, error) {
	event := new(PredictionLockRound)
	if err := _Prediction.contract.UnpackLog(event, "LockRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewAdminAddressIterator is returned from FilterNewAdminAddress and is used to iterate over the raw logs and unpacked data for NewAdminAddress events raised by the Prediction contract.
type PredictionNewAdminAddressIterator struct {
	Event *PredictionNewAdminAddress // Event containing the contract specifics and raw log

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
func (it *PredictionNewAdminAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewAdminAddress)
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
		it.Event = new(PredictionNewAdminAddress)
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
func (it *PredictionNewAdminAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewAdminAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewAdminAddress represents a NewAdminAddress event raised by the Prediction contract.
type PredictionNewAdminAddress struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdminAddress is a free log retrieval operation binding the contract event 0x137b621413925496477d46e5055ac0d56178bdd724ba8bf843afceef18268ba3.
//
// Solidity: event NewAdminAddress(address admin)
func (_Prediction *PredictionFilterer) FilterNewAdminAddress(opts *bind.FilterOpts) (*PredictionNewAdminAddressIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewAdminAddress")
	if err != nil {
		return nil, err
	}
	return &PredictionNewAdminAddressIterator{contract: _Prediction.contract, event: "NewAdminAddress", logs: logs, sub: sub}, nil
}

// WatchNewAdminAddress is a free log subscription operation binding the contract event 0x137b621413925496477d46e5055ac0d56178bdd724ba8bf843afceef18268ba3.
//
// Solidity: event NewAdminAddress(address admin)
func (_Prediction *PredictionFilterer) WatchNewAdminAddress(opts *bind.WatchOpts, sink chan<- *PredictionNewAdminAddress) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewAdminAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewAdminAddress)
				if err := _Prediction.contract.UnpackLog(event, "NewAdminAddress", log); err != nil {
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

// ParseNewAdminAddress is a log parse operation binding the contract event 0x137b621413925496477d46e5055ac0d56178bdd724ba8bf843afceef18268ba3.
//
// Solidity: event NewAdminAddress(address admin)
func (_Prediction *PredictionFilterer) ParseNewAdminAddress(log types.Log) (*PredictionNewAdminAddress, error) {
	event := new(PredictionNewAdminAddress)
	if err := _Prediction.contract.UnpackLog(event, "NewAdminAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewBufferAndIntervalSecondsIterator is returned from FilterNewBufferAndIntervalSeconds and is used to iterate over the raw logs and unpacked data for NewBufferAndIntervalSeconds events raised by the Prediction contract.
type PredictionNewBufferAndIntervalSecondsIterator struct {
	Event *PredictionNewBufferAndIntervalSeconds // Event containing the contract specifics and raw log

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
func (it *PredictionNewBufferAndIntervalSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewBufferAndIntervalSeconds)
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
		it.Event = new(PredictionNewBufferAndIntervalSeconds)
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
func (it *PredictionNewBufferAndIntervalSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewBufferAndIntervalSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewBufferAndIntervalSeconds represents a NewBufferAndIntervalSeconds event raised by the Prediction contract.
type PredictionNewBufferAndIntervalSeconds struct {
	BufferSeconds   *big.Int
	IntervalSeconds *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewBufferAndIntervalSeconds is a free log retrieval operation binding the contract event 0xe60149e0431fec12df63dfab5fce2a9cefe9a4d3df5f41cb626f579ae1f2b91a.
//
// Solidity: event NewBufferAndIntervalSeconds(uint256 bufferSeconds, uint256 intervalSeconds)
func (_Prediction *PredictionFilterer) FilterNewBufferAndIntervalSeconds(opts *bind.FilterOpts) (*PredictionNewBufferAndIntervalSecondsIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewBufferAndIntervalSeconds")
	if err != nil {
		return nil, err
	}
	return &PredictionNewBufferAndIntervalSecondsIterator{contract: _Prediction.contract, event: "NewBufferAndIntervalSeconds", logs: logs, sub: sub}, nil
}

// WatchNewBufferAndIntervalSeconds is a free log subscription operation binding the contract event 0xe60149e0431fec12df63dfab5fce2a9cefe9a4d3df5f41cb626f579ae1f2b91a.
//
// Solidity: event NewBufferAndIntervalSeconds(uint256 bufferSeconds, uint256 intervalSeconds)
func (_Prediction *PredictionFilterer) WatchNewBufferAndIntervalSeconds(opts *bind.WatchOpts, sink chan<- *PredictionNewBufferAndIntervalSeconds) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewBufferAndIntervalSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewBufferAndIntervalSeconds)
				if err := _Prediction.contract.UnpackLog(event, "NewBufferAndIntervalSeconds", log); err != nil {
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

// ParseNewBufferAndIntervalSeconds is a log parse operation binding the contract event 0xe60149e0431fec12df63dfab5fce2a9cefe9a4d3df5f41cb626f579ae1f2b91a.
//
// Solidity: event NewBufferAndIntervalSeconds(uint256 bufferSeconds, uint256 intervalSeconds)
func (_Prediction *PredictionFilterer) ParseNewBufferAndIntervalSeconds(log types.Log) (*PredictionNewBufferAndIntervalSeconds, error) {
	event := new(PredictionNewBufferAndIntervalSeconds)
	if err := _Prediction.contract.UnpackLog(event, "NewBufferAndIntervalSeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewMinBetAmountIterator is returned from FilterNewMinBetAmount and is used to iterate over the raw logs and unpacked data for NewMinBetAmount events raised by the Prediction contract.
type PredictionNewMinBetAmountIterator struct {
	Event *PredictionNewMinBetAmount // Event containing the contract specifics and raw log

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
func (it *PredictionNewMinBetAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewMinBetAmount)
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
		it.Event = new(PredictionNewMinBetAmount)
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
func (it *PredictionNewMinBetAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewMinBetAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewMinBetAmount represents a NewMinBetAmount event raised by the Prediction contract.
type PredictionNewMinBetAmount struct {
	Epoch        *big.Int
	MinBetAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewMinBetAmount is a free log retrieval operation binding the contract event 0x90eb87c560a0213754ceb3a7fa3012f01acab0a35602c1e1995adf69dabc9d50.
//
// Solidity: event NewMinBetAmount(uint256 indexed epoch, uint256 minBetAmount)
func (_Prediction *PredictionFilterer) FilterNewMinBetAmount(opts *bind.FilterOpts, epoch []*big.Int) (*PredictionNewMinBetAmountIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewMinBetAmount", epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionNewMinBetAmountIterator{contract: _Prediction.contract, event: "NewMinBetAmount", logs: logs, sub: sub}, nil
}

// WatchNewMinBetAmount is a free log subscription operation binding the contract event 0x90eb87c560a0213754ceb3a7fa3012f01acab0a35602c1e1995adf69dabc9d50.
//
// Solidity: event NewMinBetAmount(uint256 indexed epoch, uint256 minBetAmount)
func (_Prediction *PredictionFilterer) WatchNewMinBetAmount(opts *bind.WatchOpts, sink chan<- *PredictionNewMinBetAmount, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewMinBetAmount", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewMinBetAmount)
				if err := _Prediction.contract.UnpackLog(event, "NewMinBetAmount", log); err != nil {
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

// ParseNewMinBetAmount is a log parse operation binding the contract event 0x90eb87c560a0213754ceb3a7fa3012f01acab0a35602c1e1995adf69dabc9d50.
//
// Solidity: event NewMinBetAmount(uint256 indexed epoch, uint256 minBetAmount)
func (_Prediction *PredictionFilterer) ParseNewMinBetAmount(log types.Log) (*PredictionNewMinBetAmount, error) {
	event := new(PredictionNewMinBetAmount)
	if err := _Prediction.contract.UnpackLog(event, "NewMinBetAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewOperatorAddressIterator is returned from FilterNewOperatorAddress and is used to iterate over the raw logs and unpacked data for NewOperatorAddress events raised by the Prediction contract.
type PredictionNewOperatorAddressIterator struct {
	Event *PredictionNewOperatorAddress // Event containing the contract specifics and raw log

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
func (it *PredictionNewOperatorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewOperatorAddress)
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
		it.Event = new(PredictionNewOperatorAddress)
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
func (it *PredictionNewOperatorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewOperatorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewOperatorAddress represents a NewOperatorAddress event raised by the Prediction contract.
type PredictionNewOperatorAddress struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOperatorAddress is a free log retrieval operation binding the contract event 0xc47d127c07bdd56c5ccba00463ce3bd3c1bca71b4670eea6e5d0c02e4aa156e2.
//
// Solidity: event NewOperatorAddress(address operator)
func (_Prediction *PredictionFilterer) FilterNewOperatorAddress(opts *bind.FilterOpts) (*PredictionNewOperatorAddressIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewOperatorAddress")
	if err != nil {
		return nil, err
	}
	return &PredictionNewOperatorAddressIterator{contract: _Prediction.contract, event: "NewOperatorAddress", logs: logs, sub: sub}, nil
}

// WatchNewOperatorAddress is a free log subscription operation binding the contract event 0xc47d127c07bdd56c5ccba00463ce3bd3c1bca71b4670eea6e5d0c02e4aa156e2.
//
// Solidity: event NewOperatorAddress(address operator)
func (_Prediction *PredictionFilterer) WatchNewOperatorAddress(opts *bind.WatchOpts, sink chan<- *PredictionNewOperatorAddress) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewOperatorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewOperatorAddress)
				if err := _Prediction.contract.UnpackLog(event, "NewOperatorAddress", log); err != nil {
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

// ParseNewOperatorAddress is a log parse operation binding the contract event 0xc47d127c07bdd56c5ccba00463ce3bd3c1bca71b4670eea6e5d0c02e4aa156e2.
//
// Solidity: event NewOperatorAddress(address operator)
func (_Prediction *PredictionFilterer) ParseNewOperatorAddress(log types.Log) (*PredictionNewOperatorAddress, error) {
	event := new(PredictionNewOperatorAddress)
	if err := _Prediction.contract.UnpackLog(event, "NewOperatorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewOracleIterator is returned from FilterNewOracle and is used to iterate over the raw logs and unpacked data for NewOracle events raised by the Prediction contract.
type PredictionNewOracleIterator struct {
	Event *PredictionNewOracle // Event containing the contract specifics and raw log

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
func (it *PredictionNewOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewOracle)
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
		it.Event = new(PredictionNewOracle)
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
func (it *PredictionNewOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewOracle represents a NewOracle event raised by the Prediction contract.
type PredictionNewOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOracle is a free log retrieval operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address oracle)
func (_Prediction *PredictionFilterer) FilterNewOracle(opts *bind.FilterOpts) (*PredictionNewOracleIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewOracle")
	if err != nil {
		return nil, err
	}
	return &PredictionNewOracleIterator{contract: _Prediction.contract, event: "NewOracle", logs: logs, sub: sub}, nil
}

// WatchNewOracle is a free log subscription operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address oracle)
func (_Prediction *PredictionFilterer) WatchNewOracle(opts *bind.WatchOpts, sink chan<- *PredictionNewOracle) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewOracle)
				if err := _Prediction.contract.UnpackLog(event, "NewOracle", log); err != nil {
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

// ParseNewOracle is a log parse operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address oracle)
func (_Prediction *PredictionFilterer) ParseNewOracle(log types.Log) (*PredictionNewOracle, error) {
	event := new(PredictionNewOracle)
	if err := _Prediction.contract.UnpackLog(event, "NewOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewOracleUpdateAllowanceIterator is returned from FilterNewOracleUpdateAllowance and is used to iterate over the raw logs and unpacked data for NewOracleUpdateAllowance events raised by the Prediction contract.
type PredictionNewOracleUpdateAllowanceIterator struct {
	Event *PredictionNewOracleUpdateAllowance // Event containing the contract specifics and raw log

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
func (it *PredictionNewOracleUpdateAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewOracleUpdateAllowance)
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
		it.Event = new(PredictionNewOracleUpdateAllowance)
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
func (it *PredictionNewOracleUpdateAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewOracleUpdateAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewOracleUpdateAllowance represents a NewOracleUpdateAllowance event raised by the Prediction contract.
type PredictionNewOracleUpdateAllowance struct {
	OracleUpdateAllowance *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewOracleUpdateAllowance is a free log retrieval operation binding the contract event 0x93ccaceac092ffb842c46b8718667a13a80e9058dcd0bd403d0b47215b30da07.
//
// Solidity: event NewOracleUpdateAllowance(uint256 oracleUpdateAllowance)
func (_Prediction *PredictionFilterer) FilterNewOracleUpdateAllowance(opts *bind.FilterOpts) (*PredictionNewOracleUpdateAllowanceIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewOracleUpdateAllowance")
	if err != nil {
		return nil, err
	}
	return &PredictionNewOracleUpdateAllowanceIterator{contract: _Prediction.contract, event: "NewOracleUpdateAllowance", logs: logs, sub: sub}, nil
}

// WatchNewOracleUpdateAllowance is a free log subscription operation binding the contract event 0x93ccaceac092ffb842c46b8718667a13a80e9058dcd0bd403d0b47215b30da07.
//
// Solidity: event NewOracleUpdateAllowance(uint256 oracleUpdateAllowance)
func (_Prediction *PredictionFilterer) WatchNewOracleUpdateAllowance(opts *bind.WatchOpts, sink chan<- *PredictionNewOracleUpdateAllowance) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewOracleUpdateAllowance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewOracleUpdateAllowance)
				if err := _Prediction.contract.UnpackLog(event, "NewOracleUpdateAllowance", log); err != nil {
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

// ParseNewOracleUpdateAllowance is a log parse operation binding the contract event 0x93ccaceac092ffb842c46b8718667a13a80e9058dcd0bd403d0b47215b30da07.
//
// Solidity: event NewOracleUpdateAllowance(uint256 oracleUpdateAllowance)
func (_Prediction *PredictionFilterer) ParseNewOracleUpdateAllowance(log types.Log) (*PredictionNewOracleUpdateAllowance, error) {
	event := new(PredictionNewOracleUpdateAllowance)
	if err := _Prediction.contract.UnpackLog(event, "NewOracleUpdateAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionNewTreasuryFeeIterator is returned from FilterNewTreasuryFee and is used to iterate over the raw logs and unpacked data for NewTreasuryFee events raised by the Prediction contract.
type PredictionNewTreasuryFeeIterator struct {
	Event *PredictionNewTreasuryFee // Event containing the contract specifics and raw log

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
func (it *PredictionNewTreasuryFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionNewTreasuryFee)
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
		it.Event = new(PredictionNewTreasuryFee)
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
func (it *PredictionNewTreasuryFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionNewTreasuryFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionNewTreasuryFee represents a NewTreasuryFee event raised by the Prediction contract.
type PredictionNewTreasuryFee struct {
	Epoch       *big.Int
	TreasuryFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryFee is a free log retrieval operation binding the contract event 0xb1c4ee38d35556741133da7ff9b6f7ab0fa88d0406133126ff128f635490a857.
//
// Solidity: event NewTreasuryFee(uint256 indexed epoch, uint256 treasuryFee)
func (_Prediction *PredictionFilterer) FilterNewTreasuryFee(opts *bind.FilterOpts, epoch []*big.Int) (*PredictionNewTreasuryFeeIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "NewTreasuryFee", epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionNewTreasuryFeeIterator{contract: _Prediction.contract, event: "NewTreasuryFee", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryFee is a free log subscription operation binding the contract event 0xb1c4ee38d35556741133da7ff9b6f7ab0fa88d0406133126ff128f635490a857.
//
// Solidity: event NewTreasuryFee(uint256 indexed epoch, uint256 treasuryFee)
func (_Prediction *PredictionFilterer) WatchNewTreasuryFee(opts *bind.WatchOpts, sink chan<- *PredictionNewTreasuryFee, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "NewTreasuryFee", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionNewTreasuryFee)
				if err := _Prediction.contract.UnpackLog(event, "NewTreasuryFee", log); err != nil {
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

// ParseNewTreasuryFee is a log parse operation binding the contract event 0xb1c4ee38d35556741133da7ff9b6f7ab0fa88d0406133126ff128f635490a857.
//
// Solidity: event NewTreasuryFee(uint256 indexed epoch, uint256 treasuryFee)
func (_Prediction *PredictionFilterer) ParseNewTreasuryFee(log types.Log) (*PredictionNewTreasuryFee, error) {
	event := new(PredictionNewTreasuryFee)
	if err := _Prediction.contract.UnpackLog(event, "NewTreasuryFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Prediction contract.
type PredictionOwnershipTransferredIterator struct {
	Event *PredictionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PredictionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionOwnershipTransferred)
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
		it.Event = new(PredictionOwnershipTransferred)
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
func (it *PredictionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionOwnershipTransferred represents a OwnershipTransferred event raised by the Prediction contract.
type PredictionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prediction *PredictionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PredictionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PredictionOwnershipTransferredIterator{contract: _Prediction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Prediction *PredictionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PredictionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionOwnershipTransferred)
				if err := _Prediction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Prediction *PredictionFilterer) ParseOwnershipTransferred(log types.Log) (*PredictionOwnershipTransferred, error) {
	event := new(PredictionOwnershipTransferred)
	if err := _Prediction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Prediction contract.
type PredictionPauseIterator struct {
	Event *PredictionPause // Event containing the contract specifics and raw log

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
func (it *PredictionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionPause)
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
		it.Event = new(PredictionPause)
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
func (it *PredictionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionPause represents a Pause event raised by the Prediction contract.
type PredictionPause struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) FilterPause(opts *bind.FilterOpts, epoch []*big.Int) (*PredictionPauseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Pause", epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionPauseIterator{contract: _Prediction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PredictionPause, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Pause", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionPause)
				if err := _Prediction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) ParsePause(log types.Log) (*PredictionPause, error) {
	event := new(PredictionPause)
	if err := _Prediction.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Prediction contract.
type PredictionPausedIterator struct {
	Event *PredictionPaused // Event containing the contract specifics and raw log

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
func (it *PredictionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionPaused)
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
		it.Event = new(PredictionPaused)
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
func (it *PredictionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionPaused represents a Paused event raised by the Prediction contract.
type PredictionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Prediction *PredictionFilterer) FilterPaused(opts *bind.FilterOpts) (*PredictionPausedIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PredictionPausedIterator{contract: _Prediction.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Prediction *PredictionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PredictionPaused) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionPaused)
				if err := _Prediction.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Prediction *PredictionFilterer) ParsePaused(log types.Log) (*PredictionPaused, error) {
	event := new(PredictionPaused)
	if err := _Prediction.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionRewardsCalculatedIterator is returned from FilterRewardsCalculated and is used to iterate over the raw logs and unpacked data for RewardsCalculated events raised by the Prediction contract.
type PredictionRewardsCalculatedIterator struct {
	Event *PredictionRewardsCalculated // Event containing the contract specifics and raw log

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
func (it *PredictionRewardsCalculatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionRewardsCalculated)
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
		it.Event = new(PredictionRewardsCalculated)
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
func (it *PredictionRewardsCalculatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionRewardsCalculatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionRewardsCalculated represents a RewardsCalculated event raised by the Prediction contract.
type PredictionRewardsCalculated struct {
	Epoch               *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	TreasuryAmount      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRewardsCalculated is a free log retrieval operation binding the contract event 0x6dfdfcb09c8804d0058826cd2539f1acfbe3cb887c9be03d928035bce0f1a58d.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount)
func (_Prediction *PredictionFilterer) FilterRewardsCalculated(opts *bind.FilterOpts, epoch []*big.Int) (*PredictionRewardsCalculatedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "RewardsCalculated", epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionRewardsCalculatedIterator{contract: _Prediction.contract, event: "RewardsCalculated", logs: logs, sub: sub}, nil
}

// WatchRewardsCalculated is a free log subscription operation binding the contract event 0x6dfdfcb09c8804d0058826cd2539f1acfbe3cb887c9be03d928035bce0f1a58d.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount)
func (_Prediction *PredictionFilterer) WatchRewardsCalculated(opts *bind.WatchOpts, sink chan<- *PredictionRewardsCalculated, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "RewardsCalculated", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionRewardsCalculated)
				if err := _Prediction.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
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

// ParseRewardsCalculated is a log parse operation binding the contract event 0x6dfdfcb09c8804d0058826cd2539f1acfbe3cb887c9be03d928035bce0f1a58d.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount)
func (_Prediction *PredictionFilterer) ParseRewardsCalculated(log types.Log) (*PredictionRewardsCalculated, error) {
	event := new(PredictionRewardsCalculated)
	if err := _Prediction.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionStartRoundIterator is returned from FilterStartRound and is used to iterate over the raw logs and unpacked data for StartRound events raised by the Prediction contract.
type PredictionStartRoundIterator struct {
	Event *PredictionStartRound // Event containing the contract specifics and raw log

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
func (it *PredictionStartRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionStartRound)
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
		it.Event = new(PredictionStartRound)
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
func (it *PredictionStartRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionStartRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionStartRound represents a StartRound event raised by the Prediction contract.
type PredictionStartRound struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartRound is a free log retrieval operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) FilterStartRound(opts *bind.FilterOpts, epoch []*big.Int) (*PredictionStartRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionStartRoundIterator{contract: _Prediction.contract, event: "StartRound", logs: logs, sub: sub}, nil
}

// WatchStartRound is a free log subscription operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) WatchStartRound(opts *bind.WatchOpts, sink chan<- *PredictionStartRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionStartRound)
				if err := _Prediction.contract.UnpackLog(event, "StartRound", log); err != nil {
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

// ParseStartRound is a log parse operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) ParseStartRound(log types.Log) (*PredictionStartRound, error) {
	event := new(PredictionStartRound)
	if err := _Prediction.contract.UnpackLog(event, "StartRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionTokenRecoveryIterator is returned from FilterTokenRecovery and is used to iterate over the raw logs and unpacked data for TokenRecovery events raised by the Prediction contract.
type PredictionTokenRecoveryIterator struct {
	Event *PredictionTokenRecovery // Event containing the contract specifics and raw log

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
func (it *PredictionTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionTokenRecovery)
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
		it.Event = new(PredictionTokenRecovery)
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
func (it *PredictionTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionTokenRecovery represents a TokenRecovery event raised by the Prediction contract.
type PredictionTokenRecovery struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRecovery is a free log retrieval operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_Prediction *PredictionFilterer) FilterTokenRecovery(opts *bind.FilterOpts, token []common.Address) (*PredictionTokenRecoveryIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PredictionTokenRecoveryIterator{contract: _Prediction.contract, event: "TokenRecovery", logs: logs, sub: sub}, nil
}

// WatchTokenRecovery is a free log subscription operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_Prediction *PredictionFilterer) WatchTokenRecovery(opts *bind.WatchOpts, sink chan<- *PredictionTokenRecovery, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionTokenRecovery)
				if err := _Prediction.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
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

// ParseTokenRecovery is a log parse operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_Prediction *PredictionFilterer) ParseTokenRecovery(log types.Log) (*PredictionTokenRecovery, error) {
	event := new(PredictionTokenRecovery)
	if err := _Prediction.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionTreasuryClaimIterator is returned from FilterTreasuryClaim and is used to iterate over the raw logs and unpacked data for TreasuryClaim events raised by the Prediction contract.
type PredictionTreasuryClaimIterator struct {
	Event *PredictionTreasuryClaim // Event containing the contract specifics and raw log

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
func (it *PredictionTreasuryClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionTreasuryClaim)
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
		it.Event = new(PredictionTreasuryClaim)
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
func (it *PredictionTreasuryClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionTreasuryClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionTreasuryClaim represents a TreasuryClaim event raised by the Prediction contract.
type PredictionTreasuryClaim struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTreasuryClaim is a free log retrieval operation binding the contract event 0xb9197c6b8e21274bd1e2d9c956a88af5cfee510f630fab3f046300f88b422361.
//
// Solidity: event TreasuryClaim(uint256 amount)
func (_Prediction *PredictionFilterer) FilterTreasuryClaim(opts *bind.FilterOpts) (*PredictionTreasuryClaimIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "TreasuryClaim")
	if err != nil {
		return nil, err
	}
	return &PredictionTreasuryClaimIterator{contract: _Prediction.contract, event: "TreasuryClaim", logs: logs, sub: sub}, nil
}

// WatchTreasuryClaim is a free log subscription operation binding the contract event 0xb9197c6b8e21274bd1e2d9c956a88af5cfee510f630fab3f046300f88b422361.
//
// Solidity: event TreasuryClaim(uint256 amount)
func (_Prediction *PredictionFilterer) WatchTreasuryClaim(opts *bind.WatchOpts, sink chan<- *PredictionTreasuryClaim) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "TreasuryClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionTreasuryClaim)
				if err := _Prediction.contract.UnpackLog(event, "TreasuryClaim", log); err != nil {
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

// ParseTreasuryClaim is a log parse operation binding the contract event 0xb9197c6b8e21274bd1e2d9c956a88af5cfee510f630fab3f046300f88b422361.
//
// Solidity: event TreasuryClaim(uint256 amount)
func (_Prediction *PredictionFilterer) ParseTreasuryClaim(log types.Log) (*PredictionTreasuryClaim, error) {
	event := new(PredictionTreasuryClaim)
	if err := _Prediction.contract.UnpackLog(event, "TreasuryClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Prediction contract.
type PredictionUnpauseIterator struct {
	Event *PredictionUnpause // Event containing the contract specifics and raw log

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
func (it *PredictionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionUnpause)
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
		it.Event = new(PredictionUnpause)
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
func (it *PredictionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionUnpause represents a Unpause event raised by the Prediction contract.
type PredictionUnpause struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) FilterUnpause(opts *bind.FilterOpts, epoch []*big.Int) (*PredictionUnpauseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Unpause", epochRule)
	if err != nil {
		return nil, err
	}
	return &PredictionUnpauseIterator{contract: _Prediction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PredictionUnpause, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Unpause", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionUnpause)
				if err := _Prediction.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 indexed epoch)
func (_Prediction *PredictionFilterer) ParseUnpause(log types.Log) (*PredictionUnpause, error) {
	event := new(PredictionUnpause)
	if err := _Prediction.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Prediction contract.
type PredictionUnpausedIterator struct {
	Event *PredictionUnpaused // Event containing the contract specifics and raw log

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
func (it *PredictionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionUnpaused)
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
		it.Event = new(PredictionUnpaused)
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
func (it *PredictionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionUnpaused represents a Unpaused event raised by the Prediction contract.
type PredictionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Prediction *PredictionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PredictionUnpausedIterator, error) {

	logs, sub, err := _Prediction.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PredictionUnpausedIterator{contract: _Prediction.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Prediction *PredictionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PredictionUnpaused) (event.Subscription, error) {

	logs, sub, err := _Prediction.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionUnpaused)
				if err := _Prediction.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Prediction *PredictionFilterer) ParseUnpaused(log types.Log) (*PredictionUnpaused, error) {
	event := new(PredictionUnpaused)
	if err := _Prediction.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
