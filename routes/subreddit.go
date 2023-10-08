package routes

import (
	"github.com/gin-gonic/gin"
)

func subredditRouter(r *gin.RouterGroup) {
	subreddit := r.Group("/r")

	subreddit.GET("/")
}
