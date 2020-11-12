package main

import (
	"fmt"
	"net/http"
)

type handlerType int

func (h handlerType) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Ed-Key", "this is from ed")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Any code you want can go in here</h1>")
}

func main() {
	var h handlerType
	http.ListenAndServe(":8080", h)
}
