package sectionRepository

import (
	"errors"
	sectionLoader "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/section"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionMap struct {
	db map[int]models.Section
}

func NewSectionMap(db map[int]models.Section) *SectionMap {
	// defaultDb is an empty map
	sectionLoader.NewSectionJson("")
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}
	return &SectionMap{db: defaultDb}
}

func (r *SectionMap) FindAll() (v map[int]models.Section, err error) {
	v = make(map[int]models.Section)
	// copy db
	for key, value := range r.db {
		v[key] = value
	}
	return
}

func (r *SectionMap) FindByID(id int) (models.Section, error) {
	v, exist := r.db[id]
	if !exist {
		return models.Section{}, errors.New("Section not found")
	}
	return v, nil

}

func (r *SectionMap) Add(s models.Section) (models.Section, error) {
	v, exist := r.db[s.Id]
	if exist {
		return v, errors.New("Section already exists")
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

func (r *SectionMap) Delete(s models.Section) (models.Section, error) {
	v, exist := r.db[s.Id]
	if !exist {
		return models.Section{}, errors.New("Section not found")
	}
	delete(r.db, v.Id)
	return v, nil
}
