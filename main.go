package main

import (
	"net/http"

	hand "project/handler"
)

func main() {
	http.HandleFunc("/", hand.HandleRoot)
	http.ListenAndServe(":8080", nil)
}
