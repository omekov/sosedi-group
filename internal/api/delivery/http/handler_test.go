package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/omekov/sosedi-group/internal/api/delivery/http"
	"github.com/omekov/sosedi-group/internal/counter"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(&counter.Service{})

	require.IsType(t, &handler.Handler{}, h)
}

func TestNewHandler_Init(t *testing.T) {
	h := handler.NewHandler(&counter.Service{})

	router := h.Init()

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, http.StatusOK, res.StatusCode)
}
