package main

import (
	"context"
	"go.uber.org/fx"
	"net/http"
)
import "github.com/gos-apoorv/golang-web-server/server"

func main() {

	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Invoke(server.New),
		fx.Invoke(registerHooks),
	).Run()

}

func registerHooks(lifecycle fx.Lifecycle, serveMux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go http.ListenAndServe(":8080", serveMux)
				return nil
			},
		},
	)
}
