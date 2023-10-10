package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/middlewares"
	"example.com/go-htmx/utils"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

func SearchSubreddit(user middlewares.MaybeUser, c *gin.Context) {
	q := c.Query("q")

	if q == "" {
		c.JSON(http.StatusOK, gin.H{
			"items": []string{},
		})
		return
	}

	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	items, err := db.SearchSubredditPublic(context.Background(), database.SearchSubredditPublicParams{
		Name:   "%" + strings.ToLower(q) + "%",
		Limit:  int32(limit),
		Offset: (int32(page) - 1) * int32(limit),
		UserID: userId,
	})

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

func GetSubredditByID(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	items, err := db.FindSubredditByIDPublic(context.Background(), database.FindSubredditByIDPublicParams{
		ID:     int32(id),
		UserID: userId,
	})

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
	c.JSON(http.StatusOK, gin.H{"subreddit": subreddit})
}

func GetSubredditByName(user middlewares.MaybeUser, c *gin.Context) {
	name, ok := c.Params.Get("name")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	var userId int32

	if user.User != nil {
		userId = user.User.ID
	}

	items, err := db.FindSubredditByNamePublic(context.Background(), database.FindSubredditByNamePublicParams{
		Name:   name,
		UserID: userId,
	})

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
	c.JSON(http.StatusOK, gin.H{"subreddit": subreddit})
}

func GetSubredditPosts(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
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

	posts, err := db.GetSubredditPosts(context.Background(), database.GetSubredditPostsParams{
		SubredditID: int32(id),
		Limit:       int32(limit),
		Offset:      (int32(page) - 1) * int32(limit),
		UserID:      userId,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func CreateSubreddit(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.CreateSubredditParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	items, err := db.FindSubredditByName(context.Background(), body.Name)

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

	subreddit, ok := verifySubredditCreator(user, c, true)

	if !ok {
		return
	}

	err := db.UpdateSubredditTitle(context.Background(), database.UpdateSubredditTitleParams{
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

func UpdateSubredditAbout(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.UpdateSubredditAboutParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	subreddit, ok := verifySubredditCreator(user, c, true)

	if !ok {
		return
	}

	err := db.UpdateSubredditAbout(context.Background(), database.UpdateSubredditAboutParams{
		ID:    subreddit.ID,
		About: body.About,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "About updated",
	})
}

func UpdateSubredditAvatar(user *database.GetUserByIdRow, c *gin.Context) {
	subreddit, ok := verifySubredditCreator(user, c, true)

	if !ok {
		return
	}

	path, ok := utils.UploadFile(c, "image")

	if !ok {
		return
	}

	err := db.UpdateSubredditAvatar(context.Background(), database.UpdateSubredditAvatarParams{
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

func UpdateSubredditCover(user *database.GetUserByIdRow, c *gin.Context) {
	subreddit, ok := verifySubredditCreator(user, c, true)

	if !ok {
		return
	}

	path, ok := utils.UploadFile(c, "image")

	if !ok {
		return
	}

	err := db.UpdateSubredditCover(context.Background(), database.UpdateSubredditCoverParams{
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

func verifySubredditCreator(user *database.GetUserByIdRow, c *gin.Context, checkCreator bool) (*database.FindSubredditByIdRow, bool) {
	s, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return nil, false
	}

	id, err := strconv.Atoi(s)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return nil, false
	}

	items, err := db.FindSubredditById(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return nil, false
	}

	subreddit := items[0]
	if checkCreator && subreddit.CreatorID != user.ID {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Subreddit not found",
		})
		return nil, false
	}

	return &subreddit, true
}

func JoinSubreddit(user *database.GetUserByIdRow, c *gin.Context) {
	subreddit, ok := verifySubredditCreator(user, c, false)

	if !ok {
		return
	}

	items, err := db.IsAlreadyJoined(context.Background(), database.IsAlreadyJoinedParams{
		UserID:      user.ID,
		SubredditID: subreddit.ID,
	})

	if err != nil {
		panic(err)
	}

	if len(items) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Already Joined",
		})
		return
	}

	_, err = db.JoinSubreddit(context.Background(), database.JoinSubredditParams{
		UserID:      user.ID,
		SubredditID: subreddit.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Joined",
	})
}

func LeaveSubreddit(user *database.GetUserByIdRow, c *gin.Context) {
	subreddit, ok := verifySubredditCreator(user, c, false)

	if !ok {
		return
	}

	items, err := db.IsAlreadyJoined(context.Background(), database.IsAlreadyJoinedParams{
		UserID:      user.ID,
		SubredditID: subreddit.ID,
	})

	if err != nil {
		panic(err)
	}

	if len(items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You need to join first",
		})
		return
	}

	err = db.LeaveSubreddit(context.Background(), database.LeaveSubredditParams{
		UserID:      user.ID,
		SubredditID: subreddit.ID,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Left",
	})
}

func DeleteSubreddit(user *database.GetUserByIdRow, c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
