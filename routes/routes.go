package routes

import (
	"oghenekparobor/market-lens/controllers"

	"github.com/gin-gonic/gin"
)

// todo: remove this and call register directly
type Handler struct {
}

func UserRouteHandler() *Handler {
	return &Handler{}
}

func (*Handler) RegisterUserRoutes(r *gin.RouterGroup) {
	r.POST("/login", controllers.LoginUser)
	r.POST("/register", controllers.RegisterUser)
}
