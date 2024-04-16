package avsregistry

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsapkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IBLSApkRegistry"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type AvsRegistrySubscriber interface {
	SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractIBLSApkRegistryNewPubkeyRegistration, event.Subscription, error)
}

type AvsRegistryChainSubscriber struct {
	logger         logging.Logger
	blsApkRegistry blsapkreg.ContractIBLSApkRegistryFilters
}

// forces EthSubscriber to implement the chainio.Subscriber interface
var _ AvsRegistrySubscriber = (*AvsRegistryChainSubscriber)(nil)

func NewAvsRegistryChainSubscriber(
	blsApkRegistry blsapkreg.ContractIBLSApkRegistryFilters,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	return &AvsRegistryChainSubscriber{
		logger:         logger,
		blsApkRegistry: blsApkRegistry,
	}, nil
}

func BuildAvsRegistryChainSubscriber(
	blsApkRegistryAddr common.Address,
	ethWsClient eth.Client,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	blsapkreg, err := blsapkreg.NewContractIBLSApkRegistry(blsApkRegistryAddr, ethWsClient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create BLSApkRegistry contract"), err)
	}
	return NewAvsRegistryChainSubscriber(blsapkreg, logger)
}

func (s *AvsRegistryChainSubscriber) SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractIBLSApkRegistryNewPubkeyRegistration, event.Subscription, error) {
	newPubkeyRegistrationChan := make(chan *blsapkreg.ContractIBLSApkRegistryNewPubkeyRegistration)
	sub, err := s.blsApkRegistry.WatchNewPubkeyRegistration(
		&bind.WatchOpts{}, newPubkeyRegistrationChan, nil,
	)
	if err != nil {
		return nil, nil, types.WrapError(errors.New("Failed to subscribe to NewPubkeyRegistration events"), err)
	}
	return newPubkeyRegistrationChan, sub, nil
}
