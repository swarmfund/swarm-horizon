package horizon

import (
	"gitlab.com/swarmfund/horizon/db2"
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
)

type LedgerChangesAction struct {
	Action
	PagingParams db2.PageQuery
	Records      []history.Transaction
	Page         hal.Page
}

// JSON is a method for actions.JSON
func (action *LedgerChangesAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.checkAllowed,
		action.ValidateCursorWithinHistory,
		action.loadRecords,
		action.loadPage,
		func() { hal.Render(action.W, action.Page) },
	)
}

func (action *LedgerChangesAction) loadParams() {
	action.ValidateCursorAsDefault()
	action.PagingParams = action.GetPageQuery()
}

func (action *LedgerChangesAction) checkAllowed() {
	action.IsAllowed("")
}

func (action *LedgerChangesAction) loadRecords() {
	err := action.HistoryQ().Transactions().
		Page(action.PagingParams).
		Select(&action.Records)
	if err != nil {
		action.Log.WithError(err).Error("failed to load transaction records")
		action.Err = &problem.ServerError
	}
}

func (action *LedgerChangesAction) loadPage() {
	for _, record := range action.Records {
		var res resource.LedgerChanges
		if err := res.Populate(record); err != nil {
			action.Log.WithError(err).Error("failed to populate ledger changes")
			action.Err = &problem.ServerError
			return
		}

		action.Page.Add(res)
	}

	action.Page.BaseURL = action.BaseURL()
	action.Page.BasePath = action.Path()
	action.Page.Limit = action.PagingParams.Limit
	action.Page.Cursor = action.PagingParams.Cursor
	action.Page.Order = action.PagingParams.Order
	action.Page.PopulateLinks()
}
