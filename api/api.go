package api

import (
	"oghenekparobor/market-lens/routes"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	// addr string
	// db   *sql.DB
}

func MarketLensApiServer() *APIServer {
	return &APIServer{
		// addr: addr,
		// db:   db,
	}
}

func (s *APIServer) Run() error {
	// entry point of the app using gin
	r := gin.Default()

	// v1 apis
	api := r.Group("/api/v1")

	userRouteHandler := routes.UserRouteHandler()
	userRouteHandler.RegisterUserRoutes(api)

	err := r.Run() // runs over port 8080

	return err
}
