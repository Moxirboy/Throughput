package mysql

import (
	"database/sql"
	"fmt"
	"project/internal/controller/v1/dto"
)

type ClientQueryImpl struct {
	DetailsClient dto.Client
}

type ProductQueryImpl struct {
	Product     dto.Goods
	ProductName dto.RequirementGoods
}

func (pq *ProductQueryImpl) GetName(DB *sql.DB) []string {

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
