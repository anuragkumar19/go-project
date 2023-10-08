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

func CreateSubreddit(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.CreateSubredditParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	items, err := db.FindSubreddit(context.Background(), body.Name)

	if err != nil {
		panic(err)
	}

	if len(items) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Subreddit not available",
		})
		return
	}

	s, err := db.CreateSubreddit(context.Background(), database.CreateSubredditParams{
		Name:      body.Name,
		CreatorID: user.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": s[0].ID,
	})
}

func UpdateSubredditTitle(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.UpdateSubredditTitleParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	s, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	id, err := strconv.Atoi(s)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	items, err := db.FindSubredditById(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	subreddit := items[0]
	if subreddit.CreatorID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	err = db.UpdateSubredditTitle(context.Background(), database.UpdateSubredditTitleParams{
		ID:    subreddit.ID,
		Title: body.Title,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Title updated",
	})
}

func UpdateSubredditAvatar(user *database.GetUserByIdRow, c *gin.Context) {
	s, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	id, err := strconv.Atoi(s)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	items, err := db.FindSubredditById(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	subreddit := items[0]
	if subreddit.CreatorID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	path, ok := utils.UploadFile(c, "image")

	if !ok {
		return
	}

	err = db.UpdateSubredditAvatar(context.Background(), database.UpdateSubredditAvatarParams{
		ID:     subreddit.ID,
		Avatar: path,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Avatar updated",
	})

}

// TODO: REFACTOR UPLOAD IN ANOTHER FUNCTION
func UpdateSubredditCover(user *database.GetUserByIdRow, c *gin.Context) {
	s, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	id, err := strconv.Atoi(s)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	items, err := db.FindSubredditById(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	subreddit := items[0]
	if subreddit.CreatorID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	path, ok := utils.UploadFile(c, "image")

	if !ok {
		return
	}

	err = db.UpdateSubredditCover(context.Background(), database.UpdateSubredditCoverParams{
		ID:    subreddit.ID,
		Cover: path,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cover updated",
	})

}
