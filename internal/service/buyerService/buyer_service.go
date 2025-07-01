package buyerService

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// BuyerService is an interface that represents a Buyer Service
type BuyerService interface {
	//Create(buyer models.Buyer) (v *models.Buyer, err error)
	FindAll() (v map[int]models.Buyer, err error)
	//FindById(id int) (v *models.Buyer, err error)
	//Update(buyer models.Buyer) (*models.Buyer, error)
	//Delete(id int) (*models.Buyer, error)
}
