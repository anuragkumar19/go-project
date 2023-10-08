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
		str := c.Request.Header.Get("Authorization")

		s := strings.Split(str, " ")

		if len(s) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if s[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		userId, err := utils.VerifyToken(s[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		users, err := db.GetUserById(context.Background(), int32(userId))

		if err != nil {
			panic(err)
		}

		if len(users) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		handler(&users[0], c)
	}
}
