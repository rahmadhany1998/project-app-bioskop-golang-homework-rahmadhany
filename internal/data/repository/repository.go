package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type Repository struct {
	UserRepo    UserRepository
	TokenRepo   TokenRepository
	CinemaRepo  CinemaRepository
	BookingRepo BookingRepository
	SeatRepo    SeatRepository
	PaymentRepo PaymentRepository
}

func NewRepository(db *sql.DB, log *zap.Logger) Repository {
	return Repository{
		UserRepo:    NewUserRepository(db, log),
		TokenRepo:   NewTokenRepository(db, log),
		CinemaRepo:  NewCinemaRepository(db, log),
		BookingRepo: NewBookingRepository(db, log),
		SeatRepo:    NewSeatRepository(db, log),
		PaymentRepo: NewPaymentRepository(db, log),
	}
}
