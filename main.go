package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
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
	indexData.Title = "Hello World"
	indexData.Desc = "This is a Blog"
	jsonStr, _ := json.Marshal(indexData)
	resp.Write(jsonStr)
}
