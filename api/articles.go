package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Article struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

// ArticlesHandler serves the articles as JSON
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("articles", "articles.json")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Unable to read articles", http.StatusInternalServerError)
		return
	}

	var articles []Article
	if err := json.Unmarshal(data, &articles); err != nil {
		http.Error(w, "Error parsing articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
