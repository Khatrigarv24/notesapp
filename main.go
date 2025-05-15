package main

import (
	"fmt"
	"log"
	"net/http"
)


func main () {
	SetupRoutes()
	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
