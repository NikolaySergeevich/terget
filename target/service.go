package target

import (
	"context"
	"log/slog"
)

type Service struct {
	context context.Context
	logger  *slog.Logger
}

func NewService(ctx context.Context, logger *slog.Logger) *Service {

	return &Service{
		context: ctx,
		logger:  logger,
	}
}
