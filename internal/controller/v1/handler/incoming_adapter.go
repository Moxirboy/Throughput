package handler

import "net/http"
import form "project/internal/controller/v1/dto"

var product *form.Goods
var details *form.Client
var purchase *form.Purchase
var purchaseGoods *form.PurchaseGoods

func FormValues(r *http.Request) {
	product := &form.Goods{
		Name: r.FormValue("name_of_purchase"),
		Sort: r.FormValue("sort"),
	}
	_ = product
	details := &form.Client{
		Name: r.FormValue("cname"),
		Date: r.FormValue("date"),
	}
	_ = details
	purchase := &form.Purchase{
		Name:   r.FormValue("name_of_purchase"),
		Amount: r.FormValue("amount"),
	}
	_ = purchase
	purchaseGoods := &form.PurchaseGoods{
		CortPrice: r.FormValue("cost"),
	}
	_ = purchaseGoods
}
