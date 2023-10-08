package routes

import (
	database "example.com/go-htmx/db"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/protected", middlewares.WithAuthGuard(func(user *database.GetUserByIdRow, c *gin.Context) {
		c.JSON(200, user)
	}))

	authRouter(api)
}
