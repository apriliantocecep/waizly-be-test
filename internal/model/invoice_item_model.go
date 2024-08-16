package model

type InvoiceItemResponse struct {
	ID       uint    `json:"id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	ItemId   uint    `json:"item_id"`
}
