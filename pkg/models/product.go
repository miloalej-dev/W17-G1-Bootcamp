package models

type Product struct {
	//Identificador único de producto
	ID int `json:"id"`
	// Un código de producto
	ProductCode string `json:"product_code"`
	// Una descripción
	Description string `json:"description"`
	// Una descripción
	Width float64 `json:"width"`
	// Una dimensión de alto
	Height float64 `json:"height"`
	//Una dimensión de largo
	Length float64 `json:"lenght"`
	//Un peso neto
	NetWeight float64 `json:"net_weight"`
	// Un ratio de vencimiento
	ExpirationRate float64 `json:"expiration_rate"`
	// Una temperatura de congelación recomendada
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	// Un ratio de congelación, el cual indica que tan rápido se congelara
	FreezingRate float64 `json:"freezing_rate"`
	// Un Tipo de producto(ProductType) asociado
	ProductTypeID int `json:"product_type_id"`
	// Un Proveedor (Seller) asociado (NO es OBLIGATORIA)
	SellerID int `json:"seller_id,omitempty"`
}
type ProductType struct {
	ID          int
	Description string
}
