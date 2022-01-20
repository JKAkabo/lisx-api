package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"lisxAPI/models"
	"lisxAPI/repos"
	"lisxAPI/utils"
	"log"
	"strings"
)

func GetAPIKeys(c *fiber.Ctx) error {
	apiKeys, err := repos.SelectAPIKeys()
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(apiKeys)
}

func CreateAPIKey(c *fiber.Ctx) error {
	userID, _ := utils.GetUserID(c)
	var apiKeyCreate models.APIKeyCreate
	if err := c.BodyParser(&apiKeyCreate); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	nameExists, err := repos.APIKeyNameExists(apiKeyCreate.Name)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	if nameExists {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "name has already been used",
		}
	}
	id, err := repos.InsertAPIKey(
		apiKeyCreate.Name,
		strings.Replace(uuid.New().String(), "-", "", -1),
		userID,
	)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	apiKey, err := repos.SelectAPIKeyById(id)
	if err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON(apiKey)
}
