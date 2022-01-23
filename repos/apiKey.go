package repos

import (
	"lisxAPI/db"
	"lisxAPI/models"
)

func InsertAPIKey(
	name string,
	key string,
	userID int,
) (int, error) {
	row := db.DB.QueryRow(
		"insert into api_key (name, key, user_id) values ($1, $2, $3) returning id",
		name,
		key,
		userID,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func SelectAPIKeyById(id int) (apiKey models.APIKey, err error) {
	err = db.DB.Get(&apiKey, "select * from api_key where id = $1", id)
	if err != nil {
		return apiKey, err
	}
	return apiKey, nil
}

func SelectAPIKeyByKey(key string) (apiKey models.APIKey, err error) {
	err = db.DB.Get(&apiKey, "select * from api_key where key = $1", key)
	if err != nil {
		return apiKey, err
	}
	return apiKey, nil
}

func SelectAPIKeys() (apiKeys []models.APIKey, err error) {
	apiKeys = []models.APIKey{}
	err = db.DB.Select(&apiKeys, "select * from api_key")
	if err != nil {
		return apiKeys, err
	}
	return apiKeys, nil
}

func APIKeyNameExists(name string) (bool, error) {
	row := db.DB.QueryRow("select exists(select * from api_key where name = $1)", name)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
