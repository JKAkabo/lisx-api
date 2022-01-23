package repos

import (
	"lisxAPI/db"
	"lisxAPI/models"
)

func SelectUserPermissionByUserID(userID int) (userPermissions []models.UserPermission, err error) {
	userPermissions = []models.UserPermission{}
	err = db.DB.Select(&userPermissions, "select * from user_permission where user_id = $1", userID)
	return userPermissions, err
}

func SelectUserPermissionByUserIDAndResource(userID int, resource string) (userPermission models.UserPermission, err error) {
	err = db.DB.Get(&userPermission, "select * from user_permission where user_id = $1 and resource = $2", userID, resource)
	return userPermission, err
}
