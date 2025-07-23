package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type LocalityService interface {
	RetrieveAll() ([]models.Locality, error)
	Retrieve(id int) (models.Locality, error)
	RetrieveCarriers(id int) ([]map[int]int, error)
	Register(seller models.Locality) (models.Locality, error)
	Modify(seller models.Locality) (models.Locality, error)
	PartialModify(id int, fields map[string]any) (models.Locality, error)
	Remove(id int) error
	RetrieveLocalityBySeller(id int) (models.LocalitySellerCount, error)
	RetrieveAllLocalitiesBySeller() ([]models.LocalitySellerCount, error)
}
