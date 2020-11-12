package main

import (
	"fmt"
	"net/http"
)

// create a type
type parameter int

// add the compliance to the handler interface to the type declared
func (p parameter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "A response from a function impelmenting the handler interface")
}

func main() {

	// create a varoable of our type that implements the handler interface
	var p parameter

	// call ListenAndServe with the port as address and the object implementing handler as the handler
	http.ListenAndServe(":8080", p)
}
