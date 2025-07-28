package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type ProductRecordRepository struct {
	db *gorm.DB
}

func NewProductRecordRepository(db *gorm.DB) *ProductRecordRepository {
	return &ProductRecordRepository{
		db: db,
	}
}

func (s *ProductRecordRepository) FindAll() ([]models.ProductRecord, error) {
	var productRecords []models.ProductRecord

	result := s.db.Find(&productRecords)

	if result.Error != nil {
		return nil, result.Error
	}

	return productRecords, nil
}

func (s *ProductRecordRepository) FindById(id int) (models.ProductRecord, error) {
	var productRecord models.ProductRecord

	result := s.db.First(&productRecord, id)

	if result.Error != nil {
		return models.ProductRecord{}, result.Error
	}

	return productRecord, nil
}

func (s *ProductRecordRepository) Create(productRecord models.ProductRecord) (models.ProductRecord, error) {
	result := s.db.Create(&productRecord)

	if result.Error != nil {
		return models.ProductRecord{}, result.Error
	}

	return productRecord, nil
}

func (s *ProductRecordRepository) Update(productRecord models.ProductRecord) (models.ProductRecord, error) {
	result := s.db.Save(&productRecord)

	if result.Error != nil {
		return models.ProductRecord{}, result.Error
	}

	return productRecord, nil
}

func (s *ProductRecordRepository) PartialUpdate(id int, fields map[string]interface{}) (models.ProductRecord, error) {
	var productRecord models.ProductRecord

	result := s.db.First(&productRecord, id)
	if result.Error != nil {
		return models.ProductRecord{}, result.Error
	}

	result = s.db.Model(&productRecord).Updates(fields)
	if result.Error != nil {
		return models.ProductRecord{}, result.Error
	}

	return productRecord, nil
}

func (s *ProductRecordRepository) Delete(id int) error {
	result := s.db.Delete(&models.ProductRecord{}, id)

	if result.RowsAffected < 1 {
		return repository.ErrEntityNotFound
	}

	return nil
}
