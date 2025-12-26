package router

import (
	"net/http"

	"github.com/Kevinmajesta/backend_library/internal/http/handler"
	"github.com/Kevinmajesta/backend_library/pkg/route"
)

func PublicRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
	}
}

func PrivateRoutes() []*route.Route {
	return []*route.Route{}
}
