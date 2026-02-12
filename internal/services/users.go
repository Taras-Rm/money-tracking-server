package services

import (
	"context"
	"errors"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/internal/dto"
	"github.com/Taras-Rm/money-tracker-server/internal/services/models"
	"github.com/Taras-Rm/money-tracker-server/pkg/hasher"
	"github.com/Taras-Rm/money-tracker-server/pkg/token"
	"github.com/jackc/pgx/v5"
)

type userService struct {
	q *sqlc.Queries

	hasher       *hasher.Hasher
	tokenManager *token.TokenManager
}

func NewUsersService(q *sqlc.Queries, hasher *hasher.Hasher, tokenManager *token.TokenManager) Users {
	return &userService{
		q,
		hasher,
		tokenManager,
	}
}

func (s *userService) CreateUser(ctx context.Context, userData models.CreateUserInput) (*dto.UserDTO, error) {
	hashedPassword, err := s.hasher.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	_, err = s.q.GetUserByEmail(ctx, userData.Email)
	if err == nil {
		return nil, errors.New("user we such email already exists")
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	user, err := s.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: hashedPassword,
	})

	userDTO := ToUserDTO(user)

	return &userDTO, nil
}

func (s *userService) LoginUser(ctx context.Context, loginData models.LoginUserInput) (string, error) {
	user, err := s.q.GetUserByEmail(ctx, loginData.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	isPasswordCorrrect := s.hasher.VerifyPassword(loginData.Password, user.Password)
	if !isPasswordCorrrect {
		return "", errors.New("wrong password")
	}

	token, err := s.tokenManager.NewToken(int64(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
