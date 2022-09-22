package util

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

const SigningKey = "allBaseKey"

type CustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJwt(id int, username string) (string, error) {
	expireDuration := time.Now().Add(time.Hour * 24)
	claims := CustomClaims{
		ID:       id,
		Username: username,
		//Authority: authority,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireDuration),
			Issuer:    "todoLister",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SigningKey))
	return signedToken, err
}

func ParseJwt(cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return strconv.Itoa(claims.ID), nil
		}
	}

	return "", err

	/*
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
	*/
}
