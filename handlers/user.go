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

func Me(user *database.GetUserByIdRow, c *gin.Context) {
	c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	users, err := db.GetUserByIDPublic(context.Background(), int32(id))

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	user := users[0]
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUserByUsername(c *gin.Context) {
	username, ok := c.Params.Get("username")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	users, err := db.GetUserByUsernamePublic(context.Background(), username)

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	user := users[0]
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUserPosts(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
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

	posts, err := db.GetPostsOfUser(context.Background(), database.GetPostsOfUserParams{
		CreatorID: int32(id),
		Limit:     int32(limit),
		Offset:    (int32(page) - 1) * int32(limit),
		UserID:    userId,
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

func GetUserReplies(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
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

	replies, err := db.GetUserReplyPublic(context.Background(), database.GetUserReplyPublicParams{
		CreatorID: int32(id),
		Limit:     int32(limit),
		Offset:    (int32(page) - 1) * int32(limit),
		UserID:    userId,
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

func GetJoinedSubreddit(user middlewares.MaybeUser, c *gin.Context) {
	str, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
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

	items, err := db.GetJoinedSubreddit(context.Background(), int32(id))

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

func SearchUser(c *gin.Context) {
	q := c.Query("q")

	if q == "" {
		c.JSON(http.StatusOK, gin.H{
			"items": []string{},
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

	users, err := db.SearchUserPublic(context.Background(), database.SearchUserPublicParams{
		Username: "%" + strings.ToLower(q) + "%",
		Limit:    int32(limit),
		Offset:   (int32(page) - 1) * int32(limit),
	})

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"users": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UpdateName(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.UpdateNameParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	err := db.UpdateName(context.Background(), database.UpdateNameParams{
		ID:   user.ID,
		Name: body.Name,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Name updated",
	})
}

func UpdateUsername(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.UpdateUsernameParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	err := db.UpdateUsername(context.Background(), database.UpdateUsernameParams{
		ID:       user.ID,
		Username: body.Username,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Username updated",
	})
}

func UpdatePassword(user *database.GetUserByIdRow, c *gin.Context) {
	body := &validations.UpdatePasswordParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	users, err := db.LoginQuery(context.Background(), user.Email)

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		panic("user len cannot be zero")
	}

	if !utils.CheckPasswordHash(body.OldPassword, users[0].Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Old password did not match",
		})

		return
	}

	passwordHash, err := utils.HashPassword(body.NewPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password too long",
		})
		return
	}

	err = db.UpdatePassword(context.Background(), database.UpdatePasswordParams{
		ID:       user.ID,
		Password: passwordHash,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated",
	})
}

func UpdateAvatar(user *database.GetUserByIdRow, c *gin.Context) {
	path, ok := utils.UploadFile(c, "image")

	if !ok {
		return
	}

	err := db.UpdateAvatar(context.Background(), database.UpdateAvatarParams{
		ID:     user.ID,
		Avatar: path,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Uploaded",
	})
}
