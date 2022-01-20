package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lisxAPI/db"
	"lisxAPI/models"
	"lisxAPI/repos"
	"log"
	"strconv"
)

func GetUserPermissions(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Query("user-id"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	userPermissions, err := repos.SelectUserPermissionByUserID(userID)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(userPermissions)
}

func UpdateUserPermissions(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Query("user-id"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	var userPermissionUpdates []models.UserPermissionUpdate
	if err := c.BodyParser(&userPermissionUpdates); err != nil {
		log.Printf("error: %v", err)
		return fiber.ErrBadRequest
	}
	tx, err := db.DB.Begin()
	_, err = tx.Exec("delete from user_permission where user_id = $1", userID)
	if err != nil {
		_ = tx.Rollback()
		return fiber.ErrInternalServerError
	}
	for _, userPermissionUpdate := range userPermissionUpdates {
		_, err = tx.Exec(
			"insert into user_permission (user_id, resource, can_create, can_read, can_update, can_delete) values ($1, $2, $3, $4, $5, $6)",
			userID,
			userPermissionUpdate.Resource,
			userPermissionUpdate.CanCreate,
			userPermissionUpdate.CanRead,
			userPermissionUpdate.CanUpdate,
			userPermissionUpdate.CanDelete,
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
