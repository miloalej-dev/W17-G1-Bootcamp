package database

import (
	"errors"
	"gorm.io/gorm"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/go-sql-driver/mysql"
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
	// 1- Validate that there is no carrier with this cid already
	var exists bool
	err := r.db.Model(&models.Carrier{}).
		Select("1").
		Where("cid = ?", carrier.CId).
		First(&exists).Error

	if err != nil {
		// Record not found is not an error in this context. Everythinf else is
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Carrier{}, err
		}
	}

	if exists {
		return models.Carrier{}, repository.ErrEntityAlreadyExists
	}

	// 2- Create carrier
	result := r.db.Create(&carrier)
	if result.Error == nil {
		return carrier, nil
	}

	// 2.1- Check for MySql specific errors
	var mysqlErr *mysql.MySQLError
	if errors.As(result.Error, &mysqlErr) {
		switch mysqlErr.Number {
		// Locality does not exist
		case 1452: 
			return models.Carrier{}, repository.ErrLocalityNotFound
		}
	}

	// Return original error if not handled above
	return models.Carrier{}, result.Error
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
		carrier.CId = val.(string)
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
