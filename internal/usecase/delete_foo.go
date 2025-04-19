package usecase

import (
	"fmt"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
)

type (
	// DeleteFooUseCase defines the contract for deleting a Foo entity by its ID.
	DeleteFooUseCase interface {
		Execute(id int64) error
	}

	// deleteFooUseCaseImpl is the implementation of DeleteFooUseCase.
	deleteFooUseCaseImpl struct {
		repo gateway.FooRepository
	}
)

// NewDeleteFoo creates a new instance of DeleteFooUseCase.
func NewDeleteFoo(repo gateway.FooRepository) DeleteFooUseCase {
	return &deleteFooUseCaseImpl{
		repo: repo,
	}
}

// Execute deletes a Foo entity by its ID.
// Returns an error if the operation fails.
func (d *deleteFooUseCaseImpl) Execute(id int64) error {
	if err := d.repo.DeleteById(id); err != nil {
		return fmt.Errorf("failed to delete Foo by id %d, %w", id, err)
	}

	return nil
}
