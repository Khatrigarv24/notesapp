package main

import (
	"net/http"
	"notes-app/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			handlers.GetNotes(w,r)
		}else if r.Method == http.MethodPost{
			handlers.CreateNote(w,r)
		}else {
			http.Error(w, "bad method", http.StatusMethodNotAllowed)
		}
	})
}
