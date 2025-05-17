package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/models"   // make sure package name is models
	"strings"
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

func UpdateNotes(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/notes/")

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.UpdateNote(id, updatedNote); err != nil {
		http.Error(w, "Error saving note", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Note updated successfully"))
}

func DeleteNote (w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/notes/")
	if err := models.DeleteNote(id); err != nil {
		http.Error(w, "Unable to delete the note", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Note deleted successfully"))
}
