package operations

type ManageSale struct {
	Base
	SaleID           uint64 `json:"sale_id"`
	ManageSaleAction int32  `json:"manage_sale_action"`
}
