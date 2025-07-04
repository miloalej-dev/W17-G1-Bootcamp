package memory

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionMap struct {
	db map[int]models.Section
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
		return models.Section{}, errors.New("Section not found")
	}
	return v, nil

}

func (r *SectionMap) Create(s models.Section) (models.Section, error) {
	if s.Id == 0 {
		s.Id = len(r.db) + 1
	}
	v, exist := r.db[s.Id]
	if exist {
		return v, errors.New("Section already exists")
	}
	sc, err := r.FindBySection(s.SectionNumber)
	if err == nil {
		return sc, errors.New("Section already exists")
	}
	r.db[s.Id] = s
	return s, nil
}

func (r *SectionMap) Update(s models.Section) (models.Section, error) {
	v, exist := r.db[s.Id]
	if !exist {
		return models.Section{}, errors.New("section not found")
	}
	if s.SectionNumber != 0 {
		v.SectionNumber = s.SectionNumber
	}
	if s.CurrentTemperature != 0 {
		v.CurrentTemperature = s.CurrentTemperature
	}
	if s.MinimumTemperature != 0 {
		v.MinimumTemperature = s.MinimumTemperature
	}
	if s.CurrentCapacity != 0 {
		v.CurrentCapacity = s.CurrentCapacity
	}
	if s.MinimumCapacity != 0 {
		v.MinimumCapacity = s.MinimumCapacity
	}
	if s.MaximumCapacity != 0 {
		v.MinimumTemperature = s.MinimumTemperature
	}
	if s.WarehouseId != 0 {
		v.WarehouseId = s.WarehouseId
	}
	if s.ProductTypeId != 0 {
		v.ProductTypeId = s.ProductTypeId
	}
	if len(s.ProductsBatch) != 0 {
		v.ProductsBatch = s.ProductsBatch
	}
	return v, nil
}

func (r *SectionMap) PartialUpdate(id int, fields map[string]interface{}) (models.Section, error) {
	return models.Section{}, nil
}

func (r *SectionMap) Delete(id int) error {
	v, exist := r.db[id]
	if !exist {
		return errors.New("Section not found")
	}
	delete(r.db, v.Id)
	return nil
}

func NewSectionMap() *SectionMap {
	// defaultDb is an empty map
	defaultDB := make(map[int]models.Section)
	ld := json.NewSectionFile("docs/db/sections.json")
	db, err := ld.Load()
	if err != nil {
		return nil
	}
	if db != nil {
		defaultDB = db
	}
	return &SectionMap{db: defaultDB}
}

func (r *SectionMap) FindBySection(section int) (models.Section, error) {
	for _, v := range r.db {
		if v.SectionNumber == section {
			return v, nil
		}
	}
	return models.Section{}, errors.New("Section not found")
}
