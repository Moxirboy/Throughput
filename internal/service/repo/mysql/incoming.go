package mysql

import "project/internal/controller/v1/dto"

type ClientQueryImpl struct {
	DetailsClient dto.Client
}

type ProductQueryImpl struct {
	Product dto.Goods
}

type PurchaseQueryImpl struct {
	Product       dto.Goods
	DetailsClient dto.Client
	Purchase      dto.Purchase
}

type PurchaseGoodsQueryImpl struct {
	Product       dto.Goods
	Purchase      dto.Purchase
	PurchaseGoods dto.PurchaseGoods
}
