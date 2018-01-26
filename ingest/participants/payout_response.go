package participants

import "gitlab.com/swarmfund/go/xdr"

type PayoutResponse struct {
	Receiver        xdr.AccountId
	ReceiverBalance xdr.BalanceId
	ReceivedAmount  xdr.Uint64
}

func NewPayoutResponse(receiver xdr.AccountId, receiverBalance xdr.BalanceId, receivedAmount xdr.Uint64) *PayoutResponse {
	return &PayoutResponse{
		Receiver:        receiver,
		ReceiverBalance: receiverBalance,
		ReceivedAmount:  receivedAmount,
	}
}

func (p *PayoutResponse) ToParticipant() Participant {
	return Participant{p.Receiver, &p.ReceiverBalance, &map[string]interface{}{
		"received_amount": p.ReceivedAmount,
	}}
}
