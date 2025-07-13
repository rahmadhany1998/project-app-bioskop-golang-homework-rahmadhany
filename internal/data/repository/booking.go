package repository

import (
	"context"
	"database/sql"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"

	"go.uber.org/zap"
)

type BookingRepository interface {
	IsSeatBooked(ctx context.Context, cinemaID int, seatCode int, date, time string) (bool, error)
	CreateBooking(ctx context.Context, booking *entity.Booking) error
}

type bookingRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewBookingRepository(db *sql.DB, log *zap.Logger) BookingRepository {
	return &bookingRepository{
		db:  db,
		log: log,
	}
}

func (r *bookingRepository) IsSeatBooked(ctx context.Context, cinemaID int, seatCode int, date, time string) (bool, error) {
	query := `SELECT COUNT(*) FROM bookings WHERE cinema_id = $1 AND seat_id = $2 AND booking_date = $3 AND booking_time = $4`
	var count int
	err := r.db.QueryRowContext(ctx, query, cinemaID, seatCode, date, time).Scan(&count)
	return count > 0, err
}

func (r *bookingRepository) CreateBooking(ctx context.Context, b *entity.Booking) error {
	query := `INSERT INTO bookings (id, user_id, cinema_id, seat_id, booking_date, booking_time, payment_method, status, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP)`
	_, err := r.db.ExecContext(ctx, query, b.ID, b.UserID, b.CinemaID, b.SeatID, b.Date, b.Time, b.PaymentMethod, b.Status)
	return err
}
