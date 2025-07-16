package database

import (
	"gorm.io/gorm"
	//"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionRepository struct {
	db *gorm.DB
}

func NewSectionRepository(db *gorm.DB) *SectionRepository {
	return &SectionRepository{
		db: db,
	}
}

func (r *SectionRepository) FindAll() ([]models.Section, error) {
	sections := make([]models.Section, 0)
	result := r.db.Find(&sections)
	if result.Error != nil {
		return make([]models.Section, 0), result.Error
	}
	return sections, nil
}

func (r *SectionRepository) FindById(id int) (models.Section, error) {
	var section models.Section
	result := r.db.First(&section, id)
	if result.Error != nil {
		return models.Section{}, result.Error
	}
	return section, nil
}

func (r *SectionRepository) Create(section models.Section) (models.Section, error) {
	result := r.db.Create(&section)
	if result.Error != nil {
		return models.Section{}, result.Error
	}
	return section, nil
}

func (s *SectionRepository) Update(section models.Section) (models.Section, error) {
	result := s.db.Save(&section)

	if result.Error != nil {
		return models.Section{}, result.Error
	}

	return section, nil
}
func (r *SectionRepository) PartialUpdate(id int, fields map[string]interface{}) (models.Section, error) {
	var section models.Section
	result := r.db.First(&section, id)
	if result.Error != nil {
		return models.Section{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.Section{}, result.Error
	}

	if val, ok := fields["section_number"]; ok {
		section.SectionNumber = val.(string)
	}
	if val, ok := fields["current_temperature"]; ok {
		section.CurrentTemperature = val.(float64)
	}
	if val, ok := fields["minimum_temperature"]; ok {
		section.MinimumTemperature = val.(float64)
	}
	if val, ok := fields["current_capacity"]; ok {
		section.CurrentCapacity = int(val.(float64))
	}
	if val, ok := fields["minimum_capacity"]; ok {
		section.MinimumCapacity = int(val.(float64))
	}
	if val, ok := fields["maximum_capacity"]; ok {
		section.MaximumCapacity = int(val.(float64))
	}
	if val, ok := fields["warehouse_id"]; ok {
		section.WarehouseId = int(val.(float64))
	}
	if val, ok := fields["product_type_id"]; ok {
		section.ProductTypeId = int(val.(float64))
	}

	result = r.db.Save(&section)
	if result.Error != nil {
		return models.Section{}, result.Error
	}
	return section, nil
}

func (r *SectionRepository) Delete(id int) error {
	var section models.Section
	result := r.db.Delete(&section, id)
	return result.Error
}
