package dto

// CreateFooInputDTO represents the input data transfer object (DTO)
// used when creating a new Foo entity.
//
// This struct is used to capture the data sent by the client
// when making a request to create a new Foo. It contains
// the necessary fields for the creation process.
type CreateFooInputDTO struct {
	// Name represents the name of the Foo entity.
	// This field is required for creating a new Foo.
	Name string `json:"name"` // JSON tag specifies the key for the field in the request body

	// Type represents the type of the Foo entity.
	// This field is required for creating a new Foo.
	Type string `json:"type"` // JSON tag specifies the key for the field in the request body
}
