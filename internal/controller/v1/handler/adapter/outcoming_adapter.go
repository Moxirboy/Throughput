package adapter

import (
	"fmt"
	"net/http"
	"time"
)
import "project/internal/controller/v1/dto"

var Requirements dto.Requirements
var RequirementGoods dto.RequirementGoods
var RAdapter dto.RAdapter
var Submit string
var Date = time.Now()

func FormValue(r *http.Request) {
	date := time.Now()
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	Requirements = dto.Requirements{
		NameClient: r.FormValue("clients"),
		Date:       date,
	}
	RequirementGoods = dto.RequirementGoods{
		GoodName: r.FormValue("goods"),
		Amount:   r.FormValue("amount"),
		CostCell: r.FormValue("costCell"),
	}
	RAdapter = dto.RAdapter{
		Requirements:     Requirements,
		RequirementGoods: RequirementGoods,
	}
	Submit = r.FormValue("submit")
	if Submit != "" {
		fmt.Println(RequirementGoods, "\n", Requirements)
	}
}
func FormValidate(goods dto.RequirementGoods, requirement dto.Requirements) bool {
	if goods.GoodName == "" || goods.Amount == "" || goods.CostCell == "" || requirement.NameClient == "" {
		return false
	}
	return true
}
