package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizqyfahmi/gin-greetings-clean-architecture/internal/greetings/domain/usecase"
)

type GreetingsHandler interface {
	Index(c *gin.Context)
}

type GreetingsHandlerImpl struct {
	usecase usecase.GreetingsUsecase
}

func NewGreetingsHandler(
	usecase usecase.GreetingsUsecase,
) GreetingsHandler {
	return &GreetingsHandlerImpl{
		usecase: usecase,
	}
}

func (h *GreetingsHandlerImpl) Index(c *gin.Context) {
	result := h.usecase.Index()
	response := &gin.H{
		"status":  false,
		"message": result.Message,
		"data":    nil,
	}
	c.JSON(http.StatusOK, response)
}
