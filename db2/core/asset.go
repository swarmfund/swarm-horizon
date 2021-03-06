package core

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/swarmfund/horizon/db2"
)

type Asset struct {
	Code                 string `db:"code"`
	Policies             int32  `db:"policies"`
	Owner                string `db:"owner"`
	AvailableForIssuance uint64 `db:"available_for_issueance"`
	PreissuedAssetSigner string `db:"preissued_asset_signer"`
	MaxIssuanceAmount    uint64 `db:"max_issuance_amount"`
	Issued               uint64 `db:"issued"`
	LockedIssuance       uint64 `db:"locked_issuance"`
	PendingIssuance      uint64 `db:"pending_issuance"`
	Details              []byte `db:"details"`
}

func (a Asset) GetDetails() (db2.Details, error) {
	var result db2.Details
	err := json.Unmarshal(a.Details, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal asset details")
	}

	return result, nil
}
