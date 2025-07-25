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

func (r *CarrierDB) FindByCid(cid string) (models.Carrier, bool, error) {
	var carrier models.Carrier
	result := r.db.Where("cid = ?", cid).First(&carrier)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Carrier{}, false, result.Error
		}
	}
	found := result.RowsAffected >= 1
	return carrier, found, nil
}

func (r *CarrierDB) Create(carrier models.Carrier) (models.Carrier, error) {
	result := r.db.Create(&carrier)
	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) {
			if mysqlErr.Number == 1452 {
				return models.Carrier{}, repository.ErrForeignKeyViolation
			}
		}
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
