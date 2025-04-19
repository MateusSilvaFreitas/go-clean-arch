package usecase

import (
	"fmt"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase/dto"
)

type (
	// FindFooByIdUseCase defines the contract for retrieving a Foo entity by its ID.
	FindFooByIDUseCase interface {
		Execute(id int64) (dto.FooOutput, error)
	}

	// findFooByIdUseCaseImpl is the concrete implementation of FindFooByIdUseCase.
	findFooByIDUseCaseImpl struct {
		repo gateway.FooRepository
	}
)

// NewFindFooById returns a new instance of FindFooByIdUseCase.
func NewFindFooById(repo gateway.FooRepository) FindFooByIDUseCase {
	return &findFooByIDUseCaseImpl{
		repo: repo,
	}
}

// Execute retrieves a Foo by its ID and returns it as a DTO.
// If the Foo is not found or an error occurs, it returns an error.
func (f *findFooByIDUseCaseImpl) Execute(id int64) (dto.FooOutput, error) {
	foo, err := f.repo.GetById(id)

	if err != nil {
		return dto.FooOutput{}, fmt.Errorf("failed to get Foo by id %d: %w", id, err)
	}

	return dto.FooOutput{
		Name:      foo.Name,
		Type:      foo.Type,
		CreatedAt: foo.CreatedAt,
	}, nil
}
