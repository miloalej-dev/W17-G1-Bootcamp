package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// BuyerService is an interface that represents a Buyer Service
type BuyerService interface {
	Create(buyer models.Buyer) (models.Buyer, error)
	FindAll() ([]models.Buyer, error)
	FindById(id int) (models.Buyer, error)
	Update(buyer models.Buyer) (models.Buyer, error)
	Delete(id int) error
}
