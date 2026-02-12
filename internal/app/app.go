package app

import (
	"context"
	"fmt"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/internal/config"
	"github.com/Taras-Rm/money-tracker-server/internal/handlers"
	"github.com/Taras-Rm/money-tracker-server/internal/services"
	"github.com/Taras-Rm/money-tracker-server/internal/setup"
	"github.com/Taras-Rm/money-tracker-server/pkg/hasher"
	"github.com/Taras-Rm/money-tracker-server/pkg/token"
	"github.com/jackc/pgx/v5"
)

func Run() {
	config := config.Config

	ctx := context.Background()

	dbConn, err := pgx.Connect(ctx, config.DbConfig.ConnectionUri)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close(ctx)

	queries := sqlc.New(dbConn)

	user, err := queries.GetUserById(ctx, 1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(user)

	hasher := hasher.NewHasher(14)
	tokenManager := token.NewTokenManager(config.AuthConfig.Secret, config.AuthConfig.Ttl)

	services := services.NewServices(services.Dependencies{
		Db:           queries,
		Hasher:       hasher,
		TokenManager: tokenManager,
	})

	handlers := handlers.NewHandlers(services)

	server := setup.NewServer(config.ServerConfig.Port, handlers.InitHandlers())

	server.Start()
}
