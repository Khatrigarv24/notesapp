package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/models"   // make sure package name is models
	"github.com/google/uuid"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote models.Note  // Note with capital N, matching exported struct in models
	err := json.NewDecoder(r.Body).Decode(&newNote) // r.Body (uppercase B)
	if err != nil {
		http.Error(w, "Note not in correct format", http.StatusBadRequest)
		return
	}

	newNote.Id = uuid.New().String()   // fix capitalization here
	models.Notes = append(models.Notes, newNote) // use capital N for Notes slice

	if err := models.SaveNotes(); err != nil {  // SaveNotes (plural), match function name in models
		http.Error(w, "Unable to save note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNote)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Notes)  // models.Notes slice exported
}

