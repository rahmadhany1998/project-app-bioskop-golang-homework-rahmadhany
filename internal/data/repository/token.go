package repository

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
)

type TokenRepository interface {
	SaveToken(ctx context.Context, tokenID int, token string) error
	DeleteToken(ctx context.Context, token string) error
	IsTokenValid(ctx context.Context, token string) (bool, error)
}

type tokenRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewTokenRepository(db *sql.DB, log *zap.Logger) TokenRepository {
	return &tokenRepository{
		db:  db,
		log: log,
	}
}

func (r *tokenRepository) SaveToken(ctx context.Context, userID int, token string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`, userID, token)
	return err
}

func (r *tokenRepository) DeleteToken(ctx context.Context, token string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM tokens WHERE token = $1`, token)
	return err
}

func (r *tokenRepository) IsTokenValid(ctx context.Context, token string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM tokens WHERE token = $1)`, token).Scan(&exists)
	return exists, err
}
