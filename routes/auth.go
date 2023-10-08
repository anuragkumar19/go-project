package routes

import (
	"example.com/go-htmx/handlers"
	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/verify-email", handlers.VerifyEmail)
	auth.POST("/login", handlers.Login)
	auth.POST("/forgot-password", handlers.ForgotPassword)
	auth.POST("/reset-password", handlers.ResetPassword)
}
