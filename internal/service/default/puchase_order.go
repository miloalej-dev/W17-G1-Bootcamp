package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type PurchaseOrderDefault struct {
	// rp is the repository that will be used by the service
	rp repository.PurchaseOrderRepository
}

func NewPurchaseOrderDefault(rp repository.PurchaseOrderRepository) *PurchaseOrderDefault {
	return &PurchaseOrderDefault{rp: rp}
}

func (s *PurchaseOrderDefault) RetrieveAll() (v []models.PurchaseOrder, err error) {
	return s.rp.FindAll()
}

func (s *PurchaseOrderDefault) Retrieve(id int) (models.PurchaseOrder, error) {
	return s.rp.FindById(id)
}

func (s *PurchaseOrderDefault) Register(ss models.PurchaseOrder) (models.PurchaseOrder, error) {

	return s.rp.Create(ss)
}
func (s *PurchaseOrderDefault) Modify(ss models.PurchaseOrder) (models.PurchaseOrder, error) {
	return s.rp.Update(ss)
}

func (s *PurchaseOrderDefault) PartialModify(id int, fields map[string]any) (models.PurchaseOrder, error) {
	return s.rp.PartialUpdate(id, fields)
}

func (s *PurchaseOrderDefault) Remove(id int) error {
	return s.rp.Delete(id)
}

func (s *PurchaseOrderDefault) RetrieveByBuyer(id int) ([]models.PurchaseOrder, error) {
	return s.rp.FindByBuyerId(id)
}
