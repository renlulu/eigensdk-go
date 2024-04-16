// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIBLSApkRegistry

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

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// ContractIBLSApkRegistryMetaData contains all meta data concerning the ContractIBLSApkRegistry contract.
var ContractIBLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHashAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.ApkUpdate\",\"components\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromPubkeyHash\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// ContractIBLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIBLSApkRegistryMetaData.ABI instead.
var ContractIBLSApkRegistryABI = ContractIBLSApkRegistryMetaData.ABI

// ContractIBLSApkRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractIBLSApkRegistryMethods interface {
	ContractIBLSApkRegistryCalls
	ContractIBLSApkRegistryTransacts
	ContractIBLSApkRegistryFilters
}

// ContractIBLSApkRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIBLSApkRegistryCalls interface {
	GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error)

	GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error)

	GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error)

	GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error)

	GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error)

	OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)
}

// ContractIBLSApkRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIBLSApkRegistryTransacts interface {
	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error)

	RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractIBLSApkRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIBLSApkRegistryFilters interface {
	FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractIBLSApkRegistryNewPubkeyRegistrationIterator, error)
	WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractIBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error)
	ParseNewPubkeyRegistration(log types.Log) (*ContractIBLSApkRegistryNewPubkeyRegistration, error)

	FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractIBLSApkRegistryOperatorAddedToQuorumsIterator, error)
	WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractIBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error)
	ParseOperatorAddedToQuorums(log types.Log) (*ContractIBLSApkRegistryOperatorAddedToQuorums, error)

	FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator, error)
	WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractIBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error)
	ParseOperatorRemovedFromQuorums(log types.Log) (*ContractIBLSApkRegistryOperatorRemovedFromQuorums, error)
}

// ContractIBLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type ContractIBLSApkRegistry struct {
	ContractIBLSApkRegistryCaller     // Read-only binding to the contract
	ContractIBLSApkRegistryTransactor // Write-only binding to the contract
	ContractIBLSApkRegistryFilterer   // Log filterer for contract events
}

// ContractIBLSApkRegistry implements the ContractIBLSApkRegistryMethods interface.
var _ ContractIBLSApkRegistryMethods = (*ContractIBLSApkRegistry)(nil)

// ContractIBLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIBLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIBLSApkRegistryCaller implements the ContractIBLSApkRegistryCalls interface.
var _ ContractIBLSApkRegistryCalls = (*ContractIBLSApkRegistryCaller)(nil)

// ContractIBLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIBLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIBLSApkRegistryTransactor implements the ContractIBLSApkRegistryTransacts interface.
var _ ContractIBLSApkRegistryTransacts = (*ContractIBLSApkRegistryTransactor)(nil)

// ContractIBLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIBLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIBLSApkRegistryFilterer implements the ContractIBLSApkRegistryFilters interface.
var _ ContractIBLSApkRegistryFilters = (*ContractIBLSApkRegistryFilterer)(nil)

// ContractIBLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIBLSApkRegistrySession struct {
	Contract     *ContractIBLSApkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractIBLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIBLSApkRegistryCallerSession struct {
	Contract *ContractIBLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractIBLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIBLSApkRegistryTransactorSession struct {
	Contract     *ContractIBLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractIBLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIBLSApkRegistryRaw struct {
	Contract *ContractIBLSApkRegistry // Generic contract binding to access the raw methods on
}

// ContractIBLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIBLSApkRegistryCallerRaw struct {
	Contract *ContractIBLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIBLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIBLSApkRegistryTransactorRaw struct {
	Contract *ContractIBLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIBLSApkRegistry creates a new instance of ContractIBLSApkRegistry, bound to a specific deployed contract.
func NewContractIBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*ContractIBLSApkRegistry, error) {
	contract, err := bindContractIBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistry{ContractIBLSApkRegistryCaller: ContractIBLSApkRegistryCaller{contract: contract}, ContractIBLSApkRegistryTransactor: ContractIBLSApkRegistryTransactor{contract: contract}, ContractIBLSApkRegistryFilterer: ContractIBLSApkRegistryFilterer{contract: contract}}, nil
}

// NewContractIBLSApkRegistryCaller creates a new read-only instance of ContractIBLSApkRegistry, bound to a specific deployed contract.
func NewContractIBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractIBLSApkRegistryCaller, error) {
	contract, err := bindContractIBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistryCaller{contract: contract}, nil
}

// NewContractIBLSApkRegistryTransactor creates a new write-only instance of ContractIBLSApkRegistry, bound to a specific deployed contract.
func NewContractIBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIBLSApkRegistryTransactor, error) {
	contract, err := bindContractIBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistryTransactor{contract: contract}, nil
}

// NewContractIBLSApkRegistryFilterer creates a new log filterer instance of ContractIBLSApkRegistry, bound to a specific deployed contract.
func NewContractIBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIBLSApkRegistryFilterer, error) {
	contract, err := bindContractIBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistryFilterer{contract: contract}, nil
}

// bindContractIBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindContractIBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIBLSApkRegistry.Contract.ContractIBLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.ContractIBLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.ContractIBLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIBLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getApk", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractIBLSApkRegistry.Contract.GetApk(&_ContractIBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractIBLSApkRegistry.Contract.GetApk(&_ContractIBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getApkHashAtBlockNumberAndIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractIBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractIBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractIBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractIBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getApkIndicesAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractIBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractIBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractIBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractIBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getApkUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSApkRegistryApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryApkUpdate)).(*IBLSApkRegistryApkUpdate)

	return out0, err

}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	return _ContractIBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractIBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryApkUpdate, error) {
	return _ContractIBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractIBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getOperatorFromPubkeyHash", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractIBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractIBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractIBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractIBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractIBLSApkRegistry.Contract.GetOperatorId(&_ContractIBLSApkRegistry.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractIBLSApkRegistry.Contract.GetOperatorId(&_ContractIBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractIBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractIBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractIBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractIBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _ContractIBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractIBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _ContractIBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractIBLSApkRegistry.CallOpts, operator)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractIBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractIBLSApkRegistry.CallOpts, pubkeyHash)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractIBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractIBLSApkRegistry.CallOpts, pubkeyHash)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIBLSApkRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractIBLSApkRegistry.Contract.RegistryCoordinator(&_ContractIBLSApkRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIBLSApkRegistry.Contract.RegistryCoordinator(&_ContractIBLSApkRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.DeregisterOperator(&_ContractIBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.DeregisterOperator(&_ContractIBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.InitializeQuorum(&_ContractIBLSApkRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.InitializeQuorum(&_ContractIBLSApkRegistry.TransactOpts, quorumNumber)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractIBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractIBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.RegisterOperator(&_ContractIBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIBLSApkRegistry.Contract.RegisterOperator(&_ContractIBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// ContractIBLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the ContractIBLSApkRegistry contract.
type ContractIBLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *ContractIBLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractIBLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIBLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(ContractIBLSApkRegistryNewPubkeyRegistration)
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
func (it *ContractIBLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIBLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIBLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the ContractIBLSApkRegistry contract.
type ContractIBLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractIBLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIBLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistryNewPubkeyRegistrationIterator{contract: _ContractIBLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractIBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractIBLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIBLSApkRegistryNewPubkeyRegistration)
				if err := _ContractIBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*ContractIBLSApkRegistryNewPubkeyRegistration, error) {
	event := new(ContractIBLSApkRegistryNewPubkeyRegistration)
	if err := _ContractIBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIBLSApkRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the ContractIBLSApkRegistry contract.
type ContractIBLSApkRegistryOperatorAddedToQuorumsIterator struct {
	Event *ContractIBLSApkRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *ContractIBLSApkRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIBLSApkRegistryOperatorAddedToQuorums)
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
		it.Event = new(ContractIBLSApkRegistryOperatorAddedToQuorums)
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
func (it *ContractIBLSApkRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIBLSApkRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIBLSApkRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the ContractIBLSApkRegistry contract.
type ContractIBLSApkRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractIBLSApkRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _ContractIBLSApkRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistryOperatorAddedToQuorumsIterator{contract: _ContractIBLSApkRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractIBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractIBLSApkRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIBLSApkRegistryOperatorAddedToQuorums)
				if err := _ContractIBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*ContractIBLSApkRegistryOperatorAddedToQuorums, error) {
	event := new(ContractIBLSApkRegistryOperatorAddedToQuorums)
	if err := _ContractIBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the ContractIBLSApkRegistry contract.
type ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *ContractIBLSApkRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIBLSApkRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(ContractIBLSApkRegistryOperatorRemovedFromQuorums)
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
func (it *ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIBLSApkRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the ContractIBLSApkRegistry contract.
type ContractIBLSApkRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _ContractIBLSApkRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractIBLSApkRegistryOperatorRemovedFromQuorumsIterator{contract: _ContractIBLSApkRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractIBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractIBLSApkRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIBLSApkRegistryOperatorRemovedFromQuorums)
				if err := _ContractIBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractIBLSApkRegistry *ContractIBLSApkRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*ContractIBLSApkRegistryOperatorRemovedFromQuorums, error) {
	event := new(ContractIBLSApkRegistryOperatorRemovedFromQuorums)
	if err := _ContractIBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
