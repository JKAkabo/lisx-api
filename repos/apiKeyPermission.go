package repos

import (
	"lisxAPI/db"
	"lisxAPI/models"
)

func SelectAPIKeyPermissionByAPIKeyID(apiKeyID int) (apiKeyPermissions []models.APIKeyPermission, err error) {
	apiKeyPermissions = []models.APIKeyPermission{}
	err = db.DB.Select(&apiKeyPermissions, "select * from api_key_permission where api_key_id = $1", apiKeyID)
	return apiKeyPermissions, err
}

func SelectAPIKeyPermissionByAPIKeyIDAndResource(apiKeyID int, resource string) (apiKeyPermission models.APIKeyPermission, err error) {
	err = db.DB.Get(&apiKeyPermission, "select * from api_key_permission where api_key_id = $1 and resource = $2", apiKeyID, resource)
	return apiKeyPermission, err
}
