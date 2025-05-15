package models

import (
	"encoding/json"
	"os"
)

type Note struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Notes []Note = LoadNotes()

func SaveNotes() error {
	file, err := os.Create("notes.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(Notes)
}

func LoadNotes() []Note {
	file, err := os.Open("notes.json")
	if err != nil {
		return []Note{}
	}
	defer file.Close()

	var notes []Note
	err = json.NewDecoder(file).Decode(&notes)
	if err != nil {
		return []Note{}
	}
	return notes
}

