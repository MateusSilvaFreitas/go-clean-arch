package main

import (
	"log"

	"github.com/MateusSilvaFreitas/go-clean-arch/internal/infra/dependencies"
	"github.com/MateusSilvaFreitas/go-clean-arch/internal/infra/router"
	"github.com/gin-gonic/gin"
)

func main() {
	appDependencies, err := dependencies.NewAppDependencies()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	router.CreateRouter(appDependencies, r)

	r.Run(":8080")
}
