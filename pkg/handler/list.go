package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-hw/model"
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

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
