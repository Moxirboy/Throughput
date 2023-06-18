package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project/internal/controller/v1/handler/adapter"
	"project/internal/service/repo/mysql"
	"project/templates"
)

var err error

func Outcoming(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/coming.html"))
	DropdownHTMLClient := templates.GenerateDropdownHTMLClient()
	DropdownHTMLGoods := templates.GenerateDropdownHTMLGoods()
	tmpl.Execute(w, map[string]interface{}{
		"DropdownHTMLClient": template.HTML(DropdownHTMLClient),
		"DropdownHTMLGoods":  template.HTML(DropdownHTMLGoods),
	})
	r.ParseForm()
	adapter.FormValue(r)
	if mysql.AmountCheck() >= adapter.RequirementGood.Amount {
		mysql.InserterRequirementGoods()
		tmpl.Execute(w, struct{ success bool }{true})
	} else {
		response := "not enough amount!"
		_, err := w.Write([]byte(response))
		if err != nil {
			// Handle the error if writing the response fails
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
