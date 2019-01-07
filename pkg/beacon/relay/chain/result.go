package chain

import (
	"math/big"
)

// DKGResult is a result of distributed key generation protocol.
//
// Success means that the protocol execution finished with acceptable number of
// disqualified or inactive members. The group of remaining members should be
// added to the signing groups for the threshold relay.
//
// Failure means that the group creation could not finish, due to either the number
// of inactive or disqualified participants, or the presented results being
// disputed in a way where the correct outcome cannot be ascertained.
type DKGResult struct {
	// Result type of the protocol execution. True if success, false if failure.
	Success bool
	// Group public key generated by protocol execution, nil if the protocol failed.
	GroupPublicKey *big.Int
	// Disqualified members. Length of the slice and order of members are the same
	// as in the members group. Disqualified members are marked as true. It is
	// kept in this form as an optimization for an on-chain storage.
	Disqualified []bool
	// Inactive members. Length of the slice and order of members are the same
	// as in the members group. Disqualified members are marked as true. It is
	// kept in this form as an optimization for an on-chain storage.
	Inactive []bool
}

// Equals checks if two DKG results are equal.
func (r1 *DKGResult) Equals(r2 *DKGResult) bool {
	if r1 == nil || r2 == nil {
		return r1 == r2
	}
	if r1.Success != r2.Success {
		return false
	}
	if !bigIntEquals(r1.GroupPublicKey, r2.GroupPublicKey) {
		return false
	}
	if !boolSlicesEqual(r1.Disqualified, r2.Disqualified) {
		return false
	}
	if !boolSlicesEqual(r1.Inactive, r2.Inactive) {
		return false
	}
	return true
}

// bigIntEquals checks if two big.Int values are equal.
func bigIntEquals(expected *big.Int, actual *big.Int) bool {
	if expected != nil && actual != nil {
		return expected.Cmp(actual) == 0
	}
	return expected == nil && actual == nil
}

// boolSlicesEqual checks if two slices of bool are equal. Slices need to have
// the same length and have the same order of entries.
func boolSlicesEqual(expectedSlice []bool, actualSlice []bool) bool {
	if len(expectedSlice) != len(actualSlice) {
		return false
	}
	for i := range expectedSlice {
		if expectedSlice[i] != actualSlice[i] {
			return false
		}
	}
	return true
}
