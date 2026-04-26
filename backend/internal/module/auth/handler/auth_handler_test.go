package handler

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockAuthUC struct{}

func (m *mockAuthUC) Login(email, password string) (string, *entity.User, error) {
	return "fake-token", &entity.User{
		ID:    "1",
		Email: email,
		Role:  "user",
	}, nil
}

func setupAuthHandler() *AuthHandler {
	return &AuthHandler{
		authUC: &mockAuthUC{},
	}
}

func TestAuthLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := setupAuthHandler()

	reqBody := `{"email":"test@example.com","password":"123456"}`

	req := httptest.NewRequest("POST", "/dashboard/v1/auth/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.PostDashboardV1AuthLogin(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "fake-token")
	assert.Contains(t, w.Body.String(), "test@example.com")
}

type mockAuthUCFail struct{}

func (m *mockAuthUCFail) Login(email, password string) (string, *entity.User, error) {
	return "", nil, assert.AnError
}

func TestAuthLoginHandler_Failed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := &AuthHandler{
		authUC: &mockAuthUCFail{},
	}

	reqBody := `{"email":"wrong@test.com","password":"wrong"}`

	req := httptest.NewRequest("POST", "/dashboard/v1/auth/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.PostDashboardV1AuthLogin(w, req)

	assert.NotEqual(t, 200, w.Code)
}
