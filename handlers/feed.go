package handlers

import (
	"context"
	"net/http"
	"strconv"

	database "example.com/go-htmx/db"
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
