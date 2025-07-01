package sectionRepository

import (
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
