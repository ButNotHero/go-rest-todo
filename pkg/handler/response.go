package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, place string, statusCode int, message string) {
	logrus.Errorf("[%d] %s - %s", statusCode, place, message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
	})
}
