package resource

import (
	"time"

	"fmt"

	"gitlab.com/swarmfund/go/xdr"
	"gitlab.com/swarmfund/horizon/db2/history"
)

type LedgerChanges struct {
	ID              string                 `json:"id"`
	PT              string                 `json:"paging_token"`
	Ledger          int32                  `json:"ledger"`
	LedgerCloseTime time.Time              `json:"created_at"`
	Changes         xdr.LedgerEntryChanges `json:"changes"`
}

func (lc *LedgerChanges) Populate(tm history.TransactionMeta) error {
	lc.ID = fmt.Sprintf("%d", tm.ID)
	lc.Ledger = tm.LedgerSequence
	lc.PT = tm.PagingToken()
	lc.LedgerCloseTime = tm.LedgerCloseTime

	txMeta := xdr.TransactionMeta{}
	err := xdr.SafeUnmarshalBase64(tm.TxMeta, &txMeta)
	if err != nil {
		return err
	}
	if txMeta.Operations == nil {
		return nil
	}

	lc.Changes = make([]xdr.LedgerEntryChange, 0)
	for _, opMeta := range *txMeta.Operations {
		lc.Changes = append(lc.Changes, opMeta.Changes...)
	}

	return nil
}

func (lc LedgerChanges) PagingToken() string {
	return lc.PT
}
