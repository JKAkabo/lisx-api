package repos

import (
	"lisxAPI/db"
	"lisxAPI/models"
)

func InsertUser(firstName string, lastName string, username string, password string, isAdmin bool) (int, error) {
	row := db.DB.QueryRow(
		"insert into \"user\" (first_name, last_name, username, password, is_admin) values ($1, $2, $3, $4, $5) returning id",
		firstName,
		lastName,
		username,
		password,
		isAdmin,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func SelectUserById(id int) (user models.User, err error) {
	err = db.DB.Get(&user, "select * from \"user\" where id = $1", id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SelectUserByUsername(username string) (user models.User, err error) {
	err = db.DB.Get(&user, "select * from \"user\" where username = $1", username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func SelectUsers() (users []models.User, err error) {
	err = db.DB.Select(&users, "select * from \"user\"")
	if err != nil {
		return users, err
	}
	return users, nil
}
