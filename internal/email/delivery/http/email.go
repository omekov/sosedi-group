package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EmailRequest struct {
	Email string
}

// @Summary Email Check
// @Tags email
// @Description Email Check
// @ModuleID emailCheck
// @Produce  json
// @Param EmailRequest body EmailRequest true "EmailRequest"
// @Success 200 {string} string
// @Router /rest/email [post]
func (h *Handler) emailCheck(c echo.Context) error {
	num := c.Param("num")
	return c.String(http.StatusOK, num)
}
