package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-hw/model"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Create list invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	var input model.TodoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, "Create list invalid input", http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)

	if err != nil {
		newErrorResponse(c, "Create list", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

type getAllListsResponse struct {
	Lists []model.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Get all lists invalid userId", http.StatusBadRequest, "Invalid user id")
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)

	if err != nil {
		newErrorResponse(c, "Create list", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Lists: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Get list by id. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Get list by id. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.GetById(userId, listId)

	if err != nil {
		newErrorResponse(c, "Get list by id", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Update list. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Update list. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	var input model.UpdateListInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, "Update list. Invalid input", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, listId, input); err != nil {
		newErrorResponse(c, "Update list. Error when make update", http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Delete list. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Delete list. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Delete(userId, listId)

	if err != nil {
		newErrorResponse(c, "Delete list", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
