package models

import "time"

type ProductRecord struct {
	Id            int       `json:"id"`
	LastUpdate    time.Time `json:"last_update"`
	PurchasePrice float64   `json:"purchase_price"`
	SalePrice     float64   `json:"sale_price"`
	ProductsId    int       `json:"products_id"`
}

// NewProductRecord is a function that creates a new productRecord
func NewProductRecord(id int, lastUpdateDate time.Time, purchasePrice float64, salePrice float64, productId int) *ProductRecord {
	return &ProductRecord{
		Id:            id,
		LastUpdate:    lastUpdateDate,
		PurchasePrice: purchasePrice,
		SalePrice:     salePrice,
		ProductsId:    productId,
	}
}
