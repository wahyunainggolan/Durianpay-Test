package handler

import (
	"net/http"
	"strconv"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
	"github.com/durianpay/fullstack-boilerplate/internal/util"
)

type PaymentHandler struct {
	paymentUC usecase.PaymentUsecase
}

func NewPaymentHandler(uc usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUC: uc,
	}
}

func (h *PaymentHandler) GetPayments(
	w http.ResponseWriter,
	r *http.Request,
	params openapigen.GetDashboardV1PaymentsParams,
) {
	if params.Status != nil && !util.IsValidPaymentStatus(*params.Status) {
		transport.WriteAppError(w, entity.ErrorBadRequest("invalid status"))
		return
	}

	payments, err := h.paymentUC.GetPayments(params.Status)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	var res []openapigen.Payment
	for _, p := range payments {
		amountStr := strconv.FormatFloat(p.Amount, 'f', -1, 64)

		res = append(res, openapigen.Payment{
			Id:        &p.ID,
			Merchant:  &p.MerchantName,
			Amount:    &amountStr,
			Status:    &p.Status,
			CreatedAt: &p.CreatedAt,
		})
	}

	summaryEntity, err := h.paymentUC.GetSummary()
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	summary := openapigen.PaymentSummary{
		Total:      summaryEntity.Total,
		Completed:  summaryEntity.Completed,
		Processing: summaryEntity.Processing,
		Failed:     summaryEntity.Failed,
	}

	response := openapigen.PaymentListResponse{
		Summary:  &summary,
		Payments: &res,
	}

	transport.WriteSuccess(w, response)
}
