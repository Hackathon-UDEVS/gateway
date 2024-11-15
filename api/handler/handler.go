package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	// Add any required dependencies here
}

type Handlers struct {
	// Add any required dependencies here
}

type HandlerInterface interface {
	Test(c *gin.Context)
}

func NewHandler() HandlerInterface {
	return &Handlers{}
}
