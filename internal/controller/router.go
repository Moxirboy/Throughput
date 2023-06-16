package controller

import (
	"net/http"
	hand "project/internal/controller/v1/handler"
)

func Router() {
	http.HandleFunc("/kirim", hand.Incoming)
	http.HandleFunc("/chiqim", hand.Outcoming)
}
