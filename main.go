package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "How to Study Programming For Beginner", Desc: "Many newbie programmers get overhelmed when started Programming", Content: "Tech"},
		Article{Id: "2", Title: "Make Money From Youtube", Desc: "People Very Obeseve To get money from youtube", Content: "Social Media"},
		Article{Id: "3", Title: "How to Study Programming For Beginner", Desc: "Many newbie programmers get overhelmed when started Programming", Content: "Tech"},
		Article{Id: "4", Title: "How to Study Programming For Beginner", Desc: "Many newbie programmers get overhelmed when started Programming", Content: "Tech"},
	}
	handleRequests()
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit :return some Articles")
	json.NewEncoder(w).Encode(Articles)
}
