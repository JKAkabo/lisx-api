package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"lisxAPI/repos"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetUserID(c *fiber.Ctx) (int, error) {
	key := c.Params("api-key")
	if key != "" {
		apiKey, err := repos.SelectAPIKeyByKey(key)
		if err != nil {
			log.Printf("error: %v", err)
			return 0, errors.New("invalid api key")
		}
		return apiKey.UserID, nil
	}
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("missing authorization header")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("malformed authorization header")
	}
	claims, err := DecodeJWT(parts[1])
	if err != nil {
		return 0, err
	}
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["userID"]))
	if err != nil {
		return 0, errors.New("malformed authorization token")
	}
	return userID, nil
}

func DecodeJWT(t string) (claims jwt.MapClaims, err error) {
	claims = jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return claims, err
}
