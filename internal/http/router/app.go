package router

import (
	"net/http"

	"github.com/Kevinmajesta/backend_library/internal/http/handler"
	"github.com/Kevinmajesta/backend_library/pkg/route"
)

func PublicRoutes(userHandler handler.UserHandler, bookHandler handler.BookHandler,
	borrowHandler handler.BorrowHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/books",
			Handler: bookHandler.CreateBook,
		},
		{
			Method:  http.MethodGet,
			Path:    "/books/:book_id",
			Handler: bookHandler.FindBook,
		},
		{
			Method:  http.MethodPost,
			Path:    "/borrow",
			Handler: borrowHandler.BorrowBook,
		},
		{
			Method:  http.MethodPost,
			Path:    "/return/:borrow_id",
			Handler: borrowHandler.ReturnBook,
		},
	}
}

func PrivateRoutes() []*route.Route {
	return []*route.Route{}
}
