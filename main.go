package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/takuya-yone/golang-api-sandbox/handlers"
	"github.com/takuya-yone/golang-api-sandbox/sandbox"
)

func main() {

	sandbox.DB_conn()

	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
