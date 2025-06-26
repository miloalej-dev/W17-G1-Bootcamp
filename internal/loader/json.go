package loader

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// JsonFile is a struct that implements the LoaderVehicle interface
type JsonFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// NewVehicleJSONFile is a function that returns a new instance of JsonFile
func NewVehicleJSONFile(path string) *JsonFile {
	return &JsonFile{
		path: path,
	}
}

// Load is a method that loads the foos
func (l *JsonFile) Load() (v map[int]models.Foo, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var foos []models.FooDoc
	err = json.NewDecoder(file).Decode(&foos)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]models.Foo)
	for _, foo := range foos {
		v[foo.ID] = models.Foo{
			ID:   foo.ID,
			Name: foo.Name,
		}
	}

	return
}
