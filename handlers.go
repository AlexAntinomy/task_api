package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func createTask(w http.ResponseWriter, r *http.Request) {
	t := NewTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	t := GetTask(id)
	if t == nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	DeleteTask(id)
	w.WriteHeader(http.StatusNoContent)
}
