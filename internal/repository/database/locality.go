package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type LocalityRepository struct {
	db *gorm.DB
}

func NewLocalityRepository(db *gorm.DB) *LocalityRepository {
	return &LocalityRepository{db: db}
}

func (l LocalityRepository) FindBySellerId(id int) (models.LocalitySellerCount, error) {
	var locality models.LocalitySellerCount

	err := l.db.Table("localities l").
		Select("l.id, l.locality, COUNT(s.id) as seller_count").
		Joins("LEFT JOIN sellers s ON s.locality_id = l.id").
		Where("l.id = ?", id).Group("l.id").Scan(&locality).Error

	return locality, err
}

func (l LocalityRepository) FindAll() ([]models.Locality, error) {
	var localities []models.Locality
	result := l.db.Find(&localities)
	if result.Error != nil {
		return nil, result.Error
	}
	return localities, nil
}

func (l LocalityRepository) FindById(id int) (models.Locality, error) {
	var locality models.Locality

	err := l.db.Table("localities l").
		Select("l.id, l.locality, COUNT(s.id) as seller_count").
		Joins("LEFT JOIN sellers s ON s.locality_id = l.id").
		Where("l.id = ?", id).Group("l.id").Scan(&locality).Error

	return locality, err
}

func (l LocalityRepository) Create(locality models.Locality) (models.Locality, error) {
	var exists models.Locality
	result := l.db.First(&exists, locality.Id)
	if result.RowsAffected > 0 {
		return models.Locality{}, repository.ErrEntityAlreadyExists
	}
	localityCreated := l.db.Create(&locality)
	if localityCreated.Error != nil {
		return models.Locality{}, localityCreated.Error
	}
	return locality, nil
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
