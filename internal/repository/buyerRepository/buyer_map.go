package buyerRepository

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

func NewBuyerMap(db map[int]models.Buyer) *BuyerMap {
	// default db
	defaultDb := make(map[int]models.Buyer)
	if db != nil {
		defaultDb = db
	}
	return &BuyerMap{db: defaultDb}
}

// BuyerMap is a struct that represents a Buyer repository
type BuyerMap struct {
	// db is a map of buyers
	db map[int]models.Buyer
}

// FindAll is a method that returns a map of all buyers
func (r *BuyerMap) FindAll() (v map[int]models.Buyer, err error) {

	v = make(map[int]models.Buyer)
	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *BuyerMap) FindById(id int) (v *models.Buyer, err error) {
	for _, value := range r.db {
		if value.Id == id {
			return &value, nil
		}

	}
	return nil, errors.New("buyer not found")

}
func (r *BuyerMap) Create(buyer models.BuyerAtributtes) (v *models.Buyer, err error) {

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

func (r *BuyerMap) Delete(id int) (*models.Buyer, error) {
	b, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	delete(r.db, id)

	return b, nil
}

func (r *BuyerMap) Update(buyer models.Buyer) (*models.Buyer, error) {
	_, err := r.FindById(buyer.Id)
	if err != nil {
		return nil, err
	}
	r.db[buyer.Id] = buyer
	return &buyer, nil

}
