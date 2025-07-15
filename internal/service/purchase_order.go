package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type PurchaseOrderService interface {
	RetrieveAll() (v []models.PurchaseOrder, err error)
	Retrieve(id int) (models.PurchaseOrder, error)
	Register(s models.PurchaseOrder) (models.PurchaseOrder, error)
	Modify(s models.PurchaseOrder) (models.PurchaseOrder, error)
	PartialModify(id int, fields map[string]any) (models.PurchaseOrder, error)
	Remove(id int) error
}
