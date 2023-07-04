package conn

import (
	"database/sql"
	"fmt"
	"log"
	c "project/internal/configs"
)

var (
	Db, _ = DB()
)

func DB() (*sql.DB, error) {
	config, err := c.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	mysqlString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	fmt.Println("MySQL connection string:", mysqlString)

	conn, err := sql.Open("mysql", mysqlString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	return conn, nil
}
