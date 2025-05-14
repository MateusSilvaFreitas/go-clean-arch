package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/mocks"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateFoo(t *testing.T) {
	type testCase struct {
		name     string
		input    dto.CreateFooInputDTO
		mockSaveFn  func(entity.Foo) (entity.Foo, error)
		hasError bool
	}

	testCases := []testCase{
		{
			name:  "should create a new Foo",
			input: dto.CreateFooInputDTO{Name: "Test 1", Type: "Type A"},
			mockSaveFn: func(f entity.Foo) (entity.Foo, error) {
				f.ID = 1
				f.CreatedAt = time.Now()
				return f, nil
			},
			hasError: false,
		},
		{
			name: "should return error for missing name",
			input: dto.CreateFooInputDTO{Name: "", Type: "Type A"},
			mockSaveFn: func(f entity.Foo) (entity.Foo, error) {
				return f, nil
			},
			hasError: true,
		},
		{
			name: "should return error for missing type",
			input: dto.CreateFooInputDTO{Name: "Test 2", Type: ""},
			mockSaveFn: func(f entity.Foo) (entity.Foo, error) {
				return f, nil
			},
			hasError: true,
		},
		{
			name: "should return error from repository",
			input: dto.CreateFooInputDTO{Name: "Test 3", Type: "Type B"},
			mockSaveFn: func(f entity.Foo) (entity.Foo, error) {
				return entity.Foo{}, fmt.Errorf("repository error")
			},
			hasError: true,
		},
	}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				mockRepo := &mocks.FooRepositoryMock{
					SaveFn: tc.mockSaveFn,
				}
				createFooUseCase := usecase.NewCreateFoo(mockRepo)
				output, err := createFooUseCase.Execute(tc.input)

				fmt.Printf("Output: %+v\n", output)

				if tc.hasError {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tc.input.Name, output.Name)
					assert.Equal(t, tc.input.Type, output.Type)
					assert.NotEmpty(t, output.ID)
					assert.NotEmpty(t, output.CreatedAt)
				}
			
			})
		}
}