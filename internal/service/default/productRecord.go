package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

func NewProductRecordDefault(rp repository.ProductRecordRepository) *ProductRecordDefault {
	return &ProductRecordDefault{
		rp: rp,
	}
}

type ProductRecordDefault struct {
	// rp is the repository that will be used by the service
	rp repository.ProductRecordRepository
}

func (s ProductRecordDefault) RetrieveAll() ([]models.ProductRecord, error) {
	return s.rp.FindAll()
}

func (s ProductRecordDefault) Retrieve(id int) (models.ProductRecord, error) {
	return s.rp.FindById(id)
}

func (s ProductRecordDefault) Register(productRecord models.ProductRecord) (models.ProductRecord, error) {
	return s.rp.Create(productRecord)
}

func (s ProductRecordDefault) Modify(productRecord models.ProductRecord) (models.ProductRecord, error) {
	return s.rp.Update(productRecord)
}

func (s ProductRecordDefault) PartialModify(id int, fields map[string]any) (models.ProductRecord, error) {
	return s.rp.PartialUpdate(id, fields)
}

func (p ProductRecordDefault) Remove(id int) error {
	return p.rp.Delete(id)
}
