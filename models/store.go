package model

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error

	AddRecTxs(txs []*RechargeEvents) error
	UpdateTxStatus(hash string, data *RechargeEvents) error
	UpdateFreedTxStatus(hash string, data *FreedRecord) error
	FreedRecordInsert(data *FreedRecord) error
	GetFullRecord(t string) (txs []*RechargeEvents, err error)
	BidInsert(data *BidAddress) error
	QueryBid(bid *BidAddress) (data []*BidAddress, err error)
}
