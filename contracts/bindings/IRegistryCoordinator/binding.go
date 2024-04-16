// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIRegistryCoordinator

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IRegistryCoordinatorOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorInfo struct {
	OperatorId [32]byte
	Status     uint8
}

// IRegistryCoordinatorOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IRegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// ContractIRegistryCoordinatorMetaData contains all meta data concerning the ContractIRegistryCoordinator contract.
var ContractIRegistryCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ejectOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentQuorumBitmap\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorInfo\",\"components\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromId\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetParams\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapAtBlockNumberByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapHistoryLength\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapUpdateByIndex\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"components\":[{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBitmap\",\"type\":\"uint192\",\"internalType\":\"uint192\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIndexRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyRegistrationMessageHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumUpdateBlockNumber\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registries\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ChurnApproverUpdated\",\"inputs\":[{\"name\":\"prevChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newChurnApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EjectorUpdated\",\"inputs\":[{\"name\":\"prevEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newEjector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetParamsUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"operatorSetParams\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRegistryCoordinator.OperatorSetParam\",\"components\":[{\"name\":\"maxOperatorCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumBlockNumberUpdated\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ContractIRegistryCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIRegistryCoordinatorMetaData.ABI instead.
var ContractIRegistryCoordinatorABI = ContractIRegistryCoordinatorMetaData.ABI

// ContractIRegistryCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractIRegistryCoordinatorMethods interface {
	ContractIRegistryCoordinatorCalls
	ContractIRegistryCoordinatorTransacts
	ContractIRegistryCoordinatorFilters
}

// ContractIRegistryCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIRegistryCoordinatorCalls interface {
	BlsApkRegistry(opts *bind.CallOpts) (common.Address, error)

	GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error)

	GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error)

	GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error)

	GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error)

	GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error)

	GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error)

	GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error)

	IndexRegistry(opts *bind.CallOpts) (common.Address, error)

	NumRegistries(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	QuorumCount(opts *bind.CallOpts) (uint8, error)

	QuorumUpdateBlockNumber(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error)

	Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)
}

// ContractIRegistryCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIRegistryCoordinatorTransacts interface {
	EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractIRegistryCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIRegistryCoordinatorFilters interface {
	FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractIRegistryCoordinatorChurnApproverUpdatedIterator, error)
	WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error)
	ParseChurnApproverUpdated(log types.Log) (*ContractIRegistryCoordinatorChurnApproverUpdated, error)

	FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractIRegistryCoordinatorEjectorUpdatedIterator, error)
	WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorEjectorUpdated) (event.Subscription, error)
	ParseEjectorUpdated(log types.Log) (*ContractIRegistryCoordinatorEjectorUpdated, error)

	FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractIRegistryCoordinatorOperatorDeregisteredIterator, error)
	WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorDeregistered(log types.Log) (*ContractIRegistryCoordinatorOperatorDeregistered, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractIRegistryCoordinatorOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractIRegistryCoordinatorOperatorRegistered, error)

	FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator, error)
	WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseOperatorSetParamsUpdated(log types.Log) (*ContractIRegistryCoordinatorOperatorSetParamsUpdated, error)

	FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error)
	WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error)
	ParseQuorumBlockNumberUpdated(log types.Log) (*ContractIRegistryCoordinatorQuorumBlockNumberUpdated, error)
}

// ContractIRegistryCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractIRegistryCoordinator struct {
	ContractIRegistryCoordinatorCaller     // Read-only binding to the contract
	ContractIRegistryCoordinatorTransactor // Write-only binding to the contract
	ContractIRegistryCoordinatorFilterer   // Log filterer for contract events
}

// ContractIRegistryCoordinator implements the ContractIRegistryCoordinatorMethods interface.
var _ ContractIRegistryCoordinatorMethods = (*ContractIRegistryCoordinator)(nil)

// ContractIRegistryCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIRegistryCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIRegistryCoordinatorCaller implements the ContractIRegistryCoordinatorCalls interface.
var _ ContractIRegistryCoordinatorCalls = (*ContractIRegistryCoordinatorCaller)(nil)

// ContractIRegistryCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIRegistryCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIRegistryCoordinatorTransactor implements the ContractIRegistryCoordinatorTransacts interface.
var _ ContractIRegistryCoordinatorTransacts = (*ContractIRegistryCoordinatorTransactor)(nil)

// ContractIRegistryCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIRegistryCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIRegistryCoordinatorFilterer implements the ContractIRegistryCoordinatorFilters interface.
var _ ContractIRegistryCoordinatorFilters = (*ContractIRegistryCoordinatorFilterer)(nil)

// ContractIRegistryCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIRegistryCoordinatorSession struct {
	Contract     *ContractIRegistryCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractIRegistryCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIRegistryCoordinatorCallerSession struct {
	Contract *ContractIRegistryCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ContractIRegistryCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIRegistryCoordinatorTransactorSession struct {
	Contract     *ContractIRegistryCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ContractIRegistryCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIRegistryCoordinatorRaw struct {
	Contract *ContractIRegistryCoordinator // Generic contract binding to access the raw methods on
}

// ContractIRegistryCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIRegistryCoordinatorCallerRaw struct {
	Contract *ContractIRegistryCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIRegistryCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIRegistryCoordinatorTransactorRaw struct {
	Contract *ContractIRegistryCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIRegistryCoordinator creates a new instance of ContractIRegistryCoordinator, bound to a specific deployed contract.
func NewContractIRegistryCoordinator(address common.Address, backend bind.ContractBackend) (*ContractIRegistryCoordinator, error) {
	contract, err := bindContractIRegistryCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinator{ContractIRegistryCoordinatorCaller: ContractIRegistryCoordinatorCaller{contract: contract}, ContractIRegistryCoordinatorTransactor: ContractIRegistryCoordinatorTransactor{contract: contract}, ContractIRegistryCoordinatorFilterer: ContractIRegistryCoordinatorFilterer{contract: contract}}, nil
}

// NewContractIRegistryCoordinatorCaller creates a new read-only instance of ContractIRegistryCoordinator, bound to a specific deployed contract.
func NewContractIRegistryCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractIRegistryCoordinatorCaller, error) {
	contract, err := bindContractIRegistryCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorCaller{contract: contract}, nil
}

// NewContractIRegistryCoordinatorTransactor creates a new write-only instance of ContractIRegistryCoordinator, bound to a specific deployed contract.
func NewContractIRegistryCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIRegistryCoordinatorTransactor, error) {
	contract, err := bindContractIRegistryCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorTransactor{contract: contract}, nil
}

// NewContractIRegistryCoordinatorFilterer creates a new log filterer instance of ContractIRegistryCoordinator, bound to a specific deployed contract.
func NewContractIRegistryCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIRegistryCoordinatorFilterer, error) {
	contract, err := bindContractIRegistryCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorFilterer{contract: contract}, nil
}

// bindContractIRegistryCoordinator binds a generic wrapper to an already deployed contract.
func bindContractIRegistryCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIRegistryCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIRegistryCoordinator.Contract.ContractIRegistryCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.Contract.ContractIRegistryCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.Contract.ContractIRegistryCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIRegistryCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.Contract.contract.Transact(opts, method, params...)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.BlsApkRegistry(&_ContractIRegistryCoordinator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.BlsApkRegistry(&_ContractIRegistryCoordinator.CallOpts)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetCurrentQuorumBitmap(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getCurrentQuorumBitmap", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractIRegistryCoordinator.CallOpts, operatorId)
}

// GetCurrentQuorumBitmap is a free data retrieval call binding the contract method 0x871ef049.
//
// Solidity: function getCurrentQuorumBitmap(bytes32 operatorId) view returns(uint192)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetCurrentQuorumBitmap(operatorId [32]byte) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.GetCurrentQuorumBitmap(&_ContractIRegistryCoordinator.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorInfo)).(*IRegistryCoordinatorOperatorInfo)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperator(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperatorInfo, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperator(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address operator)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetOperatorFromId(opts *bind.CallOpts, operatorId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getOperatorFromId", operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address operator)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorFromId(&_ContractIRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorFromId is a free data retrieval call binding the contract method 0x296bb064.
//
// Solidity: function getOperatorFromId(bytes32 operatorId) view returns(address operator)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetOperatorFromId(operatorId [32]byte) (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorFromId(&_ContractIRegistryCoordinator.CallOpts, operatorId)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorId(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorId(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IRegistryCoordinatorOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperatorSetParam)).(*IRegistryCoordinatorOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractIRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetOperatorSetParams(quorumNumber uint8) (IRegistryCoordinatorOperatorSetParam, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorSetParams(&_ContractIRegistryCoordinator.CallOpts, quorumNumber)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorStatus(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ContractIRegistryCoordinator.Contract.GetOperatorStatus(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetQuorumBitmapAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractIRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x04ec6351.
//
// Solidity: function getQuorumBitmapAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetQuorumBitmapAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapAtBlockNumberByIndex(&_ContractIRegistryCoordinator.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetQuorumBitmapHistoryLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapHistoryLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractIRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapHistoryLength is a free data retrieval call binding the contract method 0x03fd3492.
//
// Solidity: function getQuorumBitmapHistoryLength(bytes32 operatorId) view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetQuorumBitmapHistoryLength(operatorId [32]byte) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapHistoryLength(&_ContractIRegistryCoordinator.CallOpts, operatorId)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetQuorumBitmapIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapIndicesAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractIRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xc391425e.
//
// Solidity: function getQuorumBitmapIndicesAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetQuorumBitmapIndicesAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapIndicesAtBlockNumber(&_ContractIRegistryCoordinator.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) GetQuorumBitmapUpdateByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "getQuorumBitmapUpdateByIndex", operatorId, index)

	if err != nil {
		return *new(IRegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorQuorumBitmapUpdate)).(*IRegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractIRegistryCoordinator.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByIndex is a free data retrieval call binding the contract method 0x1eb812da.
//
// Solidity: function getQuorumBitmapUpdateByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) GetQuorumBitmapUpdateByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _ContractIRegistryCoordinator.Contract.GetQuorumBitmapUpdateByIndex(&_ContractIRegistryCoordinator.CallOpts, operatorId, index)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) IndexRegistry() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.IndexRegistry(&_ContractIRegistryCoordinator.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) IndexRegistry() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.IndexRegistry(&_ContractIRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) NumRegistries() (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.NumRegistries(&_ContractIRegistryCoordinator.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) NumRegistries() (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.NumRegistries(&_ContractIRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) Owner() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.Owner(&_ContractIRegistryCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.Owner(&_ContractIRegistryCoordinator.CallOpts)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) PubkeyRegistrationMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "pubkeyRegistrationMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractIRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// PubkeyRegistrationMessageHash is a free data retrieval call binding the contract method 0x3c2a7f4c.
//
// Solidity: function pubkeyRegistrationMessageHash(address operator) view returns((uint256,uint256))
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) PubkeyRegistrationMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractIRegistryCoordinator.Contract.PubkeyRegistrationMessageHash(&_ContractIRegistryCoordinator.CallOpts, operator)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) QuorumCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "quorumCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) QuorumCount() (uint8, error) {
	return _ContractIRegistryCoordinator.Contract.QuorumCount(&_ContractIRegistryCoordinator.CallOpts)
}

// QuorumCount is a free data retrieval call binding the contract method 0x9aa1653d.
//
// Solidity: function quorumCount() view returns(uint8)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) QuorumCount() (uint8, error) {
	return _ContractIRegistryCoordinator.Contract.QuorumCount(&_ContractIRegistryCoordinator.CallOpts)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 quorumNumber) view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) QuorumUpdateBlockNumber(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "quorumUpdateBlockNumber", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 quorumNumber) view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) QuorumUpdateBlockNumber(quorumNumber uint8) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractIRegistryCoordinator.CallOpts, quorumNumber)
}

// QuorumUpdateBlockNumber is a free data retrieval call binding the contract method 0x249a0c42.
//
// Solidity: function quorumUpdateBlockNumber(uint8 quorumNumber) view returns(uint256)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) QuorumUpdateBlockNumber(quorumNumber uint8) (*big.Int, error) {
	return _ContractIRegistryCoordinator.Contract.QuorumUpdateBlockNumber(&_ContractIRegistryCoordinator.CallOpts, quorumNumber)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.Registries(&_ContractIRegistryCoordinator.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.Registries(&_ContractIRegistryCoordinator.CallOpts, arg0)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIRegistryCoordinator.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) StakeRegistry() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.StakeRegistry(&_ContractIRegistryCoordinator.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIRegistryCoordinator.Contract.StakeRegistry(&_ContractIRegistryCoordinator.CallOpts)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorTransactor) EjectOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.contract.Transact(opts, "ejectOperator", operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.Contract.EjectOperator(&_ContractIRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// EjectOperator is a paid mutator transaction binding the contract method 0x6e3b17db.
//
// Solidity: function ejectOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorTransactorSession) EjectOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIRegistryCoordinator.Contract.EjectOperator(&_ContractIRegistryCoordinator.TransactOpts, operator, quorumNumbers)
}

// ContractIRegistryCoordinatorChurnApproverUpdatedIterator is returned from FilterChurnApproverUpdated and is used to iterate over the raw logs and unpacked data for ChurnApproverUpdated events raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorChurnApproverUpdatedIterator struct {
	Event *ContractIRegistryCoordinatorChurnApproverUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIRegistryCoordinatorChurnApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRegistryCoordinatorChurnApproverUpdated)
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
		it.Event = new(ContractIRegistryCoordinatorChurnApproverUpdated)
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
func (it *ContractIRegistryCoordinatorChurnApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRegistryCoordinatorChurnApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRegistryCoordinatorChurnApproverUpdated represents a ChurnApproverUpdated event raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorChurnApproverUpdated struct {
	PrevChurnApprover common.Address
	NewChurnApprover  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterChurnApproverUpdated is a free log retrieval operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) FilterChurnApproverUpdated(opts *bind.FilterOpts) (*ContractIRegistryCoordinatorChurnApproverUpdatedIterator, error) {

	logs, sub, err := _ContractIRegistryCoordinator.contract.FilterLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorChurnApproverUpdatedIterator{contract: _ContractIRegistryCoordinator.contract, event: "ChurnApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchChurnApproverUpdated is a free log subscription operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) WatchChurnApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorChurnApproverUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIRegistryCoordinator.contract.WatchLogs(opts, "ChurnApproverUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRegistryCoordinatorChurnApproverUpdated)
				if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
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

// ParseChurnApproverUpdated is a log parse operation binding the contract event 0x315457d8a8fe60f04af17c16e2f5a5e1db612b31648e58030360759ef8f3528c.
//
// Solidity: event ChurnApproverUpdated(address prevChurnApprover, address newChurnApprover)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) ParseChurnApproverUpdated(log types.Log) (*ContractIRegistryCoordinatorChurnApproverUpdated, error) {
	event := new(ContractIRegistryCoordinatorChurnApproverUpdated)
	if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "ChurnApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRegistryCoordinatorEjectorUpdatedIterator is returned from FilterEjectorUpdated and is used to iterate over the raw logs and unpacked data for EjectorUpdated events raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorEjectorUpdatedIterator struct {
	Event *ContractIRegistryCoordinatorEjectorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIRegistryCoordinatorEjectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRegistryCoordinatorEjectorUpdated)
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
		it.Event = new(ContractIRegistryCoordinatorEjectorUpdated)
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
func (it *ContractIRegistryCoordinatorEjectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRegistryCoordinatorEjectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRegistryCoordinatorEjectorUpdated represents a EjectorUpdated event raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorEjectorUpdated struct {
	PrevEjector common.Address
	NewEjector  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEjectorUpdated is a free log retrieval operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) FilterEjectorUpdated(opts *bind.FilterOpts) (*ContractIRegistryCoordinatorEjectorUpdatedIterator, error) {

	logs, sub, err := _ContractIRegistryCoordinator.contract.FilterLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorEjectorUpdatedIterator{contract: _ContractIRegistryCoordinator.contract, event: "EjectorUpdated", logs: logs, sub: sub}, nil
}

// WatchEjectorUpdated is a free log subscription operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) WatchEjectorUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorEjectorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIRegistryCoordinator.contract.WatchLogs(opts, "EjectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRegistryCoordinatorEjectorUpdated)
				if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
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

// ParseEjectorUpdated is a log parse operation binding the contract event 0x8f30ab09f43a6c157d7fce7e0a13c003042c1c95e8a72e7a146a21c0caa24dc9.
//
// Solidity: event EjectorUpdated(address prevEjector, address newEjector)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) ParseEjectorUpdated(log types.Log) (*ContractIRegistryCoordinatorEjectorUpdated, error) {
	event := new(ContractIRegistryCoordinatorEjectorUpdated)
	if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "EjectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRegistryCoordinatorOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorOperatorDeregisteredIterator struct {
	Event *ContractIRegistryCoordinatorOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractIRegistryCoordinatorOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRegistryCoordinatorOperatorDeregistered)
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
		it.Event = new(ContractIRegistryCoordinatorOperatorDeregistered)
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
func (it *ContractIRegistryCoordinatorOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRegistryCoordinatorOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRegistryCoordinatorOperatorDeregistered represents a OperatorDeregistered event raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractIRegistryCoordinatorOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorOperatorDeregisteredIterator{contract: _ContractIRegistryCoordinator.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRegistryCoordinatorOperatorDeregistered)
				if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) ParseOperatorDeregistered(log types.Log) (*ContractIRegistryCoordinatorOperatorDeregistered, error) {
	event := new(ContractIRegistryCoordinatorOperatorDeregistered)
	if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRegistryCoordinatorOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorOperatorRegisteredIterator struct {
	Event *ContractIRegistryCoordinatorOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ContractIRegistryCoordinatorOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRegistryCoordinatorOperatorRegistered)
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
		it.Event = new(ContractIRegistryCoordinatorOperatorRegistered)
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
func (it *ContractIRegistryCoordinatorOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRegistryCoordinatorOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRegistryCoordinatorOperatorRegistered represents a OperatorRegistered event raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*ContractIRegistryCoordinatorOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorOperatorRegisteredIterator{contract: _ContractIRegistryCoordinator.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRegistryCoordinatorOperatorRegistered)
				if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) ParseOperatorRegistered(log types.Log) (*ContractIRegistryCoordinatorOperatorRegistered, error) {
	event := new(ContractIRegistryCoordinatorOperatorRegistered)
	if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator struct {
	Event *ContractIRegistryCoordinatorOperatorSetParamsUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRegistryCoordinatorOperatorSetParamsUpdated)
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
		it.Event = new(ContractIRegistryCoordinatorOperatorSetParamsUpdated)
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
func (it *ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRegistryCoordinatorOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IRegistryCoordinatorOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorOperatorSetParamsUpdatedIterator{contract: _ContractIRegistryCoordinator.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRegistryCoordinatorOperatorSetParamsUpdated)
				if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
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

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*ContractIRegistryCoordinatorOperatorSetParamsUpdated, error) {
	event := new(ContractIRegistryCoordinatorOperatorSetParamsUpdated)
	if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator is returned from FilterQuorumBlockNumberUpdated and is used to iterate over the raw logs and unpacked data for QuorumBlockNumberUpdated events raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator struct {
	Event *ContractIRegistryCoordinatorQuorumBlockNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIRegistryCoordinatorQuorumBlockNumberUpdated)
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
		it.Event = new(ContractIRegistryCoordinatorQuorumBlockNumberUpdated)
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
func (it *ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIRegistryCoordinatorQuorumBlockNumberUpdated represents a QuorumBlockNumberUpdated event raised by the ContractIRegistryCoordinator contract.
type ContractIRegistryCoordinatorQuorumBlockNumberUpdated struct {
	QuorumNumber uint8
	Blocknumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumBlockNumberUpdated is a free log retrieval operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) FilterQuorumBlockNumberUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.FilterLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &ContractIRegistryCoordinatorQuorumBlockNumberUpdatedIterator{contract: _ContractIRegistryCoordinator.contract, event: "QuorumBlockNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumBlockNumberUpdated is a free log subscription operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) WatchQuorumBlockNumberUpdated(opts *bind.WatchOpts, sink chan<- *ContractIRegistryCoordinatorQuorumBlockNumberUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _ContractIRegistryCoordinator.contract.WatchLogs(opts, "QuorumBlockNumberUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIRegistryCoordinatorQuorumBlockNumberUpdated)
				if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
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

// ParseQuorumBlockNumberUpdated is a log parse operation binding the contract event 0x46077d55330763f16269fd75e5761663f4192d2791747c0189b16ad31db07db4.
//
// Solidity: event QuorumBlockNumberUpdated(uint8 indexed quorumNumber, uint256 blocknumber)
func (_ContractIRegistryCoordinator *ContractIRegistryCoordinatorFilterer) ParseQuorumBlockNumberUpdated(log types.Log) (*ContractIRegistryCoordinatorQuorumBlockNumberUpdated, error) {
	event := new(ContractIRegistryCoordinatorQuorumBlockNumberUpdated)
	if err := _ContractIRegistryCoordinator.contract.UnpackLog(event, "QuorumBlockNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
