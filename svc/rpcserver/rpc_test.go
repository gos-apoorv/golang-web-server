package rpcserver

import (
	"net/rpc"
	"testing"
)

func TestGetUsers(t *testing.T) {
	client, _ := rpc.Dial("tcp", "localhost:8082")

	err := client.Call("Handler.GetUsers", 1, nil)
	if err != nil {
		t.Errorf("Error:1 user.GetUsers() %+v", err)
	} else {
		t.Logf("user found")
	}
}
