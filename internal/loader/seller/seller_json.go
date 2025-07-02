package seller

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// JsonFile is a struct that implements the Loader interface
type JsonFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// NewJSONFile is a function that returns a new instance of JsonFile
func NewJSONFile(path string) *JsonFile {
	return &JsonFile{
		path: path,
	}
}

func (j *JsonFile) Load() (data map[int]models.Seller, err error) {

	// open file
	file, err := os.Open(j.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var sellers []models.Seller
	err = json.NewDecoder(file).Decode(&sellers)
	if err != nil {
		return
	}

	// serialize vehicles
	data = make(map[int]models.Seller)

	for _, seller := range sellers {
		data[seller.Id] = models.Seller{
			Id:        seller.Id,
			Name:      seller.Name,
			Address:   seller.Address,
			Telephone: seller.Telephone,
		}
	}

	return

}
