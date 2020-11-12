package main

import (
	"io"
	"net/http"
)

func showLogo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<! -- external image -- >
		<img src="https://secrethub.io/img/blog/go-gopher.png" />
	`)
}

func main() {
	http.HandleFunc("/", showLogo)
	http.ListenAndServe(":8080", nil)
}
