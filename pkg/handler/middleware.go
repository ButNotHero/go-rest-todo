package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	ctxUserId  = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(c, "User identity get header", http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, "User identity get header parts length", http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseJwtToken(headerParts[1])

	if err != nil {
		newErrorResponse(c, "User identity parse jwt token", http.StatusUnauthorized, "Invalid jwt token")
		return
	}

	c.Set(ctxUserId, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(ctxUserId)

	if !ok {
		newErrorResponse(c, "getUserId user id not in context", http.StatusInternalServerError, "user id not found")
		return -1, errors.New("user id not found")
	}

	idInt, ok := id.(int)

	if !ok {
		newErrorResponse(c, "getUserId invalid conversion", http.StatusInternalServerError, "invalid user id")
		return -1, errors.New("invalid user id")
	}

	return idInt, nil
}
