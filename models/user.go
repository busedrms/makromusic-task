package models

import (
	"context"
	"makromusic-task/utils"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Create() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err = utils.DB.QueryRow(context.Background(), query, u.Username, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := utils.DB.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
