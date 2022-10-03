package postgres

import (
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func InferCodeFromError(err error) codes.Code {
	// gorm errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return codes.NotFound
	}

	// pg errors
	pgError := &pgconn.PgError{}
	if errors.As(err, &pgError) && pgError.Code == "23505" {
		// unique key violation
		return codes.AlreadyExists
	}
	return codes.Internal
}
