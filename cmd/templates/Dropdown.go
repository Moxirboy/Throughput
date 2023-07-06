package templates

import (
	"project/internal/conn"
	"project/internal/service/repo/mysql"
)

func GenerateDropdownHTMLGoods() string {

	GoodsNames := mysql.GetName(conn.Db)
	dropdownHTML := "<select name='goods'>"
	for _, name := range GoodsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
func GenerateDropdownHTMLClient() string {
	ProductsNames := mysql.GetClientNames(conn.Db)
	dropdownHTML := "<select name='clients'>"
	for _, name := range ProductsNames {
		dropdownHTML += "<option value='" + name + "'>" + name + "</option>"
	}

	dropdownHTML += "</select>"
	return dropdownHTML
}
