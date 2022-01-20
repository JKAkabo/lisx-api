package models

type UserPermissionUpdate struct {
	Resource  string `json:"resource"`
	CanCreate bool   `json:"canCreate"`
	CanRead   bool   `json:"canRead"`
	CanUpdate bool   `json:"canUpdate"`
	CanDelete bool   `json:"canDelete"`
}

type UserPermission struct {
	UserID    int    `json:"-" db:"user_id"`
	Resource  string `json:"resource" db:"resource"`
	CanCreate bool   `json:"canCreate" db:"can_create"`
	CanRead   bool   `json:"canRead" db:"can_read"`
	CanUpdate bool   `json:"canUpdate" db:"can_update"`
	CanDelete bool   `json:"canDelete" db:"can_delete"`
}