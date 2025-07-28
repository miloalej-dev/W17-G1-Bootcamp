package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type CarrierService interface {
	RetrieveAll() ([]models.Carrier, error)
	Retrieve(id int) (models.Carrier, error)
	Register(carrier models.Carrier) (models.Carrier, error)
	Modify(carrier models.Carrier) (models.Carrier, error)
	PartialModify(id int, fields map[string]any) (models.Carrier, error)
	Remove(id int) error
}
