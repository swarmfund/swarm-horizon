package ingest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/swarmfund/go/xdr"
)

func TestSetOperationDetails(t *testing.T) {
	source := "GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB"
	accountType := xdr.AccountTypeGeneral
	temp := [32]byte{
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 10,
		11, 12, 13, 14,
		15, 16, 17, 18,
		19, 20, 21, 22,
		23, 24, 25, 26,
		27, 28, 29, 30,
		31,
	}
	edd25519 := xdr.Uint256(temp)
	cases := []struct {
		name      string
		operation xdr.OperationBody
		result    xdr.OperationResultTr
	}{
		{
			name: "SetLimits",
			operation: xdr.OperationBody{
				Type: xdr.OperationTypeSetLimits,
				SetLimitsOp: &xdr.SetLimitsOp{
					Account: &xdr.AccountId{
						Type:    0,
						Ed25519: &edd25519,
					},
					AccountType: &accountType,
					Limits: xdr.Limits{
						DailyOut:   5,
						WeeklyOut:  35,
						MonthlyOut: 300,
						AnnualOut:  0,
						Ext: xdr.LimitsExt{
							V: 0,
						},
					},
					Ext: xdr.SetLimitsOpExt{
						V: 0,
					},
				},
			},
			result: xdr.OperationResultTr{
				Type: xdr.OperationTypeSetLimits,
				SetLimitsResult: &xdr.SetLimitsResult{
					Code:    5,
					Success: nil,
				},
			},
		},
		{
			name: "SetFee",
			operation: xdr.OperationBody{
				Type: xdr.OperationTypeSetFees,
				SetFeesOp: &xdr.SetFeesOp{
					Fee: &xdr.FeeEntry{
						FeeType:     1,
						Asset:       "USD",
						FixedFee:    0,
						PercentFee:  0,
						AccountId:   nil,
						AccountType: &accountType,
						Subtype:     2,
						LowerBound:  0,
						UpperBound:  5,
						Hash:        xdr.Hash{},
						Ext:         xdr.FeeEntryExt{},
					},
					IsDelete: false,
					Ext:      xdr.SetFeesOpExt{},
				},
			},
			result: xdr.OperationResultTr{
				Type: xdr.OperationTypeSetFees,
				SetFeesResult: &xdr.SetFeesResult{
					Code:    5,
					Success: nil,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				setOperationDetails(source, c.operation, &c.result)
			})
		})
	}
}
