package usecase

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"project/cmd/templates"
	"project/internal/conn"
	"project/internal/controller/v1/handler/adapter"
	"project/internal/service/repo/mysql"
)

func OrenderForm(w http.ResponseWriter) error {
	absPath, err := filepath.Abs("../templates/outcoming.html")
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.ParseFiles(absPath))
	DropdownHTMLClient := templates.GenerateDropdownHTMLClient()
	DropdownHTMLGoods := templates.GenerateDropdownHTMLGoods()
	err = tmpl.Execute(w, map[string]interface{}{
		"DropdownHTMLClient": template.HTML(DropdownHTMLClient),
		"DropdownHTMLGoods":  template.HTML(DropdownHTMLGoods),
	})
	//err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error rendering form:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}
	return nil
}
func OutCome(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from panic:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()
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
			if mysql.AmountCheck(conn.Db) >= RequirementQuery.RequirementGood.Amount {
				err := RequirementQuery.InserterRequirementGoods(conn.Db)
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
		err := OrenderForm(w)
		if err != nil {
			panic(err)
		}
	}
}
