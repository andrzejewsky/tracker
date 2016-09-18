package server

import (
	"net/http"
	"strconv"
)

// HTTPServer http server
type HTTPServer struct {
	requestBus chan *http.Request
}

// NewHTTPServer pointer to new instance
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

// StartListening run http server
func (s *HTTPServer) StartListening(port int, requestBus chan *http.Request) {
	s.requestBus = requestBus
	http.HandleFunc("/", s.requestHandler())
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func (s *HTTPServer) requestHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		go func() { s.requestBus <- request }()
		writer.WriteHeader(200)
	}
}
