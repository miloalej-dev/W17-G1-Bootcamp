package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type CarrierService interface {
	Retrieve(id int) (models.Carrier, error)
	Register(seller models.Carrier) (models.Carrier, error)
}
