package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckRequest struct {
	Text string
}

// @Summary Email Check
// @Tags email
// @Description Email Check
// @ModuleID emailCheck
// @Produce  json
// @Param CheckRequest body CheckRequest true "CheckRequest"
// @Success 200 {string} string
// @Router /rest/email/check [post]
func (h *Handler) emailCheck(c echo.Context) error {
	emailRequest := new(CheckRequest)
	if err := c.Bind(emailRequest); err != nil {
		return err
	}

	email, err := h.emailService.FindEmailFromText(emailRequest.Text)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, email)
}

// @Summary IIN Check
// @Tags email
// @Description IIN Check
// @ModuleID iinCheck
// @Produce  json
// @Param CheckRequest body CheckRequest true "CheckRequest"
// @Success 200 {string} string
// @Router /rest/email/iin [post]
func (h *Handler) iinCheck(c echo.Context) error {
	checkRequest := new(CheckRequest)
	if err := c.Bind(checkRequest); err != nil {
		return err
	}

	iin, err := h.emailService.FindIINFromText(checkRequest.Text)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, iin)
}
