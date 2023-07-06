package usecase

import "database/sql"

type RequirementUsecase interface {
	InserterRequirementGoods(DB *sql.DB) error
}

type PurchaseQuery interface {
	InsertPurchaseQuery(db *sql.DB) error
}

type PurchaseGoodsQuery interface {
	InsertPurchaseGoodsQuery(db *sql.DB) error
}
