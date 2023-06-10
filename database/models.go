package database

import (
	"fmt"
	"time"
)

type Client struct {
	Name string
	Date string
}

type Goods struct {
	ID   string
	Name string
	Sort string
}

type Purchase struct {
	Name   string
	Amount string
}
type PurchaseGoods struct {
	CortPrice string
}
type Requirement struct {
	Date time.Time
	Client
}
type RequirementGoods struct {
	Product  string
	Amount   string
	CostCell string
}

func GetClientNames() []string {
	rows, err := DB.Query("SELECT name FROM kirim.client;")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var clientNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil
		}
		clientNames = append(clientNames, name)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return clientNames
}
func GenerateDropdownHTMLClient() string {
	clientNames := GetClientNames()
	dropdownHTML := "<select name='client'>"
	for _, name := range clientNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
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
		//if err := rows.Scan(&name); err != nil {
		//	GoodsNames = append(GoodsNames, name)
		//}
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return GoodsNames
}
func GenerateDropdownHTMLGoods() string {
	GoodsNames := GetGoodsNames()
	dropdownHTML := "<select name='goods'>"
	for _, name := range GoodsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
