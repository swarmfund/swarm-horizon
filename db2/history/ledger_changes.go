package history

import (
	"time"

	"gitlab.com/swarmfund/horizon/db2"
)

// Transaction is a row of data from the `history_transactions` table.
type TransactionMeta struct {
	db2.TotalOrderID
	LedgerSequence  int32     `db:"ledger_sequence"`
	LedgerCloseTime time.Time `db:"ledger_close_time"`
	TxMeta          string    `db:"tx_meta"`
}
