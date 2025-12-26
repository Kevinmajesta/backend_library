package builder

import (
	"github.com/Kevinmajesta/backend_library/internal/http/handler"
	"github.com/Kevinmajesta/backend_library/internal/http/router"
	"github.com/Kevinmajesta/backend_library/internal/repository"
	"github.com/Kevinmajesta/backend_library/internal/service"
	"github.com/Kevinmajesta/backend_library/pkg/route"

	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	return router.PublicRoutes(userHandler)
}

func BuildPrivateRoutes(db *gorm.DB) []*route.Route {

	return router.PrivateRoutes()
}
