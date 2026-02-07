package services

import (
	"context"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/pkg/hasher"
)

type userService struct {
	q *sqlc.Queries

	hasher *hasher.Hasher
}

func NewUsersService(q *sqlc.Queries, hasher *hasher.Hasher) Users {
	return &userService{
		q,
		hasher,
	}
}

func (s *userService) CreateNewUser(ctx context.Context, user interface{}) (interface{}, error) {
	// hashedPassword, err := s.hasher.HashPassword(user)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
