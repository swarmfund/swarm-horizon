package resource

import "gitlab.com/swarmfund/go/xdr"

type LedgerEntry struct {
	LastModifiedLedgerSeq uint32
	Type                  xdr.LedgerEntryType `json:"type"`
	Account               *AccountEntry
	Asset                 *AssetEntry
	Balance               *BalanceEntry
}

func (r *LedgerEntry) Populate(xdrEntry *xdr.LedgerEntry) {
	r.Type = xdrEntry.Data.Type
	r.LastModifiedLedgerSeq = uint32(xdrEntry.LastModifiedLedgerSeq)

	switch r.Type {
	case xdr.LedgerEntryTypeAccount:
		r.Account = new(AccountEntry)
		r.Account.Populate(*xdrEntry.Data.Account)
	case xdr.LedgerEntryTypeAsset:
		r.Asset = new(AssetEntry)
		r.Asset.Populate(*xdrEntry.Data.Asset)
	case xdr.LedgerEntryTypeBalance:
		r.Balance = new(BalanceEntry)
		r.Balance.Populate(*xdrEntry.Data.Balance)
	}

}

type LedgerKey struct {
	Type    xdr.LedgerEntryType `json:"type"`
	Account *LedgerKeyAccount   `json:"account"`
	Asset   *LedgerKeyAsset     `json:"asset"`
	Balance *LedgerKeyBalance   `json:"balance"`
}

func (r *LedgerKey) Populate(ledgerKey *xdr.LedgerKey) {
	switch r.Type {
	case xdr.LedgerEntryTypeAccount:
		r.Account = new(LedgerKeyAccount)
		r.Account.Populate(*ledgerKey.Account)
	case xdr.LedgerEntryTypeAsset:
		r.Asset = new(LedgerKeyAsset)
		r.Asset.Populate(*ledgerKey.Asset)
	case xdr.LedgerEntryTypeBalance:
		r.Balance = new(LedgerKeyBalance)
		r.Balance.Populate(*ledgerKey.Balance)
	}
}
