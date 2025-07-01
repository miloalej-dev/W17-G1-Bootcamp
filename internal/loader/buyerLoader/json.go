package buyerLoader

import (
	"encoding/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"os"
)

// JsonFile is a struct that implements the LoaderBuyer interface
type JsonFile struct {
	// path is the path to the file that contains the buyers in JSON format
	path string
}

// NewBuyerJSONFile is a function that returns a new instance of JsonFile
func NewBuyerJSONFile(path string) *JsonFile {
	return &JsonFile{
		path: path,
	}
}

// Load is a method that loads the buyers
func (l *JsonFile) Load() (v map[int]models.Buyer, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var buyers []models.BuyerDoc
	err = json.NewDecoder(file).Decode(&buyers)
	if err != nil {
		return
	}

	// serialize buyers
	v = make(map[int]models.Buyer)
	for _, buyer := range buyers {
		v[buyer.Id] = models.Buyer{
			Id: buyer.Id,
			BuyerAtributtes: models.BuyerAtributtes{
				CardNumberId: buyer.CardNumberId,
				FirstName:    buyer.FirstName,
				LastName:     buyer.LastName,
			},
		}
	}

	return
}
