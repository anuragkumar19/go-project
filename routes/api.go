package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	api := r.Group("/api")

	authRouter(api)
	userRouter(api)
	subredditRouter(api)
	postRouter(api)
	repliesRouter(api)
}
