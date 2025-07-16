package service

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type CarrierService interface {
	Register(seller models.Carrier) (models.Carrier, error)
}
