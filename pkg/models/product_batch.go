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
	MinumumTemperature float64 `json:"minumum_temperature"`
	SectionsId         int     `json:"sections_id"`
	ProductsId         int     `json:"products_id"`
}

// NewProductBatch is a function that creates a new Product
func NewProductBatch(id int, batchNumber int, currentQuantity int, currentTemperature float64, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minumumTemperature float64, sectionsId int, productsId int) ProductBatch {
	return ProductBatch{
		Id:                 id,
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinumumTemperature: minumumTemperature,
		SectionsId:         sectionsId,
		ProductsId:         productsId,
	}
}
