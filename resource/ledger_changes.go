package resource

import (
	"time"

	"fmt"

	"gitlab.com/swarmfund/go/xdr"
	"gitlab.com/swarmfund/horizon/db2/history"
)

type LedgerChanges struct {
	ID              string              `json:"id"`
	PT              string              `json:"paging_token"`
	Ledger          int32               `json:"ledger"`
	LedgerCloseTime time.Time           `json:"created_at"`
	Changes         []LedgerEntryChange `json:"changes"`
}

func (lc LedgerChanges) PagingToken() string {
	return lc.PT
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

	lc.Changes = make([]LedgerEntryChange, 0)

	for _, opMeta := range txMeta.MustOperations() {
		for _, xdrChange := range opMeta.Changes {
			res := LedgerEntryChange{}
			res.Populate(xdrChange)
			lc.Changes = append(lc.Changes, res)
		}
	}

	return nil
}

type LedgerEntryChange struct {
	Type    xdr.LedgerEntryChangeType `json:"type"`
	Created *LedgerEntry              `json:"created"`
	Updated *LedgerEntry              `json:"updated"`
	Removed *LedgerKey                `json:"removed"`
	State   *LedgerEntry              `json:"state"`
}

func (r *LedgerEntryChange) Populate(xdrChange xdr.LedgerEntryChange) {
	r.Type = xdrChange.Type

	switch xdrChange.Type {
	case xdr.LedgerEntryChangeTypeCreated:
		r.Created = &LedgerEntry{}
		r.Created.Populate(xdrChange.Created)
	case xdr.LedgerEntryChangeTypeUpdated:
		r.Updated = &LedgerEntry{}
		r.Updated.Populate(xdrChange.Updated)
	case xdr.LedgerEntryChangeTypeRemoved:
		r.Removed = &LedgerKey{}
		r.Removed.Populate(xdrChange.Removed)
	case xdr.LedgerEntryChangeTypeState:
		r.State = &LedgerEntry{}
		r.State.Populate(xdrChange.State)
	}
}
