package rpcserver

import (
	"context"
	"github.com/gos-apoorv/golang-web-server/protobuf"
	"google.golang.org/protobuf/proto"
	"net/rpc"
	"testing"
)

func TestGetUsers(t *testing.T) {
	client, _ := rpc.Dial("tcp", "localhost:8082")

	err1 := client.Call("handler.GetUsers", context.Context(), protobuf.EmptyReq{})

	user := &protobuf.User{}

	err = proto.Unmarshal(err1,user)

	if err != nil {
		t.Errorf("Error:1 user.GetUsers() %+v", err)
	} else {
		t.Logf("user found")
	}
}
