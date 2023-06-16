package adapter

import (
	"net/http"
	"time"
)
import form "project/internal/controller/v1/dto"

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
