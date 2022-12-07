package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest-hw/model"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item model.TodoItem) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return -1, err
	}

	var itemId int

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", tableTodoItem)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", tableListItem)
	_, err = tx.Exec(createListItemQuery, listId, itemId)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]model.TodoItem, error) {
	var items []model.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id
                                INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		tableTodoItem, tableListItem, tableUserList)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (model.TodoItem, error) {
	var item model.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id
                                INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		tableTodoItem, tableListItem, tableUserList)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id
                                         AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		tableTodoItem, tableListItem, tableUserList)
	_, err := r.db.Exec(query, userId, itemId)

	return err
}
