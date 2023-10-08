package handlers

import (
	"context"
	"net/http"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/utils"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

func Me(user *database.GetUserByIdRow, c *gin.Context) {
	c.JSON(http.StatusOK, user)
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
