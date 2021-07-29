package rpcserver

import (
	"net/rpc"
	"testing"
)

func TestGetUsers(t *testing.T) {
	client, _ := rpc.Dial("tcp", "localhost:8082")

	err := client.Call("handler.GetUsers", 1, nil)

	//user := &protobuf.User{}
	//
	//err := proto.Unmarshal(err1,user)

	if err != nil {
		t.Errorf("Error:1 user.GetUsers() %+v", err)
	} else {
		t.Logf("user found")
	}
}
