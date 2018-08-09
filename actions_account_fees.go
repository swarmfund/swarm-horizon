package horizon

import (
	"gitlab.com/swarmfund/horizon/db2/core"
	. "gitlab.com/swarmfund/horizon/fees"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/regources"
)

// This file contains the actions:
//
// FeesForAccount: renders all the fees for a specific account
//FeesForAccount show all fees for account

type AccountFeesAction struct {
	Action
	Account *core.Account

	Records  SmartFeeTable
	Assets   []string
	Response resource.FeesResponse
}

func (action *AccountFeesAction) JSON() {
	action.Do(
		action.loadParams,
		action.loadAssets,
		action.loadFees,
		action.loadResponse,
		func() {
			hal.Render(action.W, action.Response)
		},
	)
}

func (action *AccountFeesAction) loadParams() {
	acc := action.GetCoreAccount("account_id", action.CoreQ())
	if acc != nil && action.Err == nil {
		action.Account = acc
	}
}

func (action *AccountFeesAction) loadAssets() {
	accountBalances, err := action.CoreQ().Balances().ByAddress(action.Account.AccountID).Select()
	if err != nil {
		action.Log.WithError(err).Error("failed to load balances")
		action.Err = &problem.ServerError
		return
	}

	var assets []string
	for _, balance := range accountBalances {
		assets = append(assets, balance.Asset)
	}
	action.Assets = assets
}

func (action *AccountFeesAction) loadFees() {
	forAccount, err := action.CoreQ().FeeEntries().ForAccount(action.Account.AccountID).Select()
	if err != nil {
		action.Err = &problem.ServerError
		action.Log.WithError(err).Error("can't get account fees from the database")
		return
	}

	forAccountType, err := action.CoreQ().FeeEntries().ForAccountType(&action.Account.AccountType).Select()
	if err != nil {
		action.Err = &problem.ServerError
		action.Log.WithError(err).Error("can't get account type fees from the database")
		return

	}

	//get general fees set for all, not to be confused with fees for General Account Type
	generalFees, err := action.CoreQ().FeeEntries().ForAccountType(nil).Select()
	if err != nil {
		action.Err = &problem.ServerError
		action.Log.WithError(err).Error("can't get general fees from the database")
		return
	}

	sft := NewSmartFeeTable(forAccount)
	sft.Update(forAccountType)
	sft.Update(generalFees)
	sft.AddZeroFees(action.Assets)
	action.Records = sft
}

func (action *AccountFeesAction) loadResponse() {
	byAssets := action.Records.GetValuesByAsset()
	action.Response.Fees = make(map[xdr.AssetCode][]regources.FeeEntry)
	for _, feesForAsset := range byAssets {
		for _, wrapper := range feesForAsset {
			fee := resource.NewFeeEntryFromWrapper(wrapper)
			ac := xdr.AssetCode(wrapper.Asset)
			action.Response.Fees[ac] = append(action.Response.Fees[ac], fee)
		}
	}
}
