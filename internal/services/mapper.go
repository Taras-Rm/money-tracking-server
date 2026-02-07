package services

import (
	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/internal/dto"
)

func ToUserDTO(u sqlc.User) dto.UserDTO {
	return dto.UserDTO{
		ID:        int64(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
	}
}
