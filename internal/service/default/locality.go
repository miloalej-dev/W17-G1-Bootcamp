package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type LocalityService struct {
	rp repository.LocalityRepository
}

func NewLocalityService(rp repository.LocalityRepository) *LocalityService {
	return &LocalityService{rp}
}

func (l LocalityService) RetrieveAll() ([]models.Locality, error) {
	localities, err := l.rp.FindAll()
	return localities, err
}

func (l LocalityService) Retrieve(id int) (models.Locality, error) {
	return l.rp.FindById(id)
}
func (l LocalityService) RetrieveLocalityBySeller(id int) (models.LocalitySellerCount, error) {
	return l.rp.FindLocalityBySeller(id)
}
func (l LocalityService) Register(locality models.Locality) (models.Locality, error) {
	return l.rp.Create(locality)
}

func (l LocalityService) Modify(locality models.Locality) (models.Locality, error) {
	return l.rp.Update(locality)
}

func (l LocalityService) PartialModify(id int, fields map[string]any) (models.Locality, error) {
	return l.rp.PartialUpdate(id, fields)
}

func (l LocalityService) Remove(id int) error {
	return l.rp.Delete(id)
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

func (l LocalityService) RegisterWithNames(locality models.LocalityDoc) (models.LocalityDoc, error) {
	return l.rp.CreateWithNames(locality)
}
