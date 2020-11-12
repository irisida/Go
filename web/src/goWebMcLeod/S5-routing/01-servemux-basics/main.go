package main

import (
	"io"
	"net/http"
)

type handlerType int

func (h handlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/one":
		io.WriteString(res, "Route 1")
	case "/two":
		io.WriteString(res, "Route 2")
	case "/three":
		io.WriteString(res, "Route 3")
	default:
		io.WriteString(res, "I am the default")
	}
}

// example lacks in elegance but it is easy see the difefrent aspects of the code we need
// for a routing capability
func main() {
	var h handlerType
	http.ListenAndServe(":8080", h)
}
