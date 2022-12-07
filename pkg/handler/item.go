package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-hw/model"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Create item. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Create item. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	var input model.TodoItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, "Create item. Invalid input model", http.StatusInternalServerError, err.Error())
		return
	}

	todoItemId, err := h.services.TodoItem.Create(userId, listId, input)

	if err != nil {
		newErrorResponse(c, "Create item. Error creating item. ", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": todoItemId,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Get all items. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Get all items. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)

	if err != nil {
		newErrorResponse(c, "Get all items. Error when getting items", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Get item by id. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Get item by id. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)

	if err != nil {
		newErrorResponse(c, "Get item by id. Error when getting item", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Update item. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Update item. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	var input model.UpdateListItemInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, "Update item. Invalid input", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, itemId, input); err != nil {
		newErrorResponse(c, "Update item. Error when make update", http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, ok := getUserId(c)

	if ok != nil {
		newErrorResponse(c, "Delete item by id. Invalid userId", http.StatusInternalServerError, "Invalid user id")
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, "Delete item by id. Invalid id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)

	if err != nil {
		newErrorResponse(c, "Delete item by id. Error when getting item", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
