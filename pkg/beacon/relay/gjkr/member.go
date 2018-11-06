package gjkr

import (
	"math/big"

	"github.com/keep-network/keep-core/pkg/beacon/relay/pedersen"
)

type memberCore struct {
	// ID of this group member.
	ID int
	// Group to which this member belongs.
	group *Group
	// DKG Protocol configuration parameters.
	protocolConfig *DKG
}

// CommittingMember represents one member in a threshold key sharing group, after
// it has a full list of `memberIDs` that belong to its threshold group. A
// member in this state has two maps of member shares for each member of the
// group.
type CommittingMember struct {
	*memberCore

	// Pedersen VSS scheme used to calculate commitments.
	vss *pedersen.VSS
	// Polynomial `a` coefficients generated by the member. Polynomial is of
	// degree `dishonestThreshold`, so the number of coefficients equals
	// `dishonestThreshold + 1`
	//
	// This is a private value and should not be exposed.
	secretCoefficients []*big.Int
	// Shares calculated by the current member for themself. They are defined as
	// `s_ii` and `t_ii` respectively across the protocol specification.
	//
	// These are private values and should not be exposed.
	selfSecretShareS, selfSecretShareT *big.Int
	// Shares calculated for the current member by peer group members.
	//
	// receivedSharesS are defined as `s_ji` and receivedSharesT are
	// defined as `t_ji` across the protocol specification.
	receivedSharesS, receivedSharesT map[int]*big.Int
	// Commitments to coefficients received from peer group members.
	receivedCommitments map[int][]*big.Int
}

// SharesJustifyingMember represents one member in a threshold key sharing group,
// after it completed secret shares and commitments verification and enters
// justification phase where it resolves invalid share accusations.
type SharesJustifyingMember struct {
	*CommittingMember
}

// QualifiedMember represents one member in a threshold key sharing group, after
// it completed secret shares justification. The member holds a share of group
// master private key.
type QualifiedMember struct {
	*SharesJustifyingMember

	// Member's share of the secret master private key. It is denoted as `z_ik`
	// in protocol specification.
	// TODO: unsure if we need shareT `x'_i` field, it should be removed if not used in further steps
	masterPrivateKeyShare, shareT *big.Int
}

// SharingMember represents one member in a threshold key sharing group, after it
// has been qualified to the master private key sharing group. A member shares
// public values of it's polynomial coefficients with peer members.
type SharingMember struct {
	*QualifiedMember

	// Public values of each polynomial `a` coefficient defined in secretCoefficients
	// field. It is denoted as `A_ik` in protocol specification.
	publicCoefficients []*big.Int
}

// ReconstructingMember represents one member in a threshold sharing group who
// is reconstructing individual private and public keys of disqualified group members.
type ReconstructingMember struct {
	*SharingMember

	// Disqualified members' individual private keys reconstructed from shares
	// revealed by other group members.
	// Stored as `<m, z_m>`, where:
	// - `m` is disqualified member's ID
	// - `z_m` is reconstructed individual private key of member `m`
	reconstructedIndividualPrivateKeys map[int]*big.Int
	// Individual public keys calculated from reconstructed individual private keys.
	// Stored as `<m, y_m>`, where:
	// - `m` is disqualified member's ID
	// - `y_m` is reconstructed individual public key of member `m`
	reconstructedIndividualPublicKeys map[int]*big.Int
}
