package mysql

import (
	"database/sql"
	"fmt"
	config "project/internal/configs"
)

func GetProductIDByName(DB *sql.DB, table string, productName string) (int, error) {
	var productID int
	query := "SELECT id FROM " + table + " WHERE name = ? "
	err := DB.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		return 0, err
	}

	return productID, nil
}
func GetProductNames() []string {
	DB, err := config.DB()
	rows, err := DB.Query("select name from kirim.purchase;")
	if err != nil {
		fmt.Println(64)
		panic(err)
	}
	defer rows.Close()
	var ProductsNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		ProductsNames = append(ProductsNames, name)
		if err != nil {
			fmt.Println(70)
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return ProductsNames
}
