package service

import (
	"rest-hw/model"
	"rest-hw/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item model.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)

	if err != nil {
		// list does not exist or does not belong to user
		return -1, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]model.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
