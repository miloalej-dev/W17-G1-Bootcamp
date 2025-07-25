package memory

import (
	loader "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// BuyerMap is a struct that represents a Buyer repository
type BuyerMap struct {
	// db is a map of buyers
	db map[int]models.Buyer
}

// NewBuyerMap Creates a new Buyer repository
func NewBuyerMap() *BuyerMap {

	// defaultDb is an empty map
	defaultDB := make(map[int]models.Buyer)

	// default db
	ld := loader.NewBuyerFile("docs/db/json/buyers.json")
	db, err := ld.Load()

	if err != nil {
		return nil
	}
	if db != nil {
		defaultDB = db
	}
	return &BuyerMap{db: defaultDB}
}

// FindAll is a method that returns a map of all buyers
func (r *BuyerMap) FindAll() ([]models.Buyer, error) {

	var buyers []models.Buyer
	for _, value := range r.db {
		buyers = append(buyers, value)
	}

	return buyers, nil
}

// FindById is a method that returns a buyer by its unique id
func (r *BuyerMap) FindById(id int) (models.Buyer, error) {

	for _, value := range r.db {
		if value.Id == id {
			return value, nil
		}
	}
	return models.Buyer{}, repository.ErrEntityNotFound

}

// Create  is a method that creates a new buyer
func (r *BuyerMap) Create(buyer models.Buyer) (models.Buyer, error) {

	for _, value := range r.db {
		if value.CardNumberId == buyer.CardNumberId {
			return models.Buyer{}, repository.ErrEntityAlreadyExists
		}
	}

	newId := len(r.db)

	for {
		_, err := r.FindById(newId)
		if err == nil {
			newId++
		} else {
			break
		}
	}

	buyer.Id = newId
	r.db[newId] = buyer

	return buyer, nil

}

// Update is a method that modifies an existing Buyer
func (r *BuyerMap) Update(buyer models.Buyer) (models.Buyer, error) {
	_, err := r.FindById(buyer.Id)

	if err != nil {
		return models.Buyer{}, err
	}
	r.db[buyer.Id] = buyer
	return buyer, nil

}

// Delete is a method that removes a buyer from the repository
func (r *BuyerMap) Delete(id int) error {
	_, err := r.FindById(id)
	if err != nil {
		return err
	}

	delete(r.db, id)

	return nil
}

// TODO
func (r *BuyerMap) PartialUpdate(id int, fields map[string]interface{}) (models.Buyer, error) {

	Buyer, err := r.FindById(id)
	if err != nil {
		return models.Buyer{}, err
	}

	for key, value := range fields {
		switch key {
		case "card_number_id":
			if CardNumberId, ok := value.(string); ok {
				Buyer.CardNumberId = CardNumberId
			}
		case "first_name":
			if FirstName, ok := value.(string); ok {
				Buyer.FirstName = FirstName
			}
		case "last_name":
			if Lastname, ok := value.(string); ok {
				Buyer.LastName = Lastname
			}
		}
	}

	// Update the seller in the database
	r.db[id] = Buyer

	return Buyer, nil
}
