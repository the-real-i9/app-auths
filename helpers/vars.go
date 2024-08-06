package helpers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrInternalServerError = errors.New("internal server error: check logger")

var dbPool *pgxpool.Pool
