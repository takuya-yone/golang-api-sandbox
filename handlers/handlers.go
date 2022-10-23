package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/takuya-yone/golang-api-sandbox/models"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Health Check, OK!\n")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello, world!\n")

}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {

	var reqArticle models.Article

	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle

	json.NewEncoder(w).Encode(article)

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

	json.NewEncoder(w).Encode(articleList)
	log.Println(page)

}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {

		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1

	json.NewEncoder(w).Encode(article)
	log.Println(articleID)

}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article1

	json.NewEncoder(w).Encode(article)

}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment1

	json.NewEncoder(w).Encode(comment)

}
