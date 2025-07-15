package models

import "time"

type ProductRecord struct {
	Id             int       `json:"id"`
	LastUpdateDate time.Time `json:"last_update_date"`
	PurchasePrice  float64   `json:"purchase_price"`
	SalePrice      float64   `json:"sale_price"`
	ProductId      int       `json:"product_id"`
}

// NewProductRecord is a function that creates a new productRecord
func NewProductRecord(id int, lastUpdateDate time.Time, purchasePrice float64, salePrice float64, productId int) *ProductRecord {
	return &ProductRecord{
		Id:             id,
		LastUpdateDate: lastUpdateDate,
		PurchasePrice:  purchasePrice,
		SalePrice:      salePrice,
		ProductId:      productId,
	}
}
