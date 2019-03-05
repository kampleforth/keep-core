package dkg2

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/keep-network/keep-core/pkg/altbn128"
	"github.com/keep-network/keep-core/pkg/beacon/relay/gjkr"
	"github.com/keep-network/keep-core/pkg/bls"
)

// ThresholdSigner is created from GJKR group Member when DKG protocol completed
// successfully and each group member is ready to sign. ThresholdSigner contains
// its own private key share of group public key that should never be publicly
// revealed. It also contains group's public key and ID of GJKR Member
// represented by this ThresholdSigner instance.
type ThresholdSigner struct {
	memberID             gjkr.MemberID
	groupPublicKey       *bn256.G2
	groupPrivateKeyShare *big.Int
}

// MemberID returns GJKR MemberID represented by this ThresholdSigner.
func (ts *ThresholdSigner) MemberID() gjkr.MemberID {
	return ts.memberID
}

// GroupPublicKeyBytes returns group public key representation in bytes.
func (ts *ThresholdSigner) GroupPublicKeyBytes() []byte {
	altbn128GroupPublicKey := altbn128.G2Point{ts.groupPublicKey}
	return altbn128GroupPublicKey.Compress()
}

// CalculateSignatureShare takes the message and calculates signer's signature
// share over that message.
func (ts *ThresholdSigner) CalculateSignatureShare(message []byte) *bn256.G1 {
	return bls.Sign(ts.groupPrivateKeyShare, message)
}

// CompleteSignature accepts signature shares from all group threshold signers
// and produces a final group signature from them. Input slice should contain
// signature of the current signer as well. We parameterize the threshold (number
// of honest members we require), as we recover a threshold signature, not an
// aggregate signature (one which would require all members).
func (ts *ThresholdSigner) CompleteSignature(
	signatureShares []*bls.SignatureShare,
	threshold int,
) (*bn256.G1, error) {
	return bls.RecoverSignature(signatureShares, threshold)
}
