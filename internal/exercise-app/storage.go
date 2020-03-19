package exercise

type Storage interface {
	CreateExercise(e *Exercise) (int64, error)
	GetExercises() ([]*Exercise, error)
	UpdateExercise(id int64, e *Exercise) error
	DeleteExercise(id int64) error
}
