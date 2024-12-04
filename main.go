package main

import (
	"context"
	"log/slog"
	"os"
	"target/target"
	"target/config"

	sloggin "github.com/samber/slog-gin"

	"github.com/gin-gonic/gin"
)

// @title Сервис управления товарами
// @version 2.0
// @description Сервис предоставляет функции для управления отобранными товарами.
// @BasePath /target/v1/
func main() {
	var serviceConfig config.ServiceConfiguration

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := config.MustConfig(logger, &serviceConfig)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	router := gin.Default()

	router.Use(sloggin.New(logger))
	router.Use(gin.Recovery())

	service := target.NewService(ctx, logger)

	httpTransport := target.NewHTTPTransport(ctx, logger, router, service)
	httpTransport.RegisterRoutes()

	httpTransport.ListenAndServe(serviceConfig.ServerHost, serviceConfig.ServerPort)
}