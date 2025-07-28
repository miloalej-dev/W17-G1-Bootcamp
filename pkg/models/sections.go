package models

type Section struct {
	Id                 int     `json:"id" gorm:"primaryKey"`
	SectionNumber      string  `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseId        int     `json:"warehouse_id"`
	ProductTypeId      int     `json:"product_type_id"`
}

type SectionReport struct {
	SectionId     int    `json:"section_id"`
	SectionNumber string `json:"section_number"`
	ProductsCount int    `json:"products_count"`
}
