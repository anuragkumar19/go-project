package handlers

import (
	"context"
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	database "example.com/go-htmx/db"
	"example.com/go-htmx/emails"
	"example.com/go-htmx/utils"
	"example.com/go-htmx/validations"
	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

func RegisterUser(c *gin.Context) {
	body := &validations.RegisterParameters{}

	if valid := validations.Validate(c, body); !valid {
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

	if len(checkUsername) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username taken",
		})
		return
	}

	otp := rand.Intn(899_999) + 100_000
	otpExpiry := time.Now().Add(time.Minute * 5)

	passwordHash, err := utils.HashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password too long",
		})
		return
	}

	if len(users) != 0 {
		_, err = db.UpdateUser(context.Background(), database.UpdateUserParams{
			ID:       users[0].ID,
			Name:     body.Name,
			Email:    body.Email,
			Password: passwordHash,
			Username: body.Username,
			Otp: sql.NullInt32{
				Int32: int32(otp),
				Valid: true,
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
			Password: passwordHash,
			Username: body.Username,
			Otp: sql.NullInt32{
				Int32: int32(otp),
				Valid: true,
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

	go func() {
		err = emails.SendOTP(otp, body.Email)

		if err != nil {
			panic(err)
		}
	}()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered",
	})

}

func VerifyEmail(c *gin.Context) {
	body := &validations.VerifyEmailParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	users, err := db.FindUserByEmail(context.Background(), body.Email)

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot find account with provided email",
		})
		return
	}

	user := users[0]

	if user.IsEmailVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Account already verified",
		})
		return
	}

	if !user.Otp.Valid || !user.OtpExpiry.Valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "OTP expired",
		})
		return
	}

	if user.OtpExpiry.Time.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "OTP expired",
		})
		return
	}

	if user.Otp.Int32 != int32(body.OTP) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "OTP did not match",
		})
		return
	}

	_, err = db.VerifyUser(context.Background(), database.VerifyUserParams{
		ID:              user.ID,
		IsEmailVerified: true,
		Otp: sql.NullInt32{
			Valid: false,
		},
		OtpExpiry: sql.NullTime{
			Valid: false,
		},
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Verified",
	})

}

func Login(c *gin.Context) {
	body := &validations.LoginParameters{}

	if valid := validations.Validate(c, body); !valid {
		return
	}

	users, err := db.LoginQuery(context.Background(), body.Identifier)

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	user := users[0]

	if !user.IsEmailVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email not verified yet",
		})
		return
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateJWT(int(user.ID))

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
