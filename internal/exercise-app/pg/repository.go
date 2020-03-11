package pg

import (
	"context"

	"github.com/go-boyars/gym/internal/exercise-app"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	conn *pgx.Conn
}

func NewPgRepository(conn *pgx.Conn) (*Repository, error) {
	return &Repository{conn: conn}, nil
}

func (r Repository) GetExercises() ([]*exercise.Exercise, error) {
	rows, err := r.conn.Query(context.Background(), "select name, muscule from exercise")
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
