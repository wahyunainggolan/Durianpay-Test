package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockPaymentUC struct{}

func (m *mockPaymentUC) GetPayments(status *string) ([]entity.Payment, error) {
	return []entity.Payment{
		{
			ID:           "1",
			MerchantName: "A",
			Amount:       10000,
			Status:       "completed",
		},
	}, nil
}

func (m *mockPaymentUC) GetSummary() (*entity.PaymentSummary, error) {
	return &entity.PaymentSummary{
		Total: 1, Completed: 1, Processing: 0, Failed: 0,
	}, nil
}

type mockPaymentUCError struct{}

func (m *mockPaymentUCError) GetPayments(status *string) ([]entity.Payment, error) {
	return nil, errors.New("db error")
}

func (m *mockPaymentUCError) GetSummary() (*entity.PaymentSummary, error) {
	return nil, errors.New("summary error")
}

func setupPaymentHandler(uc usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUC: uc,
	}
}

func TestGetPayments_InvalidStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := setupPaymentHandler(&mockPaymentUC{})

	req := httptest.NewRequest("GET", "/payments?status=invalid_status", nil)
	w := httptest.NewRecorder()

	stringOne := "1"
	params := openapigen.GetDashboardV1PaymentsParams{
		Status: &stringOne,
	}

	h.GetPayments(w, req, params)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPayments_UsecaseError(t *testing.T) {
	h := setupPaymentHandler(&mockPaymentUCError{})

	req := httptest.NewRequest("GET", "/payments", nil)
	w := httptest.NewRecorder()

	params := openapigen.GetDashboardV1PaymentsParams{}

	h.GetPayments(w, req, params)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

type mockPaymentUCPartial struct{}

func (m *mockPaymentUCPartial) GetPayments(status *string) ([]entity.Payment, error) {
	return []entity.Payment{
		{ID: "1", MerchantName: "A", Amount: 10000, Status: "completed"},
	}, nil
}

func (m *mockPaymentUCPartial) GetSummary() (*entity.PaymentSummary, error) {
	return nil, errors.New("summary error")
}

func TestGetPayments_SummaryError(t *testing.T) {
	h := setupPaymentHandler(&mockPaymentUCPartial{})

	req := httptest.NewRequest("GET", "/payments", nil)
	w := httptest.NewRecorder()

	params := openapigen.GetDashboardV1PaymentsParams{}

	h.GetPayments(w, req, params)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetPayments_Success(t *testing.T) {
	h := setupPaymentHandler(&mockPaymentUC{})

	req := httptest.NewRequest("GET", "/payments", nil)
	w := httptest.NewRecorder()

	params := openapigen.GetDashboardV1PaymentsParams{}

	h.GetPayments(w, req, params)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "A")
	assert.Contains(t, w.Body.String(), "completed")
}
