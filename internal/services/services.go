package services

import (
	"context"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/internal/dto"
	"github.com/Taras-Rm/money-tracker-server/internal/services/models"
	"github.com/Taras-Rm/money-tracker-server/pkg/hasher"
	"github.com/Taras-Rm/money-tracker-server/pkg/oauth"
	"github.com/Taras-Rm/money-tracker-server/pkg/token"
)

type Users interface {
	CreateUser(ctx context.Context, userData models.CreateUserInput) (*dto.UserDTO, error)
	RegisterUser(ctx context.Context, userData models.CreateUserInput) (string, error)
	LoginUser(ctx context.Context, loginData models.LoginUserInput) (string, error)
	LoginUserWithGoogle(ctx context.Context, loginData models.LoginUserWithGoogleInput) (string, error)
	GetUserByID(ctx context.Context, userId int64) (*dto.UserDTO, error)
}

type Dependencies struct {
	Db *sqlc.Queries

	Hasher       *hasher.Hasher
	TokenManager *token.TokenManager
	OAuthManager *oauth.OAuthManager
}

type Services struct {
	Users Users
}

func NewServices(deps Dependencies) *Services {
	return &Services{
		Users: NewUsersService(deps.Db, deps.Hasher, deps.TokenManager, deps.OAuthManager),
	}
}
