package routers

import (
	"xanny-go-template/internal/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	internalController := injectors.InitializeAuthController(validate)

	AuthRoutes(r, internalController)
}
