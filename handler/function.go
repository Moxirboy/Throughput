package handler

import "database/sql"

func getProductIDByName(DB *sql.DB, table string, productName string) (int, error) {
	var productID int
	query := "SELECT id FROM " + table + " WHERE name = ? "

	// Execute the query and scan the result into the productID variable
	err := DB.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		return 0, err
	}

	return productID, nil
}
