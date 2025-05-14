// Package usecase contains the business logic related to the Foo domain.
// It defines use case interfaces and their implementations, responsible
// for coordinating the creation, validation, and persistence of Foo entities.
package usecase

import (
	"fmt"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase/dto"
)

type (
	// CreateFooUseCase defines the contract for creating a new Foo entity.
	// It contains the Execute method, which receives input data, performs validation,
	// persists the entity, and returns the created Foo or an error.
	CreateFooUseCase interface {
		Execute(dto.CreateFooInputDTO) (dto.FooOutput, error)
	}

	// createFooUseCaseImpl is the concrete implementation of the CreateFooUseCase interface.
	// It depends on a FooRepository to persist the Foo entity and encapsulates
	// the business logic required to create a new Foo.
	createFooUseCaseImpl struct {
		repo gateway.FooRepository
	}
)

// NewCreateFoo creates a new instance of CreateFooUseCase.
// It receives a FooRepository implementation to handle persistence logic.
func NewCreateFoo(repo gateway.FooRepository) CreateFooUseCase {
	return &createFooUseCaseImpl{
		repo: repo,
	}
}

// Execute handles the process of creating a new Foo entity.
// It validates the input data, constructs the entity, persists it,
// and returns a DTO with the created data or an error if the operation fails.
func (c *createFooUseCaseImpl) Execute(input dto.CreateFooInputDTO) (dto.FooOutput, error) {
	foo := entity.Foo{
		Name: input.Name,
		Type: input.Type,
	}

	if err := foo.Validate(); err != nil {
		return dto.FooOutput{}, fmt.Errorf("validation failed: %w", err)
	}

	savedFoo, err := c.repo.Save(foo)

	if err != nil {
		return dto.FooOutput{}, fmt.Errorf("failed to save %w", err)
	}

	return dto.FooOutput{
		ID:        fmt.Sprintf("%d", savedFoo.ID),
		Name:      savedFoo.Name,
		Type:      savedFoo.Type,
		CreatedAt: savedFoo.CreatedAt,
	}, nil
}
