package repository

import (
	exercise "github.com/go-boyars/gym/internal/repository"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	conn *pgx.Conn
}

func NewPgRepository(conn *pgx.Conn) (*Repository, error) {
	return &Repository{conn: conn}, nil
}

func GetExercises() (*exercise.Exercise, error) {
	return nil, nil
}
