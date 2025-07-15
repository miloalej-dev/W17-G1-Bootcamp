package models

// Product represents the structure of a product.
type ProductBatch struct {
	Id                 int     `json:"id"`
	BatchNumber        int     `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  int     `json:"manufacturing_hour"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	SectionId          int     `json:"section_id"`
	ProductId          int     `json:"product_id"`
}

// NewProductBatch is a function that creates a new Product
func NewProductBatch(id int, batchNumber int, currentQuantity int, currentTemperature float64, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature float64, sectionId int, productId int) ProductBatch {
	return ProductBatch{
		Id:                 id,
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: minimumTemperature,
		SectionId:          sectionId,
		ProductId:          productId,
	}
}
