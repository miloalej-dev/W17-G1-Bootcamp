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

func (l LocalityRepository) FindLocalityBySeller(id int) (models.LocalitySellerCount, error) {
	var exists int64
	l.db.Model(&models.Locality{}).Where("id = ?", id).Count(&exists)
	if exists == 0 {
		return models.LocalitySellerCount{}, repository.ErrLocalityNotFound
	}
	var locality models.LocalitySellerCount
	query := l.db.Table("localities").
		Select("localities.id as id ,localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count").
		Joins("JOIN provinces p ON localities.province_id = p.id").
		Joins("JOIN countries c ON p.country_id = c.id").
		Joins("LEFT JOIN sellers s ON localities.id = s.locality_id").
		Where("localities.id = ?", id).
		Group("localities.id, localities.locality, p.province, c.country")

	err := query.Scan(&locality).Error
	if err != nil {
		return models.LocalitySellerCount{}, err
	}
	// Verifica si se encontraron resultados
	return locality, err
}

func (l LocalityRepository) FindAllLocality() ([]models.LocalitySellerCount, error) {
	var localitiesSellers []models.LocalitySellerCount
	query := l.db.Table("localities").
		Select("localities.id as id ,localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count").
		Joins("JOIN provinces p ON localities.province_id = p.id").
		Joins("JOIN countries c ON p.country_id = c.id").
		Joins("LEFT JOIN sellers s ON localities.id = s.locality_id").
		Group("localities.id, localities.locality, p.province, c.country")
	err := query.Scan(&localitiesSellers).Error
	if err != nil {
		return nil, err
	}
	if len(localitiesSellers) == 0 {
		return nil, repository.ErrEmptyEntity
	}

	return localitiesSellers, nil
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
	result := l.db.First(&locality, id)
	if result.Error != nil {
		return models.Locality{}, repository.ErrEntityNotFound
	}
	return locality, nil
}

func (l LocalityRepository) Create(locality models.LocalityDoc) (models.LocalityDoc, error) {
	var exists models.Locality
	result := l.db.First(&exists, locality.Id)
	if result.RowsAffected > 0 {
		return models.LocalityDoc{}, repository.ErrEntityAlreadyExists
	}
	var idProvinceTable int
	err := l.db.Table("provinces p").
		Select("p.id").
		Joins("INNER JOIN countries c ON c.id = p.country_id").
		Where("p.province = ? AND c.country = ?", locality.Province, locality.Country).Scan(&idProvinceTable).Error
	if err != nil || idProvinceTable == 0 {
		return models.LocalityDoc{}, repository.ErrProvinceNotFound
	}
	localityCreated := models.Locality{
		Id:         locality.Id,
		Locality:   locality.Locality,
		ProvinceId: idProvinceTable,
	}
	result = l.db.Create(&localityCreated)
	if result.Error != nil {
		return models.LocalityDoc{}, result.Error
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

func (l LocalityRepository) FindAllCarriers() ([]models.LocalityCarrierCount, error) {

	var carriers []models.LocalityCarrierCount
	err := l.db.Model(&models.Locality{}).
		Select("localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers'").
		Joins("LEFT JOIN carriers ON localities.id = carriers.locality_id").
		Group("localities.id").
		Find(&carriers).Error

	if err != nil {
		return make([]models.LocalityCarrierCount, 0), err
	}

	if len(carriers) == 0 {
		return make([]models.LocalityCarrierCount, 0), repository.ErrEntityNotFound
	}

	return carriers, nil
}

func (l LocalityRepository) FindCarriersByLocality(id int) ([]models.LocalityCarrierCount, error) {

	var carriers []models.LocalityCarrierCount
	err := l.db.Model(&models.Locality{}).
		Select("localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers'").
		Joins("LEFT JOIN carriers ON localities.id = carriers.locality_id").
		Where("localities.id = ?", id).
		Group("localities.id").
		Find(&carriers).Error

	if err != nil {
		return make([]models.LocalityCarrierCount, 0), err
	}

	if len(carriers) == 0 {
		return make([]models.LocalityCarrierCount, 0), repository.ErrEntityNotFound
	}

	return carriers, nil
}
