package usecase

import (
	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type PaymentUsecase interface {
	GetPayments(status *string) ([]entity.Payment, error)
	GetSummary() (*entity.PaymentSummary, error)
}

type paymentUC struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(repo repository.PaymentRepository) PaymentUsecase {
	return &paymentUC{repo: repo}
}

func (u *paymentUC) GetPayments(status *string) ([]entity.Payment, error) {
	return u.repo.GetPayments(status)
}

func (u *paymentUC) GetSummary() (*entity.PaymentSummary, error) {
	return u.repo.GetSummary()
}
