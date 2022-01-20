package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lisxAPI/db"
	"lisxAPI/models"
	"lisxAPI/repos"
	"log"
	"strconv"
)

func GetAPIKeyPermissions(c *fiber.Ctx) error {
	apiKeyID, err := strconv.Atoi(c.Query("api-key-id"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	permissions, err := repos.SelectAPIKeyPermissionByAPIKeyID(apiKeyID)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(permissions)
}

func UpdateAPIKeyPermissions(c *fiber.Ctx) error {
	apiKeyID, err := strconv.Atoi(c.Query("api-key-id"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	var apiKeyPermissionUpdates []models.APIKeyPermissionUpdate
	if err := c.BodyParser(&apiKeyPermissionUpdates); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	tx, err := db.DB.Begin()
	_, err = tx.Exec("delete from api_key_permission where api_key_id = $1", apiKeyID)
	if err != nil {
		_ = tx.Rollback()
		return fiber.ErrInternalServerError
	}
	for _, apiKeyPermissionUpdate := range apiKeyPermissionUpdates {
		_, err = tx.Exec(
			"insert into api_key_permission (api_key_id, resource, can_create, can_read, can_update, can_delete) values ($1, $2, $3, $4, $5, $6)",
			apiKeyID,
			apiKeyPermissionUpdate.Resource,
			apiKeyPermissionUpdate.CanCreate,
			apiKeyPermissionUpdate.CanRead,
			apiKeyPermissionUpdate.CanUpdate,
			apiKeyPermissionUpdate.CanDelete,
		)
		if err != nil {
			_ = tx.Rollback()
			return fiber.ErrInternalServerError
		}
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)
}
