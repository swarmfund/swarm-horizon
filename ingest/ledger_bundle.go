package ingest

import (
	"gitlab.com/swarmfund/horizon/db2/core"
	"gitlab.com/swarmfund/horizon/db2/history"
)

// Load runs queries against `core` to fill in the records of the bundle.
func (lb *LedgerBundle) Load(coreQ core.QInterface, historyQ history.QInterface) error {

	// Load Header
	err := coreQ.LedgerHeaderBySequence(&lb.Header, lb.Sequence)
	if err != nil {
		return err
	}

	// Load transactions
	err = coreQ.TransactionsByLedger(&lb.Transactions, lb.Sequence)
	if err != nil {
		return err
	}

	err = coreQ.TransactionFeesByLedger(&lb.TransactionFees, lb.Sequence)
	if err != nil {
		return err
	}

	return nil
}
