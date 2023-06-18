package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http://localhost:8080/static/css/styles.css
	http.ListenAndServe(":8080", nil)
}
