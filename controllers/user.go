package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"lisxAPI/models"
	"lisxAPI/repos"
	"log"
)

func CreateUser(c *fiber.Ctx) error {
	var userCreate models.UserCreate
	if err := c.BodyParser(&userCreate); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	existingUser, _ := repos.SelectUserByUsername(userCreate.Username)
	if existingUser.ID != 0 {
		return &fiber.Error{Code: fiber.StatusBadRequest, Message: "username has already been taken"}
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(userCreate.Password), 10)
	id, err := repos.InsertUser(
		userCreate.FirstName,
		userCreate.LastName,
		userCreate.Username,
		string(hash),
		userCreate.IsAdmin,
	)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	user, err := repos.SelectUserById(id)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	user, err := repos.SelectUserById(id)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrNotFound
	}
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := repos.SelectUsers()
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(users)
}
