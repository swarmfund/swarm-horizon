package history

import (
	sq "github.com/lann/squirrel"
	"gitlab.com/swarmfund/horizon/db2"
)

var selectTransactionMeta = sq.Select(
	"ht.id, " +
		"ht.ledger_sequence, " +
		"ht.tx_meta, " +
		"ht.ledger_close_time").
	From("history_transactions ht")

// LedgerChangesQ is a helper struct to aid in configuring queries that loads
// slices of transaction_meta structs.
type LedgerChangesQ struct {
	Err    error
	parent *Q
	sql    sq.SelectBuilder
}

type LedgerChangesQI interface {
	Page(page db2.PageQuery) LedgerChangesQI
	Select(dest interface{}) error
}

// LedgerChanges provides a helper to filter rows from the `history_transactions`
// table with pre-defined filters.  See `LedgerChangesQ` methods for the
// available filters.
func (q *Q) LedgerChanges() LedgerChangesQI {
	return &LedgerChangesQ{
		parent: q,
		sql:    selectTransactionMeta,
	}
}

// Page specifies the paging constraints for the query being built by `q`.
func (q *LedgerChangesQ) Page(page db2.PageQuery) LedgerChangesQI {
	if q.Err != nil {
		return q
	}

	q.sql, q.Err = page.ApplyTo(q.sql, "ht.id")
	return q
}

// Select loads the results of the query specified by `q` into `dest`.
func (q *LedgerChangesQ) Select(dest interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	q.Err = q.parent.Select(dest, q.sql)
	return q.Err
}
