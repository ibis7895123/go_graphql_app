package logger

import (
	"github.com/ibis7895123/go_graphql_app/src/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	switch config.Config.ENV {
	case "production":
		Logger, _ = zap.NewProduction()
	default:
		Logger, _ = zap.NewDevelopment()
	}
}
