package request

import (
	"errors"
	"net/http"
	"time"
)

type ProductRecordRequest struct {
	Id             *int       `json:"id"`
	LastUpdateDate *time.Time `json:"last_update_date"`
	PurchasePrice  *float64   `json:"purchase_price"`
	SalePrice      *float64   `json:"sale_price"`
	ProductId      *int       `json:"product_id"`
}

func (b *ProductRecordRequest) Bind(r *http.Request) error {
	if b.LastUpdateDate == nil {
		return errors.New("Last update date must be not null")
	}
	if b.PurchasePrice == nil {
		return errors.New("Purchase price  must be not null")
	}
	if b.SalePrice == nil {
		return errors.New("Sale price must be not null")
	}

	return nil
}
