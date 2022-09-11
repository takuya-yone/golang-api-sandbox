package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
