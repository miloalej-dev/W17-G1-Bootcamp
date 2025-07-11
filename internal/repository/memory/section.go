package memory

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionMap struct {
	db map[int]models.Section
}

func NewSectionMap(db map[int]models.Section) *SectionMap {
	// defaultDb is an empty map
	defaultDB := make(map[int]models.Section)
	if db != nil {
		defaultDB = db
	}
	return &SectionMap{db: defaultDB}
}
func (r *SectionMap) FindAll() ([]models.Section, error) {
	v := make([]models.Section, 0)
	// copy db
	for _, value := range r.db {
		v = append(v, value)
	}
	return v, nil
}

func (r *SectionMap) FindById(id int) (models.Section, error) {
	v, exist := r.db[id]

	if !exist {
		return models.Section{}, repository.ErrEntityNotFound
	}
	return v, nil

}

func (r *SectionMap) Create(s models.Section) (models.Section, error) {
	if s.Id == 0 {
		s.Id = len(r.db) + 1
	}
	v, exist := r.db[s.Id]
	if exist {
		return v, repository.ErrEntityAlreadyExists
	}
	sc, err := r.FindBySection(s.SectionNumber)
	if err == nil {
		return sc, repository.ErrEntityAlreadyExists
	}
	r.db[s.Id] = s
	return s, nil
}

func (r *SectionMap) Update(s models.Section) (models.Section, error) {
	v, exist := r.db[s.Id]
	if !exist {
		return models.Section{}, repository.ErrEntityNotFound
	}
	r.db[s.Id] = s
	return v, nil
}

func (r *SectionMap) PartialUpdate(id int, fields map[string]interface{}) (models.Section, error) {
	v, exist := r.db[id]
	if !exist {
		return models.Section{}, repository.ErrEntityNotFound
	}
	for key, value := range fields {
		switch key {
		case "section_number":
			if section_number, ok := value.(float64); ok {
				_, err := r.FindBySection(int(section_number))
				if err == nil {
					return models.Section{}, repository.ErrInvalidEntity

				} else {
					v.SectionNumber = int(section_number)
				}
			}
		case "current_temperature":
			if current_temperature, ok := value.(float64); ok {
				v.CurrentTemperature = int(current_temperature)
			}
		case "minimum_temperature":
			if minimum_temperature, ok := value.(float64); ok {
				v.MinimumTemperature = int(minimum_temperature)
			}
		case "current_capacity":
			if current_capacity, ok := value.(float64); ok {
				v.CurrentCapacity = int(current_capacity)
			}
		case "minimum_capacity":
			if minimum_capacity, ok := value.(float64); ok {
				v.MinimumCapacity = int(minimum_capacity)
			}
		case "maximum_capacity":
			if maximum_capacity, ok := value.(float64); ok {
				v.MaximumCapacity = int(maximum_capacity)
			}
		case "warehouse_id":
			if warehouse_id, ok := value.(float64); ok {
				v.WarehouseId = int(warehouse_id)
			}
		case "product_type_id":
			if product_type_id, ok := value.(float64); ok {
				v.ProductTypeId = int(product_type_id)
			}
		}

	}
	r.db[id] = v

	return v, nil

}

func (r *SectionMap) Delete(id int) error {
	v, exist := r.db[id]
	if !exist {
		return repository.ErrEntityNotFound
	}
	delete(r.db, v.Id)
	return nil
}

func (r *SectionMap) FindBySection(section int) (models.Section, error) {
	for _, v := range r.db {
		if v.SectionNumber == section {
			return v, nil
		}
	}
	return models.Section{}, repository.ErrEntityNotFound
}
