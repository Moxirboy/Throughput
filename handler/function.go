package handler

import "database/sql"

func getProductIDByName(DB *sql.DB, table string, productName string) (int, error) {
	var productID int
	query := "SELECT id FROM " + table + " WHERE name = ? "
	err := DB.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		return 0, err
	}

	return productID, nil
}
