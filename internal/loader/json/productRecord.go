package json

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

type ProductRecordFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// NewProductFile is a function that returns a new instance of EmployeeFile
func NewProductRecordFile(path string) *ProductRecordFile {
	return &ProductRecordFile{
		path: path,
	}
}

// Load is a method that loads the products
func (l *ProductRecordFile) Load() (v map[int]models.ProductRecord, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var productRecords []models.ProductRecord
	err = json.NewDecoder(file).Decode(&productRecords)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]models.ProductRecord)
	for _, product := range productRecords {
		v[product.Id] = models.ProductRecord{
			Id:             product.Id,
			LastUpdateDate: product.LastUpdateDate,
			PurchasePrice:  product.PurchasePrice,
			SalePrice:      product.SalePrice,
			ProductId:      product.ProductId,
		}
	}

	return
}
