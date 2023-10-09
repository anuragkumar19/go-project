package handlers

import (
	"context"
	"net/http"
	"strconv"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/utils"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

func CreatePostWithText(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.CreatePostWithTextParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	subreddit, ok := verifySubredditCreator(user, c, false)

	if !ok {
		return
	}

	posts, err := db.CreatePost(context.Background(), database.CreatePostParams{
		Title:       body.Title,
		Text:        body.Text,
		SubredditID: subreddit.ID,
		CreatorID:   user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": posts[0],
	})
}

func CreatePostWithImage(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.CreatePostWithMediaParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	subreddit, ok := verifySubredditCreator(user, c, false)

	if !ok {
		return
	}

	path, ok := utils.UploadFile(c, "image")

	if !ok {
		return
	}

	posts, err := db.CreatePost(context.Background(), database.CreatePostParams{
		Title:       body.Title,
		Text:        body.Text,
		Image:       path,
		SubredditID: subreddit.ID,
		CreatorID:   user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": posts[0],
	})
}

func CreatePostWithVideo(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.CreatePostWithMediaParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	subreddit, ok := verifySubredditCreator(user, c, false)

	if !ok {
		return
	}

	path, ok := utils.UploadFile(c, "video")

	if !ok {
		return
	}

	posts, err := db.CreatePost(context.Background(), database.CreatePostParams{
		Title:       body.Title,
		Text:        body.Text,
		Video:       path,
		SubredditID: subreddit.ID,
		CreatorID:   user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": posts[0],
	})
}

func CreatePostWithLink(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.CreatePostWithLinkParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	subreddit, ok := verifySubredditCreator(user, c, false)

	if !ok {
		return
	}

	posts, err := db.CreatePost(context.Background(), database.CreatePostParams{
		Title:       body.Title,
		Text:        body.Text,
		Link:        body.Link,
		SubredditID: subreddit.ID,
		CreatorID:   user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": posts[0],
	})
}

func VotePost(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.VotePostParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	post, ok := verifyPostCreator(user, c, false)

	if !ok {
		return
	}

	err := db.VotePost(context.Background(), database.VotePostParams{
		PostID: post.ID,
		UserID: user.ID,
		Down:   body.Down,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"down": body.Down,
	})
}

func RemovePostVote(user *database.GetUserByIdRow, c *gin.Context) {
	post, ok := verifyPostCreator(user, c, false)

	if !ok {
		return
	}

	items, err := db.GetPostVote(context.Background(), database.GetPostVoteParams{
		PostID: post.ID,
		UserID: user.ID,
	})

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No vote found to remove",
		})
	}

	err = db.RemovePostVote(context.Background(), database.RemovePostVoteParams{
		PostID: post.ID,
		UserID: user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vote removed",
	})
}

// TODO: after implementation of comments and upVote & downVote
func DeletePost(user *database.GetUserByIdRow, c *gin.Context) {

}

func verifyPostCreator(user *database.GetUserByIdRow, c *gin.Context, checkCreator bool) (*database.FindPostByIdRow, bool) {
	s, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return nil, false
	}

	id, err := strconv.Atoi(s)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return nil, false
	}

	items, err := db.FindPostById(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return nil, false
	}

	post := items[0]
	if checkCreator && post.CreatorID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return nil, false
	}

	return &post, true
}
