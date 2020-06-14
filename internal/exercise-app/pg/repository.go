package pg

import (
	"context"

	"github.com/go-boyars/gym/internal/exercise-app"
	"github.com/go-boyars/gym/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewPgRepository(pool *pgxpool.Pool) (*Repository, error) {
	return &Repository{pool: pool}, nil
}

func (r Repository) GetExercises() ([]*exercise.Exercise, error) {
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

func (r Repository) CreateUser(ctx context.Context, user models.User, hash string) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO users(
			id,
			login,
			first_name,
			middle_name,
			last_name,
			sex,
			email,
			phone,
			weight,
			height,
			pwhash
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		user.ID, user.Login, user.FirstName, user.MiddleName, user.LastName, user.Sex, user.Email, user.Phone, user.Weight, user.Heigth, hash)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetPwhash(ctx context.Context, login string) (string, error) {
	var hash string
	err := r.pool.QueryRow(ctx, "SELECT pwhash FROM users WHERE login=$1", login).Scan(&hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func (r Repository) GetUserID(ctx context.Context, login string) (string, error) {
	var userID string
	err := r.pool.QueryRow(ctx, "SELECT id FROM users WHERE login=$1", login).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}
