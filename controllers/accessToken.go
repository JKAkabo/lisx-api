package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"lisxAPI/models"
	"lisxAPI/repos"
	"log"
	"os"
	"time"
)

func CreateAccessToken(c *fiber.Ctx) error {
	var credentials models.Credentials
	if err := c.BodyParser(&credentials); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	user, err := repos.SelectUserByUsername(credentials.Username)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrUnauthorized
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrUnauthorized
	}
	claims := jwt.MapClaims{}
	expiresAt := time.Now().Add(time.Minute * 45).Unix()
	claims["exp"] = expiresAt
	claims["userID"] = user.ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON(models.AccessToken{
		Token:     tokenStr,
		ExpiresAt: int(expiresAt),
		UserID:    user.ID,
	})
}
