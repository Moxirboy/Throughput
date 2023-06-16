package db

import "fmt"

// Goods represents a goods entity

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