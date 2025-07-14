package usecase

import (
	"context"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/internal/dto"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
)

type PaymentService interface {
	ListMethods(ctx context.Context) ([]entity.PaymentMethod, error)
	ProcessPayment(ctx context.Context, req dto.PaymentRequest) (string, error)
}

type paymentService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewPaymentService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) PaymentService {
	return &paymentService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *paymentService) ListMethods(ctx context.Context) ([]entity.PaymentMethod, error) {
	return s.Repo.PaymentRepo.GetAllMethods(ctx)
}

func (s *paymentService) ProcessPayment(ctx context.Context, req dto.PaymentRequest) (string, error) {
	return s.Repo.PaymentRepo.ProcessPayment(ctx, req)
}
