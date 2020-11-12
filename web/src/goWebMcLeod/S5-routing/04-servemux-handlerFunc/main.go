package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the home handler")
}

func one(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the route 1 handler")
}

func two(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I am the route 2 handler")
}

func main() {
	// HandleFunc version

	http.HandleFunc("/", home)
	http.HandleFunc("/one", one)
	http.HandleFunc("/two", two)

	http.ListenAndServe(":8080", nil) // nil denotes the default servemux
}
