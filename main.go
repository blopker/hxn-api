package main

import (
	"encoding/json"
	"lru"
	"net/http"
)

var conf Config
var cache lru.Cache

func main() {
	conf.Load()
	cache = lru.New(512)
	http.HandleFunc("/", articleHandler)
	http.ListenAndServe(":8000", nil)
}

type article struct {
	Conent string `json:"content"`
	Domain string `json:"domain"`
	URL    string `json:"url"`
	Author string `json:"author"`
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["url"]

	if urls == nil {
		http.Error(w, "Please supply a 'url' query paremeter", http.StatusNotFound)
		return
	}

	url := urls[0]
	article, err := getArticle(url)

	if err != nil {
		http.Error(w, "Article "+url+" not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(article)
}

func getArticle(url string) (article, error) {

	resp, err := http.Get("https://www.readability.com/api/content/v1/parser?token=" + conf.Token + "&url=" + url)

	if err != nil {
		return article{}, err
	}

	defer resp.Body.Close()

	var a article

	if err := json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return article{}, err
	}

	cache.Add(url, a)
	return a, nil
}
