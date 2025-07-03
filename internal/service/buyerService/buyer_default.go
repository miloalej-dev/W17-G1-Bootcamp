package buyerService

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// NewBuyerDefault is a function that returns a new instance of BuyerDefault
func NewBuyerDefault(rp repository.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{rp: rp}
}

// BuyerDefault is a struct that represents the default service for buyers
type BuyerDefault struct {
	// rp is the repository that will be used by the service
	rp repository.BuyerRepository
}

// FindAll is a method that returns a map of all buyers
func (s *BuyerDefault) FindAll() (v map[int]models.Buyer, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *BuyerDefault) FindById(id int) (v *models.Buyer, err error) {
	v, err = s.rp.FindById(id)
	if err != nil {
		return nil, err
	}
	return v, nil

}

func (s *BuyerDefault) Create(buyer models.BuyerAtributtes) (v *models.Buyer, err error) {
	v, err = s.rp.Create(buyer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *BuyerDefault) Delete(id int) (v *models.Buyer, err error) {
	v, err = s.rp.Delete(id)
	if err != nil {
		return nil, err
	}
	return v, nil

}

func (s *BuyerDefault) Update(buyer models.Buyer) (v *models.Buyer, err error) {
	v, err = s.rp.Update(buyer)
	if err != nil {
		return nil, err
	}

	return v, nil
}
