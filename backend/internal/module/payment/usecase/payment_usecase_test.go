package usecase

import (
	"testing"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/stretchr/testify/assert"
)

type mockPaymentRepo struct{}

func (m *mockPaymentRepo) GetPayments(status *string) ([]entity.Payment, error) {
	return []entity.Payment{
		{
			ID:           "1",
			MerchantName: "A",
			Amount:       10000,
			Status:       "completed",
		},
	}, nil
}

func (m *mockPaymentRepo) GetSummary() (*entity.PaymentSummary, error) {
	return &entity.PaymentSummary{
		Total:      1,
		Completed:  1,
		Processing: 0,
		Failed:     0,
	}, nil
}

func TestGetPayments(t *testing.T) {
	repo := &mockPaymentRepo{}
	uc := NewPaymentUsecase(repo)

	status := "completed"
	data, err := uc.GetPayments(&status)

	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, "A", data[0].MerchantName)
	assert.Equal(t, "completed", data[0].Status)
}

func TestGetSummary(t *testing.T) {
	repo := &mockPaymentRepo{}
	uc := NewPaymentUsecase(repo)

	summary, err := uc.GetSummary()

	assert.NoError(t, err)
	assert.NotNil(t, summary)
	assert.Equal(t, 1, summary.Total)
	assert.Equal(t, 1, summary.Completed)
	assert.Equal(t, 0, summary.Processing)
	assert.Equal(t, 0, summary.Failed)
}
