package db

import "fmt"

// Goods represents a goods entity
type Goods struct {
	ID   string
	Name string
	Sort string
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
