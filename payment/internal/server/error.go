package server

import (
	"github.com/omise/omise-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

func InferCodeFromOmiseError(err error) codes.Code {
	omiseError := &omise.Error{}
	switch errors.As(err, &omiseError); omiseError.StatusCode {
	case 400:
		return codes.InvalidArgument
	case 403:
		return codes.FailedPrecondition
	case 404:
		return codes.NotFound
	default:
		return codes.Unavailable
	}
}
