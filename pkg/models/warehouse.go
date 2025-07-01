package models

// Wharehouse is a struct that represents a wharehouse
type Warehouse struct {
	ID   				int
	WarehouseAttributes
}

// Represents the warehouse attributes
type WarehouseAttributes struct {
	Code 				string
	Address 			string
	Telephone 			string
	MinimunCapacity 	int
	MinimumTemperature	int
}

// WharehouseDoc is a struct that represents a wharehouse in JSON format
type WarehouseDoc struct {
	ID   				int		`json:"id"`
	Code 				string	`json:"code" validate:"required"`
	Address 			string	`json:"address" validate:"required"`
	Telephone 			string	`json:"telephone" validate:"required"`
	MinimunCapacity 	int		`json:"minimun_capacity" validate:"required"`
	MinimumTemperature	int		`json:"minimun_temperature" validate:"required"`
}
