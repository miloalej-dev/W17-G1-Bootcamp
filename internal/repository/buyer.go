package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// BuyerRepository is an interface that represents a Buyer repository
type BuyerRepository interface {
	Create(buyer models.BuyerAtributtes) (v *models.Buyer, err error)
	FindAll() (buyers map[int]models.Buyer, err error)
	FindById(id int) (v *models.Buyer, err error)
	Update(buyer models.Buyer) (*models.Buyer, error)
	Delete(id int) (*models.Buyer, error)
}
