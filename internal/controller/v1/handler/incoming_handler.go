package handler

import (
	"html/template"
	"net/http"
	"project/internal/controller/v1/handler/adapter"
	repo "project/internal/service/repo/mysql"
)

func Incoming(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/incoming.html"))

	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "incoming.html", nil)
		return
	}
	r.ParseForm()
	adapter.FormValues(r)
	repo.InsertClientQuery()
	repo.InsertProductQuery()
	repo.InsertPurchaseQuery()
	repo.InsertPurchaseGoodsQuery()

	tmpl.Execute(w, struct{ Success bool }{true})
}
