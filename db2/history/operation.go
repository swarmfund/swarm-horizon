package history

import (
	"time"

	"gitlab.com/swarmfund/go/xdr"
	"gitlab.com/swarmfund/horizon/db2"
)

// Operation is a row of data from the `history_operations` table
type Operation struct {
	db2.TotalOrderID
	TransactionID    int64             `db:"transaction_id"`
	TransactionHash  string            `db:"transaction_hash"`
	ApplicationOrder int32             `db:"application_order"`
	Type             xdr.OperationType `db:"type"`
	Details          OperationDetails  `db:"details"`
	LedgerCloseTime  time.Time         `db:"ledger_close_time"`
	SourceAccount    string            `db:"source_account"`
	State            OperationState    `db:"state"`
	Identifier       int64             `db:"identifier"`
}
