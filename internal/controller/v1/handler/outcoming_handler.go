package handler

import (
	"fmt"
	"html/template"
	"net/http"
	db3 "project/internal/db"
	"time"
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
	date := time.Now()
	Requirement := db3.Requirement{
		Date: date,
		Name: r.FormValue("client"),
	}
	_ = Requirement
	RequirementGoods := db3.RequirementGoods{
		Product:  r.FormValue("goods"),
		Amount:   r.FormValue("amount"),
		CostCell: r.FormValue("cost"),
	}
	_ = RequirementGoods
	var amountCheck string
	var goodsId string
	err := db3.DB.QueryRow("select id from kirim.goods where name=?", RequirementGoods.Product).Scan(&goodsId)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()
	if err != nil {
		fmt.Println(113)
		panic(err)
	}
	err = db3.DB.QueryRow("select amount from purchase_goods where goods_id=?", goodsId).Scan(&amountCheck)
	if err != nil {
		fmt.Println(118)
		panic(err)
	}
	if amountCheck >= RequirementGoods.Amount {
		tmpl.Execute(w, struct{ success bool }{true})
		var clientId string
		err = db3.DB.QueryRow("select id from client where name=?", Requirement.Name).Scan(&clientId)
		if err != nil {
			fmt.Println(126)
			panic(err)
		}
		_, err = db3.DB.Query("insert into requirement (date,client_id) values(?,?)", Requirement.Date, clientId)
		if err != nil {
			fmt.Println(130)
			panic(err)
		}
		var RequirementId string
		err = db3.DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)
		if err != nil {
			fmt.Println(136)
			panic(err)
		}
		_, err = db3.DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", RequirementId, goodsId, RequirementGoods.Amount, RequirementGoods.CostCell)
		if err != nil {
			fmt.Println(141)
			panic(err)
		}
		tmpl.Execute(w, struct{ success bool }{true})
	} else {
		tmpl.Execute(w, struct{ success bool }{false})
	}
}
