package templates

import (
	"database/sql"
	"fmt"
	"project/internal/conn"
)

func GenerateDropdownHTMLGoods() string {

	GoodsNames := GetName(conn.Db)
	dropdownHTML := "<select name='goods'>"
	for _, name := range GoodsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
func GenerateDropdownHTMLClient() string {
	ClientsNames := GetNames(conn.Db)
	dropdownHTML := "<select name='clients'>"
	for _, name := range ClientsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}

func GetName(DB *sql.DB) []string {

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

func GetNames(DB *sql.DB) []string {

	rows, err := DB.Query("select name from kirim.client;")
	if err != nil {
		fmt.Println(64)
		panic(err)
	}
	defer rows.Close()
	var clientNames []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		clientNames = append(clientNames, name)
		if err != nil {
			fmt.Println(70)
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return clientNames
}
