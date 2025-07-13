package memory

import (
	"errors"
	loader "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// SellerMap implements a SellerRepository using an in-memory map.
// The key of the map is the seller Id.
type SellerMap struct {
	// db holds the product data. It is a private field to encapsulate storage.
	db map[int]models.Seller
}

// NewSellerMap is a constructor that creates and returns a new instance of SellerMap.
// It can be initialized with a pre-existing map of seller.
func NewSellerMap() *SellerMap {

	// defaultDb is an empty map
	defaultDB := make(map[int]models.Seller)

	ld := loader.NewSellerFile("docs/db/json/sellers.json")
	db, err := ld.Load()

	if err != nil {
		return nil
	}
	if db != nil {
		defaultDB = db
	}
	return &SellerMap{db: defaultDB}
}

func (s *SellerMap) FindAll() ([]models.Seller, error) {
	var sellers []models.Seller

	for _, seller := range s.db {
		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (s *SellerMap) FindById(id int) (models.Seller, error) {
	seller, found := s.db[id]

	if !found {
		return models.Seller{}, errors.New("seller does not exist")
	}

	return seller, nil
}

func (s *SellerMap) Create(entity models.Seller) (models.Seller, error) {
	id := len(s.db) + 1
	entity.Id = id
	s.db[id] = entity
	return entity, nil
}

func (s *SellerMap) Update(entity models.Seller) (models.Seller, error) {
	_, found := s.db[entity.Id]
	if !found {
		return models.Seller{}, errors.New("seller does not exist")
	}
	s.db[entity.Id] = entity
	return entity, nil
}

func (s *SellerMap) PartialUpdate(id int, fields map[string]interface{}) (models.Seller, error) {
	seller, found := s.db[id]

	if !found {
		return models.Seller{}, errors.New("seller does not exist")
	}

	// Update only the fields that are provided
	// this is based on the json tag of the struct
	for key, value := range fields {
		switch key {
		case "name":
			if name, ok := value.(string); ok {
				seller.Name = name
			}
		case "address":
			if address, ok := value.(string); ok {
				seller.Address = address
			}
		case "telephone":
			if telephone, ok := value.(string); ok {
				seller.Telephone = telephone
			}
		}
	}

	// Update the seller in the database
	s.db[id] = seller
	return seller, nil
}

func (s *SellerMap) Delete(id int) error {
	_, found := s.db[id]
	if !found {
		return errors.New("seller does not exist")
	}
	delete(s.db, id)
	return nil
}
