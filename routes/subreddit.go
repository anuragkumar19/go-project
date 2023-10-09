package routes

import (
	"example.com/go-htmx/handlers"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func subredditRouter(r *gin.RouterGroup) {
	subreddit := r.Group("/r")

	subreddit.POST("/", middlewares.WithAuthGuard(handlers.CreateSubreddit))
	subreddit.PUT("/:id/title", middlewares.WithAuthGuard(handlers.UpdateSubredditTitle))
	subreddit.PUT("/:id/avatar", middlewares.WithAuthGuard(handlers.UpdateSubredditAvatar))
	subreddit.PUT("/:id/cover", middlewares.WithAuthGuard(handlers.UpdateSubredditCover))
	subreddit.POST("/:id/join", middlewares.WithAuthGuard(handlers.JoinSubreddit))
	subreddit.POST("/:id/leave", middlewares.WithAuthGuard(handlers.LeaveSubreddit))
}
