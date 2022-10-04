package redis

import (
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

func InferCodeFromError(err error) codes.Code {
	if errors.Is(err, redis.Nil) {
		return codes.NotFound
	}
	return codes.Internal
}
