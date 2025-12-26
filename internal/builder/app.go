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

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	borrowRepository := repository.NewBorrowRepository(db)
	borrowService := service.NewBorrowService(db, bookRepository, borrowRepository)
	borrowHandler := handler.NewBorrowHandler(borrowService)

	return router.PublicRoutes(userHandler, bookHandler, borrowHandler)
}

func BuildPrivateRoutes(db *gorm.DB) []*route.Route {

	return router.PrivateRoutes()
}
