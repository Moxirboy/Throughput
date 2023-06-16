package handler

import (
	"fmt"
	"html/template"
	"net/http"
	db3 "project/internal/db"
)

func Outcoming(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/outcoming.html"))
	DropdownHTMLClient := db3.GenerateDropdownHTMLClient()
	DropdownHTMLGoods := db3.GenerateDropdownHTMLGoods()
	tmpl.Execute(w, map[string]interface{}{
		"DropdownHTMLClient": template.HTML(DropdownHTMLClient),
		"DropdownHTMLGoods":  template.HTML(DropdownHTMLGoods),
	})
	r.ParseForm()
	var amountCheck string
	var goodsId string
	err := db3.DB.QueryRow("select id from kirim.goods where name=?", RequirementGoods.Product).Scan(&goodsId)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()
	if err != nil {

		panic(err)
	}
	err = db3.DB.QueryRow("select amount from purchase_goods where goods_id=?", goodsId).Scan(&amountCheck)
	if err != nil {

		panic(err)
	}
	if amountCheck >= RequirementGoods.Amount {
		tmpl.Execute(w, struct{ success bool }{true})
		var clientId string
		err = db3.DB.QueryRow("select id from client where name=?", Requirement.Name).Scan(&clientId)
		if err != nil {

			panic(err)
		}
		_, err = db3.DB.Query("insert into requirement (date,client_id) values(?,?)", Requirement.Date, clientId)
		if err != nil {

			panic(err)
		}
		var RequirementId string
		err = db3.DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)
		if err != nil {
			panic(err)
		}
		_, err = db3.DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", RequirementId, goodsId, RequirementGoods.Amount, RequirementGoods.CostCell)
		if err != nil {

			panic(err)
		}
		tmpl.Execute(w, struct{ success bool }{true})
	} else {
		tmpl.Execute(w, struct{ success bool }{false})
	}
}
