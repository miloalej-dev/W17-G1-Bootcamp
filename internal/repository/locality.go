package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type LocalityRepository interface {
	Repository[int, models.Locality]
	FindLocalityBySeller(id int) (models.LocalitySellerCount, error)
	FindAllLocality() ([]models.LocalitySellerCount, error)
	FindAllCarriers() ([]models.LocalityCarrierCount, error)
	FindCarriersByLocality(id int) ([]models.LocalityCarrierCount, error)
}
