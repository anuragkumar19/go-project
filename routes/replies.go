package routes

import (
	"example.com/go-htmx/handlers"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func repliesRouter(r *gin.RouterGroup) {
	replies := r.Group("/replies")

	replies.POST("/:id/vote", middlewares.WithAuthGuard(handlers.VoteReply))
	replies.DELETE("/:id/vote", middlewares.WithAuthGuard(handlers.RemoveReplyVote))
	replies.POST("/:id/reply", middlewares.WithAuthGuard(handlers.ReplyToReply))
}
