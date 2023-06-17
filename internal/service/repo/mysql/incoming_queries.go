package mysql

import (
	"fmt"
	"project/internal/configs"
	"project/internal/controller/v1/handler/adapter"
)

var (
	DB, _ = configs.DB()
)

func GetGoodsNames() []string {

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

func InsertClientQuery() {
	insertClientQuery := "INSERT INTO client (name, created) VALUES (?, ?)"
	fmt.Println(adapter.DetailsClient.Name, adapter.DetailsClient.Date)
	var name string = adapter.DetailsClient.Name
	var date string = adapter.DetailsClient.Date
	query, err := DB.Query(insertClientQuery, name, date)
	defer query.Close()
	if err != nil {
		fmt.Println("InsertClientQuery")
		panic(err)
	}
}
func InsertProductQuery() {
	insertProductQuery := "INSERT INTO goods (name, sort) VALUES (?, ?)"
	_, err := DB.Query(insertProductQuery, adapter.Product.Name, adapter.Product.Sort)
	if err != nil {
		fmt.Println("InsertProductQuery")
		panic(err)
	}
}
func InsertPurchaseQuery() {
	clientID, err := GetProductIDByName(DB, "client", adapter.DetailsClient.Name)
	if err != nil {
		fmt.Println("clientId   InsertPurchaseQuery")
		panic(err)

	}

	insertPurchaseQuery := "INSERT INTO purchase (name, client_id) VALUES (?, ?)"
	_, err = DB.Query(insertPurchaseQuery, adapter.Product.Name, clientID)
	if err != nil {
		fmt.Println("from query   InsertPurchaseQuery")
		panic(err)
	}
}
func InsertPurchaseGoodsQuery() {
	productID, err := GetProductIDByName(DB, "goods", adapter.Product.Name)
	if err != nil {
		fmt.Println("productID   InsertPurchaseGoodsQuery")
		panic(err)
	}

	purchaseID, err := GetProductIDByName(DB, "purchase", adapter.Product.Name)
	if err != nil {
		fmt.Println("purchaseID   InsertPurchaseGoodsQuery")
		panic(err)
	}

	insertPurchaseGoodsQuery := "INSERT INTO purchase_goods (goods_id, purchase_id, amount,cort_price) VALUES (?, ?, ?,?)"
	_, err = DB.Query(insertPurchaseGoodsQuery, productID, purchaseID, adapter.Purchase.Amount, adapter.PurchaseGoods.CortPrice)
	if err != nil {
		fmt.Println("from query   InsertPurchaseGoodsQuery")
		panic(err)
	}
}
