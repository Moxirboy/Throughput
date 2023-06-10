package handler

import (
	"fmt"
	"html/template"
	"net/http"
	db "project/database"
	"time"
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
	purchaseGoods := db.PurchaseGoods{
		CortPrice: r.FormValue("cost"),
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

	insertPurchaseGoodsQuery := "INSERT INTO purchase_goods (goods_id, purchase_id, amount,cort_price) VALUES (?, ?, ?,?)"
	_, err = db.DB.Query(insertPurchaseGoodsQuery, productID, purchaseID, purchase.Amount, purchaseGoods.CortPrice)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}
func HandleRoot1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index1.html"))
	DropdownHTMLClient := db.GenerateDropdownHTMLClient()
	DropdownHTMLGoods := db.GenerateDropdownHTMLGoods()
	tmpl.Execute(w, map[string]interface{}{
		"DropdownHTMLClient": template.HTML(DropdownHTMLClient),
		"DropdownHTMLGoods":  template.HTML(DropdownHTMLGoods),
	})
	r.ParseForm()
	date := time.Now()
	Requirement := db.Requirement{
		Date: date,
		Client: db.Client{
			Name: r.FormValue("client"),
		},
	}
	_ = Requirement
	RequirementGoods := db.RequirementGoods{
		Product:  r.FormValue("goods"),
		Amount:   r.FormValue("amount"),
		CostCell: r.FormValue("cost"),
	}
	_ = RequirementGoods
	var amountCheck string
	var goodsId string
	err := db.DB.QueryRow("select id from kirim.goods where name=?", RequirementGoods.Product).Scan(&goodsId)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()
	if err != nil {
		fmt.Println(113)
		panic(err)
	}
	err = db.DB.QueryRow("select amount from purchase_goods where goods_id=?", goodsId).Scan(&amountCheck)
	if err != nil {
		fmt.Println(118)
		panic(err)
	}
	if amountCheck >= RequirementGoods.Amount {
		tmpl.Execute(w, struct{ success bool }{true})
		var clientId string
		err = db.DB.QueryRow("select id from client where name=?", Requirement.Name).Scan(&clientId)
		if err != nil {
			fmt.Println(126)
			panic(err)
		}
		_, err = db.DB.Query("insert into requirement (date,client_id) values(?,?)", Requirement.Date, clientId)
		if err != nil {
			fmt.Println(130)
			panic(err)
		}
		var RequirementId string
		err = db.DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)
		if err != nil {
			fmt.Println(136)
			panic(err)
		}
		_, err = db.DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", RequirementId, goodsId, RequirementGoods.Amount, RequirementGoods.CostCell)
		if err != nil {
			fmt.Println(141)
			panic(err)
		}
		tmpl.Execute(w, struct{ success bool }{true})
	} else {
		tmpl.Execute(w, struct{ success bool }{false})
	}
}
