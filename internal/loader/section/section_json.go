package loader

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

type JsonFileSection struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

func NewSectionJson(path string) *JsonFileSection {
	return &JsonFileSection{
		path: path,
	}
}

func (l *JsonFileSection) Load() (v map[int]models.Section, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var sections = []models.Section{}
	err = json.NewDecoder(file).Decode(&sections)
	if err != nil {
		return
	}

	v = make(map[int]models.Section)
	for _, section := range sections {
		v[section.Id] = models.Section{
			Id:                 section.Id,
			SectionNumber:      section.SectionNumber,
			CurrentTemperature: section.CurrentTemperature,
			MinimumTemperature: section.MinimumTemperature,
			CurrentCapacity:    section.CurrentCapacity,
			MinimumCapacity:    section.MinimumCapacity,
			MaximumCapacity:    section.MaximumCapacity,
			WarehouseId:        section.WarehouseId,
			ProductTypeId:      section.ProductTypeId,
			ProductsBatch:      section.ProductsBatch,
		}

	}
	return
}
