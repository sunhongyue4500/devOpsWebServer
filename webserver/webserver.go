package main

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>web server newwww demo</h1>"))
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":9000", nil)
}
