// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package name_service

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// NameServiceABI is the input ABI used to generate the binding from.
const NameServiceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"NameSet\",\"type\":\"event\"}]"

// NameServiceBin is the compiled bytecode used for deploying new contracts.
const NameServiceBin = `6060604052341561000f57600080fd5b61036d8061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317d7de7c8114610050578063c47f0027146100da575b600080fd5b341561005b57600080fd5b61006361012d565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561009f578082015183820152602001610087565b50505050905090810190601f1680156100cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100e557600080fd5b61012b60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506101d695505050505050565b005b610135610297565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101cb5780601f106101a0576101008083540402835291602001916101cb565b820191906000526020600020905b8154815290600101906020018083116101ae57829003601f168201915b505050505090505b90565b60008180516101e99291602001906102a9565b507f13c98778b0c1a086bb98d7f1986e15788b5d3a1ad4c492e1d78f1c4cc51c20cf6000604051602080825282546002600019610100600184161502019091160490820181905281906040820190849080156102865780601f1061025b57610100808354040283529160200191610286565b820191906000526020600020905b81548152906001019060200180831161026957829003601f168201915b50509250505060405180910390a150565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102ea57805160ff1916838001178555610317565b82800160010185558215610317579182015b828111156103175782518255916020019190600101906102fc565b50610323929150610327565b5090565b6101d391905b80821115610323576000815560010161032d5600a165627a7a723058203c898ea7446d30108e4fb61330a63addb11f253341522ca1e1a36a3cd35ca3310029`

// DeployNameService deploys a new Ethereum contract, binding an instance of NameService to it.
func DeployNameService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameService, error) {
	parsed, err := abi.JSON(strings.NewReader(NameServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameService{NameServiceCaller: NameServiceCaller{contract: contract}, NameServiceTransactor: NameServiceTransactor{contract: contract}}, nil
}

// NameService is an auto generated Go binding around an Ethereum contract.
type NameService struct {
	NameServiceCaller     // Read-only binding to the contract
	NameServiceTransactor // Write-only binding to the contract
}

// NameServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameServiceSession struct {
	Contract     *NameService      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameServiceCallerSession struct {
	Contract *NameServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NameServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameServiceTransactorSession struct {
	Contract     *NameServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NameServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameServiceRaw struct {
	Contract *NameService // Generic contract binding to access the raw methods on
}

// NameServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameServiceCallerRaw struct {
	Contract *NameServiceCaller // Generic read-only contract binding to access the raw methods on
}

// NameServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameServiceTransactorRaw struct {
	Contract *NameServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameService creates a new instance of NameService, bound to a specific deployed contract.
func NewNameService(address common.Address, backend bind.ContractBackend) (*NameService, error) {
	contract, err := bindNameService(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameService{NameServiceCaller: NameServiceCaller{contract: contract}, NameServiceTransactor: NameServiceTransactor{contract: contract}}, nil
}

// NewNameServiceCaller creates a new read-only instance of NameService, bound to a specific deployed contract.
func NewNameServiceCaller(address common.Address, caller bind.ContractCaller) (*NameServiceCaller, error) {
	contract, err := bindNameService(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NameServiceCaller{contract: contract}, nil
}

// NewNameServiceTransactor creates a new write-only instance of NameService, bound to a specific deployed contract.
func NewNameServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*NameServiceTransactor, error) {
	contract, err := bindNameService(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NameServiceTransactor{contract: contract}, nil
}

// bindNameService binds a generic wrapper to an already deployed contract.
func bindNameService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameService *NameServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameService.Contract.NameServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameService *NameServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameService.Contract.NameServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameService *NameServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameService.Contract.NameServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameService *NameServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameService *NameServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameService *NameServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameService.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_NameService *NameServiceCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NameService.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_NameService *NameServiceSession) GetName() (string, error) {
	return _NameService.Contract.GetName(&_NameService.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_NameService *NameServiceCallerSession) GetName() (string, error) {
	return _NameService.Contract.GetName(&_NameService.CallOpts)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_NameService *NameServiceTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _NameService.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_NameService *NameServiceSession) SetName(_name string) (*types.Transaction, error) {
	return _NameService.Contract.SetName(&_NameService.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(_name string) returns()
func (_NameService *NameServiceTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _NameService.Contract.SetName(&_NameService.TransactOpts, _name)
}
