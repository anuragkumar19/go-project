package routes

import (
	"example.com/go-htmx/handlers"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func feedRouter(r *gin.RouterGroup) {
	feed := r.Group("/feed")

	feed.GET("/", middlewares.WithAuthGuard(handlers.GetMyFeed))
	feed.GET("/r/top", middlewares.WithMaybeUser(handlers.GetTopSubreddit))
	feed.GET("/posts/trending", middlewares.WithMaybeUser(handlers.GetTrendingPosts))
}
