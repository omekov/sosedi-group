package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/omekov/sosedi-group/internal/email"
	mocks "github.com/omekov/sosedi-group/internal/email/mocks"
)

func TestEmailCheck(t *testing.T) {
	type mockBehavior func(r *mocks.MockEmailService, text, email string)
	testCases := []struct {
		name         string
		statusCode   int
		mockBehavior mockBehavior
		responseBody string
		text         string
	}{
		{
			name:       "success",
			statusCode: 200,
			mockBehavior: func(r *mocks.MockEmailService, text, email string) {
				r.EXPECT().FindEmailFromText(text).Return(email, nil)
			},
			responseBody: "test@example.com",
			text:         "Здесь ограмный текст с пробелами и обзацом \n \r test@example.com",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			emailRepo := mocks.NewMockEmailService(c)
			tt.mockBehavior(emailRepo, tt.text, "test@example.com")

			emailService := email.Service{
				Repository: emailRepo,
			}

			handler := NewHandler(&emailService)

			router := echo.New()
			router.POST("/check", handler.emailCheck)

			w := httptest.NewRecorder()
			reqJson, err := json.Marshal(CheckRequest{Text: tt.text})
			if err != nil {
				t.Error(err)
			}

			req := httptest.NewRequest(http.MethodPost, "/check", bytes.NewBuffer(reqJson))
			req.Header.Add("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.statusCode)
			assert.Equal(t, w.Body.String(), tt.responseBody)
		})
	}

}
