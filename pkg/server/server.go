package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Kevinmajesta/backend_library/pkg/response"
	"github.com/Kevinmajesta/backend_library/pkg/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(serverName string, publicRoutes, privateRoutes []*route.Route) *Server {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"}, // Ganti dengan URL frontend jika perlu, misalnya "http://localhost:3000"
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Hello, World!", nil))
	})

	v1 := e.Group(fmt.Sprintf("%s/api/v1", serverName))

	// Hanya rute privat yang memerlukan autentikasi
	if len(privateRoutes) > 0 {
		for _, v := range privateRoutes {
			v1.Add(v.Method, v.Path, v.Handler) // apply JWTProtection only for private routes
		}
	}

	// Rute publik tidak perlu autentikasi
	if len(publicRoutes) > 0 {
		for _, v := range publicRoutes {
			v1.Add(v.Method, v.Path, v.Handler) // Jangan ada JWTProtection di sini
		}
	}

	return &Server{e}
}

func (s *Server) Run() {
	runServer(s)
	gracefulShutdown(s)
}

func runServer(srv *Server) {
	go func() {
		err := srv.Start(":8080")
		log.Fatal(err)
	}()
}

func gracefulShutdown(srv *Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal("Server Shutdown:", err)
		}
	}()
}
