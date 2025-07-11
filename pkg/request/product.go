package request

import (
	"errors"
	"net/http"
)

type ProductRequest struct {
	ProductCode                    *string  `json:"product_code"`
	Description                    *string  `json:"description"`
	Width                          *float64 `json:"width"`
	Height                         *float64 `json:"height"`
	Length                         *float64 `json:"length"`
	NetWeight                      *float64 `json:"net_weight"`
	ExpirationRate                 *float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature *float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   *float64 `json:"freezing_rate"`
	ProductTypeId                  *int     `json:"product_type_id"`
	SellerId                       *int     `json:"seller_id,omitempty"`
}

func (p *ProductRequest) Bind(r *http.Request) error {
	if p.ProductCode == nil {
		return errors.New("productCode must not be null")
	}
	if p.Description == nil {
		return errors.New("description must not be null")
	}
	if p.Width == nil {
		return errors.New("width must not be null")
	}
	if p.Height == nil {
		return errors.New("height must not be null")
	}
	if p.Length == nil {
		return errors.New("length must not be null")
	}
	if p.NetWeight == nil {
		return errors.New("netWeight must not be null")
	}
	if p.ExpirationRate == nil {
		return errors.New("expirationRate must not be null")
	}
	if p.RecommendedFreezingTemperature == nil {
		return errors.New("recommendedFreezingTemperature must not be null")
	}
	if p.FreezingRate == nil {
		return errors.New("freezingRate must not be null")
	}
	if p.ProductTypeId == nil {
		return errors.New("productTypeId must not be null")
	}

	return nil
}
