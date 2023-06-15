package handler

import (
	"html/template"
	"net/http"
	db3 "project/internal/db"
)

func Incoming(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/incoming.html"))

	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "incoming.html", nil)
		return
	}

	r.ParseForm()

	product := db3.Goods{
		Name: r.FormValue("name_of_purchase"),
		Sort: r.FormValue("sort"),
	}

	details := db3.Client{
		Name: r.FormValue("cname"),
		Date: r.FormValue("date"),
	}

	purchase := db3.Purchase{
		Name:   r.FormValue("name_of_purchase"),
		Amount: r.FormValue("amount"),
	}
	purchaseGoods := db3.PurchaseGoods{
		CortPrice: r.FormValue("cost"),
	}

	insertClientQuery := "INSERT INTO client (name, created) VALUES (?, ?)"
	_, err := db3.DB.Query(insertClientQuery, details.Name, details.Date)
	if err != nil {
		panic(err)
	}

	insertProductQuery := "INSERT INTO goods (name, sort) VALUES (?, ?)"
	_, err = db3.DB.Query(insertProductQuery, product.Name, product.Sort)
	if err != nil {
		panic(err)
	}

	clientID, err := db3.GetProductIDByName(db3.DB, "client", details.Name)
	if err != nil {
		panic(err)
	}

	insertPurchaseQuery := "INSERT INTO purchase (name, client_id) VALUES (?, ?)"
	_, err = db3.DB.Query(insertPurchaseQuery, product.Name, clientID)
	if err != nil {
		panic(err)
	}

	// Retrieve the product and purchase IDs
	productID, err := db3.GetProductIDByName(db3.DB, "goods", product.Name)
	if err != nil {
		panic(err)
	}

	purchaseID, err := db3.GetProductIDByName(db3.DB, "purchase", product.Name)
	if err != nil {
		panic(err)
	}

	insertPurchaseGoodsQuery := "INSERT INTO purchase_goods (goods_id, purchase_id, amount,cort_price) VALUES (?, ?, ?,?)"
	_, err = db3.DB.Query(insertPurchaseGoodsQuery, productID, purchaseID, purchase.Amount, purchaseGoods.CortPrice)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}
