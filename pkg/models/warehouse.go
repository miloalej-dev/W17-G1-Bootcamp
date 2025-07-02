package models

// Wharehouse is a struct that represents a wharehouse
type Warehouse struct {
	ID int
	WarehouseAttributes
}

// Represents the warehouse attributes
type WarehouseAttributes struct {
	Code               string
	Address            string
	Telephone          string
	MinimumCapacity    int
	MinimumTemperature int
}

// WharehouseDoc is a struct that represents a wharehouse in JSON format
type WarehouseDoc struct {
	ID					int		`json:"id"`
	Code              	string	`json:"code"`
	Address            	string 	`json:"address"`
	Telephone          	string 	`json:"telephone"`
	MinimumCapacity    	int    	`json:"minimum_capacity"`
	MinimumTemperature 	int    	`json:"minimum_temperature"`
}
