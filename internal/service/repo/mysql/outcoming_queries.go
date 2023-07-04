package mysql

import (
	"database/sql"
	"fmt"
	"project/internal/controller/v1/handler/adapter"
)

var err error

func GetProductIDByName(DB *sql.DB, table string, productName string) (int, error) {
	var productID int
	query := "SELECT id FROM " + table + " WHERE name = ? "
	err = DB.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}
func GetClientNames(DB *sql.DB) []string {
	rows, err := DB.Query("select name from kirim.client;")
	if err != nil {
		fmt.Println(64)
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
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
func GoodsId(DB *sql.DB) string {
	var goodsId string
	err = DB.QueryRow("select id from kirim.goods where name=?", adapter.RequirementGoods.GoodName).Scan(&goodsId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("Goods ID:", goodsId)
	}
	return goodsId
}
func AmountCheck(DB *sql.DB) (amountCheck string) {
	err = DB.QueryRow("select amount from purchase_goods where goods_id=?", GoodsId(DB)).Scan(&amountCheck)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("amountCheck:", amountCheck)
	}
	return amountCheck
}
func (r *RequirementImpl) InserterRequirementGoods(DB *sql.DB) error {
	var clientId string
	err := DB.QueryRow("select id from client where name=?", r.Requirement.NameClient).Scan(&clientId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("client ID:", clientId)
	}
	_, err = DB.Query("insert into requirement (date,client_id) values(?,?)", r.Requirement.Date, clientId)
	if err != nil {

		return err
	}
	var RequirementId string
	err = DB.QueryRow("select id from requirement where client_id=?", clientId).Scan(&RequirementId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Error occurred:", err)
		}
	} else {
		// Process the retrieved goodsId
		fmt.Println("Requirement ID:", RequirementId)
	}
	_, err = DB.Query("insert into requirement_goods (requirement_id,goods_id,amount,cost_cell) values(?,?,?,?)", RequirementId, GoodsId(DB), r.RequirementGood.Amount, r.RequirementGood.CostCell)
	if err != nil {
		return err
	}
	return nil
}
