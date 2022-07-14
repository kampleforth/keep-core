package ethereum

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"math/rand"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	beaconchain "github.com/keep-network/keep-core/pkg/beacon/chain"
	"github.com/keep-network/keep-core/pkg/beacon/event"
	"github.com/keep-network/keep-core/pkg/subscription"

	"github.com/ethereum/go-ethereum/common"
	"github.com/keep-network/keep-common/pkg/chain/ethereum"
	"github.com/keep-network/keep-core/pkg/chain"
	"github.com/keep-network/keep-core/pkg/chain/ethereum/beacon/gen/contract"
)

// Definitions of contract names.
const (
	RandomBeaconContractName = "RandomBeacon"
)

// BeaconChain represents a beacon-specific chain handle.
type BeaconChain struct {
	*Chain

	randomBeacon  *contract.RandomBeacon
	sortitionPool *contract.BeaconSortitionPool

	mockRandomBeacon *mockRandomBeacon
}

// newBeaconChain construct a new instance of the beacon-specific Ethereum
// chain handle.
func newBeaconChain(
	config ethereum.Config,
	baseChain *Chain,
) (*BeaconChain, error) {
	randomBeaconAddress, err := config.ContractAddress(RandomBeaconContractName)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to resolve %s contract address: [%v]",
			RandomBeaconContractName,
			err,
		)
	}

	randomBeacon, err :=
		contract.NewRandomBeacon(
			randomBeaconAddress,
			baseChain.chainID,
			baseChain.key,
			baseChain.client,
			baseChain.nonceManager,
			baseChain.miningWaiter,
			baseChain.blockCounter,
			baseChain.transactionMutex,
		)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to attach to RandomBeacon contract: [%v]",
			err,
		)
	}

	sortitionPoolAddress, err := randomBeacon.SortitionPool()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get sortition pool address: [%v]",
			err,
		)
	}

	sortitionPool, err :=
		contract.NewBeaconSortitionPool(
			sortitionPoolAddress,
			baseChain.chainID,
			baseChain.key,
			baseChain.client,
			baseChain.nonceManager,
			baseChain.miningWaiter,
			baseChain.blockCounter,
			baseChain.transactionMutex,
		)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to attach to BeaconSortitionPool contract: [%v]",
			err,
		)
	}

	return &BeaconChain{
		Chain:            baseChain,
		randomBeacon:     randomBeacon,
		sortitionPool:    sortitionPool,
		mockRandomBeacon: newMockRandomBeacon(baseChain.blockCounter),
	}, nil
}

// GetConfig returns the expected configuration of the random beacon.
// TODO: Adjust to the random beacon v2 requirements.
func (bc *BeaconChain) GetConfig() *beaconchain.Config {
	groupSize := 64
	honestThreshold := 33
	resultPublicationBlockStep := 6
	relayEntryTimeout := groupSize * resultPublicationBlockStep

	return &beaconchain.Config{
		GroupSize:                  groupSize,
		HonestThreshold:            honestThreshold,
		ResultPublicationBlockStep: uint64(resultPublicationBlockStep),
		RelayEntryTimeout:          uint64(relayEntryTimeout),
	}
}

// OperatorToStakingProvider returns the staking provider address for the
// current operator. If the staking provider has not been registered for the
// operator, the returned address is empty and the boolean flag is set to false
// If the staking provider has been registered, the address is not empty and the
// boolean flag indicates true.
func (bc *BeaconChain) OperatorToStakingProvider() (chain.Address, bool, error) {
	stakingProvider, err := bc.randomBeacon.OperatorToStakingProvider(bc.key.Address)
	if err != nil {
		return "", false, fmt.Errorf(
			"failed to map operator %v to a staking provider: [%v]",
			bc.key.Address,
			err,
		)
	}

	if bytes.Equal(
		stakingProvider.Bytes(),
		bytes.Repeat([]byte{0}, common.AddressLength),
	) {
		return "", false, nil
	}

	return chain.Address(stakingProvider.Hex()), true, nil
}

// EligibleStake returns the current value of the staking provider's eligible
// stake. Eligible stake is defined as the currently authorized stake minus the
// pending authorization decrease. Eligible stake is what is used for operator's
// weight in the sortition pool. If the authorized stake minus the pending
// authorization decrease is below the minimum authorization, eligible stake
// is 0.
func (bc *BeaconChain) EligibleStake(stakingProvider chain.Address) (*big.Int, error) {
	eligibleStake, err := bc.randomBeacon.EligibleStake(common.HexToAddress(stakingProvider.String()))
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get eligible stake for staking provider %s: [%w]",
			stakingProvider,
			err,
		)
	}

	return eligibleStake, nil
}

// IsPoolLocked returns true if the sortition pool is locked and no state
// changes are allowed.
func (bc *BeaconChain) IsPoolLocked() (bool, error) {
	return bc.sortitionPool.IsLocked()
}

// IsOperatorInPool returns true if the current operator is registered in the
// sortition pool.
func (bc *BeaconChain) IsOperatorInPool() (bool, error) {
	return bc.randomBeacon.IsOperatorInPool(bc.key.Address)
}

// IsOperatorUpToDate checks if the operator's authorized stake is in sync
// with operator's weight in the sortition pool.
// If the operator's authorized stake is not in sync with sortition pool
// weight, function returns false.
// If the operator is not in the sortition pool and their authorized stake
// is non-zero, function returns false.
func (bc *BeaconChain) IsOperatorUpToDate() (bool, error) {
	return bc.randomBeacon.IsOperatorUpToDate(bc.key.Address)
}

// JoinSortitionPool executes a transaction to have the current operator join
// the sortition pool.
func (bc *BeaconChain) JoinSortitionPool() error {
	_, err := bc.randomBeacon.JoinSortitionPool()
	return err
}

// UpdateOperatorStatus executes a transaction to update the current
// operator's state in the sortition pool.
func (bc *BeaconChain) UpdateOperatorStatus() error {
	_, err := bc.randomBeacon.UpdateOperatorStatus(bc.key.Address)
	return err
}

// SelectGroup returns the group members for the group generated by
// the given seed. This function can return an error if the beacon chain's
// state does not allow for group selection at the moment.
func (bc *BeaconChain) SelectGroup(seed *big.Int) ([]chain.Address, error) {
	groupSize := big.NewInt(int64(bc.GetConfig().GroupSize))
	seedBytes := [32]byte{}
	seed.FillBytes(seedBytes[:])

	// TODO: Replace with a call to the RandomBeacon.selectGroup function.
	operatorsIDs, err := bc.sortitionPool.SelectGroup(groupSize, seedBytes)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot select group in the sortition pool: [%v]",
			err,
		)
	}

	operatorsAddresses, err := bc.sortitionPool.GetIDOperators(operatorsIDs)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot convert operators' IDs to addresses: [%v]",
			err,
		)
	}

	result := make([]chain.Address, len(operatorsAddresses))
	for i := range result {
		result[i] = chain.Address(operatorsAddresses[i].String())
	}

	return result, nil
}

// TODO: Implement a real OnGroupRegistered function.
func (bc *BeaconChain) OnGroupRegistered(
	handler func(groupRegistration *event.GroupRegistration),
) subscription.EventSubscription {
	return subscription.NewEventSubscription(func() {})
}

// TODO: Implement a real IsGroupRegistered function.
func (bc *BeaconChain) IsGroupRegistered(groupPublicKey []byte) (bool, error) {
	return false, nil
}

// TODO: Implement a real IsStaleGroup function.
func (bc *BeaconChain) IsStaleGroup(groupPublicKey []byte) (bool, error) {
	return false, nil
}

// TODO: Implement a real OnDKGStarted event subscription. The current
//       implementation generates a fake event every 500th block where the
//       seed is the keccak256 of the block number.
func (bc *BeaconChain) OnDKGStarted(
	handler func(event *event.DKGStarted),
) subscription.EventSubscription {
	return bc.mockRandomBeacon.OnDKGStarted(handler)
}

// TODO: Implement a real SubmitDKGResult action. The current implementation
//       just creates and pipes the DKG submission event to the handlers
//       registered in the dkgResultSubmissionHandlers map. Consider getting
//       rid of the result promise in favor of the fire-and-forget style.
func (bc *BeaconChain) SubmitDKGResult(
	participantIndex beaconchain.GroupMemberIndex,
	dkgResult *beaconchain.DKGResult,
	signatures map[beaconchain.GroupMemberIndex][]byte,
) error {
	return bc.mockRandomBeacon.SubmitDKGResult(
		participantIndex,
		dkgResult,
		signatures,
	)
}

// TODO: Implement a real OnDKGResultSubmitted event subscription. The current
//       implementation just pipes the DKG submission event generated within
//       SubmitDKGResult to the handlers registered in the
//       dkgResultSubmissionHandlers map.
func (bc *BeaconChain) OnDKGResultSubmitted(
	handler func(event *event.DKGResultSubmission),
) subscription.EventSubscription {
	return bc.mockRandomBeacon.OnDKGResultSubmitted(handler)
}

// CalculateDKGResultHash calculates Keccak-256 hash of the DKG result. Operation
// is performed off-chain.
//
// It first encodes the result using solidity ABI and then calculates Keccak-256
// hash over it. This corresponds to the DKG result hash calculation on-chain.
// Hashes calculated off-chain and on-chain must always match.
func (bc *BeaconChain) CalculateDKGResultHash(
	dkgResult *beaconchain.DKGResult,
) (beaconchain.DKGResultHash, error) {
	// Encode DKG result to the format matched with Solidity keccak256(abi.encodePacked(...))
	hash := crypto.Keccak256(dkgResult.GroupPublicKey, dkgResult.Misbehaved)
	return beaconchain.DKGResultHashFromBytes(hash)
}

// TODO: Implement a real SubmitRelayEntry function.
func (bc *BeaconChain) SubmitRelayEntry(
	entry []byte,
) error {
	return bc.mockRandomBeacon.SubmitRelayEntry(entry)
}

// TODO: Implement a real OnRelayEntrySubmitted function.
func (bc *BeaconChain) OnRelayEntrySubmitted(
	handler func(entry *event.RelayEntrySubmitted),
) subscription.EventSubscription {
	return bc.mockRandomBeacon.OnRelayEntrySubmitted(handler)
}

// TODO: Implement a real OnRelayEntryRequested function.
func (bc *BeaconChain) OnRelayEntryRequested(
	handler func(request *event.RelayEntryRequested),
) subscription.EventSubscription {
	return bc.mockRandomBeacon.OnRelayEntryRequested(handler)
}

// TODO: Implement a real ReportRelayEntryTimeout function.
func (bc *BeaconChain) ReportRelayEntryTimeout() error {
	return nil
}

// TODO: Implement a real IsEntryInProgress function.
func (bc *BeaconChain) IsEntryInProgress() (bool, error) {
	return bc.mockRandomBeacon.IsEntryInProgress()
}

// TODO: Implement a real CurrentRequestStartBlock function.
func (bc *BeaconChain) CurrentRequestStartBlock() (*big.Int, error) {
	return bc.mockRandomBeacon.CurrentRequestStartBlock()
}

// TODO: Implement a real CurrentRequestPreviousEntry function.
func (bc *BeaconChain) CurrentRequestPreviousEntry() ([]byte, error) {
	return bc.mockRandomBeacon.CurrentRequestPreviousEntry()
}

// TODO: Implement a real CurrentRequestGroupPublicKey function.
func (bc *BeaconChain) CurrentRequestGroupPublicKey() ([]byte, error) {
	return bc.mockRandomBeacon.CurrentRequestGroupPublicKey()
}

// TODO: Temporary mock that simulates the behavior of the RandomBeacon
//       contract. Should be removed eventually.
type mockRandomBeacon struct {
	blockCounter chain.BlockCounter

	dkgResultSubmissionHandlersMutex sync.Mutex
	dkgResultSubmissionHandlers      map[int]func(submission *event.DKGResultSubmission)

	activeGroupMutex sync.RWMutex
	activeGroup      []byte

	latestRelayEntryMutex sync.RWMutex
	latestRelayEntry      []byte

	currentRequestMutex         sync.RWMutex
	currentRequestStartBlock    *big.Int
	currentRequestPreviousEntry []byte
	currentRequestGroup         []byte

	relayEntrySubmissionHandlersMutex sync.Mutex
	relayEntrySubmissionHandlers      map[int]func(submission *event.RelayEntrySubmitted)
}

func newMockRandomBeacon(blockCounter chain.BlockCounter) *mockRandomBeacon {
	return &mockRandomBeacon{
		blockCounter:                 blockCounter,
		dkgResultSubmissionHandlers:  make(map[int]func(submission *event.DKGResultSubmission)),
		latestRelayEntry:             big.NewInt(1000000000).Bytes(),
		relayEntrySubmissionHandlers: make(map[int]func(submission *event.RelayEntrySubmitted)),
	}
}

func (mrb *mockRandomBeacon) OnDKGStarted(
	handler func(event *event.DKGStarted),
) subscription.EventSubscription {
	ctx, cancelCtx := context.WithCancel(context.Background())
	blocksChan := mrb.blockCounter.WatchBlocks(ctx)

	go func() {
		for {
			select {
			case block := <-blocksChan:
				// Generate an event every 500th block.
				if block%500 == 0 {
					// The seed is keccak256(block).
					blockBytes := make([]byte, 8)
					binary.BigEndian.PutUint64(blockBytes, block)
					seedBytes := crypto.Keccak256(blockBytes)
					seed := new(big.Int).SetBytes(seedBytes)

					handler(&event.DKGStarted{
						Seed:        seed,
						BlockNumber: block,
					})
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return subscription.NewEventSubscription(func() {
		cancelCtx()
	})
}

func (mrb *mockRandomBeacon) OnDKGResultSubmitted(
	handler func(event *event.DKGResultSubmission),
) subscription.EventSubscription {
	mrb.dkgResultSubmissionHandlersMutex.Lock()
	defer mrb.dkgResultSubmissionHandlersMutex.Unlock()

	// #nosec G404 (insecure random number source (rand))
	// Temporary test implementation doesn't require secure randomness.
	handlerID := rand.Int()

	mrb.dkgResultSubmissionHandlers[handlerID] = handler

	return subscription.NewEventSubscription(func() {
		mrb.dkgResultSubmissionHandlersMutex.Lock()
		defer mrb.dkgResultSubmissionHandlersMutex.Unlock()

		delete(mrb.dkgResultSubmissionHandlers, handlerID)
	})
}

func (mrb *mockRandomBeacon) SubmitDKGResult(
	participantIndex beaconchain.GroupMemberIndex,
	dkgResult *beaconchain.DKGResult,
	signatures map[beaconchain.GroupMemberIndex][]byte,
) error {
	mrb.dkgResultSubmissionHandlersMutex.Lock()
	defer mrb.dkgResultSubmissionHandlersMutex.Unlock()

	mrb.activeGroupMutex.Lock()
	defer mrb.activeGroupMutex.Unlock()

	blockNumber, err := mrb.blockCounter.CurrentBlock()
	if err != nil {
		return fmt.Errorf("failed to get the current block")
	}

	for _, handler := range mrb.dkgResultSubmissionHandlers {
		go func(handler func(*event.DKGResultSubmission)) {
			handler(&event.DKGResultSubmission{
				MemberIndex:    uint32(participantIndex),
				GroupPublicKey: dkgResult.GroupPublicKey,
				Misbehaved:     dkgResult.Misbehaved,
				BlockNumber:    blockNumber,
			})
		}(handler)
	}

	mrb.activeGroup = dkgResult.GroupPublicKey

	return nil
}

func (mrb *mockRandomBeacon) OnRelayEntryRequested(
	handler func(request *event.RelayEntryRequested),
) subscription.EventSubscription {
	ctx, cancelCtx := context.WithCancel(context.Background())
	blocksChan := mrb.blockCounter.WatchBlocks(ctx)

	go func() {
		for {
			select {
			case block := <-blocksChan:
				// Generate an event every 25th block.
				if block%25 == 0 {
					mrb.activeGroupMutex.RLock()
					mrb.latestRelayEntryMutex.RLock()
					mrb.currentRequestMutex.Lock()

					if mrb.currentRequestStartBlock == nil && len(mrb.activeGroup) > 0 {
						mrb.currentRequestStartBlock = big.NewInt(int64(block))
						mrb.currentRequestPreviousEntry = mrb.latestRelayEntry
						mrb.currentRequestGroup = mrb.activeGroup

						go handler(&event.RelayEntryRequested{
							PreviousEntry:  mrb.currentRequestPreviousEntry,
							GroupPublicKey: mrb.currentRequestGroup,
							BlockNumber:    mrb.currentRequestStartBlock.Uint64(),
						})
					}

					mrb.currentRequestMutex.Unlock()
					mrb.latestRelayEntryMutex.RUnlock()
					mrb.activeGroupMutex.RUnlock()
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return subscription.NewEventSubscription(func() {
		cancelCtx()
	})
}

func (mrb *mockRandomBeacon) OnRelayEntrySubmitted(
	handler func(entry *event.RelayEntrySubmitted),
) subscription.EventSubscription {
	mrb.relayEntrySubmissionHandlersMutex.Lock()
	defer mrb.relayEntrySubmissionHandlersMutex.Unlock()

	// #nosec G404 (insecure random number source (rand))
	// Temporary test implementation doesn't require secure randomness.
	handlerID := rand.Int()

	mrb.relayEntrySubmissionHandlers[handlerID] = handler

	return subscription.NewEventSubscription(func() {
		mrb.relayEntrySubmissionHandlersMutex.Lock()
		defer mrb.relayEntrySubmissionHandlersMutex.Unlock()

		delete(mrb.relayEntrySubmissionHandlers, handlerID)
	})
}

func (mrb *mockRandomBeacon) SubmitRelayEntry(
	entry []byte,
) error {
	mrb.relayEntrySubmissionHandlersMutex.Lock()
	defer mrb.relayEntrySubmissionHandlersMutex.Unlock()

	mrb.latestRelayEntryMutex.Lock()
	defer mrb.latestRelayEntryMutex.Unlock()

	mrb.currentRequestMutex.Lock()
	defer mrb.currentRequestMutex.Unlock()

	blockNumber, err := mrb.blockCounter.CurrentBlock()
	if err != nil {
		return fmt.Errorf("failed to get block counter: [%v]", err)
	}

	for _, handler := range mrb.relayEntrySubmissionHandlers {
		go func(handler func(submitted *event.RelayEntrySubmitted)) {
			handler(&event.RelayEntrySubmitted{BlockNumber: blockNumber})
		}(handler)
	}

	mrb.latestRelayEntry = entry
	mrb.currentRequestStartBlock = nil
	mrb.currentRequestPreviousEntry = nil
	mrb.currentRequestGroup = nil

	return nil
}

func (mrb *mockRandomBeacon) IsEntryInProgress() (bool, error) {
	mrb.currentRequestMutex.RLock()
	defer mrb.currentRequestMutex.RUnlock()

	return mrb.currentRequestStartBlock != nil, nil
}

func (mrb *mockRandomBeacon) CurrentRequestStartBlock() (*big.Int, error) {
	mrb.currentRequestMutex.RLock()
	defer mrb.currentRequestMutex.RUnlock()

	return mrb.currentRequestStartBlock, nil
}

func (mrb *mockRandomBeacon) CurrentRequestPreviousEntry() ([]byte, error) {
	mrb.currentRequestMutex.RLock()
	defer mrb.currentRequestMutex.RUnlock()

	return mrb.currentRequestPreviousEntry, nil
}

func (mrb *mockRandomBeacon) CurrentRequestGroupPublicKey() ([]byte, error) {
	mrb.currentRequestMutex.RLock()
	defer mrb.currentRequestMutex.RUnlock()

	return mrb.currentRequestGroup, nil
}
