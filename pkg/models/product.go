package models

// Product represents the structure of a product.
type Product struct {
	// Unique product identifier.
	Id int `json:"id"`
	// A unique code for the product.
	ProductCode string `json:"product_code"`
	// A description of the product.
	Description string `json:"description"`
	// The width dimension of the product.
	Width float64 `json:"width"`
	// The height dimension of the product.
	Height float64 `json:"height"`
	// The length dimension of the product.
	Length float64 `json:"length"`
	// The net weight of the product.
	NetWeight float64 `json:"net_weight"`
	// The expiration rate of the product.
	ExpirationRate float64 `json:"expiration_rate"`
	// The recommended freezing temperature for the product.
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	// The freezing rate, which indicates how fast the product will freeze.
	FreezingRate float64 `json:"freezing_rate"`
	// The ID of the associated ProductType.
	ProductTypeId int `json:"product_type_id"`
	// The ID of the associated Seller (OPTIONAL).
	SellerId *int `json:"seller_id,omitempty"`
}

// ProductType represents the type of product.
type ProductType struct {
	Id          int
	Description string
}

// NewProduct is a function that creates a new Product
func NewProduct(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate float64, recommendedFreezingTemperature float64, freezingRate float64, productTypeId int, sellerId *int) *Product {
	return &Product{
		Id:                             id,
		ProductCode:                    productCode,
		Description:                    description,
		Width:                          width,
		Height:                         height,
		Length:                         length,
		NetWeight:                      netWeight,
		ExpirationRate:                 expirationRate,
		RecommendedFreezingTemperature: recommendedFreezingTemperature,
		FreezingRate:                   freezingRate,
		ProductTypeId:                  productTypeId,
		SellerId:                       sellerId,
	}
}
