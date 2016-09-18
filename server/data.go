package server

import (
	"io"
	"net/http"
	"net/url"
)

// Request representation
type Request struct {
	Method     string
	URL        *url.URL
	Header     http.Header
	Body       io.ReadCloser
	Host       string
	Form       url.Values
	PostForm   url.Values
	Trailer    http.Header
	RemoteAddr string
	RequestURI string
}

// CreateByRequest create Request based on *http.Request
func CreateByRequest(request *http.Request) *Request {
	return &Request{
		request.Method,
		request.URL,
		request.Header,
		request.Body,
		request.Host,
		request.Form,
		request.PostForm,
		request.Trailer,
		request.RemoteAddr,
		request.RequestURI,
	}
}
