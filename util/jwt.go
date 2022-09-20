package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const SigningKey = "allBaseKey"

func GenerateJwt(id int, username ...string) (string, error) {
	// todo claim 改一下里面的内容；改完以后记得修改 User里面的jwt.Parse
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    strconv.Itoa(id),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SigningKey))
	return signedToken, err
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SigningKey), nil
	})

	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		userID := claim["iss"].(string)
		return userID, nil

	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return "", errors.New("that's not even a token")

	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		return "", errors.New("timing is everything")

	} else {
		return "", fmt.Errorf("couldn't handle this token: %v", err)
	}
}
