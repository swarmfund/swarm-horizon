package horizon

import (
	"gitlab.com/swarmfund/go/amount"
	"gitlab.com/swarmfund/horizon/charts"
	"gitlab.com/swarmfund/horizon/render/hal"
	"gitlab.com/swarmfund/horizon/render/problem"
	"gitlab.com/swarmfund/horizon/resource"
)

type ChartsAction struct {
	Action

	Code string

	Record   map[string]*charts.Histogram
	Resource resource.Charts
}

func (action *ChartsAction) JSON() {
	action.Do(
		action.loadParams,
		action.loadChart,
		action.renderResource,
		func() {
			hal.Render(action.W, action.Resource)
		},
	)
}

func (action *ChartsAction) loadParams() {
	action.Code = action.GetNonEmptyString("code")
}

func (action *ChartsAction) loadChart() {
	action.Record = action.App.charts.Get(action.Code)
	if action.Record == nil {
		action.Err = &problem.NotFound
		return
	}
}

func (action *ChartsAction) renderResource() {
	action.Resource = make(resource.Charts)
	for key, histogram := range action.Record {
		points := histogram.Render()
		for _, point := range points {
			action.Resource[key] = append(action.Resource[key], resource.Point{
				Timestamp: point.Timestamp,
				Value:     amount.String(*point.Value),
			})
		}
	}
}