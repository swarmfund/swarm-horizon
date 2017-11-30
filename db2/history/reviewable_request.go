package history

import (
	"gitlab.com/swarmfund/go/xdr"
)

// Represents Reviewable request
type ReviewableRequest struct {
	ID           uint64                    `db:"id"`
	Requestor    string                    `db:"requestor"`
	Reviewer     string                    `db:"reviewer"`
	Reference    *string                   `db:"reference"`
	RejectReason string                    `db:"reject_reason"`
	RequestType  xdr.ReviewableRequestType `db:"request_type"`
	RequestState ReviewableRequestState    `db:"request_state"`
	Hash         string                    `db:"hash"`
	Details      []byte                    `db:"details"`
}
