package main

import (
	"context"
	"net"
	"net/http"
	"net/rpc"

	"github.com/gos-apoorv/golang-web-server/httpserver"
	"github.com/gos-apoorv/golang-web-server/loggerfx"
	"github.com/gos-apoorv/golang-web-server/rpcserver"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {

	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Provide(rpcserver.New),
		fx.Invoke(httpserver.New),
		fx.Invoke(registerHooks),
		loggerfx.Module,
	).Run()

}

func registerHooks(lifecycle fx.Lifecycle, serveMux *http.ServeMux, sugar *zap.SugaredLogger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				//start the RPC Server
				listener, err := net.Listen("tcp", ":8082")
				if err != nil {
					sugar.Errorf("Error while starting RPC httpserver: %v", err)
				}
				go func() {
					for {
						rpc.Accept(listener)
					}
				}()
				sugar.Info("RPC Server running on localhost:8082")

				//start the HTTP Server
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
