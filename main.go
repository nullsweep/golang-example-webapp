package main

import (
	"encoding/json"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Filename string `json:"filename"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("articles", "articles.json")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Unable to read article list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := filepath.Base(r.URL.Path)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	listPath := filepath.Join("articles", "articles.json")
	listFile, err := ioutil.ReadFile(listPath)
	if err != nil {
		http.Error(w, "Unable to read article list", http.StatusInternalServerError)
		return
	}

	var articles []Article
	if err := json.Unmarshal(listFile, &articles); err != nil {
		http.Error(w, "Error parsing article list", http.StatusInternalServerError)
		return
	}

	var selectedArticle *Article
	for _, article := range articles {
		if article.ID == id {
			selectedArticle = &article
			break
		}
	}
	if selectedArticle == nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	mdPath := filepath.Join("articles", selectedArticle.Filename)
	mdContent, err := ioutil.ReadFile(mdPath)
	if err != nil {
		http.Error(w, "Unable to read article content", http.StatusInternalServerError)
		return
	}

	htmlContent := blackfriday.Run(mdContent)

	response := map[string]string{
		"title":   selectedArticle.Title,
		"date":    selectedArticle.Date,
		"content": string(htmlContent),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/api/articles", articlesHandler)
	http.HandleFunc("/api/article/", articleHandler)
	http.HandleFunc("/", indexHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
