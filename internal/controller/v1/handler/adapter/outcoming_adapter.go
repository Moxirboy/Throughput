package adapter

import (
	"net/http"
	"time"
)
import "project/internal/controller/v1/dto"

var Requirements dto.Requirement
var RequirementGood dto.RequirementGoods

func FormValue(r *http.Request) {
	date := time.Now()
	Requirements = dto.Requirement{
		NameClient: r.FormValue("clients"),
		Date:       date,
	}
	RequirementGood = dto.RequirementGoods{
		GoodName: r.FormValue("goods"),
		Amount:   r.FormValue("amount"),
		CostCell: r.FormValue("cost"),
	}

}
