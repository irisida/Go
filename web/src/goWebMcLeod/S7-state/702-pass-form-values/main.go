package main

import (
	"io"
	"net/http"
)

/*
	On running the server: go to localhost:8080/
	To see the query handled go to localhost:8080/q=myvalue
	Note we use the post method meaning the data is sent through
	the body and not the url(get)
*/

func home(w http.ResponseWriter, req *http.Request) {
	val := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+val)
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
