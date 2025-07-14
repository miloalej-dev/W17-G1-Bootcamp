package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type LocalityRepository struct {
	db *gorm.DB
}

func NewLocalityRepository(db *gorm.DB) *LocalityRepository {
	return &LocalityRepository{db: db}
}

func (l LocalityRepository) FindAll() ([]models.Locality, error) {
	var localities []models.Locality
	_ = l.db.Find(&localities)
	return localities, nil
}

func (l LocalityRepository) FindById(id int) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalityRepository) Create(entity models.Locality) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalityRepository) Update(entity models.Locality) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalityRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Locality, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalityRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
