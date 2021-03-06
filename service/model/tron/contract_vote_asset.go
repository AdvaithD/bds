package tron

type VoteAssetContract struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	TransactionHash   string `xorm:"transaction_hash char(64) notnull index"`
	BlockNumber       int64  `xorm:"block_number int index"`
	Timestamp         int64  `xorm:"timestamp int notnull index"`
	OwnerAddress      string `xorm:"owner_address char(64) notnull"`
	VoteAddress       string `xorm:"vote_address char(64) notnull"`
	Support           bool   `xorm:"support bool"`
	Count             int64  `xorm:"count bigint"`
	VoteAddressNumber int    `xorm:"vote_address_number int"`
}

func (c VoteAssetContract) TableName() string {
	return tableName("contract_vote_asset")
}
