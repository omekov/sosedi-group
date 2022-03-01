package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omekov/sosedi-group/internal/counter"
	counterHTTP "github.com/omekov/sosedi-group/internal/counter/delivery/http"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	counterService *counter.Service
}

func NewHandler(counterService *counter.Service) *Handler {
	return &Handler{
		counterService,
	}
}

func (h *Handler) Init() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	handlerV1 := counterHTTP.NewHandler(h.counterService)
	counterGroup := router.Group("/rest")
	{
		handlerV1.Init(counterGroup)
	}

	return router
}
