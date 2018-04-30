package resource

import (
	"gitlab.com/swarmfund/horizon/db2/core"
	"gitlab.com/tokend/go/amount"
)

type Asset struct {
	Code                 string `json:"code"`
	Owner                string `json:"owner"`
	AvailableForIssuance string `json:"available_for_issuance"`
	PreissuedAssetSigner string `json:"preissued_asset_signer"`
	MaxIssuanceAmount    string `json:"max_issuance_amount"`
	Issued               string `json:"issued"`
	PendingIssuance      string `json:"pending_issuance"`
	Policies
	Details map[string]interface{} `json:"details"`
	Sales   []Sale                 `json:"sales,omitempty"`
}

func (a *Asset) Populate(asset *core.Asset) {
	a.Code = asset.Code
	a.Owner = asset.Owner
	a.PreissuedAssetSigner = asset.PreissuedAssetSigner

	a.AvailableForIssuance = amount.StringU(asset.AvailableForIssuance)
	a.MaxIssuanceAmount = amount.StringU(asset.MaxIssuanceAmount)
	a.PendingIssuance = amount.StringU(asset.PendingIssuance)
	a.Issued = amount.StringU(asset.Issued)

	a.Policies.Populate(*asset)
	a.Details, _ = asset.GetDetails()
}

type AssetPair struct {
	BaseAsset               string `json:"base"`
	QuoteAsset              string `json:"quote"`
	CurrentPrice            string `json:"current_price"`
	PhysicalPrice           string `json:"physical_price"`
	PhysicalPriceCorrection string `json:"physical_price_correction"`
	MaxPriceStep            string `json:"max_price_step"`
	Policies
}

func (a *AssetPair) Populate(asset *core.AssetPair) {
	a.BaseAsset = asset.BaseAsset
	a.QuoteAsset = asset.QuoteAsset
	a.CurrentPrice = amount.String(asset.CurrentPrice)
	a.PhysicalPrice = amount.String(asset.PhysicalPrice)
	a.PhysicalPriceCorrection = amount.String(asset.PhysicalPriceCorrection)
	a.MaxPriceStep = amount.String(asset.MaxPriceStep)
	a.Policies.PopulateForAssetPair(*asset)
}
