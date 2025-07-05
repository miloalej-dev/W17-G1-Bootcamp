package models

type Product struct {
	//Identificador único de producto
	Id int `json:"id"`
	// Un código de producto
	ProductCode string `json:"product_code,validate:required"`
	// Una descripción
	Description string `json:"description,validate:required"`
	// Una descripción
	Width float64 `json:"width,validate:required"`
	// Una dimensión de alto
	Height float64 `json:"height,validate:required"`
	//Una dimensión de largo
	Length float64 `json:"length,validate:required"`
	//Un peso neto
	NetWeight float64 `json:"net_weight,validate:required"`
	// Un ratio de vencimiento
	ExpirationRate float64 `json:"expiration_rate,validate:required"`
	// Una temperatura de congelación recomendada
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature,validate:required"`
	// Un ratio de congelación, el cual indica que tan rápido se congelara
	FreezingRate float64 `json:"freezing_rate,validate:required"`
	// Un Tipo de producto(ProductType) asociado
	ProductTypeId int `json:"product_type_id,validate:required"`
	// Un Proveedor (Seller) asociado (NO es OBLIGATORIA)
	SellerId int `json:"seller_id,omitempty"`
}
type ProductType struct {
	Id          int
	Description string
}

// NewProduct is a function that creates a new Product
func NewProduct(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate float64, recommendedFreezingTemperature float64, freezingRate float64, productTypeId int, sellerId int) *Product {
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
