package database

import (
	"fmt"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type BuyerRepository struct {
	db *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}

func (s *BuyerRepository) FindAll() ([]models.Buyer, error) {
	var buyers []models.Buyer

	result := s.db.Find(&buyers)

	if result.Error != nil {
		return nil, result.Error
	}

	return buyers, nil
}

func (s *BuyerRepository) FindById(id int) (models.Buyer, error) {
	var buyer models.Buyer

	result := s.db.First(&buyer, id)

	if result.Error != nil {
		return models.Buyer{}, result.Error
	}

	return buyer, nil
}

func (s *BuyerRepository) Create(buyer models.Buyer) (models.Buyer, error) {
	result := s.db.Create(&buyer)

	if result.Error != nil {
		return models.Buyer{}, result.Error
	}

	return buyer, nil
}

func (s *BuyerRepository) Update(buyer models.Buyer) (models.Buyer, error) {
	result := s.db.Save(&buyer)

	if result.Error != nil {
		return models.Buyer{}, result.Error
	}

	return buyer, nil
}

func (s *BuyerRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Buyer, error) {
	var buyer models.Buyer

	result := s.db.First(&buyer, id)
	if result.Error != nil {
		return models.Buyer{}, result.Error
	}

	fmt.Println(fields)
	result2 := s.db.Model(&buyer).Updates(fields)
	if result2.Error != nil {
		return models.Buyer{}, result.Error
	}

	return buyer, nil
}

func (s *BuyerRepository) Delete(id int) error {
	result := s.db.Delete(&models.Buyer{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
