package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Panicln("Error")
	}

}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "Hello"
	indexData.Desc = "Blog"
	jsonStr, err := json.Marshal(indexData)
	if err != nil {
		log.Panicln("Error marshaling JSON:", err)
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	resp.Write(jsonStr)
}

func indexHtml(resp http.ResponseWriter, req *http.Request) {
	var indexData IndexData
	indexData.Title = "Hello"
	indexData.Desc = "This is"
	t := template.New("index.html")
	path, err := os.Getwd()
	if err != nil {
		log.Panicln("Error getting working directory:", err)
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if t, err = t.ParseFiles(path + "/template/index.html"); err != nil {
		log.Println("Error parsing template:", err)
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err = t.Execute(resp, indexData); err != nil {
		log.Println("Error executing template:", err)
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
