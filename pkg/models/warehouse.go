package models

// Wharehouse is a struct that represents a wharehouse
type Warehouse struct {
	ID   				int
	Code 				string
	Address 			string
	Telephone 			string
	MinimunCapacity 	int
	MinimumTemperature	int
}

// WharehouseDoc is a struct that represents a wharehouse in JSON format
type FooDoc struct {
	ID   				int		`json:"id"`
	Code 				string	`json:"code"`
	Address 			string	`json:"address"`
	Telephone 			string	`json:"telephone"`
	MinimunCapacity 	int		`json:"minimun_capacity"`
	MinimumTemperature	int		`json:"minimum_temperature"`
}
