package http

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/omekov/sosedi-group/internal/counter"
	mocks "github.com/omekov/sosedi-group/mocks"
)

func TestCounterAdd(t *testing.T) {
	type mockBehavior func(r *mocks.MockCounter, num uint64)
	var num uint64 = 1
	testCases := []struct {
		name         string
		statusCode   int
		mockBehavior mockBehavior
		responseBody string
		num          uint64
	}{
		{
			name:       "success",
			statusCode: 200,
			num:        num,
			mockBehavior: func(r *mocks.MockCounter, num uint64) {
				r.EXPECT().Set(context.Background(), num).Return(nil)
				r.EXPECT().Get(context.Background()).Return(strconv.Itoa(int(num)), nil)
			},
			responseBody: strconv.Itoa(int(num + 1)),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			counterRepo := mocks.NewMockCounter(c)
			tt.mockBehavior(counterRepo, tt.num+1)

			counterService := counter.Service{
				Repository: counterRepo,
			}

			handler := NewHandler(&counterService)

			router := echo.New()
			router.POST("/add/:num", handler.counterAdd)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/add/%d", tt.num), nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}
func TestCounterSub(t *testing.T) {
	type mockBehavior func(r *mocks.MockCounter, num uint64)
	var num uint64 = 12
	testCases := []struct {
		name         string
		statusCode   int
		mockBehavior mockBehavior
		responseBody string
		num          uint64
	}{
		{
			name:       "success",
			statusCode: 200,
			num:        num,
			mockBehavior: func(r *mocks.MockCounter, num uint64) {
				r.EXPECT().Set(context.Background(), num).Return(nil)
				r.EXPECT().Get(context.Background()).Return(strconv.Itoa(int(num)), nil)
			},
			responseBody: strconv.Itoa(int(num - 1)),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			counterRepo := mocks.NewMockCounter(c)
			tt.mockBehavior(counterRepo, tt.num-1)

			counterService := counter.Service{
				Repository: counterRepo,
			}

			handler := NewHandler(&counterService)

			router := echo.New()
			router.POST("/sub/:num", handler.counterSub)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/sub/%d", tt.num), nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}
}
func TestCounterVal(t *testing.T) {
	type mockBehavior func(r *mocks.MockCounter, num uint64)
	testCases := []struct {
		name         string
		statusCode   int
		mockBehavior mockBehavior
	}{
		{
			name:       "error",
			statusCode: 500,
			mockBehavior: func(r *mocks.MockCounter, num uint64) {
				r.EXPECT().Set(context.Background(), num).Return(nil)
				r.EXPECT().Get(context.Background()).Return(strconv.Itoa(int(num)), nil)
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			counterRepo := mocks.NewMockCounter(c)
			tt.mockBehavior(counterRepo, 4)

			counterService := counter.Service{
				Repository: counterRepo,
			}

			handler := NewHandler(&counterService)

			router := echo.New()
			router.GET("/val", handler.counterSub)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/val", nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.statusCode)
		})
	}
}
