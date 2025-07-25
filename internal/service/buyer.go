package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// BuyerService is an interface that represents a Buyer Service
type BuyerService interface {
	Register(buyer models.Buyer) (models.Buyer, error)
	RetrieveAll() ([]models.Buyer, error)
	Retrieve(id int) (models.Buyer, error)
	Modify(buyer models.Buyer) (models.Buyer, error)
	PartialModify(id int, fields map[string]any) (models.Buyer, error)
	Remove(id int) error
	RetrieveByPurchaseOrderReport(id int) ([]models.BuyerReport, error)
}
