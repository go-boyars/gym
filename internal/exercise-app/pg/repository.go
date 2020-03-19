package pg

import (
	"context"

	"github.com/go-boyars/gym/internal/exercise-app"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewPgRepository(pool *pgxpool.Pool) (*Repository, error) {
	return &Repository{pool: pool}, nil
}

func (r *Repository) CreateExercise(e *exercise.Exercise) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) GetExercises() ([]*exercise.Exercise, error) {
	rows, err := r.pool.Query(context.Background(), "select name, muscule from exercise")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*exercise.Exercise
	for rows.Next() {
		var name, muscule string
		err = rows.Scan(&name, &muscule)
		if err != nil {
			return nil, err
		}
		result = append(result, &exercise.Exercise{Name: name, Muscule: muscule})
	}

	return result, nil
}

func (r *Repository) UpdateExercise(id int64, e *exercise.Exercise) error {
	panic("not implemented") // TODO: Implement
}

func (r *Repository) DeleteExercise(id int64) error {
	panic("not implemented") // TODO: Implement
}
