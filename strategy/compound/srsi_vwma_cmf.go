// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package compound

import (
	"fmt"
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/strategy"
	"github.com/cinar/indicator/v2/strategy/momentum"
	"github.com/cinar/indicator/v2/strategy/trend"
)

// MacdRsiStrategy represents the configuration parameters for calculating the MACD-RSI strategy.
type SrisVwmaStrategy struct {
	VwmaStrategy *trend.VwmaStrategy

	StRsiStrategy *momentum.StochasticRsiStrategy
}

// NewMacdRsiStrategy function initializes a new MACD-RSI strategy instance with the default parameters.
func NewSrisVwmaStrategy() *SrisVwmaStrategy {
	return NewSrisVwmaStrategy()
}

// NewMacdRsiStrategyWith function initializes a new MACD-RSI strategy instance with the given parameters.
func NewMSrisVwmaCmfStrategy() *SrisVwmaStrategy {
	return &SrisVwmaStrategy{
		VwmaStrategy:  trend.NewVwmaStrategy(),
		StRsiStrategy: momentum.NewStochasticRsiStrategy(),
	}
}

// Name returns the name of the strategy.
func (m *SrisVwmaStrategy) Name() string {
	return fmt.Sprintf("SrisVwma Strategy")

}

// Compute processes the provided asset snapshots and generates a stream of actionable recommendations.
func (m *SrisVwmaStrategy) Compute(snapshots <-chan *asset.Snapshot) <-chan strategy.Action {
	snapshotsSplice := helper.Duplicate(snapshots, 2)

	vwmaActions := strategy.DenormalizeActions(
		m.VwmaStrategy.Compute(snapshotsSplice[0]),
	)

	rsiActions := strategy.DenormalizeActions(
		m.StRsiStrategy.Compute(snapshotsSplice[1]),
	)

	actions := helper.Operate(vwmaActions, rsiActions, func(macdAction, rsiAction strategy.Action) strategy.Action {
		if macdAction == rsiAction {
			return macdAction
		}

		return strategy.Hold
	})

	return actions
}

// Report processes the provided asset snapshots and generates a report annotated with the recommended actions.
func (m *SrisVwmaStrategy) Report(c <-chan *asset.Snapshot) *helper.Report {
	////
	//// snapshots[0] -> dates
	//// snapshots[1] -> closings[0] -> closings
	////                 closings[1] -> macds, signals
	////                 closings[2] -> rsi
	//// snapshots[2] -> actions     -> annotations
	////              -> outcomes
	////
	//snapshots := helper.Duplicate(c, 3)
	//
	//dates := asset.SnapshotsAsDates(snapshots[0])
	//closings := helper.Duplicate(asset.SnapshotsAsClosings(snapshots[1]), 3)
	//
	//macds, signals := m.MacdStrategy.Macd.Compute(closings[0])
	//macds = helper.Shift(macds, m.MacdStrategy.Macd.IdlePeriod(), 0)
	//signals = helper.Shift(signals, m.MacdStrategy.Macd.IdlePeriod(), 0)
	//
	//rsi := m.RsiStrategy.Rsi.Compute(closings[2])
	//rsi = helper.Shift(rsi, m.RsiStrategy.Rsi.IdlePeriod(), 0)
	//
	//actions, outcomes := strategy.ComputeWithOutcome(m, snapshots[2])
	//annotations := strategy.ActionsToAnnotations(actions)
	//outcomes = helper.MultiplyBy(outcomes, 100)
	//
	//report := helper.NewReport(m.Name(), dates)
	//report.AddChart()
	//report.AddChart()
	//report.AddChart()
	//
	//report.AddColumn(helper.NewNumericReportColumn("Close", closings[1]))
	//report.AddColumn(helper.NewNumericReportColumn("MACD", macds), 1)
	//report.AddColumn(helper.NewNumericReportColumn("Signal", signals), 1)
	//
	//report.AddColumn(helper.NewNumericReportColumn("RSI", rsi), 2)
	//report.AddColumn(helper.NewAnnotationReportColumn(annotations), 0, 1, 2)
	//
	//report.AddColumn(helper.NewNumericReportColumn("Outcome", outcomes), 3)

	return nil
}
