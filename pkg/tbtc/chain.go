package tbtc

import (
	"math/big"

	"github.com/keep-network/keep-core/pkg/chain"
	"github.com/keep-network/keep-core/pkg/operator"
	"github.com/keep-network/keep-core/pkg/protocol/group"
	"github.com/keep-network/keep-core/pkg/sortition"
	"github.com/keep-network/keep-core/pkg/subscription"
	"github.com/keep-network/keep-core/pkg/tecdsa/dkg"
)

type DKGState int

const (
	Idle DKGState = iota
	AwaitingSeed
	AwaitingResult
	Challenge
)

// GroupSelectionChain defines the subset of the TBTC chain interface that
// pertains to the group selection activities.
type GroupSelectionChain interface {
	// SelectGroup returns the group members selected for the current group
	// selection. The function returns an error if the chain's state does not
	// allow for group selection at the moment.
	SelectGroup() (*GroupSelectionResult, error)
}

// GroupSelectionResult represents a group selection result, i.e. operators
// selected to perform the DKG protocol. The result consists of two slices
// of equal length holding the chain.OperatorID and chain.Address for each
// selected operator.
type GroupSelectionResult struct {
	OperatorsIDs       chain.OperatorIDs
	OperatorsAddresses chain.Addresses
}

// DistributedKeyGenerationChain defines the subset of the TBTC chain
// interface that pertains specifically to group formation's distributed key
// generation process.
type DistributedKeyGenerationChain interface {
	// OnDKGStarted registers a callback that is invoked when an on-chain
	// notification of the DKG process start is seen.
	OnDKGStarted(
		func(event *DKGStartedEvent),
	) subscription.EventSubscription

	// OnDKGResultSubmitted registers a callback that is invoked when an on-chain
	// notification of the DKG result submission is seen.
	OnDKGResultSubmitted(
		func(event *DKGResultSubmittedEvent),
	) subscription.EventSubscription

	// SubmitDKGResult submits the DKG result to the chain, along with signatures
	// over result hash from group participants supporting the result.
	SubmitDKGResult(
		memberIndex group.MemberIndex,
		dkgResult *dkg.Result,
		signatures map[group.MemberIndex][]byte,
		groupSelectionResult *GroupSelectionResult,
	) error

	// GetDKGState returns the current state of the DKG procedure.
	GetDKGState() (DKGState, error)

	// CalculateDKGResultHash calculates 256-bit hash of DKG result in standard
	// specific for the chain. Operation is performed off-chain.
	CalculateDKGResultHash(result *dkg.Result) (dkg.ResultHash, error)
}

// DKGStartedEvent represents a DKG start event.
type DKGStartedEvent struct {
	Seed        *big.Int
	BlockNumber uint64
}

// DKGResultSubmittedEvent represents a DKG result submission event. It is emitted
// after a submitted DKG result is positively validated on the chain. It contains
// the index of the member who submitted the result and a final public key of
// the group.
type DKGResultSubmittedEvent struct {
	MemberIndex         uint32
	GroupPublicKeyBytes []byte
	Misbehaved          []uint8

	BlockNumber uint64
}

// BridgeChain defines the subset of the TBTC chain interface that pertains
// specifically to the tBTC Bridge operations.
type BridgeChain interface {
	// OnHeartbeatRequested registers a callback that is invoked when an
	// on-chain notification of the Bridge heartbeat request is seen.
	OnHeartbeatRequested(
		func(event *HeartbeatRequestedEvent),
	) subscription.EventSubscription
}

// HeartbeatRequestedEvent represents a Bridge heartbeat request event.
type HeartbeatRequestedEvent struct {
	WalletPublicKey []byte
	Messages        []*big.Int
	BlockNumber     uint64
}

// Chain represents the interface that the TBTC module expects to interact
// with the anchoring blockchain on.
type Chain interface {
	// GetConfig returns the expected configuration of the TBTC module.
	GetConfig() *ChainConfig
	// BlockCounter returns the chain's block counter.
	BlockCounter() (chain.BlockCounter, error)
	// Signing returns the chain's signer.
	Signing() chain.Signing
	// OperatorKeyPair returns the key pair of the operator assigned to this
	// chain handle.
	OperatorKeyPair() (*operator.PrivateKey, *operator.PublicKey, error)

	sortition.Chain
	GroupSelectionChain
	DistributedKeyGenerationChain
	BridgeChain
}

// ChainConfig contains the config data needed for the TBTC to operate.
type ChainConfig struct {
	// GroupSize is the target size of a group in TBTC.
	GroupSize int
	// GroupQuorum is the minimum number of active participants behaving
	// according to the protocol needed to generate a group in TBTC. This value
	// is smaller than the GroupSize and bigger than the HonestThreshold.
	GroupQuorum int
	// HonestThreshold is the minimum number of active participants behaving
	// according to the protocol needed to generate a signature.
	HonestThreshold int
}

// DishonestThreshold is the maximum number of misbehaving participants for
// which it is still possible to generate a signature.
// Misbehaviour is any misconduct to the protocol, including inactivity.
func (cc *ChainConfig) DishonestThreshold() int {
	return cc.GroupSize - cc.HonestThreshold
}
