package repository

import (
	"context"
	"database/sql"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"

	"go.uber.org/zap"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type userRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewUserRepository(db *sql.DB, log *zap.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	return err
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE username = $1`
	row := r.db.QueryRowContext(ctx, query, username)
	var user entity.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
