package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello, world!\n")

}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		log.Println(err)
		if err != nil {
			http.Error(w, "Invasssslid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	reqString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, reqString)
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	reqString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, reqString)

	// articleID := 1
	// reqString := fmt.Sprintf("Article No.%d\n", articleID)
	// io.WriteString(w, reqString)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
