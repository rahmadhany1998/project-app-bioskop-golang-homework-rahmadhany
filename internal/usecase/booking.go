package usecase

import (
	"context"
	"errors"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/internal/dto"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BookingService interface {
	CreateBooking(ctx context.Context, req dto.BookingRequest, userID int) (*entity.Booking, error)
}

type bookingService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewBookingService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) BookingService {
	return &bookingService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, req dto.BookingRequest, userID int) (*entity.Booking, error) {
	seatID, err := s.Repo.SeatRepo.GetSeatIDByCode(ctx, req.CinemaID, req.SeatID)
	if err != nil {
		return nil, err
	}
	booked, err := s.Repo.BookingRepo.IsSeatBooked(ctx, req.CinemaID, seatID, req.Date, req.Time)
	if err != nil {
		return nil, err
	}
	if booked {
		return nil, errors.New("seat already booked")
	}

	booking := &entity.Booking{
		ID:            uuid.NewString(),
		UserID:        userID,
		CinemaID:      req.CinemaID,
		SeatID:        seatID,
		Date:          req.Date,
		Time:          req.Time,
		PaymentMethod: req.PaymentMethod,
		Status:        "pending",
	}
	if err := s.Repo.BookingRepo.CreateBooking(ctx, booking); err != nil {
		return nil, err
	}
	return booking, nil
}
