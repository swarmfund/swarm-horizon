package ingest

import (
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/swarmfund/go/xdr"
	"gitlab.com/swarmfund/horizon/db2/history"
	"gitlab.com/swarmfund/horizon/ingest/participants"
)

// Run starts an attempt to ingest the range of ledgers specified in this
// session.
func (is *Session) Run() {
	is.Err = is.Ingestion.Start()
	if is.Err != nil {
		return
	}

	defer is.Ingestion.Rollback()

	for {
		isNextLegerLoaded, err := is.Cursor.NextLedger()
		if err != nil {
			is.Err = errors.Wrap(err, "failed to load next ledger")
			return
		}

		if !isNextLegerLoaded {
			break
		}

		if is.Err != nil {
			return
		}

		is.ingestLedger()
		is.flush()
	}

	if is.Err != nil {
		is.Ingestion.Rollback()
		return
	}

	is.Err = is.Ingestion.Close()
	if is.Err != nil {
		return
	}

	is.Err = is.CoreConnector.SetCursor("HORIZON", is.Cursor.LastLedger)
}

func (is *Session) flush() {
	if is.Err != nil {
		return
	}
	is.Err = is.Ingestion.Flush()
}

// ingestLedger ingests the current ledger
func (is *Session) ingestLedger() {
	if is.Err != nil {
		return
	}

	start := time.Now()
	is.Err = is.Ingestion.Ledger(
		is.Cursor.LedgerID(),
		is.Cursor.Ledger(),
		is.Cursor.SuccessfulTransactionCount(),
		is.Cursor.SuccessfulLedgerOperationCount(),
	)

	if is.Err != nil {
		return
	}

	// ingest accounts from genesis block
	if is.Cursor.LedgerSequence() == 1 {
		systemAccounts := []string{
			is.CoreInfo.MasterAccountID,
			is.CoreInfo.CommissionAccountID,
			is.CoreInfo.OperationalAccountID,
		}
		for _, address := range systemAccounts {
			_, is.Err = is.Ingestion.TryIngestAccount(address)
			if is.Err != nil {
				return
			}
		}
	}

	for is.Cursor.NextTx() {
		is.ingestTransaction()
	}

	is.Ingested++
	if is.Metrics != nil {
		is.Metrics.IngestLedgerTimer.Update(time.Since(start))
	}

	return
}

func (is *Session) storeTrades(orderBookID uint64, result xdr.ManageOfferSuccessResult) {
	if is.Err != nil {
		return
	}

	is.Err = is.Ingestion.StoreTrades(orderBookID, result, is.Cursor.Ledger().CloseTime)
}

func (is *Session) processManageInvoice(op xdr.ManageInvoiceOp, result xdr.ManageInvoiceResult) {
	if is.Err != nil {
		return
	}
	if op.InvoiceId == 0 || op.Amount != 0 {
		return
	}
	is.Ingestion.UpdateInvoice(op.InvoiceId, history.OperationStateCanceled, nil)

}

func (is *Session) permanentReject(op xdr.ReviewRequestOp) error {
	err := is.Cursor.HistoryQ().ReviewableRequests().PermanentReject(uint64(op.RequestId), string(op.Reason))
	if err != nil {
		return errors.Wrap(err, "failed to permanently reject request")
	}

	err = is.Ingestion.UpdatePayment(op.RequestId, false, nil)
	if err != nil {
		return errors.Wrap(err, "failed to permanently reject operation")
	}

	return nil
}

func (is *Session) handleCheckSaleState(result xdr.CheckSaleStateSuccess) {
	if is.Err != nil {
		return
	}

	state := history.SaleStateClosed
	if result.Effect.Effect == xdr.CheckSaleStateEffectCanceled {
		state = history.SaleStateCanceled
	}

	err := is.Cursor.HistoryQ().Sales().SetState(uint64(result.SaleId), state)
	if err != nil {
		is.Err = errors.Wrap(err, "failed to set state", logan.F{"sale_id": uint64(result.SaleId)})
		return
	}

}

func (is *Session) processManageAsset(op *xdr.ManageAssetOp) {
	if is.Err != nil {
		return
	}

	if op.Request.Action != xdr.ManageAssetActionCancelAssetRequest {
		return
	}

	err := is.Cursor.HistoryQ().ReviewableRequests().Cancel(uint64(op.RequestId))
	if err != nil {
		is.Err = errors.Wrap(err, "failed to cancel reviewable request", map[string]interface{}{
			"request_id": uint64(op.RequestId),
		})
		return
	}
}

func (is *Session) ingestOperationParticipants() {
	if is.Err != nil {
		return
	}

	// Find the participants
	var opParticipants []participants.Participant
	opParticipants, is.Err = participants.ForOperation(
		is.Ingestion.DB,
		&is.Cursor.Transaction().Envelope.Tx,
		is.Cursor.Operation(),
		*is.Cursor.OperationResult(),
		is.Cursor.Ledger(),
	)

	if is.Err != nil {
		return
	}

	if len(opParticipants) == 0 {
		return
	}

	is.Err = is.Ingestion.OperationParticipants(is.Cursor.OperationID(), opParticipants)
	if is.Err != nil {
		return
	}
}

func (is *Session) ingestTransaction() {
	if is.Err != nil {
		return
	}

	// skip ingesting failed transactions
	if !is.Cursor.Transaction().IsSuccessful() {
		return
	}

	is.Ingestion.Transaction(
		is.Cursor.Ledger(),
		is.Cursor.TransactionID(),
		is.Cursor.Transaction(),
		is.Cursor.TransactionFee(),
	)

	for is.Cursor.NextOp() {
		is.operation()
	}

	is.ingestTransactionParticipants()
}

func (is *Session) ingestTransactionParticipants() {
	if is.Err != nil {
		return
	}

	// Find the participants

	var p []xdr.AccountId
	p, is.Err = participants.ForTransaction(
		is.Ingestion.DB,
		&is.Cursor.Transaction().Envelope.Tx,
		*is.Cursor.Transaction().Result.Result.Result.Results,
		&is.Cursor.Transaction().ResultMeta,
		&is.Cursor.TransactionFee().Changes,
		is.Cursor.Ledger(),
	)
	if is.Err != nil {
		return
	}

	is.Err = is.Ingestion.TransactionParticipants(is.Cursor.TransactionID(), p)
	if is.Err != nil {
		return
	}

}

func (is *Session) processPayment(paymentOp xdr.PaymentOp, source xdr.AccountId, result xdr.PaymentResponse) {
	if is.Err != nil {
		return
	}

	invoiceReference := paymentOp.InvoiceReference
	if invoiceReference != nil {
		if invoiceReference.Accept {
			is.Ingestion.UpdateInvoice(invoiceReference.InvoiceId, history.OperationStateSuccess, nil)
		} else if !invoiceReference.Accept {
			is.Ingestion.UpdateInvoice(invoiceReference.InvoiceId, history.OperationStateRejected, nil)
		}
	}
}

func (is *Session) updateIngestedPaymentRequest(operation xdr.Operation, source xdr.AccountId) {
	if is.Err != nil {
		return
	}
	reviewPaymentOp := operation.Body.MustReviewPaymentRequestOp()
	is.Err = is.Ingestion.UpdatePaymentRequest(
		is.Cursor.Ledger(),
		uint64(reviewPaymentOp.PaymentId),
		reviewPaymentOp.Accept,
	)
	if is.Err != nil {
		return
	}
}

func (is *Session) updateIngestedPayment(operation xdr.Operation, source xdr.AccountId, result xdr.OperationResultTr) {
	if is.Err != nil {
		return
	}
	reviewPaymentOp := operation.Body.MustReviewPaymentRequestOp()
	reviewPaymentResponse := result.MustReviewPaymentRequestResult().ReviewPaymentResponse

	if reviewPaymentResponse.RelatedInvoiceId != nil {
		if reviewPaymentOp.Accept {
			is.Ingestion.UpdateInvoice(*reviewPaymentResponse.RelatedInvoiceId,
				history.OperationStateSuccess, nil)
		} else {
			is.Ingestion.UpdateInvoice(*reviewPaymentResponse.RelatedInvoiceId,
				history.OperationStateFailed, reviewPaymentOp.RejectReason)
		}
	}

	state := reviewPaymentResponse.State
	if state == xdr.PaymentStatePending {
		return
	}
	is.Err = is.Ingestion.UpdatePayment(
		reviewPaymentOp.PaymentId,
		state == xdr.PaymentStateProcessed,
		reviewPaymentOp.RejectReason,
	)
	if is.Err != nil {
		return
	}
}
