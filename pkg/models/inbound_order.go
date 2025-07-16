package models

type InboundOrder struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}

func NewInboundOrder(id int, orderNumber string, employeeId int, productBathId int, warehouseId int) *InboundOrder {
	return &InboundOrder{
		Id:             id,
		OrderNumber:    orderNumber,
		EmployeeId:     employeeId,
		ProductBatchId: productBathId,
		WarehouseId:    warehouseId,
	}
}
