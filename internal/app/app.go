package app

import (
	"context"
	"fmt"

	"github.com/Taras-Rm/money-tracker-server/db/sqlc"
	"github.com/Taras-Rm/money-tracker-server/internal/config"
	"github.com/Taras-Rm/money-tracker-server/internal/setup"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	server := setup.NewServer(config.ServerConfig.Port, router)

	server.Start()
}
