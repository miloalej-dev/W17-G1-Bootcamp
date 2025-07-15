package repository

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type CarrierRepository interface {
	Repository[int, models.Carrier]
	FindByCid(cid string) (models.Carrier, bool, error)
	FindByLocality(id int) ([]models.Carrier, error)
}
