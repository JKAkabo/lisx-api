package models

type APIKeyCreate struct {
	Name string `json:"name"`
}

type APIKey struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Key    string `json:"key" db:"key"`
	UserID int    `json:"userID" db:"user_id"`
}
