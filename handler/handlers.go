package handler

import (
	"html/template"
	"net/http"
	db "project/database"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

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

	insertClientQuery := "INSERT INTO client (name, created) VALUES (?, ?)"
	_, err := db.DB.Query(insertClientQuery, details.Name, details.Date)
	if err != nil {
		panic(err)
	}

	insertProductQuery := "INSERT INTO goods (name, sort) VALUES (?, ?)"
	_, err = db.DB.Query(insertProductQuery, product.Name, product.Sort)
	if err != nil {
		panic(err)
	}

	clientID, err := getProductIDByName(db.DB, "client", details.Name)
	if err != nil {
		panic(err)
	}

	insertPurchaseQuery := "INSERT INTO purchase (name, client_id) VALUES (?, ?)"
	_, err = db.DB.Query(insertPurchaseQuery, product.Name, clientID)
	if err != nil {
		panic(err)
	}

	// Retrieve the product and purchase IDs
	productID, err := getProductIDByName(db.DB, "goods", product.Name)
	if err != nil {
		panic(err)
	}

	purchaseID, err := getProductIDByName(db.DB, "purchase", product.Name)
	if err != nil {
		panic(err)
	}

	insertPurchaseGoodsQuery := "INSERT INTO purchase_goods (goods_id, purchase_id, amount) VALUES (?, ?, ?)"
	_, err = db.DB.Query(insertPurchaseGoodsQuery, productID, purchaseID, purchase.Amount)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}
func HandleRoot1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index2.html"))
	dropdownHTML := db.GenerateDropdownHTML()
	tmpl.Execute(w, map[string]interface{}{
		"DropdownHTML": template.HTML(dropdownHTML),
	})

}
