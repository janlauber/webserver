package main

//https://gowebexamples.com/

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Yo waddup")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/staic/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
