package ingest

import (
	"fmt"

	"encoding/hex"
	"encoding/json"

	"gitlab.com/swarmfund/go/amount"
	"gitlab.com/swarmfund/go/xdr"
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/swarmfund/horizon/utf8"
)

// operationDetails returns the details regarding the current operation, suitable
// for ingestion into a history_operation row
func (is *Session) operationDetails() history.OperationDetails {
	source := is.Cursor.OperationSourceAccount()
	return setOperationDetails(
		source.Address(),
		is.Cursor.Operation().Body,
		is.Cursor.OperationResult(),
	)
}

func setOperationDetails(source string, operation xdr.OperationBody,
	result *xdr.OperationResultTr) history.OperationDetails {

	operationDetails := history.OperationDetails{
		Type: operation.Type,
	}

	switch operationDetails.Type {
	case xdr.OperationTypeCreateAccount:
		op := operation.MustCreateAccountOp()
		operationDetails.CreateAccount = &history.CreateAccountDetails{
			Funder:      source,
			Account:     op.Destination.Address(),
			AccountType: int32(op.AccountType),
		}
	case xdr.OperationTypePayment:
		op := operation.MustPaymentOp()
		opResult := result.MustPaymentResult()

		operationDetails.Payment = &history.PaymentDetails{
			BasePayment: history.BasePayment{
				From:                  source,
				To:                    opResult.PaymentResponse.Destination.Address(),
				FromBalance:           op.SourceBalanceId.AsString(),
				ToBalance:             op.DestinationBalanceId.AsString(),
				Amount:                amount.String(int64(op.Amount)),
				Asset:                 string(opResult.PaymentResponse.Asset),
				SourcePaymentFee:      amount.String(int64(op.FeeData.SourceFee.PaymentFee)),
				DestinationPaymentFee: amount.String(int64(op.FeeData.DestinationFee.PaymentFee)),
				SourceFixedFee:        amount.String(int64(op.FeeData.SourceFee.FixedFee)),
				DestinationFixedFee:   amount.String(int64(op.FeeData.DestinationFee.FixedFee)),
				SourcePaysForDest:     op.FeeData.SourcePaysForDest,
			},
			Subject:    string(op.Subject),
			Reference:  string(op.Reference),
			QuoteAsset: string(opResult.PaymentResponse.Asset),
		}
	case xdr.OperationTypeSetOptions:
		op := operation.MustSetOptionsOp()

		operationDetails.SetOptions = &history.SetOptionsDetails{}

		if op.MasterWeight != nil {
			operationDetails.SetOptions.MasterKeyWeight = (*uint32)(op.MasterWeight)
		}

		if op.LowThreshold != nil {
			operationDetails.SetOptions.LowThreshold = (*uint32)(op.LowThreshold)
		}

		if op.MedThreshold != nil {
			operationDetails.SetOptions.MedThreshold = (*uint32)(op.MedThreshold)
		}

		if op.HighThreshold != nil {
			operationDetails.SetOptions.HighThreshold = (*uint32)(op.HighThreshold)
		}

		if op.Signer != nil {
			operationDetails.SetOptions.SignerKey = op.Signer.PubKey.Address()
		}

		if op.LimitsUpdateRequestData != nil {
			operationDetails.SetOptions.LimitsUpdateRequestDocumentHash = hex.EncodeToString(op.LimitsUpdateRequestData.DocumentHash[:])
		}
	case xdr.OperationTypeSetFees:
		op := operation.MustSetFeesOp()

		if op.Fee != nil {
			var accountID *string
			if op.Fee.AccountId != nil {
				t := op.Fee.AccountId.Address()
				accountID = &t
			}

			var accountType *int32
			if op.Fee.AccountType != nil {
				accountType = (*int32)(op.Fee.AccountType)
			}

			operationDetails.SetFees = &history.SetFeesDetails{
				Fee: &history.FeeDetails{
					AssetCode:   string(op.Fee.Asset),
					FixedFee:    amount.String(int64(op.Fee.FixedFee)),
					PercentFee:  amount.String(int64(op.Fee.PercentFee)),
					FeeType:     int64(op.Fee.FeeType),
					AccountID:   accountID,
					AccountType: accountType,
					Subtype:     int64(op.Fee.Subtype),
					LowerBound:  int64(op.Fee.LowerBound),
					UpperBound:  int64(op.Fee.UpperBound),
				},
			}
		}
	case xdr.OperationTypeManageAccount:
		op := operation.MustManageAccountOp()

		operationDetails.ManageAccount = &history.ManageAccountDetails{
			Account:              op.Account.Address(),
			BlockReasonsToAdd:    uint32(op.BlockReasonsToAdd),
			BlockReasonsToRemove: uint32(op.BlockReasonsToRemove),
		}
	case xdr.OperationTypeCreateWithdrawalRequest:
		op := operation.MustCreateWithdrawalRequestOp()
		request := op.Request

		var externalDetails map[string]interface{}
		// error is ignored on purpose, we should not block ingest in case of such error
		_ = json.Unmarshal([]byte(request.ExternalDetails), &externalDetails)

		operationDetails.CreateWithdrawalRequest = &history.CreateWithdrawalRequestDetails{
			Amount:          amount.StringU(uint64(request.Amount)),
			Balance:         request.Balance.AsString(),
			FeeFixed:        amount.StringU(uint64(request.Fee.Fixed)),
			FeePercent:      amount.StringU(uint64(request.Fee.Percent)),
			ExternalDetails: externalDetails,
			DestAsset:       string(request.Details.AutoConversion.DestAsset),
			DestAmount:      amount.StringU(uint64(request.Details.AutoConversion.ExpectedAmount)),
		}
	case xdr.OperationTypeManageBalance:
		op := operation.MustManageBalanceOp()

		//added new struct in resource/main.go and in OperationDetails
		operationDetails.ManageBalance = &history.ManageBalanceDetails{
			Destination: op.Destination.Address(),
			Action:      int32(op.Action),
		}
	case xdr.OperationTypeReviewPaymentRequest:
		op := operation.MustReviewPaymentRequestOp()

		var rejectReason *string
		if op.RejectReason != nil {
			rejectReason = (*string)(op.RejectReason)
		}

		operationDetails.ReviewPaymentRequest = &history.ReviewPaymentRequestDetails{
			PaymentID:    int64(op.PaymentId),
			Accept:       op.Accept,
			RejectReason: rejectReason,
		}
	case xdr.OperationTypeSetLimits:
		op := operation.MustSetLimitsOp()

		operationDetails.SetLimits = &history.SetLimitsDetails{
			AccountType: nil,
			Account:     op.Account.Address(),
		}

		if accountType := op.AccountType; accountType != nil {
			operationDetails.SetLimits.AccountType = (*int32)(accountType)
		}

	case xdr.OperationTypeDirectDebit:
		op := operation.MustDirectDebitOp().PaymentOp
		opResult := result.MustDirectDebitResult().MustSuccess()

		operationDetails.DirectDebit = &history.DirectDebitDetails{
			From:                  source,
			To:                    opResult.PaymentResponse.Destination.Address(),
			FromBalance:           op.SourceBalanceId.AsString(),
			ToBalance:             op.DestinationBalanceId.AsString(),
			Amount:                amount.String(int64(op.Amount)),
			SourcePaymentFee:      amount.String(int64(op.FeeData.SourceFee.PaymentFee)),
			DestinationPaymentFee: amount.String(int64(op.FeeData.DestinationFee.PaymentFee)),
			SourceFixedFee:        amount.String(int64(op.FeeData.SourceFee.FixedFee)),
			DestinationFixedFee:   amount.String(int64(op.FeeData.DestinationFee.FixedFee)),
			SourcePaysForDest:     op.FeeData.SourcePaysForDest,
			Subject:               string(op.Subject),
			Reference:             string(op.Reference),
			AssetCode:             string(opResult.PaymentResponse.Asset),
		}
	case xdr.OperationTypeManageAssetPair:
		op := operation.MustManageAssetPairOp()

		operationDetails.ManageAssetPair = &history.ManageAssetPairDetails{
			BaseAsset:               string(op.Base),
			QuoteAsset:              string(op.Quote),
			PhysicalPrice:           amount.String(int64(op.PhysicalPrice)),
			PhysicalPriceCorrection: amount.String(int64(op.PhysicalPriceCorrection)),
			MaxPriceStep:            amount.String(int64(op.MaxPriceStep)),
			Policies:                int32(op.Policies),
		}
	case xdr.OperationTypeManageOffer:
		op := operation.ManageOfferOp

		operationDetails.ManagerOffer = &history.ManagerOfferDetails{
			IsBuy:     op.IsBuy,
			Amount:    amount.String(int64(op.Amount)),
			Price:     amount.String(int64(op.Price)),
			Fee:       amount.String(int64(op.Fee)),
			OfferId:   uint64(op.OfferId),
			IsDeleted: int64(op.OfferId) != 0,
		}
	case xdr.OperationTypeManageInvoice:
		op := operation.MustManageInvoiceOp()
		opResult := result.MustManageInvoiceResult()

		operationDetails.ManageInvoice = &history.ManageInvoiceDetails{
			Amount:          amount.String(int64(op.Amount)),
			ReceiverBalance: op.ReceiverBalance.AsString(),
			Sender:          op.Sender.Address(),
			InvoiceID:       uint64(opResult.Success.InvoiceId),
			Asset:           string(opResult.Success.Asset),
		}
	case xdr.OperationTypeReviewRequest:
		op := operation.MustReviewRequestOp()

		operationDetails.ReviewRequest = &history.ReviewRequestDetails{
			Action:      int32(op.Action),
			Reason:      string(op.Reason),
			RequestHash: hex.EncodeToString(op.RequestHash[:]),
			RequestID:   uint64(op.RequestId),
			RequestType: int32(op.RequestDetails.RequestType),
		}
	case xdr.OperationTypeManageAsset:
		op := operation.MustManageAssetOp()

		operationDetails.ManageAsset = &history.ManageAssetDetails{
			RequestID: uint64(op.RequestId),
			Action:    int32(op.Request.Action),
		}
	case xdr.OperationTypeCreatePreissuanceRequest:
		// no details needed
	case xdr.OperationTypeCreateIssuanceRequest:
		op := operation.MustCreateIssuanceRequestOp()

		var externalDetails map[string]interface{}
		// error is ignored on purpose, we should not block ingest in case of such error
		_ = json.Unmarshal([]byte(op.Request.ExternalDetails), &externalDetails)

		operationDetails.CreateIssuanceRequest = &history.CreateIssuanceRequestDetails{
			Reference:       utf8.Scrub(string(op.Reference)),
			Amount:          amount.StringU(uint64(op.Request.Amount)),
			Asset:           string(op.Request.Asset),
			FeeFixed:        amount.StringU(uint64(op.Request.Fee.Fixed)),
			FeePercent:      amount.StringU(uint64(op.Request.Fee.Percent)),
			BalanceID:       op.Request.Receiver.AsString(),
			ExternalDetails: externalDetails,
		}
	case xdr.OperationTypeCreateSaleRequest:
		// no details needed
	case xdr.OperationTypeCheckSaleState:
		// no details needed
	default:
		panic(fmt.Errorf("Unknown operation type: %s", operation.Type))
	}
	return operationDetails
}
