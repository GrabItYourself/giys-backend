package authutils

import (
	"encoding/json"

	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const (
	IdentityKey = "AUTH_IDENTITY"
)

type Identity struct {
	UserId string          `json:"userId"`
	Role   models.RoleEnum `json:"role"`
}

// ExtractIdentityFromGrpcContext extracts userId and role from metadata in gRPC context
func ExtractIdentityFromGrpcContext(ctx context.Context) (*Identity, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("can't get metadata from context")
	}

	identityJSON := md.Get(IdentityKey)
	if len(identityJSON) == 0 {
		return nil, errors.Errorf("key %s is empty", IdentityKey)
	} else if len(identityJSON) > 1 {
		return nil, errors.Errorf("key %s has more than one value", IdentityKey)
	}

	var identity Identity
	if err := json.Unmarshal([]byte(identityJSON[0]), &identity); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal identity from context")
	}
	if identity.UserId == "" {
		return nil, errors.Errorf("User is empty in identity '%s'", identityJSON[0])
	}
	if identity.Role == "" {
		return nil, errors.Errorf("Role is empty in identity '%s'", identityJSON[0])
	}

	return &identity, nil
}

func EmbedIdentityToContext(ctx context.Context, identity *Identity) (context.Context, error) {
	identityJSON, err := json.Marshal(identity)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal identity to json")
	}

	return metadata.AppendToOutgoingContext(ctx, IdentityKey, string(identityJSON)), nil
}
