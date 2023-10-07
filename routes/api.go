package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	api := r.Group("/api")

	authRouter(api)
}
