package reviewablerequest

import (
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/tokend/regources"
	"gitlab.com/tokend/go/amount"
)

func PopulateIssuanceRequest(histRequest history.IssuanceRequest) (
	r *regources.IssuanceRequest, err error,
) {
	r = &regources.IssuanceRequest{}
	r.Asset = histRequest.Asset
	r.Amount = regources.Amount(amount.MustParse(histRequest.Amount))
	r.Receiver = histRequest.Receiver
	r.ExternalDetails = histRequest.ExternalDetails
	return
}
