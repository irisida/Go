package main

import (
	"io"
	"net/http"
	"os"
)

func showLogo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<! -- external image -- >
		<img src="https://secrethub.io/img/blog/go-gopher.png" />
	`)
}

func showLocal(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("go.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", showLogo)
	http.HandleFunc("/go", showLocal)
	http.ListenAndServe(":8080", nil)
}
