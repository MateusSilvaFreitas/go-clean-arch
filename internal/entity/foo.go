// Package entity defines the Foo entity and its associated methods.
package entity

import (
	"fmt"
	"strings"
	"time"
)

// Foo represents an entity in the system with an ID, Name, Type, and CreatedAt timestamp.
type Foo struct {
	ID        int64     // ID is the unique identifier for the Foo entity.
	Name      string    // Name is the name associated with the Foo entity.
	Type      string    // Type is the type associated with the Foo entity.
	CreatedAt time.Time // CreatedAt is the timestamp when the Foo entity was created.
}

// Validate checks if the Foo entity has valid data.
// It ensures that the Name and Type fields are not empty.
// If any field is invalid, an error is returned with a list of validation messages.
func (f *Foo) Validate() error {
	var errors []string

	// Check if Name is empty
	if f.Name == "" {
		errors = append(errors, "Invalid name")
	}

	// Check if Type is empty
	if f.Type == "" {
		errors = append(errors, "Invalid type")
	}

	// If there are validation errors, return them as a formatted error message
	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "; "))
	}

	// If no validation errors, return nil
	return nil
}
