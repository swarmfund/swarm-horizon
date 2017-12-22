package resource

import (
	"gitlab.com/swarmfund/go/amount"
	"gitlab.com/swarmfund/go/xdr"
)

type AssetEntry struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	Description string   `json:"description"`
	LogoId      string   `json:"logo_id"`
	Policies    Policies `json:"policies"`

	ExternalResourceLink  string `json:"external_resource_link"`
	PreissuedAssetSigner  string `json:"preissued_asset_signer"`
	AvailableForIssueance string `json:"available_for_issueance"`
	Issued                string `json:"issued"`
	MaxIssuanceAmount     string `json:"max_issuance_amount"`
}

func (r *AssetEntry) Populate(entry xdr.AssetEntry) {
	r.Code = string(entry.Code)
	r.Name = string(entry.Name)
	r.Owner = entry.Owner.Address()
	r.Description = string(entry.Description)
	r.LogoId = string(entry.LogoId)

	r.Policies.PopulateFromInt32(int32(entry.Policies))
	r.PreissuedAssetSigner = entry.PreissuedAssetSigner.Address()
	r.ExternalResourceLink = string(entry.ExternalResourceLink)

	r.AvailableForIssueance = amount.String(int64(entry.AvailableForIssueance))
	r.Issued = amount.String(int64(entry.Issued))
	r.MaxIssuanceAmount = amount.String(int64(entry.MaxIssuanceAmount))
}

type LedgerKeyAsset struct {
	AssetCode string `json:"asset_code"`
}

func (r *LedgerKeyAsset) Populate(entry xdr.LedgerKeyAsset) {
	r.AssetCode = string(entry.Code)
}
