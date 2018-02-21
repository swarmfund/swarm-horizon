package horizon

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/swarmfund/go/amount"
	"gitlab.com/swarmfund/horizon/db2/core"
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/swarmfund/horizon/exchange"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
)

// SaleShowAction renders a sale found by its ID.
type SaleShowAction struct {
	Action
	RequestID uint64
	Record    *history.Sale
	offers    []core.Offer
	balances  []core.Balance
	result    resource.Sale
}

// JSON is a method for actions.JSON
func (action *SaleShowAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.loadRecord,
		action.populateResult,
		func() {
			hal.Render(action.W, action.result)
		},
	)
}

func (action *SaleShowAction) loadParams() {
	action.RequestID = action.GetUInt64("id")
}

func (action *SaleShowAction) loadRecord() {
	var err error
	action.Record, err = action.HistoryQ().Sales().ByID(action.RequestID)
	if err != nil {
		action.Log.WithError(err).
			WithField("request_id", action.RequestID).
			Error("failed to load sale")
		action.Err = &problem.ServerError
		return
	}

	if action.Record == nil {
		action.Err = &problem.NotFound
		return
	}

	err = action.populateTotalCurrentCap()
	if err != nil {
		action.Log.WithError(err).Error("failed to populate total current cap")
		action.Err = &problem.ServerError
		return
	}

	action.offers = make([]core.Offer, 0)
	err = action.CoreQ().Offers().
		ForOrderBookID(action.Record.ID).Select(&action.offers)
	if err != nil {
		action.Log.WithError(err).
			WithField("sale_id", action.Record.ID).
			Error("failed to load offers for sale")
		action.Err = &problem.ServerError
		return
	}

	action.balances, err = action.CoreQ().Balances().
		ByAsset(action.Record.BaseAsset).Select()
	if err != nil {
		action.Log.WithError(err).
			WithField("sale_id", action.Record.ID).
			Error("failed to load base asset balances for sale")
		action.Err = &problem.ServerError
		return
	}
}

func (action *SaleShowAction) populateTotalCurrentCap() error {
	converter, err := exchange.NewConverter(action.CoreQ())
	if err != nil {
		return errors.Wrap(err, "failed to init converter")
	}

	totalCapInDefaultQuoteAsset, err := getCurrentCapInDefaultQuote(action.Record, converter)
	if err != nil {
		return errors.Wrap(err, "failed to get current cap in default quote asset")
	}

	action.Record.CurrentCap = amount.String(totalCapInDefaultQuoteAsset)

	for i := range action.Record.QuoteAssets.QuoteAssets {
		quoteAsset := &action.Record.QuoteAssets.QuoteAssets[i]
		totalCapInQuote, err := converter.TryToConvertWithOneHop(totalCapInDefaultQuoteAsset, action.Record.DefaultQuoteAsset, quoteAsset.Asset)
		if err != nil {
			return errors.Wrap(err, "failed to convert total cap in default to quote")
		}

		if totalCapInQuote == nil {
			return errors.New("failed to convert total cap in default to quote: failed to find path")
		}

		quoteAsset.TotalCurrentCap = amount.String(*totalCapInQuote)

		hardCapInQuote, err := converter.TryToConvertWithOneHop(int64(action.Record.HardCap), action.Record.DefaultQuoteAsset, quoteAsset.Asset)
		if err != nil {
			return errors.Wrap(err, "failed to convert hard cap")
		}

		if hardCapInQuote == nil {
			return errors.New("failed to convert hard cap to quote asset")
		}

		quoteAsset.HardCap = amount.String(*hardCapInQuote)
	}

	return nil
}

func (action *SaleShowAction) populateResult() {
	action.result.Populate(action.Record)
	err := action.result.PopulateStat(action.offers, action.balances)
	if err != nil {
		action.Log.WithError(err).
			WithField("request_id", action.RequestID).
			Error("failed to populate stat for sale")
		action.Err = &problem.ServerError
		return
	}
}

func getCurrentCapInDefaultQuote(sale *history.Sale, converter *exchange.Converter) (int64, error) {
	totalCapInDefaultQuoteAsset := int64(0)
	for _, quoteAsset := range sale.QuoteAssets.QuoteAssets {
		currentCap, err := amount.Parse(quoteAsset.CurrentCap)
		if err != nil {
			return 0, errors.Wrap(err, "failed to parse current cap")
		}

		currentCapInDefaultQuoteAsset, err := converter.TryToConvertWithOneHop(currentCap, quoteAsset.Asset, sale.DefaultQuoteAsset)
		if err != nil {
			return 0, errors.Wrap(err, "failed to convert current cap to default quote asset")
		}

		if currentCapInDefaultQuoteAsset == nil {
			return 0, errors.New("failed to convert to current cap: no path found")
		}

		var isOk bool
		totalCapInDefaultQuoteAsset, isOk = amount.SafePositiveSum(totalCapInDefaultQuoteAsset, *currentCapInDefaultQuoteAsset)
		if !isOk {
			return 0, errors.New("failed to find total cap in default quote asset: overflow")
		}
	}

	return totalCapInDefaultQuoteAsset, nil
}
