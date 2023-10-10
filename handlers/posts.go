package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/middlewares"
	"example.com/go-htmx/utils"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

func GetPost(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	posts, err := db.GetPostByIDPublic(context.Background(), database.GetPostByIDPublicParams{
		ID:     int32(id),
		UserID: userId,
	})

	if err != nil {
		panic(err)
	}

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	post := posts[0]
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func GetPostReplies(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	replies, err := db.GetPostReplyPublic(context.Background(), database.GetPostReplyPublicParams{
		PostID: sql.NullInt32{
			Valid: true,
			Int32: int32(id),
		},
		Limit:  int32(limit),
		Offset: (int32(page) - 1) * int32(limit),
		UserID: userId,
	})

	if err != nil {
		panic(err)
	}

	if len(replies) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"replies": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"replies": replies,
	})
}

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

	if valid := validations.ValidateForm(c, body); !valid {
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

	if valid := validations.ValidateForm(c, body); !valid {
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

func DeletePost(user *database.GetUserByIdRow, c *gin.Context) {
	post, ok := verifyPostCreator(user, c, true)

	if !ok {
		return
	}

	err := db.DeletePost(context.Background(), post.ID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post removed",
	})
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
