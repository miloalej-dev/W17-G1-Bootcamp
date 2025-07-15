package models

type OrderDetail struct {
	ID               int      `json:"id"`
	Quantity         *int     `json:"quantity"`
	CleanLinesStatus *string  `json:"clean_lines_status"`
	Temperature      *float64 `json:"temperature"`
	ProductRecordID  int      `json:"product_record_id"`
	PurchaseOrderID  int      `json:"purchase_order_id"`
}
