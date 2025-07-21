package models

import "time"

type InboundOrder struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	OrderNumber    string    `json:"order_number"`
	OrderDate      time.Time `json:"order_date"`
	EmployeeId     int       `json:"employee_id"`
	ProductBatchId int       `json:"product_batch_id"`
	WarehouseId    int       `json:"warehouse_id"`
}

func NewInboundOrder(id int, orderNumber string, orderDate time.Time, employeeId int, productBathId int, warehouseId int) *InboundOrder {
	return &InboundOrder{
		Id:             id,
		OrderNumber:    orderNumber,
		OrderDate:      orderDate,
		EmployeeId:     employeeId,
		ProductBatchId: productBathId,
		WarehouseId:    warehouseId,
	}
}
