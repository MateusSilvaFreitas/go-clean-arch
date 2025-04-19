package entity

import (
	"fmt"
	"strings"
	"time"
)

type Foo struct {
	ID        int64
	Name      string
	Type      string
	CreatedAt time.Time
}

func (f *Foo) Validate() error {
	var errors []string

	if f.Name == "" {
		errors = append(errors, "Invalid name")
	}

	if f.Type == "" {
		errors = append(errors, "Invalid type")
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "; "))
	}

	return nil
}
