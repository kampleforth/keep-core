package gjkr

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/keep-network/keep-core/pkg/beacon/relay/group"
	"github.com/keep-network/keep-core/pkg/net/ephemeral"
)

type memberCore struct {
	// ID of this group member.
	ID group.MemberIndex

	// Group to which this member belongs.
	group *group.Group

	// Evidence log provides access to messages from earlier protocol phases
	// for the sake of compliant resolution.
	evidenceLog evidenceLog

	// Cryptographic protocol parameters, the same for all members in the group.
	protocolParameters *protocolParameters
}

// LocalMember represents one member in a threshold group, prior to the
// initiation of distributed key generation process
type LocalMember struct {
	*memberCore
}

// EphemeralKeyPairGeneratingMember represents one member in a distributed key
// generating group performing ephemeral key pair generation. It has a full list
// of `memberIDs` that belong to its threshold group.
//
// Executes Phase 1 of the protocol.
type EphemeralKeyPairGeneratingMember struct {
	*LocalMember

	// Ephemeral key pairs used to create symmetric keys,
	// generated individually for each other group member.
	ephemeralKeyPairs map[group.MemberIndex]*ephemeral.KeyPair
}

// SymmetricKeyGeneratingMember represents one member in a distributed key
// generating group performing ephemeral symmetric key generation.
//
// Executes Phase 2 of the protocol.
type SymmetricKeyGeneratingMember struct {
	*EphemeralKeyPairGeneratingMember

	// Symmetric keys used to encrypt confidential information,
	// generated individually for each other group member by ECDH'ing the
	// broadcasted ephemeral public key intended for this member and the
	// ephemeral private key generated for the other member.
	symmetricKeys map[group.MemberIndex]ephemeral.SymmetricKey
}

// CommittingMember represents one member in a distributed key generation group,
// after it has fully initialized ephemeral symmetric keys with all other group
// members.
//
// Executes Phase 3 of the protocol.
type CommittingMember struct {
	*SymmetricKeyGeneratingMember

	// Polynomial `a` coefficients generated by the member. Polynomial is of
	// degree `dishonestThreshold`, so the number of coefficients equals
	// `dishonestThreshold + 1` (including constant coefficient).
	//
	// This is a private value and should not be exposed.
	secretCoefficients []*big.Int
	// Shares calculated by the current member for themself. They are defined as
	// `s_ii` and `t_ii` respectively across the protocol specification.
	//
	// These are private values and should not be exposed.
	selfSecretShareS, selfSecretShareT *big.Int
}

// CommitmentsVerifyingMember represents one member in a distributed key generation
// group, after it has received secret shares and commitments from other group
// members and it performs verification of received values.
//
// Executes Phase 4 of the protocol.
type CommitmentsVerifyingMember struct {
	*CommittingMember

	// Shares calculated for the current member by peer group members which passed
	// the validation.
	//
	// receivedQualifiedSharesS are defined as `s_ji` and receivedQualifiedSharesT are
	// defined as `t_ji` across the protocol specification.
	// TODO remove receivedQualifiedSharesT - exists only for unit tests purpose
	receivedQualifiedSharesS, receivedQualifiedSharesT map[group.MemberIndex]*big.Int
	// Commitments to secret shares polynomial coefficients received from
	// other group members.
	receivedPeerCommitments map[group.MemberIndex][]*bn256.G1
}

// SharesJustifyingMember represents one member in a threshold key sharing group,
// after it completed secret shares and commitments verification and enters
// justification phase where it resolves invalid share accusations.
//
// Executes Phase 5 of the protocol.
type SharesJustifyingMember struct {
	*CommitmentsVerifyingMember
}

// QualifiedMember represents one member in a threshold key sharing group, after
// it completed secret shares justification. The member holds a share of group
// group private key.
//
// Executes Phase 6 of the protocol.
type QualifiedMember struct {
	*SharesJustifyingMember

	// Member's share of the secret group private key. It is denoted as `z_ik`
	// in protocol specification.
	groupPrivateKeyShare *big.Int
}

// SharingMember represents one member in a threshold key sharing group, after it
// has been qualified to the group private key sharing. A member shares
// public values of it's polynomial coefficients with peer members.
//
// Executes Phase 7 and Phase 8 of the protocol.
type SharingMember struct {
	*QualifiedMember

	// Public values of each polynomial `a` coefficient defined in secretCoefficients
	// field. It is denoted as `A_ik` in protocol specification. The zeroth
	// public key share point `A_i0` is a member's public key share.
	publicKeySharePoints []*bn256.G2
	// Public key share points received from other group members which passed
	// the validation. Defined as `A_jk` across the protocol documentation.
	receivedValidPeerPublicKeySharePoints map[group.MemberIndex][]*bn256.G2
}

// PointsJustifyingMember represents one member in a threshold key sharing group,
// after it completed public key share points verification and enters justification
// phase where it resolves public key share points accusations.
//
// Executes Phase 9 of the protocol.
type PointsJustifyingMember struct {
	*SharingMember
}

// RevealingMember represents one member in a threshold sharing group who is
// revealing ephemeral private keys used to create ephemeral symmetric key
// to communicate with other members disqualified in Phase 9.
//
// Executes Phase 10 of the protocol.
type RevealingMember struct {
	*PointsJustifyingMember

	// Slice of disqualified or inactive QUAL members whose phase 3 shares need
	// to be reconstructed.
	// We snapshot this information at the beginning of phase 10 to make sure no
	// new disqualifications can affect the validation of the message with
	// ephemeral keys in phase 11. Specifically, that we do not disqualify
	// everyone when one of the group members do not deliver phase 11 message
	// at all. This can happen if we mark that member as inactive and then,
	// disqualify everyone else because all the messages do not contain a
	// revealed key for that inactive member. For this reason, it is safer
	// to use a snapshotted set of misbehaved members whose keys are expected
	// to be revealed.
	expectedMembersForReconstruction []group.MemberIndex
}

// ReconstructingMember represents one member in a threshold sharing group who
// is reconstructing individual private and public keys of disqualified group members.
//
// Executes Phase 11 of the protocol.
type ReconstructingMember struct {
	*RevealingMember

	// Revealed shares of members from the QUAL set disqualified or marked as
	// inactive in later phases, after QUAL set has been established.
	revealedMisbehavedMembersShares []*misbehavedShares
	// Disqualified members' individual private keys reconstructed from shares
	// revealed by other group members.
	// Stored as `<m, z_m>`, where:
	// - `m` is disqualified member's ID
	// - `z_m` is reconstructed individual private key of member `m`
	reconstructedIndividualPrivateKeys map[group.MemberIndex]*big.Int
	// Individual public keys calculated from reconstructed individual private keys.
	// Stored as `<m, y_m>`, where:
	// - `m` is disqualified member's ID
	// - `y_m` is reconstructed individual public key of member `m`
	reconstructedIndividualPublicKeys map[group.MemberIndex]*bn256.G2
}

// CombiningMember represents one member in a threshold sharing group who is
// combining individual public keys of group members to receive group public key.
//
// Executes Phase 12 of the protocol.
type CombiningMember struct {
	*ReconstructingMember

	// Group public key calculated from individual public keys of all group members.
	// Denoted as `Y` across the protocol specification.
	groupPublicKey *bn256.G2

	groupPublicKeySharesChannel chan map[group.MemberIndex]*bn256.G2
}

// InitializeFinalization returns a member to perform next protocol operations.
func (cm *CombiningMember) InitializeFinalization() *FinalizingMember {
	return &FinalizingMember{CombiningMember: cm}
}

// FinalizingMember represents one member in a threshold key sharing group,
// after it completed distributed key generation.
//
// Prepares a result to publish in Phase 13 of the protocol.
type FinalizingMember struct {
	*CombiningMember
}

// NewMember creates a new member in an initial state
func NewMember(
	memberID group.MemberIndex,
	groupSize,
	dishonestThreshold int,
	seed *big.Int,
) (*LocalMember, error) {
	return &LocalMember{
		memberCore: &memberCore{
			memberID,
			group.NewDkgGroup(dishonestThreshold, groupSize),
			newDkgEvidenceLog(),
			newProtocolParameters(seed),
		},
	}, nil
}

// InitializeEphemeralKeysGeneration performs a transition of a member state
// from the local state to phase 1 of the protocol.
func (lm *LocalMember) InitializeEphemeralKeysGeneration() *EphemeralKeyPairGeneratingMember {
	return &EphemeralKeyPairGeneratingMember{
		LocalMember:       lm,
		ephemeralKeyPairs: make(map[group.MemberIndex]*ephemeral.KeyPair),
	}
}

// InitializeSymmetricKeyGeneration performs a transition of the member state
// from phase 1 to phase 2. It returns a member instance ready to execute the
// next phase of the protocol.
func (ekgm *EphemeralKeyPairGeneratingMember) InitializeSymmetricKeyGeneration() *SymmetricKeyGeneratingMember {
	return &SymmetricKeyGeneratingMember{
		EphemeralKeyPairGeneratingMember: ekgm,
		symmetricKeys:                    make(map[group.MemberIndex]ephemeral.SymmetricKey),
	}
}

// InitializeCommitting returns a member to perform next protocol operations.
func (skgm *SymmetricKeyGeneratingMember) InitializeCommitting() *CommittingMember {
	return &CommittingMember{
		SymmetricKeyGeneratingMember: skgm,
	}
}

// InitializeCommitmentsVerification returns a member to perform next protocol operations.
func (cm *CommittingMember) InitializeCommitmentsVerification() *CommitmentsVerifyingMember {
	return &CommitmentsVerifyingMember{
		CommittingMember:         cm,
		receivedQualifiedSharesS: make(map[group.MemberIndex]*big.Int),
		receivedQualifiedSharesT: make(map[group.MemberIndex]*big.Int),
		receivedPeerCommitments:  make(map[group.MemberIndex][]*bn256.G1),
	}
}

// InitializeSharesJustification returns a member to perform next protocol operations.
func (cvm *CommitmentsVerifyingMember) InitializeSharesJustification() *SharesJustifyingMember {
	return &SharesJustifyingMember{cvm}
}

// InitializeQualified returns a member to perform next protocol operations.
func (sjm *SharesJustifyingMember) InitializeQualified() *QualifiedMember {
	return &QualifiedMember{SharesJustifyingMember: sjm}
}

// InitializeSharing returns a member to perform next protocol operations.
func (qm *QualifiedMember) InitializeSharing() *SharingMember {
	return &SharingMember{
		QualifiedMember:                       qm,
		receivedValidPeerPublicKeySharePoints: make(map[group.MemberIndex][]*bn256.G2),
	}
}

// InitializePointsJustification returns a member to perform next protocol operations.
func (sm *SharingMember) InitializePointsJustification() *PointsJustifyingMember {
	return &PointsJustifyingMember{sm}
}

// InitializeRevealing returns a member to perform next protocol operations.
func (pjm *PointsJustifyingMember) InitializeRevealing() *RevealingMember {
	return &RevealingMember{
		PointsJustifyingMember:           pjm,
		expectedMembersForReconstruction: make([]group.MemberIndex, 0),
	}
}

// InitializeReconstruction returns a member to perform next protocol operations.
func (rm *RevealingMember) InitializeReconstruction() *ReconstructingMember {
	return &ReconstructingMember{
		RevealingMember:                    rm,
		reconstructedIndividualPrivateKeys: make(map[group.MemberIndex]*big.Int),
		reconstructedIndividualPublicKeys:  make(map[group.MemberIndex]*bn256.G2),
	}
}

// InitializeCombining returns a member to perform next protocol operations.
func (rm *ReconstructingMember) InitializeCombining() *CombiningMember {
	return &CombiningMember{
		ReconstructingMember:        rm,
		groupPublicKeySharesChannel: make(chan map[group.MemberIndex]*bn256.G2),
	}
}

// individualPrivateKey returns current member's individual private key.
// Individual private key is zeroth polynomial coefficient `a_i0`.
func (rm *ReconstructingMember) individualPrivateKey() *big.Int {
	return rm.secretCoefficients[0]
}

// individualPublicKey returns current member's individual public key.
// Individual public key is zeroth public key share point `A_i0`.
func (rm *ReconstructingMember) individualPublicKey() *bn256.G2 {
	return rm.publicKeySharePoints[0]
}

// receivedValidPeerIndividualPublicKeys returns individual public keys received
// from other members which passed the validation. Individual public key is zeroth
// public key share point `A_j0`.
func (sm *SharingMember) receivedValidPeerIndividualPublicKeys() []*bn256.G2 {
	var receivedValidPeerIndividualPublicKeys []*bn256.G2

	for _, peerPublicKeySharePoints := range sm.receivedValidPeerPublicKeySharePoints {
		receivedValidPeerIndividualPublicKeys = append(
			receivedValidPeerIndividualPublicKeys,
			peerPublicKeySharePoints[0],
		)
	}
	return receivedValidPeerIndividualPublicKeys
}

// Result can be either the successful computation of a round of distributed key
// generation, or a notification of failure.
// It returns the generated group public key and a private key share of a group
// key along with the disqualified and inactive members (as part of including the
// group state). The group private key share is used for signing and should never
// be revealed publicly.
func (fm *FinalizingMember) Result() *Result {
	return &Result{
		Group:                       fm.group,
		GroupPublicKey:              fm.groupPublicKey, // nil if threshold not satisfied
		GroupPrivateKeyShare:        fm.groupPrivateKeyShare,
		groupPublicKeySharesChannel: fm.groupPublicKeySharesChannel,
	}
}
