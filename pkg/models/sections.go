package models

type Section struct {
	Id                 int   `json:"id"`
	SectionNumber      int   `json:"section_number"`
	CurrentTemperature int   `json:"current_temperature"`
	MinimumTemperature int   `json:"minimum_temperature"`
	CurrentCapacity    int   `json:"current_capacity"`
	MinimumCapacity    int   `json:"minimum_capacity"`
	MaximumCapacity    int   `json:"maximum_capacity"`
	WarehouseId        int   `json:"warehouse_id"`
	ProductTypeId      int   `json:"product_type_id"`
	ProductsBatch      []int `json:"products_batch"`
}
