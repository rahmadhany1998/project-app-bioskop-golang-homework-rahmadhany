package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type Repository struct {
	UserRepo  UserRepository
	TokenRepo TokenRepository
}

func NewRepository(db *sql.DB, log *zap.Logger) Repository {
	return Repository{
		UserRepo:  NewUserRepository(db, log),
		TokenRepo: NewTokenRepository(db, log),
	}
}
