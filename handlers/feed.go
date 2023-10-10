package handlers

import (
	"net/http"

	database "example.com/go-htmx/db"
	"github.com/gin-gonic/gin"
)

func GetMyFeed(user *database.GetUserByIdRow, c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
