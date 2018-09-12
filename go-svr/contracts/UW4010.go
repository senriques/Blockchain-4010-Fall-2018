// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UW4010

import (
	"math/big"
	"strings"
)

// UW4010ABI is the input ABI used to generate the binding from.
const UW4010ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensBurnt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawStudent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_pin\",\"type\":\"uint256\"}],\"name\":\"validLogin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_pin\",\"type\":\"uint256\"}],\"name\":\"newStudent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"haveCouponsStudent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStudent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStudentAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNStudent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Assignment\",\"type\":\"uint256\"}],\"name\":\"redeamToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance1\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_pin\",\"type\":\"uint256\"}],\"name\":\"updatePin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"haveCoupons\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"address\"}],\"name\":\"convToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_pin\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"newStudentCoupons\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"Student\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Assignment\",\"type\":\"uint256\"}],\"name\":\"CouponUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"Student\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Success\",\"type\":\"bool\"}],\"name\":\"PinUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// UW4010 is an auto generated Go binding around an Ethereum contract.
type UW4010 struct {
	UW4010Caller     // Read-only binding to the contract
	UW4010Transactor // Write-only binding to the contract
	UW4010Filterer   // Log filterer for contract events
}

// UW4010Caller is an auto generated read-only Go binding around an Ethereum contract.
type UW4010Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UW4010Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UW4010Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UW4010Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UW4010Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UW4010Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UW4010Session struct {
	Contract     *UW4010           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UW4010CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UW4010CallerSession struct {
	Contract *UW4010Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UW4010TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UW4010TransactorSession struct {
	Contract     *UW4010Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UW4010Raw is an auto generated low-level Go binding around an Ethereum contract.
type UW4010Raw struct {
	Contract *UW4010 // Generic contract binding to access the raw methods on
}

// UW4010CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UW4010CallerRaw struct {
	Contract *UW4010Caller // Generic read-only contract binding to access the raw methods on
}

// UW4010TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UW4010TransactorRaw struct {
	Contract *UW4010Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUW4010 creates a new instance of UW4010, bound to a specific deployed contract.
func NewUW4010(address common.Address, backend bind.ContractBackend) (*UW4010, error) {
	contract, err := bindUW4010(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UW4010{UW4010Caller: UW4010Caller{contract: contract}, UW4010Transactor: UW4010Transactor{contract: contract}, UW4010Filterer: UW4010Filterer{contract: contract}}, nil
}

// NewUW4010Caller creates a new read-only instance of UW4010, bound to a specific deployed contract.
func NewUW4010Caller(address common.Address, caller bind.ContractCaller) (*UW4010Caller, error) {
	contract, err := bindUW4010(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UW4010Caller{contract: contract}, nil
}

// NewUW4010Transactor creates a new write-only instance of UW4010, bound to a specific deployed contract.
func NewUW4010Transactor(address common.Address, transactor bind.ContractTransactor) (*UW4010Transactor, error) {
	contract, err := bindUW4010(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UW4010Transactor{contract: contract}, nil
}

// NewUW4010Filterer creates a new log filterer instance of UW4010, bound to a specific deployed contract.
func NewUW4010Filterer(address common.Address, filterer bind.ContractFilterer) (*UW4010Filterer, error) {
	contract, err := bindUW4010(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UW4010Filterer{contract: contract}, nil
}

// bindUW4010 binds a generic wrapper to an already deployed contract.
func bindUW4010(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UW4010ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UW4010 *UW4010Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UW4010.Contract.UW4010Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UW4010 *UW4010Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UW4010.Contract.UW4010Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UW4010 *UW4010Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UW4010.Contract.UW4010Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UW4010 *UW4010CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UW4010.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UW4010 *UW4010TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UW4010.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UW4010 *UW4010TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UW4010.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_UW4010 *UW4010Caller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_UW4010 *UW4010Session) INITIALSUPPLY() (*big.Int, error) {
	return _UW4010.Contract.INITIALSUPPLY(&_UW4010.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_UW4010 *UW4010CallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _UW4010.Contract.INITIALSUPPLY(&_UW4010.CallOpts)
}

// Balance1 is a free data retrieval call binding the contract method 0xaf62db3f.
//
// Solidity: function balance1( address) constant returns(uint256)
func (_UW4010 *UW4010Caller) Balance1(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "balance1", arg0)
	return *ret0, err
}

// Balance1 is a free data retrieval call binding the contract method 0xaf62db3f.
//
// Solidity: function balance1( address) constant returns(uint256)
func (_UW4010 *UW4010Session) Balance1(arg0 common.Address) (*big.Int, error) {
	return _UW4010.Contract.Balance1(&_UW4010.CallOpts, arg0)
}

// Balance1 is a free data retrieval call binding the contract method 0xaf62db3f.
//
// Solidity: function balance1( address) constant returns(uint256)
func (_UW4010 *UW4010CallerSession) Balance1(arg0 common.Address) (*big.Int, error) {
	return _UW4010.Contract.Balance1(&_UW4010.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_UW4010 *UW4010Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_UW4010 *UW4010Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _UW4010.Contract.BalanceOf(&_UW4010.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_UW4010 *UW4010CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _UW4010.Contract.BalanceOf(&_UW4010.CallOpts, _owner)
}

// ConvToAddress is a free data retrieval call binding the contract method 0xe3839203.
//
// Solidity: function convToAddress(_x address) constant returns(address)
func (_UW4010 *UW4010Caller) ConvToAddress(opts *bind.CallOpts, _x common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "convToAddress", _x)
	return *ret0, err
}

// ConvToAddress is a free data retrieval call binding the contract method 0xe3839203.
//
// Solidity: function convToAddress(_x address) constant returns(address)
func (_UW4010 *UW4010Session) ConvToAddress(_x common.Address) (common.Address, error) {
	return _UW4010.Contract.ConvToAddress(&_UW4010.CallOpts, _x)
}

// ConvToAddress is a free data retrieval call binding the contract method 0xe3839203.
//
// Solidity: function convToAddress(_x address) constant returns(address)
func (_UW4010 *UW4010CallerSession) ConvToAddress(_x common.Address) (common.Address, error) {
	return _UW4010.Contract.ConvToAddress(&_UW4010.CallOpts, _x)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_UW4010 *UW4010Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_UW4010 *UW4010Session) Decimals() (uint8, error) {
	return _UW4010.Contract.Decimals(&_UW4010.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_UW4010 *UW4010CallerSession) Decimals() (uint8, error) {
	return _UW4010.Contract.Decimals(&_UW4010.CallOpts)
}

// GetNStudent is a free data retrieval call binding the contract method 0x8d717ad2.
//
// Solidity: function getNStudent() constant returns(uint256)
func (_UW4010 *UW4010Caller) GetNStudent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "getNStudent")
	return *ret0, err
}

// GetNStudent is a free data retrieval call binding the contract method 0x8d717ad2.
//
// Solidity: function getNStudent() constant returns(uint256)
func (_UW4010 *UW4010Session) GetNStudent() (*big.Int, error) {
	return _UW4010.Contract.GetNStudent(&_UW4010.CallOpts)
}

// GetNStudent is a free data retrieval call binding the contract method 0x8d717ad2.
//
// Solidity: function getNStudent() constant returns(uint256)
func (_UW4010 *UW4010CallerSession) GetNStudent() (*big.Int, error) {
	return _UW4010.Contract.GetNStudent(&_UW4010.CallOpts)
}

// GetStudentAddr is a free data retrieval call binding the contract method 0x8637161e.
//
// Solidity: function getStudentAddr(index uint256) constant returns(address)
func (_UW4010 *UW4010Caller) GetStudentAddr(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "getStudentAddr", index)
	return *ret0, err
}

// GetStudentAddr is a free data retrieval call binding the contract method 0x8637161e.
//
// Solidity: function getStudentAddr(index uint256) constant returns(address)
func (_UW4010 *UW4010Session) GetStudentAddr(index *big.Int) (common.Address, error) {
	return _UW4010.Contract.GetStudentAddr(&_UW4010.CallOpts, index)
}

// GetStudentAddr is a free data retrieval call binding the contract method 0x8637161e.
//
// Solidity: function getStudentAddr(index uint256) constant returns(address)
func (_UW4010 *UW4010CallerSession) GetStudentAddr(index *big.Int) (common.Address, error) {
	return _UW4010.Contract.GetStudentAddr(&_UW4010.CallOpts, index)
}

// HaveCoupons is a free data retrieval call binding the contract method 0xe2fe13de.
//
// Solidity: function haveCoupons() constant returns(uint256)
func (_UW4010 *UW4010Caller) HaveCoupons(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "haveCoupons")
	return *ret0, err
}

// HaveCoupons is a free data retrieval call binding the contract method 0xe2fe13de.
//
// Solidity: function haveCoupons() constant returns(uint256)
func (_UW4010 *UW4010Session) HaveCoupons() (*big.Int, error) {
	return _UW4010.Contract.HaveCoupons(&_UW4010.CallOpts)
}

// HaveCoupons is a free data retrieval call binding the contract method 0xe2fe13de.
//
// Solidity: function haveCoupons() constant returns(uint256)
func (_UW4010 *UW4010CallerSession) HaveCoupons() (*big.Int, error) {
	return _UW4010.Contract.HaveCoupons(&_UW4010.CallOpts)
}

// HaveCouponsStudent is a free data retrieval call binding the contract method 0x60b9aa61.
//
// Solidity: function haveCouponsStudent(_to address) constant returns(uint256)
func (_UW4010 *UW4010Caller) HaveCouponsStudent(opts *bind.CallOpts, _to common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "haveCouponsStudent", _to)
	return *ret0, err
}

// HaveCouponsStudent is a free data retrieval call binding the contract method 0x60b9aa61.
//
// Solidity: function haveCouponsStudent(_to address) constant returns(uint256)
func (_UW4010 *UW4010Session) HaveCouponsStudent(_to common.Address) (*big.Int, error) {
	return _UW4010.Contract.HaveCouponsStudent(&_UW4010.CallOpts, _to)
}

// HaveCouponsStudent is a free data retrieval call binding the contract method 0x60b9aa61.
//
// Solidity: function haveCouponsStudent(_to address) constant returns(uint256)
func (_UW4010 *UW4010CallerSession) HaveCouponsStudent(_to common.Address) (*big.Int, error) {
	return _UW4010.Contract.HaveCouponsStudent(&_UW4010.CallOpts, _to)
}

// IsStudent is a free data retrieval call binding the contract method 0x82f7d392.
//
// Solidity: function isStudent( address) constant returns(uint256)
func (_UW4010 *UW4010Caller) IsStudent(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "isStudent", arg0)
	return *ret0, err
}

// IsStudent is a free data retrieval call binding the contract method 0x82f7d392.
//
// Solidity: function isStudent( address) constant returns(uint256)
func (_UW4010 *UW4010Session) IsStudent(arg0 common.Address) (*big.Int, error) {
	return _UW4010.Contract.IsStudent(&_UW4010.CallOpts, arg0)
}

// IsStudent is a free data retrieval call binding the contract method 0x82f7d392.
//
// Solidity: function isStudent( address) constant returns(uint256)
func (_UW4010 *UW4010CallerSession) IsStudent(arg0 common.Address) (*big.Int, error) {
	return _UW4010.Contract.IsStudent(&_UW4010.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_UW4010 *UW4010Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_UW4010 *UW4010Session) Name() (string, error) {
	return _UW4010.Contract.Name(&_UW4010.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_UW4010 *UW4010CallerSession) Name() (string, error) {
	return _UW4010.Contract.Name(&_UW4010.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UW4010 *UW4010Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UW4010 *UW4010Session) Owner() (common.Address, error) {
	return _UW4010.Contract.Owner(&_UW4010.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UW4010 *UW4010CallerSession) Owner() (common.Address, error) {
	return _UW4010.Contract.Owner(&_UW4010.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_UW4010 *UW4010Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_UW4010 *UW4010Session) Symbol() (string, error) {
	return _UW4010.Contract.Symbol(&_UW4010.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_UW4010 *UW4010CallerSession) Symbol() (string, error) {
	return _UW4010.Contract.Symbol(&_UW4010.CallOpts)
}

// TokensBurnt is a free data retrieval call binding the contract method 0x0614612b.
//
// Solidity: function tokensBurnt( address) constant returns(uint256)
func (_UW4010 *UW4010Caller) TokensBurnt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "tokensBurnt", arg0)
	return *ret0, err
}

// TokensBurnt is a free data retrieval call binding the contract method 0x0614612b.
//
// Solidity: function tokensBurnt( address) constant returns(uint256)
func (_UW4010 *UW4010Session) TokensBurnt(arg0 common.Address) (*big.Int, error) {
	return _UW4010.Contract.TokensBurnt(&_UW4010.CallOpts, arg0)
}

// TokensBurnt is a free data retrieval call binding the contract method 0x0614612b.
//
// Solidity: function tokensBurnt( address) constant returns(uint256)
func (_UW4010 *UW4010CallerSession) TokensBurnt(arg0 common.Address) (*big.Int, error) {
	return _UW4010.Contract.TokensBurnt(&_UW4010.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_UW4010 *UW4010Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_UW4010 *UW4010Session) TotalSupply() (*big.Int, error) {
	return _UW4010.Contract.TotalSupply(&_UW4010.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_UW4010 *UW4010CallerSession) TotalSupply() (*big.Int, error) {
	return _UW4010.Contract.TotalSupply(&_UW4010.CallOpts)
}

// ValidLogin is a free data retrieval call binding the contract method 0x0ea98f09.
//
// Solidity: function validLogin(_to address, _pin uint256) constant returns(bool)
func (_UW4010 *UW4010Caller) ValidLogin(opts *bind.CallOpts, _to common.Address, _pin *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UW4010.contract.Call(opts, out, "validLogin", _to, _pin)
	return *ret0, err
}

// ValidLogin is a free data retrieval call binding the contract method 0x0ea98f09.
//
// Solidity: function validLogin(_to address, _pin uint256) constant returns(bool)
func (_UW4010 *UW4010Session) ValidLogin(_to common.Address, _pin *big.Int) (bool, error) {
	return _UW4010.Contract.ValidLogin(&_UW4010.CallOpts, _to, _pin)
}

// ValidLogin is a free data retrieval call binding the contract method 0x0ea98f09.
//
// Solidity: function validLogin(_to address, _pin uint256) constant returns(bool)
func (_UW4010 *UW4010CallerSession) ValidLogin(_to common.Address, _pin *big.Int) (bool, error) {
	return _UW4010.Contract.ValidLogin(&_UW4010.CallOpts, _to, _pin)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_UW4010 *UW4010Transactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_UW4010 *UW4010Session) Burn(_value *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.Burn(&_UW4010.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_UW4010 *UW4010TransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.Burn(&_UW4010.TransactOpts, _value)
}

// NewStudent is a paid mutator transaction binding the contract method 0x34b19284.
//
// Solidity: function newStudent(_to address, _pin uint256) returns()
func (_UW4010 *UW4010Transactor) NewStudent(opts *bind.TransactOpts, _to common.Address, _pin *big.Int) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "newStudent", _to, _pin)
}

// NewStudent is a paid mutator transaction binding the contract method 0x34b19284.
//
// Solidity: function newStudent(_to address, _pin uint256) returns()
func (_UW4010 *UW4010Session) NewStudent(_to common.Address, _pin *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.NewStudent(&_UW4010.TransactOpts, _to, _pin)
}

// NewStudent is a paid mutator transaction binding the contract method 0x34b19284.
//
// Solidity: function newStudent(_to address, _pin uint256) returns()
func (_UW4010 *UW4010TransactorSession) NewStudent(_to common.Address, _pin *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.NewStudent(&_UW4010.TransactOpts, _to, _pin)
}

// NewStudentCoupons is a paid mutator transaction binding the contract method 0xf078326e.
//
// Solidity: function newStudentCoupons(_to address, _pin uint256, _value uint256) returns()
func (_UW4010 *UW4010Transactor) NewStudentCoupons(opts *bind.TransactOpts, _to common.Address, _pin *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "newStudentCoupons", _to, _pin, _value)
}

// NewStudentCoupons is a paid mutator transaction binding the contract method 0xf078326e.
//
// Solidity: function newStudentCoupons(_to address, _pin uint256, _value uint256) returns()
func (_UW4010 *UW4010Session) NewStudentCoupons(_to common.Address, _pin *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.NewStudentCoupons(&_UW4010.TransactOpts, _to, _pin, _value)
}

// NewStudentCoupons is a paid mutator transaction binding the contract method 0xf078326e.
//
// Solidity: function newStudentCoupons(_to address, _pin uint256, _value uint256) returns()
func (_UW4010 *UW4010TransactorSession) NewStudentCoupons(_to common.Address, _pin *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.NewStudentCoupons(&_UW4010.TransactOpts, _to, _pin, _value)
}

// RedeamToken is a paid mutator transaction binding the contract method 0xad2d544e.
//
// Solidity: function redeamToken(Assignment uint256) returns()
func (_UW4010 *UW4010Transactor) RedeamToken(opts *bind.TransactOpts, Assignment *big.Int) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "redeamToken", Assignment)
}

// RedeamToken is a paid mutator transaction binding the contract method 0xad2d544e.
//
// Solidity: function redeamToken(Assignment uint256) returns()
func (_UW4010 *UW4010Session) RedeamToken(Assignment *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.RedeamToken(&_UW4010.TransactOpts, Assignment)
}

// RedeamToken is a paid mutator transaction binding the contract method 0xad2d544e.
//
// Solidity: function redeamToken(Assignment uint256) returns()
func (_UW4010 *UW4010TransactorSession) RedeamToken(Assignment *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.RedeamToken(&_UW4010.TransactOpts, Assignment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UW4010 *UW4010Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UW4010 *UW4010Session) RenounceOwnership() (*types.Transaction, error) {
	return _UW4010.Contract.RenounceOwnership(&_UW4010.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UW4010 *UW4010TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UW4010.Contract.RenounceOwnership(&_UW4010.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_UW4010 *UW4010Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_UW4010 *UW4010Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.Transfer(&_UW4010.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_UW4010 *UW4010TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.Transfer(&_UW4010.TransactOpts, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_UW4010 *UW4010Transactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_UW4010 *UW4010Session) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _UW4010.Contract.TransferOwnership(&_UW4010.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_UW4010 *UW4010TransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _UW4010.Contract.TransferOwnership(&_UW4010.TransactOpts, _newOwner)
}

// UpdatePin is a paid mutator transaction binding the contract method 0xc7089d2c.
//
// Solidity: function updatePin(_to address, _pin uint256) returns(bool)
func (_UW4010 *UW4010Transactor) UpdatePin(opts *bind.TransactOpts, _to common.Address, _pin *big.Int) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "updatePin", _to, _pin)
}

// UpdatePin is a paid mutator transaction binding the contract method 0xc7089d2c.
//
// Solidity: function updatePin(_to address, _pin uint256) returns(bool)
func (_UW4010 *UW4010Session) UpdatePin(_to common.Address, _pin *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.UpdatePin(&_UW4010.TransactOpts, _to, _pin)
}

// UpdatePin is a paid mutator transaction binding the contract method 0xc7089d2c.
//
// Solidity: function updatePin(_to address, _pin uint256) returns(bool)
func (_UW4010 *UW4010TransactorSession) UpdatePin(_to common.Address, _pin *big.Int) (*types.Transaction, error) {
	return _UW4010.Contract.UpdatePin(&_UW4010.TransactOpts, _to, _pin)
}

// WithdrawStudent is a paid mutator transaction binding the contract method 0x0cda463f.
//
// Solidity: function withdrawStudent(_to address) returns()
func (_UW4010 *UW4010Transactor) WithdrawStudent(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _UW4010.contract.Transact(opts, "withdrawStudent", _to)
}

// WithdrawStudent is a paid mutator transaction binding the contract method 0x0cda463f.
//
// Solidity: function withdrawStudent(_to address) returns()
func (_UW4010 *UW4010Session) WithdrawStudent(_to common.Address) (*types.Transaction, error) {
	return _UW4010.Contract.WithdrawStudent(&_UW4010.TransactOpts, _to)
}

// WithdrawStudent is a paid mutator transaction binding the contract method 0x0cda463f.
//
// Solidity: function withdrawStudent(_to address) returns()
func (_UW4010 *UW4010TransactorSession) WithdrawStudent(_to common.Address) (*types.Transaction, error) {
	return _UW4010.Contract.WithdrawStudent(&_UW4010.TransactOpts, _to)
}

// UW4010BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the UW4010 contract.
type UW4010BurnIterator struct {
	Event *UW4010Burn // Event containing the contract specifics and raw log

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
func (it *UW4010BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UW4010Burn)
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
		it.Event = new(UW4010Burn)
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
func (it *UW4010BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UW4010BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UW4010Burn represents a Burn event raised by the UW4010 contract.
type UW4010Burn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_UW4010 *UW4010Filterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*UW4010BurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _UW4010.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &UW4010BurnIterator{contract: _UW4010.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_UW4010 *UW4010Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *UW4010Burn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _UW4010.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UW4010Burn)
				if err := _UW4010.contract.UnpackLog(event, "Burn", log); err != nil {
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

// UW4010CouponUsedIterator is returned from FilterCouponUsed and is used to iterate over the raw logs and unpacked data for CouponUsed events raised by the UW4010 contract.
type UW4010CouponUsedIterator struct {
	Event *UW4010CouponUsed // Event containing the contract specifics and raw log

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
func (it *UW4010CouponUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UW4010CouponUsed)
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
		it.Event = new(UW4010CouponUsed)
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
func (it *UW4010CouponUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UW4010CouponUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UW4010CouponUsed represents a CouponUsed event raised by the UW4010 contract.
type UW4010CouponUsed struct {
	Student    common.Address
	Assignment *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCouponUsed is a free log retrieval operation binding the contract event 0x5915d66237d65e9a887dbb51c150a4c9005e64cfa29d123764158f825f3c0884.
//
// Solidity: e CouponUsed(Student address, Assignment uint256)
func (_UW4010 *UW4010Filterer) FilterCouponUsed(opts *bind.FilterOpts) (*UW4010CouponUsedIterator, error) {

	logs, sub, err := _UW4010.contract.FilterLogs(opts, "CouponUsed")
	if err != nil {
		return nil, err
	}
	return &UW4010CouponUsedIterator{contract: _UW4010.contract, event: "CouponUsed", logs: logs, sub: sub}, nil
}

// WatchCouponUsed is a free log subscription operation binding the contract event 0x5915d66237d65e9a887dbb51c150a4c9005e64cfa29d123764158f825f3c0884.
//
// Solidity: e CouponUsed(Student address, Assignment uint256)
func (_UW4010 *UW4010Filterer) WatchCouponUsed(opts *bind.WatchOpts, sink chan<- *UW4010CouponUsed) (event.Subscription, error) {

	logs, sub, err := _UW4010.contract.WatchLogs(opts, "CouponUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UW4010CouponUsed)
				if err := _UW4010.contract.UnpackLog(event, "CouponUsed", log); err != nil {
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

// UW4010OwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the UW4010 contract.
type UW4010OwnershipRenouncedIterator struct {
	Event *UW4010OwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *UW4010OwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UW4010OwnershipRenounced)
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
		it.Event = new(UW4010OwnershipRenounced)
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
func (it *UW4010OwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UW4010OwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UW4010OwnershipRenounced represents a OwnershipRenounced event raised by the UW4010 contract.
type UW4010OwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_UW4010 *UW4010Filterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*UW4010OwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _UW4010.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UW4010OwnershipRenouncedIterator{contract: _UW4010.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_UW4010 *UW4010Filterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *UW4010OwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _UW4010.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UW4010OwnershipRenounced)
				if err := _UW4010.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// UW4010OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UW4010 contract.
type UW4010OwnershipTransferredIterator struct {
	Event *UW4010OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UW4010OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UW4010OwnershipTransferred)
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
		it.Event = new(UW4010OwnershipTransferred)
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
func (it *UW4010OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UW4010OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UW4010OwnershipTransferred represents a OwnershipTransferred event raised by the UW4010 contract.
type UW4010OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_UW4010 *UW4010Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UW4010OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UW4010.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UW4010OwnershipTransferredIterator{contract: _UW4010.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_UW4010 *UW4010Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UW4010OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UW4010.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UW4010OwnershipTransferred)
				if err := _UW4010.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// UW4010PinUpdatedIterator is returned from FilterPinUpdated and is used to iterate over the raw logs and unpacked data for PinUpdated events raised by the UW4010 contract.
type UW4010PinUpdatedIterator struct {
	Event *UW4010PinUpdated // Event containing the contract specifics and raw log

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
func (it *UW4010PinUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UW4010PinUpdated)
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
		it.Event = new(UW4010PinUpdated)
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
func (it *UW4010PinUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UW4010PinUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UW4010PinUpdated represents a PinUpdated event raised by the UW4010 contract.
type UW4010PinUpdated struct {
	Student common.Address
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPinUpdated is a free log retrieval operation binding the contract event 0xa0de670ec5c9472afc3eec965bffdc4239f0c3711c7f8cbbe75fe5a4df0c96ba.
//
// Solidity: e PinUpdated(Student address, Success bool)
func (_UW4010 *UW4010Filterer) FilterPinUpdated(opts *bind.FilterOpts) (*UW4010PinUpdatedIterator, error) {

	logs, sub, err := _UW4010.contract.FilterLogs(opts, "PinUpdated")
	if err != nil {
		return nil, err
	}
	return &UW4010PinUpdatedIterator{contract: _UW4010.contract, event: "PinUpdated", logs: logs, sub: sub}, nil
}

// WatchPinUpdated is a free log subscription operation binding the contract event 0xa0de670ec5c9472afc3eec965bffdc4239f0c3711c7f8cbbe75fe5a4df0c96ba.
//
// Solidity: e PinUpdated(Student address, Success bool)
func (_UW4010 *UW4010Filterer) WatchPinUpdated(opts *bind.WatchOpts, sink chan<- *UW4010PinUpdated) (event.Subscription, error) {

	logs, sub, err := _UW4010.contract.WatchLogs(opts, "PinUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UW4010PinUpdated)
				if err := _UW4010.contract.UnpackLog(event, "PinUpdated", log); err != nil {
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

// UW4010TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UW4010 contract.
type UW4010TransferIterator struct {
	Event *UW4010Transfer // Event containing the contract specifics and raw log

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
func (it *UW4010TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UW4010Transfer)
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
		it.Event = new(UW4010Transfer)
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
func (it *UW4010TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UW4010TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UW4010Transfer represents a Transfer event raised by the UW4010 contract.
type UW4010Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_UW4010 *UW4010Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UW4010TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UW4010.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UW4010TransferIterator{contract: _UW4010.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_UW4010 *UW4010Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UW4010Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UW4010.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UW4010Transfer)
				if err := _UW4010.contract.UnpackLog(event, "Transfer", log); err != nil {
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
