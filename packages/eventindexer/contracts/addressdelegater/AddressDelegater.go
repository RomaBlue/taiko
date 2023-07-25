// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addressdelegater

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

// AddressDelegaterMetaData contains all meta data concerning the AddressDelegater contract.
var AddressDelegaterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"acceptProposerDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"acceptProverDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"delegateProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"delegateProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"}],\"name\":\"proposerToDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"}],\"name\":\"proverToDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"usedDelegateProposerAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"usedDelegateProverAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AddressDelegaterABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressDelegaterMetaData.ABI instead.
var AddressDelegaterABI = AddressDelegaterMetaData.ABI

// AddressDelegater is an auto generated Go binding around an Ethereum contract.
type AddressDelegater struct {
	AddressDelegaterCaller     // Read-only binding to the contract
	AddressDelegaterTransactor // Write-only binding to the contract
	AddressDelegaterFilterer   // Log filterer for contract events
}

// AddressDelegaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressDelegaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressDelegaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressDelegaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressDelegaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressDelegaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressDelegaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressDelegaterSession struct {
	Contract     *AddressDelegater // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressDelegaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressDelegaterCallerSession struct {
	Contract *AddressDelegaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AddressDelegaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressDelegaterTransactorSession struct {
	Contract     *AddressDelegaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AddressDelegaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressDelegaterRaw struct {
	Contract *AddressDelegater // Generic contract binding to access the raw methods on
}

// AddressDelegaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressDelegaterCallerRaw struct {
	Contract *AddressDelegaterCaller // Generic read-only contract binding to access the raw methods on
}

// AddressDelegaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressDelegaterTransactorRaw struct {
	Contract *AddressDelegaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressDelegater creates a new instance of AddressDelegater, bound to a specific deployed contract.
func NewAddressDelegater(address common.Address, backend bind.ContractBackend) (*AddressDelegater, error) {
	contract, err := bindAddressDelegater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressDelegater{AddressDelegaterCaller: AddressDelegaterCaller{contract: contract}, AddressDelegaterTransactor: AddressDelegaterTransactor{contract: contract}, AddressDelegaterFilterer: AddressDelegaterFilterer{contract: contract}}, nil
}

// NewAddressDelegaterCaller creates a new read-only instance of AddressDelegater, bound to a specific deployed contract.
func NewAddressDelegaterCaller(address common.Address, caller bind.ContractCaller) (*AddressDelegaterCaller, error) {
	contract, err := bindAddressDelegater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressDelegaterCaller{contract: contract}, nil
}

// NewAddressDelegaterTransactor creates a new write-only instance of AddressDelegater, bound to a specific deployed contract.
func NewAddressDelegaterTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressDelegaterTransactor, error) {
	contract, err := bindAddressDelegater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressDelegaterTransactor{contract: contract}, nil
}

// NewAddressDelegaterFilterer creates a new log filterer instance of AddressDelegater, bound to a specific deployed contract.
func NewAddressDelegaterFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressDelegaterFilterer, error) {
	contract, err := bindAddressDelegater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressDelegaterFilterer{contract: contract}, nil
}

// bindAddressDelegater binds a generic wrapper to an already deployed contract.
func bindAddressDelegater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressDelegaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressDelegater *AddressDelegaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressDelegater.Contract.AddressDelegaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressDelegater *AddressDelegaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressDelegater.Contract.AddressDelegaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressDelegater *AddressDelegaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressDelegater.Contract.AddressDelegaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressDelegater *AddressDelegaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressDelegater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressDelegater *AddressDelegaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressDelegater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressDelegater *AddressDelegaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressDelegater.Contract.contract.Transact(opts, method, params...)
}

// ProposerToDelegate is a free data retrieval call binding the contract method 0xcc3043cd.
//
// Solidity: function proposerToDelegate(address delegater) view returns(bool approved, address delegate)
func (_AddressDelegater *AddressDelegaterCaller) ProposerToDelegate(opts *bind.CallOpts, delegater common.Address) (struct {
	Approved bool
	Delegate common.Address
}, error) {
	var out []interface{}
	err := _AddressDelegater.contract.Call(opts, &out, "proposerToDelegate", delegater)

	outstruct := new(struct {
		Approved bool
		Delegate common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Approved = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ProposerToDelegate is a free data retrieval call binding the contract method 0xcc3043cd.
//
// Solidity: function proposerToDelegate(address delegater) view returns(bool approved, address delegate)
func (_AddressDelegater *AddressDelegaterSession) ProposerToDelegate(delegater common.Address) (struct {
	Approved bool
	Delegate common.Address
}, error) {
	return _AddressDelegater.Contract.ProposerToDelegate(&_AddressDelegater.CallOpts, delegater)
}

// ProposerToDelegate is a free data retrieval call binding the contract method 0xcc3043cd.
//
// Solidity: function proposerToDelegate(address delegater) view returns(bool approved, address delegate)
func (_AddressDelegater *AddressDelegaterCallerSession) ProposerToDelegate(delegater common.Address) (struct {
	Approved bool
	Delegate common.Address
}, error) {
	return _AddressDelegater.Contract.ProposerToDelegate(&_AddressDelegater.CallOpts, delegater)
}

// ProverToDelegate is a free data retrieval call binding the contract method 0xa82437b3.
//
// Solidity: function proverToDelegate(address delegater) view returns(bool approved, address delegate)
func (_AddressDelegater *AddressDelegaterCaller) ProverToDelegate(opts *bind.CallOpts, delegater common.Address) (struct {
	Approved bool
	Delegate common.Address
}, error) {
	var out []interface{}
	err := _AddressDelegater.contract.Call(opts, &out, "proverToDelegate", delegater)

	outstruct := new(struct {
		Approved bool
		Delegate common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Approved = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ProverToDelegate is a free data retrieval call binding the contract method 0xa82437b3.
//
// Solidity: function proverToDelegate(address delegater) view returns(bool approved, address delegate)
func (_AddressDelegater *AddressDelegaterSession) ProverToDelegate(delegater common.Address) (struct {
	Approved bool
	Delegate common.Address
}, error) {
	return _AddressDelegater.Contract.ProverToDelegate(&_AddressDelegater.CallOpts, delegater)
}

// ProverToDelegate is a free data retrieval call binding the contract method 0xa82437b3.
//
// Solidity: function proverToDelegate(address delegater) view returns(bool approved, address delegate)
func (_AddressDelegater *AddressDelegaterCallerSession) ProverToDelegate(delegater common.Address) (struct {
	Approved bool
	Delegate common.Address
}, error) {
	return _AddressDelegater.Contract.ProverToDelegate(&_AddressDelegater.CallOpts, delegater)
}

// UsedDelegateProposerAddresses is a free data retrieval call binding the contract method 0x563bb512.
//
// Solidity: function usedDelegateProposerAddresses(address delegate) view returns(bool used)
func (_AddressDelegater *AddressDelegaterCaller) UsedDelegateProposerAddresses(opts *bind.CallOpts, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _AddressDelegater.contract.Call(opts, &out, "usedDelegateProposerAddresses", delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedDelegateProposerAddresses is a free data retrieval call binding the contract method 0x563bb512.
//
// Solidity: function usedDelegateProposerAddresses(address delegate) view returns(bool used)
func (_AddressDelegater *AddressDelegaterSession) UsedDelegateProposerAddresses(delegate common.Address) (bool, error) {
	return _AddressDelegater.Contract.UsedDelegateProposerAddresses(&_AddressDelegater.CallOpts, delegate)
}

// UsedDelegateProposerAddresses is a free data retrieval call binding the contract method 0x563bb512.
//
// Solidity: function usedDelegateProposerAddresses(address delegate) view returns(bool used)
func (_AddressDelegater *AddressDelegaterCallerSession) UsedDelegateProposerAddresses(delegate common.Address) (bool, error) {
	return _AddressDelegater.Contract.UsedDelegateProposerAddresses(&_AddressDelegater.CallOpts, delegate)
}

// UsedDelegateProverAddresses is a free data retrieval call binding the contract method 0xa2531733.
//
// Solidity: function usedDelegateProverAddresses(address delegate) view returns(bool used)
func (_AddressDelegater *AddressDelegaterCaller) UsedDelegateProverAddresses(opts *bind.CallOpts, delegate common.Address) (bool, error) {
	var out []interface{}
	err := _AddressDelegater.contract.Call(opts, &out, "usedDelegateProverAddresses", delegate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedDelegateProverAddresses is a free data retrieval call binding the contract method 0xa2531733.
//
// Solidity: function usedDelegateProverAddresses(address delegate) view returns(bool used)
func (_AddressDelegater *AddressDelegaterSession) UsedDelegateProverAddresses(delegate common.Address) (bool, error) {
	return _AddressDelegater.Contract.UsedDelegateProverAddresses(&_AddressDelegater.CallOpts, delegate)
}

// UsedDelegateProverAddresses is a free data retrieval call binding the contract method 0xa2531733.
//
// Solidity: function usedDelegateProverAddresses(address delegate) view returns(bool used)
func (_AddressDelegater *AddressDelegaterCallerSession) UsedDelegateProverAddresses(delegate common.Address) (bool, error) {
	return _AddressDelegater.Contract.UsedDelegateProverAddresses(&_AddressDelegater.CallOpts, delegate)
}

// AcceptProposerDelegate is a paid mutator transaction binding the contract method 0x1d9fae35.
//
// Solidity: function acceptProposerDelegate(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactor) AcceptProposerDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.contract.Transact(opts, "acceptProposerDelegate", delegate)
}

// AcceptProposerDelegate is a paid mutator transaction binding the contract method 0x1d9fae35.
//
// Solidity: function acceptProposerDelegate(address delegate) returns()
func (_AddressDelegater *AddressDelegaterSession) AcceptProposerDelegate(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.AcceptProposerDelegate(&_AddressDelegater.TransactOpts, delegate)
}

// AcceptProposerDelegate is a paid mutator transaction binding the contract method 0x1d9fae35.
//
// Solidity: function acceptProposerDelegate(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactorSession) AcceptProposerDelegate(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.AcceptProposerDelegate(&_AddressDelegater.TransactOpts, delegate)
}

// AcceptProverDelegate is a paid mutator transaction binding the contract method 0x978c8239.
//
// Solidity: function acceptProverDelegate(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactor) AcceptProverDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.contract.Transact(opts, "acceptProverDelegate", delegate)
}

// AcceptProverDelegate is a paid mutator transaction binding the contract method 0x978c8239.
//
// Solidity: function acceptProverDelegate(address delegate) returns()
func (_AddressDelegater *AddressDelegaterSession) AcceptProverDelegate(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.AcceptProverDelegate(&_AddressDelegater.TransactOpts, delegate)
}

// AcceptProverDelegate is a paid mutator transaction binding the contract method 0x978c8239.
//
// Solidity: function acceptProverDelegate(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactorSession) AcceptProverDelegate(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.AcceptProverDelegate(&_AddressDelegater.TransactOpts, delegate)
}

// DelegateProposer is a paid mutator transaction binding the contract method 0xbf4e0fca.
//
// Solidity: function delegateProposer(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactor) DelegateProposer(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.contract.Transact(opts, "delegateProposer", delegate)
}

// DelegateProposer is a paid mutator transaction binding the contract method 0xbf4e0fca.
//
// Solidity: function delegateProposer(address delegate) returns()
func (_AddressDelegater *AddressDelegaterSession) DelegateProposer(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.DelegateProposer(&_AddressDelegater.TransactOpts, delegate)
}

// DelegateProposer is a paid mutator transaction binding the contract method 0xbf4e0fca.
//
// Solidity: function delegateProposer(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactorSession) DelegateProposer(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.DelegateProposer(&_AddressDelegater.TransactOpts, delegate)
}

// DelegateProver is a paid mutator transaction binding the contract method 0x2ed6a001.
//
// Solidity: function delegateProver(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactor) DelegateProver(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.contract.Transact(opts, "delegateProver", delegate)
}

// DelegateProver is a paid mutator transaction binding the contract method 0x2ed6a001.
//
// Solidity: function delegateProver(address delegate) returns()
func (_AddressDelegater *AddressDelegaterSession) DelegateProver(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.DelegateProver(&_AddressDelegater.TransactOpts, delegate)
}

// DelegateProver is a paid mutator transaction binding the contract method 0x2ed6a001.
//
// Solidity: function delegateProver(address delegate) returns()
func (_AddressDelegater *AddressDelegaterTransactorSession) DelegateProver(delegate common.Address) (*types.Transaction, error) {
	return _AddressDelegater.Contract.DelegateProver(&_AddressDelegater.TransactOpts, delegate)
}
