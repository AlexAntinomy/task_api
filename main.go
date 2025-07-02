package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createTask(w, r)
			return
		}
		http.NotFound(w, r)
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTask(w, r)
		case http.MethodDelete:
			deleteTask(w, r)
		default:
			http.NotFound(w, r)
		}
	})
	http.ListenAndServe(":8080", nil)
}
