// Package productRepository provides an in-memory implementation of a repository for products.
// It uses a map to store product data, making it suitable for testing or simple applications
// where data persistence is not required.

package memory

import (
	loader "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
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

func NewProductMap() *ProductMap {
	// Initialize with an empty map to ensure it's not nil.
	defaultDB := make(map[int]models.Product)

	ld := loader.NewProductFile("docs/db/products.json")
	db, err := ld.Load()
	if err != nil {
		return nil
	}

	if db != nil {
		defaultDB = db
	}

	return &ProductMap{db: defaultDB}
}

// FindAll returns a copy of all products currently stored in the repository.
// It returns an error if the operation fails, which is nil in this implementation.
func (r *ProductMap) FindAll() (allProducts []models.Product, err error) {
	var products []models.Product

	if len(r.db) == 0 {
		err = repository.ErrEmptyEntity
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
	id := len(r.db) + 1
	body.Id = id
	r.db[body.Id] = body
	newProduct = r.db[body.Id]
	return
}

func (r *ProductMap) Update(body models.Product) (product models.Product, err error) {
	_, exists := r.db[body.Id]
	if !exists {
		err = repository.ErrEntityNotFound
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
		err = repository.ErrEntityNotFound

		return
	}
	// Return the found product.
	product = value
	return
}

// PartialUpdate updates specific fields of an existing product.
func (r *ProductMap) PartialUpdate(id int, fields map[string]interface{}) (models.Product, error) {
	// 1. Check if the product exists
	product, exists := r.db[id]
	if !exists {
		return models.Product{}, repository.ErrEntityNotFound
	}

	// 2. Iterate over the submitted fields and update only those.
	//    The keys ("product_code", "width", etc.) should match the JSON tags of your struct.
	for key, value := range fields {
		switch key {
		case "product_code":
			// Perform a type assertion to ensure the value is a string
			if v, ok := value.(string); ok {
				product.ProductCode = v
			}
		case "description":
			if v, ok := value.(string); ok {
				product.Description = v
			}
		case "width":
			// Numbers in JSON are decoded as float64 in Go by default
			if v, ok := value.(float64); ok {
				product.Width = v
			}
		case "height":
			if v, ok := value.(float64); ok {
				product.Height = v
			}
		case "length":
			if v, ok := value.(float64); ok {
				product.Length = v
			}
		case "net_weight":
			if v, ok := value.(float64); ok {
				product.NetWeight = v
			}
		case "expiration_rate":
			// If the field in the struct is an int, it must be cast
			if v, ok := value.(float64); ok {
				product.ExpirationRate = v
			}
		case "recommended_freezing_temperature":
			if v, ok := value.(float64); ok {
				product.RecommendedFreezingTemperature = v
			}
		case "freezing_rate":
			if v, ok := value.(float64); ok {
				product.FreezingRate = v
			}
		case "product_type_id":
			if v, ok := value.(float64); ok {
				product.ProductTypeId = int(v)
			}
		case "seller_id":
			if v, ok := value.(float64); ok {
				product.SellerId = int(v)
			}
		}
	}

	// 3. Save the updated product in the database
	r.db[id] = product

	// 4. Return the modified product and a nil error to indicate success
	return product, nil
}

func (r *ProductMap) Delete(id int) (err error) {
	_, exists := r.db[id]
	if !exists {
		// Return a descriptive error if the product is not found.
		err = repository.ErrEntityNotFound
		return
	}
	// Deletes product from the map
	delete(r.db, id)
	return
}
