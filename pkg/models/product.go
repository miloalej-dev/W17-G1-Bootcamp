package models

type Product struct {
	//Identificador único de producto
	ID int `json:"id"`
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
	ProductTypeID int `json:"product_type_id,validate:required"`
	// Un Proveedor (Seller) asociado (NO es OBLIGATORIA)
	SellerID int `json:"seller_id,omitempty"`
}
type ProductType struct {
	ID          int
	Description string
}
