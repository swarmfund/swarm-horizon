package horizon

import (
	"time"

	"gitlab.com/swarmfund/horizon/db2/core"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
)

// This file contains the actions:
//
// AccountShowAction: details for single account (including stellar-core state)

// AccountShowAction renders a account summary found by its address.
type AccountShowAction struct {
	Action
	Address     string
	Resource    resource.Account
}

// JSON is a method for actions.JSON
func (action *AccountShowAction) JSON() {
	action.Do(
		action.loadParams,
		action.checkAllowed,
		action.loadRecord,
		action.loadLimits,
		action.loadBalances,
		action.loadExternalSystemAccountIDs,
		func() {
			hal.Render(action.W, action.Resource)
		},
	)
}

func (action *AccountShowAction) loadParams() {
	action.Address = action.GetString("id")
}

func (action *AccountShowAction) checkAllowed() {
	action.IsAllowed(action.Address)
}

func (action *AccountShowAction) loadRecord() {
	coreRecord, err := action.CoreQ().
		Accounts().
		ForAddresses(action.Address).
		WithStatistics().
		First()

	if err != nil {
		action.Log.WithError(err).Error("Failed to get account from core DB")
		action.Err = &problem.ServerError
		return
	}

	if coreRecord == nil {
		action.Err = &problem.NotFound
		return
	}

	coreRecord.Statistics.ClearObsolete(time.Now().UTC())

	action.Resource.Populate(action.Ctx, *coreRecord)

	signers, err := action.GetSigners(coreRecord)
	if err != nil {
		action.Log.WithError(err).Error("Failed to get signers")
		action.Err = &problem.ServerError
		return
	}

	action.Resource.Signers.Populate(signers)
}

func (action *AccountShowAction) loadLimits() {
	limits, err := action.CoreQ().LimitsForAccount(action.Address, action.Resource.AccountTypeI)
	if err != nil {
		action.Log.WithError(err).Error("Failed to load limits for account")
		action.Err = &problem.ServerError
		return
	}

	action.Resource.Limits.Populate(limits)
}

func (action *AccountShowAction) loadBalances() {
	var balances []core.Balance
	err := action.CoreQ().
		BalancesByAddress(&balances, action.Address)
	if err != nil {
		action.Log.WithError(err).Error("Failed to get balances for account")
		action.Err = &problem.ServerError
		return
	}

	action.Resource.SetBalances(balances)
}

func (action *AccountShowAction) loadExternalSystemAccountIDs() {
	exSysIDs, err := action.CoreQ().ExternalSystemAccountID().ForAccount(action.Address).Select()
	if err != nil {
		action.Log.WithError(err).Error("Failed to load external system account ids")
		action.Err = &problem.ServerError
		return
	}

	action.Resource.ExternalSystemAccounts = make([]resource.ExternalSystemAccountID, len(exSysIDs))
	for i := range exSysIDs {
		action.Resource.ExternalSystemAccounts[i].Populate(exSysIDs[i])
	}
}
