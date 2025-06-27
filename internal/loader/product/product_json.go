package product

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
func NewProductJSONFile(path string) *JsonFile {
	return &JsonFile{
		path: path,
	}
}

// Load is a method that loads the products
func (l *JsonFile) Load() (v map[int]models.Product, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var productos []models.Product
	err = json.NewDecoder(file).Decode(&productos)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]models.Product)
	for _, producto := range productos {
		v[producto.ID] = models.Product{
			ID:                             producto.ID,
			ProductCode:                    producto.ProductCode,
			Description:                    producto.Description,
			Width:                          producto.Width,
			Height:                         producto.Height,
			Length:                         producto.Length,
			NetWeight:                      producto.NetWeight,
			ExpirationRate:                 producto.ExpirationRate,
			RecommendedFreezingTemperature: producto.RecommendedFreezingTemperature,
			FreezingRate:                   producto.FreezingRate,
			ProductTypeID:                  producto.ProductTypeID,
			SellerID:                       producto.SellerID,
		}
	}

	return
}
