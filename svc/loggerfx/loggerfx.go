package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module providing SugaredLogger to fx
var Module = fx.Options(
	fx.Provide(
		func() *zap.SugaredLogger {
			logger, _ := zap.NewProduction()
			sugar := logger.Sugar()
			return sugar
		},
	),
)
