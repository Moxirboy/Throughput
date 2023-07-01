package handler

import (
	"net/http"
	"project/internal/service/usecase"
)

func Outcome(w http.ResponseWriter, r *http.Request) {
	usecase.OutCome(w, r)
}
