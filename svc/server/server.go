package server

import "net/http"

//Handler for HTTP request
type Handler struct {
	serveMux *http.ServeMux
}

// New HTTP handler
func New(mux *http.ServeMux) *Handler {
	handler := Handler{serveMux: mux}
	handler.registerEndpoints()
	return &handler
}

//registerEndpoints for all HTTP endpoints
func (handler *Handler) registerEndpoints() {
	handler.serveMux.HandleFunc("/", handler.RenderMessage)
}

//RenderMessage to render the desired message into browser
func (handler *Handler) RenderMessage(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("Test Hello World Sample!!!"))
}
