package adapter

import (
	"net/http"
)
import form "project/internal/controller/v1/dto"

var Product form.Goods
var DetailsClient form.Client
var Purchase form.Purchase
var PurchaseGoods form.PurchaseGoods

func FormValues(r *http.Request) {

	r.ParseForm()
	Product = form.Goods{
		Name: r.FormValue("name_of_purchase"),
		Sort: r.FormValue("sort"),
	}

	DetailsClient = form.Client{
		Name: r.FormValue("cname"),
		Date: r.FormValue("date"),
	}

	Purchase = form.Purchase{
		Name:   r.FormValue("name_of_purchase"),
		Amount: r.FormValue("amount"),
	}
	_ = Purchase
	PurchaseGoods = form.PurchaseGoods{
		CortPrice: r.FormValue("cost"),
	}

}
