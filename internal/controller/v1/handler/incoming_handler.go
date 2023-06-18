package handler

import (
	"html/template"
	"net/http"
	config "project/internal/configs"
	"project/internal/controller/v1/handler/adapter"
	"project/internal/service/repo/mysql"
)

var (
	Db, _ = config.DB()
)

func Incoming(w http.ResponseWriter, r *http.Request) {
	//defer func() {
	//	if err := recover(); r != nil {
	//		fmt.Println("Recovered from panic:", err)
	//		// Perform any necessary cleanup or logging here
	//	}
	//}()
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/incoming.html"))
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "incoming.html", nil)
		return
	}
	adapter.FormValues(r)
	mysql.InsertClientQuery()
	mysql.InsertProductQuery()
	mysql.InsertPurchaseQuery()
	mysql.InsertPurchaseGoodsQuery()

	tmpl.Execute(w, struct{ Success bool }{true})
}
