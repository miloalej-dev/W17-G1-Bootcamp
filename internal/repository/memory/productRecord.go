package memory

import (
	loader "github.com/miloalej-dev/W17-G1-Bootcamp/internal/loader/json"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type ProductRecordMap struct {
	// db holds the product data. It is a private field to encapsulate storage.
	db map[int]models.ProductRecord
}

// NewProductMap is a constructor that creates and returns a new instance of ProductMap.
// It can be initialized with a pre-existing map of products.

func NewProductRecordMap() *ProductRecordMap {
	// Initialize with an empty map to ensure it's not nil.
	defaultDB := make(map[int]models.ProductRecord)

	ld := loader.NewProductRecordFile("docs/db/json/productRecords.json")
	db, err := ld.Load()
	if err != nil {
		return nil
	}

	if db != nil {
		defaultDB = db
	}

	return &ProductRecordMap{db: defaultDB}
}

// FindAll returns a copy of all productRecords currently stored in the repository.
// It returns an error if the operation fails, which is nil in this implementation.
func (r *ProductRecordMap) FindAll() ([]models.ProductRecord, error) {
	var productsRecords []models.ProductRecord

	if len(r.db) == 0 {
		return nil, repository.ErrEmptyEntity
	}

	for _, productRecord := range r.db {
		productsRecords = append(productsRecords, productRecord)
	}

	return productsRecords, nil

}

// Create adds a new product to the repository.
// It returns an error if a product with the same ID already exists.
func (r *ProductRecordMap) Create(productRecord models.ProductRecord) (models.ProductRecord, error) {

	newId := len(r.db)

	for {
		_, exists := r.db[productRecord.Id]
		if exists {
			newId++
		} else {
			break
		}
	}

	productRecord.Id = newId
	r.db[newId] = productRecord

	return productRecord, nil
}

func (r *ProductRecordMap) Update(productRecord models.ProductRecord) (models.ProductRecord, error) {
	_, exists := r.db[productRecord.Id]
	if !exists {
		return models.ProductRecord{}, repository.ErrEntityNotFound
	}
	r.db[productRecord.Id] = productRecord
	return productRecord, nil
}

// FindById searches for a product by its unique ID.
// It returns the found product or an error if the product does not exist.
func (r *ProductRecordMap) FindById(id int) (models.ProductRecord, error) {
	// Check if the product exists in the map.
	productRecord, exists := r.db[id]

	if !exists {
		// Return a descriptive error if the product is not found.
		return models.ProductRecord{}, repository.ErrEntityNotFound

	}
	// Return the found product.

	return productRecord, nil
}

// PartialUpdate updates specific fields of an existing product.
func (r *ProductRecordMap) PartialUpdate(id int, fields map[string]interface{}) (models.ProductRecord, error) {
	// 1. Check if the product exists
	productRecord, exists := r.db[id]
	if !exists {
		return models.ProductRecord{}, repository.ErrEntityNotFound
	}

	// 2. Iterate over the submitted fields and update only those.
	//    The keys ("product_code", "width", etc.) should match the JSON tags of your struct.
	for key, value := range fields {
		switch key {
		case "last_update":
			// Perform a type assertion to ensure the value is a string
			if lastUpdate, ok := value.(string); ok {
				productRecord.LastUpdate = lastUpdate
			}
		case "purchase_price":
			if purchasePrice, ok := value.(float64); ok {
				productRecord.PurchasePrice = purchasePrice
			}
		case "sale_price":
			if salePrice, ok := value.(float64); ok {
				productRecord.PurchasePrice = salePrice
			}

		case "product_id":
			if productId, ok := value.(int); ok {
				productRecord.ProductsId = productId
			}

		}

	}

	// 3. Save the updated product in the database
	r.db[id] = productRecord

	// 4. Return the modified product and a nil error to indicate success
	return productRecord, nil
}

func (r *ProductRecordMap) Delete(id int) (err error) {
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
