package repository

import "errors"

var (
	// ErrProductReportNotFound is returned when a product report cannot be genereated
	ErrProductReportNotFound = errors.New("section sent not found in system")
	// ErrProductNotFound is returned when product is not found
	ErrProductNotFound = errors.New("product not found")
	// ErrProductAlreadyExists is returned when a product already exist on creation or update
	ErrProductAlreadyExists = errors.New("product already exists")

	// ErrEntityNotFound is returned when an entity is not found
	ErrEntityNotFound = errors.New("entity not found")

	// ErrEntityAlreadyExists is returned when an entity being created already exists
	ErrEntityAlreadyExists = errors.New("entity already exists")

	// ErrInvalidEntity is returned when an entity is invalid
	ErrInvalidEntity = errors.New("invalid entity")

	// ErrProductBatch is returned when a product batch already exists
	ErrProductBatchAlreadyExists = errors.New("product batch already exists")

	// ErrEmptyEntity is returned when an entity is empty
	ErrEmptyEntity = errors.New("empty entity")

	// ErrForeignKeyViolation is returned when a foreing key does not exist
	ErrForeignKeyViolation = errors.New("referenced entity does note exist")

	// ErrSectionNotFound is returned when a section is not found in get reports
	ErrSectionNotFound = errors.New("there is no section by that id")

	// ErrNoSectionFound is returned when a section is not found in get reports
	ErrEmptyReport = errors.New("there is no reports asociated to that section_id")

	ErrProvinceNotFound = errors.New("province not found")
)
