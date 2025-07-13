package repository

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

type SeatRepository interface {
	GetSeatIDByCode(ctx context.Context, cinemaID int, seatCode string) (int, error)
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
