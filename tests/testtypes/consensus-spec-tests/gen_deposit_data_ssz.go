// Code generated by github.com/karalabe/ssz. DO NOT EDIT.

package consensus_spec_tests

import "github.com/karalabe/ssz"

// SizeSSZ returns the total size of the static ssz object.
func (obj *DepositData) SizeSSZ() uint32 {
	return 48 + 32 + 8 + 96
}

// DefineSSZ defines how an object is encoded/decoded.
func (obj *DepositData) DefineSSZ(codec *ssz.Codec) {
	ssz.DefineStaticBytes(codec, obj.Pubkey[:])                // Field  (0) -                Pubkey - 48 bytes
	ssz.DefineStaticBytes(codec, obj.WithdrawalCredentials[:]) // Field  (1) - WithdrawalCredentials - 32 bytes
	ssz.DefineUint64(codec, &obj.Amount)                       // Field  (2) -                Amount -  8 bytes
	ssz.DefineStaticBytes(codec, obj.Signature[:])             // Field  (3) -             Signature - 96 bytes
}
