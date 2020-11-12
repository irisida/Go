package main

import (
	"io"
	"net/http"
)

// create a type we will use as a handler pt1
type homeHandlerType int

// implement the function required to adhere to handler interface pt1
func (h homeHandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the home handler")
}

// create a type we will use as a handler pt2
type route1HandlerType int

// implement the function required to adhere to handler interface pt2
func (h route1HandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the route 1 handler")
}

// create a type we will use as a handler pt3
type route2HandlerType int

// implement the function required to adhere to handler interface pt3
func (h route2HandlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the route 2 handler")
}

func main() {
	var h homeHandlerType
	var h1 route1HandlerType
	var h2 route2HandlerType

	/*
		note we have removed the declaration of the ServeMux in favour
		of using the default mux in Go.
	*/

	http.Handle("/", h)
	http.Handle("/one", h1)
	http.Handle("/two", h2)

	http.ListenAndServe(":8080", nil)
}
