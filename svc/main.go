package main

import (
	"context"
	"github.com/gos-apoorv/golang-web-server/loggerfx"
	"github.com/gos-apoorv/golang-web-server/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func main() {

	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Invoke(server.New),
		fx.Invoke(registerHooks),
		loggerfx.Module,
	).Run()

}

func registerHooks(lifecycle fx.Lifecycle, serveMux *http.ServeMux,sugar *zap.SugaredLogger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				sugar.Info("Listening on localhost:8080")
				go http.ListenAndServe(":8080", serveMux)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return sugar.Sync()
			},
		},
	)
}
