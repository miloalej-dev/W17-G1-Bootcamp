package memory

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// BuyerMap is a struct that represents a Buyer repository
type BuyerMap struct {
	// db is a map of buyers
	db map[int]models.Buyer
}

// NewBuyerMap Creates a new Buyer repository
func NewBuyerMap(db map[int]models.Buyer) *BuyerMap {

	// default db
	defaultDb := make(map[int]models.Buyer)

	if db != nil {
		defaultDb = db
	}
	return &BuyerMap{db: defaultDb}
}

// FindAll is a method that returns a map of all buyers
func (r *BuyerMap) FindAll() (buyers map[int]models.Buyer, err error) {

	buyers = make(map[int]models.Buyer)
	for key, value := range r.db {
		buyers[key] = value
	}

	return
}

// FindById is a method that returns a buyer by its unique id
func (r *BuyerMap) FindById(id int) (v *models.Buyer, err error) {
	for _, value := range r.db {
		if value.Id == id {
			return &value, nil
		}

	}
	return nil, repository.ErrEntityNotFound

}

// Create  is a method that creates a new buyer
func (r *BuyerMap) Create(buyer models.BuyerAtributtes) (v *models.Buyer, err error) {

	for _, value := range r.db {
		if value.CardNumberId == buyer.CardNumberId {
			return nil, repository.ErrEntityAlreadyExists
		}

	}

	newId := len(r.db)

	for {
		_, err = r.FindById(newId)
		if err == nil {
			newId++
		} else {
			break
		}
	}

	b := models.Buyer{
		Id:              newId,
		BuyerAtributtes: buyer,
	}

	r.db[newId] = b

	return &b, nil

}

// Delete is a method that removes a buyer from the repository
func (r *BuyerMap) Delete(id int) (*models.Buyer, error) {
	b, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	delete(r.db, id)

	return b, nil
}

// Update is a method that modifies an existing Buyer
func (r *BuyerMap) Update(buyer models.Buyer) (*models.Buyer, error) {
	_, err := r.FindById(buyer.Id)

	if err != nil {
		return nil, err
	}
	r.db[buyer.Id] = buyer
	return &buyer, nil

}
