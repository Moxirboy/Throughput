package mysql

import (
	"database/sql"
	"fmt"
	"project/internal/configs"
	"project/internal/controller/v1/handler/adapter"
)

var err error
var (
	DB, _ = configs.DB()
)

func GetProductIDByName(DB *sql.DB, table string, productName string) (int, error) {
	var productID int
	query := "SELECT id FROM " + table + " WHERE name = ? "
	err = DB.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		return 0, err
	}

	return productID, nil
}
func GetClientNames() []string {
	rows, err := DB.Query("select name from kirim.client;")
	if err != nil {
		fmt.Println(64)
		panic(err)
	}
	defer rows.Close()
	var ClientNames []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		ClientNames = append(ClientNames, name)
		if err != nil {
			fmt.Println(70)
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return ClientNames
}
func GoodsId() string {
	var goodsId string
	err = DB.QueryRow("select id from kirim.goods where name=?", adapter.RequirementGood.GoodName).Scan(&goodsId)
	if err != nil {
		panic(err)
	}
	return goodsId
}
func AmountCheck() (amountCheck string) {
	err = DB.QueryRow("select amount from purchase_goods where goods_id=?", GoodsId()).Scan(&amountCheck)
	if err != nil {
		panic(err)
	}
	return amountCheck
}
func InserterRequirementGoods() {
	var clientId string
	err = DB.QueryRow("select id from client where name=?", adapter.Requirements.NameClient).Scan(&clientId)
	if err != nil {

		panic(err)
	}
	_, err = DB.Query("insert into requirement (date,client_id) values(?,?)", adapter.Requirements.Date, clientId)
	if err != nil {

		panic(err)
	}
	var RequirementId string
	err = DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)
	if err != nil {
		panic(err)
	}
	_, err = DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", RequirementId, GoodsId(), adapter.RequirementGood.Amount, adapter.RequirementGood.CostCell)
	if err != nil {

		panic(err)
	}
}
