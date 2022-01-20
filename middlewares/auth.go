package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"lisxAPI/repos"
	"lisxAPI/utils"
	"strings"
)

func RequireCreatePermission(c *fiber.Ctx) error {
	userID, err := utils.GetUserID(c)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	}
	s := strings.Split(c.Path(), "/")
	if len(s) < 2 {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "no resource found in path",
		}
	}
	resource := s[1]
	key := c.Params("api-key")
	if key != "" {
		apiKey, _ := repos.SelectAPIKeyByKey(key)
		apiKeyPermission, _ := repos.SelectAPIKeyPermissionByAPIKeyIDAndResource(apiKey.ID, resource)
		if !apiKeyPermission.CanCreate {
			return fiber.ErrForbidden
		}
	}
	userPermission, _ := repos.SelectUserPermissionByUserIDAndResource(userID, resource)
	if !userPermission.CanCreate {
		return fiber.ErrForbidden
	}
	return c.Next()
}

func RequireReadPermission(c *fiber.Ctx) error {
	userID, err := utils.GetUserID(c)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	}
	s := strings.Split(c.Path(), "/")
	if len(s) < 2 {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "no resource found in path",
		}
	}
	resource := s[1]
	key := c.Params("api-key")
	if key != "" {
		apiKey, _ := repos.SelectAPIKeyByKey(key)
		apiKeyPermission, _ := repos.SelectAPIKeyPermissionByAPIKeyIDAndResource(apiKey.ID, resource)
		if !apiKeyPermission.CanRead {
			return fiber.ErrForbidden
		}
	}
	userPermission, _ := repos.SelectUserPermissionByUserIDAndResource(userID, resource)
	if !userPermission.CanRead {
		return fiber.ErrForbidden
	}
	return c.Next()
}

func RequireUpdatePermission(c *fiber.Ctx) error {
	userID, err := utils.GetUserID(c)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	}
	s := strings.Split(c.Path(), "/")
	if len(s) < 2 {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "no resource found in path",
		}
	}
	resource := s[1]
	key := c.Params("api-key")
	if key != "" {
		apiKey, _ := repos.SelectAPIKeyByKey(key)
		apiKeyPermission, _ := repos.SelectAPIKeyPermissionByAPIKeyIDAndResource(apiKey.ID, resource)
		if !apiKeyPermission.CanUpdate {
			return fiber.ErrForbidden
		}
	}
	userPermission, err := repos.SelectUserPermissionByUserIDAndResource(userID, resource)
	if !userPermission.CanUpdate {
		return fiber.ErrForbidden
	}
	return c.Next()
}

func RequireDeletePermission(c *fiber.Ctx) error {
	userID, err := utils.GetUserID(c)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	}
	s := strings.Split(c.Path(), "/")
	if len(s) < 2 {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "no resource found in path",
		}
	}
	resource := s[1]
	key := c.Params("api-key")
	if key != "" {
		apiKey, _ := repos.SelectAPIKeyByKey(key)
		apiKeyPermission, _ := repos.SelectAPIKeyPermissionByAPIKeyIDAndResource(apiKey.ID, resource)
		if !apiKeyPermission.CanDelete {
			return fiber.ErrForbidden
		}
	}
	userPermission, err := repos.SelectUserPermissionByUserIDAndResource(userID, resource)
	if !userPermission.CanDelete {
		return fiber.ErrForbidden
	}
	return c.Next()
}

func RequireAdmin(c *fiber.Ctx) error {
	// api keys cannot have admin privileges
	if c.Get("api-key") != "" {
		return fiber.ErrForbidden
	}
	userID, err := utils.GetUserID(c)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()}
	}
	user, err := repos.SelectUserById(userID)
	if err != nil {
		return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Invalid authorization token"}
	}
	if !user.IsAdmin {
		return fiber.ErrForbidden
	}
	return c.Next()
}
