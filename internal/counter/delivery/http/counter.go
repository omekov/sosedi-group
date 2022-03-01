package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Counter Add (increment)
// @Tags counter
// @Description Redis increment
// @ModuleID counterAdd
// @Produce  json
// @Success 200 {string} string
// @Router /rest/counter/add/{num} [post]
func (h *Handler) counterAdd(c echo.Context) error {
	num := c.Param("num")
	if err := h.counterService.SetIncrement(c.Request().Context(), num); err != nil {
		return c.String(http.StatusInternalServerError, "SetIncrement")
	}

	num, err := h.counterService.Get(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Get")
	}
	return c.String(http.StatusOK, num)
}

// @Summary Counter Sub (decrement)
// @Tags counter
// @Description Redis decrement
// @ModuleID counterSub
// @Produce  json
// @Success 200 {string} string
// @Router /rest/counter/sub/{num} [post]
func (h *Handler) counterSub(c echo.Context) error {
	num := c.Param("num")
	if err := h.counterService.SetDecrement(c.Request().Context(), num); err != nil {
		return c.String(http.StatusInternalServerError, "SetDecrement")
	}

	num, err := h.counterService.Get(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Get")
	}
	return c.String(http.StatusOK, num)
}

// @Summary Counter Val
// @Tags counter
// @Description Redis get num
// @ModuleID counterVal
// @Produce  json
// @Success 200 {string} string
// @Router /rest/counter/val [get]
func (h *Handler) counterVal(c echo.Context) error {
	num, err := h.counterService.Get(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Get")
	}

	return c.String(http.StatusOK, num)
}
