package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

func (s *SellerRepository) FindAll() ([]models.Seller, error) {
	var sellers []models.Seller

	result := s.db.Find(&sellers)

	if result.Error != nil {
		return nil, result.Error
	}

	return sellers, nil
}

func (s *SellerRepository) FindById(id int) (models.Seller, error) {
	var seller models.Seller

	result := s.db.First(&seller, id)

	if result.Error != nil {
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) Create(seller models.Seller) (models.Seller, error) {
	result := s.db.Create(&seller)

	if result.Error != nil {
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) Update(seller models.Seller) (models.Seller, error) {
	result := s.db.Save(&seller)

	if result.Error != nil {
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Seller, error) {
	var seller models.Seller

	// First, find the seller to update
	result := s.db.First(&seller, id)
	if result.Error != nil {
		return models.Seller{}, result.Error
	}

	// Update only the specified fields
	result = s.db.Model(&seller).Updates(fields)
	if result.Error != nil {
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) Delete(id int) error {
	result := s.db.Delete(&models.Seller{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
