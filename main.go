package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("www")))
	log.Println("Listening at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
