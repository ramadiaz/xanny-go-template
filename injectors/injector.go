// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	exampleControllers "xanny-go-template/api/example/controllers"
	exampleRepositories "xanny-go-template/api/example/repositories"
	exampleServices "xanny-go-template/api/example/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var exampleFeatureSet = wire.NewSet(
	exampleRepositories.NewComponentRepository,
	exampleServices.NewComponentServices,
	exampleControllers.NewCompController,
)

func InitializeExampleController(db *gorm.DB, validate *validator.Validate) exampleControllers.CompControllers {
	wire.Build(exampleFeatureSet)
	return nil
}
