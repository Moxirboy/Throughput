package handler

import (
	"net/http"
	"time"
)
import form "project/internal/controller/v1/dto"
import HTML "project/internal/service/repo"

var Requirement *form.Requirement
var RequirementGoods *form.RequirementGoods

func FormValue(r *http.Request) {
	date := time.Now()
	Requirement := &form.Requirement{
		Date: date,
		Name: r.FormValue("client"),
	}
	_ = Requirement
	RequirementGoods := &form.RequirementGoods{
		Product:  r.FormValue("goods"),
		Amount:   r.FormValue("amount"),
		CostCell: r.FormValue("cost"),
	}
	_ = RequirementGoods
}

func GenerateDropdownHTMLGoods() string {
	GoodsNames := HTML.GetGoodsNames()
	dropdownHTML := "<select name='goods'>"
	for _, name := range GoodsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
func GenerateDropdownHTMLProduct() string {
	ProductsNames := HTML.GetProductNames()
	dropdownHTML := "<select name='goods'>"
	for _, name := range ProductsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
