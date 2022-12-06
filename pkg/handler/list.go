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
		newErrorResponse(c, "Create list invalid userId", http.StatusInternalServerError, "Invalid userId")
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
		newErrorResponse(c, "Get all lists invalid userId", http.StatusBadRequest, "Invalid userId")
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
		newErrorResponse(c, "Get list by id. Invalid userId", http.StatusInternalServerError, "Invalid userId")
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

}

func (h *Handler) deleteList(c *gin.Context) {

}
