package services

import (
	"context"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/pkg/hasher"
)

type Users interface {
	CreateNewUser(ctx context.Context, user interface{}) (interface{}, error)
}

type Dependencies struct {
	Db *sqlc.Queries

	Hasher *hasher.Hasher
}

type Services struct {
	Users Users
}

func NewServices(deps Dependencies) *Services {
	return &Services{
		Users: NewUsersService(deps.Db, deps.Hasher),
	}
}
