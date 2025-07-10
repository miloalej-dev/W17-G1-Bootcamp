package _default

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
func (s *BuyerDefault) RetrieveAll() ([]models.Buyer, error) {
	v, err := s.rp.FindAll()

	return v, err
}

func (s *BuyerDefault) Retrieve(id int) (models.Buyer, error) {
	v, err := s.rp.FindById(id)
	if err != nil {
		return models.Buyer{}, err
	}
	return v, nil

}

func (s *BuyerDefault) Register(buyer models.Buyer) (models.Buyer, error) {
	v, err := s.rp.Create(buyer)
	if err != nil {
		return models.Buyer{}, err
	}
	return v, nil
}

func (s *BuyerDefault) Remove(id int) error {
	err := s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil

}

func (s *BuyerDefault) Modify(buyer models.Buyer) (models.Buyer, error) {
	v, err := s.rp.Update(buyer)
	if err != nil {
		return models.Buyer{}, err
	}

	return v, nil
}

func (s *BuyerDefault) PartialModify(id int, fields map[string]any) (models.Buyer, error) {
	v, err := s.rp.PartialUpdate(id, fields)
	if err != nil {
		return models.Buyer{}, err
	}

	return v, nil
}
