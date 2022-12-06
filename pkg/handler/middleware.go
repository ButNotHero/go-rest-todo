package handler

import (
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
