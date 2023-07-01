package usecase

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project/internal/conn"
	"project/internal/controller/v1/handler/adapter"
	"project/internal/service/repo/mysql"
	"project/templates"
)

func IrenderForm(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/incoming.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error rendering form:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func OrenderForm(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("C:/Users/User/Documents/GitHub/project/templates/outcoming.html"))
	DropdownHTMLClient := templates.GenerateDropdownHTMLClient()
	DropdownHTMLGoods := templates.GenerateDropdownHTMLGoods()
	err := tmpl.Execute(w, map[string]interface{}{
		"DropdownHTMLClient": template.HTML(DropdownHTMLClient),
		"DropdownHTMLGoods":  template.HTML(DropdownHTMLGoods),
	})
	//err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error rendering form:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func Income(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		adapter.FormValues(r)

		clientQuery := &mysql.ClientQueryImpl{
			DetailsClient: adapter.Client,
		}

		productQuery := &mysql.ProductQueryImpl{
			Product: adapter.Product,
		}

		purchaseQuery := &mysql.PurchaseQueryImpl{
			Product:       adapter.Product,
			DetailsClient: adapter.Client,
			Purchase:      adapter.Purchase,
		}

		purchaseGoodsQuery := &mysql.PurchaseGoodsQueryImpl{
			Product:       adapter.Product,
			Purchase:      adapter.Purchase,
			PurchaseGoods: adapter.PurchaseGoods,
		}
		defer func() {
			if err := recover(); r != nil {
				fmt.Println("Recovered from panic:", err)
			}
		}()

		err := clientQuery.InsertClientQuery(conn.DB)
		if err != nil {
			fmt.Println(43)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = productQuery.InsertProductQuery(conn.DB)
		if err != nil {
			fmt.Println(50)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = purchaseQuery.InsertPurchaseQuery(conn.DB)
		if err != nil {
			fmt.Println(57)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = purchaseGoodsQuery.InsertPurchaseGoodsQuery(conn.DB)
		if err != nil {
			fmt.Println(64)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println("Data inserted successfully.")
		http.Redirect(w, r, "/kirim", http.StatusSeeOther)
	} else {
		IrenderForm(w)
	}
}

func OutCome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		adapter.FormValue(r)
		RequirementQuery := &mysql.RequirementImpl{
			Requirement:     adapter.Requirements,
			RequirementGood: adapter.RequirementGoods,
		}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic occurred:", err)
			}
		}()
		if adapter.FormValidate(adapter.RequirementGoods, adapter.Requirements) {
			fmt.Println("validating")
			if mysql.AmountCheck(conn.DB) >= RequirementQuery.RequirementGood.Amount {
				err := RequirementQuery.InserterRequirementGoods(conn.DB)
				if err != nil {
					fmt.Println("Error inserting client:", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				fmt.Fprintln(w, "success")
			} else {
				fmt.Fprintln(w, "not enough amount!")
				response := "not enough amount!"
				_, err := w.Write([]byte(response))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

			}
		}
	} else {
		OrenderForm(w)
	}
}
