package exercise

import (
	"context"

	"github.com/go-boyars/gym/internal/models"
)

type Storage interface {
	CreateUser(context.Context, models.User, string) error
	GetExercises() ([]*Exercise, error) // TODO remove pointer
}
