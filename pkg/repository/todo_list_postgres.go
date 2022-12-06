package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest-hw/model"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list model.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return -1, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", tableTodoList)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	err = row.Scan(&id)

	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", tableUserList)
	_, err = tx.Exec(createUserListQuery, userId, id)

	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}

	return id, tx.Commit()
}
