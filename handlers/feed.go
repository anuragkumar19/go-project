package handlers

import (
	"context"
	"net/http"
	"strconv"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/middlewares"
	"github.com/gin-gonic/gin"
)

func GetMyFeed(user *database.GetUserByIdRow, c *gin.Context) {
	ids, err := db.GetJoinedSubreddit(context.Background(), user.ID)

	if err != nil {
		panic(err)
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	posts, err := db.GetFeedPosts(context.Background(), database.GetFeedPostsParams{
		Column1: ids,
		Limit:   int32(limit),
		Offset:  (int32(page) - 1) * int32(limit),
		UserID:  user.ID,
	})

	if err != nil {
		panic(err)
	}

	if len(posts) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"posts": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetTrendingPosts(user middlewares.MaybeUser, c *gin.Context) {
	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	posts, err := db.GetTrendingPostsPublic(context.Background(), userId)

	if err != nil {
		panic(err)
	}

	if len(posts) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"posts": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetTopSubreddit(user middlewares.MaybeUser, c *gin.Context) {
	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	items, err := db.GetTopSubredditPublic(context.Background(), userId)

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"items": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}
