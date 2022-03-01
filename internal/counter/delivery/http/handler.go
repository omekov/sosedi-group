package http

import (
	"github.com/labstack/echo/v4"
	"github.com/omekov/sosedi-group/internal/counter"
)

type Handler struct {
	counterService *counter.Service
}

func NewHandler(service *counter.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Init(api *echo.Group) {
	counterApi := api.Group("/counter")
	{
		counterApi.POST("/add/:num", h.counterAdd)
		counterApi.POST("/sub/:num", h.counterSub)
		counterApi.GET("/val", h.counterVal)
	}
}
