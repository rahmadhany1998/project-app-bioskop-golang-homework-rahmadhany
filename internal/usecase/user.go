package usecase

import (
	"context"
	"errors"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/repository"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/codes"
	"project-app-bioskop-golang-homework-rahmadhany/pkg/utils"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, username, password string) (*entity.User, string, error)
	Logout(ctx context.Context, token string) error
}

type userService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewUserService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) UserService {
	return &userService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *userService) Register(ctx context.Context, user *entity.User) error {
	existing, _ := s.Repo.UserRepo.GetUserByUsername(ctx, user.Username)
	if existing != nil {
		s.Logger.Warn("Username already exists", zap.String("username", user.Username))
		return errors.New("username already exists")
	}
	hashed, err := codes.GeneratePassword(user.Password)
	if err != nil {
		s.Logger.Error("Failed to hash password", zap.Error(err))
		return err
	}
	user.Password = *hashed
	err = s.Repo.UserRepo.CreateUser(ctx, user)
	if err != nil {
		s.Logger.Error("Failed to create user", zap.Error(err))
	}
	return err
}

func (s *userService) Login(ctx context.Context, username, password string) (*entity.User, string, error) {
	user, err := s.Repo.UserRepo.GetUserByUsername(ctx, username)
	if err != nil {
		s.Logger.Warn("User not found", zap.String("username", username))
		return nil, "", errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.Logger.Warn("Incorrect password", zap.String("username", username))
		return nil, "", errors.New("invalid username or password")
	}
	token, err := utils.GenerateJWT(user.ID, user.Username, s.Config)
	if err != nil {
		s.Logger.Error("Failed to generate token", zap.Error(err))
		return nil, "", err
	}
	if err := s.Repo.TokenRepo.SaveToken(ctx, user.ID, token); err != nil {
		s.Logger.Error("Failed to save token", zap.Error(err))
	}
	return user, token, nil
}

func (s *userService) Logout(ctx context.Context, token string) error {
	err := s.Repo.TokenRepo.DeleteToken(ctx, token)
	if err != nil {
		s.Logger.Error("Failed to delete token", zap.Error(err))
	}
	return err
}
