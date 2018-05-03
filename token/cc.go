// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// HumanStandardTokenABI is the input ABI used to generate the binding from.
const HumanStandardTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"name\":\"_tokenName\",\"type\":\"string\"},{\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":false,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// HumanStandardTokenBin is the compiled bytecode used for deploying new contracts.
const HumanStandardTokenBin = `0x606060405260408051908101604052600481527f48302e31000000000000000000000000000000000000000000000000000000006020820152600690805161004b9291602001906100e9565b50341561005757600080fd5b604051610bd1380380610bd183398101604052808051919060200180518201919060200180519190602001805190910190505b600160a060020a033316600090815260016020526040812085905584905560038380516100bb9291602001906100e9565b506004805460ff191660ff841617905560058180516100de9291602001906100e9565b505b50505050610189565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012a57805160ff1916838001178555610157565b82800160010185558215610157579182015b8281111561015757825182559160200191906001019061013c565b5b50610164929150610168565b5090565b61018691905b80821115610164576000815560010161016e565b5090565b90565b610a39806101986000396000f300606060405236156100935763ffffffff60e060020a60003504166306fdde0381146100a6578063095ea7b31461013157806318160ddd1461016757806323b872dd1461018c578063313ce567146101c857806354fd4d50146101f157806370a082311461027c57806395d89b41146102ad578063a9059cbb14610338578063cae9ca511461036e578063dd62ed3e146103e7575b341561009e57600080fd5b5b600080fd5b005b34156100b157600080fd5b6100b961041e565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f65780820151818401525b6020016100dd565b50505050905090810190601f1680156101235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013c57600080fd5b610153600160a060020a03600435166024356104bc565b604051901515815260200160405180910390f35b341561017257600080fd5b61017a610529565b60405190815260200160405180910390f35b341561019757600080fd5b610153600160a060020a0360043581169060243516604435610530565b604051901515815260200160405180910390f35b34156101d357600080fd5b6101db610629565b60405160ff909116815260200160405180910390f35b34156101fc57600080fd5b6100b9610632565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f65780820151818401525b6020016100dd565b50505050905090810190601f1680156101235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561028757600080fd5b61017a600160a060020a03600435166106d0565b60405190815260200160405180910390f35b34156102b857600080fd5b6100b96106ef565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f65780820151818401525b6020016100dd565b50505050905090810190601f1680156101235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561034357600080fd5b610153600160a060020a036004351660243561078d565b604051901515815260200160405180910390f35b341561037957600080fd5b61015360048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061083795505050505050565b604051901515815260200160405180910390f35b34156103f257600080fd5b61017a600160a060020a03600435811690602435166109e0565b60405190815260200160405180910390f35b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104b45780601f10610489576101008083540402835291602001916104b4565b820191906000526020600020905b81548152906001019060200180831161049757829003601f168201915b505050505081565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b6000545b90565b600160a060020a0383166000908152600160205260408120548290108015906105805750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561058c5750600082115b1561061d57600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610621565b5060005b5b9392505050565b60045460ff1681565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104b45780601f10610489576101008083540402835291602001916104b4565b820191906000526020600020905b81548152906001019060200180831161049757829003601f168201915b505050505081565b600160a060020a0381166000908152600160205260409020545b919050565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104b45780601f10610489576101008083540402835291602001916104b4565b820191906000526020600020905b81548152906001019060200180831161049757829003601f168201915b505050505081565b600160a060020a0333166000908152600160205260408120548290108015906107b65750600082115b1561082857600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610523565b506000610523565b5b92915050565b600160a060020a03338116600090815260026020908152604080832093871680845293909152808220859055909190517f72656365697665417070726f76616c28616464726573732c75696e743235362c81527f616464726573732c6279746573290000000000000000000000000000000000006020820152602e01604051809103902060e060020a9004338530866040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a03168152602001828051906020019080838360005b838110156109395780820151818401525b602001610920565b50505050905090810190601f1680156109665780820380516001836020036101000a031916815260200191505b5094505050505060006040518083038160008761646e5a03f192505050151561098e57600080fd5b83600160a060020a031633600160a060020a03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405190815260200160405180910390a35060015b9392505050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a72305820cb11f7795bb7d23ef017716f9f10c0c2bfe79452a2607b600b61512116f3bf310029`

// DeployHumanStandardToken deploys a new Ethereum contract, binding an instance of HumanStandardToken to it.
func DeployHumanStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend, _initialAmount *big.Int, _tokenName string, _decimalUnits uint8, _tokenSymbol string) (common.Address, *types.Transaction, *HumanStandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(HumanStandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HumanStandardTokenBin), backend, _initialAmount, _tokenName, _decimalUnits, _tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HumanStandardToken{HumanStandardTokenCaller: HumanStandardTokenCaller{contract: contract}, HumanStandardTokenTransactor: HumanStandardTokenTransactor{contract: contract}}, nil
}

// HumanStandardToken is an auto generated Go binding around an Ethereum contract.
type HumanStandardToken struct {
	HumanStandardTokenCaller     // Read-only binding to the contract
	HumanStandardTokenTransactor // Write-only binding to the contract
}

// HumanStandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type HumanStandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HumanStandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HumanStandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HumanStandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HumanStandardTokenSession struct {
	Contract     *HumanStandardToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HumanStandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HumanStandardTokenCallerSession struct {
	Contract *HumanStandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HumanStandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HumanStandardTokenTransactorSession struct {
	Contract     *HumanStandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HumanStandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type HumanStandardTokenRaw struct {
	Contract *HumanStandardToken // Generic contract binding to access the raw methods on
}

// HumanStandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HumanStandardTokenCallerRaw struct {
	Contract *HumanStandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// HumanStandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HumanStandardTokenTransactorRaw struct {
	Contract *HumanStandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHumanStandardToken creates a new instance of HumanStandardToken, bound to a specific deployed contract.
func NewHumanStandardToken(address common.Address, backend bind.ContractBackend) (*HumanStandardToken, error) {
	contract, err := bindHumanStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HumanStandardToken{HumanStandardTokenCaller: HumanStandardTokenCaller{contract: contract}, HumanStandardTokenTransactor: HumanStandardTokenTransactor{contract: contract}}, nil
}

// NewHumanStandardTokenCaller creates a new read-only instance of HumanStandardToken, bound to a specific deployed contract.
func NewHumanStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*HumanStandardTokenCaller, error) {
	contract, err := bindHumanStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &HumanStandardTokenCaller{contract: contract}, nil
}

// NewHumanStandardTokenTransactor creates a new write-only instance of HumanStandardToken, bound to a specific deployed contract.
func NewHumanStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*HumanStandardTokenTransactor, error) {
	contract, err := bindHumanStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &HumanStandardTokenTransactor{contract: contract}, nil
}

// bindHumanStandardToken binds a generic wrapper to an already deployed contract.
func bindHumanStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HumanStandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HumanStandardToken *HumanStandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HumanStandardToken.Contract.HumanStandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HumanStandardToken *HumanStandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.HumanStandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HumanStandardToken *HumanStandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.HumanStandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HumanStandardToken *HumanStandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HumanStandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HumanStandardToken *HumanStandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HumanStandardToken *HumanStandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_HumanStandardToken *HumanStandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_HumanStandardToken *HumanStandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _HumanStandardToken.Contract.Allowance(&_HumanStandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_HumanStandardToken *HumanStandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _HumanStandardToken.Contract.Allowance(&_HumanStandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_HumanStandardToken *HumanStandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_HumanStandardToken *HumanStandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _HumanStandardToken.Contract.BalanceOf(&_HumanStandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_HumanStandardToken *HumanStandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _HumanStandardToken.Contract.BalanceOf(&_HumanStandardToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_HumanStandardToken *HumanStandardTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_HumanStandardToken *HumanStandardTokenSession) Decimals() (uint8, error) {
	return _HumanStandardToken.Contract.Decimals(&_HumanStandardToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_HumanStandardToken *HumanStandardTokenCallerSession) Decimals() (uint8, error) {
	return _HumanStandardToken.Contract.Decimals(&_HumanStandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenSession) Name() (string, error) {
	return _HumanStandardToken.Contract.Name(&_HumanStandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenCallerSession) Name() (string, error) {
	return _HumanStandardToken.Contract.Name(&_HumanStandardToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenSession) Symbol() (string, error) {
	return _HumanStandardToken.Contract.Symbol(&_HumanStandardToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenCallerSession) Symbol() (string, error) {
	return _HumanStandardToken.Contract.Symbol(&_HumanStandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_HumanStandardToken *HumanStandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_HumanStandardToken *HumanStandardTokenSession) TotalSupply() (*big.Int, error) {
	return _HumanStandardToken.Contract.TotalSupply(&_HumanStandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_HumanStandardToken *HumanStandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _HumanStandardToken.Contract.TotalSupply(&_HumanStandardToken.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HumanStandardToken.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenSession) Version() (string, error) {
	return _HumanStandardToken.Contract.Version(&_HumanStandardToken.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_HumanStandardToken *HumanStandardTokenCallerSession) Version() (string, error) {
	return _HumanStandardToken.Contract.Version(&_HumanStandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.Approve(&_HumanStandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.Approve(&_HumanStandardToken.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _HumanStandardToken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.ApproveAndCall(&_HumanStandardToken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.ApproveAndCall(&_HumanStandardToken.TransactOpts, _spender, _value, _extraData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.Transfer(&_HumanStandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.Transfer(&_HumanStandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.TransferFrom(&_HumanStandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_HumanStandardToken *HumanStandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HumanStandardToken.Contract.TransferFrom(&_HumanStandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b5b61043e8061001f6000396000f300606060405236156100755763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007a57806318160ddd146100b057806323b872dd146100d557806370a0823114610111578063a9059cbb14610142578063dd62ed3e14610178575b600080fd5b341561008557600080fd5b61009c600160a060020a03600435166024356101af565b604051901515815260200160405180910390f35b34156100bb57600080fd5b6100c361021c565b60405190815260200160405180910390f35b34156100e057600080fd5b61009c600160a060020a0360043581169060243516604435610223565b604051901515815260200160405180910390f35b341561011c57600080fd5b6100c3600160a060020a036004351661031c565b60405190815260200160405180910390f35b341561014d57600080fd5b61009c600160a060020a036004351660243561033b565b604051901515815260200160405180910390f35b341561018357600080fd5b6100c3600160a060020a03600435811690602435166103e5565b60405190815260200160405180910390f35b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b6000545b90565b600160a060020a0383166000908152600160205260408120548290108015906102735750600160a060020a0380851660009081526002602090815260408083203390941683529290522054829010155b801561027f5750600082115b1561031057600160a060020a03808416600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529490915290819020805486900390559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610314565b5060005b5b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600160a060020a0333166000908152600160205260408120548290108015906103645750600082115b156103d657600160a060020a033381166000818152600160205260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3506001610216565b506000610216565b5b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b929150505600a165627a7a72305820ba975f9fd4a809b65df50109fc16ac4541cfaf4b79f8bae109cdeb260e64ed520029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}
