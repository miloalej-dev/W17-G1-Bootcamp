package models

type Employee struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

type EmployeeInboundOrdersReport struct {
	Employee
	InboundOrdersCount int `json:"inbound_orders_count"`
}
