package server

import (
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) toProtoUsers(shopOwners []models.ShopOwner) []*shopproto.User {
	protoUsers := make([]*shopproto.User, len(shopOwners))
	for _, owner := range shopOwners {
		protoUsers = append(protoUsers, &shopproto.User{
			Id:       owner.User.Id,
			Email:    owner.User.Email,
			Role:     string(owner.User.Role),
			GoogleId: owner.User.GoogleId,
		})
	}

	return protoUsers
}

func (s *Server) toShopOwners(emails []string) ([]models.ShopOwner, error) {
	owners := make([]models.ShopOwner, len(emails))
	for index, email := range emails {
		user, err := s.repo.GetUserByEmail(email)

		// out, err := json.Marshal(user)
		// if err != nil {
		// }
		// logger.Info(string(out))

		if err != nil {
			return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't get user, "+email).Error())
		}
		// FIXME since there is no shopId, the db mutatuons e.g. deletion in EditShopOwners() won't work
		owners[index] = models.ShopOwner{
			UserId: user.Id,
			User:   *user,
		}
	}
	return owners, nil
}
