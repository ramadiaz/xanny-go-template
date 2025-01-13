// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"xanny-go-template/api/example/controllers"
	"xanny-go-template/api/example/repositories"
	"xanny-go-template/api/example/services"
)

// Injectors from injector.go:

func InitializeExampleController(db *gorm.DB, validate *validator.Validate) controllers.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	compService := services.NewComponentServices(compRepositories, db, validate)
	compControllers := controllers.NewCompController(compService)
	return compControllers
}

// injector.go:

var exampleFeatureSet = wire.NewSet(repositories.NewComponentRepository, services.NewComponentServices, controllers.NewCompController)
