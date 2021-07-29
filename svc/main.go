package main

import (
	"context"
	"github.com/gos-apoorv/golang-web-server/httpserver"
	"github.com/gos-apoorv/golang-web-server/loggerfx"
	"github.com/gos-apoorv/golang-web-server/protobuf"
	"github.com/gos-apoorv/golang-web-server/rpcserver"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
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

func registerHooks(lifecycle fx.Lifecycle, serveMux *http.ServeMux, sugar *zap.SugaredLogger,
	rpcserver rpcserver.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				//start the RPC Server
				listener, err := net.Listen("tcp", ":8082")
				if err != nil {
					sugar.Errorf("Error while starting RPC httpserver: %v", err)
				}
				var opts []grpc.ServerOption
				gprcServer := grpc.NewServer(opts...)
				protobuf.RegisterUsersServer(gprcServer, rpcserver)
				go gprcServer.Serve(listener)

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
