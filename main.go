package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string = "эщкере"

func main() {
	http.HandleFunc("/", handleGet)
	http.HandleFunc("/update", handlePost)

	fmt.Println("сервер запущен")
	http.ListenAndServe(":8080", nil)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "не использован GET", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "привет, %s", task)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "не использован POST", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Task string `json:"task"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "неправильный JSON", http.StatusBadRequest)
		return
	}

	task = request.Task

	fmt.Fprintf(w, "задача обновлена: %s", task)
}
