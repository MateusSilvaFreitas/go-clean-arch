package dependencies

import (
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/infra/db"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/interface/controller"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/usecase"
)

// AppDependencies holds all application-wide dependencies that are initialized once and reused,
// such as handlers, repositories, and use cases. It serves as a centralized place to wire the
// components together, following the Dependency Injection pattern.
type AppDependencies struct {
	FooHandler controller.FooHandler
}

// NewAppDependencies initializes the application's dependencies and returns a pointer to
// AppDependencies. It connects to the MySQL database, sets up repositories, use cases,
// and controllers.
//
// Returns:
//   - *AppDependencies: a struct containing all initialized application dependencies.
//   - error: an error if any of the dependencies could not be created (e.g., DB connection failure).
func NewAppDependencies() (*AppDependencies, error) {
	mysqlDatabase, err := db.NewMySqlConnector()
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	fooRepo := db.NewMySqlFooRepository(mysqlDatabase)

	// Initialize use cases
	cf := usecase.NewCreateFoo(fooRepo)
	ff := usecase.NewFindFooById(fooRepo)
	df := usecase.NewDeleteFoo(fooRepo)

	// Initialize controllers/handlers
	fooHandler := controller.NewFooHandler(cf, ff, df)

	return &AppDependencies{FooHandler: fooHandler}, nil
}
