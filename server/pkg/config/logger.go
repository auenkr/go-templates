package config

import "go.uber.org/zap"

func NewLogger(config *Config) (*zap.Logger, error) {
	if config.Environment == config.PROD_ENVIRONMENT {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
