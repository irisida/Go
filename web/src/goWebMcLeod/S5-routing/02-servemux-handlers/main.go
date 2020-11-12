package main

import (
	"io"
	"net/http"
)

type homeHandlerType int

func (h homeHandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the home handler")
}

type route1HandlerType int

func (h route1HandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the route 1 handler")
}

type route2HandlerType int

func (h route2HandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the route 2 handler")
}

func main() {
	var h homeHandlerType
	var h1 route1HandlerType
	var h2 route2HandlerType

	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/one", h1)
	mux.Handle("/two", h2)

	http.ListenAndServe(":8080", mux)
}
