package rpcserver

import (
	"context"
	"github.com/gos-apoorv/golang-web-server/protobuf"
)

type Handler  = protobuf.UsersServer

//handler interface exposes the function of User Server Method
type handler struct {
	protobuf.UnimplementedUsersServer
}

//New returns the RPC handler
func New() (Handler, error) {
	return &handler{},nil
}

//GetUsers return the list of users
func (h *handler) GetUsers(ctx context.Context, req *protobuf.EmptyReq) (*protobuf.GetUsersResponse,error) {
	return &protobuf.GetUsersResponse{
		Users: []*protobuf.User{
			{
				Name: "test user",
				Age: 11,
			},
		},
	}, nil
}
