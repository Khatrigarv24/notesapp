package main

import (
	"log"
	"net/http"
	"notes-app/models"
)


func main () {
	models.ConnectDB("mongodb+srv://garv3144:Hira.0612@notes.dgtsaqq.mongodb.net/?retryWrites=true&w=majority&appName=notes")
	RegisterRoutes()
	log.Println("Server started on port 8080....")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
