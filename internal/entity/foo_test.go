package entity_test

import (
	"testing"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/entity"
)

func TestFooValidation(t *testing.T) {
	foo := []struct {
		name string
		input entity.Foo
		hasError bool
	}{
		{name: "valid name and type", input: entity.Foo{Name: "Test 1", Type: "Type A"}, hasError: false},
		{name: "invalid name and valid type", input: entity.Foo{Name: "", Type: "Type A"}, hasError: true},
		{name: "valid name and invalid type", input: entity.Foo{Name: "Test 2", Type: ""}, hasError: true},
		{name: "invalid name and type", input: entity.Foo{Name: "", Type: ""}, hasError: true},
		{name: "valid name an type empty with spaces", input: entity.Foo{Name: "  ", Type: "  "}, hasError: false},
	}

	for _, testCase := range foo {
		err := testCase.input.Validate()

		if testCase.hasError && err == nil {
			t.Errorf("Test case \"%v\": expected error but got nil", testCase.name)
		}
		if !testCase.hasError && err != nil {
			t.Errorf("Test case \"%v\": did not expect error but got: %v", testCase.name, err)
		}
		

	}
}