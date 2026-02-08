package services

import (
	"context"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/internal/services/models"
	"github.com/Taras-Rm/money-tracker-server/pkg/hasher"
)

type Users interface {
	CreateUser(ctx context.Context, user models.CreateUserInput) (interface{}, error)
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
