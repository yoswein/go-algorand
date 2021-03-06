// Code generated by "stringer -type=TxnField"; DO NOT EDIT.

package logic

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sender-0]
	_ = x[Fee-1]
	_ = x[FirstValid-2]
	_ = x[FirstValidTime-3]
	_ = x[LastValid-4]
	_ = x[Note-5]
	_ = x[Lease-6]
	_ = x[Receiver-7]
	_ = x[Amount-8]
	_ = x[CloseRemainderTo-9]
	_ = x[VotePK-10]
	_ = x[SelectionPK-11]
	_ = x[VoteFirst-12]
	_ = x[VoteLast-13]
	_ = x[VoteKeyDilution-14]
	_ = x[Type-15]
	_ = x[TypeEnum-16]
	_ = x[XferAsset-17]
	_ = x[AssetAmount-18]
	_ = x[AssetSender-19]
	_ = x[AssetReceiver-20]
	_ = x[AssetCloseTo-21]
	_ = x[GroupIndex-22]
	_ = x[TxID-23]
	_ = x[invalidTxnField-24]
}

const _TxnField_name = "SenderFeeFirstValidFirstValidTimeLastValidNoteLeaseReceiverAmountCloseRemainderToVotePKSelectionPKVoteFirstVoteLastVoteKeyDilutionTypeTypeEnumXferAssetAssetAmountAssetSenderAssetReceiverAssetCloseToGroupIndexTxIDinvalidTxnField"

var _TxnField_index = [...]uint8{0, 6, 9, 19, 33, 42, 46, 51, 59, 65, 81, 87, 98, 107, 115, 130, 134, 142, 151, 162, 173, 186, 198, 208, 212, 227}

func (i TxnField) String() string {
	if i < 0 || i >= TxnField(len(_TxnField_index)-1) {
		return "TxnField(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TxnField_name[_TxnField_index[i]:_TxnField_index[i+1]]
}
