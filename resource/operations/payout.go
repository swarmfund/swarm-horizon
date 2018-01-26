package operations

type Payout struct {
	Base
	Asset              string `json:"asset"`
	SourceBalanceID    string `json:"source_balance_id"`
	MaxPayoutAmount    string `json:"max_payout_amount"`
	ActualPayoutAmount string `json:"actual_payout_amount"`
	FixedFee           string `json:"fixed_fee"`
	PercentFee         string `json:"percent_fee"`
}
