package repository

import (
	"context"
	"database/sql"
	"errors"
	"example/model"
)

type User struct {
	DB *sql.DB
}

func (u User) Register(username, password string) (user model.User) {
	query := `
		INSERT INTO users (username, password, role)
		VALUES (?, ?, ?)
	`
	result, err := u.DB.ExecContext(context.Background(), query, username, password, "admin")
	if err != nil {
		panic(err)
	}

	newId, _ := result.LastInsertId()
	user.Id = int(newId)
	user.Username = username
	user.Password = password
	user.Role = "admin"

	return user
}

func (u User) FindByUsername(username string) (user model.User, err error) {
	query := `
		SELECT * FROM users WHERE username = ?
	`
	rows, err := u.DB.QueryContext(context.Background(), query, username)
	if err != nil {
		panic(err)
	}

	if !rows.Next() {
		return user, errors.New("user not found")
	}

	rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	return user, err
}
