package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type OrderDetailDefault struct {
	// rp is the repository that will be used by the service
	rp repository.OrderDetailRepository
}

func NewOrderDetailDefault(rp repository.OrderDetailRepository) *OrderDetailDefault {
	return &OrderDetailDefault{rp: rp}
}

func (s *OrderDetailDefault) FindAll() (v []models.OrderDetail, err error) {
	return s.rp.FindAll()
}

func (s *OrderDetailDefault) FindByID(id int) (models.OrderDetail, error) {
	return s.rp.FindById(id)
}

func (s *OrderDetailDefault) Create(ss models.OrderDetail) (models.OrderDetail, error) {

	return s.rp.Create(ss)
}
func (s *OrderDetailDefault) Update(ss models.OrderDetail) (models.OrderDetail, error) {
	return s.rp.Update(ss)
}
func (s *OrderDetailDefault) Delete(id int) error {
	return s.rp.Delete(id)
}
