package database

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Seller{}, repository.ErrEntityNotFound
	}

	return seller, nil
}

func (s *SellerRepository) Create(seller models.Seller) (models.Seller, error) {
	result := s.db.Create(&seller)

	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.Seller{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) Update(seller models.Seller) (models.Seller, error) {
	result := s.db.Save(&seller)

	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.Seller{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Seller, error) {
	var seller models.Seller

	// First, find the seller to update
	result := s.db.First(&seller, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Seller{}, repository.ErrEntityNotFound
	}

	// Update only the specified fields
	result = s.db.Model(&seller).Updates(fields)
	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
		return models.Seller{}, repository.ErrForeignKeyViolation
	case result.Error != nil:
		return models.Seller{}, result.Error
	}

	return seller, nil
}

func (s *SellerRepository) Delete(id int) error {
	result := s.db.Delete(&models.Seller{}, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return repository.ErrEntityNotFound
	}

	return nil
}
