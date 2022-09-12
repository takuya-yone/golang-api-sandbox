package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	log.Println(r.Method)

}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Article List\n")
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID := 1
	reqString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, reqString)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
