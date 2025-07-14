package usecase

import (
	"context"
	"math"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
)

type CinemaService interface {
	GetAll(ctx context.Context, page, limit int) ([]entity.Cinema, int, int, error)
	GetByID(ctx context.Context, id int) (*entity.Cinema, error)
	GetSeatStatus(ctx context.Context, cinemaID int, date, time string) ([]entity.Seat, error)
}

type cinemaService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewCinemaService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) CinemaService {
	return &cinemaService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *cinemaService) GetAll(ctx context.Context, page, limit int) ([]entity.Cinema, int, int, error) {
	if page < 1 {
		page = 1
	}

	totalRecords, err := s.Repo.CinemaRepo.CountAll(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	cinemas, err := s.Repo.CinemaRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return cinemas, totalRecords, totalPages, nil
}

func (s *cinemaService) GetByID(ctx context.Context, id int) (*entity.Cinema, error) {
	return s.Repo.CinemaRepo.GetByID(ctx, id)
}

func (s *cinemaService) GetSeatStatus(ctx context.Context, cinemaID int, date, time string) ([]entity.Seat, error) {
	return s.Repo.SeatRepo.GetSeatStatusBySchedule(ctx, cinemaID, date, time)
}
