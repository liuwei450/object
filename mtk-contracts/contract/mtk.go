// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// StakeContractMetaData contains all meta data concerning the StakeContract contract.
var StakeContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_mtkToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"StakeNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumMtkContracts.StakingPeriod\",\"name\":\"period\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumMtkContracts.StakingPeriod\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"apy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMtkContracts.StakingPeriod\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"durations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumMtkContracts.StakingPeriod\",\"name\":\"period\",\"type\":\"uint8\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeIdToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"enumMtkContracts.StakingPeriod\",\"name\":\"period\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeContractMetaData.ABI instead.
var StakeContractABI = StakeContractMetaData.ABI

// StakeContract is an auto generated Go binding around an Ethereum contract.
type StakeContract struct {
	StakeContractCaller     // Read-only binding to the contract
	StakeContractTransactor // Write-only binding to the contract
	StakeContractFilterer   // Log filterer for contract events
}

// StakeContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeContractSession struct {
	Contract     *StakeContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeContractCallerSession struct {
	Contract *StakeContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakeContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeContractTransactorSession struct {
	Contract     *StakeContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakeContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeContractRaw struct {
	Contract *StakeContract // Generic contract binding to access the raw methods on
}

// StakeContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeContractCallerRaw struct {
	Contract *StakeContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakeContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeContractTransactorRaw struct {
	Contract *StakeContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeContract creates a new instance of StakeContract, bound to a specific deployed contract.
func NewStakeContract(address common.Address, backend bind.ContractBackend) (*StakeContract, error) {
	contract, err := bindStakeContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeContract{StakeContractCaller: StakeContractCaller{contract: contract}, StakeContractTransactor: StakeContractTransactor{contract: contract}, StakeContractFilterer: StakeContractFilterer{contract: contract}}, nil
}

// NewStakeContractCaller creates a new read-only instance of StakeContract, bound to a specific deployed contract.
func NewStakeContractCaller(address common.Address, caller bind.ContractCaller) (*StakeContractCaller, error) {
	contract, err := bindStakeContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeContractCaller{contract: contract}, nil
}

// NewStakeContractTransactor creates a new write-only instance of StakeContract, bound to a specific deployed contract.
func NewStakeContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeContractTransactor, error) {
	contract, err := bindStakeContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeContractTransactor{contract: contract}, nil
}

// NewStakeContractFilterer creates a new log filterer instance of StakeContract, bound to a specific deployed contract.
func NewStakeContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeContractFilterer, error) {
	contract, err := bindStakeContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeContractFilterer{contract: contract}, nil
}

// bindStakeContract binds a generic wrapper to an already deployed contract.
func bindStakeContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeContract *StakeContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeContract.Contract.StakeContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeContract *StakeContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeContract.Contract.StakeContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeContract *StakeContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeContract.Contract.StakeContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeContract *StakeContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeContract *StakeContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeContract *StakeContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeContract.Contract.contract.Transact(opts, method, params...)
}

// Apy is a free data retrieval call binding the contract method 0x1f1accb2.
//
// Solidity: function apy(uint8 ) view returns(uint256)
func (_StakeContract *StakeContractCaller) Apy(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeContract.contract.Call(opts, &out, "apy", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apy is a free data retrieval call binding the contract method 0x1f1accb2.
//
// Solidity: function apy(uint8 ) view returns(uint256)
func (_StakeContract *StakeContractSession) Apy(arg0 uint8) (*big.Int, error) {
	return _StakeContract.Contract.Apy(&_StakeContract.CallOpts, arg0)
}

// Apy is a free data retrieval call binding the contract method 0x1f1accb2.
//
// Solidity: function apy(uint8 ) view returns(uint256)
func (_StakeContract *StakeContractCallerSession) Apy(arg0 uint8) (*big.Int, error) {
	return _StakeContract.Contract.Apy(&_StakeContract.CallOpts, arg0)
}

// Durations is a free data retrieval call binding the contract method 0x0ae355d3.
//
// Solidity: function durations(uint8 ) view returns(uint256)
func (_StakeContract *StakeContractCaller) Durations(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeContract.contract.Call(opts, &out, "durations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Durations is a free data retrieval call binding the contract method 0x0ae355d3.
//
// Solidity: function durations(uint8 ) view returns(uint256)
func (_StakeContract *StakeContractSession) Durations(arg0 uint8) (*big.Int, error) {
	return _StakeContract.Contract.Durations(&_StakeContract.CallOpts, arg0)
}

// Durations is a free data retrieval call binding the contract method 0x0ae355d3.
//
// Solidity: function durations(uint8 ) view returns(uint256)
func (_StakeContract *StakeContractCallerSession) Durations(arg0 uint8) (*big.Int, error) {
	return _StakeContract.Contract.Durations(&_StakeContract.CallOpts, arg0)
}

// StakeIdToOwner is a free data retrieval call binding the contract method 0x3f9d8950.
//
// Solidity: function stakeIdToOwner(uint256 ) view returns(address)
func (_StakeContract *StakeContractCaller) StakeIdToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeContract.contract.Call(opts, &out, "stakeIdToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeIdToOwner is a free data retrieval call binding the contract method 0x3f9d8950.
//
// Solidity: function stakeIdToOwner(uint256 ) view returns(address)
func (_StakeContract *StakeContractSession) StakeIdToOwner(arg0 *big.Int) (common.Address, error) {
	return _StakeContract.Contract.StakeIdToOwner(&_StakeContract.CallOpts, arg0)
}

// StakeIdToOwner is a free data retrieval call binding the contract method 0x3f9d8950.
//
// Solidity: function stakeIdToOwner(uint256 ) view returns(address)
func (_StakeContract *StakeContractCallerSession) StakeIdToOwner(arg0 *big.Int) (common.Address, error) {
	return _StakeContract.Contract.StakeIdToOwner(&_StakeContract.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakeContract *StakeContractCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeContract.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakeContract *StakeContractSession) StakingToken() (common.Address, error) {
	return _StakeContract.Contract.StakingToken(&_StakeContract.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakeContract *StakeContractCallerSession) StakingToken() (common.Address, error) {
	return _StakeContract.Contract.StakingToken(&_StakeContract.CallOpts)
}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256 stakeId, uint256 amount, uint256 startTime, uint256 endTime, uint256 rewardRate, bool isActive, uint8 period)
func (_StakeContract *StakeContractCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeId    *big.Int
	Amount     *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
	RewardRate *big.Int
	IsActive   bool
	Period     uint8
}, error) {
	var out []interface{}
	err := _StakeContract.contract.Call(opts, &out, "userStakes", arg0, arg1)

	outstruct := new(struct {
		StakeId    *big.Int
		Amount     *big.Int
		StartTime  *big.Int
		EndTime    *big.Int
		RewardRate *big.Int
		IsActive   bool
		Period     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Period = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256 stakeId, uint256 amount, uint256 startTime, uint256 endTime, uint256 rewardRate, bool isActive, uint8 period)
func (_StakeContract *StakeContractSession) UserStakes(arg0 common.Address, arg1 *big.Int) (struct {
	StakeId    *big.Int
	Amount     *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
	RewardRate *big.Int
	IsActive   bool
	Period     uint8
}, error) {
	return _StakeContract.Contract.UserStakes(&_StakeContract.CallOpts, arg0, arg1)
}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256 stakeId, uint256 amount, uint256 startTime, uint256 endTime, uint256 rewardRate, bool isActive, uint8 period)
func (_StakeContract *StakeContractCallerSession) UserStakes(arg0 common.Address, arg1 *big.Int) (struct {
	StakeId    *big.Int
	Amount     *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
	RewardRate *big.Int
	IsActive   bool
	Period     uint8
}, error) {
	return _StakeContract.Contract.UserStakes(&_StakeContract.CallOpts, arg0, arg1)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period) returns()
func (_StakeContract *StakeContractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, period uint8) (*types.Transaction, error) {
	return _StakeContract.contract.Transact(opts, "stake", amount, period)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period) returns()
func (_StakeContract *StakeContractSession) Stake(amount *big.Int, period uint8) (*types.Transaction, error) {
	return _StakeContract.Contract.Stake(&_StakeContract.TransactOpts, amount, period)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period) returns()
func (_StakeContract *StakeContractTransactorSession) Stake(amount *big.Int, period uint8) (*types.Transaction, error) {
	return _StakeContract.Contract.Stake(&_StakeContract.TransactOpts, amount, period)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId) returns()
func (_StakeContract *StakeContractTransactor) Withdraw(opts *bind.TransactOpts, stakeId *big.Int) (*types.Transaction, error) {
	return _StakeContract.contract.Transact(opts, "withdraw", stakeId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId) returns()
func (_StakeContract *StakeContractSession) Withdraw(stakeId *big.Int) (*types.Transaction, error) {
	return _StakeContract.Contract.Withdraw(&_StakeContract.TransactOpts, stakeId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId) returns()
func (_StakeContract *StakeContractTransactorSession) Withdraw(stakeId *big.Int) (*types.Transaction, error) {
	return _StakeContract.Contract.Withdraw(&_StakeContract.TransactOpts, stakeId)
}

// StakeContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakeContract contract.
type StakeContractStakedIterator struct {
	Event *StakeContractStaked // Event containing the contract specifics and raw log

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
func (it *StakeContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeContractStaked)
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
		it.Event = new(StakeContractStaked)
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
func (it *StakeContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeContractStaked represents a Staked event raised by the StakeContract contract.
type StakeContractStaked struct {
	User      common.Address
	StakeId   *big.Int
	Amount    *big.Int
	Period    uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xcc10169be2ad544347561e230939849af48d1714c052d7fe247d12f3decb4896.
//
// Solidity: event Staked(address indexed user, uint256 stakeId, uint256 amount, uint8 period, uint256 timestamp)
func (_StakeContract *StakeContractFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StakeContractStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeContract.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeContractStakedIterator{contract: _StakeContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xcc10169be2ad544347561e230939849af48d1714c052d7fe247d12f3decb4896.
//
// Solidity: event Staked(address indexed user, uint256 stakeId, uint256 amount, uint8 period, uint256 timestamp)
func (_StakeContract *StakeContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakeContractStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeContract.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeContractStaked)
				if err := _StakeContract.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xcc10169be2ad544347561e230939849af48d1714c052d7fe247d12f3decb4896.
//
// Solidity: event Staked(address indexed user, uint256 stakeId, uint256 amount, uint8 period, uint256 timestamp)
func (_StakeContract *StakeContractFilterer) ParseStaked(log types.Log) (*StakeContractStaked, error) {
	event := new(StakeContractStaked)
	if err := _StakeContract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the StakeContract contract.
type StakeContractWithdrawnIterator struct {
	Event *StakeContractWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeContractWithdrawn)
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
		it.Event = new(StakeContractWithdrawn)
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
func (it *StakeContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeContractWithdrawn represents a Withdrawn event raised by the StakeContract contract.
type StakeContractWithdrawn struct {
	User        common.Address
	StakeId     *big.Int
	Principal   *big.Int
	TotalAmount *big.Int
	StakeIndex  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x94ffd6b85c71b847775c89ef6496b93cee961bdc6ff827fd117f174f06f745ae.
//
// Solidity: event Withdrawn(address indexed user, uint256 stakeId, uint256 principal, uint256 totalAmount, uint256 stakeIndex)
func (_StakeContract *StakeContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*StakeContractWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeContract.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeContractWithdrawnIterator{contract: _StakeContract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x94ffd6b85c71b847775c89ef6496b93cee961bdc6ff827fd117f174f06f745ae.
//
// Solidity: event Withdrawn(address indexed user, uint256 stakeId, uint256 principal, uint256 totalAmount, uint256 stakeIndex)
func (_StakeContract *StakeContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeContractWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeContract.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeContractWithdrawn)
				if err := _StakeContract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x94ffd6b85c71b847775c89ef6496b93cee961bdc6ff827fd117f174f06f745ae.
//
// Solidity: event Withdrawn(address indexed user, uint256 stakeId, uint256 principal, uint256 totalAmount, uint256 stakeIndex)
func (_StakeContract *StakeContractFilterer) ParseWithdrawn(log types.Log) (*StakeContractWithdrawn, error) {
	event := new(StakeContractWithdrawn)
	if err := _StakeContract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
