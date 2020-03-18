package pg

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestGetExercises(t *testing.T) {
	assert := assert.New(t)
	pool, err := pgxpool.Connect(context.Background(), "postgresql://boyar:go-boyars@localhost:5432/gym")
	assert.NoError(err)
	defer pool.Close()

	type testExercise struct {
		name    string
		muscule string
	}
	expExercises := []testExercise{
		testExercise{name: "first", muscule: "brain"},
		testExercise{name: "second", muscule: "tongue"},
		testExercise{name: "third", muscule: "ass"},
	}

	// 1. populate db with exercises (3)
	_, err = pool.Exec(
		context.Background(),
		"insert into exercise (name, muscule) values ($1, $2),($3, $4),($5, $6)",
		expExercises[0].name,
		expExercises[0].muscule,
		expExercises[1].name,
		expExercises[1].muscule,
		expExercises[2].name,
		expExercises[2].muscule,
	)
	assert.NoError(err)

	// 2. call GetExercises
	r, err := NewPgRepository(pool)
	assert.NoError(err)
	exercises, err := r.GetExercises()
	assert.NoError(err)

	// 3. check result of step 2 (3)
	flag := 0
	for _, exercise := range exercises {
		for _, expExercise := range expExercises {
			if exercise.Name == expExercise.name && exercise.Muscule == expExercise.muscule {
				flag++
				continue
				// expExercises = append(expExercises[:i], expExercises[i+1:]...)
			}

		}
	}
	assert.Equal(len(expExercises), flag, "Inconsistancy in db is found")

	// 4. delete from exercises
	_, err = pool.Exec(
		context.Background(),
		"delete from exercise",
	)
	assert.NoError(err)
}
