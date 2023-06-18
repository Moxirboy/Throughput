package configs

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

func DB() (*sql.DB, error) {
	err := godotenv.Load("C:/Users/User/Documents/GitHub/project/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	// Access the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	mysqlString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	instance, err := sql.Open("mysql", mysqlString)
	if err != nil {
		fmt.Println("dfgh")
		panic(err)
	}
	return instance, nil
}
