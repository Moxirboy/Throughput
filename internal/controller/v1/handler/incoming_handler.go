package handler

import (
	"html/template"
	"net/http"
	"project/internal/controller/v1/handler/adapter"
	repo "project/internal/service/repo/mysql"
)

func Incoming(w http.ResponseWriter, r *http.Request) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered from panic:", r)
	//		// Perform any necessary cleanup or logging here
	//	}
	//}()
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/incoming.html"))
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "incoming.html", nil)
		return
	}
	adapter.FormValues(r)
	repo.InsertClientQuery()
	repo.InsertProductQuery()
	repo.InsertPurchaseQuery()
	repo.InsertPurchaseGoodsQuery()

	tmpl.Execute(w, struct{ Success bool }{true})
}
