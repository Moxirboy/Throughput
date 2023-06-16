package main

import (
	"log"
	"net/http"
	rout "project/internal/controller"
)

func main() {
	rout.Router()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
