package repository

import (
	"context"
	"database/sql"
	"project-app-bioskop-golang-homework-rahmadhany/internal/data/entity"

	"go.uber.org/zap"
)

type CinemaRepository interface {
	GetAll(ctx context.Context, page, limit int) ([]entity.Cinema, error)
	GetByID(ctx context.Context, id int) (*entity.Cinema, error)
	CountAll(ctx context.Context) (int, error)
}

type cinemaRepository struct {
	db  *sql.DB
	log *zap.Logger
}

func NewCinemaRepository(db *sql.DB, log *zap.Logger) CinemaRepository {
	return &cinemaRepository{
		db:  db,
		log: log,
	}
}

func (r *cinemaRepository) GetAll(ctx context.Context, page, limit int) ([]entity.Cinema, error) {
	offset := (page - 1) * limit
	query := `SELECT id, name, location FROM cinemas ORDER BY id ASC LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var cinemas []entity.Cinema
	for rows.Next() {
		var c entity.Cinema
		if err := rows.Scan(&c.ID, &c.Name, &c.Location); err != nil {
			r.log.Error("error : ", zap.Error(err))
			return nil, err
		}
		cinemas = append(cinemas, c)
	}
	return cinemas, nil
}

func (r *cinemaRepository) GetByID(ctx context.Context, id int) (*entity.Cinema, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, location, seats_count FROM cinemas WHERE id = $1", id)
	var c entity.Cinema
	if err := row.Scan(&c.ID, &c.Name, &c.Location, &c.SeatsCount); err != nil {
		r.log.Error("error : ", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

func (r *cinemaRepository) CountAll(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM cinemas").Scan(&count)
	return count, err
}
