package models

// Wharehouseis a struct that represents a wharehouse in JSON format
type Warehouse struct {
	Id					int		`json:"id"`
	Code              	string	`json:"code"`
	Address            	string 	`json:"address"`
	Telephone          	string 	`json:"telephone"`
	MinimumCapacity    	int    	`json:"minimum_capacity"`
	MinimumTemperature 	int    	`json:"minimum_temperature"`
}

func NewWarehouse(id int, code, address, telephone string, minimumCapacity, minimumTemperature int) *Warehouse {
	return &Warehouse{
		Id: id,
		Code: code,
		Address: address,
		Telephone: telephone,
		MinimumCapacity: minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}
}
