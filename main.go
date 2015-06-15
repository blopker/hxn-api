package main

import (
	"encoding/json"
	"net/http"

	"github.com/blopker/hxn-api/lru"
)

var conf Config
var cache *lru.Cache

func main() {
	conf.Load()
	cache, _ = lru.New(512)
	http.HandleFunc("/", articleHandler)
	http.ListenAndServe(":"+conf.Port, nil)
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
		w.Write([]byte("Please supply a 'url' query paremeter"))
		return
	}

	url := urls[0]
	article, err := getArticle(url)

	if err != nil {
		http.Error(w, "Article "+url+" not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(article)
}

func getArticle(url string) (article, error) {
	var a article

	if a, ok := cache.Get(url); ok {
		return a.(article), nil
	}

	resp, err := http.Get("https://www.readability.com/api/content/v1/parser?token=" + conf.Token + "&url=" + url)

	if err != nil {
		return article{}, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return article{}, err
	}

	cache.Add(url, a)
	return a, nil
}
