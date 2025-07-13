package usecase

import (
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
)

type Service struct {
	UserService   UserService
	CinemaService CinemaService
}

func NewService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) Service {
	return Service{
		UserService:   NewUserService(repo, logger, config),
		CinemaService: NewCinemaService(repo, logger, config),
	}
}
