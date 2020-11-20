package main

import (
	"io"
	"net/http"
)

/*
	on running the server: go to localhost:8080/ to see
	to see the query handled go to localhost:8080/q=myvalue
	sample new line to test git integration with vim
*/

func home(w http.ResponseWriter, req *http.Request) {
	val := req.FormValue("q")
	io.WriteString(w, "Do my search: "+val)
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
