package horizon

import (
	"gitlab.com/swarmfund/horizon/db2"
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource/reviewablerequest"
)

// ReviewableRequestIndexAction renders slice of reviewable requests
type ReviewableRequestIndexAction struct {
	Action
	Reviewer     string
	Requestor    string
	Asset        string
	State        *int64
	Type         *int64
	Records      []history.ReviewableRequest
	PagingParams db2.PageQuery
	Page         hal.Page
}

// JSON is a method for actions.JSON
func (action *ReviewableRequestIndexAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.checkAllowed,
		action.loadRecord,
		action.loadPage,
		func() {
			hal.Render(action.W, action.Page)
		},
	)
}

func (action *ReviewableRequestIndexAction) loadParams() {
	action.PagingParams = action.GetPageQuery()
	action.Reviewer = action.GetString("reviewer")
	action.Requestor = action.GetString("requestor")
	action.State = action.GetOptionalInt64("state")
	action.Type = action.GetOptionalInt64("type")
	action.Asset = action.GetString("asset")
}

func (action *ReviewableRequestIndexAction) checkAllowed() {
	action.IsAllowed(action.Requestor, action.Reviewer)
}

func (action *ReviewableRequestIndexAction) loadRecord() {
	q := action.HistoryQ().ReviewableRequests()

	if action.Reviewer != "" {
		q = q.ForReviewer(action.Reviewer)
	}

	if action.Requestor != "" {
		q = q.ForRequestor(action.Requestor)
	}

	if action.Asset != "" {
		q = q.ForAsset(action.Asset)
	}

	if action.State != nil {
		q = q.ForState(*action.State)
	}

	if action.Type != nil {
		q = q.ForType(*action.Type)
	}

	q = q.Page(action.PagingParams)
	var err error
	action.Records, err = q.Select()
	if err != nil {
		action.Log.WithError(err).Error("failed to load reviewable requests")
		action.Err = &problem.ServerError
		return
	}
}

func (action *ReviewableRequestIndexAction) loadPage() {
	for i := range action.Records {
		var res reviewablerequest.ReviewableRequest
		res.Populate(&action.Records[i])
		action.Page.Add(&res)
	}

	action.Page.BaseURL = action.BaseURL()
	action.Page.BasePath = action.Path()
	action.Page.Limit = action.PagingParams.Limit
	action.Page.Cursor = action.PagingParams.Cursor
	action.Page.Order = action.PagingParams.Order
	action.Page.PopulateLinks()
}
