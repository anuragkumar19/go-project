package middlewares

import (
	"context"
	"net/http"
	"strings"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/utils"
	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

func WithAuthGuard(handler func(user *database.GetUserByIdRow, c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := validateToken(c)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		handler(user, c)
	}
}

type MaybeUser struct {
	Valid bool
	User  *database.GetUserByIdRow
}

func WithMaybeUser(handler func(u MaybeUser, c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := validateToken(c)

		if !ok {
			handler(MaybeUser{
				Valid: false,
				User:  nil,
			}, c)
			return
		}

		handler(MaybeUser{
			Valid: true,
			User:  user,
		}, c)
	}
}

func validateToken(c *gin.Context) (*database.GetUserByIdRow, bool) {
	str := c.Request.Header.Get("Authorization")

	s := strings.Split(str, " ")

	if len(s) != 2 {
		return nil, false
	}

	if s[0] != "Bearer" {
		return nil, false
	}

	userId, err := utils.VerifyToken(s[1])

	if err != nil {
		return nil, false
	}

	users, err := db.GetUserById(context.Background(), int32(userId))

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return nil, false
	}

	return &users[0], true
}
