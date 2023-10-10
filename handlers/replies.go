package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

func GetReply(c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return
	}

	replies, err := db.GetReplyByIdPublic(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(replies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return
	}

	reply := replies[0]
	c.JSON(http.StatusOK, gin.H{"reply": reply})
}

func GetReplyReplies(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func ReplyToPost(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.ReplyParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	post, ok := verifyPostCreator(user, c, false)

	if !ok {
		return
	}

	replies, err := db.CreateReply(context.Background(), database.CreateReplyParams{
		CreatorID: user.ID,
		PostID: sql.NullInt32{
			Valid: true,
			Int32: post.ID,
		},
		ParentReplyID: sql.NullInt32{
			Valid: false,
		},
		Content: body.Content,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": replies[0],
	})
}

func ReplyToReply(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.ReplyParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	reply, ok := verifyReplyCreator(user, c, false)

	if !ok {
		return
	}

	replies, err := db.CreateReply(context.Background(), database.CreateReplyParams{
		CreatorID: user.ID,
		PostID: sql.NullInt32{
			Valid: false,
		},
		ParentReplyID: sql.NullInt32{
			Valid: true,
			Int32: reply.ID,
		},
		Content: body.Content,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": replies[0],
	})
}

func VoteReply(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.VotePostParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	reply, ok := verifyReplyCreator(user, c, false)

	if !ok {
		return
	}

	err := db.VoteReply(context.Background(), database.VoteReplyParams{
		ReplyID: reply.ID,
		UserID:  user.ID,
		Down:    body.Down,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"down": body.Down,
	})
}

func RemoveReplyVote(user *database.GetUserByIdRow, c *gin.Context) {
	reply, ok := verifyReplyCreator(user, c, false)

	if !ok {
		return
	}

	items, err := db.GetReplyVote(context.Background(), database.GetReplyVoteParams{
		ReplyID: reply.ID,
		UserID:  user.ID,
	})

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No vote found to remove",
		})
	}

	err = db.RemoveReplyVote(context.Background(), database.RemoveReplyVoteParams{
		ReplyID: reply.ID,
		UserID:  user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vote removed",
	})
}

// TODO:...
func DeleteReply(user *database.GetUserByIdRow, c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func verifyReplyCreator(user *database.GetUserByIdRow, c *gin.Context, checkCreator bool) (*database.FindReplyByIdRow, bool) {
	s, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return nil, false
	}

	id, err := strconv.Atoi(s)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return nil, false
	}

	items, err := db.FindReplyById(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return nil, false
	}

	reply := items[0]
	if checkCreator && reply.CreatorID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Reply not found",
		})
		return nil, false
	}

	return &reply, true
}
