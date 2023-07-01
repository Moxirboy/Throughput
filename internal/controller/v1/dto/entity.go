package dto

import (
	"database/sql"
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

type IAdapter struct {
	Product       Goods
	Client        Client
	Purchase      Purchase
	PurchaseGoods PurchaseGoods
}
type RAdapter struct {
	Requirements     Requirements
	RequirementGoods RequirementGoods
}
type RequirementUsecase interface {
	InserterRequirementGoods(DB *sql.DB) error
}
type ClientQuery interface {
	InsertClientQuery(db *sql.DB) error
}

type ProductQuery interface {
	InsertProductQuery(db *sql.DB) error
}

type PurchaseQuery interface {
	InsertPurchaseQuery(db *sql.DB) error
}

type PurchaseGoodsQuery interface {
	InsertPurchaseGoodsQuery(db *sql.DB) error
}
