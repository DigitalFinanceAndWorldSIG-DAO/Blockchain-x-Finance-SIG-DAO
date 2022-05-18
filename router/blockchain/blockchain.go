// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlockchainABI is the input ABI used to generate the binding from.
const BlockchainABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"cleartext\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cryphtext\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cleartext\",\"type\":\"string\"}],\"name\":\"getcleartext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cryphtext\",\"type\":\"string\"}],\"name\":\"getcryphtext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockchainBin is the compiled bytecode used for deploying new contracts.
var BlockchainBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506108c8806100606000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063466a746e146100675780638da5cb5b14610085578063a3904a34146100cf578063ba9984aa14610152578063c4b67a96146101cb578063eda0040e1461024e575b600080fd5b61006f6102c7565b6040518082815260200191505060405180910390f35b61008d6102cd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100d76102f2565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101175780820151818401526020810190506100fc565b50505050905090810190601f1680156101445780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c96004803603602081101561016857600080fd5b810190808035906020019064010000000081111561018557600080fd5b82018360208201111561019757600080fd5b803590602001918460018302840111640100000000831117156101b957600080fd5b9091929391929390505050610390565b005b6101d361059b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102135780820151818401526020810190506101f8565b50505050905090810190601f1680156102405780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c56004803603602081101561026457600080fd5b810190808035906020019064010000000081111561028157600080fd5b82018360208201111561029357600080fd5b803590602001918460018302840111640100000000831117156102b557600080fd5b9091929391929390505050610639565b005b60035481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103885780601f1061035d57610100808354040283529160200191610388565b820191906000526020600020905b81548152906001019060200180831161036b57829003601f168201915b505050505081565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610452576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807fe8baabe4bbbde8aea4e8af81e5a4b1e8b4a5000000000000000000000000000081525060200191505060405180910390fd5b8181600191906104639291906107ed565b5060016000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156105465780601f1061051b57610100808354040283529160200191610546565b820191906000526020600020905b81548152906001019060200180831161052957829003601f168201915b5050935050505060405160208183030381529060405280519060200120600381905550600354600483836040518083838082843780830192505050925050509081526020016040518091039020819055505050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106315780601f1061060657610100808354040283529160200191610631565b820191906000526020600020905b81548152906001019060200180831161061457829003601f168201915b505050505081565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807fe8baabe4bbbde8aea4e8af81e5a4b1e8b4a5000000000000000000000000000081525060200191505060405180910390fd5b81816002919061070c9291906107ed565b5081816000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252858582818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405160208183030381529060405280519060200120600381905550600354600583836040518083838082843780830192505050925050509081526020016040518091039020819055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061082e57803560ff191683800117855561085c565b8280016001018555821561085c579182015b8281111561085b578235825591602001919060010190610840565b5b509050610869919061086d565b5090565b61088f91905b8082111561088b576000816000905550600101610873565b5090565b9056fea26469706673582212204b37b97157fc69a93f491ed7db971bac866fdf99cf427b5ce9b9b888b09744ea64736f6c634300060a0033"

// DeployBlockchain deploys a new contract, binding an instance of Blockchain to it.
func DeployBlockchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blockchain, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlockchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

func AsyncDeployBlockchain(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(BlockchainBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Blockchain is an auto generated Go binding around a Solidity contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around a Solidity contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around a Solidity contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around a Solidity contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// Cleartext is a free data retrieval call binding the contract method 0xa3904a34.
//
// Solidity: function cleartext() constant returns(string)
func (_Blockchain *BlockchainCaller) Cleartext(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Blockchain.contract.Call(opts, out, "cleartext")
	return *ret0, err
}

// Cleartext is a free data retrieval call binding the contract method 0xa3904a34.
//
// Solidity: function cleartext() constant returns(string)
func (_Blockchain *BlockchainSession) Cleartext() (string, error) {
	return _Blockchain.Contract.Cleartext(&_Blockchain.CallOpts)
}

// Cleartext is a free data retrieval call binding the contract method 0xa3904a34.
//
// Solidity: function cleartext() constant returns(string)
func (_Blockchain *BlockchainCallerSession) Cleartext() (string, error) {
	return _Blockchain.Contract.Cleartext(&_Blockchain.CallOpts)
}

// Cryphtext is a free data retrieval call binding the contract method 0xc4b67a96.
//
// Solidity: function cryphtext() constant returns(string)
func (_Blockchain *BlockchainCaller) Cryphtext(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Blockchain.contract.Call(opts, out, "cryphtext")
	return *ret0, err
}

// Cryphtext is a free data retrieval call binding the contract method 0xc4b67a96.
//
// Solidity: function cryphtext() constant returns(string)
func (_Blockchain *BlockchainSession) Cryphtext() (string, error) {
	return _Blockchain.Contract.Cryphtext(&_Blockchain.CallOpts)
}

// Cryphtext is a free data retrieval call binding the contract method 0xc4b67a96.
//
// Solidity: function cryphtext() constant returns(string)
func (_Blockchain *BlockchainCallerSession) Cryphtext() (string, error) {
	return _Blockchain.Contract.Cryphtext(&_Blockchain.CallOpts)
}

// HashCode is a free data retrieval call binding the contract method 0x466a746e.
//
// Solidity: function hashCode() constant returns(bytes32)
func (_Blockchain *BlockchainCaller) HashCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Blockchain.contract.Call(opts, out, "hashCode")
	return *ret0, err
}

// HashCode is a free data retrieval call binding the contract method 0x466a746e.
//
// Solidity: function hashCode() constant returns(bytes32)
func (_Blockchain *BlockchainSession) HashCode() ([32]byte, error) {
	return _Blockchain.Contract.HashCode(&_Blockchain.CallOpts)
}

// HashCode is a free data retrieval call binding the contract method 0x466a746e.
//
// Solidity: function hashCode() constant returns(bytes32)
func (_Blockchain *BlockchainCallerSession) HashCode() ([32]byte, error) {
	return _Blockchain.Contract.HashCode(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Blockchain *BlockchainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Blockchain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Blockchain *BlockchainSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Blockchain *BlockchainCallerSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// Getcleartext is a paid mutator transaction binding the contract method 0xeda0040e.
//
// Solidity: function getcleartext(string _cleartext) returns()
func (_Blockchain *BlockchainTransactor) Getcleartext(opts *bind.TransactOpts, _cleartext string) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.contract.Transact(opts, "getcleartext", _cleartext)
}

func (_Blockchain *BlockchainTransactor) AsyncGetcleartext(handler func(*types.Receipt, error), opts *bind.TransactOpts, _cleartext string) (*types.Transaction, error) {
	return _Blockchain.contract.AsyncTransact(opts, handler, "getcleartext", _cleartext)
}

// Getcleartext is a paid mutator transaction binding the contract method 0xeda0040e.
//
// Solidity: function getcleartext(string _cleartext) returns()
func (_Blockchain *BlockchainSession) Getcleartext(_cleartext string) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.Getcleartext(&_Blockchain.TransactOpts, _cleartext)
}

func (_Blockchain *BlockchainSession) AsyncGetcleartext(handler func(*types.Receipt, error), _cleartext string) (*types.Transaction, error) {
	return _Blockchain.Contract.AsyncGetcleartext(handler, &_Blockchain.TransactOpts, _cleartext)
}

// Getcleartext is a paid mutator transaction binding the contract method 0xeda0040e.
//
// Solidity: function getcleartext(string _cleartext) returns()
func (_Blockchain *BlockchainTransactorSession) Getcleartext(_cleartext string) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.Getcleartext(&_Blockchain.TransactOpts, _cleartext)
}

func (_Blockchain *BlockchainTransactorSession) AsyncGetcleartext(handler func(*types.Receipt, error), _cleartext string) (*types.Transaction, error) {
	return _Blockchain.Contract.AsyncGetcleartext(handler, &_Blockchain.TransactOpts, _cleartext)
}

// Getcryphtext is a paid mutator transaction binding the contract method 0xba9984aa.
//
// Solidity: function getcryphtext(string _cryphtext) returns()
func (_Blockchain *BlockchainTransactor) Getcryphtext(opts *bind.TransactOpts, _cryphtext string) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.contract.Transact(opts, "getcryphtext", _cryphtext)
}

func (_Blockchain *BlockchainTransactor) AsyncGetcryphtext(handler func(*types.Receipt, error), opts *bind.TransactOpts, _cryphtext string) (*types.Transaction, error) {
	return _Blockchain.contract.AsyncTransact(opts, handler, "getcryphtext", _cryphtext)
}

// Getcryphtext is a paid mutator transaction binding the contract method 0xba9984aa.
//
// Solidity: function getcryphtext(string _cryphtext) returns()
func (_Blockchain *BlockchainSession) Getcryphtext(_cryphtext string) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.Getcryphtext(&_Blockchain.TransactOpts, _cryphtext)
}

func (_Blockchain *BlockchainSession) AsyncGetcryphtext(handler func(*types.Receipt, error), _cryphtext string) (*types.Transaction, error) {
	return _Blockchain.Contract.AsyncGetcryphtext(handler, &_Blockchain.TransactOpts, _cryphtext)
}

// Getcryphtext is a paid mutator transaction binding the contract method 0xba9984aa.
//
// Solidity: function getcryphtext(string _cryphtext) returns()
func (_Blockchain *BlockchainTransactorSession) Getcryphtext(_cryphtext string) (*types.Transaction, *types.Receipt, error) {
	return _Blockchain.Contract.Getcryphtext(&_Blockchain.TransactOpts, _cryphtext)
}

func (_Blockchain *BlockchainTransactorSession) AsyncGetcryphtext(handler func(*types.Receipt, error), _cryphtext string) (*types.Transaction, error) {
	return _Blockchain.Contract.AsyncGetcryphtext(handler, &_Blockchain.TransactOpts, _cryphtext)
}
