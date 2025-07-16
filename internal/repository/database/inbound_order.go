package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type InboundOrderRepository struct {
	db *gorm.DB
}

func NewInboundOrderRepository(db *gorm.DB) *InboundOrderRepository {
	return &InboundOrderRepository{db: db}
}

func (i *InboundOrderRepository) FindAll() ([]models.InboundOrder, error) {
	var inboundOrders []models.InboundOrder

	result := i.db.Find(&inboundOrders)

	if result.Error != nil {
		return nil, result.Error
	}

	return inboundOrders, nil
}

func (i *InboundOrderRepository) FindById(id int) (models.InboundOrder, error) {
	var inboundOrder models.InboundOrder

	result := i.db.First(&inboundOrder, id)

	if result.Error != nil {
		return models.InboundOrder{}, result.Error
	}

	return inboundOrder, nil
}

func (i *InboundOrderRepository) Create(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	result := i.db.Create(&inboundOrder)

	if result.Error != nil {
		return models.InboundOrder{}, result.Error
	}

	return inboundOrder, nil
}

func (i *InboundOrderRepository) Update(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	result := i.db.Save(&inboundOrder)

	if result.Error != nil {
		return models.InboundOrder{}, result.Error
	}

	return inboundOrder, nil
}

func (i *InboundOrderRepository) PartialUpdate(id int, fields map[string]interface{}) (models.InboundOrder, error) {
	var inboundOrder models.InboundOrder

	// First, find the seller to update
	result := i.db.First(&inboundOrder, id)
	if result.Error != nil {
		return models.InboundOrder{}, result.Error
	}

	// Update only the specified fields
	result = i.db.Model(&inboundOrder).Updates(fields)
	if result.Error != nil {
		return models.InboundOrder{}, result.Error
	}

	return inboundOrder, nil
}

func (i *InboundOrderRepository) Delete(id int) error {
	result := i.db.Delete(&models.InboundOrder{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
