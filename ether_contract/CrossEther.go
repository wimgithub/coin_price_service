// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CrossEther

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CrossEtherABI is the input ABI used to generate the binding from.
const CrossEtherABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"getUserEthByTrxAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getUserTrxByEthAddr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOwner\",\"type\":\"bool\"}],\"name\":\"ownerPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ownerSetErc20Addr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ownerWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"setTrxAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CrossEtherBin is the compiled bytecode used for deploying new contracts.
var CrossEtherBin = "0x608060405234801561001057600080fd5b50604051610f2d380380610f2d8339818101604052602081101561003357600080fd5b810190808051906020019092919050505060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610e41806100ec6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a6021ace11610066578063a6021ace146102a8578063d103bebf146102dc578063d9c88e1414610320578063dbaf2b841461036e578063ef299b0b1461042957610093565b80630248184d146100985780632ddbd13a146100e8578063449b466f1461010657806354ee14c9146101eb575b600080fd5b6100e6600480360360408110156100ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610457565b005b6100f0610547565b6040518082815260200191505060405180910390f35b6101bf6004803603602081101561011c57600080fd5b810190808035906020019064010000000081111561013957600080fd5b82018360208201111561014b57600080fd5b8035906020019184600183028401116401000000008311171561016d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061054d565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61022d6004803603602081101561020157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105e0565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561026d578082015181840152602081019050610252565b50505050905090810190601f16801561029a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102b06106c1565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61031e600480360360208110156102f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106e7565b005b61036c6004803603604081101561033657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107c1565b005b6104276004803603602081101561038457600080fd5b81019080803590602001906401000000008111156103a157600080fd5b8201836020820111156103b357600080fd5b803590602001918460018302840111640100000000831117156103d557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061098c565b005b6104556004803603602081101561043f57600080fd5b8101908080359060200190929190505050610b4d565b005b600115156000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146104b357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156104ed57600080fd5b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60015481565b60006003826040518082805190602001908083835b602083106105855780518252602082019150602081019050602083039250610562565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106b55780601f1061068a576101008083540402835291602001916106b5565b820191906000526020600020905b81548152906001019060200180831161069857829003601f168201915b50505050509050919050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600115156000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461074357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561077d57600080fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600115156000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461081d57600080fd5b6000811415801561085b5750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b61086457600080fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156108f757600080fd5b505af115801561090b573d6000803e3d6000fd5b505050506040513d602081101561092157600080fd5b8101908080519060200190929190505050507f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648282604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050565b600073ffffffffffffffffffffffffffffffffffffffff166003826040518082805190602001908083835b602083106109da57805182526020820191506020810190506020830392506109b7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a4e57600080fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209080519060200190610aa1929190610d6e565b50336003826040518082805190602001908083835b60208310610ad95780518252602082019150602081019050602083039250610ab6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008111610b5a57600080fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610c0b57600080fd5b505af1158015610c1f573d6000803e3d6000fd5b505050506040513d6020811015610c3557600080fd5b810190808051906020019092919050505050610c5c81600154610d4f90919063ffffffff16565b600181905550610cb481600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d4f90919063ffffffff16565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def33382604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a150565b600080828401905083811015610d6457600080fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610daf57805160ff1916838001178555610ddd565b82800160010185558215610ddd579182015b82811115610ddc578251825591602001919060010190610dc1565b5b509050610dea9190610dee565b5090565b5b80821115610e07576000816000905550600101610def565b509056fea2646970667358221220e156073184e40dde38ab755a4093f8976ac007dffe8033d9452c1bc9b8f6bc4a64736f6c634300060c0033"

// DeployCrossEther deploys a new Ethereum contract, binding an instance of CrossEther to it.
func DeployCrossEther(auth *bind.TransactOpts, backend bind.ContractBackend, erc20 common.Address) (common.Address, *types.Transaction, *CrossEther, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossEtherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrossEtherBin), backend, erc20)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossEther{CrossEtherCaller: CrossEtherCaller{contract: contract}, CrossEtherTransactor: CrossEtherTransactor{contract: contract}, CrossEtherFilterer: CrossEtherFilterer{contract: contract}}, nil
}

// CrossEther is an auto generated Go binding around an Ethereum contract.
type CrossEther struct {
	CrossEtherCaller     // Read-only binding to the contract
	CrossEtherTransactor // Write-only binding to the contract
	CrossEtherFilterer   // Log filterer for contract events
}

// CrossEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossEtherSession struct {
	Contract     *CrossEther       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossEtherCallerSession struct {
	Contract *CrossEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrossEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossEtherTransactorSession struct {
	Contract     *CrossEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrossEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossEtherRaw struct {
	Contract *CrossEther // Generic contract binding to access the raw methods on
}

// CrossEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossEtherCallerRaw struct {
	Contract *CrossEtherCaller // Generic read-only contract binding to access the raw methods on
}

// CrossEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossEtherTransactorRaw struct {
	Contract *CrossEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossEther creates a new instance of CrossEther, bound to a specific deployed contract.
func NewCrossEther(address common.Address, backend bind.ContractBackend) (*CrossEther, error) {
	contract, err := bindCrossEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossEther{CrossEtherCaller: CrossEtherCaller{contract: contract}, CrossEtherTransactor: CrossEtherTransactor{contract: contract}, CrossEtherFilterer: CrossEtherFilterer{contract: contract}}, nil
}

// NewCrossEtherCaller creates a new read-only instance of CrossEther, bound to a specific deployed contract.
func NewCrossEtherCaller(address common.Address, caller bind.ContractCaller) (*CrossEtherCaller, error) {
	contract, err := bindCrossEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossEtherCaller{contract: contract}, nil
}

// NewCrossEtherTransactor creates a new write-only instance of CrossEther, bound to a specific deployed contract.
func NewCrossEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossEtherTransactor, error) {
	contract, err := bindCrossEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossEtherTransactor{contract: contract}, nil
}

// NewCrossEtherFilterer creates a new log filterer instance of CrossEther, bound to a specific deployed contract.
func NewCrossEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossEtherFilterer, error) {
	contract, err := bindCrossEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossEtherFilterer{contract: contract}, nil
}

// bindCrossEther binds a generic wrapper to an already deployed contract.
func bindCrossEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossEther *CrossEtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossEther.Contract.CrossEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossEther *CrossEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossEther.Contract.CrossEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossEther *CrossEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossEther.Contract.CrossEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossEther *CrossEtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossEther *CrossEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossEther *CrossEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossEther.Contract.contract.Transact(opts, method, params...)
}

// ERC20Address is a free data retrieval call binding the contract method 0xa6021ace.
//
// Solidity: function ERC20Address() view returns(address)
func (_CrossEther *CrossEtherCaller) ERC20Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossEther.contract.Call(opts, &out, "ERC20Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Address is a free data retrieval call binding the contract method 0xa6021ace.
//
// Solidity: function ERC20Address() view returns(address)
func (_CrossEther *CrossEtherSession) ERC20Address() (common.Address, error) {
	return _CrossEther.Contract.ERC20Address(&_CrossEther.CallOpts)
}

// ERC20Address is a free data retrieval call binding the contract method 0xa6021ace.
//
// Solidity: function ERC20Address() view returns(address)
func (_CrossEther *CrossEtherCallerSession) ERC20Address() (common.Address, error) {
	return _CrossEther.Contract.ERC20Address(&_CrossEther.CallOpts)
}

// GetUserEthByTrxAddr is a free data retrieval call binding the contract method 0x449b466f.
//
// Solidity: function getUserEthByTrxAddr(string addr) view returns(address)
func (_CrossEther *CrossEtherCaller) GetUserEthByTrxAddr(opts *bind.CallOpts, addr string) (common.Address, error) {
	var out []interface{}
	err := _CrossEther.contract.Call(opts, &out, "getUserEthByTrxAddr", addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserEthByTrxAddr is a free data retrieval call binding the contract method 0x449b466f.
//
// Solidity: function getUserEthByTrxAddr(string addr) view returns(address)
func (_CrossEther *CrossEtherSession) GetUserEthByTrxAddr(addr string) (common.Address, error) {
	return _CrossEther.Contract.GetUserEthByTrxAddr(&_CrossEther.CallOpts, addr)
}

// GetUserEthByTrxAddr is a free data retrieval call binding the contract method 0x449b466f.
//
// Solidity: function getUserEthByTrxAddr(string addr) view returns(address)
func (_CrossEther *CrossEtherCallerSession) GetUserEthByTrxAddr(addr string) (common.Address, error) {
	return _CrossEther.Contract.GetUserEthByTrxAddr(&_CrossEther.CallOpts, addr)
}

// GetUserTrxByEthAddr is a free data retrieval call binding the contract method 0x54ee14c9.
//
// Solidity: function getUserTrxByEthAddr(address addr) view returns(string)
func (_CrossEther *CrossEtherCaller) GetUserTrxByEthAddr(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _CrossEther.contract.Call(opts, &out, "getUserTrxByEthAddr", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserTrxByEthAddr is a free data retrieval call binding the contract method 0x54ee14c9.
//
// Solidity: function getUserTrxByEthAddr(address addr) view returns(string)
func (_CrossEther *CrossEtherSession) GetUserTrxByEthAddr(addr common.Address) (string, error) {
	return _CrossEther.Contract.GetUserTrxByEthAddr(&_CrossEther.CallOpts, addr)
}

// GetUserTrxByEthAddr is a free data retrieval call binding the contract method 0x54ee14c9.
//
// Solidity: function getUserTrxByEthAddr(address addr) view returns(string)
func (_CrossEther *CrossEtherCallerSession) GetUserTrxByEthAddr(addr common.Address) (string, error) {
	return _CrossEther.Contract.GetUserTrxByEthAddr(&_CrossEther.CallOpts, addr)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_CrossEther *CrossEtherCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossEther.contract.Call(opts, &out, "total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_CrossEther *CrossEtherSession) Total() (*big.Int, error) {
	return _CrossEther.Contract.Total(&_CrossEther.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_CrossEther *CrossEtherCallerSession) Total() (*big.Int, error) {
	return _CrossEther.Contract.Total(&_CrossEther.CallOpts)
}

// OwnerPermissions is a paid mutator transaction binding the contract method 0x0248184d.
//
// Solidity: function ownerPermissions(address newOwner, bool isOwner) returns()
func (_CrossEther *CrossEtherTransactor) OwnerPermissions(opts *bind.TransactOpts, newOwner common.Address, isOwner bool) (*types.Transaction, error) {
	return _CrossEther.contract.Transact(opts, "ownerPermissions", newOwner, isOwner)
}

// OwnerPermissions is a paid mutator transaction binding the contract method 0x0248184d.
//
// Solidity: function ownerPermissions(address newOwner, bool isOwner) returns()
func (_CrossEther *CrossEtherSession) OwnerPermissions(newOwner common.Address, isOwner bool) (*types.Transaction, error) {
	return _CrossEther.Contract.OwnerPermissions(&_CrossEther.TransactOpts, newOwner, isOwner)
}

// OwnerPermissions is a paid mutator transaction binding the contract method 0x0248184d.
//
// Solidity: function ownerPermissions(address newOwner, bool isOwner) returns()
func (_CrossEther *CrossEtherTransactorSession) OwnerPermissions(newOwner common.Address, isOwner bool) (*types.Transaction, error) {
	return _CrossEther.Contract.OwnerPermissions(&_CrossEther.TransactOpts, newOwner, isOwner)
}

// OwnerSetErc20Addr is a paid mutator transaction binding the contract method 0xd103bebf.
//
// Solidity: function ownerSetErc20Addr(address addr) returns()
func (_CrossEther *CrossEtherTransactor) OwnerSetErc20Addr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CrossEther.contract.Transact(opts, "ownerSetErc20Addr", addr)
}

// OwnerSetErc20Addr is a paid mutator transaction binding the contract method 0xd103bebf.
//
// Solidity: function ownerSetErc20Addr(address addr) returns()
func (_CrossEther *CrossEtherSession) OwnerSetErc20Addr(addr common.Address) (*types.Transaction, error) {
	return _CrossEther.Contract.OwnerSetErc20Addr(&_CrossEther.TransactOpts, addr)
}

// OwnerSetErc20Addr is a paid mutator transaction binding the contract method 0xd103bebf.
//
// Solidity: function ownerSetErc20Addr(address addr) returns()
func (_CrossEther *CrossEtherTransactorSession) OwnerSetErc20Addr(addr common.Address) (*types.Transaction, error) {
	return _CrossEther.Contract.OwnerSetErc20Addr(&_CrossEther.TransactOpts, addr)
}

// OwnerWithdraw is a paid mutator transaction binding the contract method 0xd9c88e14.
//
// Solidity: function ownerWithdraw(address addr, uint256 value) returns()
func (_CrossEther *CrossEtherTransactor) OwnerWithdraw(opts *bind.TransactOpts, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossEther.contract.Transact(opts, "ownerWithdraw", addr, value)
}

// OwnerWithdraw is a paid mutator transaction binding the contract method 0xd9c88e14.
//
// Solidity: function ownerWithdraw(address addr, uint256 value) returns()
func (_CrossEther *CrossEtherSession) OwnerWithdraw(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossEther.Contract.OwnerWithdraw(&_CrossEther.TransactOpts, addr, value)
}

// OwnerWithdraw is a paid mutator transaction binding the contract method 0xd9c88e14.
//
// Solidity: function ownerWithdraw(address addr, uint256 value) returns()
func (_CrossEther *CrossEtherTransactorSession) OwnerWithdraw(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossEther.Contract.OwnerWithdraw(&_CrossEther.TransactOpts, addr, value)
}

// Recharge is a paid mutator transaction binding the contract method 0xef299b0b.
//
// Solidity: function recharge(uint256 value) returns()
func (_CrossEther *CrossEtherTransactor) Recharge(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _CrossEther.contract.Transact(opts, "recharge", value)
}

// Recharge is a paid mutator transaction binding the contract method 0xef299b0b.
//
// Solidity: function recharge(uint256 value) returns()
func (_CrossEther *CrossEtherSession) Recharge(value *big.Int) (*types.Transaction, error) {
	return _CrossEther.Contract.Recharge(&_CrossEther.TransactOpts, value)
}

// Recharge is a paid mutator transaction binding the contract method 0xef299b0b.
//
// Solidity: function recharge(uint256 value) returns()
func (_CrossEther *CrossEtherTransactorSession) Recharge(value *big.Int) (*types.Transaction, error) {
	return _CrossEther.Contract.Recharge(&_CrossEther.TransactOpts, value)
}

// SetTrxAddress is a paid mutator transaction binding the contract method 0xdbaf2b84.
//
// Solidity: function setTrxAddress(string addr) returns()
func (_CrossEther *CrossEtherTransactor) SetTrxAddress(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return _CrossEther.contract.Transact(opts, "setTrxAddress", addr)
}

// SetTrxAddress is a paid mutator transaction binding the contract method 0xdbaf2b84.
//
// Solidity: function setTrxAddress(string addr) returns()
func (_CrossEther *CrossEtherSession) SetTrxAddress(addr string) (*types.Transaction, error) {
	return _CrossEther.Contract.SetTrxAddress(&_CrossEther.TransactOpts, addr)
}

// SetTrxAddress is a paid mutator transaction binding the contract method 0xdbaf2b84.
//
// Solidity: function setTrxAddress(string addr) returns()
func (_CrossEther *CrossEtherTransactorSession) SetTrxAddress(addr string) (*types.Transaction, error) {
	return _CrossEther.Contract.SetTrxAddress(&_CrossEther.TransactOpts, addr)
}

// CrossEtherReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the CrossEther contract.
type CrossEtherReceiveIterator struct {
	Event *CrossEtherReceive // Event containing the contract specifics and raw log

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
func (it *CrossEtherReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossEtherReceive)
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
		it.Event = new(CrossEtherReceive)
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
func (it *CrossEtherReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossEtherReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossEtherReceive represents a Receive event raised by the CrossEther contract.
type CrossEtherReceive struct {
	Addr  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0xd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def3.
//
// Solidity: event Receive(address addr, uint256 value)
func (_CrossEther *CrossEtherFilterer) FilterReceive(opts *bind.FilterOpts) (*CrossEtherReceiveIterator, error) {

	logs, sub, err := _CrossEther.contract.FilterLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return &CrossEtherReceiveIterator{contract: _CrossEther.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0xd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def3.
//
// Solidity: event Receive(address addr, uint256 value)
func (_CrossEther *CrossEtherFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *CrossEtherReceive) (event.Subscription, error) {

	logs, sub, err := _CrossEther.contract.WatchLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossEtherReceive)
				if err := _CrossEther.contract.UnpackLog(event, "Receive", log); err != nil {
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

// ParseReceive is a log parse operation binding the contract event 0xd6717f327e0cb88b4a97a7f67a453e9258252c34937ccbdd86de7cb840e7def3.
//
// Solidity: event Receive(address addr, uint256 value)
func (_CrossEther *CrossEtherFilterer) ParseReceive(log types.Log) (*CrossEtherReceive, error) {
	event := new(CrossEtherReceive)
	if err := _CrossEther.contract.UnpackLog(event, "Receive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrossEtherWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CrossEther contract.
type CrossEtherWithdrawIterator struct {
	Event *CrossEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *CrossEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossEtherWithdraw)
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
		it.Event = new(CrossEtherWithdraw)
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
func (it *CrossEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossEtherWithdraw represents a Withdraw event raised by the CrossEther contract.
type CrossEtherWithdraw struct {
	Addr  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 value)
func (_CrossEther *CrossEtherFilterer) FilterWithdraw(opts *bind.FilterOpts) (*CrossEtherWithdrawIterator, error) {

	logs, sub, err := _CrossEther.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &CrossEtherWithdrawIterator{contract: _CrossEther.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 value)
func (_CrossEther *CrossEtherFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CrossEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _CrossEther.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossEtherWithdraw)
				if err := _CrossEther.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 value)
func (_CrossEther *CrossEtherFilterer) ParseWithdraw(log types.Log) (*CrossEtherWithdraw, error) {
	event := new(CrossEtherWithdraw)
	if err := _CrossEther.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
