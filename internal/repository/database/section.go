package database

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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
	if val, ok := fields["warehouses_id"]; ok {
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

func (r *SectionRepository) FindSectionReport(id int) (models.SectionReport, error) {
	var report models.SectionReport
	var section models.Section
	exists := r.db.First(&section, id)

	if exists.Error != nil {
		return models.SectionReport{}, repository.ErrSectionNotFound
	}

	// gorm query requierment 3
	result := r.db.Table("sections as s"). // alias for table sections
						Select("s.id as section_id, s.section_number, COUNT(p.id) as products_count"). // map data found to struct
						Joins("INNER JOIN product_batches as p ON s.id = p.section_id").               // join product_batches
						Where("s.id = ?", id).                                                         // filter by id given
						Group("s.id, s.section_number").                                               // group to make COUNT() work proppertly
						Scan(&report)                                                                  // scan the result into the model variable

	if result.Error != nil {
		return models.SectionReport{}, result.Error
	}

	// if there is no registry, means there is no Section with that id
	if result.RowsAffected == 0 {
		return models.SectionReport{}, repository.ErrEmptyReport
	}

	return report, nil
}

// GetAllSectionReports obtiene el reporte de productos para TODAS las secciones.
func (r *SectionRepository) FindAllSectionReports() ([]models.SectionReport, error) {
	var reports []models.SectionReport

	// La consulta es casi idéntica, pero sin el .Where() y escaneando en un slice.
	result := r.db.Table("sections as s").
		Select("s.id as section_id, s.section_number, COUNT(p.id) as products_count").
		Joins("INNER JOIN product_batches as p ON s.id = p.section_id").
		Group("s.id, s.section_number").
		Scan(&reports)

	if result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}
