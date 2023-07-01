package mysql

import (
	"database/sql"
	"fmt"
)

func GetGoodsNames(DB *sql.DB) []string {

	rows, err := DB.Query("select name from kirim.goods;")
	if err != nil {
		fmt.Println(64)
		panic(err)
	}
	defer rows.Close()
	var GoodsNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		GoodsNames = append(GoodsNames, name)
		if err != nil {
			fmt.Println(70)
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return GoodsNames
}

func (cq *ClientQueryImpl) InsertClientQuery(db *sql.DB) error {
	insertClientQuery := "INSERT INTO kirim.client (name) VALUES (?)"
	_, err := db.Exec(insertClientQuery, cq.DetailsClient.Name)
	if err != nil {
		return err
	}
	return nil
}

func (pq *ProductQueryImpl) InsertProductQuery(db *sql.DB) error {
	insertProductQuery := "INSERT INTO kirim.goods (name, sort) VALUES (?, ?)"
	_, err := db.Exec(insertProductQuery, pq.Product.Name, pq.Product.Sort)
	if err != nil {
		return err
	}
	return nil
}

func (pq *PurchaseQueryImpl) InsertPurchaseQuery(db *sql.DB) error {
	clientID, err := GetProductIDByName(db, "client", pq.DetailsClient.Name)
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
	productID, err := GetProductIDByName(db, "goods", pgq.Product.Name)
	if err != nil {
		panic(err)
		return err
	}

	purchaseID, err := GetProductIDByName(db, "purchase", pgq.Product.Name)
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
