package main

import (
	"net/http"
	"os"
)

func home(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("home.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

func one(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "one.png")
}

func two(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("two.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/one", one)
	http.HandleFunc("/two", two)
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}
