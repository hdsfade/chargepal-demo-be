// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

var Decimal = big.NewInt(1e18)

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

// ChargepalTokenMetaData contains all meta data concerning the ChargepalToken contract.
var ChargepalTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20InvalidCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TopUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"}],\"name\":\"addBurner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBurner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"}],\"name\":\"removeBurner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChargepalTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ChargepalTokenMetaData.ABI instead.
var ChargepalTokenABI = ChargepalTokenMetaData.ABI

// ChargepalToken is an auto generated Go binding around an Ethereum contract.
type ChargepalToken struct {
	ChargepalTokenCaller     // Read-only binding to the contract
	ChargepalTokenTransactor // Write-only binding to the contract
	ChargepalTokenFilterer   // Log filterer for contract events
}

// ChargepalTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChargepalTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargepalTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChargepalTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargepalTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChargepalTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargepalTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChargepalTokenSession struct {
	Contract     *ChargepalToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChargepalTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChargepalTokenCallerSession struct {
	Contract *ChargepalTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ChargepalTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChargepalTokenTransactorSession struct {
	Contract     *ChargepalTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ChargepalTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChargepalTokenRaw struct {
	Contract *ChargepalToken // Generic contract binding to access the raw methods on
}

// ChargepalTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChargepalTokenCallerRaw struct {
	Contract *ChargepalTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ChargepalTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChargepalTokenTransactorRaw struct {
	Contract *ChargepalTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChargepalToken creates a new instance of ChargepalToken, bound to a specific deployed contract.
func NewChargepalToken(address common.Address, backend bind.ContractBackend) (*ChargepalToken, error) {
	contract, err := bindChargepalToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChargepalToken{ChargepalTokenCaller: ChargepalTokenCaller{contract: contract}, ChargepalTokenTransactor: ChargepalTokenTransactor{contract: contract}, ChargepalTokenFilterer: ChargepalTokenFilterer{contract: contract}}, nil
}

// NewChargepalTokenCaller creates a new read-only instance of ChargepalToken, bound to a specific deployed contract.
func NewChargepalTokenCaller(address common.Address, caller bind.ContractCaller) (*ChargepalTokenCaller, error) {
	contract, err := bindChargepalToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenCaller{contract: contract}, nil
}

// NewChargepalTokenTransactor creates a new write-only instance of ChargepalToken, bound to a specific deployed contract.
func NewChargepalTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ChargepalTokenTransactor, error) {
	contract, err := bindChargepalToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenTransactor{contract: contract}, nil
}

// NewChargepalTokenFilterer creates a new log filterer instance of ChargepalToken, bound to a specific deployed contract.
func NewChargepalTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ChargepalTokenFilterer, error) {
	contract, err := bindChargepalToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenFilterer{contract: contract}, nil
}

// bindChargepalToken binds a generic wrapper to an already deployed contract.
func bindChargepalToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChargepalTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChargepalToken *ChargepalTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChargepalToken.Contract.ChargepalTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChargepalToken *ChargepalTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargepalToken.Contract.ChargepalTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChargepalToken *ChargepalTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChargepalToken.Contract.ChargepalTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChargepalToken *ChargepalTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChargepalToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChargepalToken *ChargepalTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargepalToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChargepalToken *ChargepalTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChargepalToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChargepalToken *ChargepalTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChargepalToken *ChargepalTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChargepalToken.Contract.Allowance(&_ChargepalToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ChargepalToken *ChargepalTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ChargepalToken.Contract.Allowance(&_ChargepalToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChargepalToken *ChargepalTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChargepalToken *ChargepalTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChargepalToken.Contract.BalanceOf(&_ChargepalToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ChargepalToken *ChargepalTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ChargepalToken.Contract.BalanceOf(&_ChargepalToken.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ChargepalToken *ChargepalTokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ChargepalToken *ChargepalTokenSession) Cap() (*big.Int, error) {
	return _ChargepalToken.Contract.Cap(&_ChargepalToken.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_ChargepalToken *ChargepalTokenCallerSession) Cap() (*big.Int, error) {
	return _ChargepalToken.Contract.Cap(&_ChargepalToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChargepalToken *ChargepalTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChargepalToken *ChargepalTokenSession) Decimals() (uint8, error) {
	return _ChargepalToken.Contract.Decimals(&_ChargepalToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChargepalToken *ChargepalTokenCallerSession) Decimals() (uint8, error) {
	return _ChargepalToken.Contract.Decimals(&_ChargepalToken.CallOpts)
}

// IsBurner is a free data retrieval call binding the contract method 0x4334614a.
//
// Solidity: function isBurner(address ) view returns(bool)
func (_ChargepalToken *ChargepalTokenCaller) IsBurner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "isBurner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBurner is a free data retrieval call binding the contract method 0x4334614a.
//
// Solidity: function isBurner(address ) view returns(bool)
func (_ChargepalToken *ChargepalTokenSession) IsBurner(arg0 common.Address) (bool, error) {
	return _ChargepalToken.Contract.IsBurner(&_ChargepalToken.CallOpts, arg0)
}

// IsBurner is a free data retrieval call binding the contract method 0x4334614a.
//
// Solidity: function isBurner(address ) view returns(bool)
func (_ChargepalToken *ChargepalTokenCallerSession) IsBurner(arg0 common.Address) (bool, error) {
	return _ChargepalToken.Contract.IsBurner(&_ChargepalToken.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_ChargepalToken *ChargepalTokenCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_ChargepalToken *ChargepalTokenSession) IsMinter(arg0 common.Address) (bool, error) {
	return _ChargepalToken.Contract.IsMinter(&_ChargepalToken.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_ChargepalToken *ChargepalTokenCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _ChargepalToken.Contract.IsMinter(&_ChargepalToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChargepalToken *ChargepalTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChargepalToken *ChargepalTokenSession) Name() (string, error) {
	return _ChargepalToken.Contract.Name(&_ChargepalToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ChargepalToken *ChargepalTokenCallerSession) Name() (string, error) {
	return _ChargepalToken.Contract.Name(&_ChargepalToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChargepalToken *ChargepalTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChargepalToken *ChargepalTokenSession) Owner() (common.Address, error) {
	return _ChargepalToken.Contract.Owner(&_ChargepalToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChargepalToken *ChargepalTokenCallerSession) Owner() (common.Address, error) {
	return _ChargepalToken.Contract.Owner(&_ChargepalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChargepalToken *ChargepalTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChargepalToken *ChargepalTokenSession) Symbol() (string, error) {
	return _ChargepalToken.Contract.Symbol(&_ChargepalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ChargepalToken *ChargepalTokenCallerSession) Symbol() (string, error) {
	return _ChargepalToken.Contract.Symbol(&_ChargepalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChargepalToken *ChargepalTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChargepalToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChargepalToken *ChargepalTokenSession) TotalSupply() (*big.Int, error) {
	return _ChargepalToken.Contract.TotalSupply(&_ChargepalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ChargepalToken *ChargepalTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ChargepalToken.Contract.TotalSupply(&_ChargepalToken.CallOpts)
}

// AddBurner is a paid mutator transaction binding the contract method 0xf44637ba.
//
// Solidity: function addBurner(address _burner) returns()
func (_ChargepalToken *ChargepalTokenTransactor) AddBurner(opts *bind.TransactOpts, _burner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "addBurner", _burner)
}

// AddBurner is a paid mutator transaction binding the contract method 0xf44637ba.
//
// Solidity: function addBurner(address _burner) returns()
func (_ChargepalToken *ChargepalTokenSession) AddBurner(_burner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.AddBurner(&_ChargepalToken.TransactOpts, _burner)
}

// AddBurner is a paid mutator transaction binding the contract method 0xf44637ba.
//
// Solidity: function addBurner(address _burner) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) AddBurner(_burner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.AddBurner(&_ChargepalToken.TransactOpts, _burner)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_ChargepalToken *ChargepalTokenTransactor) AddMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "addMinter", _minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_ChargepalToken *ChargepalTokenSession) AddMinter(_minter common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.AddMinter(&_ChargepalToken.TransactOpts, _minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address _minter) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) AddMinter(_minter common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.AddMinter(&_ChargepalToken.TransactOpts, _minter)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Approve(&_ChargepalToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Approve(&_ChargepalToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Burn(&_ChargepalToken.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Burn(&_ChargepalToken.TransactOpts, _from, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Mint(&_ChargepalToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Mint(&_ChargepalToken.TransactOpts, _to, _amount)
}

// RemoveBurner is a paid mutator transaction binding the contract method 0x02846858.
//
// Solidity: function removeBurner(address _burner) returns()
func (_ChargepalToken *ChargepalTokenTransactor) RemoveBurner(opts *bind.TransactOpts, _burner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "removeBurner", _burner)
}

// RemoveBurner is a paid mutator transaction binding the contract method 0x02846858.
//
// Solidity: function removeBurner(address _burner) returns()
func (_ChargepalToken *ChargepalTokenSession) RemoveBurner(_burner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.RemoveBurner(&_ChargepalToken.TransactOpts, _burner)
}

// RemoveBurner is a paid mutator transaction binding the contract method 0x02846858.
//
// Solidity: function removeBurner(address _burner) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) RemoveBurner(_burner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.RemoveBurner(&_ChargepalToken.TransactOpts, _burner)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address _minter) returns()
func (_ChargepalToken *ChargepalTokenTransactor) RemoveMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "removeMinter", _minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address _minter) returns()
func (_ChargepalToken *ChargepalTokenSession) RemoveMinter(_minter common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.RemoveMinter(&_ChargepalToken.TransactOpts, _minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address _minter) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) RemoveMinter(_minter common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.RemoveMinter(&_ChargepalToken.TransactOpts, _minter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChargepalToken *ChargepalTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChargepalToken *ChargepalTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChargepalToken.Contract.RenounceOwnership(&_ChargepalToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChargepalToken.Contract.RenounceOwnership(&_ChargepalToken.TransactOpts)
}

// TopUp is a paid mutator transaction binding the contract method 0x50017f3e.
//
// Solidity: function topUp(uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenTransactor) TopUp(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "topUp", _amount)
}

// TopUp is a paid mutator transaction binding the contract method 0x50017f3e.
//
// Solidity: function topUp(uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenSession) TopUp(_amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.TopUp(&_ChargepalToken.TransactOpts, _amount)
}

// TopUp is a paid mutator transaction binding the contract method 0x50017f3e.
//
// Solidity: function topUp(uint256 _amount) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) TopUp(_amount *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.TopUp(&_ChargepalToken.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Transfer(&_ChargepalToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.Transfer(&_ChargepalToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.TransferFrom(&_ChargepalToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ChargepalToken *ChargepalTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ChargepalToken.Contract.TransferFrom(&_ChargepalToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChargepalToken *ChargepalTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChargepalToken *ChargepalTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.TransferOwnership(&_ChargepalToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChargepalToken *ChargepalTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChargepalToken.Contract.TransferOwnership(&_ChargepalToken.TransactOpts, newOwner)
}

// ChargepalTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ChargepalToken contract.
type ChargepalTokenApprovalIterator struct {
	Event *ChargepalTokenApproval // Event containing the contract specifics and raw log

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
func (it *ChargepalTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargepalTokenApproval)
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
		it.Event = new(ChargepalTokenApproval)
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
func (it *ChargepalTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargepalTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargepalTokenApproval represents a Approval event raised by the ChargepalToken contract.
type ChargepalTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChargepalToken *ChargepalTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ChargepalTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChargepalToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenApprovalIterator{contract: _ChargepalToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChargepalToken *ChargepalTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChargepalTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ChargepalToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargepalTokenApproval)
				if err := _ChargepalToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ChargepalToken *ChargepalTokenFilterer) ParseApproval(log types.Log) (*ChargepalTokenApproval, error) {
	event := new(ChargepalTokenApproval)
	if err := _ChargepalToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargepalTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChargepalToken contract.
type ChargepalTokenOwnershipTransferredIterator struct {
	Event *ChargepalTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChargepalTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargepalTokenOwnershipTransferred)
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
		it.Event = new(ChargepalTokenOwnershipTransferred)
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
func (it *ChargepalTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargepalTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargepalTokenOwnershipTransferred represents a OwnershipTransferred event raised by the ChargepalToken contract.
type ChargepalTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChargepalToken *ChargepalTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChargepalTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChargepalToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenOwnershipTransferredIterator{contract: _ChargepalToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChargepalToken *ChargepalTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChargepalTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChargepalToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargepalTokenOwnershipTransferred)
				if err := _ChargepalToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChargepalToken *ChargepalTokenFilterer) ParseOwnershipTransferred(log types.Log) (*ChargepalTokenOwnershipTransferred, error) {
	event := new(ChargepalTokenOwnershipTransferred)
	if err := _ChargepalToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargepalTokenTopUpIterator is returned from FilterTopUp and is used to iterate over the raw logs and unpacked data for TopUp events raised by the ChargepalToken contract.
type ChargepalTokenTopUpIterator struct {
	Event *ChargepalTokenTopUp // Event containing the contract specifics and raw log

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
func (it *ChargepalTokenTopUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargepalTokenTopUp)
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
		it.Event = new(ChargepalTokenTopUp)
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
func (it *ChargepalTokenTopUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargepalTokenTopUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargepalTokenTopUp represents a TopUp event raised by the ChargepalToken contract.
type ChargepalTokenTopUp struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopUp is a free log retrieval operation binding the contract event 0x29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df.
//
// Solidity: event TopUp(address indexed sender, uint256 amount)
func (_ChargepalToken *ChargepalTokenFilterer) FilterTopUp(opts *bind.FilterOpts, sender []common.Address) (*ChargepalTokenTopUpIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChargepalToken.contract.FilterLogs(opts, "TopUp", senderRule)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenTopUpIterator{contract: _ChargepalToken.contract, event: "TopUp", logs: logs, sub: sub}, nil
}

// WatchTopUp is a free log subscription operation binding the contract event 0x29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df.
//
// Solidity: event TopUp(address indexed sender, uint256 amount)
func (_ChargepalToken *ChargepalTokenFilterer) WatchTopUp(opts *bind.WatchOpts, sink chan<- *ChargepalTokenTopUp, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChargepalToken.contract.WatchLogs(opts, "TopUp", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargepalTokenTopUp)
				if err := _ChargepalToken.contract.UnpackLog(event, "TopUp", log); err != nil {
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

// ParseTopUp is a log parse operation binding the contract event 0x29e6701fa948667bae25d3e2ef7966f2a11754e98aa0dc52c55133430caee2df.
//
// Solidity: event TopUp(address indexed sender, uint256 amount)
func (_ChargepalToken *ChargepalTokenFilterer) ParseTopUp(log types.Log) (*ChargepalTokenTopUp, error) {
	event := new(ChargepalTokenTopUp)
	if err := _ChargepalToken.contract.UnpackLog(event, "TopUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargepalTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ChargepalToken contract.
type ChargepalTokenTransferIterator struct {
	Event *ChargepalTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ChargepalTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargepalTokenTransfer)
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
		it.Event = new(ChargepalTokenTransfer)
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
func (it *ChargepalTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargepalTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargepalTokenTransfer represents a Transfer event raised by the ChargepalToken contract.
type ChargepalTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChargepalToken *ChargepalTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChargepalTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChargepalToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChargepalTokenTransferIterator{contract: _ChargepalToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChargepalToken *ChargepalTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChargepalTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChargepalToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargepalTokenTransfer)
				if err := _ChargepalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ChargepalToken *ChargepalTokenFilterer) ParseTransfer(log types.Log) (*ChargepalTokenTransfer, error) {
	event := new(ChargepalTokenTransfer)
	if err := _ChargepalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
