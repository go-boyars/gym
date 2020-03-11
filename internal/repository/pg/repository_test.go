package repository

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestCheckConnection(t *testing.T) {
	assert := assert.New(t)
	conn, err := pgx.Connect(context.Background(), "postgresql://boyar:go-boyars@localhost:5432/gym")
	assert.NoError(err)
	_ = conn
}
