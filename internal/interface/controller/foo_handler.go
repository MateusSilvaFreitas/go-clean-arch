package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/gateway"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase/dto"
	"github.com/gin-gonic/gin"
)

type (
	// FooHandler defines the methods that the FooController should implement
	FooHandler interface {
		CreateFoo(ctx *gin.Context)     // Method to handle the creation of Foo
		FindFooByID(ctx *gin.Context)   // Method to handle the retrieval of Foo by ID
		DeleteFooByID(ctx *gin.Context) // Method to handle the deletion of Foo by ID
	}

	// fooHandlerImpl implements the FooHandler interface
	fooHandlerImpl struct {
		CreateFooUseCase   usecase.CreateFooUseCase
		FindFooByIDUseCase usecase.FindFooByIDUseCase
		DeleteFooUseCase   usecase.DeleteFooUseCase
	}
)

// NewFooHandler is a constructor function that returns an instance of FooHandler with the required use cases
func NewFooHandler(cf usecase.CreateFooUseCase,
	ff usecase.FindFooByIDUseCase,
	df usecase.DeleteFooUseCase,
) FooHandler {
	return &fooHandlerImpl{
		CreateFooUseCase:   cf,
		FindFooByIDUseCase: ff,
		DeleteFooUseCase:   df,
	}
}

// CreateFoo handles the HTTP request to create a Foo
// It binds the input JSON to the CreateFooInputDTO, calls the use case to create Foo,
// and returns a JSON response with the created Foo or an error message.
func (f *fooHandlerImpl) CreateFoo(ctx *gin.Context) {
	var fooInput dto.CreateFooInputDTO

	// Bind the JSON body to the CreateFooInputDTO struct
	if err := ctx.ShouldBindJSON(&fooInput); err != nil {
		convertError(http.StatusBadRequest, err, ctx) // Return BadRequest if binding fails
		return
	}

	// Execute the use case to create Foo
	fooOutput, err := f.CreateFooUseCase.Execute(fooInput)

	// Handle errors from the use case execution
	if err != nil {
		convertError(http.StatusInternalServerError, err, ctx) // Return InternalServerError on failure
		return
	}

	// Return the created Foo with status 201 Created
	ctx.IndentedJSON(http.StatusCreated, fooOutput)
}

// FindFooByID handles the HTTP request to find a Foo by its ID
// It parses the ID from the URL parameters, calls the use case to find the Foo,
// and returns the Foo or an error message (404 if not found).
func (f *fooHandlerImpl) FindFooByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)

	// If the ID is invalid, return a BadRequest error
	if err != nil {
		convertError(http.StatusBadRequest, err, ctx)
		return
	}

	// Execute the use case to find the Foo by ID
	fooOutput, err := f.FindFooByIDUseCase.Execute(id)

	// Handle errors from the use case execution
	if err != nil {
		// Return 404 if the Foo is not found
		if errors.Is(err, gateway.ErrorFooNotFound) {
			convertError(http.StatusNotFound, err, ctx)
			return
		}
		convertError(http.StatusInternalServerError, err, ctx) // Return InternalServerError for other errors
		return
	}

	// Return the found Foo with status 200 OK
	ctx.IndentedJSON(http.StatusOK, fooOutput)
}

// DeleteFooByID handles the HTTP request to delete a Foo by its ID
// It parses the ID from the URL parameters, calls the use case to delete the Foo,
// and returns a 204 No Content response or an error message.
func (f *fooHandlerImpl) DeleteFooByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 0)

	// If the ID is invalid, return a BadRequest error
	if err != nil {
		convertError(http.StatusBadRequest, err, ctx)
		return
	}

	// Execute the use case to delete the Foo by ID
	err = f.DeleteFooUseCase.Execute(id)

	// Handle errors from the use case execution
	if err != nil {
		convertError(http.StatusInternalServerError, err, ctx) // Return InternalServerError on failure
		return
	}

	// Return status 204 No Content if the deletion is successful
	ctx.Status(http.StatusNoContent)
}

// convertError is a helper function that formats error messages and returns them as JSON responses
// It takes the HTTP status code, the error, and the context, and sends a structured JSON response.
func convertError(status int, err error, ctx *gin.Context) {
	ctx.IndentedJSON(status, gin.H{"message": err.Error()})
}
