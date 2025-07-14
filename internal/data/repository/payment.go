package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"
	"project-app-bioskop-golang-homework-rahmadhany/internal/dto"

	"go.uber.org/zap"
)

type PaymentRepository interface {
	GetAllMethods(ctx context.Context) ([]entity.PaymentMethod, error)
	ProcessPayment(ctx context.Context, req dto.PaymentRequest) (string, error)
}

type paymentRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewPaymentRepository(db *sql.DB, log *zap.Logger) PaymentRepository {
	return &paymentRepository{
		db:  db,
		log: log,
	}
}

func (r *paymentRepository) GetAllMethods(ctx context.Context) ([]entity.PaymentMethod, error) {
	query := `SELECT id, name FROM payment_methods`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []entity.PaymentMethod
	for rows.Next() {
		var method entity.PaymentMethod
		if err := rows.Scan(&method.ID, &method.Name); err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}
	return methods, nil
}

func (r *paymentRepository) ProcessPayment(ctx context.Context, req dto.PaymentRequest) (string, error) {
	// Simulasi gagal jika kartu tidak valid
	if req.PaymentMethod == "credit_card" && req.PaymentDetails["cardNumber"] == "0000-0000-0000-0000" {
		return "", fmt.Errorf("insufficient funds")
	}

	txID := "txn" + req.BookingID

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	insertPayment := `INSERT INTO payments (booking_id, payment_method, transaction_id, status, details) VALUES ($1, $2, $3, $4, $5)`
	detailsJSON, _ := json.Marshal(req.PaymentDetails)
	_, err = tx.ExecContext(ctx, insertPayment, req.BookingID, req.PaymentMethod, txID, "paid", string(detailsJSON))
	if err != nil {
		tx.Rollback()
		return "", err
	}

	updateBooking := `UPDATE bookings SET status = $1 WHERE id = $2`
	_, err = tx.ExecContext(ctx, updateBooking, "paid", req.BookingID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return txID, nil
}
