package handlers

import (
	"context"
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/emails"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

func RegisterUser(c *gin.Context) {
	body := &validations.RegisterParameters{}

	valid := validations.Validate(c, body)

	if !valid {
		return
	}

	// check email
	users, err := db.FindUserByEmail(context.Background(), body.Email)

	if err != nil {
		panic(err)
	}

	if len(users) > 0 && users[0].IsEmailVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already registered",
		})
		return
	}

	checkUsername, err := db.FindUserByUsername(context.Background(), body.Username)

	if err != nil {
		panic(err)
	}

	if len(checkUsername) != 0 && (len(users) == 0 || checkUsername[0].Username != users[0].Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username taken",
		})
		return
	}

	otp := rand.Intn(899_999) + 100_000
	otpExpiry := time.Now().Add(time.Minute * 5)

	if len(users) != 0 {
		_, err = db.UpdateUser(context.Background(), database.UpdateUserParams{
			ID:       users[0].ID,
			Name:     body.Name,
			Email:    body.Email,
			Password: body.Password,
			Username: body.Username,
			Otp: sql.NullString{
				String: strconv.Itoa(otp),
				Valid:  true,
			},
			OtpExpiry: sql.NullTime{
				Time:  otpExpiry,
				Valid: true,
			},
		})

	} else {
		_, err = db.CreateUser(context.Background(), database.CreateUserParams{
			Name:     body.Name,
			Email:    body.Email,
			Password: body.Password,
			Username: body.Username,
			Otp: sql.NullString{
				String: strconv.Itoa(otp),
				Valid:  true,
			},
			OtpExpiry: sql.NullTime{
				Time:  otpExpiry,
				Valid: true,
			},
		})

	}

	if err != nil {
		panic(err)
	}

	err = emails.SendOTP(otp, body.Email)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered",
	})

}
