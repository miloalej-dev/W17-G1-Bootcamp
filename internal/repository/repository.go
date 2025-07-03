package repository

// Repository is a generic repository interface for CRUD operations
// It is parameterized by K for the key type and T for the entity type.
type Repository[K comparable, T any] interface {
	// FindAll returns all entities
	FindAll() ([]T, error)

	// FindById returns an entity by K
	FindById(id K) (T, error)

	// Create adds a new entity and returns the created entity
	Create(entity T) (T, error)

	// Update updates an existing entity and returns the updated entity
	Update(entity T) (T, error)

	// PartialUpdate updates specific fields of an existing entity and returns the updated entity
	PartialUpdate(id K, fields map[string]interface{}) (T, error)

	// Delete removes an entity by K
	Delete(id K) error
}
