package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// CarrierDefault is a struct that represents the default service for warehouses
type CarrierDefault struct {
	// rp is the repository that will be used by the service
	rp repository.CarrierRepository
}

func (s *CarrierDefault) RetrieveAll() ([]models.Carrier, error) {
	return s.rp.FindAll()
}

func (s *CarrierDefault) Retrieve(id int) (models.Carrier, error) {
	return s.rp.FindById(id)
}

// NewCarrierDefault is a function that returns a new instance of CarrierDefault
func NewCarrierDefault(rp repository.CarrierRepository) *CarrierDefault {
	return &CarrierDefault{rp: rp}
}

func (s *CarrierDefault) Register(entity models.Carrier) (models.Carrier, error) {
	return s.rp.Create(entity)
}

func (s *CarrierDefault) Modify(entity models.Carrier) (models.Carrier, error) {
	return s.rp.Update(entity)
}

func (s *CarrierDefault) PartialModify(id int, fields map[string]interface{}) (models.Carrier, error) {
	return s.rp.PartialUpdate(id, fields)
}

func (s *CarrierDefault) Remove(id int) error {
	return s.rp.Delete(id)
}
