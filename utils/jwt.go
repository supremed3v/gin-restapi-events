package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse error")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("JWT Token is not valid")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		if !tokenIsValid {
			return errors.New("JWT Token is not valid")
		}
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil

}
