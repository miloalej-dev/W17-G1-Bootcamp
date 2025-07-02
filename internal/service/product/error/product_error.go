// Package errorProduct defines common errors that can occur during product operations.
// These exported error variables can be used by other packages to check for specific
// failure conditions when interacting with the product repository or service.

package errorProduct

import "errors"

// ErrorCreate is returned when an attempt to create a product fails because
// a product with the same ID already exists in the repository.
var ErrorCreate = errors.New("There is already a product with that ID")

// ErrorNotFound is returned when a lookup operation fails to find a product
// with the specified ID.
var ErrorNotFound = errors.New("Product Not found with that ID")
