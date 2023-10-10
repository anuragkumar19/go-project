package routes

import (
	"example.com/go-htmx/handlers"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func postRouter(r *gin.RouterGroup) {
	posts := r.Group("/posts")

	posts.GET("/:id", handlers.GetPost)
	posts.GET("/:id/replies", handlers.GetPostReplies)
	posts.POST("/:id/vote", middlewares.WithAuthGuard(handlers.VotePost))
	posts.DELETE("/:id/vote", middlewares.WithAuthGuard(handlers.RemovePostVote))
	posts.POST("/:id/reply", middlewares.WithAuthGuard(handlers.ReplyToPost))
	posts.DELETE("/:id", middlewares.WithAuthGuard(handlers.DeletePost))
}
