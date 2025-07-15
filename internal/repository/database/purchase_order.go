package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"gorm.io/gorm"
	"time"
)

type PurchaseOrderRepository struct {
	db *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}

// FindAll retrieves all purchase orders
func (r *PurchaseOrderRepository) FindAll() ([]models.PurchaseOrder, error) {
	purchaseOrders := make([]models.PurchaseOrder, 0)
	result := r.db.Find(&purchaseOrders)
	if result.Error != nil {
		return []models.PurchaseOrder{}, result.Error
	}
	return purchaseOrders, nil
}

// FindById retrieves a purchase order by its ID
func (r *PurchaseOrderRepository) FindById(id int) (models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder
	result := r.db.First(&purchaseOrder, id)
	if result.Error != nil {
		return models.PurchaseOrder{}, result.Error
	}
	return purchaseOrder, nil
}

// Create inserts a new purchase order
func (r *PurchaseOrderRepository) Create(po models.PurchaseOrder) (models.PurchaseOrder, error) {
	result := r.db.Create(&po)
	if result.Error != nil {
		return models.PurchaseOrder{}, result.Error
	}
	return po, nil
}

// Update replaces an existing purchase order by its struct
func (r *PurchaseOrderRepository) Update(po models.PurchaseOrder) (models.PurchaseOrder, error) {
	result := r.db.Save(&po)
	if result.Error != nil {
		return models.PurchaseOrder{}, result.Error
	}
	return po, nil
}

// PartialUpdate updates only the provided fields
func (r *PurchaseOrderRepository) PartialUpdate(id int, fields map[string]interface{}) (models.PurchaseOrder, error) {
	var po models.PurchaseOrder
	result := r.db.First(&po, id)
	if result.Error != nil {
		return models.PurchaseOrder{}, result.Error
	}

	// Apply each field conditionally
	if val, ok := fields["order_number"]; ok {
		po.OrderNumber = val.(string)
	}
	if val, ok := fields["order_date"]; ok {
		po.OrderDate = val.(time.Time)
	}
	if val, ok := fields["tracing_code"]; ok {
		po.TracingCode = val.(string)
	}
	if val, ok := fields["buyer_id"]; ok {
		po.BuyerID = int(val.(float64))
	}
	if val, ok := fields["warehouse_id"]; ok {
		po.WarehouseID = int(val.(float64))
	}
	if val, ok := fields["carrier_id"]; ok {
		po.CarrierID = int(val.(float64))
	}
	if val, ok := fields["order_status_id"]; ok {
		po.OrderStatusID = int(val.(float64))
	}

	result = r.db.Save(&po)
	if result.Error != nil {
		return models.PurchaseOrder{}, result.Error
	}
	return po, nil
}

// Delete removes a purchase order by ID
func (r *PurchaseOrderRepository) Delete(id int) error {
	result := r.db.Delete(&models.PurchaseOrder{}, id)
	return result.Error
}

func (r *PurchaseOrderRepository) FindByBuyerId(id int) ([]models.PurchaseOrder, error) {
	var orders []models.PurchaseOrder
	result := r.db.Where("buyer_id = ?", id).Find(&orders)
	if result.Error != nil {
		return []models.PurchaseOrder{}, result.Error
	}
	return orders, nil
}
