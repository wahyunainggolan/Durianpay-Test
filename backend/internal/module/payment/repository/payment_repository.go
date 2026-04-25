package repository

import (
	"database/sql"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
)

type PaymentRepository interface {
	GetPayments(status *string) ([]entity.Payment, error)
}

type paymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) PaymentRepository {
	return &paymentRepo{db: db}
}

func (r *paymentRepo) GetPayments(status *string) ([]entity.Payment, error) {
	query := "SELECT id, merchant, amount, status, created_at FROM payments WHERE 1=1"
	args := []interface{}{}

	if status != nil {
		query += " AND status = ?"
		args = append(args, *status)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entity.Payment

	for rows.Next() {
		var p entity.Payment
		err := rows.Scan(&p.ID, &p.MerchantName, &p.Amount, &p.Status, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}
