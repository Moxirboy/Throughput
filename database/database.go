package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB, err = sql.Open("mysql", "root:moxirboy@tcp(127.0.0.1:3306)/kirim")
