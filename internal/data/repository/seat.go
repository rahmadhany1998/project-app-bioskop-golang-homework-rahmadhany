package repository

import (
	"context"
	"database/sql"
	"errors"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"

	"go.uber.org/zap"
)

type SeatRepository interface {
	GetSeatIDByCode(ctx context.Context, cinemaID int, seatCode string) (int, error)
	GetSeatStatusBySchedule(ctx context.Context, cinemaID int, date, time string) ([]entity.Seat, error)
}

type seatRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewSeatRepository(db *sql.DB, log *zap.Logger) SeatRepository {
	return &seatRepository{
		db:  db,
		log: log,
	}
}

func (r *seatRepository) GetSeatIDByCode(ctx context.Context, cinemaID int, seatCode string) (int, error) {
	query := `SELECT id FROM seats WHERE cinema_id = $1 AND seat_code = $2`
	var seatID int
	err := r.db.QueryRowContext(ctx, query, cinemaID, seatCode).Scan(&seatID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("seat not found")
		}
		return 0, err
	}
	return seatID, nil
}

func (r *seatRepository) GetSeatStatusBySchedule(ctx context.Context, cinemaID int, date, time string) ([]entity.Seat, error) {
	query := `
		SELECT s.seat_code,
			CASE
				WHEN b.seat_id IS NULL THEN 'available'
				ELSE 'booked'
			END AS status
		FROM seats s
		LEFT JOIN bookings b ON s.id = b.seat_id
			AND b.booking_date = $2
			AND b.booking_time = $3
			AND b.cinema_id = $1
		WHERE s.cinema_id = $1
		ORDER BY s.seat_code`

	rows, err := r.db.QueryContext(ctx, query, cinemaID, date, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.Seat
	for rows.Next() {
		var seat entity.Seat
		err := rows.Scan(&seat.SeatCode, &seat.Status)
		if err != nil {
			return nil, err
		}
		results = append(results, seat)
	}
	return results, nil
}
