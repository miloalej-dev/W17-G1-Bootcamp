package models

type Product struct {
	//Identificador único de producto
	ID int
	// Un código de producto
	ProductCode string
	// Una descripción
	Description string
	// Una descripción
	Width float64
	// Una dimensión de alto
	Height float64
	//Una dimensión de largo
	Length float64
	//Un peso neto
	NetWeight float64
	// Un ratio de vencimiento
	ExpirationRate float64
	// Una temperatura de congelación recomendada
	RecommendedFreezingTemperature float64
	// Un ratio de congelación, el cual indica que tan rápido se congelara
	FreezingRate float64
	// Un Tipo de producto(ProductType) asociado
	ProductTypeID int
	// Un Proveedor (Seller) asociado (NO es OBLIGATORIA)
	SellerID int
}
type ProductType struct {
	ID          int
	Description string
}
