package routes

import (
	"example.com/go-htmx/handlers"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/me", middlewares.WithAuthGuard(handlers.Me))
	user.PUT("/name", middlewares.WithAuthGuard(handlers.UpdateName))
	user.PUT("/username", middlewares.WithAuthGuard(handlers.UpdateUsername))
	user.PUT("/password", middlewares.WithAuthGuard(handlers.UpdatePassword))
	user.PUT("/avatar", middlewares.WithAuthGuard(handlers.UpdateAvatar))
}
