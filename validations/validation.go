package validations

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

func Validate[T interface{}](c *gin.Context, body *T) (valid bool) {
	conform := modifiers.New()
	validator := validator.New()

	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return false
	}

	if err := conform.Struct(context.Background(), body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return false
	}

	if err := validator.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return false
	}

	return true
}

func ValidateForm[T interface{}](c *gin.Context, body *T) (valid bool) {
	conform := modifiers.New()
	validator := validator.New()

	if err := c.Bind(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return false
	}

	if err := conform.Struct(context.Background(), body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return false
	}

	if err := validator.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return false
	}

	return true
}
