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
	http.HandleFunc("/notes/", func (w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "PUT":
			handlers.UpdateNotes(w,r)
		case "DELETE":
			handlers.DeleteNote(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
