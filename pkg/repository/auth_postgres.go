package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest-hw/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", tableAuthUser)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (int, error) {
	var userId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", tableAuthUser)
	err := r.db.Get(&userId, query, username, password)

	return userId, err
}
