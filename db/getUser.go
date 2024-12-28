package db

import (
	"database/sql"
	"ecomerce/data"
)

func GetUser(email string) (data.User, error) {
	user := data.User{}

	stm := `SELECT id, firstname, lastname, phonenumber, email, password FROM users WHERE email = ?`

	err := DB.QueryRow(stm, email).Scan(&user.Id, &user.FName, &user.LName, &user.Phone, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return user, err
	} else if err != nil {
		return user, err
	}

	return user, nil
}