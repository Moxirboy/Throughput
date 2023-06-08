package handler

import (
	"html/template"
	"net/http"
	db "project/database"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	_ = tmpl
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	r.ParseForm()

	product := db.Goods{
		Name: r.FormValue("name_of_purchase"),
		Sort: r.FormValue("sort"),
	}
	details := db.Client{
		Name: r.FormValue("cname"),
		Date: r.FormValue("date"),
	}
	purchase := db.Purchase{
		Name:   r.FormValue("name_of_purchase"),
		Amount: r.FormValue("amount"),
	}
	insert, err := db.DB.Query("insert into client (name,created) values(?,?)", details.Name, details.Date)
	if err != nil {
		panic(err)
	}
	_, err = db.DB.Query("insert into goods (name,sort) values(?,?)", product.Name, product.Sort)
	if err != nil {
		panic(err)
	}
	clientID, err := getProductIDByName(db.DB, "client", details.Name)
	_, err = db.DB.Query("insert into purchase(name,client_id) values(?,?)", product.Name, clientID)
	if err != nil {
		panic(err)
	}
	productID, err := getProductIDByName(db.DB, "goods", product.Name)
	if err != nil {
		panic(err)
	}
	purchaseID, err := getProductIDByName(db.DB, "purchase", product.Name)
	if err != nil {
		panic(err)
	}
	_, err = db.DB.Query("insert into purchase_goods (goods_id,purchase_id,amount) values(?,?,?)", productID, purchaseID, purchase.Amount)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}
