package models

type UserCreate struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Type      string `json:"type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

type User struct {
	ID                  int    `json:"id" db:"id"`
	FirstName           string `json:"firstName" db:"first_name"`
	LastName            string `json:"lastName" db:"last_name"`
	Type                string `json:"type" db:"type"`
	Username            string `json:"username" db:"username"`
	Password            string `json:"-" db:"password"`
	ForcePasswordChange bool   `json:"forcePasswordChange" db:"force_password_change"`
	IsAdmin             bool   `json:"isAdmin" db:"is_admin"`
}
