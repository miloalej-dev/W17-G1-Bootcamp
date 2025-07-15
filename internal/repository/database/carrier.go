package database

import (
	"gorm.io/gorm"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// Carrier repository
type CarrierDB struct {
	db *gorm.DB
}

// Creates a new Carrier repository
func NewCarrierDB(db *gorm.DB) *CarrierDB {
	return &CarrierDB{db: db}
}

func (r *CarrierDB) FindAll() ([]models.Carrier, error) {
	carriers := make([]models.Carrier, 0)
	result := r.db.Find(&carriers)
	if result.Error != nil {
		return make([]models.Carrier, 0), result.Error
	}
	return carriers, nil
}

func (r *CarrierDB) FindById(id int) (models.Carrier, error) {
	var carrier models.Carrier
	result := r.db.First(&carrier, id)
	if result.Error != nil {
		return models.Carrier{}, result.Error
	}
	return carrier, nil
}

func (r *CarrierDB) Create(carrier models.Carrier) (models.Carrier, error) {
	result := r.db.Create(&carrier)
	if result.Error != nil {
		return models.Carrier{}, result.Error
	}
	return carrier, nil
}

func (r *CarrierDB) Update(carrier models.Carrier) (models.Carrier, error) {
	return carrier, nil
}

func (r *CarrierDB) PartialUpdate(id int, fields map[string]interface{}) (models.Carrier, error) {
	var carrier models.Carrier
	result := r.db.First(&carrier, id)
	if result.Error != nil {
		return models.Carrier{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.Carrier{}, result.Error
	}

	if val, ok := fields["cid"]; ok {
		carrier.CarrierCode = val.(string)
	}
	if val, ok := fields["company_name"]; ok {
		carrier.Address = val.(string)
	}
	if val, ok := fields["address"]; ok {
		carrier.Address = val.(string)
	}
	if val, ok := fields["telephone"]; ok {
		carrier.Telephone = val.(string)
	}
	if val, ok := fields["locality_id"]; ok {
		carrier.LocalityId = int(val.(float64))
	}

	result = r.db.Save(&carrier)
	if result.Error != nil {
		return models.Carrier{}, result.Error
	}
	return carrier, nil
}

func (r *CarrierDB) Delete(id int) error {
	var carrier models.Carrier
	result := r.db.Delete(&carrier, id)
	return result.Error
}
