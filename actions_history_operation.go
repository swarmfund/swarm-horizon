package horizon

import (
	"gitlab.com/swarmfund/horizon/db2"
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
	"gitlab.com/tokend/go/xdr"
)

// This file contains the actions:
//
// OperationIndexAction: pages of operations
// OperationShowAction: single operation by id

// OperationIndexAction renders a page of operations resources, identified by
// a normal page query and optionally filtered by an account, ledger, or
// transaction.
type HistoryOperationIndexAction struct {
	Action
	Types        []xdr.OperationType
	PagingParams db2.PageQuery
	Records      []history.Operation
	Participants map[int64]*history.OperationParticipants
	Page         hal.Page
	AssetFilter  string
}

// JSON is a method for actions.JSON
func (action *HistoryOperationIndexAction) JSON() {
	action.Do(
		action.loadParams,
		action.loadRecords,
		action.loadParticipants,
		action.loadPage,
		func() {
			hal.Render(action.W, action.Page)
		},
	)
}

func (action *HistoryOperationIndexAction) loadParams() {
	action.AssetFilter = action.GetString("asset")
	action.PagingParams = action.GetPageQuery()
}

func (action *HistoryOperationIndexAction) loadRecords() {
	ops := action.HistoryQ().Operations()

	if action.AssetFilter != "" {
		ops.JoinOnBalance().ForAsset(action.AssetFilter)
	}

	if len(action.Types) > 0 {
		ops.ForTypes(action.Types)
	}

	err := ops.Page(action.PagingParams).Select(&action.Records)

	if err != nil {
		action.Log.WithError(err).Error("failed to get operations")
		action.Err = &problem.ServerError
		return
	}
}

// loadParticipants for this action is needed only for asset provenance feature
func (action *HistoryOperationIndexAction) loadParticipants() {
	isPublicAsset, err := action.isPublicAsset()
	if err != nil {
		action.Log.WithError(err).Error("failed to check if asset is public")
		action.Err = &problem.ServerError
		return
	}
	if !isPublicAsset {
		return
	}

	// initializing our operation -> participants map
	action.Participants = map[int64]*history.OperationParticipants{}
	for _, operation := range action.Records {
		if operation.Type != xdr.OperationTypeManageOffer {
			continue
		}
		action.Participants[operation.ID] = &history.OperationParticipants{
			operation.Type,
			[]*history.Participant{},
		}
	}

	// workaround for load participants
	action.IsAdmin = true
	action.LoadParticipants("", action.Participants)
	// reverting workaround, just in case
	action.IsAdmin = false
}

// isPublicAsset checks for asset details to ensure if it's allowed to expose manageOfferParticipants. Client-app
// defines if the asset is public.
func (action *HistoryOperationIndexAction) isPublicAsset() (bool, error) {
	// provenance is considered to be used only when we need to get the trace for specific asset,
	// so if no such filter is present we can omit populating participants
	if action.AssetFilter == "" {
		return false, nil
	}

	asset, err := action.CoreQ().Assets().ByCode(action.AssetFilter)
	if err != nil {
		return false, err
	}
	if asset == nil {
		return false, nil
	}

	details, err := asset.GetDetails()
	if err != nil {
		return false, err
	}

	// isPublic field is being set by client-app
	if details["isPublic"] != true {
		return false, nil
	}

	return true, nil
}

func (action *HistoryOperationIndexAction) loadPage() {
	for _, record := range action.Records {
		var res hal.Pageable
		opParticipants := action.Participants[record.ID]
		if opParticipants != nil {
			// HACK: call `NewOperation` to expose participants ids for asset traceability.
			res, action.Err = resource.NewOperation(action.Ctx, record, opParticipants.Participants)
		} else {
			res, action.Err = resource.NewPublicOperation(action.Ctx, record, nil)
		}
		if action.Err != nil {
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

// OperationShowAction renders a ledger found by its sequence number.
type HistoryOperationShowAction struct {
	Action
	ID           int64
	Record       history.Operation
	Resource     interface{}
	Participants map[int64]*history.OperationParticipants
}

func (action *HistoryOperationShowAction) loadParams() {
	action.ID = action.GetInt64("id")
}

func (action *HistoryOperationShowAction) loadRecord() {
	action.Err = action.HistoryQ().OperationByID(&action.Record, action.ID)
	if action.Err != nil {
		return
	}

	action.Participants = map[int64]*history.OperationParticipants{
		action.ID: {},
	}
	switch action.Record.Type {
	case xdr.OperationTypeManageOffer:
		// workaround for load participants
		action.IsAdmin = true
		action.LoadParticipants("", action.Participants)
		// reverting workaround, just in case
		action.IsAdmin = false
	}
}

func (action *HistoryOperationShowAction) loadResource() {
	action.Resource, action.Err = resource.NewPublicOperation(
		action.Ctx, action.Record, action.Participants[action.Record.ID].Participants)
}

// JSON is a method for actions.JSON
func (action *HistoryOperationShowAction) JSON() {
	action.Do(
		action.loadParams,
		action.loadRecord,
		action.loadResource,
		func() {
			hal.Render(action.W, action.Resource)
		},
	)
}
