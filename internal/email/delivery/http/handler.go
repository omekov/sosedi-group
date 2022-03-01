package http

import (
	"github.com/labstack/echo/v4"
	"github.com/omekov/sosedi-group/internal/email"
)

type Handler struct {
	counterService *email.Service
}

func NewHandler(service *email.Service) *Handler {
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
