package config

import (
	"fmt"
	"log/slog"

	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/kelseyhightower/envconfig"
)

type ServiceConfiguration struct {
	ServerHost            string `envconfig:"SERVER_HOST"`
	ServerPort            string `envconfig:"SERVER_PORT"`
	TimeoutInMilliseconds int    `envconfig:"TIMEOUT_IN_MILLISECONDS"`
}

func MustConfig(logger *slog.Logger, serviceConfig *ServiceConfiguration) error {
	err := envconfig.Process("", serviceConfig)
	if err != nil {
		logger.Error(ERROR_FAILED_TO_PROCESS_SERVICE_CONFIGS, "msg", err.Error())
		return fmt.Errorf("%s", ERROR_FAILED_TO_PROCESS_SERVICE_CONFIGS)
	}
	return nil
}