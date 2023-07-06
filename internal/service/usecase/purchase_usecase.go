package usecase

import (
	"database/sql"
	"fmt"
	"project/internal/controller/v1/dto"
	"project/internal/service/repo/mysql"
)

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

func (pq *PurchaseQueryImpl) InsertPurchaseQuery(db *sql.DB) error {
	clientID, err := mysql.GetIDByName(db, "client", pq.DetailsClient.Name)
	if err != nil {
		return err
	}

	insertPurchaseQuery := "INSERT INTO kirim.purchase (name, client_id) VALUES (?, ?)"
	_, err = db.Exec(insertPurchaseQuery, pq.Product.Name, clientID)
	if err != nil {
		return err
	}
	return nil
}

func (pgq *PurchaseGoodsQueryImpl) InsertPurchaseGoodsQuery(db *sql.DB) error {
	productID, err := mysql.GetIDByName(db, "goods", pgq.Product.Name)
	if err != nil {
		panic(err)
		return err
	}

	purchaseID, err := mysql.GetIDByName(db, "purchase", pgq.Product.Name)
	if err != nil {
		panic(err)
		return err
	}

	insertPurchaseGoodsQuery := "INSERT INTO kirim.purchase_goods (goods_id, purchase_id, amount, cort_price) VALUES (?, ?, ?, ?)"
	_, err = db.Query(insertPurchaseGoodsQuery, productID, purchaseID, pgq.Purchase.Amount, pgq.PurchaseGoods.CortPrice)
	if err != nil {
		fmt.Println(80, "     ", err)
		return err
	}
	return nil
}
