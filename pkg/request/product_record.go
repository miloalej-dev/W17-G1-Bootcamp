package request

import (
	"errors"
	"net/http"
)

type ProductRecordRequest struct {
	Id            *int     `json:"id"`
	LastUpdate    *string  `json:"last_update"`
	PurchasePrice *float64 `json:"purchase_price"`
	SalePrice     *float64 `json:"sale_price"`
	ProductId     *int     `json:"product_id"`
}

func (b *ProductRecordRequest) Bind(r *http.Request) error {

	if b.LastUpdate == nil || *b.LastUpdate == "" {
		return errors.New("last update date must be not null")
	}
	if b.PurchasePrice == nil || *b.PurchasePrice < 0 {
		return errors.New("purchase price  must be not null and greater than 0")
	}
	if b.SalePrice == nil || *b.SalePrice < 0 {
		return errors.New("sale price must be not null and greater than 0")
	}
	if b.ProductId == nil || *b.ProductId < 0 {

		return errors.New("products Id must be not null and greater than 0")
	}

	return nil
}
