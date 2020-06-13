package exercise

type Storage interface {
	GetExercises() ([]*Exercise, error) // TODO remove pointer
}
