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

func (r *TodoListPostgres) GetAll(userId int) ([]model.TodoList, error) {
	var lists []model.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1", tableTodoList, tableUserList)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
