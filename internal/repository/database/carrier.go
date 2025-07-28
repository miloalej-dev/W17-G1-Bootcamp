package database

import (
	"errors"
	"gorm.io/gorm"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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
		return nil, result.Error
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
		// Gorm treats not found as an error. In this case we explicitly want to find zero rows
		// so this is not an error. Everything else is
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Carrier{}, err
		}
	}

	if exists {
		return models.Carrier{}, repository.ErrEntityAlreadyExists
	}

	// 2- Create carrier
	result := r.db.Create(&carrier)
	switch {
	case errors.Is(result.Error, gorm.ErrForeignKeyViolated):
			return models.Carrier{}, repository.ErrLocalityNotFound
	}
	return carrier, result.Error
}

func (r *CarrierDB) Update(carrier models.Carrier) (models.Carrier, error) {
	var exists bool
	err := r.db.Model(&models.Carrier{}).
		Select("1").
		Where("`carriers`.`cid` = ? AND `carriers`.`id` <> ?", carrier.CId, carrier.ID).
		First(&exists).Error

	if err != nil {
		// Gorm treats not found as an error. In this case we explicitly want to find zero rows
		// so this is not an error. Everything else is
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Carrier{}, err
		}
	}

	if exists {
		return models.Carrier{}, repository.ErrEntityAlreadyExists
	}

	result := r.db.Save(&carrier)
	if result.Error == nil {
		return carrier, nil
	}
	return models.Carrier{}, result.Error
}

func (r *CarrierDB) PartialUpdate(id int, fields map[string]interface{}) (models.Carrier, error) {
	// 1- Validate that there is no carrier with this cid already
	if val, ok := fields["cid"]; ok {
		var exists bool
		err := r.db.Model(&models.Carrier{}).
			Select("1").
			Where("`carriers`.`cid` = ? AND `carriers`.`id` <> ?", val.(string), id).
			First(&exists).Error

		if err != nil {
			// Gorm treats not found as an error. In this case we explicitly want to find zero rows
			// so this is not an error. Everything else is
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Carrier{}, err
			}
		}

		if exists {
			return models.Carrier{}, repository.ErrEntityAlreadyExists
		}
	}

	var carrier models.Carrier
	result := r.db.First(&carrier, id)
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return models.Carrier{}, repository.ErrEntityNotFound
	case result.Error != nil:
			return models.Carrier{}, result.Error
	}

	if val, ok := fields["cid"]; ok {
		carrier.CId = val.(string)
	}
	if val, ok := fields["company_name"]; ok {
		carrier.CompanyName = val.(string)
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
	if result.RowsAffected < 1 {
		return repository.ErrEntityNotFound
	}
	return result.Error
}
