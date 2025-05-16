package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/models"   // make sure package name is models
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	result, err := models.CreateNote(note)
	if err != nil {
		http.Error(w, "failed to make note", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := models.GetAllNotes()
	if err != nil {
		http.Error(w, "failed to get notes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

