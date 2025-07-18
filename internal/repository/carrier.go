package repository

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type CarrierRepository interface {
	Repository[int, models.Carrier]
}
