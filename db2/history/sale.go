package history

import (
	"gitlab.com/swarmfund/horizon/db2"
	"time"
)

type Sale struct {
	ID                uint64      `db:"id"`
	OwnerID           string      `db:"owner_id"`
	BaseAsset         string      `db:"base_asset"`
	DefaultQuoteAsset string      `db:"default_quote_asset"`
	StartTime         time.Time   `db:"start_time"`
	EndTime           time.Time   `db:"end_time"`
	SoftCap           uint64      `db:"soft_cap"`
	HardCap           uint64      `db:"hard_cap"`
	CurrentCap        uint64      `db:"current_cap"`
	Details           db2.Details `db:"details"`
	State             SaleState   `db:"state"`
	QuoteAssets       db2.Details `db:"quote_assets"`
}
