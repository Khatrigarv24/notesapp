package main

import (
	"net/http"
	"notes-app/handlers"
)

func RegisterRoutes() {
    http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            handlers.GetNotes(w, r)
        case "POST":
            handlers.CreateNote(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
}
