package router

import (
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/infra/dependencies"
	"github.com/gin-gonic/gin"
)

// CreateRouter registers all application routes and binds them to their respective handlers.
// It organizes the routes using Gin's route grouping for better readability and scalability.
//
// Parameters:
//   - appDependencies: pointer to AppDependencies, which contains the handler instances.
//   - router: the Gin engine instance used to define the HTTP routes.
func CreateRouter(appDependencies *dependencies.AppDependencies, router *gin.Engine) {
	fooRoutes := router.Group("/foo")
	{
		fooRoutes.POST("", appDependencies.FooHandler.CreateFoo)           // Creates a new Foo
		fooRoutes.GET("/:id", appDependencies.FooHandler.FindFooByID)      // Retrieves a Foo by ID
		fooRoutes.DELETE("/:id", appDependencies.FooHandler.DeleteFooByID) // Deletes a Foo by ID
	}
}
