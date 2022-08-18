package tbtc

import (
	"fmt"
	"math/big"

	"github.com/keep-network/keep-common/pkg/persistence"
	"github.com/keep-network/keep-core/pkg/generator"
	"github.com/keep-network/keep-core/pkg/internal/testutils"
	"github.com/keep-network/keep-core/pkg/net"
	"github.com/keep-network/keep-core/pkg/protocol/group"
	"github.com/keep-network/keep-core/pkg/tecdsa/dkg"
)

// TODO: Unit tests for `node.go`.

// node represents the current state of an ECDSA node.
type node struct {
	chain          Chain
	netProvider    net.Provider
	walletRegistry *walletRegistry
	dkgExecutor    *dkg.Executor
	protocolLatch  *generator.ProtocolLatch
}

func newNode(
	chain Chain,
	netProvider net.Provider,
	persistence persistence.Handle,
	scheduler *generator.Scheduler,
	config Config,
) *node {
	walletRegistry := newWalletRegistry(persistence)

	dkgExecutor := dkg.NewExecutor(
		logger,
		scheduler,
		config.PreParamsPoolSize,
		config.PreParamsGenerationTimeout,
		config.PreParamsGenerationDelay,
		config.PreParamsGenerationConcurrency,
	)

	latch := generator.NewProtocolLatch()
	scheduler.RegisterProtocol(latch)

	return &node{
		chain:          chain,
		netProvider:    netProvider,
		walletRegistry: walletRegistry,
		dkgExecutor:    dkgExecutor,
		protocolLatch:  latch,
	}
}

// joinDKGIfEligible takes a seed value and undergoes the process of the
// distributed key generation if this node's operator proves to be eligible for
// the group generated by that seed. This is an interactive on-chain process,
// and joinDKGIfEligible can block for an extended period of time while it
// completes the on-chain operation.
func (n *node) joinDKGIfEligible(seed *big.Int, startBlockNumber uint64) {
	logger.Infof(
		"checking eligibility for DKG with seed [0x%x]",
		seed,
	)

	selectedSigningGroupOperators, err := n.chain.SelectGroup(seed)
	if err != nil {
		logger.Errorf(
			"failed to select group with seed [0x%x]: [%v]",
			seed,
			err,
		)
		return
	}

	chainConfig := n.chain.GetConfig()

	if len(selectedSigningGroupOperators) > chainConfig.GroupSize {
		logger.Errorf(
			"group size larger than supported: [%v]",
			len(selectedSigningGroupOperators),
		)
		return
	}

	signing := n.chain.Signing()

	_, operatorPublicKey, err := n.chain.OperatorKeyPair()
	if err != nil {
		logger.Errorf("failed to get operator public key: [%v]", err)
		return
	}

	operatorAddress, err := signing.PublicKeyToAddress(operatorPublicKey)
	if err != nil {
		logger.Errorf("failed to get operator address: [%v]", err)
		return
	}

	indexes := make([]uint8, 0)
	for index, operator := range selectedSigningGroupOperators {
		// See if we are amongst those chosen
		if operator == operatorAddress {
			indexes = append(indexes, uint8(index))
		}
	}

	// Create temporary broadcast channel name for DKG using the
	// group selection seed with the protocol name as prefix.
	channelName := fmt.Sprintf("%s-%s", ProtocolName, seed.Text(16))

	if len(indexes) > 0 {
		logger.Infof(
			"joining DKG with seed [0x%x] and controlling [%v] group members",
			seed,
			len(indexes),
		)

		broadcastChannel, err := n.netProvider.BroadcastChannelFor(channelName)
		if err != nil {
			logger.Errorf("failed to get broadcast channel: [%v]", err)
			return
		}

		membershipValidator := group.NewMembershipValidator(
			&testutils.MockLogger{},
			selectedSigningGroupOperators,
			signing,
		)

		err = broadcastChannel.SetFilter(membershipValidator.IsInGroup)
		if err != nil {
			logger.Errorf(
				"could not set filter for channel [%v]: [%v]",
				broadcastChannel.Name(),
				err,
			)
		}

		blockCounter, err := n.chain.BlockCounter()
		if err != nil {
			logger.Errorf("failed to get block counter: [%v]", err)
			return
		}

		for _, index := range indexes {
			// Capture the member index for the goroutine. The group member
			// index should be in range [1, groupSize] so we need to add 1.
			memberIndex := index + 1

			go func() {
				n.protocolLatch.Lock()
				defer n.protocolLatch.Unlock()

				result, _, err := n.dkgExecutor.Execute(
					seed,
					startBlockNumber,
					memberIndex,
					chainConfig.GroupSize,
					chainConfig.DishonestThreshold(),
					blockCounter,
					broadcastChannel,
					membershipValidator,
				)
				if err != nil {
					// TODO: Add retries into the mix.
					logger.Errorf(
						"[member:%v] failed to execute dkg: [%v]",
						memberIndex,
						err,
					)
					return
				}

				// TODO: Snapshot the key material before doing on-chain result
				//       submission.

				// TODO: Submit the result using the chain layer.

				// TODO: The final `signingGroupOperators` may differ from
				//       the original `selectedSigningGroupOperators`.
				//       Consider that when integrating the retry algorithm.
				signer := newSigner(
					result.PrivateKeyShare.PublicKey(),
					selectedSigningGroupOperators,
					memberIndex,
					result.PrivateKeyShare,
				)

				err = n.walletRegistry.registerSigner(signer)
				if err != nil {
					logger.Errorf(
						"failed to register %s: [%v]",
						signer,
						err,
					)
					return
				}

				logger.Infof("registered %s", signer)
			}()
		}
	} else {
		logger.Infof("not eligible for DKG with seed [0x%x]", seed)
	}
}
