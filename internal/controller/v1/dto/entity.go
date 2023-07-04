package dto

import (
	"time"
)

type Requirements struct {
	NameClient string `form:"clients"`
	Date       time.Time
}
type RequirementGoods struct {
	GoodName string `form:"goods"`
	Amount   string `form:"amount"`
	CostCell string `form:"costCell"`
}
type Goods struct {
	Name string
	Sort string
}

type Client struct {
	Name string
	Date string
}

type Purchase struct {
	Name   string
	Amount any
}

type PurchaseGoods struct {
	CortPrice any
}
