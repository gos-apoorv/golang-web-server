package rpcserver

import "net/rpc"

//User struct for message
type User struct {
	Name string
	Age  int32
}

//Handler interface definition
type Handler struct {
}

//New returns the RPC handler
func New() *Handler {
	handler := &Handler{}
	err := rpc.Register(handler)
	if err != nil {
		panic(err)
	}
	return handler
}

//GetUsers return the list of users
func (handler *Handler) GetUsers(payload int, reply *string) error {
	// add logic to return users
	return nil
}
