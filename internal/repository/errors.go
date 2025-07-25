package repository

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

	// ErrForeignKeyViolation is returned when a foreing key does not exist
	ErrForeignKeyViolation = errors.New("referenced entity does note exist")

	// ErrSectionNotFound is returned when a section is not found in get reports
	ErrSectionNotFound = errors.New("there is no section by that id")

	// ErrNoSectionFound is returned when a section is not found in get reports
	ErrEmptyReport = errors.New("there is no reports asociated to that section_id")
)