package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/database"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type LocalityService struct {
	rp database.LocalityRepository
}

func NewLocalityService(rp *database.LocalityRepository) *LocalityService {
	return &LocalityService{*rp}
}

func (l LocalityService) RetrieveAll() ([]models.Locality, error) {
	localities, err := l.rp.FindAll()
	return localities, err
}

func (l LocalityService) Retrieve(id int) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}
func (l LocalityService) RetrieveLocalityBySeller(id int) (models.LocalitySellerCount, error) {
	return l.rp.FindLocalityBySeller(id)
}
func (l LocalityService) Register(seller models.LocalityDoc) (models.LocalityDoc, error) {
	return l.rp.Create(seller)
}

func (l LocalityService) Modify(seller models.Locality) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalityService) PartialModify(id int, fields map[string]any) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalityService) Remove(id int) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalityService) RetrieveCarriers() ([]models.LocalityCarrierCount, error) {
	return l.rp.FindAllCarriers()
}

func (l LocalityService) RetrieveCarriersByLocality(id int) ([]models.LocalityCarrierCount, error) {
	return l.rp.FindCarriersByLocality(id)
}

func (l LocalityService) RetrieveAllLocalitiesBySeller() ([]models.LocalitySellerCount, error) {
	return l.rp.FindAllLocality()
}
