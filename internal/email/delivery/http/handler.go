package http

import (
	"github.com/labstack/echo/v4"
	"github.com/omekov/sosedi-group/internal/email"
)

type Handler struct {
	emailService email.EmailService
}

func NewHandler(service email.EmailService) *Handler {
	return &Handler{service}
}

func (h *Handler) Init(api *echo.Group) {
	emailApi := api.Group("/email")
	{
		emailApi.POST("/check", h.emailCheck)
		emailApi.POST("/iin", h.iinCheck)
	}
}
