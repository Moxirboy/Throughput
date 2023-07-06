package usecase

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"project/internal/conn"
	"project/internal/controller/v1/handler/adapter"
	"project/internal/service/repo/mysql"
)

func IrenderForm(w http.ResponseWriter) {
	absPath, err := filepath.Abs("../templates/incoming.html")
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.ParseFiles(absPath))
	err = tmpl.Execute(w, nil)
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

		purchaseQuery := &PurchaseQueryImpl{
			Product:       adapter.Product,
			DetailsClient: adapter.Client,
			Purchase:      adapter.Purchase,
		}

		purchaseGoodsQuery := &PurchaseGoodsQueryImpl{
			Product:       adapter.Product,
			Purchase:      adapter.Purchase,
			PurchaseGoods: adapter.PurchaseGoods,
		}
		defer func() {
			if err := recover(); r != nil {
				fmt.Println("Recovered from panic:", err)
			}
		}()

		err := clientQuery.InsertClientQuery(conn.Db)
		if err != nil {
			fmt.Println(43)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = productQuery.InsertProductQuery(conn.Db)
		if err != nil {
			fmt.Println(50)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = purchaseQuery.InsertPurchaseQuery(conn.Db)
		if err != nil {
			fmt.Println(57)
			fmt.Println("Error inserting client:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = purchaseGoodsQuery.InsertPurchaseGoodsQuery(conn.Db)
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
