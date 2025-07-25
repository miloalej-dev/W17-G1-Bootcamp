package service

import "errors"

var (
	// ErrEntityNotFound is returned when an entity is not found
	ErrEntityNotFound = errors.New("entity not found")

	// ErrEntityAlreadyExists is returned when an entity being created already exists
	ErrEntityAlreadyExists = errors.New("entity already exists")

	// ErrInvalidEntity is returned when an entity is invalid
	ErrInvalidEntity = errors.New("invalid entity")

	// ErrEmptyEntity is returned when an entity is empty
	ErrEmptyEntity = errors.New("empty entity")

	ErrProductIdConflict = errors.New("the product with the given id does not exists")
	ErrProductNotFound   = errors.New("product not found")
)


