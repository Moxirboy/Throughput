package dto

import "time"

type Client struct {
	Name string
	Date string
}
type Goods struct {
	ID   string
	Name string
	Sort string
}

type Purchase struct {
	Name   string
	Amount string
}

type PurchaseGoods struct {
	CortPrice string
}
type Requirement struct {
	NameClient string
	Date       time.Time
}
type RequirementGoods struct {
	GoodName string
	Amount   string
	CostCell string
}
