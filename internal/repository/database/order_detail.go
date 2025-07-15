package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) *OrderDetailRepository {
	return &OrderDetailRepository{db: db}
}

// FindAll retrieves all order details
func (r *OrderDetailRepository) FindAll() ([]models.OrderDetail, error) {
	orderDetails := make([]models.OrderDetail, 0)
	result := r.db.Find(&orderDetails)
	if result.Error != nil {
		return []models.OrderDetail{}, result.Error
	}
	return orderDetails, nil
}

// FindById retrieves a specific order detail by ID
func (r *OrderDetailRepository) FindById(id int) (models.OrderDetail, error) {
	var od models.OrderDetail
	result := r.db.First(&od, id)
	if result.Error != nil {
		return models.OrderDetail{}, result.Error
	}
	return od, nil
}

// Create inserts a new order detail
func (r *OrderDetailRepository) Create(od models.OrderDetail) (models.OrderDetail, error) {
	result := r.db.Create(&od)
	if result.Error != nil {
		return models.OrderDetail{}, result.Error
	}
	return od, nil
}

// Update updates an entire order detail
func (r *OrderDetailRepository) Update(od models.OrderDetail) (models.OrderDetail, error) {
	result := r.db.Save(&od)
	if result.Error != nil {
		return models.OrderDetail{}, result.Error
	}
	return od, nil
}

// PartialUpdate modifies only specific fields
func (r *OrderDetailRepository) PartialUpdate(id int, fields map[string]interface{}) (models.OrderDetail, error) {
	var od models.OrderDetail
	result := r.db.First(&od, id)
	if result.Error != nil {
		return models.OrderDetail{}, result.Error
	}

	if val, ok := fields["quantity"]; ok {
		q := int(val.(float64))
		od.Quantity = &q
	}
	if val, ok := fields["clean_lines_status"]; ok {
		status := val.(string)
		od.CleanLinesStatus = &status
	}
	if val, ok := fields["temperature"]; ok {
		t := val.(float64)
		od.Temperature = &t
	}
	if val, ok := fields["product_record_id"]; ok {
		od.ProductRecordID = int(val.(float64))
	}
	if val, ok := fields["purchase_order_id"]; ok {
		od.PurchaseOrderID = int(val.(float64))
	}

	result = r.db.Save(&od)
	if result.Error != nil {
		return models.OrderDetail{}, result.Error
	}
	return od, nil
}

// Delete removes an order detail by ID
func (r *OrderDetailRepository) Delete(id int) error {
	result := r.db.Delete(&models.OrderDetail{}, id)
	return result.Error
}
