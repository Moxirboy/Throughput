package mysql

import (
	"database/sql"
	"fmt"
)

var err error

func (pq *ProductQueryImpl) GetId(DB *sql.DB) string {
	var goodsId string
	err = DB.QueryRow("select id from kirim.goods where name=?", pq.ProductName.GoodName).Scan(&goodsId)

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
func GetIDByName(DB *sql.DB, table string, productName string) (int, error) {
	var productID int
	query := "SELECT id FROM " + table + " WHERE name = ? "
	err = DB.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}
func GetNames(DB *sql.DB) []string {
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
func (pq *ProductQueryImpl) GetAmount(DB *sql.DB) (amountCheck string) {
	err = DB.QueryRow("select amount from purchase_goods where goods_id=?", pq.GetId(DB)).Scan(&amountCheck)

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
