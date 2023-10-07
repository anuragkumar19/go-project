package routes

import (
	"example.com/go-htmx/handlers"
	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", handlers.RegisterUser)
}
