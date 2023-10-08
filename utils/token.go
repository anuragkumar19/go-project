package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(ID int) (tokenString string, err error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // 1 month

	// Sign and get the complete encoded token as a string
	tokenString, err = token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func VerifyToken(tokenString string) (int, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if token.Valid {
		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, err
		}

		// Access specific claim values
		userId := claims["id"].(float64)

		return int(userId), nil
	} else {
		return 0, fmt.Errorf("token not valid")
	}
}
