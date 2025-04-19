package dto

import "time"

// FooOutput represents the data transfer object (DTO)
// used to send the output data for a Foo entity.
//
// This struct is used when sending the details of a Foo
// entity to the client, typically in response to a successful
// operation like creating, retrieving, or updating a Foo.
type FooOutput struct {
	// ID represents the unique identifier of the Foo entity.
	// It is returned to the client after performing operations like
	// creating or retrieving a Foo.
	ID string `json:"id"` // JSON tag specifies the key for the field in the response body.

	// Name represents the name of the Foo entity.
	// This field is part of the output to show the Foo's name.
	Name string `json:"name"` // JSON tag specifies the key for the field in the response body.

	// Type represents the type of the Foo entity.
	// This field is part of the output to show the Foo's type.
	Type string `json:"type"` // JSON tag specifies the key for the field in the response body.

	// CreatedAt represents the timestamp when the Foo entity was created.
	// This field is included in the output to show the creation time of the Foo.
	CreatedAt time.Time `json:"created_at"` // JSON tag specifies the key for the field in the response body.
}
