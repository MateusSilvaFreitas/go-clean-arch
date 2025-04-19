package usecase

import (
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase/dto"
)

type (
	CreateFooUseCase interface {
		Execute(dto.CreateFooInputDTO) (dto.FooOutput, error)
	}

	createFooUseCaseImpl struct {
		repo *gateway.FooRepository
	}
)

func newCreateFoo(repo *gateway.FooRepository) CreateFooUseCase {
	return &createFooUseCaseImpl{
		repo: repo,
	}
}

func (c *createFooUseCaseImpl) Execute(input dto.CreateFooInputDTO) (dto.FooOutput, error) {
	return dto.FooOutput{}, nil
}
