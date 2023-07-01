package handler

import (
	"net/http"
	"project/internal/service/usecase"
)

func Incoming(w http.ResponseWriter, r *http.Request) {
	usecase.Income(w, r)
}
