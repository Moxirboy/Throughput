package main

import (
	"log"
	"net/http"
	hand "project/internal/controller/v1/handler"
)

func main() {
	http.HandleFunc("/kirim", hand.Incoming)
	http.HandleFunc("/chiqim", hand.Outcoming)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
