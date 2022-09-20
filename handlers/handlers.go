package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/takuya-yone/golang-api-sandbox/models"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello, world!\n")

}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Posting Article...\n")

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}

	w.Write(jsonData)

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

	articleList := []models.Article{models.Article1, models.Article2}

	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("ail to encode json (page %d)\n", page)

		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	w.Write(jsonData)

}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {

		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	// article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Posting Comment...\n")
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
