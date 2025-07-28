package database

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Buyer{}, repository.ErrEntityNotFound
	}

	result = s.db.Model(&buyer).Updates(fields)
	if result.Error != nil {
		return models.Buyer{}, result.Error
	}

	return buyer, nil
}

func (s *BuyerRepository) Delete(id int) error {
	result := s.db.Delete(&models.Buyer{}, id)

	if result.RowsAffected < 1 {
		return repository.ErrEntityNotFound
	}

	return nil
}
func (r *BuyerRepository) FindByPurchaseOrderReport(id int) ([]models.BuyerReport, error) {
	var reports []models.BuyerReport

	if id == 0 {
		// Obtener todos los buyers con su conteo de órdenes
		err := r.db.
			Table("buyers").
			Select("buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name, COUNT(purchase_orders.id) AS purchase_orders_count").
			Joins("LEFT JOIN purchase_orders ON purchase_orders.buyer_id = buyers.id").
			Group("buyers.id").
			Scan(&reports).Error

		if err != nil {
			return nil, err
		}
	} else {
		_, err := r.FindById(id)
		if err != nil {
			return reports, err
		}
		// Obtener un solo buyer con su conteo de órdenes
		var report models.BuyerReport
		err = r.db.
			Table("buyers").
			Select("buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name,  COUNT(purchase_orders.id) AS purchase_orders_count").
			Joins("LEFT JOIN purchase_orders ON purchase_orders.buyers_id = buyers.id").
			Where("buyers.id = ?", id).
			Group("buyers.id").
			Scan(&report).Error

		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}
