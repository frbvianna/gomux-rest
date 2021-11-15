package main

import (
	"encoding/json"
	"fmt"
	art "go-rest-api/articles"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage!")
	fmt.Println("hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit: return all articles")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	_ = enc.Encode(art.Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range art.Articles {
		if article.Id == key {
			enc := json.NewEncoder(w)
			enc.SetIndent("", "    ")
			_ = enc.Encode(article)
		}
	}

	fmt.Println("hit: return article ID ", key)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	article := new(art.Article)

	reqBody, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	_ = json.Unmarshal(reqBody, article)

	art.Articles = append(art.Articles, *article)

}

func handleRequests() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/all", returnAllArticles)
	muxRouter.HandleFunc("/article/{id}", returnSingleArticle)
	muxRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	art.Articles = []art.Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
