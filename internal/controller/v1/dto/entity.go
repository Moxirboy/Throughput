package handler

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
	Date   time.Time
	Client Client
	Name   any
}
type RequirementGoods struct {
	Product  string
	Amount   string
	CostCell string
}
