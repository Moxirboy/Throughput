package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project/internal/configs"
	"project/internal/controller/v1/handler/adapter"
	"project/templates"
)

var (
	DB, _ = configs.DB()
)

func Outcoming(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/coming.html"))
	DropdownHTMLClient := templates.GenerateDropdownHTMLProduct()
	DropdownHTMLGoods := templates.GenerateDropdownHTMLGoods()
	tmpl.Execute(w, map[string]interface{}{
		"DropdownHTMLClient": template.HTML(DropdownHTMLClient),
		"DropdownHTMLGoods":  template.HTML(DropdownHTMLGoods),
	})
	r.ParseForm()
	var amountCheck string
	var goodsId string
	err := DB.QueryRow("select id from kirim.goods where name=?", adapter.RequirementGoods.Product).Scan(&goodsId)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()
	if err != nil {

		panic(err)
	}
	err = DB.QueryRow("select amount from purchase_goods where goods_id=?", goodsId).Scan(&amountCheck)
	if err != nil {

		panic(err)
	}
	if amountCheck >= adapter.RequirementGoods.Amount {
		tmpl.Execute(w, struct{ success bool }{true})
		var clientId string
		err = DB.QueryRow("select id from client where name=?", adapter.Requirement.Name).Scan(&clientId)
		if err != nil {

			panic(err)
		}
		_, err = DB.Query("insert into requirement (date,client_id) values(?,?)", adapter.Requirement.Date, clientId)
		if err != nil {

			panic(err)
		}
		var RequirementId string
		err = DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)
		if err != nil {
			panic(err)
		}
		_, err = DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", "RequirementId, goodsId, RequirementGoods.Amount, RequirementGoods.CostCell")
		if err != nil {

			panic(err)
		}
		tmpl.Execute(w, struct{ success bool }{true})
	} else {
		tmpl.Execute(w, struct{ success bool }{false})
	}
}

//func InsertValues(values ...string) string {
//	var valuesString string
//	for _, values := range values {
//		valuesString += values
//	}
//	return valuesString
//}
//func Insert(
//	tableName string,
//	values string,
//	columnName ...string,
//) {
//	var column string
//	var questionMarks string
//	var numberOfValues int
//	var names string
//	for numberOfValues, names = range columnName {
//		column += names
//		column += ","
//	}
//	for i := 0; i < numberOfValues; i++ {
//		questionMarks += "?"
//		if i == numberOfValues-1 {
//
//		} else {
//			questionMarks += ","
//		}
//	}
//	for _, names := range columnName {
//		column += names
//		column += ","
//	}
//	var query = "insert into " + tableName + "(" + column + ") values(" + questionMarks + ")"
//	_ = query
//	_, err = DB.Query(query, values)
//}
