package models

// Wharehouseis a struct that represents a wharehouse in JSON format
type Warehouse struct {
	Id					int		`json:"id"`
	WarehouseCode      	string	`json:"warehouse_code"`
	Address            	string 	`json:"address"`
	Telephone          	string 	`json:"telephone"`
	MinimumCapacity    	int    	`json:"minimum_capacity"`
	MinimumTemperature 	int    	`json:"minimum_temperature"`
	LocalityId			int		`json:"locality_id"`
}

func NewWarehouse(id int, code, address, telephone string, minimumCapacity, minimumTemperature, locality_id int) *Warehouse {
	return &Warehouse{
		Id: id,
		WarehouseCode: code,
		Address: address,
		Telephone: telephone,
		MinimumCapacity: minimumCapacity,
		MinimumTemperature: minimumTemperature,
		LocalityId: locality_id,
	}
}
