package gateway

import "github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"

// FooRepository defines the contract for interacting with
// persistent storage for Foo entities.
//
// Implementations of this interface are responsible for handling
// the creation, retrieval, and deletion of Foo records from the data source.
type FooRepository interface {
	// Save persists a new Foo entity to the storage.
	// Returns the saved Foo with generated ID and timestamp, or an error.
	Save(foo entity.Foo) (entity.Foo, error)

	// GetById retrieves a Foo entity by its ID.
	// Returns the found Foo or an error. If no entity is found,
	// implementations should return gateway.ErrorFooNotFound.
	GetById(id int64) (entity.Foo, error)

	// DeleteById removes a Foo entity by its ID.
	// Returns an error if the entity does not exist or deletion fails.
	DeleteById(id int64) error
}
