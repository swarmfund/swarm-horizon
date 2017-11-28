package horizon

import (
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
	"gitlab.com/swarmfund/horizon/db2/history"
)

type AccountBalancesAction struct {
	Action

	AccountID string

	Records []history.Balance
	Resource []resource.BalancePublic
}

func (action *AccountBalancesAction) JSON() {
	action.Do(
		action.loadParams,
		action.loadRecords,
		action.loadResource,
		func () {
			hal.Render(action.W, action.Resource)
		},
	)
}

func (action *AccountBalancesAction) loadParams() {
	action.AccountID = action.GetNonEmptyString("id")
}

func (action *AccountBalancesAction) loadRecords() {
	if err := action.HistoryQ().Balances().ForAccount(action.AccountID).Select(&action.Records); err != nil {
		action.Log.WithError(err).Error("failed to get balances")
		action.Err = &problem.ServerError
		return
	}

	if len(action.Records) == 0 {
		action.Err = &problem.NotFound
		return
	}
}

func (action *AccountBalancesAction) loadResource() {
	for _, record := range action.Records {
		var r resource.BalancePublic
		r.Populate(record)
		action.Resource = append(action.Resource, r)
	}
}


