// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testing

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

// SwapperMetaData contains all meta data concerning the Swapper contract.
var SwapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getBlackfuryERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610045575f80fd5b50608051610b116100725f395f818160fe015281816102650152818161028a01526103740152610b115ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80631ecd53731461005957806347e7ef2414610089578063714ba40c146100a55780639d456b62146100c3578063d004f0f7146100df575b5f80fd5b610073600480360381019061006e91906104c1565b6100fb565b6040516100809190610586565b60405180910390f35b6100a3600480360381019061009e919061060d565b61019e565b005b6100ad610263565b6040516100ba919061066b565b60405180910390f35b6100dd60048036038101906100d89190610684565b610287565b005b6100f960048036038101906100f4919061071c565b610371565b005b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a333e57c84846040518363ffffffff1660e01b81526004016101579291906107b4565b602060405180830381865afa158015610172573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061019691906107ea565b905092915050565b5f8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016101dc93929190610833565b6020604051808303815f875af11580156101f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061021c919061089d565b90508061025e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025590610912565b60405180910390fd5b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663096b406985853333876040518663ffffffff1660e01b81526004016102e9959493929190610930565b6020604051808303815f875af1158015610305573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610329919061089d565b90508061036b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610362906109ec565b60405180910390fd5b50505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b96d8bec843333866040518563ffffffff1660e01b81526004016103d19493929190610a0a565b6020604051808303815f875af11580156103ed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610411919061089d565b905080610453576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044a90610abd565b60405180910390fd5b505050565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261048157610480610460565b5b8235905067ffffffffffffffff81111561049e5761049d610464565b5b6020830191508360018202830111156104ba576104b9610468565b5b9250929050565b5f80602083850312156104d7576104d6610458565b5b5f83013567ffffffffffffffff8111156104f4576104f361045c565b5b6105008582860161046c565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61054e6105496105448461050c565b61052b565b61050c565b9050919050565b5f61055f82610534565b9050919050565b5f61057082610555565b9050919050565b61058081610566565b82525050565b5f6020820190506105995f830184610577565b92915050565b5f6105a98261050c565b9050919050565b6105b98161059f565b81146105c3575f80fd5b50565b5f813590506105d4816105b0565b92915050565b5f819050919050565b6105ec816105da565b81146105f6575f80fd5b50565b5f81359050610607816105e3565b92915050565b5f806040838503121561062357610622610458565b5b5f610630858286016105c6565b9250506020610641858286016105f9565b9150509250929050565b5f61065582610555565b9050919050565b6106658161064b565b82525050565b5f60208201905061067e5f83018461065c565b92915050565b5f805f6040848603121561069b5761069a610458565b5b5f84013567ffffffffffffffff8111156106b8576106b761045c565b5b6106c48682870161046c565b935093505060206106d7868287016105f9565b9150509250925092565b5f6106eb8261059f565b9050919050565b6106fb816106e1565b8114610705575f80fd5b50565b5f81359050610716816106f2565b92915050565b5f806040838503121561073257610731610458565b5b5f61073f85828601610708565b9250506020610750858286016105f9565b9150509250929050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f610793838561075a565b93506107a083858461076a565b6107a983610778565b840190509392505050565b5f6020820190508181035f8301526107cd818486610788565b90509392505050565b5f815190506107e4816106f2565b92915050565b5f602082840312156107ff576107fe610458565b5b5f61080c848285016107d6565b91505092915050565b61081e8161059f565b82525050565b61082d816105da565b82525050565b5f6060820190506108465f830186610815565b6108536020830185610815565b6108606040830184610824565b949350505050565b5f8115159050919050565b61087c81610868565b8114610886575f80fd5b50565b5f8151905061089781610873565b92915050565b5f602082840312156108b2576108b1610458565b5b5f6108bf84828501610889565b91505092915050565b7f537761707065723a207472616e7366657246726f6d206661696c6564000000005f82015250565b5f6108fc601c8361075a565b9150610907826108c8565b602082019050919050565b5f6020820190508181035f830152610929816108f0565b9050919050565b5f6080820190508181035f830152610949818789610788565b90506109586020830186610815565b6109656040830185610815565b6109726060830184610824565b9695505050505050565b7f537761707065723a207472616e73666572436f696e546f4552433230206661695f8201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b5f6109d660238361075a565b91506109e18261097c565b604082019050919050565b5f6020820190508181035f830152610a03816109ca565b9050919050565b5f608082019050610a1d5f830187610577565b610a2a6020830186610815565b610a376040830185610815565b610a446060830184610824565b95945050505050565b7f537761707065723a207472616e736665724552433230546f436f696e206661695f8201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b5f610aa760238361075a565b9150610ab282610a4d565b604082019050919050565b5f6020820190508181035f830152610ad481610a9b565b905091905056fea264697066735822122068b9d2fa8a14a0af555f17a4f1586e8501588459f2fc595598d18b58e73178cd64736f6c63430008140033",
}

// SwapperABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapperMetaData.ABI instead.
var SwapperABI = SwapperMetaData.ABI

// SwapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapperMetaData.Bin instead.
var SwapperBin = SwapperMetaData.Bin

// DeploySwapper deploys a new Ethereum contract, binding an instance of Swapper to it.
func DeploySwapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swapper, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// Swapper is an auto generated Go binding around an Ethereum contract.
type Swapper struct {
	SwapperCaller     // Read-only binding to the contract
	SwapperTransactor // Write-only binding to the contract
	SwapperFilterer   // Log filterer for contract events
}

// SwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapperSession struct {
	Contract     *Swapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapperCallerSession struct {
	Contract *SwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapperTransactorSession struct {
	Contract     *SwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapperRaw struct {
	Contract *Swapper // Generic contract binding to access the raw methods on
}

// SwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapperCallerRaw struct {
	Contract *SwapperCaller // Generic read-only contract binding to access the raw methods on
}

// SwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapperTransactorRaw struct {
	Contract *SwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapper creates a new instance of Swapper, bound to a specific deployed contract.
func NewSwapper(address common.Address, backend bind.ContractBackend) (*Swapper, error) {
	contract, err := bindSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// NewSwapperCaller creates a new read-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperCaller(address common.Address, caller bind.ContractCaller) (*SwapperCaller, error) {
	contract, err := bindSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperCaller{contract: contract}, nil
}

// NewSwapperTransactor creates a new write-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapperTransactor, error) {
	contract, err := bindSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperTransactor{contract: contract}, nil
}

// NewSwapperFilterer creates a new log filterer instance of Swapper, bound to a specific deployed contract.
func NewSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapperFilterer, error) {
	contract, err := bindSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapperFilterer{contract: contract}, nil
}

// bindSwapper binds a generic wrapper to an already deployed contract.
func bindSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.SwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transact(opts, method, params...)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperCaller) Erc20Module(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "erc20Module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperSession) Erc20Module() (common.Address, error) {
	return _Swapper.Contract.Erc20Module(&_Swapper.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperCallerSession) Erc20Module() (common.Address, error) {
	return _Swapper.Contract.Erc20Module(&_Swapper.CallOpts)
}

// GetBlackfuryERC20 is a free data retrieval call binding the contract method 0x1ecd5373.
//
// Solidity: function getBlackfuryERC20(string denom) view returns(address)
func (_Swapper *SwapperCaller) GetBlackfuryERC20(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getBlackfuryERC20", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBlackfuryERC20 is a free data retrieval call binding the contract method 0x1ecd5373.
//
// Solidity: function getBlackfuryERC20(string denom) view returns(address)
func (_Swapper *SwapperSession) GetBlackfuryERC20(denom string) (common.Address, error) {
	return _Swapper.Contract.GetBlackfuryERC20(&_Swapper.CallOpts, denom)
}

// GetBlackfuryERC20 is a free data retrieval call binding the contract method 0x1ecd5373.
//
// Solidity: function getBlackfuryERC20(string denom) view returns(address)
func (_Swapper *SwapperCallerSession) GetBlackfuryERC20(denom string) (common.Address, error) {
	return _Swapper.Contract.GetBlackfuryERC20(&_Swapper.CallOpts, denom)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Deposit(&_Swapper.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Deposit(&_Swapper.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Swap(opts *bind.TransactOpts, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap", denom, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperSession) Swap(denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, denom, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Swap(denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, denom, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Swap0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap0", token, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperSession) Swap0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap0(&_Swapper.TransactOpts, token, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Swap0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap0(&_Swapper.TransactOpts, token, amount)
}
