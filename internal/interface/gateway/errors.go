package gateway

import "errors"

// ErrorFooNotFound is a predefined error used to indicate that
// a Foo entity was not found in the data store.
//
// This error is typically returned by repository implementations when
// a lookup by ID or other criteria does not yield any results.
var ErrorFooNotFound = errors.New("foo not found")
