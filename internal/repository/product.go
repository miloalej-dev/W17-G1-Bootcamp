// Package repository defines the interfaces and implementations for product data access.
package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// ProductRepository Inherits the basic CRUD methods from the Generic Repository[int, models.Product]
type ProductRepository interface {
	// Repository is a generic repository interface for CRUD operations
	Repository[int, models.Product]
	FindRecordsCountByProductId(id int) (models.ProductReport, error)
}
