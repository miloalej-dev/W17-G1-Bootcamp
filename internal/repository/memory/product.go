// Package productRepository provides an in-memory implementation of a repository for products.
// It uses a map to store product data, making it suitable for testing or simple applications
// where data persistence is not required.

package memory

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

// ProductMap implements a product repository using an in-memory map.
// The key of the map is the product ID.
type ProductMap struct {
	// db holds the product data. It is a private field to encapsulate storage.
	db map[int]models.Product
}

// NewProductMap is a constructor that creates and returns a new instance of ProductMap.
// It can be initialized with a pre-existing map of products.

func NewProductMap(db map[int]models.Product) *ProductMap {
	// Initialize with an empty map to ensure it's not nil.
	defaultDb := make(map[int]models.Product)
	if db != nil {
		// If an initial database is provided, use it.
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}

// FindAll returns a copy of all products currently stored in the repository.
// It returns an error if the operation fails, which is nil in this implementation.
func (r *ProductMap) FindAll() (allProducts []models.Product, err error) {
	var products []models.Product

	if len(r.db) == 0 {
		err = errors.New("table empty")
		return
	}

	for _, product := range r.db {
		products = append(products, product)
	}

	allProducts = products
	return
}

// Create adds a new product to the repository.
// It returns an error if a product with the same ID already exists.
func (r *ProductMap) Create(body models.Product) (newProduct models.Product, err error) {
	// Create the new product to the map in database.
	r.db[body.Id] = body
	newProduct = r.db[body.Id]
	return
}

func (r *ProductMap) Update(body models.Product) (product models.Product, err error) {
	_, exists := r.db[body.Id]
	if !exists {
		err = errors.New("product does not exist")
		return
	}
	r.db[body.Id] = body
	return body, nil
}

// FindById searches for a product by its unique ID.
// It returns the found product or an error if the product does not exist.
func (r *ProductMap) FindById(id int) (product models.Product, err error) {
	// Check if the product exists in the map.
	value, exists := r.db[id]
	if !exists {
		// Return a descriptive error if the product is not found.
		err = errors.New("product not found")
		return
	}
	// Return the found product.
	product = value
	return
}

func (r *ProductMap) PartialUpdate(id int, fields map[string]interface{}) (product models.Product, err error) {
	return
}

func (r *ProductMap) PartialUpdateV2(id int, body models.Product) (product models.Product, err error) {
	value, exists := r.db[id]
	if !exists {
		// Return a descriptive error if the product is not found.
		err = errors.New("product not found")
		return
	}

	if body.ProductCode != "" {
		value.ProductCode = body.ProductCode
	}
	if body.Description != "" {
		value.Description = body.Description
	}

	if body.Width != 0 {
		value.Width = body.Width
	}

	if body.Height != 0 {
		value.Height = body.Height
	}
	if body.Length != 0 {
		value.Length = body.Length
	}

	if body.NetWeight != 0 {
		value.NetWeight = body.NetWeight
	}
	if body.ExpirationRate != 0 {
		value.ExpirationRate = body.ExpirationRate
	}
	if body.RecommendedFreezingTemperature != 0 {
		value.RecommendedFreezingTemperature = body.RecommendedFreezingTemperature
	}
	if body.FreezingRate != 0 {
		value.FreezingRate = body.FreezingRate
	}
	if body.ProductTypeId != 0 {
		value.ProductTypeId = body.ProductTypeId
	}
	if body.SellerId != 0 {
		value.SellerId = body.SellerId
	}
	// Overwrite Product in map with new values
	r.db[id] = value
	product = r.db[id]
	return
}
func (r *ProductMap) Delete(id int) (err error) {
	_, exists := r.db[id]
	if !exists {
		// Return a descriptive error if the product is not found.
		err = errors.New("1")
		return
	}
	// Deletes product from the map
	delete(r.db, id)
	return
}
