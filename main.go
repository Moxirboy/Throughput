package main

import (
	"net/http"

	hand "project/handler"
)

func main() {
	http.HandleFunc("/", hand.HandleRoot)
	http.HandleFunc("/client", hand.HandleRoot1)
	http.ListenAndServe(":8080", nil)
}
