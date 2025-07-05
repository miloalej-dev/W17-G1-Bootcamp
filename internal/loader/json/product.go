package json

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// ProductFile is a struct that implements the LoaderVehicle interface
type ProductFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// NewProductFile is a function that returns a new instance of EmployeeFile
func NewProductFile(path string) *ProductFile {
	return &ProductFile{
		path: path,
	}
}

// Load is a method that loads the products
func (l *ProductFile) Load() (v map[int]models.Product, err error) {
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
		v[producto.Id] = models.Product{
			Id:                             producto.Id,
			ProductCode:                    producto.ProductCode,
			Description:                    producto.Description,
			Width:                          producto.Width,
			Height:                         producto.Height,
			Length:                         producto.Length,
			NetWeight:                      producto.NetWeight,
			ExpirationRate:                 producto.ExpirationRate,
			RecommendedFreezingTemperature: producto.RecommendedFreezingTemperature,
			FreezingRate:                   producto.FreezingRate,
			ProductTypeId:                  producto.ProductTypeId,
			SellerId:                       producto.SellerId,
		}
	}

	return
}
