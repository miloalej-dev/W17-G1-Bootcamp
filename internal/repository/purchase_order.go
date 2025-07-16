package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type PurchaseOrderRepository interface {
	Repository[int, models.PurchaseOrder]
	FindByBuyerId(id int) ([]models.PurchaseOrder, error)
}
