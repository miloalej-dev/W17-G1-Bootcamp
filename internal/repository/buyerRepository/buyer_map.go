package buyerRepository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

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
