package main

import (
	"net/http"

	hand "project/handler"
)

func main() {
	http.HandleFunc("/kirim", hand.HandleRoot)
	http.HandleFunc("/chiqim", hand.HandleRoot1)
	http.ListenAndServe(":8080", nil)
}
