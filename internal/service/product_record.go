package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type ProductRecordService interface {
	RetrieveAll() ([]models.ProductRecord, error)
	Retrieve(id int) (models.ProductRecord, error)
	Register(productRecord models.ProductRecord) (models.ProductRecord, error)
	Modify(productRecord models.ProductRecord) (models.ProductRecord, error)
	PartialModify(id int, fields map[string]any) (models.ProductRecord, error)
	Remove(id int) error
}
