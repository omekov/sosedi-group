package http

import (
	"github.com/labstack/echo/v4"
	"github.com/omekov/sosedi-group/internal/firstname"
)

type Handler struct {
	counterService *firstname.Service
}

func NewHandler(service *firstname.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Init(api *echo.Group) {
	// substrApi := api.Group("/substr")
	// {

	// }

	// selfApi := api.Group("/self")
	// {

	// }
}
