package main

import (
	"net/http"
)

/*
	special case is the presence of index.html
*/

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
